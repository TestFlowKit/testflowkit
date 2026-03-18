import {
  extractMacrosFromContent,
  extractVariableNames,
} from './macroIndex';

describe('extractVariableNames', () => {
  it('extracts a single variable', () => {
    const result = extractVariableNames(['When the user fills the ${username} field']);
    expect(result).toEqual(['username']);
  });

  it('extracts multiple variables in first-appearance order', () => {
    const result = extractVariableNames([
      'When the user fills the ${username} field with ${value}',
      'And the user fills the ${password} field',
    ]);
    expect(result).toEqual(['username', 'value', 'password']);
  });

  it('de-duplicates repeated variable names', () => {
    const result = extractVariableNames([
      'When the user uses ${token} as auth',
      'Then the result for ${token} should be ok',
    ]);
    expect(result).toEqual(['token']);
  });

  it('returns empty array when no variables present', () => {
    const result = extractVariableNames(['When the user clicks the login button']);
    expect(result).toEqual([]);
  });

  it('trims whitespace around variable names', () => {
    const result = extractVariableNames(['${ username }']);
    expect(result).toEqual(['username']);
  });
});

describe('extractMacrosFromContent', () => {
  const source = '/tmp/test.feature';

  it('extracts a macro with variables', () => {
    const content = `
Feature: Test
@macro
Scenario: user login with credentials
  When the user fills the \${username} field
  And the user fills the \${password} field
  And the user clicks login
Scenario: some other test
  Given something
    `.trim();

    const macros = extractMacrosFromContent(content, source);
    expect(macros).toHaveLength(1);
    expect(macros[0].name).toBe('user login with credentials');
    expect(macros[0].variables).toEqual(['username', 'password']);
    expect(macros[0].sourceFile).toBe(source);
  });

  it('extracts a macro with no variables', () => {
    const content = `
@macro
Scenario: user logout
  When the user clicks logout
    `.trim();

    const macros = extractMacrosFromContent(content, source);
    expect(macros).toHaveLength(1);
    expect(macros[0].name).toBe('user logout');
    expect(macros[0].variables).toEqual([]);
  });

  it('extracts multiple macros', () => {
    const content = `
@macro
Scenario: login with \${user}
  Given something with \${user}

@macro
Scenario: navigate to page \${page}
  When navigation to \${page}
    `.trim();

    const macros = extractMacrosFromContent(content, source);
    expect(macros).toHaveLength(2);
    expect(macros[0].variables).toEqual(['user']);
    expect(macros[1].variables).toEqual(['page']);
  });

  it('ignores scenarios not tagged with @macro', () => {
    const content = `
Scenario: normal test
  Given something with \${notAMacroVar}
    `.trim();

    const macros = extractMacrosFromContent(content, source);
    expect(macros).toHaveLength(0);
  });

  it('is case-insensitive for @macro tag', () => {
    const content = `
@Macro
Scenario: tagged macro
  When step with \${var}
    `.trim();

    const macros = extractMacrosFromContent(content, source);
    expect(macros).toHaveLength(1);
  });

  it('handles @macro alongside other tags on the same line', () => {
    const content = `
@macro @smoke
Scenario: macro with extra tags
  When step with \${val}
    `.trim();

    const macros = extractMacrosFromContent(content, source);
    expect(macros).toHaveLength(1);
    expect(macros[0].variables).toEqual(['val']);
  });
});
