import { z } from "zod";
import { CLI_VERSION_RE } from "../catalog/version.js";
import { HandlerParams, TkitTool } from "./tool.js";

const inputSchema = z.object({});

export class GetCliVersionCommandsTool implements TkitTool<typeof inputSchema> {
  getName(): string {
    return "get_cli_version_commands";
  }

  getDescription(): string {
    return (
      "Returns the shell commands to run to detect the installed tkit CLI version. " +
      "Execute them in order, stop at the first that succeeds, extract the version with the provided regex, " +
      "then pass it as cliVersion to get_step_catalog."
    );
  }

  getInputSchema() {
    return inputSchema;
  }

  async handler(_: HandlerParams<z.infer<typeof inputSchema>>) {
    return {
      content: [
        {
          type: "text" as const,
          text: JSON.stringify(
            {
              commands: ["tkit version", "npx --yes @testflowkit/cli version"],
              versionRegex: CLI_VERSION_RE.source,
              instructions:
                "Run each command in the project terminal. Extract the version string using versionRegex. " +
                "Pass the matched version as cliVersion to get_step_catalog.",
            },
            null,
            2,
          ),
        },
      ],
    };
  }
}
