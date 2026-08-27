import { z } from "zod";
import { extractMacros } from "../features/extractMacros.js";
import { HandlerParams, TkitTool } from "./tool.js";

const inputSchema = z.object({
  nameContains: z
    .string()
    .optional()
    .describe(
      "Case-insensitive substring filter on the macro (scenario) name.",
    ),
});

export class ListMacrosTool implements TkitTool<typeof inputSchema> {
  getName(): string {
    return "list_macros";
  }

  getDescription(): string {
    return (
      "List reusable @macro scenarios already defined in this project's Gherkin feature files. " +
      "Call this before writing a new scenario so you reuse an existing macro instead of duplicating steps. " +
      "Each result includes the required ${variable} placeholders and a ready-to-paste callExample: " +
      "call a macro by using its exact scenario name as a step, followed by a two-column data table " +
      "mapping each variable to a value."
    );
  }

  getInputSchema() {
    return inputSchema;
  }

  async handler(params: HandlerParams<z.infer<typeof inputSchema>>) {
    const { config, configDir, input } = params;
    const { nameContains } = input;

    const all = extractMacros(config.featuresGlob, configDir);
    const macros = nameContains
      ? all.filter((m) =>
          m.name.toLowerCase().includes(nameContains.toLowerCase()),
        )
      : all;

    return {
      content: [
        {
          type: "text" as const,
          text: JSON.stringify({ count: macros.length, macros }, null, 2),
        },
      ],
    };
  }
}
