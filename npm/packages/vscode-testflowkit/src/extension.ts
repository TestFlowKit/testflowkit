import * as vscode from 'vscode';
import { SentenceIndex } from './sentenceIndex';
import { MacroIndex } from './macroIndex';
import { CompletionProvider } from './completionProvider';
import { HoverProvider } from './hoverProvider';

const GHERKIN_SELECTOR: vscode.DocumentSelector = [
  { language: 'gherkin', scheme: 'file' },
  // Some installations detect .feature files under a generic language id.
  { language: 'cucumber', scheme: 'file' },
  // Fallback: match by file extension unconditionally.
  { pattern: '**/*.feature', scheme: 'file' },
];

export async function activate(
  context: vscode.ExtensionContext,
): Promise<void> {
  const sentenceIndex = new SentenceIndex();
  const macroIndex = new MacroIndex();

  // ── Load indices ────────────────────────────────────────────────────────────
  await sentenceIndex.initialize(context);

  await macroIndex.initialize(() => {
    // Fired whenever a .feature file changes — completion list stays fresh.
    vscode.window.setStatusBarMessage(
      '$(sync~spin) TestFlowKit: refreshing macro index…',
      2000,
    );
  });

  showLoadedMessage(sentenceIndex.size, macroIndex.getAll().length);

  // ── Register completion provider ────────────────────────────────────────────
  context.subscriptions.push(
    vscode.languages.registerCompletionItemProvider(
      GHERKIN_SELECTOR,
      new CompletionProvider(sentenceIndex, macroIndex),
      ' ', // trigger after space — catches "Given ", "When ", etc.
    ),
  );

  // ── Register hover provider ─────────────────────────────────────────────────
  context.subscriptions.push(
    vscode.languages.registerHoverProvider(
      GHERKIN_SELECTOR,
      new HoverProvider(sentenceIndex, macroIndex),
    ),
  );

  // ── Refresh commands ────────────────────────────────────────────────────────
  context.subscriptions.push(
    vscode.commands.registerCommand(
      'testflowkit.refreshSentences',
      async () => {
        await sentenceIndex.refresh(context);
        vscode.window.showInformationMessage(
          `TestFlowKit: ${sentenceIndex.size} sentences loaded.`,
        );
      },
    ),
  );

  context.subscriptions.push(
    vscode.commands.registerCommand(
      'testflowkit.refreshMacros',
      async () => {
        await macroIndex.refresh();
        vscode.window.showInformationMessage(
          `TestFlowKit: ${macroIndex.getAll().length} macros loaded.`,
        );
      },
    ),
  );

  // ── Dispose macro watcher on deactivation ──────────────────────────────────
  context.subscriptions.push({ dispose: () => macroIndex.dispose() });
}

export function deactivate(): void {
  // Nothing to do — disposables are handled by the subscriptions array above.
}

function showLoadedMessage(sentenceCount: number, macroCount: number): void {
  if (sentenceCount === 0) {
    vscode.window.showWarningMessage(
      'TestFlowKit: no sentence definitions found. ' +
        'Check the "testflowkit.sentencesPath" setting or run "make generate_doc".',
    );
    return;
  }

  vscode.window.setStatusBarMessage(
    `$(check) TestFlowKit: ${sentenceCount} sentences, ${macroCount} macros ready`,
    4000,
  );
}
