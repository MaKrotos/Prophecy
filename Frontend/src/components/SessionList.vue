<template>
  <div class="session-list-component">
    <h3 class="section-title">🎮 {{ t('home_view.sessions.title') }}</h3>

    <div v-if="loading" class="loading">
      {{ t('home_view.sessions.loading') }}
    </div>

    <div v-else-if="sessions.length === 0" class="no-sessions">
      {{ t('home_view.sessions.no_sessions') }}
    </div>

    <div v-else class="sessions-grid">
      <AnimatedCard v-for="(session, index) in sessions" :key="session.id" :index="index" :animation-delay="0.1"
        custom-class="session-card" @click="viewSession(session)">
        <div class="session-header">
          <h4 class="session-name">{{ session.name }}</h4>
        </div>

        <div class="session-details">
          <div class="detail-item">
            <span class="detail-label">{{ t('home_view.sessions.architect') }}:</span>
            <span class="detail-value">{{ session.architect_name }}</span>
          </div>
        </div>
      </AnimatedCard>
    </div>

    <button v-if="isArchitect" class="create-session-button" @click="goToCreateSession">
      + {{ t('home_view.sessions.create_session') }}
    </button>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useApi } from '../telegram/composables/useApi'
import AnimatedCard from './AnimatedCard.vue'
import { useLocalization } from '@/locales/index.js'
import { getUserInfoFromToken } from '../telegram/auth/user'

const { t } = useLocalization()
const router = useRouter()
const { apiGet } = useApi()

const sessions = ref([])
const loading = ref(false)
const userInfo = ref(null)

// Проверяем, является ли пользователь архитектором
const isArchitect = computed(() => {
  if (!userInfo.value) return false
  return userInfo.value.role && userInfo.value.role.String === 'Архитектор'
})

// Загрузка сессий
const loadSessions = async () => {
  try {
    loading.value = true
    const response = await apiGet('sessions')

    if (response.ok) {
      const data = await response.json()
      // Проверяем, что data - массив
      if (Array.isArray(data)) {
        // Ограничиваем количество отображаемых сессий до 5
        sessions.value = data.slice(0, 5)
      }
    }
  } catch (error) {
    console.error('Ошибка при запросе сессий:', error)
  } finally {
    loading.value = false
  }
}

// Просмотр сессии
const viewSession = (session) => {
  router.push(`/sessions/${session.id}`)
}

// Переход к созданию сессии
const goToCreateSession = () => {
  router.push('/sessions/create')
}

// Первоначальная загрузка
onMounted(() => {
  // Получаем информацию о пользователе из JWT токена
  userInfo.value = getUserInfoFromToken()
  loadSessions()
})
</script>

<style scoped>
.session-list-component {
  margin-top: 24px;
}

.section-title {
  color: var(--tg-theme-text-color, #000000);
  font-size: 1.2rem;
  font-weight: 600;
  margin-bottom: 16px;
  transition: color 0.3s ease;
}

.loading,
.no-sessions {
  color: var(--tg-theme-hint-color, #666666);
  font-size: 1rem;
  text-align: center;
  padding: 20px;
  transition: color 0.3s ease;
}

.sessions-grid {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 16px;
}

.session-card {
  cursor: pointer;
  background: var(--tg-theme-secondary-bg-color, white);
  border-radius: 12px;
  padding: 16px;
  box-shadow: 0 1px 6px rgba(0, 0, 0, 0.08);
  border: 1px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
  transition: all 0.3s ease;
}

.session-header {
  margin-bottom: 12px;
}

.session-name {
  color: var(--tg-theme-text-color, #333333);
  font-size: 1rem;
  font-weight: 600;
  margin: 0;
  transition: color 0.3s ease;
}

.session-details {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.detail-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 0.85rem;
}

.detail-label {
  color: var(--tg-theme-hint-color, #666666);
  transition: color 0.3s ease;
}

.detail-value {
  color: var(--tg-theme-text-color, #333333);
  font-weight: 500;
  transition: color 0.3s ease;
}

.create-session-button {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
  border: none;
  border-radius: 12px;
  padding: 12px 20px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  width: 100%;
}

.create-session-button:hover {
  opacity: 0.9;
  transform: translateY(-1px);
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .session-card {
    padding: 12px;
  }

  .session-name {
    font-size: 0.95rem;
  }
}
</style>