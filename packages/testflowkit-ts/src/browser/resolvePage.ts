import { snakeCase } from '../config/elements.js';
import type { TestFlowKitWorld } from '../world.js';

/**
 * Resolves a logical page name to a URL: absolute http(s) values are used verbatim,
 * otherwise the path is joined onto the Playwright config's baseURL. Mirrors
 * config.GetFrontendURL in the Go version.
 */
export function resolvePageUrl(world: TestFlowKitWorld, pageName: string): string {
  const pages = world.config.frontend?.pages ?? {};
  const path = pages[snakeCase(pageName)];

  if (path === undefined) {
    if (!world.browserSettings.baseURL) {
      throw new Error(`Page "${pageName}" is not defined and no baseURL is configured`);
    }
    return world.browserSettings.baseURL;
  }

  if (/^https?:\/\//.test(path)) {
    return path;
  }

  if (!world.browserSettings.baseURL) {
    throw new Error(`Page "${pageName}" resolves to a relative path but no baseURL is configured`);
  }

  return new URL(path.replace(/^\//, ''), world.browserSettings.baseURL.replace(/\/?$/, '/')).toString();
}
