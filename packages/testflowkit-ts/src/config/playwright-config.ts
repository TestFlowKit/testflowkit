import { pathToFileURL } from 'node:url';
import type { PlaywrightTestConfig } from '@playwright/test';

export interface BrowserSettings {
  baseURL?: string;
  headless: boolean;
  viewport: { width: number; height: number } | null;
  locale?: string;
  timezoneId?: string;
  userAgent?: string;
  timeout: number;
}

const DEFAULT_TIMEOUT_MS = 30_000;

/**
 * Dynamically imports the consumer's playwright.config.ts (a real, type-checked
 * Playwright config) and extracts the subset of `use`/`timeout` that TestFlowKit's
 * browser hooks need. Browser-level settings live here, not in testflowkit.yml.
 */
export async function loadBrowserSettings(configPath: string): Promise<BrowserSettings> {
  const module = (await import(pathToFileURL(configPath).href)) as {
    default: PlaywrightTestConfig;
  };
  const config = module.default;
  const use = config.use ?? {};

  return {
    baseURL: use.baseURL,
    headless: use.headless ?? true,
    viewport: use.viewport ?? null,
    locale: use.locale,
    timezoneId: use.timezoneId,
    userAgent: use.userAgent,
    timeout: config.timeout ?? DEFAULT_TIMEOUT_MS,
  };
}
