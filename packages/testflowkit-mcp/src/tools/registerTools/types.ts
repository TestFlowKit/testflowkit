import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import type { AgentConfig } from "../../config/types.js";
import type { CatalogSession } from "../../context/session.js";

export type RegisterTool = (
  server: McpServer,
  config: AgentConfig,
  configDir: string,
  session: CatalogSession,
) => void;
