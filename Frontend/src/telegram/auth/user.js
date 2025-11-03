/**
 * Функции для работы с пользовательскими данными из JWT токена
 */

/**
 * Парсит JWT токен и возвращает полезную нагрузку (payload)
 * @param {string} token - JWT токен
 * @returns {Object|null} Объект с данными пользователя или null
 */
export function parseJWTToken(token) {
  console.log("🔍 Парсинг JWT токена");
  if (!token) {
    console.log("⚠️ Токен отсутствует");
    return null;
  }

  try {
    const base64Url = token.split(".")[1];
    console.log("🔍 base64Url части токена:", base64Url);
    const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
    console.log("🔍 base64 после замены символов:", base64);
    const jsonPayload = decodeURIComponent(
      atob(base64)
        .split("")
        .map((c) => "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2))
        .join("")
    );
    console.log("🔍 Раскодированный payload:", jsonPayload);

    const parsedPayload = JSON.parse(jsonPayload);
    console.log("✅ Токен успешно распарсен:", parsedPayload);
    return parsedPayload;
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
  console.log(
    "🔍 Получение информации о пользователе из токена:",
    token ? "Токен найден" : "Токен отсутствует"
  );
  if (!token) return null;

  const payload = parseJWTToken(token);
  console.log("🔍 Распарсенный токен:", payload);
  if (!payload) return null;

  // Проверяем, не истек ли токен
  const currentTime = Math.floor(Date.now() / 1000);
  if (payload.exp <= currentTime) {
    console.log("⚠️ Токен истек");
    return null;
  }

  const userInfo = {
    id: payload.user_id,
    telegramId: payload.telegram_id,
    generatedName: payload.generated_name,
    isAdmin: payload.is_admin || false,
    role: payload.role !== undefined ? payload.role : null,
    issuedAt: payload.iat,
    expiresAt: payload.exp,
    issuer: payload.iss,
  };

  console.log("✅ Информация о пользователе из токена:", userInfo);
  return userInfo;
}
