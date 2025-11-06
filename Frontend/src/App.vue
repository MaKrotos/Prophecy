<template>
  <div id="app">
    <!-- Показываем основной layout только если в Telegram WebApp и есть JWT токен -->
    <MainLayout v-if="isTelegram && isInitialized && hasValidToken" />
    
    <!-- Показываем сообщение о необходимости Telegram, если не в WebApp -->
    <TelegramOnlyLayout v-else-if="!isTelegram && isInitialized" :telegram-link="telegramBotLink" />
    
    <!-- Лоадер во время инициализации -->
    <div v-else class="app-loader">
      <div class="loader-content">
        <div class="loader-spinner"></div>
        <p class="loader-text">{{ loaderMessage }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import MainLayout from '/src/layouts/MainLayout.vue'
import TelegramOnlyLayout from '/src/layouts/onlyTelegramUse.vue'
import { useTelegramWebAppSingleton } from '/src/telegram/composables/useTelegramWebAppSingleton'
import { useLocalization, initLocalization } from '/src/locales/index.js'

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

const isInitialized = ref(false)
const loaderMessage = ref(t('app.initializing'))
const telegramBotLink = ref(`https://t.me/${import.meta.env.VITE_TELEGRAM_BOT_USERNAME || 'your_bot_username'}`)
const startParam = ref(null)

/**
 * Инициализация приложения
 */
const initializeApp = async () => {
  console.log('🚀 App mounted, initializing...')

  try {
    // Инициализируем локализацию
    await initLocalization()

    // Если это Telegram, ждем его готовности
    if (isTelegram.value) {
      await initializeTelegramApp()
    } else {
      console.log('🌐 Это не Telegram WebApp')
    }

    isInitialized.value = true
    console.log('✅ App initialized', {
      isTelegram: isTelegram.value,
      isTelegramReady: isTelegramReady?.value,
      authHash: !!authHash.value,
      hasValidToken: hasValidToken()
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

  // Авторизация только если Telegram готов и есть хэш
  if (isTelegramReady?.value && authHash.value) {
    loaderMessage.value = t('app.loading')
    console.log('📡 Отправка данных аутентификации на сервер...')

    await sendAuthToServer('auth/telegram', 3)
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
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.loader-content {
  text-align: center;
  color: white;
}

.loader-spinner {
  width: 50px;
  height: 50px;
  border: 4px solid rgba(255, 255, 255, 0.3);
  border-top: 4px solid white;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto 20px;
}

.loader-text {
  font-size: 16px;
  opacity: 0.9;
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