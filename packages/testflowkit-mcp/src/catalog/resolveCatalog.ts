import type { ResolvedCatalog, StepDefinition } from "../config/types.js";
import { getCliVersion, resolveCliBinary } from "./version.js";
import { fetchFromLocalCli } from "./fetchLocalCli.js";

function parseCatalog(json: string): StepDefinition[] {
  const data = JSON.parse(json) as unknown;
  if (!Array.isArray(data)) {
    throw new TypeError("step-definitions.json must be a JSON array");
  }
  return data as StepDefinition[];
}

export async function resolveCatalog(
  configDir: string,
  opts: { cliBinary?: string } = {},
): Promise<ResolvedCatalog> {
  const warnings: string[] = [];
  const cliVersion = getCliVersion();
  const binary = opts.cliBinary ?? resolveCliBinary();
  const content = fetchFromLocalCli(binary, configDir);
  const steps = parseCatalog(content);

  return {
    steps,
    meta: {
      source: "local",
      cliVersion,
      warnings,
    },
  };
}
