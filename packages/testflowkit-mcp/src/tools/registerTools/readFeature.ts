import { z } from "zod";
import { readFeatureFile } from "../../features/readFeature.js";
import type { RegisterTool } from "./types.js";

export const registerReadFeatureTool: RegisterTool = (
  server,
  config,
  configDir,
) => {
  server.registerTool(
    "read_feature",
    {
      description: "Read the content of a specific feature file.",
      inputSchema: z.object({
        path: z
          .string()
          .describe(
            "Relative path to the feature file from the project root (e.g. features/registration.feature)",
          ),
      }),
    },
    async ({ path: featurePath }) => {
      const content = readFeatureFile(
        featurePath,
        config.project.features_glob,
        configDir,
      );
      return {
        content: [{ type: "text", text: content }],
      };
    },
  );
};
