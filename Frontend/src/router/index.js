import { createRouter, createWebHashHistory } from "vue-router";
import mainRoutes from "./mainRoutes";
import sessionRoutes from "./sessionRoutes";
import playerRoutes from "./playerRoutes";
import qrRoutes from "./qrRoutes";
import otherRoutes from "./otherRoutes";
import clanRoutes from "./clanRoutes";

const routes = [
  ...mainRoutes,
  ...sessionRoutes,
  ...playerRoutes,
  ...qrRoutes,
  ...otherRoutes,
  ...clanRoutes
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
