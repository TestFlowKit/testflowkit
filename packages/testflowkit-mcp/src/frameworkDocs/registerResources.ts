import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import { formatFrameworkDocsList, listFrameworkDocs } from "./indexer.js";
import type { FrameworkDocsIndex } from "./types.js";

export const FRAMEWORK_DOCS_INDEX_URI = "docs://framework/features/index";

export function frameworkDocUri(slug: string): string {
  return `docs://framework/features/${slug}`;
}

export function registerFrameworkDocResources(
  server: McpServer,
  index: FrameworkDocsIndex,
): void {
  server.registerResource(
    "framework-features-index",
    FRAMEWORK_DOCS_INDEX_URI,
    {
      title: "Framework documentation index",
      description:
        "Index of TestFlowKit framework documentation (macros, random data, global hooks, variables, API testing, frontend testing).",
      mimeType: "text/markdown",
    },
    async () => ({
      contents: [
        {
          uri: FRAMEWORK_DOCS_INDEX_URI,
          mimeType: "text/markdown",
          text: formatFrameworkDocsList(index),
        },
      ],
    }),
  );

  for (const doc of listFrameworkDocs(index)) {
    const uri = frameworkDocUri(doc.slug);

    server.registerResource(
      `framework-feature-${doc.slug}`,
      uri,
      {
        title: doc.title,
        description: doc.description,
        mimeType: "text/markdown",
      },
      async () => ({
        contents: [
          {
            uri,
            mimeType: "text/markdown",
            text: doc.content,
          },
        ],
      }),
    );
  }
}
