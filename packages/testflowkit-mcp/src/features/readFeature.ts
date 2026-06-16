import fs from "node:fs";
import { resolveFeaturePath } from "./writeFeature.js";

export function readFeatureFile(
  featurePath: string,
  featuresGlob: string,
  baseDir: string,
): string {
  const abs = resolveFeaturePath(featurePath, featuresGlob, baseDir);
  if (!fs.existsSync(abs)) {
    throw new Error(`Feature file not found: "${abs}"`);
  }
  return fs.readFileSync(abs, "utf-8");
}
