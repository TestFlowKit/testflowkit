/**
 * Minimal VS Code API stub for Jest unit tests.
 * Only the classes/objects used by the tested modules are implemented.
 */

class SnippetString {
  value: string;
  constructor(value: string = '') {
    this.value = value;
  }
}

class MarkdownString {
  value: string;
  isTrusted: boolean = false;

  constructor(value: string = '', _supportThemeIcons?: boolean) {
    this.value = value;
  }

  appendMarkdown(text: string): this {
    this.value += text;
    return this;
  }

  appendCodeblock(code: string, language?: string): this {
    this.value += `\`\`\`${language ?? ''}\n${code}\n\`\`\`\n`;
    return this;
  }

  appendText(text: string): this {
    this.value += text;
    return this;
  }
}

const CompletionItemKind = {
  Text: 0,
  Method: 1,
  Function: 2,
  Constructor: 3,
  Field: 4,
  Variable: 5,
  Class: 6,
  Interface: 7,
  Module: 8,
  Property: 9,
  Unit: 10,
  Value: 11,
  Enum: 12,
  Keyword: 13,
  Snippet: 14,
  Color: 15,
  File: 16,
  Reference: 17,
  Folder: 18,
};

class CompletionItem {
  label: string;
  kind?: number;
  detail?: string;
  documentation?: MarkdownString;
  insertText?: SnippetString | string;
  sortText?: string;
  filterText?: string;
  range?: unknown;

  constructor(label: string, kind?: number) {
    this.label = label;
    this.kind = kind;
  }
}

class Range {
  constructor(
    public start: unknown,
    public end: unknown,
  ) {}
}

class Position {
  constructor(
    public line: number,
    public character: number,
  ) {}
}

class Hover {
  constructor(
    public contents: unknown,
    public range?: unknown,
  ) {}
}

const workspace = {
  getConfiguration: jest.fn().mockReturnValue({
    get: jest.fn().mockReturnValue(''),
  }),
  workspaceFolders: [],
  findFiles: jest.fn().mockResolvedValue([]),
  createFileSystemWatcher: jest.fn().mockReturnValue({
    onDidChange: jest.fn(),
    onDidCreate: jest.fn(),
    onDidDelete: jest.fn(),
    dispose: jest.fn(),
  }),
};

const window = {
  showInformationMessage: jest.fn(),
  showWarningMessage: jest.fn(),
  setStatusBarMessage: jest.fn(),
};

const languages = {
  registerCompletionItemProvider: jest.fn(),
  registerHoverProvider: jest.fn(),
};

const commands = {
  registerCommand: jest.fn(),
};

module.exports = {
  SnippetString,
  MarkdownString,
  CompletionItemKind,
  CompletionItem,
  Range,
  Position,
  Hover,
  workspace,
  window,
  languages,
  commands,
};
