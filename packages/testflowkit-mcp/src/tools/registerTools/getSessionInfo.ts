import type { RegisterTool } from "./types.js";

const MCP_PACKAGE_VERSION = "0.1.0";

export const registerGetSessionInfoTool: RegisterTool = (
  server,
  config,
  configDir,
  session,
) => {
  server.registerTool(
    "get_session_info",
    {
      description:
        "Returns the current MCP session state: resolved tkit CLI version, version source, and server metadata. " +
        "Useful for debugging catalog resolution issues.",
    },
    async () => {
      const probeCliEnabled =
        !config.step_catalog.file && !config.step_catalog.url;

      return {
        content: [
          {
            type: "text" as const,
            text: JSON.stringify(
              {
                cliVersion: session.cliVersion,
                source: session.source,
                initializedAt: session.initializedAt,
                mcpPackageVersion: MCP_PACKAGE_VERSION,
                configDir,
                probeCliEnabled,
                catalogMode: config.step_catalog.file
                  ? "file"
                  : config.step_catalog.url
                    ? "url"
                    : "release",
              },
              null,
              2,
            ),
          },
        ],
      };
    },
  );
};
