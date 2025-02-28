import { defineContentConfig, defineCollection } from "@nuxt/content";
import { sentenceValidationSchema } from "./data/sentence";

export default defineContentConfig({
  collections: {
    sentence: defineCollection({
      source: "sentences/**/*.json",
      type: "data",
      schema: sentenceValidationSchema,
    }),
  },
});
