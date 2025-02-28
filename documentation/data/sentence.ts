import { z } from "@nuxt/content";

export const sentenceValidationSchema = z.object({
  sentence: z.string(),
  description: z.string(),
  category: z.string(),
  gherkinExample: z.string(),
  variables: z.array(
    z.object({
      name: z.string(),
      type: z.string(),
    }),
  ),
});

export type Sentence = {
  sentence: string;
  description: string;
  category: string;
  gherkinExample: string;
  variables: Array<{
    name: string;
    description?: string;
    type: string;
  }>;
};
