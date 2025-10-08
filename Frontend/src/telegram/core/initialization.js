import { isInTelegramWebView } from './detection.js';
import { isValidTelegramWebApp } from './validation.js'
import { applyThemeColors, setupThemeListeners } from '../theme/manager.js';
import { extractAuthData, logAuthData } from '../auth/data.js';
import { showTelegramOnlyMessage } from '../ui/messages.js';

/**
 * Основная функция инициализации Telegram WebApp
 */

export function initTelegramWebApp(callbacks = {}) {
  const {
    onUserDetected = () => {},
    onReady = () => {},
    onError = () => {},
    onHashReceived = () => {},
    onThemeChanged = () => {},
    onNotInTelegram = () => {},
    onViewportChanged = () => {},
  } = callbacks;

  if (typeof window === "undefined") {
    onError(new Error("Window object not available"));
    return;
  }

  // Проверяем, запущено ли в Telegram
  if (!isInTelegramWebView()) {
    console.warn("⚠️ App is not running in Telegram WebView");
    onNotInTelegram();
    return;
  }

  // Если скрипт уже загружен
  if (window.Telegram?.WebApp) {
    initializeTelegramFeatures(callbacks);
    return;
  }

  loadTelegramSDK(callbacks);
}

function loadTelegramSDK(callbacks) {
  const script = document.createElement("script");
  script.src = "https://telegram.org/js/telegram-web-app.js";
  script.async = true;

  script.onload = () => {
    console.log("Telegram WebApp SDK loaded");
    initializeTelegramFeatures(callbacks);
  };

  script.onerror = () => {
    console.error("Failed to load Telegram WebApp SDK");
    callbacks.onError(new Error("Failed to load Telegram WebApp SDK"));
  };

  document.head.appendChild(script);
}

function initializeTelegramFeatures(callbacks) {
  try {
    if (!window.Telegram?.WebApp) {
      throw new Error("Telegram WebApp API not available");
    }

    const webApp = window.Telegram.WebApp;

    // Дополнительная проверка через initData
    if (!isValidTelegramWebApp(webApp)) {
      console.warn("⚠️ Telegram WebApp data is invalid or missing");
      callbacks.onNotInTelegram();
      showTelegramOnlyMessage();
      return;
    }

    // Базовые настройки WebApp
    setupWebApp(webApp);
    
    // Настройка темы
    setupThemeListeners(webApp, callbacks.onThemeChanged);
    
    // Обработка auth данных
    processAuthData(webApp, callbacks);
    
    // Логирование информации для отладки
    logDebugInfo(webApp);

    callbacks.onReady();
  } catch (error) {
    console.error("Error initializing Telegram WebApp:", error);
    callbacks.onError(error);
  }
}

function setupWebApp(webApp) {
  webApp.ready?.();
  webApp.expand();
  webApp.enableClosingConfirmation();
}

function processAuthData(webApp, callbacks) {
  const authData = extractAuthData(webApp);
  
  if (authData.hash) {
    console.log("🔐 Authentication HASH for server verification:", authData.hash);
    callbacks.onHashReceived(authData.hash, authData.initData);
  }

  if (authData.user) {
    console.log("👤 Telegram User Data:", authData.user);
    callbacks.onUserDetected(authData.user);
  } else {
    console.log("ℹ️ No Telegram user data found");
  }

  logAuthData(authData);
}

function logDebugInfo(webApp) {
  console.log("🌐 Telegram WebApp version:", webApp.version);
  console.log("📱 Platform:", webApp.platform);
  console.log("🎨 Theme:", webApp.colorScheme);
  console.log("🎨 Theme Params:", webApp.themeParams);
}