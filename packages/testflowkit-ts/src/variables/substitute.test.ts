import assert from 'node:assert/strict';
import { test } from 'node:test';
import { substituteEnv, substituteVariables } from './substitute.js';

test('substituteEnv replaces {{ env.X }} placeholders', () => {
  const result = substituteEnv('base_url: "{{ env.host }}"', { host: 'http://localhost:3000' });
  assert.equal(result, 'base_url: "http://localhost:3000"');
});

test('substituteEnv leaves unknown env placeholders untouched', () => {
  const result = substituteEnv('{{ env.missing }}', {});
  assert.equal(result, '{{ env.missing }}');
});

test('substituteEnv ignores non-env placeholders', () => {
  const result = substituteEnv('{{ someVar }}', {});
  assert.equal(result, '{{ someVar }}');
});

test('substituteVariables replaces scenario variables', () => {
  const vars = new Map([['token', 'abc123']]);
  assert.equal(substituteVariables('Bearer {{ token }}', vars), 'Bearer abc123');
});

test('substituteVariables replaces env placeholders using the provided env map', () => {
  const result = substituteVariables('{{ env.host }}', new Map(), { host: 'http://x' });
  assert.equal(result, 'http://x');
});
