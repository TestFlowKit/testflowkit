import assert from 'node:assert/strict';
import { test } from 'node:test';
import { getElementSelectors, isElementDefined, parseSelector, snakeCase, suffixWithUnderscore } from './elements.js';

test('snakeCase converts camelCase and spaces', () => {
  assert.equal(snakeCase('loginPage'), 'login_page');
  assert.equal(snakeCase('Login Page'), 'login_page');
  assert.equal(snakeCase('submit'), 'submit');
});

test('suffixWithUnderscore lowercases and joins', () => {
  assert.equal(suffixWithUnderscore('Submit', 'button'), 'submit_button');
});

test('parseSelector detects xpath prefix', () => {
  assert.deepEqual(parseSelector('xpath://div'), { type: 'xpath', value: '//div' });
  assert.deepEqual(parseSelector('#foo'), { type: 'css', value: '#foo' });
});

test('getElementSelectors resolves from the page group first', () => {
  const elements = {
    login_e2e: { submit_button: ['#login-submit'] },
    common: { submit_button: ['#generic-submit'] },
  };
  const selectors = getElementSelectors(elements, 'login_e2e', 'submit_button');
  assert.deepEqual(selectors, [{ type: 'css', value: '#login-submit' }]);
});

test('getElementSelectors falls back to common when not found on the page', () => {
  const elements = {
    login_e2e: {},
    common: { submit_button: ['#generic-submit'] },
  };
  const selectors = getElementSelectors(elements, 'login_e2e', 'submit_button');
  assert.deepEqual(selectors, [{ type: 'css', value: '#generic-submit' }]);
});

test('getElementSelectors returns empty array when not found anywhere', () => {
  assert.deepEqual(getElementSelectors({}, 'login_e2e', 'missing'), []);
});

test('isElementDefined scans all groups', () => {
  const elements = { common: { submit_button: ['#s'] } };
  assert.equal(isElementDefined(elements, 'submit_button'), true);
  assert.equal(isElementDefined(elements, 'unknown'), false);
});
