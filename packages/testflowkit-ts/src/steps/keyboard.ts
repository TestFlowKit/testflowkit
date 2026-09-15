import { When } from '@cucumber/cucumber';
import type { TestFlowKitWorld } from '../world.js';

const KEY_MAP: Record<string, string> = {
  Enter: 'Enter',
  Tab: 'Tab',
  Delete: 'Delete',
  Escape: 'Escape',
  Space: 'Space',
  'Arrow Up': 'ArrowUp',
  'Arrow Right': 'ArrowRight',
  'Arrow Down': 'ArrowDown',
  'Arrow Left': 'ArrowLeft',
};

// Registered once under When: Cucumber matches step text against every registered
// pattern regardless of the Given/When/Then keyword used in the .feature file, so this
// also works written as Given/Then/And/But without a second registration.
When('the user presses the {string} key', async function (this: TestFlowKitWorld, key: string) {
  const playwrightKey = KEY_MAP[key];
  if (!playwrightKey) {
    throw new Error(`Unsupported key "${key}". Supported keys: ${Object.keys(KEY_MAP).join(', ')}`);
  }
  await this.page.keyboard.press(playwrightKey);
});
