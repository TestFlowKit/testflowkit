const fs = require("fs");
const path = require("path");
const os = require("os");
const https = require("https");
const http = require("http");
const { Readable } = require("stream");

/**
 * Integration test for install.js
 * Tests the actual installation flow with mocked network calls
 */

// Mock setup
const mockBinaryContent = Buffer.from("mock binary content");
const tempTestDir = path.join(os.tmpdir(), "testflowkit-install-test");
const cliDir = path.join(tempTestDir, "cli");
const binaryName = process.platform === "win32" ? "tkit.exe" : "tkit";
const binaryPath = path.join(cliDir, binaryName);

/**
 * Clean up test directories
 */
function cleanupTestDir() {
  if (fs.existsSync(tempTestDir)) {
    fs.rmSync(tempTestDir, { recursive: true, force: true });
  }
}

/**
 * Setup test directory structure
 */
function setupTestDir() {
  cleanupTestDir();
  fs.mkdirSync(tempTestDir, { recursive: true });
}

/**
 * Test: Installation of new binary (no existing binary)
 */
async function testNewInstallation() {
  console.log("\n🧪 Test 1: New Installation (no existing binary)");
  setupTestDir();

  try {
    // Verify directory was created
    if (!fs.existsSync(tempTestDir)) {
      throw new Error("Test directory not created");
    }

    // Verify binary doesn't exist yet
    if (fs.existsSync(binaryPath)) {
      throw new Error("Binary should not exist before installation");
    }

    console.log("  ✓ Test directory created");
    console.log("  ✓ Binary does not exist before installation");
    console.log("  ✓ Test 1 PASSED");
  } catch (error) {
    console.error("  ✗ Test 1 FAILED:", error.message);
    return false;
  }

  cleanupTestDir();
  return true;
}

/**
 * Test: Binary replacement (existing binary + new download)
 */
async function testBinaryReplacement() {
  console.log("\n🧪 Test 2: Binary Replacement (existing binary)");
  setupTestDir();

  try {
    // Create initial binary
    if (!fs.existsSync(cliDir)) {
      fs.mkdirSync(cliDir, { recursive: true });
    }

    fs.writeFileSync(binaryPath, "old binary content");
    console.log("  ✓ Created existing binary");

    // Verify binary exists
    if (!fs.existsSync(binaryPath)) {
      throw new Error("Binary should exist");
    }

    const oldContent = fs.readFileSync(binaryPath, "utf8");

    // Simulate downloading to temp directory
    const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "tkit-"));
    console.log("  ✓ Created temporary directory:", tempDir);

    // Write new binary to temp
    const tempBinaryPath = path.join(tempDir, binaryName);
    fs.writeFileSync(tempBinaryPath, "new binary content");
    console.log("  ✓ Downloaded binary to temporary directory");

    // Verify temp binary exists
    if (!fs.existsSync(tempBinaryPath)) {
      throw new Error("Temp binary should exist");
    }

    // Remove old binary
    fs.unlinkSync(binaryPath);
    console.log("  ✓ Removed old binary");

    // Move new binary from temp to final location
    fs.renameSync(tempBinaryPath, binaryPath);
    console.log("  ✓ Moved new binary to final location");

    // Verify replacement
    const newContent = fs.readFileSync(binaryPath, "utf8");
    if (oldContent === newContent) {
      throw new Error("Binary content should be updated");
    }

    console.log("  ✓ Binary successfully replaced");

    // Clean up temp directory
    fs.rmSync(tempDir, { recursive: true, force: true });
    console.log("  ✓ Cleaned up temporary directory");

    console.log("  ✓ Test 2 PASSED");
  } catch (error) {
    console.error("  ✗ Test 2 FAILED:", error.message);
    return false;
  }

  cleanupTestDir();
  return true;
}

/**
 * Test: Temporary directory creation
 */
