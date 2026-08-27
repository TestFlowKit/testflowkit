import fs from "node:fs";
import path from "node:path";
import { listFeatureFiles } from "./globFeatures.js";

const MACRO_TAG = "@macro";
const VARIABLE_PATTERN = /\$\{([^}]+)\}/g;
const TAG_LINE_PATTERN = /^@\S+(\s+@\S+)*$/;
const SCENARIO_LINE_PATTERN = /^\s*Scenario(?: Outline)?:\s*(.+?)\s*$/;
const BOUNDARY_LINE_PATTERN = /^\s*(Scenario(?: Outline)?|Background|Feature):/;
const STEP_KEYWORD_PATTERN = /^(Given|When|Then|And|But)\s+/;

export type MacroInfo = {
  name: string;
  file: string;
  line: number;
  variables: string[];
  steps: string[];
  callExample: string;
};

/**
 * Scans every Gherkin feature file under featuresGlob for `@macro`-tagged
 * scenarios (macros are resolved across the whole gherkin_location, not
 * per-file, matching the Go framework's own macro resolution).
 */
export function extractMacros(
  featuresGlob: string,
  baseDir: string,
): MacroInfo[] {
  const files = listFeatureFiles(featuresGlob, baseDir);
  const macros: MacroInfo[] = [];

  for (const file of files) {
    const content = fs.readFileSync(file, "utf-8");
    macros.push(
      ...extractMacrosFromContent(content, path.relative(baseDir, file)),
    );
  }

  return macros;
}

export function extractMacrosFromContent(
  content: string,
  relativePath: string,
): MacroInfo[] {
  const lines = content.split(/\r?\n/);
  const macros: MacroInfo[] = [];

  let pendingTags: string[] = [];
  let i = 0;

  while (i < lines.length) {
    const trimmed = lines[i].trim();

    if (trimmed === "") {
      i++;
      continue;
    }

    if (TAG_LINE_PATTERN.test(trimmed)) {
      pendingTags.push(...trimmed.split(/\s+/));
      i++;
      continue;
    }

    const scenarioMatch = SCENARIO_LINE_PATTERN.exec(lines[i]);
    if (scenarioMatch) {
      const isMacro = pendingTags.includes(MACRO_TAG);
      const name = scenarioMatch[1];
      const startLine = i + 1;
      pendingTags = [];
      i++;

      const bodyLines: string[] = [];
      while (i < lines.length && !isBoundary(lines[i])) {
        bodyLines.push(lines[i]);
        i++;
      }

      if (isMacro) {
        macros.push(buildMacroInfo(name, relativePath, startLine, bodyLines));
      }
      continue;
    }

    // Feature:, Background:, or free-text description lines reset any pending tags.
    pendingTags = [];
    i++;
  }

  return macros;
}

function isBoundary(line: string): boolean {
  const trimmed = line.trim();
  return TAG_LINE_PATTERN.test(trimmed) || BOUNDARY_LINE_PATTERN.test(line);
}

function buildMacroInfo(
  name: string,
  file: string,
  line: number,
  bodyLines: string[],
): MacroInfo {
  const bodyText = bodyLines.join("\n");
  const variables = extractVariables(bodyText);
  const steps = bodyLines
    .map((l) => l.trim())
    .filter((l) => STEP_KEYWORD_PATTERN.test(l));

  return {
    name,
    file,
    line,
    variables,
    steps,
    callExample: buildCallExample(name, variables),
  };
}

function extractVariables(text: string): string[] {
  const seen = new Set<string>();
  const result: string[] = [];
  for (const match of text.matchAll(VARIABLE_PATTERN)) {
    const name = match[1].trim();
    if (!seen.has(name)) {
      seen.add(name);
      result.push(name);
    }
  }
  return result;
}

function buildCallExample(name: string, variables: string[]): string {
  if (variables.length === 0) {
    return `Given ${name}`;
  }
  const rows = variables.map((v) => `    | ${v} | <value> |`).join("\n");
  return `Given ${name}\n${rows}`;
}
