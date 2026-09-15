import type { Locator, Page } from 'playwright';
import { getElementSelectors, type Selector } from '../config/elements.js';
import type { TestFlowKitWorld } from '../world.js';

function toPlaywrightSelector(selector: Selector): string {
  return selector.type === 'xpath' ? `xpath=${selector.value}` : selector.value;
}

/**
 * Races every fallback selector for `elementName` (first one to attach wins), mirroring
 * the concurrent selector race in the Go browser driver.
 */
export async function resolveElement(world: TestFlowKitWorld, elementName: string): Promise<Locator> {
  const elements = world.config.frontend?.elements ?? {};
  const selectors = getElementSelectors(elements, world.currentPageName, elementName);
  if (selectors.length === 0) {
    throw new Error(`No selector configured for element "${elementName}"`);
  }

  return raceSelectors(world.page, selectors, world.browserSettings.timeout);
}

async function raceSelectors(page: Page, selectors: Selector[], timeout: number): Promise<Locator> {
  const attempts = selectors.map(async (selector) => {
    const locator = page.locator(toPlaywrightSelector(selector));
    await locator.first().waitFor({ state: 'attached', timeout });
    return locator.first();
  });

  try {
    return await Promise.any(attempts);
  } catch {
    throw new Error(`Element not found using any of the configured selectors: ${JSON.stringify(selectors)}`);
  }
}

/**
 * Counts matches for `elementName`, using the first configured selector that matches at
 * least one element (mirrors the Go "active selector" fallback used for element counts).
 */
export async function countElements(world: TestFlowKitWorld, elementName: string): Promise<number> {
  const elements = world.config.frontend?.elements ?? {};
  const selectors = getElementSelectors(elements, world.currentPageName, elementName);
  if (selectors.length === 0) {
    throw new Error(`No selector configured for element "${elementName}"`);
  }

  for (const selector of selectors) {
    const count = await world.page.locator(toPlaywrightSelector(selector)).count();
    if (count > 0) {
      return count;
    }
  }
  return 0;
}
