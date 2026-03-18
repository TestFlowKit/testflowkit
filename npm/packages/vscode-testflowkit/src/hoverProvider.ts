import * as vscode from 'vscode';
import type { SentenceIndex } from './sentenceIndex';
import type { MacroIndex } from './macroIndex';
import { buildCompletionDocs } from './snippetEngine';
import { buildMacroCompletionDocs } from './macroSnippet';

/** Step keywords to strip before matching. */
const STEP_KEYWORDS = ['given', 'when', 'then', 'and', 'but', '*'];

/**
 * Hover provider for Gherkin .feature files.
 *
 * - When the cursor is on a step line that matches a known sentence pattern,
 *   shows the sentence's description, variables, and example.
 * - When the line contains a macro call name, shows the macro's variable list.
 */
export class HoverProvider implements vscode.HoverProvider {
  constructor(
    private readonly sentences: SentenceIndex,
    private readonly macros: MacroIndex,
  ) {}

  provideHover(
    document: vscode.TextDocument,
    position: vscode.Position,
    _token: vscode.CancellationToken,
  ): vscode.Hover | null {
    const lineText = document.lineAt(position.line).text.trim();
    const stepText = stripStepKeyword(lineText);

    if (!stepText) {
      return null;
    }

    // 1 — Try exact sentence pattern match
    const sentence = this.sentences.findByStepText(stepText);
    if (sentence) {
      return new vscode.Hover(buildCompletionDocs(sentence));
    }

    // 2 — Try macro name match (exact)
    const macro = this.macros.findByName(stepText);
    if (macro) {
      return new vscode.Hover(buildMacroCompletionDocs(macro));
    }

    return null;
  }
}

function stripStepKeyword(line: string): string {
  const lower = line.toLowerCase();

  for (const kw of STEP_KEYWORDS) {
    const prefix = kw + ' ';
    if (lower.startsWith(prefix)) {
      return line.slice(prefix.length).trim();
    }
  }

  return '';
}
