import HomeView from "../views/HomeView.vue";

const mainRoutes = [
  {
    path: "/",
    name: "home",
    component: HomeView,
    meta: {
      title: "Главная",
    },
  },
  {
    path: "/profile",
    name: "profile",
    component: () => import("../views/ProfileView.vue"),
    meta: {
      title: "Профиль",
    },
  },
  {
    path: "/settings",
    name: "settings",
    component: () => import("../views/SettingsView.vue"),
    meta: {
      title: "Настройки",
    },
  },
];

export default mainRoutes;