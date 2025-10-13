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
    // Декодируем base64url
    const decodedPayload = decodeURIComponent(
      atob(payloadPart.replace(/-/g, "+").replace(/_/g, "/"))
        .split("")
        .map((c) => "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2))
        .join("")
    );

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
 * Выполняет авторизованный запрос к API с возможностью повторных попыток
 * @param {string} url - URL для запроса
 * @param {Object} options - Опции запроса
 * @param {number} maxRetries - Максимальное количество попыток
 * @returns {Promise} Результат запроса
 */
export async function authenticatedFetch(url, options = {}, maxRetries = 3) {
  let lastError;

  for (let attempt = 0; attempt <= maxRetries; attempt++) {
    try {
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
    } catch (error) {
      lastError = error;
      console.warn(`Попытка ${attempt + 1} не удалась:`, error.message);

      // Если это не последняя попытка, ждем перед повтором
      if (attempt < maxRetries) {
        // Экспоненциальная задержка: 1s, 2s, 4s, ...
        const delay = Math.pow(2, attempt) * 1000;
        await new Promise((resolve) => setTimeout(resolve, delay));
      }
    }
  }

  // Если все попытки исчерпаны, выбрасываем последнюю ошибку
  throw lastError;
}
