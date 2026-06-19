#!/usr/bin/env tsx
import fs from "node:fs";
import path from "node:path";
import { parse as parseYaml, stringify as stringifyYaml } from "yaml";
import process from "node:process";

const PACKAGE_ROOT = process.cwd();
const DOCS_CONTENT_ROOT = path.join(
  PACKAGE_ROOT,
  "..",
  "..",
  "documentation/content/docs",
);

const OUTPUT_DIR = path.join(PACKAGE_ROOT, "docs/features");

const DOC_MAPPINGS = [
  {
    sourceDir: "patterns",
    source: "macros.md",
    target: "macros.md",
    title: "Macros",
  },
  {
    sourceDir: "patterns",
    source: "random-data.md",
    target: "random_data.md",
    title: "Random Data",
  },
  {
    sourceDir: "patterns",
    source: "global-hooks.md",
    target: "global_hooks.md",
    title: "Global Hooks",
  },
  {
    sourceDir: "patterns",
    source: "variables.md",
    target: "variables.md",
    title: "Variables",
  },
  {
    sourceDir: "guides",
    source: "api-testing.md",
    target: "api_testing.md",
    title: "API Testing",
  },
  {
    sourceDir: "guides",
    source: "frontend-testing.md",
    target: "frontend_testing.md",
    title: "Frontend Testing",
  },
] as const;

const LINK_REWRITES: Array<[RegExp, string]> = [
  [/\]\(\/docs\/patterns\/variables\)/g, "](./variables.md)"],
  [/\]\(\/docs\/patterns\/random-data\)/g, "](./random_data.md)"],
  [/\]\(\/docs\/patterns\/global-hooks\)/g, "](./global_hooks.md)"],
  [/\]\(\/docs\/patterns\/macros\)/g, "](./macros.md)"],
];

function parseFrontmatter(content: string): {
  frontmatter: Record<string, unknown>;
  body: string;
} {
  const match = content.match(/^---\r?\n([\s\S]*?)\r?\n---\r?\n?([\s\S]*)$/);
  if (!match) {
    throw new Error("Missing YAML frontmatter");
  }
  return {
    frontmatter: parseYaml(match[1]) as Record<string, unknown>,
    body: match[2],
  };
}

function rewriteLinks(body: string): string {
  let result = body;
  for (const [pattern, replacement] of LINK_REWRITES) {
    result = result.replace(pattern, replacement);
  }
  return result;
}

function transformDoc(
  sourcePath: string,
  titleOverride: string,
): { frontmatter: { title: string; description: string }; body: string } {
  const raw = fs.readFileSync(sourcePath, "utf8");
  const { frontmatter, body } = parseFrontmatter(raw);
  const description = String(frontmatter.description ?? "").trim();
  if (!description) {
    throw new Error(`Missing description in ${sourcePath}`);
  }

  return {
    frontmatter: {
      title: titleOverride,
      description,
    },
    body: rewriteLinks(body),
  };
}

function formatDoc(
  frontmatter: { title: string; description: string },
  body: string,
): string {
  const yamlBlock = stringifyYaml(frontmatter).trimEnd();
  return `---\n${yamlBlock}\n---\n\n${body.replace(/^\n+/, "")}`;
}

export function syncFrameworkDocs(): void {
  fs.mkdirSync(OUTPUT_DIR, { recursive: true });

  for (const { sourceDir, source, target, title } of DOC_MAPPINGS) {
    const sourcePath = path.join(DOCS_CONTENT_ROOT, sourceDir, source);
    if (!fs.existsSync(sourcePath)) {
      throw new Error(`Source file not found: ${sourcePath}`);
    }

    const transformed = transformDoc(sourcePath, title);
    const outputPath = path.join(OUTPUT_DIR, target);
    fs.writeFileSync(
      outputPath,
      formatDoc(transformed.frontmatter, transformed.body),
      "utf8",
    );
  }

  const written = fs
    .readdirSync(OUTPUT_DIR)
    .filter((name) => name.endsWith(".md"));
  if (written.length !== DOC_MAPPINGS.length) {
    throw new Error(
      `Expected ${DOC_MAPPINGS.length} docs, found ${written.length}`,
    );
  }
}

if (import.meta.url === `file://${process.argv[1]}`) {
  syncFrameworkDocs();
  console.log(`Synced ${DOC_MAPPINGS.length} framework docs to ${OUTPUT_DIR}`);
}
