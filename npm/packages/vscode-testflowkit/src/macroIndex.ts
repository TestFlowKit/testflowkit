import * as vscode from 'vscode';
import * as fs from 'fs';
import * as path from 'path';
import type { Macro } from './types';

/**
 * Scans all .feature files in the workspace for @macro-tagged scenarios and
 * extracts each macro's name plus the ordered, de-duplicated list of ${var}
 * placeholder names found in the macro's step text.
 *
 * The format mirrors pkg/gherkinparser/macro.go:
 *   - Scenarios tagged with @macro define reusable macros.
 *   - Variable placeholders use ${variable_name} syntax.
 *   - When a macro is called, the caller supplies values via a two-column
 *     data table: | variable_name | value |
 */
export class MacroIndex {
  private macros: Macro[] = [];
  private watchers: vscode.FileSystemWatcher[] = [];

  /** Initial load from all workspace .feature files. */
  async initialize(onChange: () => void): Promise<void> {
    await this.reload();
    this.startWatching(onChange);
  }

  /** Force reload — called when the user triggers the refresh command. */
  async refresh(): Promise<void> {
    await this.reload();
  }

  /** Dispose file-system watchers. */
  dispose(): void {
    for (const w of this.watchers) {
      w.dispose();
    }
    this.watchers = [];
  }

  /** Return all discovered macros. */
  getAll(): Macro[] {
    return [...this.macros];
  }

  /** Look up a macro by its exact scenario name. */
  findByName(name: string): Macro | undefined {
    return this.macros.find(m => m.name === name);
  }

  /** Token-based search across macro names. */
  findMatches(query: string): Macro[] {
    const tokens = query
      .toLowerCase()
      .split(/\s+/)
      .filter(t => t.length > 0);

    if (tokens.length === 0) {
      return [...this.macros];
    }

    return this.macros.filter(m => {
      const lower = m.name.toLowerCase();
      return tokens.every(t => lower.includes(t));
    });
  }

  // ---------------------------------------------------------------------------
  // Internal helpers
  // ---------------------------------------------------------------------------

  private async reload(): Promise<void> {
    const featureFiles = await findFeatureFiles();
    const freshMacros: Macro[] = [];

    for (const filePath of featureFiles) {
      const macrosInFile = extractMacrosFromFile(filePath);
      freshMacros.push(...macrosInFile);
    }

    this.macros = freshMacros;
  }

  private startWatching(onChange: () => void): void {
    // Watch for changes/creation/deletion of .feature files.
    const watcher = vscode.workspace.createFileSystemWatcher('**/*.feature');

    const handler = async () => {
      await this.reload();
      onChange();
    };

    watcher.onDidChange(handler);
    watcher.onDidCreate(handler);
    watcher.onDidDelete(handler);

    this.watchers.push(watcher);
  }
}

// ---------------------------------------------------------------------------
// File-system helpers
// ---------------------------------------------------------------------------

async function findFeatureFiles(): Promise<string[]> {
  const uris = await vscode.workspace.findFiles(
    '**/*.feature',
    '**/node_modules/**',
  );
  return uris.map(u => u.fsPath);
}

/**
 * Parse a single .feature file and return all @macro scenarios found.
 */
export function extractMacrosFromFile(filePath: string): Macro[] {
  let content: string;
  try {
    content = fs.readFileSync(filePath, 'utf8');
  } catch {
    return [];
  }

  return extractMacrosFromContent(content, filePath);
}

/**
 * Pure extraction logic — separated so it can be unit-tested without the
 * file system.
 */
export function extractMacrosFromContent(
  content: string,
  sourceFile: string,
): Macro[] {
  const macros: Macro[] = [];
  const lines = content.split('\n');

  let nextScenarioIsMacro = false;

  for (let i = 0; i < lines.length; i++) {
    const line = lines[i].trim();

    // Collect @macro tag (may appear alone or alongside other tags)
    if (line.startsWith('@')) {
      if (hasMacroTag(line)) {
        nextScenarioIsMacro = true;
      }
      continue;
    }

    // Scenario / Scenario Outline declaration
    if (nextScenarioIsMacro && isScenarioLine(line)) {
      nextScenarioIsMacro = false;
      const name = extractScenarioName(line);
      if (!name) {
        continue;
      }

      // Collect the body of this scenario (until next Scenario / Feature / @tag)
      const bodyLines = collectScenarioBody(lines, i + 1);
      const variables = extractVariableNames(bodyLines);

      macros.push({ name, variables, sourceFile });
    } else if (!line.startsWith('@')) {
      // Reset if we hit a non-tag, non-scenario line before a scenario
      nextScenarioIsMacro = false;
    }
  }

  return macros;
}

function hasMacroTag(tagLine: string): boolean {
  return tagLine
    .split(/\s+/)
    .map(t => t.toLowerCase())
    .includes('@macro');
}

const SCENARIO_RE = /^Scenario(?:\s+Outline)?:\s*(.*)/i;

function isScenarioLine(line: string): boolean {
  return SCENARIO_RE.test(line);
}

function extractScenarioName(line: string): string {
  const m = SCENARIO_RE.exec(line);
  return m ? m[1].trim() : '';
}

/** Collect lines that belong to this scenario's body. */
function collectScenarioBody(lines: string[], startIndex: number): string[] {
  const body: string[] = [];
  for (let i = startIndex; i < lines.length; i++) {
    const trimmed = lines[i].trim();
    // Stop at another scenario, feature, or tag block
    if (
      SCENARIO_RE.test(trimmed) ||
      trimmed.startsWith('Feature:') ||
      trimmed.startsWith('@')
    ) {
      break;
    }
    body.push(lines[i]);
  }
  return body;
}

/**
 * Extract de-duplicated, first-appearance-ordered variable names
 * from ${variable_name} tokens in the given lines.
 *
 * This matches the regex `\$\{([^}]+)\}` used in
 * pkg/gherkinparser/macro.go (macroVarPattern).
 */
export function extractVariableNames(lines: string[]): string[] {
  const seen = new Set<string>();
  const varRe = /\$\{([^}]+)\}/g;

  for (const line of lines) {
    let match: RegExpExecArray | null;
    while ((match = varRe.exec(line)) !== null) {
      const name = match[1].trim();
      if (name && !seen.has(name)) {
        seen.add(name);
      }
    }
  }

  return Array.from(seen);
}
