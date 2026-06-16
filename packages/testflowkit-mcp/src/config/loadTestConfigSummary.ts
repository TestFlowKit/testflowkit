import fs from "node:fs";
import path from "node:path";
import { parse as parseYaml } from "yaml";

const LEGACY_CONFIG = "config.yml";
const DEFAULT_CONFIG = "testflowkit.yml";

const SECRET_FIELDS = new Set([
  "token",
  "client_secret",
  "password",
  "api_key",
  "secret",
  "credential",
]);

function redactSecrets(obj: unknown, depth = 0): unknown {
  if (depth > 10 || obj === null || typeof obj !== "object") return obj;
  if (Array.isArray(obj)) {
    return obj.map((v) => redactSecrets(v, depth + 1));
  }
  const result: Record<string, unknown> = {};
  for (const [k, v] of Object.entries(obj as Record<string, unknown>)) {
    if (SECRET_FIELDS.has(k.toLowerCase())) {
      result[k] = "***";
    } else {
      result[k] = redactSecrets(v, depth + 1);
    }
  }
  return result;
}

function stripEnvSection(
  raw: Record<string, unknown>,
): Record<string, unknown> {
  const { env: _env, ...rest } = raw;
  return rest;
}

export type TestConfigSummary = {
  apis: ApiSummary[];
  frontend: FrontendSummary | null;
};

export type ApiSummary = {
  name: string;
  type: string;
  operationsOrEndpoints: string[];
};

export type FrontendSummary = {
  pages: string[];
  elementGroups: string[];
};

export function loadTestConfigSummary(
  testConfigPath: string,
  configDir: string,
): TestConfigSummary {
  let resolved = path.isAbsolute(testConfigPath)
    ? testConfigPath
    : path.resolve(configDir, testConfigPath);

  if (!fs.existsSync(resolved)) {
    const legacy = path.join(path.dirname(resolved), LEGACY_CONFIG);
    if (fs.existsSync(legacy)) {
      resolved = legacy;
    } else {
      const def = path.join(path.dirname(resolved), DEFAULT_CONFIG);
      if (fs.existsSync(def)) resolved = def;
    }
  }

  if (!fs.existsSync(resolved)) {
    return { apis: [], frontend: null };
  }

  const raw = parseYaml(fs.readFileSync(resolved, "utf-8")) as Record<
    string,
    unknown
  >;
  const stripped = stripEnvSection(raw);
  const safe = redactSecrets(stripped) as Record<string, unknown>;

  const apis = extractApis(safe);
  const frontend = extractFrontend(safe);
  return { apis, frontend };
}

function extractApis(cfg: Record<string, unknown>): ApiSummary[] {
  const apisBlock = cfg["apis"] as
    | { definitions?: Record<string, unknown> }
    | undefined;
  if (!apisBlock?.definitions) return [];

  return Object.entries(apisBlock.definitions).map(([name, def]) => {
    const d = def as Record<string, unknown>;
    const type = (d["type"] as string) ?? "unknown";

    let operationsOrEndpoints: string[] = [];
    if (d["operations"]) {
      operationsOrEndpoints = Object.keys(d["operations"] as object).map(
        (op) => `${name}.${op}`,
      );
    } else if (d["endpoints"]) {
      operationsOrEndpoints = Object.keys(d["endpoints"] as object).map(
        (ep) => `${name}.${ep}`,
      );
    }

    return { name, type, operationsOrEndpoints };
  });
}

function extractFrontend(cfg: Record<string, unknown>): FrontendSummary | null {
  const fe = cfg["frontend"] as Record<string, unknown> | undefined;
  if (!fe) return null;

  const pages = fe["pages"] ? Object.keys(fe["pages"] as object) : [];

  const elementGroups = fe["elements"]
    ? Object.keys(fe["elements"] as object)
    : [];

  return { pages, elementGroups };
}
