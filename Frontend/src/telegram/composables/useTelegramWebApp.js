import { ref, onMounted } from "vue";
import {
  initTelegramWebApp,
  getTelegramAuthData,
  applyThemeColors,
  isDarkTheme as checkDarkTheme,
  sendAuthToServer as sendAuthToServerUtil,
  prepareAuthPayload
} from "../index.js";

export function useTelegramWebApp() {
  const telegramUser = ref(null);
  const isTelegram = ref(false);
  const authHash = ref(null);
  const authData = ref(null);
  const themeParams = ref(null);
  const isDarkTheme = ref(false);

  onMounted(() => {
    initTelegramWebApp({
      onUserDetected: (user) => {
        telegramUser.value = user;
        console.log("Telegram user detected:", user);
      },
      onReady: () => {
        isTelegram.value = true;
        console.log("Telegram WebApp ready");
        
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
      document.body.classList.add('tg-theme-dark');
      document.body.classList.remove('tg-theme-light');
    } else {
      document.body.classList.add('tg-theme-light');
      document.body.classList.remove('tg-theme-dark');
    }

    console.log("🎨 Theme applied to app:", {
      theme,
      isDark: isDarkTheme.value
    });
  };

  const sendAuthToServer = async (endpoint = "/api/telegram-auth") => {
    const currentAuthData = getTelegramAuthData();
    if (!currentAuthData?.hash) {
      console.warn("No auth hash available");
      return null;
    }

    const payload = prepareAuthPayload({
      ...currentAuthData,
      themeParams: themeParams.value
    });

    return await sendAuthToServerUtil(payload, endpoint);
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

  return {
    telegramUser,
    isTelegram,
    authHash,
    authData,
    themeParams,
    isDarkTheme,
    sendAuthToServer,
    getAuthData: getCurrentAuthData,
    refreshTheme,
    applyThemeToApp,
  };
}