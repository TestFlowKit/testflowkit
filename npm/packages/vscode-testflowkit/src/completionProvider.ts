import * as vscode from 'vscode';
import type { SentenceIndex } from './sentenceIndex';
import type { MacroIndex } from './macroIndex';
import { sentenceToSnippet, buildCompletionDocs } from './snippetEngine';
import { macroToSnippet, buildMacroCompletionDocs } from './macroSnippet';

/** Gherkin step keywords that trigger completions. */
const STEP_KEYWORDS = ['given', 'when', 'then', 'and', 'but', '*'];

/**
 * Provides autocomplete for step definitions (sentences) and macro calls in
 * Gherkin .feature files.
 *
 * Items:
 *   - Sentence items: insert a VS Code SnippetString with tab-stops for every
 *     {string}/{number} wildcard. Documentation shows description + variables +
 *     example.
 *   - Macro items: insert the macro call followed by a two-column data table
 *     whose rows contain the macro's ${var} variable names. The value column in
 *     each row is a tab-stop so the user can press Tab to fill each value in order.
 */
export class CompletionProvider implements vscode.CompletionItemProvider {
  constructor(
    private readonly sentences: SentenceIndex,
    private readonly macros: MacroIndex,
  ) {}

  provideCompletionItems(
    document: vscode.TextDocument,
    position: vscode.Position,
    _token: vscode.CancellationToken,
    _context: vscode.CompletionContext,
  ): vscode.CompletionItem[] {
    const lineText = document.lineAt(position.line).text;
    const stepContext = parseStepContext(lineText, position.character);

    if (!stepContext) {
      return [];
    }

    const { keyword, query, indent } = stepContext;
    const items: vscode.CompletionItem[] = [];

    // ── Standard sentence completions ────────────────────────────────────────
    const matchedSentences = this.sentences.findMatches(query);
    for (const sentence of matchedSentences) {
      const item = new vscode.CompletionItem(
        sentence.sentence,
        vscode.CompletionItemKind.Snippet,
      );

      item.detail = sentence.categories.join(', ');
      item.documentation = buildCompletionDocs(sentence);
      item.insertText = sentenceToSnippet(sentence);
      // Replace the step text after the keyword
      item.range = buildReplaceRange(document, position, keyword, lineText);
      item.sortText = `1_${sentence.sentence}`;
      item.filterText = sentence.sentence;

      items.push(item);
    }

    // ── Macro completions ─────────────────────────────────────────────────────
    const config = vscode.workspace.getConfiguration('testflowkit');
    if (config.get<boolean>('enableMacroCompletion', true)) {
      const matchedMacros = this.macros.findMatches(query);
      for (const macro of matchedMacros) {
        const item = new vscode.CompletionItem(
          macro.name,
          vscode.CompletionItemKind.Function,
        );

        item.detail = `macro (${macro.variables.length} variable${macro.variables.length === 1 ? '' : 's'})`;
        item.documentation = buildMacroCompletionDocs(macro);
        item.insertText = macroToSnippet(macro, indent);
        item.range = buildReplaceRange(document, position, keyword, lineText);
        item.sortText = `2_${macro.name}`;
        item.filterText = macro.name;

        items.push(item);
      }
    }

    return items;
  }
}

// ---------------------------------------------------------------------------
// Parsing helpers
// ---------------------------------------------------------------------------

interface StepContext {
  keyword: string;
  query: string;   // text after the keyword
  indent: string;  // leading whitespace of the current line
}

/**
 * Returns the step keyword and the query text if the cursor is on a step line,
 * or null if the current line is not a step.
 */
function parseStepContext(
  lineText: string,
  cursorChar: number,
): StepContext | null {
  const beforeCursor = lineText.slice(0, cursorChar);
  const trimmed = beforeCursor.trimStart();
  const indent = beforeCursor.slice(0, beforeCursor.length - trimmed.length);

  for (const kw of STEP_KEYWORDS) {
    // Allow "Given ", "given "
    const prefix = kw + ' ';
    if (trimmed.toLowerCase().startsWith(prefix)) {
      const query = trimmed.slice(prefix.length);
      return { keyword: kw, query, indent };
    }
  }

  return null;
}

/**
 * Compute the range that a completion item should replace — from just after the
 * keyword+space up to the cursor.
 */
function buildReplaceRange(
  document: vscode.TextDocument,
  position: vscode.Position,
  keyword: string,
  lineText: string,
): vscode.Range {
  const lowerLine = lineText.trimStart().toLowerCase();
  const afterKeyword = keyword.length + 1; // +1 for the space
  const lineStart = lineText.length - lineText.trimStart().length;
  const replaceStart = lineStart + afterKeyword;

  return new vscode.Range(
    new vscode.Position(position.line, replaceStart),
    position,
  );
}
