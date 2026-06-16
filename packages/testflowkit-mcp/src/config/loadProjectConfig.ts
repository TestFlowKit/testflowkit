import fs from "node:fs";
import path from "node:path";
import { parse as parseYaml } from "yaml";
import { Logger } from "../logger.js";
import {
  ProjectConfigSchema,
  toFeaturesGlob,
  type AgentConfig,
  type ProjectConfig,
} from "./types.js";

const PRIMARY_CONFIG_FILENAME = "testflowkit.yml";
const LEGACY_CONFIG_FILENAME = "config.yml";

export interface LoadedProjectConfig {
  config: ProjectConfig;
  configDir: string;
}

export function loadProjectConfig(
  cwd: string = process.cwd(),
): LoadedProjectConfig {
  const configPath = resolveConfigPath(cwd);
  const raw = fs.readFileSync(configPath, "utf-8");
  const parsed = parseYaml(raw) as unknown;
  const result = ProjectConfigSchema.safeParse(parsed);

  if (!result.success) {
    const issues = result.error.issues
      .map((i) => `  - ${i.path.join(".")}: ${i.message}`)
      .join("\n");
    throw new Error(
      `Invalid "${path.basename(configPath)}" at "${configPath}":\n${issues}`,
    );
  }

  const configDir = path.dirname(configPath);

  const rawAgent = result.data.agent;
  let agent: AgentConfig | undefined;
  if (rawAgent) {
    agent = {
      defaultTagsForDraft: rawAgent.default_tags_for_draft,
      runCommand: rawAgent.run_command,
      stepCatalog: rawAgent.step_catalog,
    };
    if (rawAgent.step_catalog?.file || rawAgent.step_catalog?.url) {
      Logger.warn(
        "agent.step_catalog.file / url are not yet implemented — the MCP server will use the local tkit CLI export instead.",
      );
    }
  }

  const config: ProjectConfig = {
    configPath,
    featuresGlob: toFeaturesGlob(result.data.settings.gherkin_location),
    agent,
  };

  return { config, configDir };
}

function resolveConfigPath(cwd: string): string {
  const primary = path.join(cwd, PRIMARY_CONFIG_FILENAME);
  if (fs.existsSync(primary)) {
    return primary;
  }

  const legacy = path.join(cwd, LEGACY_CONFIG_FILENAME);
  if (fs.existsSync(legacy)) {
    return legacy;
  }

  throw new Error(
    `Could not find "${PRIMARY_CONFIG_FILENAME}" in "${cwd}".\n` +
      `Create a "${PRIMARY_CONFIG_FILENAME}" in the working directory.`,
  );
}
