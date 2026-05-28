import fs from "node:fs";
import path from "node:path";
import { parse as parseYaml } from "yaml";
import { AgentConfigSchema, type AgentConfig } from "./types.js";

const MAX_WALK_DEPTH = 5;
const AGENT_CONFIG_FILENAME = "testflowkit.agent.yml";
const AGENT_CONFIG_ENV = "TESTFLOWKIT_AGENT_CONFIG";

export interface LoadedAgentConfig {
  config: AgentConfig;
  configDir: string;
}

function walkForConfig(startDir: string): string | null {
  let dir = startDir;
  for (let depth = 0; depth < MAX_WALK_DEPTH; depth++) {
    const candidate = path.join(dir, AGENT_CONFIG_FILENAME);
    if (fs.existsSync(candidate)) {
      return candidate;
    }
    const parent = path.dirname(dir);
    if (parent === dir) break;
    dir = parent;
  }
  return null;
}

function interpolateEnv(value: string): string {
  return value.replace(
    /\{\{\s*env\.([A-Z0-9_]+)\s*\}\}/g,
    (_, name: string) => {
      return process.env[name] ?? "";
    },
  );
}

function interpolateConfigEnv(config: AgentConfig): AgentConfig {
  if (config.step_catalog.url) {
    config.step_catalog.url = interpolateEnv(config.step_catalog.url);
  }
  return config;
}

export function loadAgentConfig(
  cwd: string = process.cwd(),
): LoadedAgentConfig {
  const configPath = resolveConfigPath(cwd);

  const raw = fs.readFileSync(configPath, "utf-8");
  const parsed = parseYaml(raw) as unknown;
  const result = AgentConfigSchema.safeParse(parsed);

  if (!result.success) {
    const issues = result.error.issues
      .map((i) => `  - ${i.path.join(".")}: ${i.message}`)
      .join("\n");
    throw new Error(
      `Invalid "testflowkit.agent.yml" at "${configPath}":\n${issues}`,
    );
  }

  const config = interpolateConfigEnv(result.data);
  const configDir = path.dirname(configPath);
  return { config, configDir };
}

function resolveConfigPath(cwd: string): string {
  const envOverride = process.env[AGENT_CONFIG_ENV];
  if (envOverride) {
    const configPath = path.isAbsolute(envOverride)
      ? envOverride
      : path.resolve(cwd, envOverride);
    if (!fs.existsSync(configPath)) {
      throw new Error(
        `TESTFLOWKIT_AGENT_CONFIG points to "${configPath}" but the file does not exist.`,
      );
    }
    return configPath;
  }
  const found = walkForConfig(cwd);
  if (!found) {
    throw new Error(
      `Could not find "${AGENT_CONFIG_FILENAME}" in "${cwd}" or any of its ${MAX_WALK_DEPTH} parent directories.\n` +
        `Create a "testflowkit.agent.yml" at your project root, or set TESTFLOWKIT_AGENT_CONFIG to its path.`,
    );
  }
  return found;
}
