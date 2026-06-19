import { z } from "zod";
import { writeFeatureFile } from "../features/writeFeature.js";
import { HandlerParams, TkitTool } from "./tool.js";

const inputSchema = z.object({
  path: z
    .string()
    .describe(
      "Relative path for the Gherkin feature file from the project root (e.g. features/auth/registration.feature)",
    ),
  content: z.string().describe("Full Gherkin content for the feature file"),
  createDirs: z
    .boolean()
    .optional()
    .default(true)
    .describe("Create parent directories if they do not exist"),
});

export class WriteGherkinFileTool implements TkitTool<typeof inputSchema> {
  getName(): string {
    return "write_gherkin_file";
  }

  getDescription(): string {
    return "Create or overwrite a Gherkin feature file. The path must be within settings.gherkin_location from testflowkit.yml. Parent directories are created automatically.";
  }

  getTagHint(defaultTagsForDraft?: string): string {
    return defaultTagsForDraft
      ? ` Tag every new scenario with: ${defaultTagsForDraft}.`
      : "";
  }

  getInputSchema() {
    return inputSchema;
  }

  async handler(params: HandlerParams<z.infer<typeof inputSchema>>) {
    const { config, configDir, input } = params;
    const { path: featurePath, content, createDirs } = input;

    writeFeatureFile(
      featurePath,
      content,
      config.featuresGlob,
      configDir,
      createDirs,
    );

    return {
      content: [
        {
          type: "text" as const,
          text: JSON.stringify({ written: featurePath, ok: true }),
        },
      ],
    };
  }
}
