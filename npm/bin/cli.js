#!/usr/bin/env node

"use strict";

const { spawn } = require("node:child_process");
const { getBinaryPath } = require("../lib/getBinaryPath");

const binaryPath = getBinaryPath();

const child = spawn(binaryPath, process.argv.slice(2), {
  stdio: "inherit",
  windowsHide: true,
});

child.on("exit", (code) => {
  process.exit(code ?? 0);
});
