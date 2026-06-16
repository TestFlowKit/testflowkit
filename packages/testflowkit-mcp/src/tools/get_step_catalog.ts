import { z } from "zod";
import { filterStepsByCategory } from "../catalog/categories.js";
import { resolveCatalog } from "../catalog/resolveCatalog.js";
import { HandlerParams, TkitTool } from "./tool.js";

const inputSchema = z.object({
  category: z
    .string()
    .optional()
    .describe(
      "Filter steps to this category. Use get_step_categories to list available values.",
    ),
});

export class GetStepCatalogTool implements TkitTool<typeof inputSchema> {
  getName(): string {
    return "get_step_catalog";
  }

  getDescription(): string {
    return (
      "Retrieve the TestFlowKit step definitions catalog via the installed tkit CLI. Returns registered Gherkin sentence patterns with descriptions, categories, examples, and variables. " +
      "Optionally pass category (from get_step_categories) to return only steps in that category."
    );
  }

  getInputSchema() {
    return inputSchema;
  }

  async handler(params: HandlerParams<z.infer<typeof inputSchema>>) {
    const { configDir, input } = params;
    const { category } = input;

    const resolved = await resolveCatalog(configDir);
    const steps = category
      ? filterStepsByCategory(resolved.steps, category)
      : resolved.steps;

    return {
      content: [
        {
          type: "text" as const,
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
  }
}
