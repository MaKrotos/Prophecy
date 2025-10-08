/**
 * Утилиты для работы с цветами тем
 */

export function hexToRgb(hex) {
  const result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex);
  return result ? {
    r: parseInt(result[1], 16),
    g: parseInt(result[2], 16),
    b: parseInt(result[3], 16)
  } : null;
}

export function calculateBrightness(hexColor) {
  const rgb = hexToRgb(hexColor);
  if (!rgb) return 255;
  
  return (rgb.r * 299 + rgb.g * 587 + rgb.b * 114) / 1000;
}

export function applyThemeClassToBody(themeParams) {
  if (!themeParams?.bg_color) return;
  
  const brightness = calculateBrightness(themeParams.bg_color);
  const isDark = brightness < 128;
  
  if (isDark) {
    document.body.classList.add('tg-theme-dark');
    document.body.classList.remove('tg-theme-light');
  } else {
    document.body.classList.add('tg-theme-light');
    document.body.classList.remove('tg-theme-dark');
  }
  
  return isDark;
}