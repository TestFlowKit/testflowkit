/**
 * Minimal VS Code API stub for Jest unit tests.
 * Only the classes/objects used by the tested modules are implemented.
 */
declare class SnippetString {
    value: string;
    constructor(value?: string);
}
declare class MarkdownString {
    value: string;
    isTrusted: boolean;
    constructor(value?: string, _supportThemeIcons?: boolean);
    appendMarkdown(text: string): this;
    appendCodeblock(code: string, language?: string): this;
    appendText(text: string): this;
}
declare const CompletionItemKind: {
    Text: number;
    Method: number;
    Function: number;
    Constructor: number;
    Field: number;
    Variable: number;
    Class: number;
    Interface: number;
    Module: number;
    Property: number;
    Unit: number;
    Value: number;
    Enum: number;
    Keyword: number;
    Snippet: number;
    Color: number;
    File: number;
    Reference: number;
    Folder: number;
};
declare class CompletionItem {
    label: string;
    kind?: number;
    detail?: string;
    documentation?: MarkdownString;
    insertText?: SnippetString | string;
    sortText?: string;
    filterText?: string;
    range?: unknown;
    constructor(label: string, kind?: number);
}
declare class Range {
    start: unknown;
    end: unknown;
    constructor(start: unknown, end: unknown);
}
declare class Position {
    line: number;
    character: number;
    constructor(line: number, character: number);
}
declare class Hover {
    contents: unknown;
    range?: unknown | undefined;
    constructor(contents: unknown, range?: unknown | undefined);
}
declare const workspace: {
    getConfiguration: jest.Mock<any, any, any>;
    workspaceFolders: never[];
    findFiles: jest.Mock<any, any, any>;
    createFileSystemWatcher: jest.Mock<any, any, any>;
};
declare const window: {
    showInformationMessage: jest.Mock<any, any, any>;
    showWarningMessage: jest.Mock<any, any, any>;
    setStatusBarMessage: jest.Mock<any, any, any>;
};
declare const languages: {
    registerCompletionItemProvider: jest.Mock<any, any, any>;
    registerHoverProvider: jest.Mock<any, any, any>;
};
declare const commands: {
    registerCommand: jest.Mock<any, any, any>;
};
