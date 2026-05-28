import { loadTestConfigSummary } from "../../config/loadTestConfigSummary.js";
import type { RegisterTool } from "./types.js";

export const registerReadTestConfigTool: RegisterTool = (
  server,
  config,
  configDir,
) => {
  server.registerTool(
    "read_test_config",
    {
      description:
        "Read a summary of the project's testflowkit.yml (or config.yml). Returns API names, operation/endpoint references (in api_name.operation format), frontend pages, and element group names. Secrets are redacted.",
    },
    async () => {
      const summary = loadTestConfigSummary(
        config.project.test_config,
        configDir,
      );
      return {
        content: [
          {
            type: "text",
            text: JSON.stringify(summary, null, 2),
          },
        ],
      };
    },
  );
};
