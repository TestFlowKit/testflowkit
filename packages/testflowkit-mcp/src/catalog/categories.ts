import type { StepDefinition } from "../config/types.js";

export function getStepCategories(steps: StepDefinition[]): string[] {
  const categories = new Set<string>();
  for (const step of steps) {
    for (const category of step.categories) {
      categories.add(category);
    }
  }
  return [...categories].sort();
}

export function filterStepsByCategory(
  steps: StepDefinition[],
  category: string,
): StepDefinition[] {
  return steps.filter((step) => step.categories.includes(category));
}
