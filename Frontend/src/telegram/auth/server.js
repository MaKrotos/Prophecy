/**
 * Функции для взаимодействия с сервером
 */

export async function sendAuthToServer(
  authData,
  endpoint = "/api/auth/telegram",
  maxRetries = 3
) {
  if (!authData?.hash) {
    console.warn("No auth hash available");
    return null;
  }

  let lastError;

  // Пытаемся отправить запрос несколько раз
  for (let attempt = 0; attempt <= maxRetries; attempt++) {
    try {
      // Подготавливаем данные для отправки в формате, который ожидает сервер
      const payload = {
        initData: window.Telegram?.WebApp?.initData,
      };

      console.log(
        `📤 Sending auth data to server (attempt ${attempt + 1}/${
          maxRetries + 1
        }):`,
        payload
      );

      const response = await fetch(endpoint, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
      });

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(
          `Server responded with ${response.status}: ${errorText}`
        );
      }

      const result = await response.json();
      console.log("✅ Auth data successfully sent to server:", result);

      // Если сервер вернул токен, сохраняем его
      if (result.token) {
        console.log("🔐 JWT token received from server");
      }

      return result;
    } catch (error) {
      lastError = error;
      console.error(
        `❌ Failed to send auth data to server (attempt ${attempt + 1}):`,
        error.message
      );

      // Если это не последняя попытка, ждем перед повтором
      if (attempt < maxRetries) {
        // Экспоненциальная задержка: 1s, 2s, 4s, ...
        const delay = Math.pow(2, attempt) * 1000;
        console.log(`⏳ Waiting ${delay}ms before next attempt...`);
        await new Promise((resolve) => setTimeout(resolve, delay));
      }
    }
  }

  // Если все попытки исчерпаны, выбрасываем последнюю ошибку
  throw lastError;
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
