import { ref, onMounted } from "vue";
import {
  initTelegramWebApp,
  getTelegramAuthData,
  applyThemeColors,
  isDarkTheme as checkDarkTheme,
  sendAuthToServer as sendAuthToServerUtil,
  prepareAuthPayload,
} from "../index.js";
import { saveJWTToken, clearJWTToken, hasValidToken } from "../auth/jwt.js";

export function useTelegramWebApp() {
  const telegramUser = ref(null);
  const isTelegram = ref(false);
  const authHash = ref(null);
  const authData = ref(null);
  const themeParams = ref(null);
  const isDarkTheme = ref(false);
  const jwtToken = ref(null);
  const isTelegramReady = ref(false);
  const initializationError = ref(null);

  onMounted(() => {
    initTelegramWebAppWithRetry();
  });

  const initTelegramWebAppWithRetry = async (retryCount = 0, maxRetries = 15) => {
    try {
      console.log(`🔄 Попытка инициализации Telegram WebApp #${retryCount + 1}`);

      initTelegramWebApp({
        onUserDetected: (user) => {
          telegramUser.value = user;
          console.log("✅ Telegram user detected:", user);
        },
        onReady: () => {
          isTelegram.value = true;
          isTelegramReady.value = true;
          initializationError.value = null;
          console.log("✅ Telegram WebApp READY");

          // Проверяем наличие валидного токена
          if (hasValidToken()) {
            jwtToken.value = localStorage.getItem("jwt_token");
            console.log("✅ Найден валидный JWT токен");
          }

          // Получаем параметры темы после готовности
          const tgData = getTelegramAuthData();
          if (tgData?.themeParams) {
            themeParams.value = tgData.themeParams;
            applyThemeToApp(themeParams.value);
          }
        },
        onError: (error) => {
          console.error("❌ Telegram WebApp error:", error);
          initializationError.value = error;
          
          // При ошибке пробуем повторить
          if (retryCount < maxRetries) {
            setTimeout(() => {
              initTelegramWebAppWithRetry(retryCount + 1, maxRetries);
            }, 200);
          }
        },
        onHashReceived: (hash, initData) => {
          authHash.value = hash;
          authData.value = initData;
          console.log("✅ Authentication hash received and stored");
        },
        onThemeChanged: (theme) => {
          themeParams.value = theme;
          applyThemeToApp(theme);
          console.log("🎨 Theme changed:", theme);
        },
      });

      // Проверяем готовность через короткий интервал
      setTimeout(() => {
        if (!isTelegramReady.value && retryCount < maxRetries) {
          console.log(`⏳ Telegram не готов, повторная попытка #${retryCount + 1}`);
          initTelegramWebAppWithRetry(retryCount + 1, maxRetries);
        } else if (!isTelegramReady.value) {
          console.warn(`⚠️ Достигнут лимит попыток (${maxRetries}), продолжаем без Telegram`);
        }
      }, 300);

    } catch (error) {
      console.error("❌ Ошибка при инициализации Telegram WebApp:", error);
      initializationError.value = error;
      if (retryCount < maxRetries) {
        setTimeout(() => {
          initTelegramWebAppWithRetry(retryCount + 1, maxRetries);
        }, 200);
      }
    }
  };

  // Функция ожидания готовности Telegram
  const waitForTelegramReady = (timeout = 5000) => {
    return new Promise((resolve, reject) => {
      if (isTelegramReady.value) {
        resolve(true);
        return;
      }

      const startTime = Date.now();
      const checkInterval = setInterval(() => {
        if (isTelegramReady.value) {
          clearInterval(checkInterval);
          resolve(true);
        } else if (Date.now() - startTime > timeout) {
          clearInterval(checkInterval);
          console.warn(`⏰ Таймаут ожидания Telegram WebApp (${timeout}ms)`);
          reject(new Error(`Telegram WebApp initialization timeout (${timeout}ms)`));
        }
      }, 50);
    });
  };

  const applyThemeToApp = (theme) => {
    if (!theme) return;

    applyThemeColors(theme);
    isDarkTheme.value = checkDarkTheme(theme);

    if (isDarkTheme.value) {
      document.body.classList.add("tg-theme-dark");
      document.body.classList.remove("tg-theme-light");
    } else {
      document.body.classList.add("tg-theme-light");
      document.body.classList.remove("tg-theme-dark");
    }

    console.log("🎨 Theme applied to app");
  };

  const sendAuthToServer = async (endpoint = "/api/auth/telegram") => {
    // Ждем готовности Telegram перед отправкой
    try {
      await waitForTelegramReady(3000);
    } catch (error) {
      console.warn("⚠️ Telegram не готов, но продолжаем авторизацию:", error);
    }

    const currentAuthData = getTelegramAuthData();
    if (!currentAuthData?.hash) {
      console.warn("No auth hash available");
      throw new Error("No authentication data available");
    }

    const payload = prepareAuthPayload({
      ...currentAuthData,
      themeParams: themeParams.value,
    });

    try {
      const result = await sendAuthToServerUtil(payload, endpoint);

      if (result && result.token) {
        saveJWTToken(result.token);
        jwtToken.value = result.token;
      }

      return result;
    } catch (error) {
      clearJWTToken();
      jwtToken.value = null;
      throw error;
    }
  };

  const getCurrentAuthData = () => {
    return getTelegramAuthData();
  };

  const refreshTheme = () => {
    const tgData = getTelegramAuthData();
    if (tgData?.themeParams) {
      themeParams.value = tgData.themeParams;
      applyThemeToApp(themeParams.value);
    }
  };

  const logout = () => {
    clearJWTToken();
    jwtToken.value = null;
    console.log("📤 Пользователь вышел из системы");
  };

  return {
    telegramUser,
    isTelegram,
    authHash,
    authData,
    themeParams,
    isDarkTheme,
    jwtToken,
    isTelegramReady,
    initializationError,
    sendAuthToServer,
    getAuthData: getCurrentAuthData,
    refreshTheme,
    applyThemeToApp,
    logout,
    hasValidToken,
    waitForTelegramReady,
  };
}