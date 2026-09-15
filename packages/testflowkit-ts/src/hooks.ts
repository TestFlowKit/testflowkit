import { resolve } from 'node:path';
import { After, AfterAll, Before, BeforeAll, Status } from '@cucumber/cucumber';
import { chromium, type Browser } from 'playwright';
import { loadConfig } from './config/loader.js';
import { loadBrowserSettings, type BrowserSettings } from './config/playwright-config.js';
import type { TestFlowKitConfig } from './config/types.js';
import type { SharedFixtures, TestFlowKitWorld } from './world.js';

const CONFIG_PATH = resolve(process.cwd(), process.env.TESTFLOWKIT_CONFIG ?? 'testflowkit.yml');
const PLAYWRIGHT_CONFIG_PATH = resolve(
  process.cwd(),
  process.env.TESTFLOWKIT_PLAYWRIGHT_CONFIG ?? 'playwright.config.ts',
);

let browser: Browser;
let config: TestFlowKitConfig;
let browserSettings: BrowserSettings;

BeforeAll(async function () {
  config = loadConfig(CONFIG_PATH);
  browserSettings = await loadBrowserSettings(PLAYWRIGHT_CONFIG_PATH);
  browser = await chromium.launch({ headless: browserSettings.headless });
});

Before(async function (this: TestFlowKitWorld) {
  const fixtures: SharedFixtures = { browser, browserSettings, config };
  this.fixtures = fixtures;

  this.context = await browser.newContext({
    baseURL: browserSettings.baseURL,
    viewport: browserSettings.viewport,
    locale: browserSettings.locale,
    timezoneId: browserSettings.timezoneId,
    userAgent: browserSettings.userAgent,
  });
  this.contexts = [this.context];
  this.page = await this.context.newPage();
  this.page.setDefaultTimeout(browserSettings.timeout);
  this.currentPageName = '';
  this.variables = new Map();
});

After(async function (this: TestFlowKitWorld, { result }) {
  if (result?.status === Status.FAILED) {
    const screenshot = await this.page.screenshot();
    await this.attach(screenshot, 'image/png');
  }
  await Promise.all(this.contexts.map((context) => context.close()));
});

AfterAll(async function () {
  await browser.close();
});
