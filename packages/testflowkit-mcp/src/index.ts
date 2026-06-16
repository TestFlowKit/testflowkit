#!/usr/bin/env node
import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import {
  loadProjectConfig,
  LoadedProjectConfig,
} from "./config/loadProjectConfig.js";
import { initializeSession } from "./context/session.js";
import { Logger } from "./logger.js";
import { registerTools } from "./tools/index.js";

async function main(): Promise<void> {
  const { config, configDir } = loadConfig();
  const session = initializeSession();

  const server = new McpServer({
    name: "testflowkit",
    version: "0.1.0",
  });


  registerTools(server, config, configDir);

  const transport = new StdioServerTransport();
  await server.connect(transport);

  Logger.log(`Started. Config: ${config.configPath}`);
  if (session.cliVersion) {
    Logger.log(
      `CLI version at boot: ${session.cliVersion} (source: ${session.source})`,
    );
  }
}

function loadConfig(): LoadedProjectConfig {
  try {
    return loadProjectConfig(process.cwd());
  } catch (err) {
    Logger.error(`Config error: ${String(err)}`);
    process.exit(1);
  }
}

main().catch((err) => {
  Logger.error(`Fatal: ${String(err)}`);
  process.exit(1);
});
