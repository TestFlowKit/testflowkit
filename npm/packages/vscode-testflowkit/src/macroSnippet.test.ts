import { buildMacroSnippetText } from './macroSnippet';
import type { Macro } from './types';

const src = '/tmp/macros.feature';

describe('buildMacroSnippetText', () => {
  it('inserts only the macro name when it has no variables', () => {
    const macro: Macro = { name: 'user logout', variables: [], sourceFile: src };
    expect(buildMacroSnippetText(macro, '')).toBe('user logout$0');
  });

  it('inserts macro name + data table rows with tab-stops for each variable', () => {
    const macro: Macro = {
      name: 'user login with credentials',
      variables: ['username', 'password'],
      sourceFile: src,
    };

    const result = buildMacroSnippetText(macro, '  ');
    expect(result).toBe(
      [
        'user login with credentials',
        '    | username | ${1:username_value} |',
        '    | password | ${2:password_value} |',
        '$0',
      ].join('\n'),
    );
  });

  it('generates sequential tab-stop indices across variables', () => {
    const macro: Macro = {
      name: 'complete workflow',
      variables: ['username', 'password', 'page'],
      sourceFile: src,
    };

    const result = buildMacroSnippetText(macro, '');
    expect(result).toContain('${1:');
    expect(result).toContain('${2:');
    expect(result).toContain('${3:');
    expect(result).not.toContain('${4:');
  });

  it('preserves indentation: table rows use indent + 2 spaces', () => {
    const macro: Macro = {
      name: 'do something',
      variables: ['key'],
      sourceFile: src,
    };

    const result = buildMacroSnippetText(macro, '    ');
    // "    " indent + "  " = 6 spaces before |
    expect(result).toMatch(/^do something\n {6}\| key \|/);
  });
});
