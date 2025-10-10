/**
 * Функции для работы с JWT токенами
 */

/**
 * Сохраняет JWT токен в localStorage
 * @param {string} token - JWT токен
 */
export function saveJWTToken(token) {
  if (token) {
    localStorage.setItem("jwt_token", token);
    console.log("✅ JWT токен сохранен в localStorage");
  }
}

/**
 * Получает JWT токен из localStorage
 * @returns {string|null} JWT токен или null
 */
export function getJWTToken() {
  return localStorage.getItem("jwt_token");
}

/**
 * Удаляет JWT токен из localStorage
 */
export function clearJWTToken() {
  localStorage.removeItem("jwt_token");
  console.log("🗑️ JWT токен удален из localStorage");
}

/**
 * Проверяет валидность JWT токена
 * @param {string} token - JWT токен
 * @returns {boolean} true если токен валиден, false если нет
 */
export function isTokenValid(token) {
  if (!token || typeof token !== "string") return false;

  // Проверяем, что токен состоит из 3 частей, разделенных точками
  const parts = token.split(".");
  if (parts.length !== 3) return false;

  const payloadPart = parts[1];
  if (!payloadPart) return false;

  try {
    // Проверяем, что payload часть может быть декодирована
    const decodedPayload = atob(payloadPart);
    const payload = JSON.parse(decodedPayload);

    // Проверяем, что в payload есть поле exp
    if (!payload.exp) return false;

    const currentTime = Math.floor(Date.now() / 1000);
    return payload.exp > currentTime;
  } catch (e) {
    console.error("Ошибка проверки токена:", e);
    return false;
  }
}

/**
 * Проверяет, есть ли валидный токен в localStorage
 * @returns {boolean} true если есть валидный токен, false если нет
 */
export function hasValidToken() {
  const token = getJWTToken();
  return isTokenValid(token);
}

/**
 * Добавляет JWT токен к заголовкам запроса
 * @param {Object} headers - Объект заголовков
 * @returns {Object} Объект заголовков с добавленным Authorization
 */
export function addAuthHeader(headers = {}) {
  const token = getJWTToken();

  if (token && isTokenValid(token)) {
    return {
      ...headers,
      Authorization: `Bearer ${token}`,
    };
  }

  return headers;
}

/**
 * Выполняет авторизованный запрос к API
 * @param {string} url - URL для запроса
 * @param {Object} options - Опции запроса
 * @returns {Promise} Результат запроса
 */
export async function authenticatedFetch(url, options = {}) {
  const headers = addAuthHeader(options.headers || {});

  const response = await fetch(url, {
    ...options,
    headers,
  });

  if (response.status === 401) {
    // Токен недействителен, очищаем его
    clearJWTToken();
    throw new Error("Токен авторизации недействителен");
  }

  return response;
}
