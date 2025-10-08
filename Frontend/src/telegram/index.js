export * from './core/detection.js';
export * from './core/initialization.js';
export * from './core/validation.js';
export * from './theme/manager.js';
export * from './auth/data.js';
export * from './auth/server.js';
export * from './ui/messages.js';

// Основные функции для обратной совместимости
export { initTelegramWebApp } from './core/initialization.js';
export { getTelegramAuthData } from './auth/data.js';
export { getTelegramUser } from './auth/data.js';
export { applyTelegramThemeToDocument } from './theme/manager.js';
export { showTelegramOnlyMessage } from './ui/messages.js';