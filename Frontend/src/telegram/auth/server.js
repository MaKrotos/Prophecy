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

  // Для запроса аутентификации используем прямой fetch,
  // чтобы избежать циклической зависимости
  const apiUrl = `/api/${endpoint.replace(/^\//, '')}`;
  try {
    console.log(`📤 Sending auth data to server:`, authData);

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
