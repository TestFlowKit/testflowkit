import * as vscode from 'vscode';
import type { Macro } from './types';

/**
 * Builds the snippet text for a macro invocation.
 *
 * When a macro has variables the snippet inserts the macro call line followed
 * by a two-column Gherkin data table where:
 *   - Column 1: the variable name (static, non-editable)
 *   - Column 2: an editable tab-stop so the user can Tab through each value
 *
 * Example output for macro "user login with credentials" with variables
 * ["username", "password"]:
 *
 *   user login with credentials
 *     | username | ${1:username_value} |
 *     | password | ${2:password_value} |$0
 *
 * When the macro has no variables, only the call line is inserted.
 */
export function macroToSnippet(macro: Macro, indent: string): vscode.SnippetString {
  const text = buildMacroSnippetText(macro, indent);
  return new vscode.SnippetString(text);
}

/**
 * Pure function version for testing.
 *
 * @param macro   The macro definition containing name and variable list.
 * @param indent  Whitespace prefix used for the caller line (the table rows
 *                will be indented by two extra spaces to align with step text).
 */
export function buildMacroSnippetText(macro: Macro, indent: string = ''): string {
  if (macro.variables.length === 0) {
    return `${macro.name}$0`;
  }

  const tableIndent = `${indent}  `;
  const rows = macro.variables.map((varName, idx) => {
    const tabstop = idx + 1;
    return `${tableIndent}| ${varName} | \${${tabstop}:${escapeSnippet(varName + '_value')}} |`;
  });

  return [macro.name, ...rows, '$0'].join('\n');
}

function escapeSnippet(text: string): string {
  return text.replace(/[\\$}]/g, '\\$&');
}

/**
 * Builds the MarkdownString shown in the completion docs panel for a macro.
 */
export function buildMacroCompletionDocs(macro: Macro): vscode.MarkdownString {
  const md = new vscode.MarkdownString('', true);
  md.isTrusted = true;

  md.appendMarkdown(`**Macro** — reusable scenario defined with \`@macro\`.\n\n`);

  if (macro.variables.length > 0) {
    md.appendMarkdown(
      `A data table will be inserted automatically with ${macro.variables.length} variable row(s).\n\n`,
    );
    md.appendMarkdown('**Variables:**\n\n');
    for (const v of macro.variables) {
      md.appendMarkdown(`- \`${v}\`\n`);
    }
  } else {
    md.appendMarkdown('This macro has no variables.\n');
  }

  md.appendMarkdown(`\n*Defined in: \`${macro.sourceFile}\`*`);

  return md;
}
