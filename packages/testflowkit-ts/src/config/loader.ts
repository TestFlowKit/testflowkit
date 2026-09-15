import { readFileSync } from 'node:fs';
import { parse as parseYaml } from 'yaml';
import { substituteEnv } from '../variables/substitute.js';
import type { FrontendElements, FrontendPages, TestFlowKitConfig } from './types.js';

interface RawYamlConfig {
  settings?: { gherkin_location?: string; tags?: string };
  env?: Record<string, string>;
  frontend?: {
    elements?: FrontendElements;
    pages?: FrontendPages;
  };
  files?: {
    base_directory?: string;
    definitions?: Record<string, string>;
  };
}

/**
 * Loads testflowkit.yml: parses the `env` block, then substitutes `{{ env.X }}` over the
 * raw file text (so any value, including nested page/element entries, can reference env
 * vars) before the final YAML parse — matching the Go loader's two-pass approach.
 */
export function loadConfig(path: string): TestFlowKitConfig {
  const raw = readFileSync(path, 'utf-8');
  const prelim = parseYaml(raw) as RawYamlConfig;
  const env = { ...(prelim.env ?? {}), ...(process.env as Record<string, string>) };

  const substituted = substituteEnv(raw, env);
  const parsed = parseYaml(substituted) as RawYamlConfig;

  if (!parsed.settings?.gherkin_location) {
    throw new Error(`${path}: settings.gherkin_location is required`);
  }

  return {
    settings: {
      gherkinLocation: parsed.settings.gherkin_location,
      tags: parsed.settings.tags,
    },
    env,
    frontend: parsed.frontend
      ? {
          elements: parsed.frontend.elements ?? {},
          pages: parsed.frontend.pages ?? {},
        }
      : undefined,
    files: parsed.files
      ? {
          baseDirectory: parsed.files.base_directory ?? '',
          definitions: parsed.files.definitions ?? {},
        }
      : undefined,
  };
}
