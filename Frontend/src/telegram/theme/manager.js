import { calculateBrightness, hexToRgb } from './colors.js';

/**
 * Управление темами Telegram WebApp
 */

export function applyThemeColors(themeParams) {
  if (!themeParams) return;

  console.log("🎨 Applying Telegram theme colors:", themeParams);

  const root = document.documentElement;
  
  // Основные цвета
  applyThemeColor(root, themeParams.bg_color, [
    '--tg-theme-bg-color',
    '--tg-background-color'
  ]);
  
  applyThemeColor(root, themeParams.text_color, [
    '--tg-theme-text-color',
    '--tg-text-color'
  ]);
  
  applyThemeColor(root, themeParams.hint_color, [
    '--tg-theme-hint-color',
    '--tg-hint-color'
  ]);
  
  applyThemeColor(root, themeParams.link_color, [
    '--tg-theme-link-color',
    '--tg-link-color'
  ]);
  
  applyThemeColor(root, themeParams.button_color, [
    '--tg-theme-button-color',
    '--tg-button-color'
  ]);
  
  applyThemeColor(root, themeParams.button_text_color, [
    '--tg-theme-button-text-color',
    '--tg-button-text-color'
  ]);
  
  applyThemeColor(root, themeParams.secondary_bg_color, [
    '--tg-theme-secondary-bg-color',
    '--tg-secondary-bg-color'
  ]);

  return themeParams;
}

function applyThemeColor(root, color, cssVars) {
  if (!color) return;
  
  cssVars.forEach(cssVar => {
    root.style.setProperty(cssVar, color);
  });
}

export function setupThemeListeners(webApp, onThemeChanged = () => {}) {
  if (!webApp) return;
  
  webApp.onEvent('themeChanged', (themeParams) => {
    applyThemeColors(themeParams);
    updateWebAppColors(webApp, themeParams);
    onThemeChanged(themeParams);
  });
  
  webApp.onEvent('viewportChanged', handleViewportChange);
}

export function updateWebAppColors(webApp, themeParams) {
  if (!webApp) return;
  
  const headerColor = themeParams.bg_color || '#2b2d42';
  webApp.setHeaderColor(headerColor);
  webApp.setBackgroundColor(themeParams.bg_color || '#ffffff');
}

function handleViewportChange() {
  const webApp = window.Telegram.WebApp;
  console.log("📱 Viewport changed:", {
    height: webApp.viewportHeight,
    stableHeight: webApp.viewportStableHeight,
    isExpanded: webApp.isExpanded
  });
}

export function applyTelegramThemeToDocument() {
  if (!window.Telegram?.WebApp) return;
  
  const webApp = window.Telegram.WebApp;
  const themeParams = webApp.themeParams;
  
  if (!themeParams) return;

  const root = document.documentElement;
  
  // Основные цвета темы
  if (themeParams.bg_color) {
    root.style.setProperty('--tg-theme-bg-color', themeParams.bg_color);
    document.body.style.backgroundColor = themeParams.bg_color;
  }
  
  if (themeParams.text_color) {
    root.style.setProperty('--tg-theme-text-color', themeParams.text_color);
    document.body.style.color = themeParams.text_color;
  }
  
  applyThemeColors(themeParams);

  return themeParams;
}

export function getTelegramThemeParams() {
  return window.Telegram?.WebApp?.themeParams || null;
}

export function setTelegramBackgroundColor(color) {
  if (!window.Telegram?.WebApp) return;
  window.Telegram.WebApp.setBackgroundColor(color);
}

export function setTelegramHeaderColor(color) {
  if (!window.Telegram?.WebApp) return;
  window.Telegram.WebApp.setHeaderColor(color);
}

export function isDarkTheme(themeParams) {
  if (!themeParams?.bg_color) return false;
  return calculateBrightness(themeParams.bg_color) < 128;
}