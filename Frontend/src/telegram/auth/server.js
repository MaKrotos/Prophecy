/**
 * Функции для взаимодействия с сервером
 */

export async function sendAuthToServer(authData, endpoint = "/api/telegram-auth") {
  if (!authData?.hash) {
    console.warn("No auth hash available");
    return null;
  }

  try {
    const response = await fetch(endpoint, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(authData),
    });

    if (!response.ok) {
      throw new Error(`Server responded with ${response.status}`);
    }

    const result = await response.json();
    console.log("✅ Auth data successfully sent to server:", result);
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