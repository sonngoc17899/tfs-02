import { createRouter, createWebHistory } from "vue-router";


const routes = [

  {
    path: "/",
    name: "Shopping",
    component: () =>
      import("../views/Shopping.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
