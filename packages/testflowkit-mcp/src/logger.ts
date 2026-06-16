import * as process from "node:process";
type LogLevel = "log" | "debug" | "warn" | "error";

const PREFIX = "[testflowkit-mcp]";
const DEBUG_ENABLED = process.env.TESTFLOWKIT_MCP_DEBUG === "1";

export class Logger {
  private static write(level: LogLevel, message: string): void {
    process.stderr.write(`${PREFIX} [${level}] ${message}\n`);
  }

  static log(message: string): void {
    Logger.write("log", message);
  }

  static debug(message: string): void {
    if (!DEBUG_ENABLED) return;
    Logger.write("debug", message);
  }

  static warn(message: string): void {
    Logger.write("warn", message);
  }

  static error(message: string): void {
    Logger.write("error", message);
  }
}
