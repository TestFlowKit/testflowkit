import { execSync } from "node:child_process";

const VALIDATE_TIMEOUT_MS = 60_000;

export type ValidateResult = {
  ok: boolean;
  exitCode: number;
  output: string;
};

/**
 * Runs `tkit validate` for the project's testflowkit.yml.
 *
 * Uses /bin/sh so the command inherits the user's PATH from the shell environment.
 * cwd is set to the project directory so relative config paths resolve correctly.
 *
 * Unlike catalog/schema fetches, a non-zero exit here means "validation
 * failed" — an expected outcome, not a tool error — so it is captured and
 * returned instead of thrown.
 */
export function runValidate(
  cliBinary: string,
  cwd: string,
  configPath: string,
  tags?: string,
): ValidateResult {
  const tagsArg = tags ? ` --tags ${shellQuote(tags)}` : "";
  const command = `${shellQuote(cliBinary)} validate --config ${shellQuote(configPath)}${tagsArg}`;
  try {
    const output = execSync(command, {
      shell: "/bin/sh",
      cwd,
      encoding: "utf-8",
      timeout: VALIDATE_TIMEOUT_MS,
      stdio: ["ignore", "pipe", "pipe"],
    });
    return { ok: true, exitCode: 0, output: output.trim() };
  } catch (error) {
    const err = error as NodeJS.ErrnoException & {
      stdout?: string;
      stderr?: string;
      status?: number | null;
    };
    const output = `${err.stdout ?? ""}${err.stderr ?? ""}`.trim();
    return {
      ok: false,
      exitCode: typeof err.status === "number" ? err.status : 1,
      output: output || err.message,
    };
  }
}

// Single-quote the value and escape embedded single quotes for /bin/sh.
function shellQuote(value: string): string {
  return `'${value.replace(/'/g, `'\\''`)}'`;
}
