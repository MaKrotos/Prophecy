import { prepareAuthPayload, sendAuthToServer as sendAuthToServerUtil } from '../index.js';
import { saveJWTToken, clearJWTToken } from './jwt.js';
import { retryWithBackoff } from '../utils/retry.js';

/**
 * Сервис для управления аутентификацией Telegram
 */
export class AuthService {
  /**
   * Отправка данных аутентификации на сервер с повторными попытками
   */
  static async sendAuthToServer(authData, endpoint = '/api/auth/telegram', maxRetries = 3) {
    if (!authData?.hash) {
      throw new Error('No authentication data available');
    }

    const payload = prepareAuthPayload({
      ...authData,
      themeParams: authData.themeParams
    });

    const result = await retryWithBackoff(
      async () => {
        const response = await sendAuthToServerUtil(payload, endpoint);
        
        if (response?.token) {
          saveJWTToken(response.token);
          return response;
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
    return result;
  }

  /**
   * Проверка соответствия Telegram ID в JWT и WebApp данных
   */
  static checkTelegramIdConsistency(telegramUser, getUserInfoFromToken) {
    const tokenUserInfo = getUserInfoFromToken();

    // Если токена нет, проверка не требуется
    if (!tokenUserInfo) {
      console.log('ℹ️ Нет JWT токена для проверки');
      return { consistent: true, needsReauth: false };
    }

    const tokenTelegramId = tokenUserInfo.telegramId;
    const webAppTelegramId = telegramUser?.id;

    // Если нет данных WebApp, проверка не может быть выполнена
    if (!webAppTelegramId) {
      console.log('⚠️ Нет данных Telegram WebApp для проверки');
      return { consistent: true, needsReauth: false };
    }

    if (tokenTelegramId === webAppTelegramId) {
      console.log('✅ Telegram ID из токена и WebApp совпадают');
      return { consistent: true, needsReauth: false };
    } else {
      console.warn('❌ Telegram ID из токена и WebApp НЕ совпадают', {
        tokenTelegramId,
        webAppTelegramId,
      });
      
      clearJWTToken();
      return { consistent: false, needsReauth: true };
    }
  }
}