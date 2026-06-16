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
    title: 'Getting Started',
    children: [
      { path: '/docs/getting-started/introduction', title: 'Introduction' },
      { path: '/docs/getting-started/installation', title: 'Installation' },
      { path: '/docs/getting-started/quick-start', title: 'Quick Start' },
    ],
  },
  {
    title: 'Guides',
    children: [
      { path: '/docs/guides/writing-tests', title: 'Writing Tests' },
      { path: '/docs/guides/frontend-testing', title: 'Frontend Testing' },
      { path: '/docs/guides/api-testing', title: 'API Testing' },
      { path: '/docs/guides/ide-agent', title: 'IDE Agent' },
    ],
  },
  {
    title: 'Configuration',
    children: [
      { path: '/docs/config/overview', title: 'testflowkit.yml' },
      { path: '/docs/config/selectors', title: 'Selectors' },
    ],
  },
  {
    title: 'Patterns',
    children: [
      { path: '/docs/patterns/variables', title: 'Variables' },
      { path: '/docs/patterns/random-data', title: 'Random Data' },
      { path: '/docs/patterns/macros', title: 'Macros' },
      { path: '/docs/patterns/global-hooks', title: 'Global Hooks' },
      { path: '/docs/patterns/skip-tag', title: 'Skipping Tests' },
    ],
  },
  {
    title: 'Reference',
    children: [
      { path: '/docs/reference/cli', title: 'CLI Reference' },
      { path: '/sentences', title: 'Step Catalog', external: true },
    ],
  },
  {
    title: 'Troubleshooting',
    children: [
      { path: '/docs/troubleshooting/common-issues', title: 'Common Issues' },
      { path: '/docs/troubleshooting/platform-issues', title: 'Platform Issues' },
      { path: '/docs/troubleshooting/migration-guide', title: 'Migration Guide' },
    ],
  },
];

/** Flat list of doc pages in sidebar order (for prev/next navigation). */
export const allDocPages: NavItem[] = docNavigation.flatMap((group) =>
  group.children.filter((item) => !item.external),
);

export const docHubPaths = {
  qa: [
    { path: '/docs/getting-started/quick-start', title: 'Quick Start' },
    { path: '/docs/guides/writing-tests', title: 'Writing Tests' },
    { path: '/docs/guides/frontend-testing', title: 'Frontend Testing' },
    { path: '/sentences', title: 'Step Catalog' },
  ],
  developer: [
    { path: '/docs/getting-started/installation', title: 'Installation' },
    { path: '/docs/config/overview', title: 'testflowkit.yml' },
    { path: '/docs/guides/api-testing', title: 'API Testing' },
    { path: '/docs/reference/cli', title: 'CLI Reference' },
  ],
  ai: [
    { path: '/docs/guides/ide-agent', title: 'IDE Agent' },
    { path: '/docs/config/overview', title: 'testflowkit.yml' },
    { path: '/sentences', title: 'Step Catalog' },
  ],
};
