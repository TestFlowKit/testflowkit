import path from "node:path";
import { z } from "zod";
import { listFeatureFiles } from "../features/globFeatures.js";
import { HandlerParams, TkitTool } from "./tool.js";

const inputSchema = z.object({});

export class ListFeaturesTool implements TkitTool<typeof inputSchema> {
  getName(): string {
    return "list_features";
  }

  getDescription(): string {
    return "List all feature files in the project matching settings.gherkin_location from testflowkit.yml.";
  }

  getInputSchema() {
    return inputSchema;
  }

  async handler(params: HandlerParams<z.infer<typeof inputSchema>>) {
    const { config, configDir } = params;

    const files = listFeatureFiles(config.featuresGlob, configDir);
    const relative = files.map((f) => path.relative(configDir, f));

    return {
      content: [
        {
          type: "text" as const,
          text: JSON.stringify(
            { count: relative.length, files: relative },
            null,
            2,
          ),
        },
      ],
    };
  }
}
