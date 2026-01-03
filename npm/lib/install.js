#!/usr/bin/env node

const https = require("https");
const fs = require("fs");
const path = require("path");
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

    protocol
      .get(url, (response) => {
        // Handle redirects (GitHub releases use redirects)
        if (
          response.statusCode >= 300 &&
          response.statusCode < 400 &&
          response.headers.location
        ) {
          downloadFile(response.headers.location, maxRedirects - 1)
            .then(resolve)
            .catch(reject);
          return;
        }

        if (response.statusCode !== 200) {
          reject(
            new Error(`Failed to download: HTTP ${response.statusCode}`)
          );
          return;
        }

        const chunks = [];
        response.on("data", (chunk) => chunks.push(chunk));
        response.on("end", () => resolve(Buffer.concat(chunks)));
        response.on("error", reject);
      })
      .on("error", reject);
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
  const binDir = path.join(__dirname, "..", "bin");
  const binaryName = getBinaryName();
  const binaryPath = path.join(binDir, binaryName);

  // Check if binary already exists
  if (fs.existsSync(binaryPath)) {
    console.log(`✓ ${BINARY_NAME} binary already exists`);
    return;
  }

  // Ensure bin directory exists
  if (!fs.existsSync(binDir)) {
    fs.mkdirSync(binDir, { recursive: true });
  }

  const downloadUrl = getDownloadUrl();
  console.log(`Downloading ${BINARY_NAME} v${VERSION}...`);
  console.log(`  Platform: ${process.platform} (${PLATFORM_MAP[process.platform]})`);
  console.log(`  Architecture: ${process.arch} (${ARCH_MAP[process.arch]})`);
  console.log(`  URL: ${downloadUrl}`);

  try {
    // Download the zip file
    const zipBuffer = await downloadFile(downloadUrl);
    console.log(`Downloaded ${(zipBuffer.length / 1024 / 1024).toFixed(2)} MB`);

    // Extract to bin directory
    console.log("Extracting...");
    await extractZip(zipBuffer, binDir);

    // Make executable
    makeExecutable(binaryPath);

    // Verify installation
    if (fs.existsSync(binaryPath)) {
      console.log(`✓ Successfully installed ${BINARY_NAME} v${VERSION}`);
    } else {
      throw new Error("Binary not found after extraction");
    }
  } catch (error) {
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

// Run installation
install();
