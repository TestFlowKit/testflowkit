import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";

export function resolveFrameworkDocsDir(): string {
  const here = path.dirname(fileURLToPath(import.meta.url));
  const candidates = [
    path.resolve(here, "./docs/features"),
    path.resolve(here, "../docs/features"),
    path.resolve(here, "../../docs/features"),
  ];

  for (const candidate of candidates) {
    if (fs.existsSync(candidate)) {
      return candidate;
    }
  }

  throw new Error(
    `Framework docs directory not found. Tried: ${candidates.join(", ")}`,
  );
}
