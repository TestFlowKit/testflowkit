/**
 * @testflowkit/cli
 *
 * This package provides the TestFlowKit CLI (tkit) for running
 * end-to-end tests using Gherkin syntax.
 *
 * Usage:
 *   npx @testflowkit/cli run
 *   npx @testflowkit/cli init
 *   npx @testflowkit/cli validate
 *   npx @testflowkit/cli --version
 *
 * For programmatic usage, you can get the path to the binary:
 */

const path = require("path");
const fs = require("fs");

const BINARY_NAME = process.platform === "win32" ? "tkit.exe" : "tkit";
const binaryPath = path.join(__dirname, "..", "bin", BINARY_NAME);

/**
 * Get the path to the tkit binary
 * @returns {string} Absolute path to the tkit binary
 * @throws {Error} If the binary is not found
 */
function getBinaryPath() {
  if (!fs.existsSync(binaryPath)) {
    throw new Error(
      `tkit binary not found. Please reinstall @testflowkit/cli`
    );
  }
  return binaryPath;
}

/**
 * Check if the binary is installed
 * @returns {boolean} True if the binary exists
 */
function isInstalled() {
  return fs.existsSync(binaryPath);
}

module.exports = {
  getBinaryPath,
  isInstalled,
  binaryPath,
};
