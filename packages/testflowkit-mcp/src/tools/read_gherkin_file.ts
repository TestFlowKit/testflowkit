import { z } from "zod";
import { readFeatureFile } from "../features/readFeature.js";
import { HandlerParams, TkitTool } from "./tool.js";

const inputSchema = z.object({
  path: z
    .string()
    .describe(
      "Relative path to the Gherkin feature file from the project root (e.g. features/registration.feature)",
    ),
});

export class ReadGherkinFileTool implements TkitTool<typeof inputSchema> {
  getName(): string {
    return "read_gherkin_file";
  }

  getDescription(): string {
    return "Read the content of a specific Gherkin feature file.";
  }

  getInputSchema() {
    return inputSchema;
  }

  async handler(params: HandlerParams<z.infer<typeof inputSchema>>) {
    const { config, configDir, input } = params;

    const content = readFeatureFile(
      input.path,
      config.featuresGlob,
      configDir,
    );

    return {
      content: [{ type: "text" as const, text: content }],
    };
  }
}
