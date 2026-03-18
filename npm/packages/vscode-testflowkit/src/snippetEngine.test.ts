import { buildSnippetText } from './snippetEngine';
import type { SentenceVariable } from './types';

describe('buildSnippetText', () => {
  it('replaces a single {string} wildcard with an ordered tab-stop', () => {
    const vars: SentenceVariable[] = [{ name: 'path', type: 'string' }];
    const result = buildSnippetText('the response should have field {string}', vars);
    expect(result).toBe('the response should have field ${1:path}$0');
  });

  it('replaces two {string} wildcards with sequential tab-stops', () => {
    const vars: SentenceVariable[] = [
      { name: 'text', type: 'string' },
      { name: 'name', type: 'string' },
    ];
    const result = buildSnippetText(
      'the user enters {string} into the {string} field',
      vars,
    );
    expect(result).toBe(
      'the user enters ${1:text} into the ${2:name} field$0',
    );
  });

  it('handles {int} and {float} wildcards', () => {
    const vars: SentenceVariable[] = [{ name: 'code', type: 'int' }];
    const result = buildSnippetText(
      'the response status code should be {int}',
      vars,
    );
    expect(result).toBe('the response status code should be ${1:code}$0');
  });

  it('falls back to param1, param2 when variables list is shorter than wildcards', () => {
    const result = buildSnippetText(
      'the {string} field should be {string}',
      [{ name: 'field', type: 'string' }],
    );
    expect(result).toBe('the ${1:field} field should be ${2:param2}$0');
  });

  it('returns plain text with $0 when there are no wildcards', () => {
    const result = buildSnippetText('I send the request', []);
    expect(result).toBe('I send the request$0');
  });

  it('escapes dollar signs in variable names', () => {
    const vars: SentenceVariable[] = [{ name: 'some$var', type: 'string' }];
    const result = buildSnippetText('{string}', vars);
    expect(result).toBe('${1:some\\$var}$0');
  });
});
