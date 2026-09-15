import type { FrontendElements } from './types.js';

export type SelectorType = 'css' | 'xpath';

export interface Selector {
  type: SelectorType;
  value: string;
}

const XPATH_PREFIX = 'xpath:';

export function parseSelector(raw: string): Selector {
  const trimmed = raw.trim();
  if (trimmed.startsWith(XPATH_PREFIX)) {
    return { type: 'xpath', value: trimmed.slice(XPATH_PREFIX.length) };
  }
  return { type: 'css', value: trimmed };
}

/** snake_case, matching the Go `stringutils.SnakeCase` used before every element lookup. */
export function snakeCase(input: string): string {
  return input
    .trim()
    .replace(/([a-z0-9])([A-Z])/g, '$1_$2')
    .replace(/[\s-]+/g, '_')
    .toLowerCase();
}

export function suffixWithUnderscore(label: string, suffix: string): string {
  return `${label.trim().toLowerCase()}_${suffix.trim()}`;
}

/**
 * Resolves an element name against `elements[pageName]`, falling back to `elements["common"]`,
 * mirroring the Go chain-of-responsibility in internal/config/frontend.go.
 */
export function getElementSelectors(
  elements: FrontendElements,
  pageName: string,
  elementName: string,
): Selector[] {
  const key = snakeCase(elementName);
  const fromPage = elements[pageName]?.[key];
  if (fromPage && fromPage.length > 0) {
    return fromPage.map(parseSelector);
  }
  const fromCommon = elements.common?.[key];
  if (fromCommon && fromCommon.length > 0) {
    return fromCommon.map(parseSelector);
  }
  return [];
}

/** True if `elementName` is defined under any group — used for fail-fast validation. */
export function isElementDefined(elements: FrontendElements, elementName: string): boolean {
  const key = snakeCase(elementName);
  return Object.values(elements).some((group) => key in group);
}
