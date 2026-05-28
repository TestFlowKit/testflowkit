import { z } from "zod";
import { resolveCatalog } from "../../catalog/resolveCatalog.js";
import { resolveCliVersion } from "../../context/session.js";
import type { RegisterTool } from "./types.js";

export const registerGetStepCatalogTool: RegisterTool = (
  server,
  config,
  configDir,
  session,
) => {
  server.registerTool(
    "get_step_catalog",
    {
      description:
        "Retrieve the TestFlowKit step definitions catalog. Returns registered Gherkin sentence patterns with descriptions, categories, examples, and variables. " +
        "Use optional category to narrow results to one category. " +
        "Call initialize_session first if the tkit CLI is not available on the server's PATH.",
      inputSchema: z.object({
        forceRefresh: z.boolean().optional().default(false),
        category: z
          .string()
          .optional()
          .describe(
            "Optional category filter (exact match, case-insensitive), for example: frontend, restapi, graphql",
          ),
        cliVersion: z
          .string()
          .optional()
          .describe(
            'Override the tkit CLI version for this call only (e.g. "1.4.2"). ' +
              "Prefer calling initialize_session once instead of passing this on every call.",
          ),
      }),
    },
    async ({ forceRefresh, category, cliVersion: cliVersionOverride }) => {
      const { version: cliVersion } = resolveCliVersion(
        session,
        config,
        cliVersionOverride,
      );

      if (
        !cliVersion &&
        !config.step_catalog.file &&
        !config.step_catalog.url
      ) {
        return {
          content: [
            {
              type: "text" as const,
              text: JSON.stringify(
                {
                  error: "cli_version_required",
                  message:
                    "Run `tkit version` in the project terminal, then call initialize_session.",
                  hints: [
                    'initialize_session({ cliVersion: "x.y.z" })',
                    "or set TESTFLOWKIT_CLI_VERSION in mcp.json env",
                    "or set step_catalog.cli_version in testflowkit.agent.yml",
                    "or set step_catalog.file for offline/canary use",
                  ],
                },
                null,
                2,
              ),
            },
          ],
          isError: true,
        };
      }

      const resolved = await resolveCatalog(config, configDir, {
        forceRefresh,
        cliVersion,
      });
      const normalizedCategory = category?.trim().toLowerCase();
      const steps = normalizedCategory
        ? resolved.steps.filter((step) =>
            step.categories.some((c) => c.toLowerCase() === normalizedCategory),
          )
        : resolved.steps;

      return {
        content: [
          {
            type: "text",
            text: JSON.stringify(
              {
                meta: resolved.meta,
                category: category ?? null,
                count: steps.length,
                steps,
              },
              null,
              2,
            ),
          },
        ],
      };
    },
  );
};
