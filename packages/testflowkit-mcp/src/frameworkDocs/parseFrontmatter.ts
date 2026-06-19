import { parse as parseYaml } from "yaml";

export type ParsedFrontmatter = {
  title: string;
  description: string;
};

export function parseFrontmatter(content: string): {
  frontmatter: ParsedFrontmatter;
  body: string;
} {
  const match = content.match(/^---\r?\n([\s\S]*?)\r?\n---\r?\n?([\s\S]*)$/);
  if (!match) {
    throw new Error("Missing YAML frontmatter");
  }

  const raw = parseYaml(match[1]) as Record<string, unknown>;
  const title = String(raw.title ?? "").trim();
  const description = String(raw.description ?? "").trim();

  if (!title || !description) {
    throw new Error("Frontmatter must include title and description");
  }

  return {
    frontmatter: { title, description },
    body: match[2],
  };
}
