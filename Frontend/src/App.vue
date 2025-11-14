<template>
  <div id="app">
    <transition name="fade" mode="out-in">
      <!-- Показываем основной layout только если в Telegram WebApp и есть JWT токен -->
      <MainLayout v-if="isTelegram && isInitialized && hasValidToken" />
      
      <!-- Показываем сообщение о необходимости Telegram, если не в WebApp -->
      <TelegramOnlyLayout v-else-if="!isTelegram && isInitialized" :telegram-link="telegramBotLink" />
      
      <!-- Лоадер во время инициализации -->
      <div v-else class="app-loader">
        <ThemedCard card-type="default" class="loader-card">
          <div class="loader-content">
            <div class="loader-spinner"></div>
            <p class="loader-text">{{ loaderMessage }}</p>
          </div>
        </ThemedCard>
      </div>
    </transition>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

import { ref, onMounted } from 'vue'
import MainLayout from '/src/layouts/MainLayout.vue'
import TelegramOnlyLayout from '/src/layouts/onlyTelegramUse.vue'
import ThemedCard from '/src/components/ThemedCard.vue'
import { useTelegramWebAppSingleton } from '/src/telegram/composables/useTelegramWebAppSingleton'
import { useLocalization, initLocalization } from '/src/locales/index.js'
import { useTelegramWebApp } from '/src/telegram/composables/useTelegramWebApp.js'

const { t } = useLocalization()
const router = useRouter()

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

const isInitialized = ref(false)
const loaderMessage = ref(t('app.initializing'))
const telegramBotLink = ref(`https://t.me/${import.meta.env.VITE_TELEGRAM_BOT_USERNAME || 'your_bot_username'}`)
const startParam = ref(null)

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
 * Извлечение параметров Telegram WebApp из URL
 */
const extractTelegramParamsFromUrl = () => {
  try {
    // Получаем текущий URL
    const url = window.location.href;
    console.log('🔍 Текущий URL:', url);
    
    // Проверяем, содержит ли URL параметры Telegram WebApp после хэша
    const hashIndex = url.indexOf('#');
    if (hashIndex !== -1) {
      const hashPart = url.substring(hashIndex + 1);
      console.log('🔍 Часть URL после хэша:', hashPart);
      
      // Проверяем, содержит ли часть после хэша параметры Telegram WebApp
      if (hashPart.includes('tgWebAppData=') || hashPart.includes('tgWebAppStartParam=')) {
        console.log('🔍 Найдены параметры Telegram WebApp в URL');
        
        // Извлекаем параметры из части после хэша
        // Обрабатываем различные форматы URL
        let paramsString = '';
        if (hashPart.includes('?')) {
          // Формат: /#/path?param1=value1&param2=value2
          paramsString = hashPart.split('?')[1];
        } else if (hashPart.includes('&') || hashPart.includes('=')) {
          // Формат: /#/param1=value1&param2=value2
          paramsString = hashPart;
        } else {
          // Формат: /#/tgWebAppData=value1&tgWebAppStartParam=value2
          // Ищем начало параметров
          const tgWebAppDataIndex = hashPart.indexOf('tgWebAppData=');
          const tgWebAppStartParamIndex = hashPart.indexOf('tgWebAppStartParam=');
          
          if (tgWebAppDataIndex !== -1) {
            paramsString = hashPart.substring(tgWebAppDataIndex);
          } else if (tgWebAppStartParamIndex !== -1) {
            paramsString = hashPart.substring(tgWebAppStartParamIndex);
          }
        }
        
        if (paramsString) {
          try {
            const params = new URLSearchParams(paramsString);
            const tgWebAppData = params.get('tgWebAppData');
            const tgWebAppStartParam = params.get('tgWebAppStartParam');
            
            console.log('🔍 Извлеченные параметры:', { tgWebAppData, tgWebAppStartParam });
            
            // Если есть параметры, сохраняем их для дальнейшей обработки
            if (tgWebAppStartParam) {
              startParam.value = tgWebAppStartParam;
              console.log('✅ Сохранен параметр startParam из URL:', startParam.value);
            }
          } catch (parseError) {
            console.error('❌ Ошибка при парсинге параметров из URL:', parseError);
          }
        }
      }
    }
  } catch (error) {
    console.error('❌ Ошибка при извлечении параметров Telegram WebApp из URL:', error);
  }
};

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
    await initLocalization()

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
      }

      else
      {
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
    })

  } catch (error) {
    console.error('❌ Ошибка инициализации:', error)
    authError.value = error.message || t('app.authError')
    isInitialized.value = true
  }
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


