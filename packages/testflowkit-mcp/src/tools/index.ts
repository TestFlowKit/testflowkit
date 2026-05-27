import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import type { AgentConfig } from "../config/types.js";
import { GetCliVersionCommandsTool } from "./get_cli_version_commands.js";
import { GetGuidelinesTool } from "./get-guidelines/get_guidelines.js";
import { GetRunCommandTool } from "./get_run_command.js";
import { GetStepCatalogTool } from "./get_step_catalog.js";
import { ListFeaturesTool } from "./list_features.js";
import { ReadFeatureTool } from "./read_feature.js";
import { ReadTestConfigTool } from "./read_test_config.js";
import { SearchSentencesTool } from "./search_sentences.js";
import { TkitTool } from "./tool.js";
import { WriteFeatureTool } from "./write_feature.js";
import { WriteMacroTool } from "./write_macro.js";

export function registerTools(
  server: McpServer,
  config: AgentConfig,
  configDir: string,
): void {
  const tools: Array<TkitTool<any>> = [
    new GetCliVersionCommandsTool(),
    new GetRunCommandTool(),
    new GetStepCatalogTool(),
    new SearchSentencesTool(),
    new ReadTestConfigTool(),
    new ListFeaturesTool(),
    new ReadFeatureTool(),
    new GetGuidelinesTool(),
    new WriteFeatureTool(),
    new WriteMacroTool(),
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
