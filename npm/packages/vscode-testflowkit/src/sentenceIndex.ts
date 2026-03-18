import * as vscode from 'vscode';
import * as path from 'path';
import * as fs from 'fs';
import type { Sentence } from './types';

/**
 * Loads, caches, and queries all sentence JSON documents from
 * documentation/content/sentences/**\/*.json (or a user-configured path).
 *
 * Sentences come from the generator at scripts/doc_generator/main.go and
 * follow the schema in documentation/data/sentence.ts.
 */
export class SentenceIndex {
  private sentences: Sentence[] = [];

  /** Load all sentence JSON files from the workspace. */
  async initialize(context: vscode.ExtensionContext): Promise<void> {
    this.sentences = await loadAllSentences(context);
  }

  /** Refresh by reloading every sentence file. */
  async refresh(context: vscode.ExtensionContext): Promise<void> {
    this.sentences = await loadAllSentences(context);
  }

  /** Return all sentences whose text contains each token of the query string. */
  findMatches(query: string): Sentence[] {
    const tokens = query
      .toLowerCase()
      .split(/\s+/)
      .filter(t => t.length > 0);

    if (tokens.length === 0) {
      return [...this.sentences];
    }

    return this.sentences.filter(s => {
      const lower = s.sentence.toLowerCase();
      return tokens.every(t => lower.includes(t));
    });
  }

  /** Look up the first sentence whose pattern matches the full step text. */
  findByStepText(stepText: string): Sentence | undefined {
    const lower = stepText.toLowerCase().trim();
    return this.sentences.find(s => matchesSentencePattern(s.sentence, lower));
  }

  get size(): number {
    return this.sentences.length;
  }
}

// ---------------------------------------------------------------------------
// Internal helpers
// ---------------------------------------------------------------------------

async function loadAllSentences(context: vscode.ExtensionContext): Promise<Sentence[]> {
  const sentencesDir = resolveSentencesDir(context);
  if (!sentencesDir) {
    return [];
  }

  const jsonFiles = collectJsonFiles(sentencesDir);
  const results: Sentence[] = [];

  for (const filePath of jsonFiles) {
    const sentence = parseSentenceFile(filePath);
    if (sentence) {
      results.push(sentence);
    }
  }

  return results;
}

function resolveSentencesDir(context: vscode.ExtensionContext): string | null {
  const config = vscode.workspace.getConfiguration('testflowkit');
  const customPath = config.get<string>('sentencesPath', '').trim();

  if (customPath) {
    const workspaceRoot = vscode.workspace.workspaceFolders?.[0]?.uri.fsPath;
    if (!workspaceRoot) {
      return null;
    }
    const resolved = path.isAbsolute(customPath)
      ? customPath
      : path.join(workspaceRoot, customPath);
    return fs.existsSync(resolved) ? resolved : null;
  }

  // Auto-detect: walk up from the extension's install location and from each
  // workspace folder looking for documentation/content/sentences.
  const candidates: string[] = [];

  for (const folder of vscode.workspace.workspaceFolders ?? []) {
    candidates.push(
      path.join(folder.uri.fsPath, 'documentation', 'content', 'sentences'),
    );
  }

  for (const candidate of candidates) {
    if (fs.existsSync(candidate)) {
      return candidate;
    }
  }

  return null;
}

function collectJsonFiles(dir: string): string[] {
  const results: string[] = [];

  function walk(current: string) {
    let entries: fs.Dirent[];
    try {
      entries = fs.readdirSync(current, { withFileTypes: true });
    } catch {
      return;
    }

    for (const entry of entries) {
      const full = path.join(current, entry.name);
      if (entry.isDirectory()) {
        walk(full);
      } else if (entry.isFile() && entry.name.endsWith('.json')) {
        results.push(full);
      }
    }
  }

  walk(dir);
  return results;
}

function parseSentenceFile(filePath: string): Sentence | null {
  try {
    const raw = fs.readFileSync(filePath, 'utf8');
    const obj = JSON.parse(raw) as Record<string, unknown>;

    if (
      typeof obj.sentence !== 'string' ||
      typeof obj.description !== 'string'
    ) {
      return null;
    }

    return {
      sentence: obj.sentence,
      description: obj.description,
      categories: Array.isArray(obj.categories)
        ? (obj.categories as string[])
        : [],
      gherkinExample:
        typeof obj.gherkinExample === 'string' ? obj.gherkinExample : '',
      variables: Array.isArray(obj.variables)
        ? (obj.variables as Array<{ name: string; type: string; description?: string }>)
        : [],
    };
  } catch {
    return null;
  }
}

/**
 * Returns true when the given step text matches the sentence pattern.
 * Wildcards ({string}, {number}, {int}, {float}, {word}, {bigdecimal})
 * are converted to a simple token-level regex.
 */
export function matchesSentencePattern(
  sentencePattern: string,
  stepText: string,
): boolean {
  // Escape regex metacharacters, then replace wildcard placeholders.
  const escaped = sentencePattern
    .replace(/[.*+?^${}()|[\]\\]/g, '\\$&')
    // After escaping {, } become \{, \} — convert known escaped wildcards back.
    .replace(/\\\{string\\\}/g, '"[^"]*"')
    .replace(/\\\{int\\\}|\\\{number\\\}|\\\{float\\\}|\\\{bigdecimal\\\}/g, '-?\\d+(?:\\.\\d+)?')
    .replace(/\\\{word\\\}/g, '\\S+');

  try {
    const re = new RegExp(`^${escaped}$`, 'i');
    return re.test(stepText);
  } catch {
    return false;
  }
}
