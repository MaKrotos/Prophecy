import { useApi } from '../composables/useApi.js';

/**
 * Функции для взаимодействия с сервером
 */

export async function sendAuthToServer(
  authData,
  endpoint = "auth/telegram"
) {
  if (!authData?.hash) {
    console.warn("No auth hash available");
    throw new Error("No authentication data available");
  }

  const { apiPost } = useApi();
  try {
    console.log(`📤 Sending auth data to server:`, authData);

    const response = await apiPost(endpoint, {
      initData: window.Telegram?.WebApp?.initData,
    }, {
      // Отключаем автоматическую аутентификацию для этого запроса,
      // так как мы отправляем данные для получения токена
      autoAuth: false
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(
        `Server responded with ${response.status}: ${errorText}`
      );
    }

    const result = await response.json();
    console.log("✅ Auth data successfully sent to server:", result);

    return result;
  } catch (error) {
    console.error(`❌ Failed to send auth data to server:`, error.message);
    throw error;
  }
}

export function prepareAuthPayload(authData) {
  return {
    hash: authData.hash,
    initData: authData.initData,
    user: authData.user,
    authDate: authData.authDate,
    queryId: authData.queryId,
    theme: authData.themeParams,
    timestamp: Date.now(),
  };
}
