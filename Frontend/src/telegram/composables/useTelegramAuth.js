import { ref } from 'vue';
import { AuthService } from '../auth/authService.js';
import { getUserInfoFromToken } from '../auth/user.js';

/**
 * Композиция для управления аутентификацией Telegram
 */
export function useTelegramAuth(telegramUser, getTelegramAuthData, hasValidToken, saveJWTToken, clearJWTToken) {
  const authError = ref(null);
  const isAuthenticating = ref(false);
  const needsReauth = ref(false);

  /**
   * Проверка соответствия Telegram ID
   */
  const checkTelegramIdConsistency = () => {
    const result = AuthService.checkTelegramIdConsistency(telegramUser.value, getUserInfoFromToken);
    needsReauth.value = result.needsReauth;
    return result.consistent;
  };

  /**
   * Выполнение аутентификации
   */
  const performAuth = async (endpoint = '/api/auth/telegram', maxRetries = 3) => {
    isAuthenticating.value = true;
    authError.value = null;

    try {
      const authData = getTelegramAuthData();
      
      if (!authData?.hash) {
        throw new Error('Нет хэша аутентификации. Перезайдите в приложение.');
      }

      const result = await AuthService.sendAuthToServer(authData, endpoint, maxRetries);
      needsReauth.value = false;
      return result;
    } catch (error) {
      authError.value = error.message;
      throw error;
    } finally {
      isAuthenticating.value = false;
    }
  };

  /**
   * Повторная аутентификация
   */
  const retryAuth = async (endpoint = '/api/auth/telegram', maxRetries = 3) => {
    authError.value = null;
    
    try {
      const result = await performAuth(endpoint, maxRetries);
      console.log('✅ Повторная авторизация успешна');
      return result;
    } catch (error) {
      console.error('❌ Повторная авторизация не удалась после всех попыток:', error);
      throw error;
    }
  };

  const clearAuthError = () => {
    authError.value = null;
  };

  return {
    authError,
    isAuthenticating,
    needsReauth,
    checkTelegramIdConsistency,
    performAuth,
    retryAuth,
    clearAuthError
  };
}