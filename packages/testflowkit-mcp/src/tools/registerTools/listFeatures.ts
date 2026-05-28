import path from "node:path";
import { listFeatureFiles } from "../../features/globFeatures.js";
import type { RegisterTool } from "./types.js";

export const registerListFeaturesTool: RegisterTool = (
  server,
  config,
  configDir,
) => {
  server.registerTool(
    "list_features",
    {
      description:
        "List all feature files in the project matching the configured features_glob pattern.",
    },
    async () => {
      const files = listFeatureFiles(config.project.features_glob, configDir);
      const relative = files.map((f) => path.relative(configDir, f));
      return {
        content: [
          {
            type: "text",
            text: JSON.stringify(
              { count: relative.length, files: relative },
              null,
              2,
            ),
          },
        ],
      };
    },
  );
};
