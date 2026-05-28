import type { StepDefinition } from "../config/types.js";

const MAX_RESULTS = 30;

export function searchSentences(
  steps: StepDefinition[],
  query: string,
  category?: string
): StepDefinition[] {
  const terms = query
    .toLowerCase()
    .split(/\s+/)
    .filter((t) => t.length > 1);

  let filtered = steps;

  if (category) {
    const cat = category.toLowerCase();
    filtered = filtered.filter((s) =>
      s.categories.some((c) => c.toLowerCase() === cat)
    );
  }

  if (terms.length === 0) {
    return filtered.slice(0, MAX_RESULTS);
  }

  const scored = filtered
    .map((step) => {
      const target = `${step.sentence} ${step.description} ${step.categories.join(" ")}`.toLowerCase();
      const matches = terms.filter((t) => target.includes(t)).length;
      return { step, matches };
    })
    .filter(({ matches }) => matches > 0)
    .sort((a, b) => b.matches - a.matches);

  return scored.slice(0, MAX_RESULTS).map(({ step }) => step);
}
