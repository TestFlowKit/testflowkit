import assert from "node:assert/strict";
import { mkdtempSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { test } from "node:test";
import { loadConfig } from "../src/config/loader.js";

function writeConfig(contents: string): string {
  const dir = mkdtempSync(join(tmpdir(), "testflowkit-"));
  const path = join(dir, "testflowkit.yml");
  writeFileSync(path, contents, "utf-8");
  return path;
}

test("loadConfig parses settings and frontend elements/pages", () => {
  const path = writeConfig(`
settings:
  gherkin_location: "features"
frontend:
  elements:
    common:
      submit_button:
        - "#submit"
  pages:
    login_e2e: "login"
`);

  const config = loadConfig(path);
  assert.equal(config.settings.gherkinLocation, "features");
  assert.deepEqual(config.frontend?.elements.common.submit_button, ["#submit"]);
  assert.equal(config.frontend?.pages.login_e2e, "login");
});

test("loadConfig substitutes {{ env.X }} from the inline env block", () => {
  const path = writeConfig(`
settings:
  gherkin_location: "features"
env:
  google_url: "https://google.com"
frontend:
  elements: {}
  pages:
    google: "{{ env.google_url }}"
`);

  const config = loadConfig(path);
  assert.equal(config.frontend?.pages.google, "https://google.com");
});

test("loadConfig throws when gherkin_location is missing", () => {
  const path = writeConfig("settings: {}\n");
  assert.throws(() => loadConfig(path), /gherkin_location is required/);
});
