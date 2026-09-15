import type { Page } from 'playwright';
import type { TestFlowKitWorld } from '../world.js';

/** All pages currently open across every browser context created in this scenario, in the order they were opened. */
export function allPages(world: TestFlowKitWorld): Page[] {
  return world.contexts.flatMap((context) => context.pages());
}

const DURATION_PATTERN = /^(\d+)(ms|s|m|h)$/;

/** Parses simple durations like "500ms", "5s", "2m", "1h" into milliseconds. */
export function parseDurationMs(value: string): number {
  const match = DURATION_PATTERN.exec(value.trim());
  if (!match) {
    throw new Error(`Invalid duration format: "${value}"`);
  }
  const amount = Number(match[1]);
  const unitMs: Record<string, number> = { ms: 1, s: 1000, m: 60_000, h: 3_600_000 };
  return amount * unitMs[match[2]!]!;
}

const POLL_INTERVAL_MS = 100;

/** Polls until a new page appears in any tracked context, or throws once `timeoutMs` elapses. */
export async function waitForNewPage(world: TestFlowKitWorld, timeoutMs: number): Promise<Page> {
  const initialCount = allPages(world).length;
  const deadline = Date.now() + timeoutMs;

  while (Date.now() < deadline) {
    const pages = allPages(world);
    if (pages.length > initialCount) {
      return pages.at(-1)!;
    }
    await new Promise((resolve) => setTimeout(resolve, POLL_INTERVAL_MS));
  }

  throw new Error(`No new window opened within ${timeoutMs}ms`);
}
