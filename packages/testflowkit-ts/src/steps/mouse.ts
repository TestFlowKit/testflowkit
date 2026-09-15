import { When } from '@cucumber/cucumber';
import { resolveElement } from '../browser/resolveElement.js';
import { suffixWithUnderscore } from '../config/elements.js';
import type { TestFlowKitWorld } from '../world.js';

// Registered once each under When: Cucumber matches step text against every registered
// pattern regardless of the Given/When/Then keyword used in the .feature file, so these
// also work written as Given/Then/And/But without a second registration.
When('the user clicks the {string} button', async function (this: TestFlowKitWorld, label: string) {
  const element = await resolveElement(this, suffixWithUnderscore(label, 'button'));
  await element.click();
});

When('the user clicks the {string} element', async function (this: TestFlowKitWorld, label: string) {
  const element = await resolveElement(this, suffixWithUnderscore(label, 'element'));
  await element.click();
});

When('the user clicks the {string} link', async function (this: TestFlowKitWorld, label: string) {
  const element = await resolveElement(this, suffixWithUnderscore(label, 'link'));
  await element.click();
});

When('the user hovers on {string}', async function (this: TestFlowKitWorld, label: string) {
  const element = await resolveElement(this, label);
  await element.hover();
});

When('the user double clicks on {string}', async function (this: TestFlowKitWorld, label: string) {
  const element = await resolveElement(this, label);
  await element.dblclick();
});

When('the user right clicks on {string}', async function (this: TestFlowKitWorld, label: string) {
  const element = await resolveElement(this, label);
  await element.click({ button: 'right' });
});

When(
  'the user clicks on {string} which contains {string}',
  async function (this: TestFlowKitWorld, _label: string, text: string) {
    const element = this.page.getByText(text).first();
    await element.click();
  },
);

When(
  'the user double clicks on {string} which contains {string}',
  async function (this: TestFlowKitWorld, _label: string, text: string) {
    const element = this.page.getByText(text).first();
    await element.dblclick();
  },
);

When(
  'the user right clicks on {string} which contains {string}',
  async function (this: TestFlowKitWorld, _label: string, text: string) {
    const element = this.page.getByText(text).first();
    await element.click({ button: 'right' });
  },
);

When(
  'the user hovers on {string} which contains {string}',
  async function (this: TestFlowKitWorld, _label: string, text: string) {
    const element = this.page.getByText(text).first();
    await element.hover();
  },
);
