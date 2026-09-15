import { Then, When, type DataTable } from '@cucumber/cucumber';
import { countElements, resolveElement } from '../browser/resolveElement.js';
import { suffixWithUnderscore } from '../config/elements.js';
import type { TestFlowKitWorld } from '../world.js';
import { substituteVariables } from '../variables/substitute.js';

const TAG_BY_ELEMENT_TYPE: Record<string, string> = {
  link: 'a',
  button: 'button',
  element: '*',
};

Then(
  /^the user should see a (link|button|element) which contains "([^"]*)"$/,
  async function (this: TestFlowKitWorld, elementType: string, text: string) {
    const tag = TAG_BY_ELEMENT_TYPE[elementType]!;
    const value = substituteVariables(text, this.variables);
    const element = this.page.locator(tag).filter({ hasText: value }).first();
    if (!(await element.isVisible().catch(() => false))) {
      throw new Error(`No ${elementType} is visible with text "${value}"`);
    }
  },
);

Then(
  'the user should see {int} {string} elements on the page',
  async function (this: TestFlowKitWorld, expectedCount: number, elementName: string) {
    const actualCount = await countElements(this, elementName);
    if (actualCount !== expectedCount) {
      throw new Error(`${expectedCount} ${elementName} expected but ${actualCount} ${elementName} found`);
    }
  },
);

Then(
  'the user should see {string} details on the page',
  async function (this: TestFlowKitWorld, elementName: string, table: DataTable) {
    const container = await resolveElement(this, suffixWithUnderscore(elementName, 'details'));
    const text = ((await container.textContent()) ?? '').replace(/\s*\n\s*/g, ' ').trim();

    const errors: string[] = [];
    for (const [name, rawValue] of Object.entries(table.rowsHash())) {
      const value = substituteVariables(rawValue, this.variables);
      if (!text.includes(value)) {
        errors.push(`${elementName} ${name} not found`);
      }
    }
    if (errors.length > 0) {
      throw new Error(errors.join('\n'));
    }
  },
);

When('the user scrolls to the {string} element', async function (this: TestFlowKitWorld, elementName: string) {
  const element = await resolveElement(this, suffixWithUnderscore(elementName, 'element'));
  await element.scrollIntoViewIfNeeded();
});
