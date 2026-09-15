import { join } from 'node:path';
import type { TestFlowKitConfig } from './types.js';

/** Resolves logical file names to filesystem paths under files.base_directory. */
export function resolveFilePaths(config: TestFlowKitConfig, fileNames: string[]): string[] {
  const files = config.files;
  if (!files) {
    throw new Error('No file definitions configured');
  }

  const notFound: string[] = [];
  const paths = fileNames.map((fileName) => {
    const relativePath = files.definitions[fileName];
    if (relativePath === undefined) {
      notFound.push(fileName);
      return '';
    }
    return join(files.baseDirectory, relativePath);
  });

  if (notFound.length > 0) {
    throw new Error(`Files do not exist: ${notFound.join(', ')}`);
  }

  return paths;
}
