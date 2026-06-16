import assert from "node:assert/strict";
import fs from "node:fs";
import os from "node:os";
import path from "node:path";
import { describe, it } from "node:test";
import { fetchFromLocalCli } from "./fetchLocalCli.js";

function makeTempDir(): string {
  return fs.mkdtempSync(path.join(os.tmpdir(), "tkit-mcp-fetchcli-"));
}

function writeFakeCliScript(dir: string, output: string): string {
  const scriptPath = path.join(dir, "fake-tkit.sh");
  fs.writeFileSync(scriptPath, `#!/bin/sh\necho '${output}'`, { mode: 0o755 });
  return scriptPath;
}

describe("fetchFromLocalCli", () => {
  it("returns trimmed stdout from a successful CLI command", () => {
    const tempDir = makeTempDir();
    try {
      const catalogJson =
        '[{"sentence":"test step","description":"desc","categories":["restapi"],"example":"Given test step"}]';
      const scriptPath = writeFakeCliScript(tempDir, catalogJson);
      const result = fetchFromLocalCli(scriptPath, tempDir);
      assert.equal(result, catalogJson);
    } finally {
      fs.rmSync(tempDir, { recursive: true, force: true });
    }
  });

  it("throws when the binary does not exist", () => {
    assert.throws(() =>
      fetchFromLocalCli("/nonexistent/binary/tkit", process.cwd()),
    );
  });

  it("throws when the command exits with a non-zero code", () => {
    const tempDir = makeTempDir();
    try {
      const scriptPath = path.join(tempDir, "failing-tkit.sh");
      fs.writeFileSync(scriptPath, "#!/bin/sh\nexit 1", { mode: 0o755 });
      assert.throws(() => fetchFromLocalCli(scriptPath, tempDir));
    } finally {
      fs.rmSync(tempDir, { recursive: true, force: true });
    }
  });

  it("runs the command in the given cwd", () => {
    const tempDir = makeTempDir();
    try {
      // Script prints its $PWD so we can verify cwd was set correctly
      const scriptPath = path.join(tempDir, "pwd-tkit.sh");
      fs.writeFileSync(scriptPath, "#!/bin/sh\npwd", { mode: 0o755 });
      const result = fetchFromLocalCli(scriptPath, tempDir);
      assert.equal(result, fs.realpathSync(tempDir));
    } finally {
      fs.rmSync(tempDir, { recursive: true, force: true });
    }
  });
});
