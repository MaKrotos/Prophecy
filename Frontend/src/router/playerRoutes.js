const playerRoutes = [
  {
    path: "/player-sessions",
    name: "player-sessions",
    component: () => import("../views/PlayerSessionsView.vue"),
    meta: {
      title: "Мои Сессии",
      order: 2,
    },
  },
  {
    path: "/session/role",
    name: "session-role",
    component: () => import("../views/SessionRoleView.vue"),
    meta: {
      title: "Моя роль",
      order: 12,
    },
  },
  {
    path: "/session/friends",
    name: "session-friends",
    component: () => import("../views/SessionFriendsView.vue"),
    meta: {
      title: "Список друзей",
      order: 13,
    },
  },
];

export default playerRoutes;