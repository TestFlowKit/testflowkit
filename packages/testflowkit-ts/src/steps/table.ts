import { Then, When, type DataTable } from '@cucumber/cucumber';
import type { Locator, Page } from 'playwright';
import type { TestFlowKitWorld } from '../world.js';
import { substituteVariables } from '../variables/substitute.js';

function buildRowLocator(page: Page, cellSelector: 'td' | 'th', values: string[]): Locator {
  return values.reduce(
    (row, value) => row.filter({ has: page.locator(cellSelector, { hasText: value }) }),
    page.locator('tr'),
  );
}

async function findTableRow(
  world: TestFlowKitWorld,
  cellSelector: 'td' | 'th',
  values: string[],
): Promise<Locator> {
  const row = buildRowLocator(world.page, cellSelector, values).first();
  try {
    await row.waitFor({ state: 'attached', timeout: world.browserSettings.timeout });
  } catch {
    throw new Error(`Row not found containing: ${values.join(', ')}`);
  }
  if (!(await row.isVisible())) {
    throw new Error(`Row is not visible containing: ${values.join(', ')}`);
  }
  return row;
}

function rowValues(world: TestFlowKitWorld, row: Record<string, string>): string[] {
  return Object.values(row).map((value) => substituteVariables(value, world.variables));
}

When(
  'the user clicks on the row containing the following elements',
  async function (this: TestFlowKitWorld, table: DataTable) {
    for (const row of table.hashes()) {
      const element = await findTableRow(this, 'td', rowValues(this, row));
      await element.click();
    }
  },
);

Then(
  'the user should see a row containing the following elements',
  async function (this: TestFlowKitWorld, table: DataTable) {
    for (const row of table.hashes()) {
      await findTableRow(this, 'td', rowValues(this, row));
    }
  },
);

Then(
  'the user should not see a row containing the following elements',
  async function (this: TestFlowKitWorld, table: DataTable) {
    for (const row of table.hashes()) {
      const values = rowValues(this, row);
      let found = true;
      try {
        await findTableRow(this, 'td', values);
      } catch {
        found = false;
      }
      if (found) {
        throw new Error(`Row containing the specified elements was found but should not be visible: ${values.join(', ')}`);
      }
    }
  },
);

Then(
  'the user should see a table with the following headers',
  async function (this: TestFlowKitWorld, table: DataTable) {
    const headers = Object.values(table.rowsHash()).map((value) => substituteVariables(value, this.variables));
    await findTableRow(this, 'th', headers);
  },
);
