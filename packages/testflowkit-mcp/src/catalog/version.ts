import { execSync } from "node:child_process";
import { Logger } from "../logger.js";

const CANARY_SUFFIX_RE = /-canary\.[a-f0-9]+$/;
const GIT_DESCRIBE_SUFFIX_RE = /-\d+-g[a-f0-9]+$/i;
// Matches stable, canary (3.7.0-canary.fed5ed9), and git-describe canary (3.7.0-canary.fed5ed9-2-gba72690)
const VERSION_RE =
  /(\d+\.\d+\.\d+(?:-canary\.[a-f0-9]+(?:-\d+-g[a-f0-9]+)?)?)/i;
export const CLI_VERSION_RE =
  /\d+\.\d+\.\d+(?:-canary\.[a-f0-9]+(?:-\d+-g[a-f0-9]+)?)?/i;

export type CliVersionProbeAttempt = {
  command: string;
  output: string | null;
  matchedVersion: string | null;
  error: string | null;
};

export type CliVersionProbeResult = {
  version: string | null;
  attempts: CliVersionProbeAttempt[];
};

export function probeCliVersion(): CliVersionProbeResult {
  const cmds = ["tkit version", "npx --yes @testflowkit/cli version"];
  const attempts: CliVersionProbeAttempt[] = [];

  for (const cmd of cmds) {
    try {
      const output = execSync(cmd, {
        encoding: "utf-8",
        timeout: 10_000,
        stdio: ["ignore", "pipe", "ignore"],
      }).trim();

      Logger.debug(`[catalog] getCliVersion output from "${cmd}": ${output}`);

      const match = VERSION_RE.exec(output);
      const matchedVersion = match ? match[1] : null;
      attempts.push({
        command: cmd,
        output,
        matchedVersion,
        error: null,
      });

      if (matchedVersion) {
        return {
          version: matchedVersion,
          attempts,
        };
      }

      Logger.debug(`[catalog] getCliVersion no semver match for "${cmd}"`);
    } catch (error) {
      const message = error instanceof Error ? error.message : String(error);
      attempts.push({
        command: cmd,
        output: null,
        matchedVersion: null,
        error: message,
      });
      Logger.debug(`[catalog] getCliVersion failed for "${cmd}": ${message}`);
    }
  }

  Logger.debug("[catalog] getCliVersion could not resolve CLI version");
  return {
    version: null,
    attempts,
  };
}

export function getCliVersion(): string | null {
  return probeCliVersion().version;
}

export function resolveCliBinary(): string {
  return process.env.TESTFLOWKIT_CLI_PATH ?? "tkit";
}

export function toSemverBase(version: string): string {
  // Strip git-describe suffix first (e.g. -2-gba72690), then canary suffix
  return version
    .replace(GIT_DESCRIBE_SUFFIX_RE, "")
    .replace(CANARY_SUFFIX_RE, "");
}

export function hasGitDescribeSuffix(version: string): boolean {
  return GIT_DESCRIBE_SUFFIX_RE.test(version);
}

export function stripGitDescribeSuffix(version: string): string {
  return version.replace(GIT_DESCRIBE_SUFFIX_RE, "");
}

export function isCanary(version: string): boolean {
  return version.includes("-canary.");
}

export function versionMismatchLevel(
  cliVersion: string,
  catalogVersion: string,
): "ok" | "patch" | "minor" | "major" {
  const cli = toSemverBase(cliVersion).split(".").map(Number);
  const cat = catalogVersion.split(".").map(Number);

  if (cli[0] !== cat[0]) return "major";
  if (cli[1] !== cat[1]) return "minor";
  if (cli[2] !== cat[2]) return "patch";
  return "ok";
}
