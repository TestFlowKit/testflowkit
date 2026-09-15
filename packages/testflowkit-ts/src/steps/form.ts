import { When } from '@cucumber/cucumber';
import { resolveElement } from '../browser/resolveElement.js';
import { suffixWithUnderscore } from '../config/elements.js';
import { resolveFilePaths } from '../config/files.js';
import type { TestFlowKitWorld } from '../world.js';
import { substituteVariables } from '../variables/substitute.js';

function splitAndTrim(value: string): string[] {
  return value.split(',').map((item) => item.trim());
}

// Registered once each under When: Cucumber matches step text against every registered
// pattern regardless of the Given/When/Then keyword used in the .feature file, so these
// also work written as Given/Then/And/But without a second registration.
When(
  'the user enters {string} into the {string} field',
  async function (this: TestFlowKitWorld, text: string, fieldLabel: string) {
    const element = await resolveElement(this, suffixWithUnderscore(fieldLabel, 'field'));
    await element.fill(substituteVariables(text, this.variables));
  },
);

When('the user clears the {string} field', async function (this: TestFlowKitWorld, fieldLabel: string) {
  const element = await resolveElement(this, suffixWithUnderscore(fieldLabel, 'field'));
  await element.fill('');
});

When('the user checks the {string} checkbox', async function (this: TestFlowKitWorld, label: string) {
  const element = await resolveElement(this, suffixWithUnderscore(label, 'checkbox'));
  await element.check();
});

When('the user unchecks the {string} checkbox', async function (this: TestFlowKitWorld, label: string) {
  const element = await resolveElement(this, suffixWithUnderscore(label, 'checkbox'));
  await element.uncheck();
});

When('the user selects the {string} radio button', async function (this: TestFlowKitWorld, label: string) {
  const element = await resolveElement(this, suffixWithUnderscore(label, 'radio_button'));
  await element.check();
});

When(
  'the user selects the option with text {string} from the {string} dropdown',
  async function (this: TestFlowKitWorld, text: string, dropdownLabel: string) {
    const element = await resolveElement(this, suffixWithUnderscore(dropdownLabel, 'dropdown'));
    await element.selectOption({ label: substituteVariables(text, this.variables) });
  },
);

When(
  'the user selects the option with value {string} from the {string} dropdown',
  async function (this: TestFlowKitWorld, value: string, dropdownLabel: string) {
    const element = await resolveElement(this, suffixWithUnderscore(dropdownLabel, 'dropdown'));
    await element.selectOption({ value: substituteVariables(value, this.variables) });
  },
);

When(
  'the user selects the options with text {string} from the {string} dropdown',
  async function (this: TestFlowKitWorld, texts: string, dropdownLabel: string) {
    const element = await resolveElement(this, suffixWithUnderscore(dropdownLabel, 'dropdown'));
    const labels = splitAndTrim(texts).map((text) => substituteVariables(text, this.variables));
    await element.selectOption(labels.map((label) => ({ label })));
  },
);

When(
  'the user selects the options with values {string} from the {string} dropdown',
  async function (this: TestFlowKitWorld, values: string, dropdownLabel: string) {
    const element = await resolveElement(this, suffixWithUnderscore(dropdownLabel, 'dropdown'));
    const optionValues = splitAndTrim(values).map((value) => substituteVariables(value, this.variables));
    await element.selectOption(optionValues);
  },
);

When(
  'the user selects the option at index {int} from the {string} dropdown',
  async function (this: TestFlowKitWorld, index: number, dropdownLabel: string) {
    const element = await resolveElement(this, suffixWithUnderscore(dropdownLabel, 'dropdown'));
    await element.selectOption({ index });
  },
);

When(
  'the user uploads the {string} file into the {string} field',
  async function (this: TestFlowKitWorld, fileName: string, fieldLabel: string) {
    const [filePath] = resolveFilePaths(this.config, [fileName]);
    const element = await resolveElement(this, suffixWithUnderscore(fieldLabel, 'field'));
    await element.setInputFiles(filePath!);
  },
);

When(
  'the user uploads the {string} files into the {string} field',
  async function (this: TestFlowKitWorld, fileNames: string, fieldLabel: string) {
    const filePaths = resolveFilePaths(this.config, splitAndTrim(fileNames));
    const element = await resolveElement(this, suffixWithUnderscore(fieldLabel, 'field'));
    await element.setInputFiles(filePaths);
  },
);
