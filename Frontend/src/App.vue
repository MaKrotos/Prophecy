<template>
  <div id="app">
    <!-- Показываем основной layout только если в Telegram WebApp и есть JWT токен -->
    <MainLayout v-if="isTelegram && isInitialized && hasValidToken" />

    <!-- Показываем страницу ошибки авторизации, если есть ошибка -->
    <AuthErrorLayout v-else-if="isTelegram && isInitialized && authError" :error-message="authError" @retry="retryAuth"
      @try-later="tryLater" />

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
import { ref, onMounted, onErrorCaptured, computed } from 'vue'
import MainLayout from '/src/layouts/MainLayout.vue'
import TelegramOnlyLayout from '/src/layouts/onlyTelegramUse.vue'
import AuthErrorLayout from '/src/layouts/AuthErrorLayout.vue'
import { useTelegramWebApp } from '/src/telegram/composables/useTelegramWebApp'

// Используем хук Telegram WebApp
const {
  telegramUser,
  isTelegram,
  authHash,
  authData,
  themeParams,
  isDarkTheme,
  sendAuthToServer,
  getAuthData,
  refreshTheme
} = useTelegramWebApp()

const isInitialized = ref(false)
const authError = ref(null)
const loaderMessage = ref('Инициализация приложения...')
const telegramBotLink = ref('https://t.me/your_bot_username') // Замените на реальную ссылку

// Проверяем наличие валидного JWT токена
const hasValidToken = computed(() => {
  const token = localStorage.getItem('jwt_token')
  if (!token) return false

  // Здесь можно добавить проверку срока действия токена
  try {
    const payload = JSON.parse(atob(token.split('.')[1]))
    const currentTime = Math.floor(Date.now() / 1000)
    return payload.exp > currentTime
  } catch (e) {
    console.error('Ошибка проверки токена:', e)
    return false
  }
})

// Функция для сохранения JWT токена
const saveJWTToken = (token) => {
  if (token) {
    localStorage.setItem('jwt_token', token)
    console.log('✅ JWT токен сохранен в localStorage')
  }
}

// Функция для очистки JWT токена
const clearJWTToken = () => {
  localStorage.removeItem('jwt_token')
  console.log('🗑️ JWT токен удален из localStorage')
}

onMounted(async () => {
  console.log('🚀 App mounted, initializing...')

  // Даем время на инициализацию Telegram WebApp
  setTimeout(async () => {
    isInitialized.value = true
    console.log('✅ App initialized:', {
      isTelegram: isTelegram.value,
      user: telegramUser.value,
      theme: themeParams.value
    })

    // Если в Telegram, отправляем данные аутентификации на сервер
    if (isTelegram.value && authHash.value) {
      loaderMessage.value = 'Авторизация через Telegram...'
      console.log('📡 Sending auth data to server...')

      try {
        const result = await sendAuthToServer('/auth/telegram')

        if (result && result.token) {
          // Сохраняем JWT токен
          saveJWTToken(result.token)
          console.log('✅ Успешная авторизация, токен сохранен')
          authError.value = null
        } else {
          throw new Error('Сервер не вернул JWT токен')
        }
      } catch (error) {
        console.error('❌ Failed to send auth data:', error)
        authError.value = error.message || 'Не удалось выполнить авторизацию. Попробуйте позже.'
        clearJWTToken() // Очищаем старый токен при ошибке
      }
    }
  }, 1000)
})

// Обработчик глобальных ошибок
onErrorCaptured((error, instance, info) => {
  console.error('💥 Global error captured:', error, info)
  // Можно отправить ошибку в сервис мониторинга
  return false
})

// Функция для повторной попытки авторизации
const retryAuth = async () => {
  authError.value = null
  loaderMessage.value = 'Повторная авторизация...'
  isInitialized.value = false

  // Небольшая задержка для отображения лоадера
  await new Promise(resolve => setTimeout(resolve, 500))

  // Повторная попытка авторизации
  try {
    const result = await sendAuthToServer('/auth/telegram')

    if (result && result.token) {
      saveJWTToken(result.token)
      console.log('✅ Повторная авторизация успешна')
      authError.value = null
    } else {
      throw new Error('Сервер не вернул JWT токен')
    }
  } catch (error) {
    console.error('❌ Повторная авторизация не удалась:', error)
    authError.value = error.message || 'Не удалось выполнить авторизацию. Попробуйте позже.'
    clearJWTToken()
  } finally {
    isInitialized.value = true
  }
}

// Функция для отложенной попытки
const tryLater = () => {
  // Просто очищаем ошибку и показываем основной интерфейс
  // В реальном приложении можно добавить логику отложенной авторизации
  authError.value = null
  console.log('🕒 Пользователь выбрал "Попробовать позже"')
}

// Глобальные функции для отладки (можно убрать в продакшене)
if (import.meta.env.DEV) {
  window.$telegram = {
    getUser: () => telegramUser.value,
    getAuthData: () => getAuthData(),
    refreshTheme: () => refreshTheme(),
    isTelegram: () => isTelegram.value,
    hasValidToken: () => hasValidToken.value,
    clearToken: () => clearJWTToken()
  }
}
</script>

<style>
@import './assets/css/app.css';
/* Глобальные стили остаются здесь */


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