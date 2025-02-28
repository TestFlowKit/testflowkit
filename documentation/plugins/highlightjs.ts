import "highlight.js/styles/atom-one-dark.css";

import hljs from "highlight.js";
import bash from "highlight.js/lib/languages/bash";
import gherkin from "highlight.js/lib/languages/gherkin";
import yaml from "highlight.js/lib/languages/yaml";

hljs.registerLanguage("bash", bash);
hljs.registerLanguage("gherkin", gherkin);
hljs.registerLanguage("yaml", yaml);

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.directive("highlight", {
    mounted(el) {
      const blocks = el.querySelectorAll("pre code");
      blocks.forEach((block: HTMLElement) => {
        hljs.highlightElement(block);
      });
    },
  });
});
