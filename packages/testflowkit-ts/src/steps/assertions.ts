import { Then } from '@cucumber/cucumber';
import type { Locator } from 'playwright';
import { resolveElement } from '../browser/resolveElement.js';
import { suffixWithUnderscore } from '../config/elements.js';
import type { TestFlowKitWorld } from '../world.js';
import { substituteVariables } from '../variables/substitute.js';

/** True when every label in `labels` is currently selected in the `<select>` element. */
async function hasOptionsSelected(select: Locator, labels: string[]): Promise<boolean> {
  const selectedLabels: string[] = await select.evaluate((el) => {
    const selectEl = el as unknown as { selectedOptions: ArrayLike<{ label: string }> };
    return Array.from(selectEl.selectedOptions).map((option) => option.label);
  });
  return labels.every((label) => selectedLabels.includes(label));
}

// Registered once each under Then: Cucumber matches step text against every registered
// pattern regardless of the Given/When/Then keyword used in the .feature file, so these
// also work written as Given/When/And/But without a second registration.
Then('the {string} should be visible', async function (this: TestFlowKitWorld, elementName: string) {
  const element = await resolveElement(this, elementName);
  if (!(await element.isVisible())) {
    throw new Error(`Expected "${elementName}" to be visible`);
  }
});

Then('the {string} should not be visible', async function (this: TestFlowKitWorld, elementName: string) {
  const element = await resolveElement(this, elementName);
  if (await element.isVisible()) {
    throw new Error(`Expected "${elementName}" not to be visible`);
  }
});

Then('the {string} should exist', async function (this: TestFlowKitWorld, elementName: string) {
  await resolveElement(this, elementName);
});

Then('the {string} should not exist', async function (this: TestFlowKitWorld, elementName: string) {
  let found = true;
  try {
    await resolveElement(this, elementName);
  } catch {
    found = false;
  }
  if (found) {
    throw new Error(`Expected "${elementName}" not to exist`);
  }
});

Then(
  'the {string} should contain the text {string}',
  async function (this: TestFlowKitWorld, elementName: string, expected: string) {
    const element = await resolveElement(this, elementName);
    const actual = (await element.textContent()) ?? '';
    const value = substituteVariables(expected, this.variables);
    if (!actual.includes(value)) {
      throw new Error(`Expected "${elementName}" ("${actual}") to contain "${value}"`);
    }
  },
);

Then(
  'the {string} should not contain the text {string}',
  async function (this: TestFlowKitWorld, elementName: string, expected: string) {
    const element = await resolveElement(this, elementName);
    const actual = (await element.textContent()) ?? '';
    const value = substituteVariables(expected, this.variables);
    if (actual.includes(value)) {
      throw new Error(`Expected "${elementName}" ("${actual}") not to contain "${value}"`);
    }
  },
);

Then('the current URL should contain {string}', async function (this: TestFlowKitWorld, expected: string) {
  const value = substituteVariables(expected, this.variables);
  if (!this.page.url().includes(value)) {
    throw new Error(`Expected current URL "${this.page.url()}" to contain "${value}"`);
  }
});

Then('the current URL should not contain {string}', async function (this: TestFlowKitWorld, expected: string) {
  const value = substituteVariables(expected, this.variables);
  if (this.page.url().includes(value)) {
    throw new Error(`Expected current URL "${this.page.url()}" not to contain "${value}"`);
  }
});

Then('the page title should be {string}', async function (this: TestFlowKitWorld, expected: string) {
  const value = substituteVariables(expected, this.variables);
  const actual = await this.page.title();
  if (actual !== value) {
    throw new Error(`Expected page title to be "${value}" but was "${actual}"`);
  }
});

Then('the page title should not be {string}', async function (this: TestFlowKitWorld, expected: string) {
  const value = substituteVariables(expected, this.variables);
  const actual = await this.page.title();
  if (actual === value) {
    throw new Error(`Expected page title not to be "${value}"`);
  }
});

Then(
  /^the "([^"]*)" checkbox should be (checked|unchecked)$/,
  async function (this: TestFlowKitWorld, label: string, state: 'checked' | 'unchecked') {
    const element = await resolveElement(this, suffixWithUnderscore(label, 'checkbox'));
    const isChecked = await element.isChecked();
    if (state === 'checked' && !isChecked) {
      throw new Error(`Expected "${label}" checkbox to be checked`);
    }
    if (state === 'unchecked' && isChecked) {
      throw new Error(`Expected "${label}" checkbox to be unchecked`);
    }
  },
);

