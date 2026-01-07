#!/usr/bin/env node

const { spawn } = require("child_process");
const path = require("path");

const BINARY_NAME = process.platform === "win32" ? "tkit.exe" : "tkit";
// Store the actual binary in a 'cli' subdirectory to avoid naming conflict
const binaryPath = path.join(__dirname, "..", "cli", BINARY_NAME);

// Pass all arguments to the binary
const child = spawn(binaryPath, process.argv.slice(2), {
  stdio: "inherit",
  windowsHide: true,
});

child.on("exit", (code) => {
  process.exit(code || 0);
});
