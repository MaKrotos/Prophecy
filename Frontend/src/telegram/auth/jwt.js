/**
 * Функции для работы с JWT токенами
 */

/**
 * Сохраняет JWT токен в localStorage
 * @param {string} token - JWT токен
 */
export function saveJWTToken(token) {
  console.log(
    "🔍 Сохранение JWT токена в localStorage:",
    token ? "Токен предоставлен" : "Токен отсутствует"
  );
  if (token) {
    localStorage.setItem("jwt_token", token);
    console.log("✅ JWT токен сохранен в localStorage");
  } else {
    console.log("⚠️ Токен отсутствует, сохранение не выполнено");
  }
}

/**
 * Получает JWT токен из localStorage
 * @returns {string|null} JWT токен или null
 */
export function getJWTToken() {
  const token = localStorage.getItem("jwt_token");
  console.log(
    "🔍 Получение JWT токена из localStorage:",
    token ? "Токен найден" : "Токен отсутствует"
  );
  return token;
}

/**
 * Удаляет JWT токен из localStorage
 */
export function clearJWTToken() {
  console.log("🔍 Удаление JWT токена из localStorage");
  localStorage.removeItem("jwt_token");
  console.log("🗑️ JWT токен удален из localStorage");
}

/**
 * Проверяет валидность JWT токена
 * @param {string} token - JWT токен
 * @returns {boolean} true если токен валиден, false если нет
 */
export function isTokenValid(token) {
  console.log("🔍 Проверка валидности токена");
  if (!token || typeof token !== "string") {
    console.log("⚠️ Токен отсутствует или не является строкой");
    return false;
  }

  // Проверяем, что токен состоит из 3 частей, разделенных точками
  const parts = token.split(".");
  if (parts.length !== 3) {
    console.log("⚠️ Токен не состоит из 3 частей");
    return false;
  }

  const payloadPart = parts[1];
  if (!payloadPart) {
    console.log("⚠️ Отсутствует payload часть токена");
    return false;
  }

  try {
    // Декодируем base64url
    const decodedPayload = decodeURIComponent(
      atob(payloadPart.replace(/-/g, "+").replace(/_/g, "/"))
        .split("")
        .map((c) => "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2))
        .join("")
    );
    console.log("🔍 Раскодированный payload:", decodedPayload);

    const payload = JSON.parse(decodedPayload);
    console.log("🔍 Распарсенный payload:", payload);

    // Проверяем, что в payload есть поле exp
    if (!payload.exp) {
      console.log("⚠️ В payload отсутствует поле exp");
      return false;
    }

    const currentTime = Math.floor(Date.now() / 1000);
    const isValid = payload.exp > currentTime;
    console.log("🔍 Сравнение времени окончания действия токена:", {
      exp: payload.exp,
      currentTime: currentTime,
      isValid: isValid,
    });
    return isValid;
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
  console.log(
    "🔍 Проверка наличия валидного токена:",
    token ? "Токен найден" : "Токен отсутствует"
  );
  const isValid = isTokenValid(token);
  console.log("🔍 Результат проверки валидности токена:", isValid);
  return isValid;
}

/**
 * Добавляет JWT токен к заголовкам запроса
 * @param {Object} headers - Объект заголовков
 * @returns {Object} Объект заголовков с добавленным Authorization
 */
export function addAuthHeader(headers = {}) {
  const token = getJWTToken();
  console.log(
    "🔍 Добавление заголовка авторизации, токен:",
    token ? "Токен найден" : "Токен отсутствует"
  );

  if (token && isTokenValid(token)) {
    const authHeaders = {
      ...headers,
      Authorization: `Bearer ${token}`,
    };
    console.log("✅ Заголовок авторизации добавлен:", authHeaders);
    return authHeaders;
  } else {
    console.log(
      "⚠️ Заголовок авторизации не добавлен, токен отсутствует или недействителен"
    );
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
  console.log("🔍 Выполнение авторизованного запроса:", url);
  let lastError;

  for (let attempt = 0; attempt <= maxRetries; attempt++) {
    try {
      const headers = addAuthHeader(options.headers || {});
      console.log("🔍 Заголовки запроса:", headers);
      const response = await fetch(url, {
        ...options,
        headers,
      });
      console.log("🔍 Ответ от сервера:", response.status);

      if (response.status === 401) {
        // Токен недействителен, очищаем его
        console.log("⚠️ Токен авторизации недействителен, очистка токена");
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
        console.log(`⏳ Ожидание ${delay}ms перед следующей попыткой`);
        await new Promise((resolve) => setTimeout(resolve, delay));
      }
    }
  }

  // Если все попытки исчерпаны, выбрасываем последнюю ошибку
  console.error(
    "❌ Все попытки запроса исчерпаны, последняя ошибка:",
    lastError.message
  );
  throw lastError;
}
