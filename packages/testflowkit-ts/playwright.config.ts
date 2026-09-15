import { defineConfig } from "@playwright/test";
import * as process from "node:process";
export default defineConfig({
  timeout: 30_000,
  use: {
    baseURL: process.env.FRONTEND_BASE_URL ?? "http://localhost:3000",
    headless: true,
    viewport: { width: 1280, height: 720 },
    locale: "en-US",
    timezoneId: "America/New_York",
  },
});
