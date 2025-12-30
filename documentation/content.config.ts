import { defineContentConfig, defineCollection, z } from "@nuxt/content";
import { sentenceValidationSchema } from "./data/sentence";

export default defineContentConfig({
  collections: {
    docs: defineCollection({
      source: "docs/**/*.md",
      type: "page",
      schema: z.object({
        title: z.string(),
        description: z.string().optional(),
        navigation: z.object({
          title: z.string().optional(),
        }).optional(),
      }),
    }),
    sentence: defineCollection({
      source: "sentences/**/*.json",
      type: "data",
      schema: sentenceValidationSchema,
    }),
  },
});
