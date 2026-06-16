import { execSync } from "node:child_process";

const LOCAL_CLI_TIMEOUT_MS = 30_000;

/**
 * Calls the testflowkit CLI to export step definitions as JSON.
 *
 * Uses /bin/sh so the command inherits the user's PATH from the shell environment.
 * cwd is set to the project directory so relative config paths resolve correctly.
 * If the binary is not found via PATH, set TESTFLOWKIT_CLI_PATH to its absolute path.
 *
 * Throws on non-zero exit code or if the binary is not found.
 */
export function fetchFromLocalCli(cliBinary: string, cwd: string): string {
  return execSync(`${cliBinary} export-step-definitions --format json`, {
    shell: "/bin/sh",
    cwd,
    encoding: "utf-8",
    timeout: LOCAL_CLI_TIMEOUT_MS,
    stdio: ["ignore", "pipe", "ignore"],
  }).trim();
}
