/**
 * Shared type definitions mirroring the sentence JSON schema produced by
 * scripts/doc_generator/main.go and validated in documentation/data/sentence.ts.
 */

export type SentenceVariable = {
  name: string;
  type: string;
  description?: string;
};

export type Sentence = {
  sentence: string;
  description: string;
  categories: string[];
  gherkinExample: string;
  variables: SentenceVariable[];
};

/**
 * A macro extracted from a @macro-tagged scenario in a workspace .feature file.
 * Variables are the de-duplicated, ordered list of ${var} placeholder names.
 */
export type Macro = {
  name: string;
  variables: string[];   // de-duplicated, first-appearance order from macro body
  sourceFile: string;    // absolute path of the .feature file defining this macro
};
