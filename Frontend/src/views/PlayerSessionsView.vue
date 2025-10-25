<template>
  <div class="page">
    <div class="header-section">
      <h2 class="page-title">🎮 {{ t('player_sessions_view.title') }}</h2>
      <p class="page-description">{{ t('player_sessions_view.description') }}</p>
    </div>

    <AnimatedCardList
      :items="sessions"
      :loading="loading"
      :no-more-items="!loading && sessions.length === 0"
      key-field="id"
      card-class="session-card"
      :animation-delay="0.1"
      :loading-text="t('player_sessions_view.loading')"
      :no-more-items-text="t('player_sessions_view.no_sessions')">
      
      <template #card="{ item: session }">
        <SessionCard :session="session" @click="viewSession" />
      </template>
    </AnimatedCardList>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useApi } from '../telegram/composables/useApi'
import AnimatedCardList from '../components/AnimatedCardList.vue'
import SessionCard from '../components/SessionCard.vue'
import { useLocalization } from '@/locales/index.js'

const { t } = useLocalization()
const router = useRouter()
const { apiGet } = useApi()

const sessions = ref([])
const loading = ref(false)

// Загрузка сессий игрока
const loadPlayerSessions = async () => {
  try {
    loading.value = true
    const response = await apiGet('players/sessions')
    
    if (response.ok) {
      const data = await response.json()
      // Проверяем, что data - массив
      if (Array.isArray(data)) {
        sessions.value = data
      } else {
        console.warn('Получен неожиданный формат данных:', data)
      }
    } else {
      console.error('Ошибка при загрузке сессий игрока:', response.status)
      if (window.Telegram && window.Telegram.WebApp) {
        window.Telegram.WebApp.showAlert(t('player_sessions_view.load_error'))
      }
    }
  } catch (error) {
    console.error('Ошибка при запросе сессий игрока:', error)
    if (window.Telegram && window.Telegram.WebApp) {
      window.Telegram.WebApp.showAlert(t('player_sessions_view.load_error_general'))
    }
  } finally {
    loading.value = false
  }
}

// Просмотр сессии
const viewSession = (session) => {
  router.push(`/sessions/${session.id}`)
}

// Первоначальная загрузка
onMounted(() => {
  loadPlayerSessions()
})
</script>

<style scoped>
.page {
  padding: 16px;
  background-color: var(--tg-theme-bg-color, #f5f5f5);
  transition: background-color 0.3s ease;
}

.header-section {
  margin-bottom: 24px;
}

.page-title {
  color: var(--tg-theme-text-color, #000000);
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 8px;
  transition: color 0.3s ease;
}

.page-description {
  color: var(--tg-theme-hint-color, #666666);
  font-size: 1rem;
  margin-bottom: 16px;
  transition: color 0.3s ease;
}

/* Плавные переходы для всех элементов */
.page,
.page-title,
.page-description {
  transition: all 0.3s ease;
}
</style>