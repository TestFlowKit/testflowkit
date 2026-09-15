import { Then, When } from '@cucumber/cucumber';
import { resolveElement } from '../browser/resolveElement.js';
import type { TestFlowKitWorld } from '../world.js';
import { substituteVariables } from './substitute.js';

// Registered once each under When/Then: Cucumber matches step text against every
// registered pattern regardless of the Given/When/Then keyword used in the .feature
// file, so these also work written with any keyword without a second registration.
When('I store the value {string} into {string} variable', function (this: TestFlowKitWorld, value: string, name: string) {
  this.variables.set(name, substituteVariables(value, this.variables));
});

When(
  'I store the content of {string} into {string} variable',
  async function (this: TestFlowKitWorld, elementName: string, varName: string) {
    const element = await resolveElement(this, elementName);
    const text = (await element.textContent()) ?? '';
    this.variables.set(varName, text.trim());
  },
);

When('I display the value of variable {string}', function (this: TestFlowKitWorld, name: string) {
  console.log(`[testflowkit] ${name} = ${this.variables.get(name) ?? ''}`);
});

Then('the variable {string} should contain {string}', function (this: TestFlowKitWorld, name: string, expected: string) {
  const actual = this.variables.get(name) ?? '';
  const value = substituteVariables(expected, this.variables);
  if (!actual.includes(value)) {
    throw new Error(`Expected variable "${name}" ("${actual}") to contain "${value}"`);
  }
});

Then(
  'the variable {string} should not contain {string}',
  function (this: TestFlowKitWorld, name: string, expected: string) {
    const actual = this.variables.get(name) ?? '';
    const value = substituteVariables(expected, this.variables);
    if (actual.includes(value)) {
      throw new Error(`Expected variable "${name}" ("${actual}") not to contain "${value}"`);
    }
  },
);
