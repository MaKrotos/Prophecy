import { ref } from 'vue';
import { applyThemeColors, isDarkTheme as checkDarkTheme } from '../index.js';

/**
 * Композиция для управления темой Telegram
 */
export function useTelegramTheme() {
  const themeParams = ref(null);
  const isDarkTheme = ref(false);

  /**
   * Применение темы к приложению
   */
  const applyThemeToApp = (theme) => {
    if (!theme) return;

    applyThemeColors(theme);
    themeParams.value = theme;
    isDarkTheme.value = checkDarkTheme(theme);

    // Обновление классов темы
    if (isDarkTheme.value) {
      document.body.classList.add('tg-theme-dark');
      document.body.classList.remove('tg-theme-light');
    } else {
      document.body.classList.add('tg-theme-light');
      document.body.classList.remove('tg-theme-dark');
    }

    console.log('🎨 Theme applied to app');
  };

  /**
   * Обновление темы
   */
  const refreshTheme = (getTelegramAuthData) => {
    const tgData = getTelegramAuthData();
    if (tgData?.themeParams) {
      applyThemeToApp(tgData.themeParams);
    }
  };

  return {
    themeParams,
    isDarkTheme,
    applyThemeToApp,
    refreshTheme
  };
}