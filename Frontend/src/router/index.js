import { createRouter, createWebHashHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";

const routes = [
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
    path: "/player-sessions",
    name: "player-sessions",
    component: () => import("../views/PlayerSessionsView.vue"),
    meta: {
      title: "Мои Сессии",
      order: 2,
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
  {
    path: "/scan",
    name: "scan",
    component: () => import("../views/ScanView.vue"),
    meta: {
      title: "Сканер",
      order: 5,
    },
  },
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
  {
    path: "/session/my-qr",
    name: "session-my-qr",
    component: () => import("../views/SessionMyQRView.vue"),
    meta: {
      title: "Мой QR код",
      order: 14,
    },
  },
  {
    path: "/session/qr-scanner",
    name: "session-qr-scanner",
    component: () => import("../views/SessionQRScannerView.vue"),
    meta: {
      title: "Сканер QR кодов",
      order: 15,
    },
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

// Глобальный обработчик для заголовков страниц
router.beforeEach((to, from, next) => {
  console.log(
    "🔍 Роутер: переход со страницы",
    from.path,
    "на страницу",
    to.path
  );
  // Можно также установить заголовок документа здесь
  document.title = to.meta.title || "Мое Приложение";
  next();
});

// Глобальный обработчик для отладки навигации
router.afterEach((to, from) => {
  console.log(
    "✅ Роутер: переход завершен со страницы",
    from.path,
    "на страницу",
    to.path
  );
});

export default router;
