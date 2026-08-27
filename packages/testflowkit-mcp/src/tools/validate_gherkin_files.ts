import { z } from "zod";
import { resolveCliBinary } from "../catalog/version.js";
import { runValidate } from "../features/runValidate.js";
import { HandlerParams, TkitTool } from "./tool.js";

const inputSchema = z.object({
  tags: z
    .string()
    .optional()
    .describe(
      "Cucumber tag expression to scope validation (e.g. the agent.default_tags_for_draft value), " +
        "so unrelated pre-existing failures elsewhere in the project don't block this check.",
    ),
});

export class ValidateGherkinFilesTool implements TkitTool<typeof inputSchema> {
  getName(): string {
    return "validate_gherkin_files";
  }

  getDescription(): string {
    return (
      "Statically validate the project's Gherkin feature files via the installed tkit CLI (tkit validate). " +
      "Reports undefined steps (including calls to a macro that doesn't exist), malformed or missing page/element " +
      "references, and undefined environment variables. Does not execute any browser or API action. " +
      "Call this after write_gherkin_file to catch mistakes before running the suite with tkit run."
    );
  }

  getInputSchema() {
    return inputSchema;
  }

  async handler(params: HandlerParams<z.infer<typeof inputSchema>>) {
    const { config, configDir, input } = params;
    const cliBinary = resolveCliBinary();

    const result = runValidate(
      cliBinary,
      configDir,
      config.configPath,
      input.tags,
    );

    return {
      content: [
        {
          type: "text" as const,
          text: JSON.stringify(result, null, 2),
        },
      ],
    };
  }
}
