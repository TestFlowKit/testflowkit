import assert from "node:assert/strict";
import fs from "node:fs";
import os from "node:os";
import path from "node:path";
import { describe, it } from "node:test";
import { extractMacros } from "./extractMacros.js";

function mkTmp(): string {
  return fs.mkdtempSync(path.join(os.tmpdir(), "tfk-mcp-macros-"));
}

function writeFeature(dir: string, name: string, content: string): void {
  fs.writeFileSync(path.join(dir, name), content, "utf-8");
}

describe("extractMacros", () => {
  it("extracts a macro scenario with placeholders and a call example", () => {
    const dir = mkTmp();
    try {
      writeFeature(
        dir,
        "login.feature",
        [
          "Feature: macros",
          "",
          "  @macro",
          "  Scenario: Login as user",
          '    Given the user goes to the "login" page',
          '    When the user enters "${email}" into the "email" field',
          '    And the user enters "${password}" into the "password" field',
          "",
          "  Scenario: Not a macro",
          "    Given a regular step",
        ].join("\n"),
      );

      const macros = extractMacros("*.feature", dir);

      assert.equal(macros.length, 1);
      assert.equal(macros[0].name, "Login as user");
      assert.deepEqual(macros[0].variables, ["email", "password"]);
      assert.equal(macros[0].steps.length, 3);
      assert.match(macros[0].callExample, /Given Login as user/);
      assert.match(macros[0].callExample, /\| email \| <value> \|/);
      assert.match(macros[0].callExample, /\| password \| <value> \|/);
    } finally {
      fs.rmSync(dir, { recursive: true, force: true });
    }
  });

  it("supports multiple tags across multiple lines above Scenario", () => {
    const dir = mkTmp();
    try {
      writeFeature(
        dir,
        "tagged.feature",
        [
          "Feature: macros",
          "",
          "  @api",
          "  @macro",
          "  Scenario: prepare request",
          '    Given I prepare a request to "jsonplaceholder.create_post"',
        ].join("\n"),
      );

      const macros = extractMacros("*.feature", dir);
      assert.equal(macros.length, 1);
      assert.equal(macros[0].name, "prepare request");
    } finally {
      fs.rmSync(dir, { recursive: true, force: true });
    }
  });

  it("extracts variables from doc-string bodies", () => {
    const dir = mkTmp();
    try {
      writeFeature(
        dir,
        "docstring.feature",
        [
          "Feature: macros",
          "",
          "  @macro",
          "  Scenario: create post",
          '    Given I prepare a request to "jsonplaceholder.create_post"',
          "    When I set the request body to:",
          '      """',
          "      {",
          '      "title": "${title}",',
          '      "body": "${body}",',
          '      "userId": ${userId}',
          "      }",
          '      """',
          "    And I send the request",
        ].join("\n"),
      );

      const macros = extractMacros("*.feature", dir);
      assert.equal(macros.length, 1);
      assert.deepEqual(macros[0].variables, ["title", "body", "userId"]);
    } finally {
      fs.rmSync(dir, { recursive: true, force: true });
    }
  });

  it("tolerates free-text description lines under Scenario", () => {
    const dir = mkTmp();
    try {
      writeFeature(
        dir,
        "described.feature",
        [
          "Feature: macros",
          "",
          "    @macro",
          "    Scenario: I try to create a new post with the following details:",
          "        this macro was created for test a bug in macro application",
          "        the macro did not include the docstring",
          "",
          '        Given I prepare a request to "jsonplaceholder.create_post"',
          "        And I send the request",
        ].join("\n"),
      );

      const macros = extractMacros("*.feature", dir);
      assert.equal(macros.length, 1);
      assert.equal(
        macros[0].name,
        "I try to create a new post with the following details:",
      );
      assert.equal(macros[0].steps.length, 2);
    } finally {
      fs.rmSync(dir, { recursive: true, force: true });
    }
  });

  it("excludes non-macro scenarios and scans across multiple files", () => {
    const dir = mkTmp();
    try {
      writeFeature(
        dir,
        "a.feature",
        ["Feature: a", "", "  Scenario: plain", "    Given a step"].join("\n"),
      );
      writeFeature(
        dir,
        "b.feature",
        [
          "Feature: b",
          "",
          "  @macro",
          "  Scenario: reusable setup",
          "    Given a setup step",
        ].join("\n"),
      );

      const macros = extractMacros("*.feature", dir);
      assert.equal(macros.length, 1);
      assert.equal(macros[0].name, "reusable setup");
      assert.equal(macros[0].file, "b.feature");
    } finally {
      fs.rmSync(dir, { recursive: true, force: true });
    }
  });
});
