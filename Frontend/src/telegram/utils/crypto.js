/**
 * Функции для шифрования и дешифрования данных
 */

/**
 * Шифрует данные с использованием простого XOR шифрования
 * @param {string} data - Данные для шифрования
 * @param {string} key - Ключ шифрования
 * @returns {string} Зашифрованные данные в формате base64
 */
export function encryptData(data, key) {
  console.log("🔍 Шифрование данных");
  try {
    // Преобразуем данные и ключ в байтовые массивы
    const dataBytes = new TextEncoder().encode(data);
    const keyBytes = new TextEncoder().encode(key);
    
    // Выполняем XOR шифрование
    const encryptedBytes = dataBytes.map((byte, index) => 
      byte ^ keyBytes[index % keyBytes.length]
    );
    
    // Преобразуем результат в base64
    const encryptedData = btoa(String.fromCharCode(...encryptedBytes));
    console.log("✅ Данные зашифрованы");
    return encryptedData;
  } catch (e) {
    console.error("Ошибка шифрования данных:", e);
    return null;
  }
}

/**
 * Дешифрует данные с использованием простого XOR шифрования
 * @param {string} encryptedData - Зашифрованные данные в формате base64
 * @param {string} key - Ключ шифрования
 * @returns {string} Расшифрованные данные
 */
export function decryptData(encryptedData, key) {
  console.log("🔍 Дешифрование данных");
  try {
    // Преобразуем base64 в байтовый массив
    const encryptedBytes = new Uint8Array(
      atob(encryptedData).split("").map(char => char.charCodeAt(0))
    );
    
    // Преобразуем ключ в байтовый массив
    const keyBytes = new TextEncoder().encode(key);
    
    // Выполняем XOR дешифрование
    const decryptedBytes = encryptedBytes.map((byte, index) => 
      byte ^ keyBytes[index % keyBytes.length]
    );
    
    // Преобразуем результат в строку
    const decryptedData = new TextDecoder().decode(decryptedBytes);
    console.log("✅ Данные расшифрованы");
    return decryptedData;
  } catch (e) {
    console.error("Ошибка дешифрования данных:", e);
    return null;
  }
}

/**
 * Генерирует ключ шифрования на основе данных сессии и времени
 * @param {number} playerId - ID игрока
 * @param {number} sessionId - ID сессии
 * @param {number} timestamp - Временная метка
 * @returns {string} Ключ шифрования
 */
export function generateEncryptionKey(playerId, sessionId, timestamp) {
  console.log("🔍 Генерация ключа шифрования");
  try {
    // Создаем ключ на основе ID игрока, ID сессии и времени
    const keyString = `${playerId}-${sessionId}-${timestamp}`;
    // Хешируем ключ с помощью простого алгоритма
    let hash = 0;
    for (let i = 0; i < keyString.length; i++) {
      const char = keyString.charCodeAt(i);
      hash = ((hash << 5) - hash) + char;
      hash = hash & hash; // Преобразуем в 32-битное целое
    }
    
    // Преобразуем хеш в строку
    const key = Math.abs(hash).toString();
    console.log("✅ Ключ шифрования сгенерирован");
    return key;
  } catch (e) {
    console.error("Ошибка генерации ключа шифрования:", e);
    return null;
  }
}