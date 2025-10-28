import HomeView from "../views/HomeView.vue";

const mainRoutes = [
  {
    path: "/",
    name: "home",
    component: HomeView,
    meta: {
      title: "Главная",
      order: 1,
    },
  },
  {
    path: "/profile",
    name: "profile",
    component: () => import("../views/ProfileView.vue"),
    meta: {
      title: "Профиль",
      order: 3,
    },
  },
  {
    path: "/settings",
    name: "settings",
    component: () => import("../views/SettingsView.vue"),
    meta: {
      title: "Настройки",
      order: 4,
    },
  },
];

export default mainRoutes;