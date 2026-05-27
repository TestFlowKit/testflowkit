import { z } from "zod";
import { loadTestConfigSummary } from "../config/loadTestConfigSummary.js";
import { HandlerParams, TkitTool } from "./tool.js";

const inputSchema = z.object({});

export class ReadTestConfigTool implements TkitTool<typeof inputSchema> {
  getName(): string {
    return "read_test_config";
  }

  getDescription(): string {
    return "Read a summary of the project's testflowkit.yml (or config.yml). Returns API names, operation/endpoint references (in api_name.operation format), frontend pages, and element group names. Secrets are redacted.";
  }

  getInputSchema() {
    return inputSchema;
  }

  async handler(params: HandlerParams<z.infer<typeof inputSchema>>) {
    const { config, configDir } = params;

    const summary = loadTestConfigSummary(
      config.project.test_config,
      configDir,
    );

    return {
      content: [
        {
          type: "text" as const,
          text: JSON.stringify(summary, null, 2),
        },
      ],
    };
  }
}