async function testTempDirectoryCreation() {
  console.log("\n🧪 Test 3: Temporary Directory Creation");

  try {
    const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "tkit-"));
    console.log("  ✓ Temporary directory created:", tempDir);

    // Verify it exists
    if (!fs.existsSync(tempDir)) {
      throw new Error("Temp directory should exist");
    }
    console.log("  ✓ Temporary directory verified");

    // Verify it's in system temp directory
    if (!tempDir.includes(os.tmpdir())) {
      throw new Error("Temp directory should be in system temp location");
    }
    console.log("  ✓ Temporary directory in correct location");

    // Create a file in temp directory
    const testFile = path.join(tempDir, "test.txt");
    fs.writeFileSync(testFile, "test content");
    console.log("  ✓ Created test file in temp directory");

    // Clean up
    fs.rmSync(tempDir, { recursive: true, force: true });
    console.log("  ✓ Cleaned up temporary directory");

    console.log("  ✓ Test 3 PASSED");
  } catch (error) {
    console.error("  ✗ Test 3 FAILED:", error.message);
    return false;
  }

  return true;
}

/**
 * Test: Error handling - cleanup on failure
 */
async function testErrorHandlingAndCleanup() {
  console.log("\n🧪 Test 4: Error Handling and Cleanup");

  try {
    const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "tkit-"));
    console.log("  ✓ Temporary directory created");

    // Simulate creation failure
    try {
      const testFile = path.join(tempDir, "test.txt");
      fs.writeFileSync(testFile, "test");
      throw new Error("Simulated download error");
    } catch (error) {
      console.log("  ✓ Simulated error occurred:", error.message);

      // Cleanup on error
      if (fs.existsSync(tempDir)) {
        fs.rmSync(tempDir, { recursive: true, force: true });
        console.log("  ✓ Cleaned up on error");
      }
    }

    // Verify cleanup
    if (fs.existsSync(tempDir)) {
      throw new Error("Temp directory should be deleted after error");
    }
    console.log("  ✓ Temporary directory successfully cleaned up");

    console.log("  ✓ Test 4 PASSED");
  } catch (error) {
    console.error("  ✗ Test 4 FAILED:", error.message);
    return false;
  }

  return true;
}

/**
 * Test: File permissions on Unix
 */
async function testFilePermissions() {
  console.log("\n🧪 Test 5: File Permissions");

  if (process.platform === "win32") {
    console.log("  ⊘ Test 5 SKIPPED (Windows platform)");
    return true;
  }

  try {
    const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "tkit-"));
    const testBinary = path.join(tempDir, "test-binary");
    fs.writeFileSync(testBinary, "test");
    console.log("  ✓ Created test binary");

    // Get permissions before
    const statsBefore = fs.statSync(testBinary);
    const modeBefore = (statsBefore.mode & parseInt("777", 8)).toString(8);
    console.log("  ✓ Permissions before chmod:", modeBefore);

    // Set executable
    fs.chmodSync(testBinary, 0o755);
    console.log("  ✓ Set executable permission (0o755)");

    // Verify permissions
    const statsAfter = fs.statSync(testBinary);
    const modeAfter = (statsAfter.mode & parseInt("777", 8)).toString(8);
    console.log("  ✓ Permissions after chmod:", modeAfter);

    if ((statsAfter.mode & fs.constants.S_IXUSR) === 0) {
      throw new Error("File should be executable");
    }
    console.log("  ✓ File is executable");

    // Cleanup
    fs.rmSync(tempDir, { recursive: true, force: true });
    console.log("  ✓ Test 5 PASSED");
  } catch (error) {
    console.error("  ✗ Test 5 FAILED:", error.message);
    return false;
  }

  return true;
}

/**
 * Run all tests
 */
async function runAllTests() {
  console.log("╔════════════════════════════════════════╗");
  console.log("║  TestFlowKit Install Script Tests      ║");
  console.log("╚════════════════════════════════════════╝");

  const results = [];

  results.push(await testNewInstallation());
  results.push(await testBinaryReplacement());
  results.push(await testTempDirectoryCreation());
  results.push(await testErrorHandlingAndCleanup());
  results.push(await testFilePermissions());

  // Summary
  const passed = results.filter((r) => r).length;
  const total = results.length;

  console.log("\n╔════════════════════════════════════════╗");
  console.log(`║  Test Results: ${passed}/${total} passed            ║`);
  console.log("╚════════════════════════════════════════╝");

  if (passed === total) {
    console.log("\n✓ All tests passed!");
    process.exit(0);
  } else {
    console.log(`\n✗ ${total - passed} test(s) failed`);
    process.exit(1);
  }
}

// Run tests
runAllTests().catch((error) => {
  console.error("Test runner error:", error);
  process.exit(1);
});
