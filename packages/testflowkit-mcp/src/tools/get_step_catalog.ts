import { z } from "zod";
import { resolveCatalog } from "../catalog/resolveCatalog.js";
import { HandlerParams, TkitTool } from "./tool.js";

const inputSchema = z.object({
  forceRefresh: z.boolean().optional().default(false),
  cliVersion: z
    .string()
    .optional()
    .describe(
      'Installed tkit CLI version (e.g. "1.4.2"). Obtain it by running get_cli_version_commands first.',
    ),
});

export class GetStepCatalogTool implements TkitTool<typeof inputSchema> {
  getName(): string {
    return "get_step_catalog";
  }

  getDescription(): string {
    return (
      "Retrieve the TestFlowKit step definitions catalog. Returns all registered Gherkin sentence patterns with descriptions, categories, examples, and variables. " +
      "Optionally pass cliVersion (obtained from get_cli_version_commands) to skip server-side CLI detection."
    );
  }

  getInputSchema() {
    return inputSchema;
  }

  async handler(params: HandlerParams<z.infer<typeof inputSchema>>) {
    const { config, configDir, input } = params;
    const { forceRefresh, cliVersion } = input;

    const resolved = await resolveCatalog(
      config,
      configDir,
      forceRefresh,
      cliVersion,
    );

    return {
      content: [
        {
          type: "text" as const,
          text: JSON.stringify(
            {
              meta: resolved.meta,
              count: resolved.steps.length,
              steps: resolved.steps,
            },
            null,
            2,
          ),
        },
      ],
    };
  }
}
