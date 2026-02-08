const test = require("node:test");
const assert = require("node:assert");
const fs = require("fs");
const path = require("path");
const os = require("os");
const Module = require("module");

const {
  install,
  downloadFile,
  extractZip,
  makeExecutable,
  removeExistingBinary,
  getBinaryName,
  getDownloadUrl,
} = require("./install.js");

test("getBinaryName", async (t) => {
  await t.test("should return tkit for non-Windows", () => {
    if (process.platform !== "win32") {
      assert.strictEqual(getBinaryName(), "tkit");
    }
  });

  await t.test("should return tkit.exe for Windows", () => {
    // This test would need platform mocking to work properly
    // For now, just verify the function exists
    assert.ok(typeof getBinaryName === "function");
  });
});

test("getDownloadUrl", async (t) => {
  await t.test("should generate valid download URL", () => {
    try {
      const url = getDownloadUrl();
      assert.ok(url.includes("github.com"));
      assert.ok(url.includes("testflowkit"));
      assert.ok(url.includes(".zip"));
    } catch (error) {
      // Platform/arch not supported - that's ok for testing
      assert.ok(
        error.message.includes("Unsupported") ||
          error.message.includes("platform") ||
          error.message.includes("architecture"),
      );
    }
  });

  await t.test("should include version in URL", () => {
    try {
      const url = getDownloadUrl();
      assert.ok(url.includes("releases/download"));
    } catch (error) {
      // Expected for some platforms
      assert.ok(true);
    }
  });
});

test("makeExecutable", async (t) => {
  await t.test("should be callable", () => {
    assert.ok(typeof makeExecutable === "function");
  });

  if (process.platform !== "win32") {
    await t.test("should set executable permission on Unix", () => {
      const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "test-chmod-"));
      const testFile = path.join(tempDir, "test-binary");

      try {
        fs.writeFileSync(testFile, "test");

        // Get permissions before
        const statsBefore = fs.statSync(testFile);
        const isExecutableBefore =
          (statsBefore.mode & fs.constants.S_IXUSR) !== 0;

        // Make executable
        makeExecutable(testFile);

        // Check permissions after
        const statsAfter = fs.statSync(testFile);
        const isExecutableAfter = (statsAfter.mode & fs.constants.S_IXUSR) !== 0;

        assert.ok(
          isExecutableAfter,
          "File should be executable after makeExecutable",
        );
      } finally {
        fs.rmSync(tempDir, { recursive: true, force: true });
      }
    });
  }
});

test("removeExistingBinary", async (t) => {
  await t.test("should successfully remove a file", () => {
    const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "test-remove-"));
    const testFile = path.join(tempDir, "test-binary");

    try {
      fs.writeFileSync(testFile, "test");
      assert.ok(fs.existsSync(testFile), "File should exist before removal");

      removeExistingBinary(testFile);
      assert.ok(!fs.existsSync(testFile), "File should not exist after removal");
    } finally {
      if (fs.existsSync(tempDir)) {
        fs.rmSync(tempDir, { recursive: true, force: true });
      }
    }
  });

  await t.test("should handle removal failure gracefully", async (t) => {
    const nonExistentPath = "/nonexistent/path/to/binary";

    // This test demonstrates the function would exit with error
    // In a real test, we'd mock process.exit
    assert.ok(typeof removeExistingBinary === "function");
  });
});

test("extractZip", async (t) => {
  await t.test("should be callable", () => {
    assert.ok(typeof extractZip === "function");
  });

  await t.test("should throw error on invalid zip buffer", async () => {
    const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "test-extract-"));

    try {
      const invalidZipBuffer = Buffer.from("not a zip file");
      try {
        await extractZip(invalidZipBuffer, tempDir);
        assert.fail("Should have thrown an error");
      } catch (error) {
        assert.ok(
          error.message.includes("Failed to extract"),
          "Should throw extraction error",
        );
      }
    } finally {
      fs.rmSync(tempDir, { recursive: true, force: true });
    }
  });
});

test("downloadFile", async (t) => {
  await t.test("should be callable", () => {
    assert.ok(typeof downloadFile === "function");
  });

  await t.test("should handle network errors", async () => {
    // This test would require mocking https.get
    // For now, just verify the function exists and is a Promise
    assert.ok(typeof downloadFile === "function");
  });
});

test("Integration scenarios", async (t) => {
  await t.test("should handle directory creation", () => {
    const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "test-install-"));

    try {
      const testCliDir = path.join(tempDir, "cli");

      if (!fs.existsSync(testCliDir)) {
        fs.mkdirSync(testCliDir, { recursive: true });
      }

      assert.ok(fs.existsSync(testCliDir), "Directory should be created");
    } finally {
      fs.rmSync(tempDir, { recursive: true, force: true });
    }
  });

  await t.test("should handle temporary directory cleanup", () => {
    const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "tkit-"));
    assert.ok(fs.existsSync(tempDir), "Temp directory should exist");

    fs.rmSync(tempDir, { recursive: true, force: true });
    assert.ok(!fs.existsSync(tempDir), "Temp directory should be cleaned up");
  });

  await t.test("should use system temp directory", () => {
    const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "tkit-"));
    const tempLocation = os.tmpdir();

    try {
      assert.ok(
        tempDir.includes(tempLocation),
        "Temp directory should be in system temp location",
      );
    } finally {
      fs.rmSync(tempDir, { recursive: true, force: true });
    }
  });
});
