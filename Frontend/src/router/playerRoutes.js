const playerRoutes = [
  {
    path: "/player-sessions",
    name: "player-sessions",
    component: () => import("../views/PlayerSessionsView.vue"),
    meta: {
      title: "Мои Сессии",
    },
  },
  {
    path: "/session/role",
    name: "session-role",
    component: () => import("../views/SessionRoleView.vue"),
    meta: {
      title: "Моя роль",
    },
  },
  {
    path: "/session/friends",
    name: "session-friends",
    component: () => import("../views/SessionFriendsView.vue"),
    meta: {
      title: "Список друзей",
    },
  },
  {
    path: "/session/role/:id",
    name: "session-role-with-id",
    component: () => import("../views/SessionRoleView.vue"),
    meta: {
      title: "Моя роль",
    },
  },
  {
    path: "/session/friends/:id",
    name: "session-friends-with-id",
    component: () => import("../views/SessionFriendsView.vue"),
    meta: {
      title: "Список друзей",
    },
  },
];

export default playerRoutes;