import fs from "node:fs";
import path from "node:path";

/**
 * Resolve and security-check a feature file path.
 * - Must be under baseDir (no path traversal).
 * - Must match the features_glob pattern.
 */
export function resolveFeaturePath(
  featurePath: string,
  featuresGlob: string,
  baseDir: string
): string {
  // Resolve against baseDir
  const abs = path.isAbsolute(featurePath)
    ? featurePath
    : path.resolve(baseDir, featurePath);

  // Normalise (collapses ..)
  const normalised = path.normalize(abs);

  // Must be inside baseDir
  const base = path.resolve(baseDir);
  if (!normalised.startsWith(base + path.sep) && normalised !== base) {
    throw new Error(
      `Security: path "${featurePath}" resolves outside the workspace root "${base}".`
    );
  }

  // Must match features_glob
  const rel = path.relative(base, normalised);
  const globPattern = path.isAbsolute(featuresGlob)
    ? path.relative(base, featuresGlob)
    : featuresGlob;

  if (!path.matchesGlob(rel, globPattern)) {
    throw new Error(
      `Path "${rel}" does not match features_glob pattern "${globPattern}".`
    );
  }

  return normalised;
}

export function writeFeatureFile(
  featurePath: string,
  content: string,
  featuresGlob: string,
  baseDir: string,
  createDirs = true
): void {
  const abs = resolveFeaturePath(featurePath, featuresGlob, baseDir);

  if (createDirs) {
    fs.mkdirSync(path.dirname(abs), { recursive: true });
  }

  fs.writeFileSync(abs, content, "utf-8");
}
