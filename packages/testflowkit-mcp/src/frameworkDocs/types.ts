export type FrameworkDocEntry = {
  slug: string;
  title: string;
  description: string;
  filePath: string;
  content: string;
};

export type FrameworkDocsIndex = Map<string, FrameworkDocEntry>;
