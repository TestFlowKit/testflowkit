import type { AgentConfig } from "../config/types.js";
import {
  CLI_VERSION_RE,
  toSemverBase,
  isCanary,
  getCliVersion,
} from "../catalog/version.js";

export type CliVersionSource =
  | "session"
  | "env"
  | "config"
  | "probe"
  | "override";

export type CatalogSession = {
  cliVersion: string | null;
  source: CliVersionSource | null;
  initializedAt: string | null;
};

export function createCatalogSession(config: AgentConfig): CatalogSession {
  // Priority 3: TESTFLOWKIT_CLI_VERSION env var
  const envVersion = process.env["TESTFLOWKIT_CLI_VERSION"];
  if (envVersion && CLI_VERSION_RE.test(envVersion)) {
    return {
      cliVersion: envVersion,
      source: "env",
      initializedAt: new Date().toISOString(),
    };
  }

  // Priority 4: step_catalog.cli_version in testflowkit.agent.yml
  const configVersion = config.step_catalog.cli_version;
  if (configVersion && CLI_VERSION_RE.test(configVersion)) {
    return {
      cliVersion: configVersion,
      source: "config",
      initializedAt: new Date().toISOString(),
    };
  }

  // Priority 5: probe (last resort, may fail in sandboxed envs)
  const probed = getCliVersion();
  if (probed) {
    return {
      cliVersion: probed,
      source: "probe",
      initializedAt: new Date().toISOString(),
    };
  }

  return {
    cliVersion: null,
    source: null,
    initializedAt: null,
  };
}

export function setSessionCliVersion(
  session: CatalogSession,
  version: string,
): void {
  if (!CLI_VERSION_RE.test(version)) {
    throw new Error(
      `Invalid cliVersion "${version}". Expected semver like "1.4.2" or "3.6.1-canary.abc1234".`,
    );
  }
  session.cliVersion = version;
  session.source = "session";
  session.initializedAt = new Date().toISOString();
}

export function resolveCliVersion(
  session: CatalogSession,
  _config: AgentConfig,
  override?: string,
): { version: string | null; source: CliVersionSource | null } {
  // Priority 1: explicit per-tool override
  if (override && CLI_VERSION_RE.test(override)) {
    return { version: override, source: "override" };
  }

  // Priority 2–5: session (already holds env / config / probe at boot, or agent-set)
  if (session.cliVersion) {
    return { version: session.cliVersion, source: session.source };
  }

  return { version: null, source: null };
}

export { toSemverBase, isCanary };
