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
    title: "How-to",
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
      {
        path: "/docs/how-to/test-multiple-environments",
        title: "Test Multiple Environments",
      },
      {
        path: "/docs/how-to/share-data-between-scenarios",
        title: "Share Data Between Scenarios",
      },
      { path: "/docs/how-to/skip-a-test", title: "Skip a Test" },
    ],
  },
  {
    title: "Guides",
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
      { path: "/docs/patterns/skip-tag", title: "Skipping Tests" },
    ],
  },
  {
    title: "Reference",
    children: [
      { path: "/docs/reference/cli", title: "CLI Reference" },
      { path: "/sentences", title: "Step Catalog", external: true },
    ],
  },
  {
    title: "Troubleshooting",
    children: [
      { path: "/docs/troubleshooting/common-issues", title: "Common Issues" },
      {
        path: "/docs/troubleshooting/platform-issues",
        title: "Platform Issues",
      },
      {
        path: "/docs/troubleshooting/migration-guide",
        title: "Migration Guide",
      },
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
];
