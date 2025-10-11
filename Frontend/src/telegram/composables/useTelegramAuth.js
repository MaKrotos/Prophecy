import { ref } from 'vue';
import { useApi } from './useApi.js';
import { retryWithBackoff } from '../utils/retry.js';
import { saveJWTToken, clearJWTToken } from '../auth/jwt.js';
import { getUserInfoFromToken } from '../auth/user.js';
import { prepareAuthPayload } from '../auth/server.js';

/**
 * Композиция для управления аутентификацией Telegram
 */
export function useTelegramAuth(telegramUser, getTelegramAuthData, hasValidToken) {
  const authError = ref(null);
  const isAuthenticating = ref(false);
  const needsReauth = ref(false);

  /**
   * Проверка соответствия Telegram ID
   */
  const checkTelegramIdConsistency = () => {
    const tokenUserInfo = getUserInfoFromToken();

    // Если токена нет, проверка не требуется
    if (!tokenUserInfo) {
      console.log('ℹ️ Нет JWT токена для проверки');
      needsReauth.value = false;
      return { consistent: true, needsReauth: false };
    }

    const tokenTelegramId = tokenUserInfo.telegramId;
    const webAppTelegramId = telegramUser?.value?.id;

    // Если нет данных WebApp, проверка не может быть выполнена
    if (!webAppTelegramId) {
      console.log('⚠️ Нет данных Telegram WebApp для проверки');
      needsReauth.value = false;
      return { consistent: true, needsReauth: false };
    }

    if (tokenTelegramId === webAppTelegramId) {
      console.log('✅ Telegram ID из токена и WebApp совпадают');
      needsReauth.value = false;
      return { consistent: true, needsReauth: false };
    } else {
      console.warn('❌ Telegram ID из токена и WebApp НЕ совпадают', {
        tokenTelegramId,
        webAppTelegramId,
      });
      
      clearJWTToken();
      needsReauth.value = true;
      return { consistent: false, needsReauth: true };
    }
  };

  /**
   * Выполнение аутентификации
   */
  const performAuth = async (endpoint = 'auth/telegram', maxRetries = 3) => {
    isAuthenticating.value = true;
    authError.value = null;

    try {
      const authData = getTelegramAuthData();
      
      if (!authData?.hash) {
        throw new Error('Нет хэша аутентификации. Перезайдите в приложение.');
      }

      const payload = prepareAuthPayload({
        ...authData,
        themeParams: authData.themeParams
      });

      const result = await retryWithBackoff(
        async () => {
          // Для запроса аутентификации используем прямой fetch,
          // чтобы избежать циклической зависимости
          const apiUrl = `/api/${endpoint.replace(/^\//, '')}`;
          
          const response = await fetch(apiUrl, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({
              initData: window.Telegram?.WebApp?.initData,
            })
          });
          
          if (!response.ok) {
            const errorText = await response.text();
            throw new Error(
              `Server responded with ${response.status}: ${errorText}`
            );
          }
          
          const data = await response.json();
          
          // Если сервер вернул токен, сохраняем его
          if (data.token) {
            saveJWTToken(data.token);
            console.log("🔐 JWT token received from server");
            return data;
          } else {
            throw new Error('Сервер не вернул токен');
          }
        },
        maxRetries,
        {
          onRetry: (attempt, maxAttempts, delay, error) => {
            console.warn(`❌ Попытка авторизации #${attempt}/${maxAttempts} не удалась:`, error.message);
          }
        }
      );
      
      console.log('✅ Авторизация успешна');
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
  const retryAuth = async (endpoint = 'auth/telegram', maxRetries = 3) => {
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