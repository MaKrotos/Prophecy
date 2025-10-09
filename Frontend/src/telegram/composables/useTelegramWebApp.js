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

  onMounted(() => {
    initTelegramWebApp({
      onUserDetected: (user) => {
        telegramUser.value = user;
        console.log("Telegram user detected:", user);
      },
      onReady: () => {
        isTelegram.value = true;
        console.log("Telegram WebApp ready");

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
        console.error("Telegram WebApp error:", error);
      },
      onHashReceived: (hash, initData) => {
        authHash.value = hash;
        authData.value = initData;

        console.log("✅ Authentication hash received and stored");
        console.log("🔐 Hash to send to server:", hash);
      },
      onThemeChanged: (theme) => {
        themeParams.value = theme;
        applyThemeToApp(theme);
        console.log("🎨 Theme changed:", theme);
      },
    });
  });

  const applyThemeToApp = (theme) => {
    if (!theme) return;

    applyThemeColors(theme);
    isDarkTheme.value = checkDarkTheme(theme);

    // Добавляем класс для темы
    if (isDarkTheme.value) {
      document.body.classList.add("tg-theme-dark");
      document.body.classList.remove("tg-theme-light");
    } else {
      document.body.classList.add("tg-theme-light");
      document.body.classList.remove("tg-theme-dark");
    }

    console.log("🎨 Theme applied to app:", {
      theme,
      isDark: isDarkTheme.value,
    });
  };

  const sendAuthToServer = async (endpoint = "/auth/telegram") => {
    const currentAuthData = getTelegramAuthData();
    if (!currentAuthData?.hash) {
      console.warn("No auth hash available");
      return null;
    }

    const payload = prepareAuthPayload({
      ...currentAuthData,
      themeParams: themeParams.value,
    });

    try {
      const result = await sendAuthToServerUtil(payload, endpoint);

      // Если сервер вернул токен, сохраняем его
      if (result && result.token) {
        saveJWTToken(result.token);
        jwtToken.value = result.token;
      }

      return result;
    } catch (error) {
      // При ошибке очищаем токен
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
    sendAuthToServer,
    getAuthData: getCurrentAuthData,
    refreshTheme,
    applyThemeToApp,
    logout,
    hasValidToken,
  };
}
