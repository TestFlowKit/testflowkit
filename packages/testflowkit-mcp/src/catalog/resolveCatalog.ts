import fs from "node:fs";
import path from "node:path";
import type {
  AgentConfig,
  ResolvedCatalog,
  StepDefinition,
} from "../config/types.js";
import {
  getCliVersion,
  resolveCliBinary,
  toSemverBase,
  isCanary,
  hasGitDescribeSuffix,
  stripGitDescribeSuffix,
  versionMismatchLevel,
} from "./version.js";
import { fetchReleaseAsset, fetchUrl } from "./fetchRelease.js";
import { fetchFromLocalCli } from "./fetchLocalCli.js";

const CACHE_DIR = ".testflowkit/cache";
const FALLBACK_CACHE_FILENAME = "step-definitions.json";
const RELEASE_REPOSITORY = "TestFlowKit/testflowkit";
const RELEASE_ASSET = "step-definitions.json";

function buildCachePath(configDir: string, cliVersion: string | null): string {
  const cacheDir = path.resolve(configDir, CACHE_DIR);
  if (!cliVersion) {
    return path.join(cacheDir, FALLBACK_CACHE_FILENAME);
  }

  const baseVersion = toSemverBase(cliVersion);
  return path.join(cacheDir, `step-definitions-${baseVersion}.json`);
}

function parseCatalog(json: string): StepDefinition[] {
  const data = JSON.parse(json) as unknown;
  if (!Array.isArray(data)) {
    throw new TypeError("step-definitions.json must be a JSON array");
  }
  return data as StepDefinition[];
}

function readCache(cachePath: string): StepDefinition[] | null {
  try {
    const content = fs.readFileSync(cachePath, "utf-8");
    return parseCatalog(content);
  } catch {
    return null;
  }
}

function writeCache(cachePath: string, content: string): void {
  try {
    fs.mkdirSync(path.dirname(cachePath), { recursive: true });
    fs.writeFileSync(cachePath, content, "utf-8");
  } catch {
    // non-fatal: cache write failure should not block the tool
  }
}

export async function resolveCatalog(
  config: AgentConfig,
  configDir: string,
  opts: {
    forceRefresh?: boolean;
    cliVersion?: string | null;
    cliBinary?: string;
  } = {},
): Promise<ResolvedCatalog> {
  const forceRefresh = opts.forceRefresh ?? false;
  const cliVersionOverride = opts.cliVersion ?? undefined;
  const warnings: string[] = [];

  // 1. Local file
  if (config.step_catalog.file) {
    const filePath = path.isAbsolute(config.step_catalog.file)
      ? config.step_catalog.file
      : path.resolve(configDir, config.step_catalog.file);

    const content = fs.readFileSync(filePath, "utf-8");
    const steps = parseCatalog(content);
    return {
      steps,
      meta: {
        source: "file",
        catalogVersion: null,
        cliVersion: null,
        warnings,
      },
    };
  }

  // 2. Explicit URL
  if (config.step_catalog.url) {
    const cachePath = buildCachePath(configDir, null);
    const content = await fetchUrl(config.step_catalog.url);
    const steps = parseCatalog(content);
    writeCache(cachePath, content);
    return {
      steps,
      meta: { source: "url", catalogVersion: null, cliVersion: null, warnings },
    };
  }

  // 3. Release fetch (by CLI version)
  const cliVersion = cliVersionOverride ?? getCliVersion();
  const cachePath = buildCachePath(configDir, cliVersion);

  if (!cliVersion) {
    // Fall back to cache if CLI not found
    if (!forceRefresh) {
      const cached = readCache(cachePath);
      if (cached) {
        warnings.push(
          "Could not determine installed tkit version. Using cached catalog.",
        );
        return {
          steps: cached,
          meta: {
            source: "cache",
            catalogVersion: null,
            cliVersion: null,
            warnings,
          },
        };
      }
    }
    throw new Error(
      "Could not determine the installed tkit version and no cache exists.\n" +
        "Install the tkit CLI or specify step_catalog.file in testflowkit.agent.yml.",
    );
  }

  const baseVersion = toSemverBase(cliVersion);

  if (isCanary(cliVersion)) {
    warnings.push(
      `Canary CLI "${cliVersion}" detected — fetching catalog for semver base "${baseVersion}".`,
    );
  }

  // Check cache first (shared by local CLI and GitHub release)
  if (!forceRefresh) {
    const cached = readCache(cachePath);
    if (cached) {
      return {
        steps: cached,
        meta: {
          source: "cache",
          catalogVersion: baseVersion,
          cliVersion,
          warnings,
        },
      };
    }
  }

  // Try local CLI before fetching from GitHub release
  try {
    const binary = opts.cliBinary ?? resolveCliBinary();
    const content = fetchFromLocalCli(binary, configDir);
    const steps = parseCatalog(content);
    writeCache(cachePath, content);
    return {
      steps,
      meta: {
        source: "local",
        catalogVersion: baseVersion,
        cliVersion,
        warnings,
      },
    };
  } catch {
    // Local CLI not available or failed — fall through to GitHub release
  }

  // Fetch from GitHub Release.
  // For versions with a git-describe suffix (e.g. 3.7.0-canary.fed5ed9-2-gba72690),
  // try the full CLI version tag first, then fall back to the canary base tag
  // (e.g. 3.7.0-canary.fed5ed9) before giving up.
  let content: string;
  if (hasGitDescribeSuffix(cliVersion)) {
    const fallbackVersion = stripGitDescribeSuffix(cliVersion);
    try {
      content = await fetchReleaseAsset(
        RELEASE_REPOSITORY,
        cliVersion,
        RELEASE_ASSET,
      );
    } catch {
      warnings.push(
        `Release not found for "${cliVersion}" — retrying with "${fallbackVersion}".`,
      );
      content = await fetchReleaseAsset(
        RELEASE_REPOSITORY,
        fallbackVersion,
        RELEASE_ASSET,
      );
    }
  } else {
    content = await fetchReleaseAsset(
      RELEASE_REPOSITORY,
      baseVersion,
      RELEASE_ASSET,
    );
  }
  const steps = parseCatalog(content);

  const level = versionMismatchLevel(cliVersion, baseVersion);
  if (level === "patch") {
    warnings.push(
      `Patch version difference between CLI "${cliVersion}" and catalog "${baseVersion}".`,
    );
  }

  writeCache(cachePath, content);

  return {
    steps,
    meta: {
      source: "release",
      catalogVersion: baseVersion,
      cliVersion,
      warnings,
    },
  };
}
