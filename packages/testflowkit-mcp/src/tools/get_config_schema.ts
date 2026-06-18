import { z } from "zod";
import { fetchConfigSchemaFromCli } from "../catalog/fetchConfigSchemaFromCli.js";
import { resolveCliBinary } from "../catalog/version.js";
import { HandlerParams, TkitTool } from "./tool.js";

const inputSchema = z.object({});

const configSchemaPayloadSchema = z.object({
  root: z.string(),
  version: z.string(),
  schema: z.record(z.unknown()),
});

export class GetConfigSchemaTool implements TkitTool<typeof inputSchema> {
  getName(): string {
    return "get_config_schema";
  }

  getDescription(): string {
    return "Return the full TestFlowKit configuration schema exported by the installed tkit CLI. Use this to generate or validate testflowkit.yml with authoritative field types, required flags, and constraints.";
  }

  getInputSchema() {
    return inputSchema;
  }

  async handler(params: HandlerParams<z.infer<typeof inputSchema>>) {
    const { configDir } = params;
    const cliBinary = resolveCliBinary();
    const output = fetchConfigSchemaFromCli(cliBinary, configDir);
    const payload = configSchemaPayloadSchema.parse(JSON.parse(output));

    return {
      content: [
        {
          type: "text" as const,
          text: JSON.stringify(payload, null, 2),
        },
      ],
    };
  }
}
