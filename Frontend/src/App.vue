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
import { onMounted } from 'vue'
import MainLayout from '/src/layouts/MainLayout.vue'
import TelegramOnlyLayout from '/src/layouts/onlyTelegramUse.vue'
import ThemedCard from '/src/components/ThemedCard.vue'
import { useAppInitialization } from '/src/telegram/composables/useAppInitialization'

const {
  isInitialized,
  loaderMessage,
  telegramBotLink,
  isTelegram,
  hasValidToken,
  initializeApp
} = useAppInitialization()

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