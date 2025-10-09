/**
 * Функции для взаимодействия с сервером
 */

export async function sendAuthToServer(
  authData,
  endpoint = "/api/auth/telegram"
) {
  if (!authData?.hash) {
    console.warn("No auth hash available");
    return null;
  }

  try {
    // Подготавливаем данные для отправки в формате, который ожидает сервер
    const payload = {
      id: authData.user?.id,
      first_name: authData.user?.first_name,
      last_name: authData.user?.last_name,
      username: authData.user?.username,
      photo_url: authData.user?.photo_url,
      auth_date: authData.authDate,
      hash: authData.hash,
    };

    console.log("📤 Sending auth data to server:", payload);

    const response = await fetch(endpoint, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(payload),
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`Server responded with ${response.status}: ${errorText}`);
    }

    const result = await response.json();
    console.log("✅ Auth data successfully sent to server:", result);

    // Если сервер вернул токен, сохраняем его
    if (result.token) {
      console.log("🔐 JWT token received from server");
    }

    return result;
  } catch (error) {
    console.error("❌ Failed to send auth data to server:", error);
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
