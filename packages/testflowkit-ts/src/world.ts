import { setWorldConstructor, World as CucumberWorld, type IWorldOptions } from '@cucumber/cucumber';
import type { Browser, BrowserContext, Page } from 'playwright';
import type { BrowserSettings } from './config/playwright-config.js';
import type { TestFlowKitConfig } from './config/types.js';

export interface SharedFixtures {
  browser: Browser;
  browserSettings: BrowserSettings;
  config: TestFlowKitConfig;
}

export class TestFlowKitWorld extends CucumberWorld {
  context!: BrowserContext;
  page!: Page;
  /** Every browser context opened in this scenario (the initial one, plus any private tabs), used to enumerate open windows/tabs for switching steps. */
  contexts: BrowserContext[] = [];
  currentPageName = '';
  variables = new Map<string, string>();
  fixtures!: SharedFixtures;

  constructor(options: IWorldOptions) {
    super(options);
  }

  get config(): TestFlowKitConfig {
    return this.fixtures.config;
  }

  get browserSettings(): BrowserSettings {
    return this.fixtures.browserSettings;
  }
}

setWorldConstructor(TestFlowKitWorld);
