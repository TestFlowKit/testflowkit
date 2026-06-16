import fs from "node:fs";
import path from "node:path";

export function listFeatureFiles(
  featuresGlob: string,
  baseDir: string
): string[] {
  const absGlob = path.isAbsolute(featuresGlob)
    ? featuresGlob
    : path.join(baseDir, featuresGlob);

  return fs.globSync(absGlob);
}
