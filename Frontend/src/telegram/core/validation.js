import { isInTelegramWebView } from './detection.js';

/**
 * Функции валидации Telegram WebApp данных
 */

export function isValidTelegramWebApp(webApp = window.Telegram?.WebApp) {
  if (!webApp) return false;
  
  const hasInitData = !!webApp.initData;
  const hasPlatform = !!webApp.platform;
  const hasVersion = !!webApp.version;
  
  const initDataUnsafe = webApp.initDataUnsafe;
  const hasValidInitData = 
    hasInitData && 
    initDataUnsafe && 
    (initDataUnsafe.user || initDataUnsafe.query_id || initDataUnsafe.auth_date);
  
  console.log("🔍 Telegram WebApp validation:", {
    hasInitData,
    hasPlatform,
    hasVersion,
    hasValidInitData,
    initDataUnsafe
  });
  
  return hasInitData && hasPlatform && hasVersion && hasValidInitData;
}

export function validateTelegramInitData(initDataUnsafe) {
  if (!initDataUnsafe) return false;
  
  return {
    hasUser: !!initDataUnsafe.user,
    hasQueryId: !!initDataUnsafe.query_id,
    hasAuthDate: !!initDataUnsafe.auth_date,
    isValid: !!(initDataUnsafe.user || initDataUnsafe.query_id || initDataUnsafe.auth_date)
  };
}