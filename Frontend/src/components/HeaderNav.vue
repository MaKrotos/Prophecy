<script setup>
import { useRouter } from 'vue-router'
import ThemedButton from '@/components/ThemedButton.vue'

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
        <ThemedButton buttonType="icon" class="icon-button" @click="showRules">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor">
            <path
              d="M11 18h2v-2h-2v2zm1-16C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm0-14c-2.21 0-4 1.79-4 4h2c0-1.1.9-2 2-2s2 .9 2 2c0 2-3 1.75-3 5h2c0-2.25 3-2.5 3-5 0-2.21-1.79-4-4-4z" />
          </svg>
        </ThemedButton>
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