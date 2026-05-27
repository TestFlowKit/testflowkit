import { z } from "zod";
import { resolveCatalog } from "../catalog/resolveCatalog.js";
import { searchSentences } from "../catalog/searchSentences.js";
import { HandlerParams, TkitTool } from "./tool.js";

const inputSchema = z.object({
  query: z.string().describe("Keywords to search for (space-separated)"),
  category: z
    .string()
    .optional()
    .describe(
      "Filter by category: frontend, restapi, graphql, variable, assertions, form, navigation, visual, mouse, keyboard",
    ),
});

export class SearchSentencesTool implements TkitTool<typeof inputSchema> {
  getName(): string {
    return "search_sentences";
  }

  getDescription(): string {
    return "Search the TestFlowKit step catalog by keyword and optional category. Returns up to 30 matching sentences. Categories include: frontend, restapi, graphql, variable, assertions, form, navigation, visual, mouse, keyboard.";
  }

  getInputSchema() {
    return inputSchema;
  }

  async handler(params: HandlerParams<z.infer<typeof inputSchema>>) {
    const { config, configDir, input } = params;
    const { query, category } = input;

    const resolved = await resolveCatalog(config, configDir);
    const results = searchSentences(resolved.steps, query, category);

    return {
      content: [
        {
          type: "text" as const,
          text: JSON.stringify(
            {
              query,
              category: category ?? null,
              count: results.length,
              steps: results,
            },
            null,
            2,
          ),
        },
      ],
    };
  }
}
