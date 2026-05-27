import { z } from "zod";
import type { AgentConfig } from "../../config/types.js";
import { HandlerParams, TkitTool } from "../tool.js";

type Concept = {
  name: string;
  description: string;
  documentation: string;
  capability: {
    key: string;
    enabled: boolean;
    whenDisabled: string;
  };
};

type Guidelines = {
  version: 1;
  concepts: {
    macro: Concept;
  };
};

function buildGuidelines(config: AgentConfig): Guidelines {
  return {
    version: 1,
    concepts: {
      macro: {
        name: "Macro",
        description:
          "A macro is a reusable scenario template that can be invoked with different parameters. Macros are defined as scenarios tagged with @macro.",
        documentation:
          "Use macros to avoid duplication of common scenario patterns, and to create higher-level abstractions in your test suite. Define a macro once, then invoke it from other scenarios with different variable values.", 
        capability: {
          key: "agent.capabilities.macros",
          enabled: config.agent.capabilities.macros,
          whenDisabled:
            "Macro creation is denied by write_macro and write_feature rejects content containing @macro.",
        },
        ,
        reading: {
          listTool: "list_features",
          readTool: "read_feature",
          identifyMacroRule:
            "A scenario is a macro definition when it has the @macro tag.",
        },
        promptTemplate:
          "Create a reusable login macro with variables email and password, then provide one invocation example scenario.",
      },
    },
  };
}

const inputSchema = z.object({
  concept: z
    .string()
    .optional()
    .describe(
      "Guideline concept name (e.g. macro). Unknown or empty values return all available guidelines.",
    ),
});

export class GetGuidelinesTool implements TkitTool<typeof inputSchema> {
  getName(): string {
    return "get_guidelines";
  }

  getDescription(): string {
    return "Return bundled authoring and reading guidelines for a concept. Pass concept (e.g. macro) to get one concept; unknown or empty concept returns all guidelines.";
  }

  getInputSchema() {
    return inputSchema;
  }

  async handler(params: HandlerParams<z.infer<typeof inputSchema>>) {
    const { config, input } = params;
    const { concept } = input;

    const guidelines = buildGuidelines(config);
    const normalized = concept?.trim().toLowerCase();

    if (!normalized || !(normalized in guidelines.concepts)) {
      return {
        content: [
          {
            type: "text" as const,
            text: JSON.stringify(guidelines, null, 2),
          },
        ],
      };
    }

    const conceptGuidelines =
      guidelines.concepts[normalized as keyof Guidelines["concepts"]];

    return {
      content: [
        {
          type: "text" as const,
          text: JSON.stringify(
            {
              version: guidelines.version,
              concept: normalized,
              guidelines: conceptGuidelines,
            },
            null,
            2,
          ),
        },
      ],
    };
  }
}