Then(
  /^the "([^"]*)" radio button should be (selected|unselected)$/,
  async function (this: TestFlowKitWorld, label: string, state: 'selected' | 'unselected') {
    const element = await resolveElement(this, suffixWithUnderscore(label, 'radio_button'));
    const isChecked = await element.isChecked();
    if (state === 'selected' && !isChecked) {
      throw new Error(`Expected "${label}" radio button to be selected`);
    }
    if (state === 'unselected' && isChecked) {
      throw new Error(`Expected "${label}" radio button to be unselected`);
    }
  },
);

Then(
  'the value of the {string} field should be {string}',
  async function (this: TestFlowKitWorld, fieldLabel: string, expected: string) {
    const element = await resolveElement(this, suffixWithUnderscore(fieldLabel, 'field'));
    const actual = await element.inputValue();
    const value = substituteVariables(expected, this.variables);
    if (actual !== value) {
      throw new Error(`Expected "${fieldLabel}" field to be "${value}" but was "${actual}"`);
    }
  },
);

Then(
  'the value of the {string} field should not be {string}',
  async function (this: TestFlowKitWorld, fieldLabel: string, expected: string) {
    const element = await resolveElement(this, suffixWithUnderscore(fieldLabel, 'field'));
    const actual = await element.inputValue();
    const value = substituteVariables(expected, this.variables);
    if (actual === value) {
      throw new Error(`Expected "${fieldLabel}" field value not to be "${value}"`);
    }
  },
);

Then(
  'the {string} dropdown should have {string} selected',
  async function (this: TestFlowKitWorld, dropdownLabel: string, optionLabels: string) {
    const element = await resolveElement(this, suffixWithUnderscore(dropdownLabel, 'dropdown'));
    const labels = optionLabels.split(',').map((label) => substituteVariables(label.trim(), this.variables));
    const isSelected = await hasOptionsSelected(element, labels);
    if (!isSelected) {
      throw new Error(`Expected "${optionLabels}" to be selected in "${dropdownLabel}" dropdown`);
    }
  },
);

Then(
  'the {string} dropdown should not have {string} selected',
  async function (this: TestFlowKitWorld, dropdownLabel: string, optionLabels: string) {
    const element = await resolveElement(this, suffixWithUnderscore(dropdownLabel, 'dropdown'));
    const labels = optionLabels.split(',').map((label) => substituteVariables(label.trim(), this.variables));
    const isSelected = await hasOptionsSelected(element, labels);
    if (isSelected) {
      throw new Error(`Expected "${optionLabels}" not to be selected in "${dropdownLabel}" dropdown`);
    }
  },
);

Then(
  'the text of the {string} element should be exactly {string}',
  async function (this: TestFlowKitWorld, elementName: string, expected: string) {
    const element = await resolveElement(this, elementName);
    const actual = ((await element.textContent()) ?? '').trim();
    const value = substituteVariables(expected, this.variables);
    if (actual !== value) {
      throw new Error(`Expected "${elementName}" text to be exactly "${value}" but was "${actual}"`);
    }
  },
);

Then(
  'the text of the {string} element should not be exactly {string}',
  async function (this: TestFlowKitWorld, elementName: string, expected: string) {
    const element = await resolveElement(this, elementName);
    const actual = ((await element.textContent()) ?? '').trim();
    const value = substituteVariables(expected, this.variables);
    if (actual === value) {
      throw new Error(`Expected "${elementName}" text not to be exactly "${value}"`);
    }
  },
);

Then(
  'the {string} attribute of the {string} element should be {string}',
  async function (this: TestFlowKitWorld, attributeName: string, elementName: string, expected: string) {
    const element = await resolveElement(this, elementName);
    const actual = await element.getAttribute(attributeName);
    const value = substituteVariables(expected, this.variables);
    if (actual !== value) {
      throw new Error(`Expected "${attributeName}" of "${elementName}" to be "${value}" but was "${actual}"`);
    }
  },
);

Then(
  'the {string} attribute of the {string} element should not be {string}',
  async function (this: TestFlowKitWorld, attributeName: string, elementName: string, expected: string) {
    const element = await resolveElement(this, elementName);
    const actual = await element.getAttribute(attributeName);
    const value = substituteVariables(expected, this.variables);
    if (actual === value) {
      throw new Error(`Expected "${attributeName}" of "${elementName}" not to be "${value}"`);
    }
  },
);
