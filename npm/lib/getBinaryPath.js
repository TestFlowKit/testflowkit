"use strict";

const path = require("node:path");

const SUPPORTED_PLATFORMS = {
  "linux-x64": { pkg: "@testflowkit/cli-linux-x64", bin: "tkit" },
  "linux-arm64": { pkg: "@testflowkit/cli-linux-arm64", bin: "tkit" },
  "darwin-x64": { pkg: "@testflowkit/cli-darwin-x64", bin: "tkit" },
  "darwin-arm64": { pkg: "@testflowkit/cli-darwin-arm64", bin: "tkit" },
  "win32-x64": { pkg: "@testflowkit/cli-win32-x64", bin: "tkit.exe" },
  "win32-arm64": { pkg: "@testflowkit/cli-win32-arm64", bin: "tkit.exe" },
};

/**
 * Returns the absolute path to the platform-specific tkit binary.
 * Throws a descriptive error if the platform is unsupported or the
 * optional native package was not installed.
 *
 * @param {string} [platform] - override process.platform (for testing)
 * @param {string} [arch]     - override process.arch (for testing)
 * @returns {string} absolute path to the binary
 */
function getBinaryPath(platform, arch) {
  const plt = platform || process.platform;
  const arc = arch || process.arch;
  const key = `${plt}-${arc}`;
  const entry = SUPPORTED_PLATFORMS[key];

  if (!entry) {
    throw new Error(
      `Unsupported platform: ${plt}-${arc}. ` +
        `TestFlowKit CLI does not provide a binary for this environment.`
    );
  }

  try {
    return require.resolve(path.join(entry.pkg, "bin", entry.bin));
  } catch {
    throw new Error(
      `Could not find the native package ${entry.pkg}. ` +
        `Please ensure you had an internet connection during 'npm install'. ` +
        `If you are on an unsupported platform (${plt}-${arc}), ` +
        `please open an issue at https://github.com/TestFlowKit/testflowkit/issues`
    );
  }
}

module.exports = { getBinaryPath, SUPPORTED_PLATFORMS };
