/**
 * Функции для работы с пользовательскими данными из JWT токена
 */

/**
 * Парсит JWT токен и возвращает полезную нагрузку (payload)
 * @param {string} token - JWT токен
 * @returns {Object|null} Объект с данными пользователя или null
 */
export function parseJWTToken(token) {
  if (!token) return null;

  try {
    const base64Url = token.split(".")[1];
    const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
    const jsonPayload = decodeURIComponent(
      atob(base64)
        .split("")
        .map((c) => "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2))
        .join("")
    );

    return JSON.parse(jsonPayload);
  } catch (e) {
    console.error("Ошибка парсинга токена:", e);
    return null;
  }
}

/**
 * Получает информацию о пользователе из JWT токена
 * @returns {Object|null} Объект с информацией о пользователе или null
 */
export function getUserInfoFromToken() {
  const token = localStorage.getItem("jwt_token");
  if (!token) return null;

  const payload = parseJWTToken(token);
  if (!payload) return null;

  // Проверяем, не истек ли токен
  const currentTime = Math.floor(Date.now() / 1000);
  if (payload.exp <= currentTime) {
    return null;
  }

  return {
    id: payload.user_id,
    telegramId: payload.telegram_id,
    generatedName: payload.generated_name,
    isAdmin: payload.is_admin || false,
    issuedAt: payload.iat,
    expiresAt: payload.exp,
    issuer: payload.iss,
  };
}
