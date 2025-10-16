import { isValidTelegramWebApp } from "../core/validation.js";

/**
 * Функции для работы с данными аутентификации Telegram
 */

export function getTelegramAuthData() {
  if (!isValidTelegramWebApp()) return null;

  const webApp = window.Telegram.WebApp;
  return extractAuthData(webApp);
}

export function extractAuthData(webApp) {
  console.log("🔍 extractAuthData called with webApp:", webApp);
  const initData = webApp.initData;
  const initDataUnsafe = webApp.initDataUnsafe;
  const hash = initData ? new URLSearchParams(initData).get("hash") : null;
  const startParam = initDataUnsafe?.start_param;

  console.log("🔍 Telegram WebApp initDataUnsafe:", initDataUnsafe);
  console.log("🔗 Start param from Telegram:", startParam);

  return {
    hash,
    initData,
    initDataUnsafe,
    user: initDataUnsafe?.user,
    authDate: initDataUnsafe?.auth_date,
    queryId: initDataUnsafe?.query_id,
    themeParams: initDataUnsafe?.theme_params,
    startParam: startParam,
  };
}

export function getTelegramUser() {
  return isValidTelegramWebApp()
    ? window.Telegram.WebApp.initDataUnsafe?.user
    : null;
}

export function logAuthData(authData) {
  if (!authData) return;

  console.log("=== Telegram WebApp Data for Server Verification ===");
  console.log("Init Data (raw):", authData.initData);
  console.log("Init Data Unsafe:", authData.initDataUnsafe);

  if (authData.hash) {
    console.log("📤 Send this hash to your server for validation");
  }

  if (authData.user) {
    console.log("🆔 User ID:", authData.user.id);
    console.log("👤 Username:", authData.user.username);
    console.log("📛 First Name:", authData.user.first_name);
    console.log("📛 Last Name:", authData.user.last_name);
  }
}
