import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useTelegramWebAppSingleton } from './useTelegramWebAppSingleton'
import { useTelegramWebApp } from './useTelegramWebApp'
import { useLocalization, initLocalization } from '/src/locales/index.js'
import { useUrlParams } from './useUrlParams.js'

/**
 * Композабл для управления инициализацией приложения
 * @returns {Object} Объект с реактивными переменными и методами
 */
export function useAppInitialization() {
  const router = useRouter()
  const { t } = useLocalization()
  
  const {
    telegramUser,
    isTelegram,
    authHash,
    themeParams,
    isTelegramReady,
    sendAuthToServer,
    waitForTelegramReady,
    hasValidToken,
    checkTelegramIdConsistency
  } = useTelegramWebAppSingleton()
  
  const {
    parseAndHandleStartParam
  } = useTelegramWebApp()
  
  const { startParam, extractTelegramParamsFromUrl } = useUrlParams()
  
  const isInitialized = ref(false)
  const loaderMessage = ref(t('app.initializing'))
  const authError = ref(null)
  const telegramBotLink = ref(`https://t.me/${import.meta.env.VITE_TELEGRAM_BOT_USERNAME || 'your_bot_username'}`)

  /**
   * Проверяет, первый ли раз пользователь открывает приложение
   * @returns {boolean} true, если первый раз, false если нет
   */
  const isFirstTimeUser = () => {
    const hasVisited = localStorage.getItem('hasVisitedApp');
    return !hasVisited;
  }

  /**
   * Отмечает, что пользователь уже посещал приложение
   */
  const markAsVisited = () => {
    localStorage.setItem('hasVisitedApp', 'true');
  }

  /**
   * Инициализация Telegram приложения
   */
  const initializeTelegramApp = async () => {
    loaderMessage.value = t('app.loading')
    console.log('⏳ Ожидание готовности Telegram WebApp...')

    await waitForTelegramReady(5000)
    console.log('✅ Telegram WebApp готов')

    // Проверяем соответствие Telegram ID
    const isConsistent = checkTelegramIdConsistency()
    if (!isConsistent) {
      loaderMessage.value = t('app.retryAuth')
      console.log('🔄 Необходима повторная авторизация из-за несоответствия Telegram ID')
    }

    // Проверяем, есть ли параметр startParam
    if (startParam.value) {
      console.log('🔗 Параметр startParam доступен при инициализации Telegram:', startParam.value);
    }

    // Авторизация только если Telegram готов и есть хэш
    if (isTelegramReady?.value && authHash.value) {
      loaderMessage.value = t('app.loading')
      console.log('📡 Отправка данных аутентификации на сервер...')

      try {
        await sendAuthToServer('auth/telegram', 3)
        console.log('✅ Данные аутентификации успешно отправлены на сервер')
      } catch (error) {
        console.error('❌ Ошибка при отправке данных аутентификации на сервер:', error)
        throw new Error(t('auth.error.message'))
      }
    } else if (!authHash.value) {
      console.warn('⚠️ Нет хэша аутентификации')
      throw new Error(t('auth.error.message'))
    }
  }

  /**
   * Инициализация приложения
   */
  const initializeApp = async () => {
    console.log('🚀 App mounted, initializing...')

    try {
      // Извлекаем параметры Telegram WebApp из URL перед инициализацией
      extractTelegramParamsFromUrl();
      
      // Логируем состояние параметров перед инициализацией
      console.log('🔍 Состояние параметров перед инициализацией:', {
        isTelegram: isTelegram.value,
        startParam: startParam.value,
        hasAuthHash: !!authHash.value
      });
      
      // Инициализируем локализацию
      await initLocalization() // Инициализация локализации

      // Если это Telegram, ждем его готовности
      if (isTelegram.value) {
        await initializeTelegramApp()
      } else {
        // Даже если это не Telegram WebApp, передаем параметр startParam если он есть
        if (startParam.value) {
          console.log('🔗 Найден параметр startParam вне Telegram:', startParam.value);
          // Параметр startParam будет использован композаблом useTelegramWebApp при инициализации
          // даже если приложение запущено не в Telegram WebApp
          
          // Пытаемся обработать параметр startParam даже вне Telegram
          try {
            // Используем функцию из композабла для обработки параметра
            if (typeof parseAndHandleStartParam === 'function') {
              parseAndHandleStartParam(startParam.value);
              console.log('✅ Параметр startParam обработан вне Telegram');
            } else {
              console.warn('⚠️ Функция parseAndHandleStartParam недоступна');
            }
          } catch (error) {
            console.error('❌ Ошибка при обработке startParam вне Telegram:', error);
          }
        }
        console.log('🌐 Это не Telegram WebApp')
      }

      // Проверяем, первый ли раз пользователь открывает приложение
      // Делаем это только если нет startParam (пользователь не переходит по реферальной ссылке)
      // и только для Telegram WebApp
      if (!startParam.value && isTelegram.value) {
        if (isFirstTimeUser()) {
          console.log('Первый раз открывает приложение, перенаправляем на страницу правил');
          markAsVisited();
          // Перенаправляем на страницу правил, но только если мы не на ней уже
          if (window.location.hash !== '#/rules') {
            router.push('/rules');
          }
        } else {
          router.push('/');
        }
      }

      isInitialized.value = true
      console.log('✅ App initialized', {
        isTelegram: isTelegram.value,
        isTelegramReady: isTelegramReady?.value,
        authHash: !!authHash.value,
        hasValidToken: hasValidToken(),
        startParam: startParam.value
      });

    } catch (error) {
      console.error('❌ Ошибка инициализации:', error)
      authError.value = error.message || t('app.authError')
      isInitialized.value = true
    }
  }

  return {
    isInitialized,
    loaderMessage,
    authError,
    telegramBotLink,
    isTelegram,
    hasValidToken,
    initializeApp
  }
}
