// Registers hooks and step definitions as a side effect of importing this module.
import './hooks.js';
import './steps/navigation.js';
import './steps/form.js';
import './steps/mouse.js';
import './steps/keyboard.js';
import './steps/assertions.js';
import './steps/visual.js';
import './steps/table.js';
import './variables/steps.js';

export { TestFlowKitWorld } from './world.js';
export type { TestFlowKitConfig, FrontendElements, FrontendPages } from './config/types.js';
export { loadConfig } from './config/loader.js';
export { loadBrowserSettings } from './config/playwright-config.js';
