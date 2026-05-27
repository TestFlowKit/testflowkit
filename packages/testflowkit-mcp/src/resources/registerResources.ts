import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";
import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";

type DocResource = {
  name: string;
  uri: string;
  relativePath: string;
  title: string;
  description: string;
};

const DOC_RESOURCES: DocResource[] = [
  {
    name: "guidelines-macros",
    uri: "testflowkit://guidelines/macros",
    relativePath: "docs/macros.md",
    title: "Macros Guidelines",
    description: "TestFlowKit macro authoring and usage documentation.",
  },
  {
    name: "guidelines-ide-agent",
    uri: "testflowkit://guidelines/ide-agent",
    relativePath: "docs/ide-agent.md",
    title: "IDE Agent Guidelines",
    description: "TestFlowKit IDE agent setup and MCP usage guidance.",
  },
  {
    name: "guidelines-copilot-instructions",
    uri: "testflowkit://guidelines/copilot-instructions",
    relativePath: "docs/copilot-instructions.md",
    title: "Copilot Instructions Template",
    description: "Boilerplate instruction rules for TestFlowKit Copilot usage.",
  },
];

type RegisterResourcesOptions = {
  docsBaseDir?: string;
};

function getDefaultDocsBaseDir(): string {
  const moduleDir = path.dirname(fileURLToPath(import.meta.url));
  return path.resolve(moduleDir, "..", "..");
}

export function registerResources(
  server: McpServer,
  configDir: string,
  options?: RegisterResourcesOptions,
): void {
  const docsBaseDir = options?.docsBaseDir ?? getDefaultDocsBaseDir();

  for (const resource of DOC_RESOURCES) {
    server.registerResource(
      resource.name,
      resource.uri,
      {
        title: resource.title,
        description: resource.description,
        mimeType: "text/markdown",
      },
      async () => {
        const bundledPath = path.resolve(docsBaseDir, resource.relativePath);
        const workspacePath = path.resolve(configDir, resource.relativePath);

        const abs = fs.existsSync(bundledPath)
          ? bundledPath
          : fs.existsSync(workspacePath)
            ? workspacePath
            : null;

        if (!abs) {
          return {
            contents: [
              {
                uri: resource.uri,
                mimeType: "text/markdown",
                text:
                  `# Resource not available\n\n` +
                  `The documentation file was not found at: ${resource.relativePath}\n` +
                  `You can still use the tool get_guidelines for bundled guidance.`,
              },
            ],
          };
        }

        return {
          contents: [
            {
              uri: resource.uri,
              mimeType: "text/markdown",
              text: fs.readFileSync(abs, "utf-8"),
            },
          ],
        };
      },
    );
  }
}
