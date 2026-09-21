# TestFlowKit documentation site

Nuxt 4 + `@nuxt/content` site, deployed to GitHub Pages: https://testflowkit.github.io/testflowkit/

## Run locally

```bash
yarn install
yarn dev        # http://localhost:3000
yarn generate   # static build, checks that every page renders
```

## Where things live

| What | Where |
|---|---|
| Doc pages (Markdown) | `content/docs/<section>/<page>.md` |
| Step catalog (one JSON per step) | `content/sentences/<category>/*.json` |
| Sidebar, prev/next, docs home shortcuts | `navigation.ts` |
| Old URL redirects | `nuxt.config.ts` |

Do not start a page with a `# Title` line: the page template already renders `title` and `description` from the frontmatter as the H1 and subtitle.

## Adding a how-to

Create `content/docs/how-to/<verb-the-task>.md` with the sections **Goal**, steps with code, **Verify**, **See also**, then add it to `navigation.ts`. If you move or rename a page, add a redirect in `nuxt.config.ts`.

The project README is at [../readme.md](../readme.md).
