import { z } from "zod";
import { writeFeatureFile } from "../../features/writeFeature.js";
import type { RegisterTool } from "./types.js";

export const registerWriteFeatureTool: RegisterTool = (
  server,
  config,
  configDir,
) => {
  server.registerTool(
    "write_feature",
    {
      description:
        "Create or overwrite a feature file. The path must be within the features_glob directory. Parent directories are created automatically.",
      inputSchema: z.object({
        path: z
          .string()
          .describe(
            "Relative path for the feature file from the project root (e.g. features/auth/registration.feature)",
          ),
        content: z
          .string()
          .describe("Full Gherkin content for the feature file"),
        createDirs: z
          .boolean()
          .optional()
          .default(true)
          .describe("Create parent directories if they do not exist"),
      }),
    },
    async ({ path: featurePath, content, createDirs }) => {
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
            type: "text",
            text: JSON.stringify({ written: featurePath, ok: true }),
          },
        ],
      };
    },
  );
};
