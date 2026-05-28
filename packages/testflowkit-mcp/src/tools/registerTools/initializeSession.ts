import { z } from "zod";
import {
  setSessionCliVersion,
  toSemverBase,
  isCanary,
} from "../../context/session.js";
import type { RegisterTool } from "./types.js";

export const registerInitializeSessionTool: RegisterTool = (
  server,
  _config,
  _configDir,
  session,
) => {
  server.registerTool(
    "initialize_session",
    {
      description:
        "Set the installed tkit CLI version for the current MCP session. " +
        "Call this once per connection after running `tkit version` in the project terminal. " +
        "All subsequent catalog tools (get_step_catalog, get_step_categories) will use this version " +
        "without requiring it to be passed on every call.",
      inputSchema: z.object({
        cliVersion: z
          .string()
          .describe(
            'Installed tkit CLI version string (e.g. "1.4.2" or "3.6.1-canary.abc1234"). ' +
              "Obtain by running `tkit version` or `npx --yes @testflowkit/cli version` in the project terminal.",
          ),
      }),
    },
    async ({ cliVersion }) => {
      try {
        setSessionCliVersion(session, cliVersion);
        const semverBase = toSemverBase(cliVersion);
        const canary = isCanary(cliVersion);
        return {
          content: [
            {
              type: "text" as const,
              text: JSON.stringify(
                {
                  ok: true,
                  cliVersion,
                  semverBase,
                  isCanary: canary,
                  message:
                    "Session ready. You can now call get_step_catalog and get_step_categories without passing cliVersion again.",
                },
                null,
                2,
              ),
            },
          ],
        };
      } catch (err) {
        return {
          content: [
            {
              type: "text" as const,
              text: JSON.stringify(
                {
                  ok: false,
                  error: "invalid_cli_version",
                  message: err instanceof Error ? err.message : String(err),
                },
                null,
                2,
              ),
            },
          ],
          isError: true,
        };
      }
    },
  );
};
