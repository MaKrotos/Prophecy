import { createRouter, createWebHashHistory } from "vue-router";
import mainRoutes from "./mainRoutes";
import sessionRoutes from "./sessionRoutes";
import playerRoutes from "./playerRoutes";
import qrRoutes from "./qrRoutes";
import otherRoutes from "./otherRoutes";
import clanRoutes from "./clanRoutes";

// Объединяем маршруты в правильном порядке
// Важно, чтобы /sessions/join/:referral_link шел перед /sessions/:id
const routes = [
  ...mainRoutes,
  playerRoutes[0], // /player-sessions
  sessionRoutes[0], // /sessions
  sessionRoutes[1], // /sessions/create
  qrRoutes[0], // /scan
  qrRoutes[1], // /session/my-qr/:id
  qrRoutes[2], // /session/qr-scanner
  otherRoutes[0], // /users
  otherRoutes[1], // /rules
  otherRoutes[2], // /sessions/join/:referral_link (важно, чтобы шел перед /sessions/:id)
  playerRoutes[1], // /session/role
  playerRoutes[2], // /session/friends
  playerRoutes[3], // /session/role/:id
  playerRoutes[4], // /session/friends/:id
  sessionRoutes[2], // /sessions/:id (должен идти после /sessions/join/:referral_link)
  clanRoutes[0], // /clan
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
