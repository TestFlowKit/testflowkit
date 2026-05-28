import type { RegisterTool } from "./types.js";

export const registerGetRunCommandTool: RegisterTool = (
  server,
  config,
  configDir,
) => {
  server.registerTool(
    "get_run_command",
    {
      description:
        "Returns the configured tkit test run command. Execute it in the project terminal and observe the output.",
    },
    async () => ({
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
    }),
  );
};
