/**
 * Функции для обнаружения Telegram WebView окружения
 */

export function isInTelegramWebView() {
  if (typeof window === "undefined") return false;

  const userAgent = navigator.userAgent.toLowerCase();
  const isTelegramWebView =
    /telegram|webview/i.test(userAgent) ||
    /telegram|tweb/i.test(navigator.userAgent) ||
    !!window.TelegramWebviewProxy ||
    !!window.Telegram?.WebApp ||
    // Дополнительные проверки для Telegram WebApp
    (window.Telegram && typeof window.Telegram === "object") ||
    userAgent.includes("telegram") ||
    userAgent.includes("tweb") ||
    // Проверка для веб-версии Telegram
    userAgent.includes("electron") ||
    // Проверка на наличие специфичных для Telegram объектов
    (window.Telegram && window.Telegram.WebApp);

  console.log("🔍 Telegram WebView detection:", {
    userAgent: navigator.userAgent,
    userAgentLower: userAgent,
    hasTelegramWebviewProxy: !!window.TelegramWebviewProxy,
    hasTelegramWebApp: !!window.Telegram?.WebApp,
    isTelegramWebView,
  });

  return isTelegramWebView;
}

export function isTelegramEnvironment() {
  return isInTelegramWebView() && isValidTelegramWebApp();
}

export function requireTelegramWebApp() {
  if (!isInTelegramWebView() || !isValidTelegramWebApp()) {
    throw new Error("This application can only be used within Telegram");
  }
  return window.Telegram.WebApp;
}
