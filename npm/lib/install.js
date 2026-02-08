#!/usr/bin/env node

const https = require("https");
const http = require("http");

const fs = require("fs");
const path = require("path");
const os = require("os");
const AdmZip = require("adm-zip");
const PACKAGE_JSON = require("../package.json");
const VERSION = PACKAGE_JSON.version;
const BINARY_NAME = "tkit";
const GITHUB_RELEASE_URL = `https://github.com/TestFlowKit/testflowkit/releases/download/${VERSION}`;

// Platform and architecture mapping from Node.js to Go naming conventions
const PLATFORM_MAP = {
  darwin: "darwin",
  linux: "linux",
  win32: "windows",
};

const ARCH_MAP = {
  x64: "amd64",
  arm64: "arm64",
};

/**
 * Get the binary name for the current platform
 */
function getBinaryName() {
  return process.platform === "win32" ? `${BINARY_NAME}.exe` : BINARY_NAME;
}

/**
 * Get the download URL for the current platform and architecture
 */
function getDownloadUrl() {
  const platform = PLATFORM_MAP[process.platform];
  const arch = ARCH_MAP[process.arch];

  if (!platform) {
    throw new Error(`Unsupported platform: ${process.platform}`);
  }

  if (!arch) {
    throw new Error(`Unsupported architecture: ${process.arch}`);
  }

  const filename = `${BINARY_NAME}-${platform}-${arch}.zip`;
  return `${GITHUB_RELEASE_URL}/${filename}`;
}

/**
 * Follow redirects and download file
 */
function downloadFile(url, maxRedirects = 5) {
  return new Promise((resolve, reject) => {
    if (maxRedirects === 0) {
      reject(new Error("Too many redirects"));
      return;
    }

    const protocol = url.startsWith("https") ? https : require("http");

    protocol.get(url, (response) => handleDownloadResponse(response, resolve, reject, maxRedirects)).on("error", reject);
  });
}

/**
 * Extract zip file using adm-zip (cross-platform, no OS dependency)
 */
async function extractZip(zipBuffer, destDir) {
  try {
    const zip = new AdmZip(zipBuffer);
    zip.extractAllTo(destDir, true);
  } catch (error) {
    throw new Error(`Failed to extract zip: ${error.message}`);
  }
}

/**
 *
 * @param {http.IncomingMessage} response
 * @returns
 */
function handleDownloadResponse(response, resolve, reject, maxRedirects) {
  // Handle redirects (GitHub releases use redirects)
  const mustHandleRedirects =
    response.statusCode >= 300 &&
    response.statusCode < 400 &&
    response.headers.location;

  if (mustHandleRedirects) {
    downloadFile(response.headers.location, maxRedirects - 1)
      .then(resolve)
      .catch(reject);
    return;
  }

  if (response.statusCode !== 200) {
    reject(new Error(`Failed to download: HTTP ${response.statusCode}`));
    return;
  }

  const chunks = [];
  const totalBytes = parseInt(response.headers["content-length"] || "0", 10);
  let downloadedBytes = 0;
  let nextLogPercent = 0;

  if (totalBytes > 0) {
    console.log("Download progress: 0%");
    nextLogPercent = 10;
  }

  response.on("data", (chunk) => {
    chunks.push(chunk);
    downloadedBytes += chunk.length;

    if (totalBytes > 0) {
      const percent = Math.floor((downloadedBytes / totalBytes) * 100);
      while (percent >= nextLogPercent && nextLogPercent <= 100) {
        console.log(`Download progress: ${nextLogPercent}%`);
        nextLogPercent += 10;
      }
    }
  });

  response.on("end", () => {
    if (totalBytes <= 0 || nextLogPercent <= 100) {
      console.log("Download progress: 100%");
    }
    resolve(Buffer.concat(chunks));
  });
  response.on("error", reject);
}

/**
 * Make the binary executable (Unix only)
 */
function makeExecutable(filePath) {
  if (process.platform !== "win32") {
    fs.chmodSync(filePath, 0o755);
  }
}

/**
 * Main installation function
 */
async function install() {
  const binDir = path.join(__dirname, "..", "cli");
  const binaryName = getBinaryName();
  const binaryPath = path.join(binDir, binaryName);
  const binaryExists = fs.existsSync(binaryPath);

  // Ensure bin directory exists
  if (!fs.existsSync(binDir)) {
    fs.mkdirSync(binDir, { recursive: true });
  }

  const downloadUrl = getDownloadUrl();
  console.log(`Downloading ${BINARY_NAME} v${VERSION}...`);
  console.log(
    `  Platform: ${process.platform} (${PLATFORM_MAP[process.platform]})`,
  );
  console.log(`  Architecture: ${process.arch} (${ARCH_MAP[process.arch]})`);
  console.log(`  URL: ${downloadUrl}`);

  if (binaryExists) {
    console.log(
      `ℹ Existing ${BINARY_NAME} binary found, downloading new version...`,
    );
  }

  let tempDir;
  try {
    // Create temporary directory in system temp location
    tempDir = fs.mkdtempSync(path.join(os.tmpdir(), `${BINARY_NAME}-`));

    // Download the zip file
    const zipBuffer = await downloadFile(downloadUrl);
    console.log(`Downloaded ${(zipBuffer.length / 1024 / 1024).toFixed(2)} MB`);

    // Extract to temporary directory first
    console.log("Extracting...");
    await extractZip(zipBuffer, tempDir);

    // Get the extracted binary path from temp directory
    const tempBinaryPath = path.join(tempDir, binaryName);

    if (!fs.existsSync(tempBinaryPath)) {
      throw new Error("Binary not found after extraction");
    }

    // Make executable
    makeExecutable(tempBinaryPath);

    // Remove old binary if it exists, then move new binary
    if (binaryExists) {
      try {
        removeExistingBinary(binaryPath);
      } catch (error) {
        // If removal fails, clean up temp and rethrow
        fs.rmSync(tempDir, { recursive: true, force: true });
        throw error;
      }
    }

    // Move new binary from temp to final location
    fs.renameSync(tempBinaryPath, binaryPath);

    // Clean up temporary directory
    fs.rmSync(tempDir, { recursive: true, force: true });

    // Verify installation
    if (fs.existsSync(binaryPath)) {
      console.log(`✓ Successfully installed ${BINARY_NAME} v${VERSION}`);
    } else {
      throw new Error("Binary not found after installation");
    }
  } catch (error) {
    // Clean up temp directory on error
    if (tempDir && fs.existsSync(tempDir)) {
      fs.rmSync(tempDir, { recursive: true, force: true });
    }
    console.error(`✗ Failed to install ${BINARY_NAME}: ${error.message}`);
    console.error("");
    console.error("You can manually download the binary from:");
    console.error(GITHUB_RELEASE_URL);
    console.error("");
    console.error("Then place it in:");
    console.error(`  ${binDir}`);
    process.exit(1);
  }
}

function removeExistingBinary(binaryPath) {
  try {
    fs.unlinkSync(binaryPath);
  } catch (error) {
    console.error(
      `✗ Failed to remove existing ${BINARY_NAME} binary: ${error.message}`,
    );
    process.exit(1);
  }
}


// Run installation
install();
