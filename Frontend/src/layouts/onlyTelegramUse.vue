<template>
  <div class="telegram-only-layout">
    <div class="container">
      <ThemedCard card-type="default" class="content-card">
        <div class="header">
          <div class="telegram-icon">
            <svg width="64" height="64" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M12 2C6.48 2 2 6.48 2 12C2 17.52 6.48 22 12 22C17.52 22 22 17.52 22 12C22 6.48 17.52 2 12 2Z"
                fill="#0088CC" />
              <path
                d="M16.64 8.08C16.49 8.03 16.33 8 16.17 8C15.86 8 15.56 8.1 15.31 8.28L8.94 12.65C8.24 13.11 7.74 13.85 7.56 14.69C7.38 15.53 7.53 16.41 7.99 17.11C8.45 17.81 9.19 18.31 10.03 18.49C10.87 18.67 11.75 18.52 12.45 18.06L18.82 13.69C19.32 13.35 19.64 12.82 19.69 12.23C19.74 11.64 19.5 11.06 19.05 10.65L16.64 8.08Z"
                fill="white" />
              <path
                d="M9.75 14.5C9.61 14.5 9.47 14.47 9.34 14.41C9.21 14.35 9.09 14.27 9 14.16C8.91 14.05 8.84 13.92 8.8 13.78C8.76 13.64 8.75 13.49 8.77 13.35C8.79 13.21 8.84 13.07 8.92 12.95C9 12.83 9.1 12.73 9.22 12.65L12.45 10.45C12.67 10.3 12.94 10.25 13.19 10.31C13.44 10.37 13.65 10.53 13.77 10.75C13.89 10.97 13.91 11.23 13.82 11.46C13.73 11.69 13.54 11.87 13.3 11.95L10.07 14.15C9.98 14.2 9.87 14.23 9.75 14.23V14.5Z"
                fill="#0088CC" />
            </svg>
          </div>
          <h1 class="title">{{ t('telegram.only.title') }}</h1>
          <p class="subtitle">{{ t('telegram.only.subtitle') }}</p>
        </div>

        <div class="content">
          <ThemedCard card-type="secondary" class="info-card">
            <h2>{{ t('telegram.only.instructions.title') }}</h2>
            <ol class="instructions">
              <li>{{ t('telegram.only.instructions.step1') }}</li>
              <li v-html="t('telegram.only.instructions.step2')"></li>
              <li>{{ t('telegram.only.instructions.step3') }}</li>
            </ol>
          </ThemedCard>

          <div class="bot-link">
            <ThemedButton buttonType="primary" class="telegram-button" @click="openTelegram">
              <span class="button-icon">📱</span>
              {{ t('telegram.only.button') }}
            </ThemedButton>
          </div>

          <ThemedCard card-type="secondary" class="features">
            <h3>{{ t('telegram.only.features.title') }}</h3>
            <ul class="features-list">
              <li v-html="t('telegram.only.features.feature1')"></li>
              <li v-html="t('telegram.only.features.feature2')"></li>
              <li v-html="t('telegram.only.features.feature3')"></li>
              <li v-html="t('telegram.only.features.feature4')"></li>
            </ul>
          </ThemedCard>
        </div>

        <div class="footer">
          <p class="footer-text">
            {{ t('telegram.only.footer') }}
          </p>
        </div>
      </ThemedCard>
    </div>
  </div>
</template>

<script setup>
import { useLocalization } from '@/locales/index.js';
import ThemedButton from '@/components/ThemedButton.vue';
import ThemedCard from '@/components/ThemedCard.vue';
const { t } = useLocalization();

const openTelegram = () => {
  // Можно добавить аналитику или дополнительную логику
  console.log('Redirecting to Telegram bot...');
};
</script>

<style scoped>
.telegram-only-layout {
  background-color: var(--tg-theme-bg-color, #ffffff);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  min-height: 100vh;
}

.container {
  width: 100%;
  max-width: 500px;
}

.content-card {
  background: var(--tg-theme-secondary-bg-color, white);
  border-radius: 12px;
  padding: 32px;
  box-shadow: 0 1px 6px rgba(0, 0, 0, 0.08);
  border: 1px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
  text-align: center;
}

.header {
  margin-bottom: 30px;
}

.telegram-icon {
  margin-bottom: 20px;
}

.title {
  color: var(--tg-theme-text-color, #000000);
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 10px;
  transition: color 0.3s ease;
}

.subtitle {
  color: var(--tg-theme-hint-color, #666666);
  font-size: 1.1rem;
  margin-bottom: 0;
  transition: color 0.3s ease;
}

.content {
  margin: 30px 0;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.info-card,
.features {
  background: var(--tg-theme-secondary-bg-color, #f8f9fa);
  border-radius: 12px;
  padding: 24px;
  text-align: left;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.05);
  border: 1px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
  transition: all 0.3s ease;
}

.info-card:hover,
.features:hover {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.info-card h2,
.features h3 {
  color: var(--tg-theme-text-color, #333333);
  margin-bottom: 15px;
  font-size: 1.3rem;
  font-weight: 600;
  transition: color 0.3s ease;
}

.instructions,
.features-list {
  color: var(--tg-theme-text-color, #555555);
  line-height: 1.6;
  padding-left: 20px;
  transition: color 0.3s ease;
}

.instructions li,
.features-list li {
  margin-bottom: 8px;
}

.bot-link {
  margin: 20px 0;
}

.telegram-button {
  width: 100%;
  justify-content: center;
}

.footer {
  margin-top: 30px;
  padding-top: 20px;
  border-top: 1px solid var(--tg-theme-secondary-bg-color, #eeeeee);
}

.footer-text {
  color: var(--tg-theme-hint-color, #888888);
  font-size: 0.9rem;
  line-height: 1.4;
  transition: color 0.3s ease;
}

/* Адаптивность */
@media (max-width: 480px) {
  .telegram-only-layout {
    padding: 16px;
  }

  .content-card {
    padding: 24px;
  }

  .title {
    font-size: 1.75rem;
  }

  .info-card,
  .features {
    padding: 20px;
  }
}

/* Дополнительные стили для очень маленьких экранов */
@media (max-height: 600px) {
  .telegram-only-layout {
    align-items: flex-start;
    padding-top: 20px;
  }
}

/* Анимации для элементов */
.content-card {
  animation: fadeInUp 0.4s ease-out;
}

.info-card,
.features {
  opacity: 0;
  transform: translateY(20px);
  animation: fadeInUp 0.5s ease-out forwards;
}

.info-card {
  animation-delay: 0.1s;
}

.bot-link {
  opacity: 0;
  transform: translateY(20px);
  animation: fadeInUp 0.5s ease-out 0.2s forwards;
}

.features {
  animation-delay: 0.3s;
}

@keyframes fadeInUp {
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Плавные переходы для всех элементов */
.telegram-only-layout,
.content-card,
.info-card,
.features,
.title,
.subtitle,
.info-card h2,
.features h3,
.instructions,
.features-list,
.footer-text {
  transition: all 0.3s ease;
}
</style>