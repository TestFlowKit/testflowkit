/** Maps group names (page names, plus "common") to element names and their selector fallbacks. */
export type FrontendElements = Record<string, Record<string, string[]>>;

/** Maps page names to URL paths appended to the Playwright config's baseURL, or absolute URLs. */
export type FrontendPages = Record<string, string>;

export interface FrontendYamlConfig {
  elements: FrontendElements;
  pages: FrontendPages;
}

export interface Settings {
  gherkinLocation: string;
  tags?: string;
}

export interface FilesConfig {
  baseDirectory: string;
  definitions: Record<string, string>;
}

export interface TestFlowKitConfig {
  settings: Settings;
  env: Record<string, string>;
  frontend?: FrontendYamlConfig;
  files?: FilesConfig;
}
