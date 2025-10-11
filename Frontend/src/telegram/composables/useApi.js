import { ref } from 'vue'
import { authenticatedFetch, hasValidToken, getJWTToken } from '../auth/jwt.js'
import { useTelegramWebApp } from './useTelegramWebApp.js'

/**
 * Композиция для выполнения API запросов с автоматическим добавлением токена
 */
export function useApi() {
  const { jwtToken, sendAuthToServer } = useTelegramWebApp()
  
  /**
   * Выполняет API запрос с автоматическим добавлением префикса /api/ и токена
   * @param {string} url - URL для запроса (без префикса /api/)
   * @param {Object} options - Опции запроса
   * @param {boolean} autoAuth - Автоматически получать токен, если его нет
   * @returns {Promise} Результат запроса
   */
  const apiFetch = async (url, options = {}, autoAuth = true) => {

    const apiUrl = url.startsWith('/api/') ? url : `/api/${url.replace(/^\//, '')}`
    
    // Проверяем наличие токена
    if (autoAuth && !hasValidToken()) {
      try {
        // Пытаемся получить токен
        const authData = await sendAuthToServer()
        if (!authData?.token) {
          throw new Error('Не удалось получить токен аутентификации')
        }
      } catch (error) {
        console.error('Ошибка при получении токена:', error)
        throw error
      }
    }
    
    // Выполняем запрос с использованием authenticatedFetch
    return authenticatedFetch(apiUrl, options)
  }
  
  /**
   * Выполняет GET запрос к API
   * @param {string} url - URL для запроса
   * @param {Object} options - Дополнительные опции запроса
   * @returns {Promise} Результат запроса
   */
  const apiGet = (url, options = {}) => {
    return apiFetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        ...options.headers
      },
      ...options
    })
  }
  
  /**
   * Выполняет POST запрос к API
   * @param {string} url - URL для запроса
   * @param {Object} data - Данные для отправки
   * @param {Object} options - Дополнительные опции запроса
   * @returns {Promise} Результат запроса
   */
  const apiPost = (url, data, options = {}) => {
    return apiFetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        ...options.headers
      },
      body: JSON.stringify(data),
      ...options
    })
  }
  
  /**
   * Выполняет PUT запрос к API
   * @param {string} url - URL для запроса
   * @param {Object} data - Данные для отправки
   * @param {Object} options - Дополнительные опции запроса
   * @returns {Promise} Результат запроса
   */
  const apiPut = (url, data, options = {}) => {
    return apiFetch(url, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        ...options.headers
      },
      body: JSON.stringify(data),
      ...options
    })
  }
  
  /**
   * Выполняет DELETE запрос к API
   * @param {string} url - URL для запроса
   * @param {Object} options - Дополнительные опции запроса
   * @returns {Promise} Результат запроса
   */
  const apiDelete = (url, options = {}) => {
    return apiFetch(url, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
        ...options.headers
      },
      ...options
    })
  }
  
  return {
    apiFetch,
    apiGet,
    apiPost,
    apiPut,
    apiDelete
  }
}