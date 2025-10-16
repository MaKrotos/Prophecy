import { ref, onMounted } from "vue";
import { initTelegramWebApp, getTelegramAuthData } from "../index.js";
import { hasValidToken, saveJWTToken, clearJWTToken } from "../auth/jwt.js";
import { useTelegramAuth } from "./useTelegramAuth.js";
import { useTelegramTheme } from "./useTelegramTheme.js";
import { waitForCondition } from "../utils/retry.js";
import { useRouter } from "vue-router";

/**
 * Главная композиция для работы с Telegram WebApp
 */
export function useTelegramWebApp() {
  // Router
  const router = useRouter();

  // Reactive state
  const telegramUser = ref(null);
  const isTelegram = ref(false);
  const isTelegramReady = ref(false);
  const initializationError = ref(null);

  // Инициализация модулей
  const { themeParams, isDarkTheme, applyThemeToApp, refreshTheme } =
    useTelegramTheme();
  const {
    authError,
    isAuthenticating,
    needsReauth,
    checkTelegramIdConsistency,
    performAuth,
    retryAuth,
    clearAuthError,
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
  const startParam = ref(null);
  const jwtToken = ref(
    hasValidToken() ? localStorage.getItem("jwt_token") : null
  );

  onMounted(() => {
    initTelegramWebAppWithRetry();
  });

  /**
   * Инициализация Telegram WebApp с повторными попытками
   */
  const initTelegramWebAppWithRetry = async (
    retryCount = 0,
    maxRetries = 15
  ) => {
    try {
      console.log(
        `🔄 Попытка инициализации Telegram WebApp #${retryCount + 1}`
      );

      initTelegramWebApp({
        onUserDetected: (user) => {
          telegramUser.value = user;
          console.log("✅ Telegram user detected:", user);
        },
        onReady: () => {
          handleTelegramReady();
        },
        onError: (error) => {
          handleTelegramError(error, retryCount, maxRetries);
        },
        onHashReceived: (hash, initData, startParamValue) => {
          console.log("🔍 onHashReceived called with:", {
            hash,
            initData,
            startParamValue,
          });
          authHash.value = hash;
          authData.value = initData;
          startParam.value = startParamValue;
          console.log("✅ Authentication hash received and stored");
          console.log(
            "🔗 Start param received in onHashReceived:",
            startParamValue
          );
          // Дополнительная проверка значения startParam
          if (startParamValue) {
            console.log(
              "🔍 Значение startParam в onHashReceived:",
              startParamValue
            );
          } else {
            console.log("⚠️ Значение startParam пустое в onHashReceived");
          }
        },
        onThemeChanged: (theme) => {
          themeParams.value = theme;
          applyThemeToApp(theme);
          console.log("🎨 Theme changed:", theme);
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
    console.log("✅ Telegram WebApp READY");

    // Проверяем наличие валидного токена
    if (hasValidToken()) {
      jwtToken.value = localStorage.getItem("jwt_token");
      console.log("✅ Найден валидный JWT токен");
    }

    // Применяем тему
    const tgData = getTelegramAuthData();
    if (tgData?.themeParams) {
      applyThemeToApp(tgData.themeParams);
    }

    console.log("🔍 Telegram WebApp ready, startParam:", startParam.value);

    // Если есть startParam, обрабатываем его
    if (startParam.value) {
      console.log("🔗 Найден параметр startParam:", startParam.value);
      console.log("🔍 Начинаем обработку startParam при инициализации");
      // Парсим и обрабатываем параметр
      parseAndHandleStartParam(startParam.value);
      console.log("✅ Обработка startParam при инициализации завершена");
      console.log("🔍 Значение startParam после обработки:", startParam.value);
    } else {
      console.log("⚠️ Параметр startParam отсутствует");
    }
  };

  /**
   * Парсинг и обработка параметра startParam
   */
  const parseAndHandleStartParam = (param) => {
    console.log("🔍 Начало обработки параметра startParam:", param);
    if (!param) {
      console.log("⚠️ Параметр startParam пустой или отсутствует");
      return;
    }

    // Разбиваем параметр на части по "_"
    const parts = param.split("_");
    const action = parts[0];
    const args = parts.slice(1);

    console.log("🔧 Обработка параметра startParam:", { action, args });

    // Обработчики для разных действий
    const handlers = {
      join_session: (sessionId) => {
        console.log(
          "🔍 Обработка действия join_session с sessionId:",
          sessionId
        );
        if (sessionId) {
          console.log(
            "🔗 Перенаправление на страницу присоединения к сессии:",
            sessionId
          );
          // Небольшая задержка для уверенности в инициализации
          setTimeout(() => {
            router.push(`/sessions/join/${sessionId}`);
          }, 100);
        } else {
          console.log("⚠️ Не указан sessionId для действия join_session");
        }
        // Обнуляем startParam, чтобы он не срабатывал при повторной инициализации
        console.log("🧹 Обнуление startParam после обработки join_session");
        startParam.value = null;
      },
      // Примеры других обработчиков
      invite_user: (userId) => {
        if (userId) {
          console.log("👥 Приглашение пользователя:", userId);
          // Здесь можно добавить логику приглашения пользователя
          // Например, показать модальное окно с подтверждением приглашения
        }
        // Обнуляем startParam, чтобы он не срабатывал при повторной инициализации
        console.log("🧹 Обнуление startParam после обработки invite_user");
        startParam.value = null;
      },
      share_content: (contentId) => {
        if (contentId) {
          console.log("📤 Общий доступ к контенту:", contentId);
          // Здесь можно добавить логику общего доступа к контенту
          // Например, перенаправить на страницу контента и показать опции общего доступа
        }
        // Обнуляем startParam, чтобы он не срабатывал при повторной инициализации
        console.log("🧹 Обнуление startParam после обработки share_content");
        startParam.value = null;
      },
    };

    // Вызываем соответствующий обработчик
    if (handlers[action]) {
      console.log("✅ Вызов обработчика для действия:", action);
      handlers[action](...args);
      console.log("✅ Обработчик для действия завершен:", action);
    } else {
      console.warn("⚠️ Неизвестное действие в параметре startParam:", action);
      // Обнуляем startParam, чтобы он не срабатывал при повторной инициализации
      console.log("🧹 Обнуление startParam после неизвестного действия");
      startParam.value = null;
    }
  };

  /**
   * Обработка ошибок Telegram
   */
  const handleTelegramError = (error, retryCount, maxRetries) => {
    console.error("❌ Telegram WebApp error:", error);
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
        console.log(
          `⏳ Telegram не готов, повторная попытка #${retryCount + 1}`
        );
        initTelegramWebAppWithRetry(retryCount + 1, maxRetries);
      } else if (!isTelegramReady.value) {
        console.warn(
          `⚠️ Достигнут лимит попыток (${maxRetries}), продолжаем без Telegram`
        );
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
    console.log("📤 Пользователь вышел из системы");
  };

  return {
    // State
    telegramUser,
    isTelegram,
    authHash,
    authData,
    startParam,
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
    clearAuthError,
  };
}
