import fs from "node:fs";
import path from "node:path";
import { Logger } from "../logger.js";
import { parseFrontmatter } from "./parseFrontmatter.js";
import { resolveFrameworkDocsDir } from "./resolveDocsDir.js";
import { titleToSlug } from "./slug.js";
import type { FrameworkDocEntry, FrameworkDocsIndex } from "./types.js";

const EXPECTED_DOC_COUNT = 6;

export function buildFrameworkDocsIndex(
  docsDir = resolveFrameworkDocsDir(),
): FrameworkDocsIndex {
  if (!fs.existsSync(docsDir)) {
    throw new Error(`Framework docs directory not found: ${docsDir}`);
  }

  const files = fs
    .readdirSync(docsDir)
    .filter((name) => name.endsWith(".md"))
    .sort();

  if (files.length !== EXPECTED_DOC_COUNT) {
    throw new Error(
      `Expected ${EXPECTED_DOC_COUNT} framework docs, found ${files.length} in ${docsDir}`,
    );
  }

  const index: FrameworkDocsIndex = new Map();

  for (const fileName of files) {
    const filePath = path.join(docsDir, fileName);
    const content = fs.readFileSync(filePath, "utf8");
    const { frontmatter } = parseFrontmatter(content);
    const slug = titleToSlug(frontmatter.title);
    const fileSlug = fileName.replace(/\.md$/, "");

    if (fileSlug !== slug) {
      Logger.log(
        `Framework doc slug mismatch: ${fileName} -> title slug "${slug}"`,
      );
    }

    if (index.has(slug)) {
      throw new Error(`Duplicate framework doc slug: ${slug}`);
    }

    index.set(slug, {
      slug,
      title: frontmatter.title,
      description: frontmatter.description,
      filePath,
      content,
    });
  }

  return index;
}

export function listFrameworkDocs(
  index: FrameworkDocsIndex,
): FrameworkDocEntry[] {
  return [...index.values()].sort((a, b) => a.slug.localeCompare(b.slug));
}

export function getFrameworkDoc(
  index: FrameworkDocsIndex,
  slug: string,
): FrameworkDocEntry | undefined {
  return index.get(slug);
}

export function formatFrameworkDocsList(index: FrameworkDocsIndex): string {
  const lines = ["# Framework Documentation", ""];
  for (const doc of listFrameworkDocs(index)) {
    lines.push(
      `- **${doc.slug}** — ${doc.title}: ${doc.description}`,
      `  URI: docs://framework/features/${doc.slug}`,
      "",
    );
  }
  return lines.join("\n").trimEnd();
}
