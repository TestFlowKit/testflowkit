import { z } from "zod";

export const AgentConfigSchema = z.object({
  version: z.literal(1),
  project: z
    .object({
      test_config: z.string().default("./testflowkit.yml"),
      features_glob: z.string().default("features/**/*.feature"),
    })
    .default({}),
  step_catalog: z
    .object({
      cli_version: z.string().optional(),
      file: z.string().optional(),
      url: z.string().url().optional(),
    })
    .strict()
    .default({}),
  agent: z
    .object({
      capabilities: z
        .object({
          macros: z.boolean().default(true),
        })
        .default({}),
      default_tags_for_draft: z.string().default("@wip @ai-generated"),
      run_command: z
        .string()
        .default("tkit run -c testflowkit.yml --tags @wip"),
    })
    .default({}),
});

export type AgentConfig = z.infer<typeof AgentConfigSchema>;

export type StepDefinition = {
  sentence: string;
  description: string;
  categories: string[];
  example: string;
  variables?: StepVariable[];
};

export type StepVariable = {
  name: string;
  description: string;
  type: string;
};

export type CatalogMetadata = {
  source: "file" | "url" | "local" | "release" | "cache";
  catalogVersion: string | null;
  cliVersion: string | null;
  warnings: string[];
};

export type ResolvedCatalog = {
  steps: StepDefinition[];
  meta: CatalogMetadata;
};
