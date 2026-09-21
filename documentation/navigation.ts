export type NavItem = {
  path: string;
  title: string;
  external?: boolean;
};

export type NavGroup = {
  title: string;
  children: NavItem[];
};

export const docNavigation: NavGroup[] = [
  {
    title: "Getting Started",
    children: [
      { path: "/docs/getting-started/introduction", title: "Introduction" },
      { path: "/docs/getting-started/installation", title: "Installation" },
      { path: "/docs/getting-started/quick-start", title: "Quick Start" },
    ],
  },
  {
    title: "Recipes: API",
    children: [
      {
        path: "/docs/how-to/add-a-rest-endpoint",
        title: "Add a REST Endpoint",
      },
      {
        path: "/docs/how-to/add-a-graphql-operation",
        title: "Add a GraphQL Operation",
      },
      {
        path: "/docs/how-to/configure-authentication",
        title: "Configure Authentication",
      },
      {
        path: "/docs/how-to/authenticate-before-tests",
        title: "Authenticate Before Tests",
      },
      {
        path: "/docs/how-to/send-a-json-body-or-file",
        title: "Send a JSON Body or File",
      },
    ],
  },
  {
    title: "Recipes: UI",
    children: [
      { path: "/docs/how-to/fill-in-a-form", title: "Fill in a Form" },
      {
        path: "/docs/how-to/wait-for-an-element",
        title: "Wait for an Element",
      },
      {
        path: "/docs/how-to/upload-a-file-in-the-browser",
        title: "Upload a File (Browser)",
      },
      {
        path: "/docs/how-to/capture-a-failure-screenshot",
        title: "Screenshot on Failure",
      },
    ],
  },
  {
    title: "Recipes: Organize & Run",
    children: [
      {
        path: "/docs/how-to/test-multiple-environments",
        title: "Test Multiple Environments",
      },
      {
        path: "/docs/how-to/share-data-between-scenarios",
        title: "Share Data Between Scenarios",
      },
      { path: "/docs/how-to/skip-a-test", title: "Skip a Test" },
      { path: "/docs/how-to/run-in-parallel", title: "Run in Parallel" },
      { path: "/docs/how-to/run-in-ci", title: "Run in CI" },
      {
        path: "/docs/how-to/debug-a-failing-scenario",
        title: "Debug a Failing Scenario",
      },
    ],
  },
  {
    title: "Concepts & Guides",
    children: [
      { path: "/docs/guides/writing-tests", title: "Writing Tests" },
      { path: "/docs/guides/frontend-testing", title: "Frontend Testing" },
      { path: "/docs/guides/api-testing", title: "API Testing" },
      { path: "/docs/guides/ide-agent", title: "IDE Agent" },
    ],
  },
  {
    title: "Configuration",
    children: [
      { path: "/docs/config/overview", title: "testflowkit.yml" },
      { path: "/docs/config/selectors", title: "Selectors" },
    ],
  },
  {
    title: "Patterns",
    children: [
      { path: "/docs/patterns/variables", title: "Variables" },
      { path: "/docs/patterns/random-data", title: "Random Data" },
      { path: "/docs/patterns/macros", title: "Macros" },
      { path: "/docs/patterns/global-hooks", title: "Global Hooks" },
    ],
  },
  {
    title: "Reference",
    children: [
      { path: "/docs/reference/cli", title: "CLI Reference" },
      { path: "/docs/reference/reporters", title: "Reporters" },
      { path: "/docs/reference/glossary", title: "Glossary" },
      { path: "/sentences", title: "Step Catalog", external: true },
    ],
  },
  {
    title: "Help",
    children: [
      { path: "/docs/troubleshooting/faq", title: "FAQ" },
      { path: "/docs/troubleshooting/common-issues", title: "Common Issues" },
      {
        path: "/docs/troubleshooting/platform-issues",
        title: "Platform Issues",
      },
      {
        path: "/docs/troubleshooting/migration-guide",
        title: "Migration Guide",
      },
      { path: "/docs/changelog", title: "Changelog" },
    ],
  },
];

/** Flat list of doc pages in sidebar order (for prev/next navigation). */
export const allDocPages: NavItem[] = docNavigation.flatMap((group) =>
  group.children.filter((item) => !item.external),
);

export const docHubPaths = {
  qa: [
    { path: "/docs/getting-started/quick-start", title: "Quick Start" },
    { path: "/docs/guides/writing-tests", title: "Writing Tests" },
    { path: "/docs/guides/frontend-testing", title: "Frontend Testing" },
    { path: "/sentences", title: "Step Catalog" },
  ],
  developer: [
    { path: "/docs/getting-started/installation", title: "Installation" },
    { path: "/docs/how-to/add-a-rest-endpoint", title: "Add a REST Endpoint" },
    {
      path: "/docs/how-to/configure-authentication",
      title: "Configure Authentication",
    },
    { path: "/docs/reference/cli", title: "CLI Reference" },
  ],
  ai: [
    { path: "/docs/guides/ide-agent", title: "IDE Agent" },
    { path: "/docs/config/overview", title: "testflowkit.yml" },
    { path: "/sentences", title: "Step Catalog" },
  ],
};

/** Task-oriented shortcuts surfaced on the docs home page, above the role-based hub. */
export const howToShortcuts: NavItem[] = [
  { path: "/docs/how-to/add-a-rest-endpoint", title: "Add a REST endpoint" },
  {
    path: "/docs/how-to/add-a-graphql-operation",
    title: "Add a GraphQL operation",
  },
  {
    path: "/docs/how-to/configure-authentication",
    title: "Configure authentication",
  },
  {
    path: "/docs/how-to/authenticate-before-tests",
    title: "Authenticate before tests",
  },
  {
    path: "/docs/how-to/send-a-json-body-or-file",
    title: "Send a JSON body or file",
  },
  {
    path: "/docs/how-to/test-multiple-environments",
    title: "Test multiple environments",
  },
  {
    path: "/docs/how-to/share-data-between-scenarios",
    title: "Share data between scenarios",
  },
  { path: "/docs/how-to/skip-a-test", title: "Skip a test" },
  { path: "/docs/how-to/fill-in-a-form", title: "Fill in a form" },
  { path: "/docs/how-to/wait-for-an-element", title: "Wait for an element" },
  {
    path: "/docs/how-to/upload-a-file-in-the-browser",
    title: "Upload a file in the browser",
  },
  {
    path: "/docs/how-to/capture-a-failure-screenshot",
    title: "Capture a screenshot on failure",
  },
  { path: "/docs/how-to/run-in-parallel", title: "Run in parallel" },
  { path: "/docs/how-to/run-in-ci", title: "Run in CI" },
  {
    path: "/docs/how-to/debug-a-failing-scenario",
    title: "Debug a failing scenario",
  },
];
