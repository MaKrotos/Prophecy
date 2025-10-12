import { ref, computed } from 'vue'

// Поддерживаемые языки
export const supportedLocales = ['ru', 'en']

// Текущая локаль (по умолчанию русская)
export const currentLocale = ref('ru')

// Загруженные переводы
export const translations = ref({})

// Предварительно загружаем все JSON файлы локалей с помощью import.meta.glob
const localeModules = import.meta.glob('./*.json', { import: 'default', eager: true })

// Асинхронная загрузка переводов
export const loadLocaleMessages = async (locale) => {
  try {
    // Получаем сообщения из предварительно загруженных модулей
    const modulePath = `./${locale}.json`
    if (localeModules[modulePath]) {
      translations.value[locale] = localeModules[modulePath]
      console.log(`✅ Загружены переводы для локали: ${locale}`)
    } else {
      throw new Error(`Locale file not found: ${locale}.json`)
    }
  } catch (error) {
    console.error(`❌ Ошибка загрузки переводов для локали ${locale}:`, error)
  }
}

// Инициализация локализации
export const initLocalization = async () => {
  // Загружаем все поддерживаемые локали
  for (const locale of supportedLocales) {
    await loadLocaleMessages(locale)
  }
  
  // Определяем начальную локаль из Telegram WebApp, если доступна
  const telegramLocale = getTelegramLocale()
  if (telegramLocale && supportedLocales.includes(telegramLocale)) {
    currentLocale.value = telegramLocale
  } else {
    // По умолчанию русский
    currentLocale.value = 'ru'
  }
}

// Получение языка из Telegram WebApp
export const getTelegramLocale = () => {
  if (typeof window !== 'undefined' && window.Telegram?.WebApp?.initDataUnsafe?.user?.language_code) {
    const langCode = window.Telegram.WebApp.initDataUnsafe.user.language_code
    // Преобразуем код языка к поддерживаемым локалям
    // Например, 'en-US' -> 'en', 'ru-RU' -> 'ru'
    const shortLangCode = langCode.split('-')[0].toLowerCase()
    return shortLangCode
  }
  return null
}

// Функция для получения перевода по ключу
export const t = (key, params = {}) => {
  const locale = currentLocale.value
  const messages = translations.value[locale]
  
  if (!messages) {
    console.warn(`Нет переводов для локали: ${locale}`)
    return key
  }
  
  // Ищем перевод по ключу (поддержка вложенных ключей через точку)
  const keys = key.split('.')
  let translation = messages
  
  for (const k of keys) {
    if (translation && typeof translation === 'object' && k in translation) {
      translation = translation[k]
    } else {
      console.warn(`Перевод не найден для ключа: ${key} в локали: ${locale}`)
      return key
    }
  }
  
  // Если переданы параметры, заменяем их в строке
  if (typeof translation === 'string' && Object.keys(params).length > 0) {
    return translation.replace(/\{(\w+)\}/g, (match, key) => {
      return params[key] !== undefined ? params[key] : match
    })
  }
  
  return translation
}

// Композиция для использования в компонентах
export const useLocalization = () => {
  const locale = computed(() => currentLocale.value)
  
  const changeLocale = (newLocale) => {
    if (supportedLocales.includes(newLocale)) {
      currentLocale.value = newLocale
      console.log(`🌐 Локаль изменена на: ${newLocale}`)
    } else {
      console.warn(`🚫 Локаль ${newLocale} не поддерживается`)
    }
  }
  
  return {
    locale,
    changeLocale,
    t
  }
}