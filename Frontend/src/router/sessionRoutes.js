const sessionRoutes = [
  {
    path: "/sessions",
    name: "sessions",
    component: () => import("../views/SessionsView.vue"),
    meta: {
      title: "Сессии",
      order: 7,
    },
  },
  {
    path: "/sessions/create",
    name: "create-session",
    component: () => import("../views/CreateSessionView.vue"),
    meta: {
      title: "Создать сессию",
      order: 8,
    },
  },
  {
    path: "/sessions/:id",
    name: "session-detail",
    component: () => import("../views/SessionDetailView.vue"),
    meta: {
      title: "Детали сессии",
      order: 9,
    },
  },
];

export default sessionRoutes;