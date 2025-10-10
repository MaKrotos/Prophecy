import { ref, onMounted } from 'vue';
import { initTelegramWebApp, getTelegramAuthData } from '../index.js';
import { hasValidToken, saveJWTToken, clearJWTToken } from '../auth/jwt.js';
import { useTelegramAuth } from './useTelegramAuth.js';
import { useTelegramTheme } from './useTelegramTheme.js';
import { waitForCondition } from '../utils/retry.js';

/**
 * Главная композиция для работы с Telegram WebApp
 */
export function useTelegramWebApp() {
  // Reactive state
  const telegramUser = ref(null);
  const isTelegram = ref(false);
  const isTelegramReady = ref(false);
  const initializationError = ref(null);

  // Инициализация модулей
  const { themeParams, isDarkTheme, applyThemeToApp, refreshTheme } = useTelegramTheme();
  const { 
    authError, 
    isAuthenticating, 
    needsReauth, 
    checkTelegramIdConsistency, 
    performAuth, 
    retryAuth, 
    clearAuthError 
  } = useTelegramAuth(
    telegramUser, 
    getTelegramAuthData, 
    hasValidToken, 
    saveJWTToken, 
    clearJWTToken
  );

  // Auth data
  const authHash = ref(null);
  const authData = ref(null);
  const jwtToken = ref(hasValidToken() ? localStorage.getItem('jwt_token') : null);

  onMounted(() => {
    initTelegramWebAppWithRetry();
  });

  /**
   * Инициализация Telegram WebApp с повторными попытками
   */
  const initTelegramWebAppWithRetry = async (retryCount = 0, maxRetries = 15) => {
    try {
      console.log(`🔄 Попытка инициализации Telegram WebApp #${retryCount + 1}`);

      initTelegramWebApp({
        onUserDetected: (user) => {
          telegramUser.value = user;
          console.log('✅ Telegram user detected:', user);
        },
        onReady: () => {
          handleTelegramReady();
        },
        onError: (error) => {
          handleTelegramError(error, retryCount, maxRetries);
        },
        onHashReceived: (hash, initData) => {
          authHash.value = hash;
          authData.value = initData;
          console.log('✅ Authentication hash received and stored');
        },
        onThemeChanged: (theme) => {
          themeParams.value = theme;
          applyThemeToApp(theme);
          console.log('🎨 Theme changed:', theme);
        },
      });

      // Проверяем готовность
      checkReadiness(retryCount, maxRetries);
    } catch (error) {
      handleTelegramError(error, retryCount, maxRetries);
    }
  };

  /**
   * Обработка готовности Telegram
   */
  const handleTelegramReady = () => {
    isTelegram.value = true;
    isTelegramReady.value = true;
    initializationError.value = null;
    console.log('✅ Telegram WebApp READY');

    // Проверяем наличие валидного токена
    if (hasValidToken()) {
      jwtToken.value = localStorage.getItem('jwt_token');
      console.log('✅ Найден валидный JWT токен');
    }

    // Применяем тему
    const tgData = getTelegramAuthData();
    if (tgData?.themeParams) {
      applyThemeToApp(tgData.themeParams);
    }
  };

  /**
   * Обработка ошибок Telegram
   */
  const handleTelegramError = (error, retryCount, maxRetries) => {
    console.error('❌ Telegram WebApp error:', error);
    initializationError.value = error;

    if (retryCount < maxRetries) {
      setTimeout(() => {
        initTelegramWebAppWithRetry(retryCount + 1, maxRetries);
      }, 200);
    }
  };

  /**
   * Проверка готовности Telegram
   */
  const checkReadiness = (retryCount, maxRetries) => {
    setTimeout(() => {
      if (!isTelegramReady.value && retryCount < maxRetries) {
        console.log(`⏳ Telegram не готов, повторная попытка #${retryCount + 1}`);
        initTelegramWebAppWithRetry(retryCount + 1, maxRetries);
      } else if (!isTelegramReady.value) {
        console.warn(`⚠️ Достигнут лимит попыток (${maxRetries}), продолжаем без Telegram`);
      }
    }, 300);
  };

  /**
   * Ожидание готовности Telegram
   */
  const waitForTelegramReady = (timeout = 5000) => {
    return waitForCondition(() => isTelegramReady.value, timeout);
  };

  /**
   * Выход из системы
   */
  const logout = () => {
    clearJWTToken();
    jwtToken.value = null;
    console.log('📤 Пользователь вышел из системы');
  };

  return {
    // State
    telegramUser,
    isTelegram,
    authHash,
    authData,
    themeParams,
    isDarkTheme,
    jwtToken,
    isTelegramReady,
    initializationError,
    
    // Auth
    authError,
    isAuthenticating,
    needsReauth,
    hasValidToken,
    
    // Methods
    sendAuthToServer: performAuth,
    getAuthData: getTelegramAuthData,
    refreshTheme: () => refreshTheme(getTelegramAuthData),
    applyThemeToApp,
    logout,
    waitForTelegramReady,
    saveJWTToken,
    clearJWTToken,
    checkTelegramIdConsistency,
    retryAuth,
    clearAuthError
  };
}