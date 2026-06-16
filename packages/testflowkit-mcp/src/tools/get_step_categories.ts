import { z } from "zod";
import { getStepCategories } from "../catalog/categories.js";
import { resolveCatalog } from "../catalog/resolveCatalog.js";
import { HandlerParams, TkitTool } from "./tool.js";

const inputSchema = z.object({});

export class GetStepCategoriesTool implements TkitTool<typeof inputSchema> {
  getName(): string {
    return "get_step_categories";
  }

  getDescription(): string {
    return (
      "List available step definition categories from the TestFlowKit catalog via the installed tkit CLI. " +
      "Use a category name with get_step_catalog to filter steps."
    );
  }

  getInputSchema() {
    return inputSchema;
  }

  async handler(params: HandlerParams<z.infer<typeof inputSchema>>) {
    const { configDir } = params;

    const resolved = await resolveCatalog(configDir);
    const categories = getStepCategories(resolved.steps);

    return {
      content: [
        {
          type: "text" as const,
          text: JSON.stringify(
            {
              meta: resolved.meta,
              count: categories.length,
              categories,
            },
            null,
            2,
          ),
        },
      ],
    };
  }
}
