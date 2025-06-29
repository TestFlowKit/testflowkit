import "highlight.js/styles/atom-one-dark.css";

import hljs from "highlight.js";
import bash from "highlight.js/lib/languages/bash";
import gherkin from "highlight.js/lib/languages/gherkin";
import yaml from "highlight.js/lib/languages/yaml";

hljs.registerLanguage("bash", bash);
hljs.registerLanguage("gherkin", gherkin);
hljs.registerLanguage("yaml", yaml);

export default defineNuxtPlugin((nuxtApp) => {
  const highlightElement = (el: HTMLElement) => {
    const blocks = el.querySelectorAll("pre code");
    blocks.forEach((block) => {
      hljs.highlightElement(block as HTMLElement);
    });
  };

  const rehighlight = () => {
    const blocks = document.querySelectorAll("pre code");
    blocks.forEach((block) => {
      hljs.highlightElement(block as HTMLElement);
    });
  };

  nuxtApp.vueApp.directive("highlight", {
    mounted(el) {
      highlightElement(el);
    },
    updated(el) {
      highlightElement(el);
    },
  });

  return {
    provide: {
      rehighlight,
    },
  };
});
