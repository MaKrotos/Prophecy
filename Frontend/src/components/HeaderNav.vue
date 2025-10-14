<script setup>
import { useRouter } from 'vue-router'

defineProps({
  title: String
})

const router = useRouter()

const showRules = () => {
  router.push('/rules')
}
</script>

<template>
  <header class="header-nav">
    <div class="header-content">
      <h1 class="header-title">{{ title }}</h1>
      <div class="header-actions">
        <button class="icon-button" @click="showRules">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor">
            <path
              d="M11 18h2v-2h-2v2zm1-16C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4h2c0-1.1.9-2 2-2s2 .9 2 2c0 2-3 1.75-3 5h2c0-2.25 3-2.5 3-5 0-2.21-1.79-4-4-4z" />
          </svg>
        </button>
      </div>
    </div>
  </header>
</template>

<style scoped>
.header-nav {
  background: var(--tg-theme-button-color, linear-gradient(135deg, #667eea 0%, #764ba2 100%));
  color: var(--tg-theme-button-text-color, white);
  padding: 16px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  position: sticky;
  top: 0;
  z-index: 100;
  transition: all 0.3s ease;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-title {
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--tg-theme-button-text-color, white);
  transition: color 0.3s ease;
}

.header-actions {
  display: flex;
  gap: 12px;
}

.icon-button {
  background: none;
  border: none;
  color: var(--tg-theme-button-text-color, white);
  padding: 8px;
  border-radius: 50%;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.icon-button:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

/* Стили для темной темы Telegram */
:global(.tg-theme-dark) .header-nav {
  background: var(--tg-theme-button-color, #3e6d9c);
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
}

:global(.tg-theme-dark) .icon-button:hover {
  background-color: rgba(255, 255, 255, 0.15);
}

/* Стили для светлой темы Telegram */
:global(.tg-theme-light) .header-nav {
  background: var(--tg-theme-button-color, #667eea);
}

/* Альтернативный вариант с градиентом, если цвет кнопки не задан */
.header-nav {
  background: var(--tg-theme-button-color,
      linear-gradient(135deg, #667eea 0%, #764ba2 100%));
}

/* Улучшенная анимация для кнопок */
.icon-button {
  position: relative;
  overflow: hidden;
}

.icon-button::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: currentColor;
  opacity: 0;
  border-radius: 50%;
  transform: scale(0.8);
  transition: all 0.3s ease;
}

.icon-button:hover::before {
  opacity: 0.1;
  transform: scale(1);
}

/* Адаптивность для мобильных устройств */
@media (max-width: 480px) {
  .header-nav {
    padding: 12px 16px;
  }

  .header-title {
    font-size: 1.1rem;
  }

  .icon-button {
    padding: 6px;
  }

  .icon-button svg {
    width: 18px;
    height: 18px;
  }
}

@media (max-width: 360px) {
  .header-nav {
    padding: 10px 12px;
  }

  .header-title {
    font-size: 1rem;
  }

  .header-actions {
    gap: 8px;
  }
}

/* Плавные переходы для всех элементов */
.header-nav,
.header-title,
.icon-button {
  transition: all 0.3s ease;
}

/* Дополнительные эффекты для лучшей интеграции с Telegram */
.header-nav {
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
}

/* Контрастность текста для лучшей читаемости */
.header-title {
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

:global(.tg-theme-dark) .header-title {
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
}
</style>