import {
  createWebHistory,
  createRouter,
  type RouteRecordRaw,
} from "vue-router";

import HomeView from "./HomeView.vue";

export const routeNames = {
  home: "home",
};

const routes: readonly RouteRecordRaw[] = [
  { path: "/", component: HomeView, name: routeNames.home },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
