import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import type { AgentConfig } from "../config/types.js";
import type { CatalogSession } from "../context/session.js";
import {
  registerInitializeSessionTool,
  registerGetSessionInfoTool,
  registerGetRunCommandTool,
  registerGetStepCatalogTool,
  registerGetStepCategoriesTool,
  registerReadTestConfigTool,
  registerListFeaturesTool,
  registerReadFeatureTool,
  registerWriteFeatureTool,
} from "./registerTools/index.js";

export function registerTools(
  server: McpServer,
  config: AgentConfig,
  configDir: string,
  session: CatalogSession,
): void {
  registerInitializeSessionTool(server, config, configDir, session);
  registerGetSessionInfoTool(server, config, configDir, session);
  registerGetRunCommandTool(server, config, configDir, session);
  registerGetStepCatalogTool(server, config, configDir, session);
  registerGetStepCategoriesTool(server, config, configDir, session);
  registerReadTestConfigTool(server, config, configDir, session);
  registerListFeaturesTool(server, config, configDir, session);
  registerReadFeatureTool(server, config, configDir, session);
  registerWriteFeatureTool(server, config, configDir, session);
}
