const clanRoutes = [
  {
    path: "/clan",
    name: "clan",
    component: () => import("../views/ClanView.vue"),
    meta: {
      title: "Клан",
    },
  },
  {
    path: "/sessions/:id/clan",
    name: "session-clan",
    component: () => import("../views/ClanView.vue"),
    meta: {
      title: "Клан",
    },
  },
];

export default clanRoutes;