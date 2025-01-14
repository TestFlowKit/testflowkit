import "./style.css";
import "highlight.js/styles/atom-one-dark.css";
import hljs from "highlight.js/lib/core";
import gherkin from "highlight.js/lib/languages/gherkin";
import bash from "highlight.js/lib/languages/bash";
import hljsVuePlugin from "@highlightjs/vue-plugin";
import { createApp } from "vue";
import App from "./App.vue";
import { router } from "./views";

hljs.registerLanguage("gherkin", gherkin);
hljs.registerLanguage("bash", bash);

const app = createApp(App);

app.use(router);
app.use(hljsVuePlugin);

app.mount("#app");
