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
