import * as vscode from 'vscode';
import type { Sentence, SentenceVariable } from './types';

/**
 * Converts a sentence pattern like:
 *   "the {string} should be visible"
 * into a VS Code SnippetString:
 *   "the ${1:name} should be visible"
 *
 * Each {type} placeholder is replaced with a tab-stop using the corresponding
 * variable name from the sentence metadata. If there are more wildcards than
 * named variables, a fallback name is used.
 * A final $0 tab-stop is appended so the cursor rests at the end.
 */
export function sentenceToSnippet(sentence: Sentence): vscode.SnippetString {
  const snippetText = buildSnippetText(sentence.sentence, sentence.variables);
  return new vscode.SnippetString(snippetText);
}

/**
 * Pure function version — returns the raw snippet text string.
 * Used directly by tests and internally by sentenceToSnippet.
 */
export function buildSnippetText(
  sentencePattern: string,
  variables: SentenceVariable[],
): string {
  // Wildcards we recognise in generated sentences
  const wildcard = /\{(?:string|int|number|float|bigdecimal|word)\}/g;

  let variableIndex = 0;
  let tabstopIndex = 1;

  const result = sentencePattern.replace(wildcard, () => {
    const variable = variables[variableIndex];
    variableIndex++;
    tabstopIndex++;

    const name = variable?.name ?? `param${variableIndex}`;
    return `\${${tabstopIndex - 1}:${escapeName(name)}}`;
  });

  return `${result}$0`;
}

/** Escape characters that have special meaning inside a SnippetString. */
function escapeName(name: string): string {
  // Only $ and } are special inside a placeholder display text.
  return name.replace(/[\\$}]/g, '\\$&');
}

/**
 * Builds the MarkdownString shown in the completion item's documentation panel.
 */
export function buildCompletionDocs(sentence: Sentence): vscode.MarkdownString {
  const md = new vscode.MarkdownString('', true);
  md.isTrusted = true;

  // Description
  md.appendMarkdown(`${sentence.description}\n\n`);

  // Categories badge
  if (sentence.categories.length > 0) {
    md.appendMarkdown(
      `**Categories:** ${sentence.categories.join(', ')}\n\n`,
    );
  }

  // Variables table
  if (sentence.variables.length > 0) {
    md.appendMarkdown('**Variables:**\n\n');
    md.appendMarkdown('| Name | Type | Description |\n');
    md.appendMarkdown('|------|------|-------------|\n');
    for (const v of sentence.variables) {
      const desc = v.description ?? '';
      md.appendMarkdown(`| \`${v.name}\` | \`${v.type}\` | ${desc} |\n`);
    }
    md.appendMarkdown('\n');
  }

  // Example
  if (sentence.gherkinExample) {
    md.appendMarkdown('**Example:**\n\n');
    md.appendCodeblock(sentence.gherkinExample.trim(), 'gherkin');
  }

  return md;
}
