import { CallToolResult } from "@modelcontextprotocol/sdk/types";
import { ProjectConfig } from "../config/types.js";
import z from "zod";

export type HandlerParams<T> = {
  config: ProjectConfig;
  configDir: string;
  input: T;
};

export interface TkitTool<T extends z.ZodType> {
  getName(): string;
  getDescription(): string;
  getInputSchema(): T;
  handler: (params: HandlerParams<z.infer<T>>) => Promise<CallToolResult>;
}
