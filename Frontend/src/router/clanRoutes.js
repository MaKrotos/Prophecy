const clanRoutes = [
  {
    path: "/clan",
    name: "clan",
    component: () => import("../views/ClanView.vue"),
    meta: {
      title: "Клан",
      order: 6,
    },
  },
];

export default clanRoutes;