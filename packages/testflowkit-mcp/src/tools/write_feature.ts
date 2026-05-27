import { z } from "zod";
import { writeFeatureFile } from "../features/writeFeature.js";
import { HandlerParams, TkitTool } from "./tool.js";
import { hasMacroTag } from "./write_macro.js";

const inputSchema = z.object({
  path: z
    .string()
    .describe(
      "Relative path for the feature file from the project root (e.g. features/auth/registration.feature)",
    ),
  content: z.string().describe("Full Gherkin content for the feature file"),
  createDirs: z
    .boolean()
    .optional()
    .default(true)
    .describe("Create parent directories if they do not exist"),
});

export class WriteFeatureTool implements TkitTool<typeof inputSchema> {
  getName(): string {
    return "write_feature";
  }

  getDescription(): string {
    return "Create or overwrite a feature file. The path must be within the features_glob directory. Parent directories are created automatically.";
  }

  getInputSchema() {
    return inputSchema;
  }

  async handler(params: HandlerParams<z.infer<typeof inputSchema>>) {
    const { config, configDir, input } = params;
    const { path: featurePath, content, createDirs } = input;

    if (!config.agent.capabilities.macros && hasMacroTag(content)) {
      throw new Error(
        "Macro creation is disabled. Set agent.capabilities.macros: true in testflowkit.agent.yml to enable macro writes.",
      );
    }

    writeFeatureFile(
      featurePath,
      content,
      config.project.features_glob,
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
