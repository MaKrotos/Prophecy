const sessionRoutes = [
  {
    path: "/sessions",
    name: "sessions",
    component: () => import("../views/SessionsView.vue"),
    meta: {
      title: "Сессии",
    },
  },
  {
    path: "/sessions/create",
    name: "create-session",
    component: () => import("../views/CreateSessionView.vue"),
    meta: {
      title: "Создать сессию",
    },
  },
  {
    path: "/sessions/:id(\\d+)",
    name: "session-detail",
    component: () => import("../views/SessionDetailView.vue"),
    meta: {
      title: "Детали сессии",
    },
  },
];

export default sessionRoutes;