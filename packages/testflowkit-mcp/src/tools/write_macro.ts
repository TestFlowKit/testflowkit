import { z } from "zod";
import { writeFeatureFile } from "../features/writeFeature.js";
import { CallToolResult } from "@modelcontextprotocol/sdk/types";
import { HandlerParams, TkitTool } from "./tool.js";

const MACRO_TAG_RE = /(^|\n)\s*@macro(\s|$)/i;

export function hasMacroTag(content: string): boolean {
  return MACRO_TAG_RE.test(content);
}

export class WriteMacroTool implements TkitTool<typeof inputSchema> {
  getName(): string {
    return "write_macro";
  }

  getDescription(): string {
    return "Create or overwrite a macro feature file. The content must include an @macro scenario definition and path must be within features_glob.";
  }

  getInputSchema() {
    return inputSchema;
  }

  handler(
    params: HandlerParams<z.infer<typeof inputSchema>>,
  ): Promise<CallToolResult> {
    const { config, configDir, input } = params;

    const { path: featurePath, content, createDirs } = input;

    if (!config.agent.capabilities.macros) {
      throw new Error(
        "Macro creation is disabled. Set agent.capabilities.macros: true in testflowkit.agent.yml to enable macro writes.",
      );
    }

    if (!hasMacroTag(content)) {
      throw new Error(
        `${this.getName()} requires at least one scenario tagged with "@macro".`,
      );
    }

    writeFeatureFile(
      featurePath,
      content,
      config.project.features_glob,
      configDir,
      createDirs,
    );

    return Promise.resolve({
      content: [
        {
          type: "text",
          text: JSON.stringify(
            { written: featurePath, ok: true, kind: "macro" },
            null,
            2,
          ),
        },
      ],
    });
  }
}

const inputSchema = z.object({
  path: z
    .string()
    .describe(
      "Relative path for the macro feature file from the project root (e.g. features/macros/auth.feature)",
    ),
  content: z
    .string()
    .describe("Full Gherkin content containing at least one @macro scenario"),
  createDirs: z
    .boolean()
    .default(true)
    .describe("Create parent directories if they do not exist"),
});