onMounted(() => {
  initializeApp()
})
</script>

<style>
/* Стили остаются без изменений */
@import './assets/css/app.css';

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  transition: background-color 0.3s ease, color 0.3s ease, border-color 0.3s ease;
}

body {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  background-color: var(--tg-theme-bg-color, #f5f5f5);
  color: var(--tg-theme-text-color, #000000);
  transition: background-color 0.3s ease, color 0.3s ease;
  overflow: hidden;
}

/* Стили лоадера */
.app-loader {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
  width: 100vw;
  background-color: var(--tg-theme-bg-color, #ffffff);
  padding: 20px;
}

.loader-card {
  background: var(--tg-theme-secondary-bg-color, white);
  border-radius: 12px;
  padding: 32px;
  box-shadow: 0 1px 6px rgba(0, 0, 0, 0.08);
  border: 1px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
  max-width: 300px;
  width: 100%;
  text-align: center;
}

.loader-content {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.loader-spinner {
  width: 40px;
  height: 40px;
  border: 3px solid var(--tg-theme-secondary-bg-color, #f0f0f0);
  border-top: 3px solid var(--tg-theme-button-color, #667eea);
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 20px;
}

.loader-text {
  color: var(--tg-theme-text-color, #000000);
  font-size: 16px;
  font-weight: 500;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}



/* Анимации перехода между страницами */

/* Slide Left Animation */
.slide-left-enter-active,
.slide-left-leave-active,
.slide-right-enter-active,
.slide-right-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}

.slide-left-enter-from {
  opacity: 0;
  transform: translateX(30px);
}

.slide-left-leave-to {
  opacity: 0;
  transform: translateX(-30px);
}

.slide-left-enter-to,
.slide-left-leave-from {
  opacity: 1;
  transform: translateX(0);
}

/* Slide Right Animation */
.slide-right-enter-from {
  opacity: 0;
  transform: translateX(-30px);
}

.slide-right-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

.slide-right-enter-to,
.slide-right-leave-from {
  opacity: 1;
  transform: translateX(0);
}

/* Fade Animation */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.25s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* Scale Animation */
.scale-enter-active,
.scale-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.scale-enter-from {
  opacity: 0;
  transform: scale(0.9);
}

.scale-leave-to {
  opacity: 0;
  transform: scale(1.1);
}

/* Slide Up Animation */
.slide-up-enter-active,
.slide-up-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.slide-up-enter-from {
  opacity: 0;
  transform: translateY(20px);
}

.slide-up-leave-to {
  opacity: 0;
  transform: translateY(-20px);
}

/* Улучшенная анимация для контента внутри страниц */
.page-content {
  animation: pageFadeIn 0.4s ease-out;
}

@keyframes pageFadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Анимация для карточек с задержкой */
.stagger-animation>* {
  opacity: 0;
  transform: translateY(20px);
  animation: staggerFadeIn 0.5s ease-out forwards;
}

.stagger-animation>*:nth-child(1) {
  animation-delay: 0.1s;
}

.stagger-animation>*:nth-child(2) {
  animation-delay: 0.2s;
}

.stagger-animation>*:nth-child(3) {
  animation-delay: 0.3s;
}

.stagger-animation>*:nth-child(4) {
  animation-delay: 0.4s;
}

.stagger-animation>*:nth-child(5) {
  animation-delay: 0.5s;
}

.stagger-animation>*:nth-child(6) {
  animation-delay: 0.6s;
}

@keyframes staggerFadeIn {
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Оптимизация производительности анимаций */
.slide-left-enter-active,
.slide-left-leave-active,
.slide-right-enter-active,
.slide-right-leave-active {
  will-change: transform, opacity;
}

/* Адаптация анимаций для мобильных устройств */
@media (max-width: 768px) {
  .slide-left-enter-from {
    transform: translateX(20px);
  }

  .slide-left-leave-to {
    transform: translateX(-20px);
  }

  .slide-right-enter-from {
    transform: translateX(-20px);
  }

  .slide-right-leave-to {
    transform: translateX(20px);
  }
}

/* Отключение анимаций для пользователей, которые предпочитают их не использовать */
@media (prefers-reduced-motion: reduce) {
  * {
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 0.01ms !important;
  }
}

* {
  -webkit-tap-highlight-color: transparent !important;
  tap-highlight-color: transparent !important;
  outline: none !important;
}
</style>