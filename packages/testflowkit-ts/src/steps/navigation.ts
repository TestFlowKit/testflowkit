import { Given, Then, When } from '@cucumber/cucumber';
import { resolvePageUrl } from '../browser/resolvePage.js';
import { snakeCase } from '../config/elements.js';
import { allPages, parseDurationMs, waitForNewPage } from '../browser/windows.js';
import type { TestFlowKitWorld } from '../world.js';
import { substituteVariables } from '../variables/substitute.js';

// Registered once under Given: Cucumber matches step text against all registered
// patterns regardless of the Given/When/Then keyword used in the .feature file, so a
// step can be written with any keyword without needing a second registration here.
Given('the user goes to the {string} page', async function (this: TestFlowKitWorld, pageName: string) {
  const url = resolvePageUrl(this, substituteVariables(pageName, this.variables));
  await this.page.goto(url);
  this.currentPageName = snakeCase(pageName);
});

When('the user navigate to the URL {string}', async function (this: TestFlowKitWorld, url: string) {
  await this.page.goto(substituteVariables(url, this.variables));
  this.currentPageName = '';
});

Then('the user should be navigated to the {string} page', async function (this: TestFlowKitWorld, pageName: string) {
  const expected = resolvePageUrl(this, substituteVariables(pageName, this.variables));
  const actual = this.page.url();
  if (!actual.startsWith(expected)) {
    throw new Error(`Expected to be on "${expected}" but was on "${actual}"`);
  }
});

When('the user navigates back', async function (this: TestFlowKitWorld) {
  await this.page.goBack();
});

When('the user refreshes the page', async function (this: TestFlowKitWorld) {
  await this.page.reload();
});

When('the user waits for {int} seconds', async function (this: TestFlowKitWorld, seconds: number) {
  await this.page.waitForTimeout(seconds * 1000);
});

Given('the user is on the homepage', async function (this: TestFlowKitWorld) {
  const url = resolvePageUrl(this, 'homepage');
  await this.page.goto(url);
  this.currentPageName = 'homepage';
});

When('the user opens a new browser tab', async function (this: TestFlowKitWorld) {
  this.page = await this.context.newPage();
});

When('the user opens a new private browser tab', async function (this: TestFlowKitWorld) {
  const context = await this.fixtures.browser.newContext({
    baseURL: this.browserSettings.baseURL,
    viewport: this.browserSettings.viewport,
    locale: this.browserSettings.locale,
    timezoneId: this.browserSettings.timezoneId,
    userAgent: this.browserSettings.userAgent,
  });
  this.contexts.push(context);
  this.context = context;
  this.page = await context.newPage();
});

When(
  'the user waits for a new window to open within {string}',
  async function (this: TestFlowKitWorld, waitTime: string) {
    await waitForNewPage(this, parseDurationMs(waitTime));
  },
);

When('the user switches to the newly opened window', async function (this: TestFlowKitWorld) {
  this.page = await waitForNewPage(this, this.browserSettings.timeout);
});

When('the user switches to the most recently opened window', async function (this: TestFlowKitWorld) {
  const pages = allPages(this);
  if (pages.length < 2) {
    throw new Error(`No additional windows found to switch to (only ${pages.length} window open)`);
  }
  this.page = pages.at(-1)!;
});

When('the user switches back to the original window', async function (this: TestFlowKitWorld) {
  const pages = allPages(this);
  if (pages.length < 2) {
    throw new Error('Only one window is open, no original window to switch back to');
  }
  this.page = pages[0]!;
  await this.page.bringToFront();
});
