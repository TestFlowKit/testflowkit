import { z } from "zod";
import { HandlerParams, TkitTool } from "./tool.js";

const inputSchema = z.object({});

export class GetRunCommandTool implements TkitTool<typeof inputSchema> {
  getName(): string {
    return "get_run_command";
  }

  getDescription(): string {
    return "Returns the configured tkit test run command. Execute it in the project terminal and observe the output.";
  }

  getInputSchema() {
    return inputSchema;
  }

  async handler(params: HandlerParams<z.infer<typeof inputSchema>>) {
    const { config, configDir } = params;

    return {
      content: [
        {
          type: "text" as const,
          text: JSON.stringify(
            {
              command: config.agent.run_command,
              cwd: configDir,
              instructions:
                "Run this command in the project directory (cwd). " +
                "Report the exit code, stdout and stderr back to the user.",
            },
            null,
            2,
          ),
        },
      ],
    };
  }
}
