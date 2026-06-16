import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import type { ProjectConfig } from "../config/types.js";
import { GetStepCategoriesTool } from "./get_step_categories.js";
import { GetStepCatalogTool } from "./get_step_catalog.js";
import { ListFeaturesTool } from "./list_features.js";
import { ReadFeatureTool } from "./read_feature.js";
import { ReadTestConfigTool } from "./read_test_config.js";
import { TkitTool } from "./tool.js";
import { WriteFeatureTool } from "./write_feature.js";

export function registerTools(
  server: McpServer,
  config: ProjectConfig,
  configDir: string,
): void {
  const tools: Array<TkitTool<any>> = [
    new GetStepCategoriesTool(),
    new GetStepCatalogTool(),
    new ReadTestConfigTool(),
    new ListFeaturesTool(),
    new ReadFeatureTool(),
    new WriteFeatureTool(),
  ];

  for (const tool of tools) {
    server.registerTool(
      tool.getName(),
      {
        description: tool.getDescription(),
        inputSchema: tool.getInputSchema(),
      },
      async (input: any) =>
        tool.handler({
          config,
          configDir,
          input,
        }),
    );
  }
}
