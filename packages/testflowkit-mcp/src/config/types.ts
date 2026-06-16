import { z } from "zod";

export const ProjectConfigSchema = z.object({
  settings: z
    .object({
      gherkin_location: z.string().min(1).default("./features"),
    })
    .default({}),
  agent: z
    .object({
      default_tags_for_draft: z.string().optional(),
      run_command: z.string().optional(),
      step_catalog: z
        .object({
          file: z.string().optional(),
          url: z.string().optional(),
        })
        .optional(),
    })
    .optional(),
});

export type AgentConfig = {
  defaultTagsForDraft?: string;
  runCommand?: string;
  stepCatalog?: { file?: string; url?: string };
};

export type ProjectConfig = {
  configPath: string;
  featuresGlob: string;
  agent?: AgentConfig;
};

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
  source: "local";
  cliVersion: string | null;
  warnings: string[];
};

export type ResolvedCatalog = {
  steps: StepDefinition[];
  meta: CatalogMetadata;
};

export function toFeaturesGlob(gherkinLocation: string): string {
  const trimmed = gherkinLocation.trim().replace(/\/+$/, "");
  if (trimmed.includes("*") || trimmed.endsWith(".feature")) {
    return trimmed;
  }
  return `${trimmed}/**/*.feature`;
}
