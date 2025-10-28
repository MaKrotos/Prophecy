const otherRoutes = [
  {
    path: "/users",
    name: "users",
    component: () => import("../views/UsersView.vue"),
    meta: {
      title: "Все пользователи",
      order: 6,
    },
  },
  {
    path: "/rules",
    name: "rules",
    component: () => import("../views/RulesView.vue"),
    meta: {
      title: "Правила Игры",
      order: 11,
    },
  },
  {
    path: "/sessions/join/:referral_link",
    name: "session-join",
    component: () => import("../views/SessionJoinView.vue"),
    meta: {
      title: "Присоединиться к сессии",
      order: 10,
    },
    beforeEnter: (to, from, next) => {
      console.log(
        "🔍 Роутер: Переход на страницу присоединения к сессии",
        to.params.referral_link
      );
      next();
    },
  },
];

export default otherRoutes;