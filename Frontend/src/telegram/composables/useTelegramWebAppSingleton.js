import { useTelegramWebApp } from "./useTelegramWebApp.js";

// Глобальный инстанс композиции
let telegramWebAppInstance = null;

// Функция для получения или создания инстанса
function getTelegramWebAppInstance() {
  if (!telegramWebAppInstance) {
    telegramWebAppInstance = useTelegramWebApp();
  }
  return telegramWebAppInstance;
}

/**
 * Singleton-версия композиции для работы с Telegram WebApp
 * Гарантирует, что состояние сохраняется между переходами по страницам
 */
export function useTelegramWebAppSingleton() {
  return getTelegramWebAppInstance();
}

// Также экспортируем инстанс напрямую для случаев, когда нужно получить доступ к состоянию вне композиции
export { telegramWebAppInstance };