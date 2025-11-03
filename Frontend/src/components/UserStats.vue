<template>
  <div v-if="isAdmin" class="user-stats">
    <div class="stats-card">
      <h3>📊 {{ t('user_stats.title') }}</h3>
      <div class="stats-content">
        <div class="stat-item">
          <span class="stat-label">👥 {{ t('user_stats.total_users') }}</span>
          <span class="stat-value">{{ totalUsers }}</span>
        </div>
        <div class="stat-item">
          <span class="stat-label">👑 {{ t('user_stats.admins') }}</span>
          <span class="stat-value">{{ adminUsers }}</span>
        </div>
        <div v-for="(count, role) in roleStats" :key="role" class="stat-item">
          <span class="stat-label">👤 {{ getRoleName(role) }}</span>
          <span class="stat-value">{{ count }}</span>
        </div>
      </div>
      <div class="stats-actions">
        <ThemedButton @click="fetchStats" class="refresh-btn" :disabled="loading">
          <span v-if="loading">⏳ {{ t('user_stats.loading') }}</span>
          <span v-else>🔄 {{ t('user_stats.refresh') }}</span>
        </ThemedButton>
        <ThemedButtonLined buttonType="secondary" @click="viewAllUsers" class="view-all-btn">
          👥 {{ t('user_stats.view_all') }}
        </ThemedButtonLined>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useTelegramWebAppSingleton } from '../telegram/composables/useTelegramWebAppSingleton'
import { getUserInfoFromToken } from '../telegram/auth/user'
import { useApi } from '../telegram/composables/useApi'
import { useLocalization } from '@/locales/index.js'
import ThemedButton from './ThemedButton.vue'
import ThemedButtonLined from './ThemedButtonLined.vue'

const { t } = useLocalization()

const router = useRouter()
const { sendAuthToServer, jwtToken } = useTelegramWebAppSingleton()
const { apiGet } = useApi()

// Получаем информацию о пользователе из токена
const userInfo = computed(() => getUserInfoFromToken())
const isAdmin = computed(() => userInfo.value?.isAdmin || false)

const totalUsers = ref(0)
const adminUsers = ref(0)
const roleStats = ref({})
const loading = ref(false)

// Функция для получения названия роли по числовому значению
const getRoleName = (role) => {
  switch (parseInt(role)) {
    case 1:
      return t('users_view.role_1')
    case 0:
      return t('users_view.role_0')
    default:
      return t('users_view.user')
  }
}

// Переход к просмотру всех пользователей
const viewAllUsers = () => {
  router.push('/users')
}

// Получаем статистику при монтировании компонента
const fetchStats = async (showErrors = true) => {
  try {
    loading.value = true

    // Запрос к бэкенду для получения статистики через новую composable функцию
    const response = await apiGet('users/stats')

    if (response.ok) {
      const data = await response.json()
      totalUsers.value = data.total_users || 0
      adminUsers.value = data.admin_users || 0
      roleStats.value = data.role_stats || {}
    } else {
      console.error('Ошибка при получении статистики:', response.status)
      // Показываем уведомление об ошибке только если это ручное обновление
      if (showErrors) {
        alert(t('user_stats.load_error'))
      }
    }
  } catch (error) {
    console.error('Ошибка при запросе статистики:', error)
    // Показываем уведомление об ошибке только если это ручное обновление
    if (showErrors) {
      alert(t('user_stats.load_error_general'))
    }
  } finally {
    loading.value = false
  }
}

// Получаем статистику при монтировании компонента
onMounted(() => {
  if (isAdmin.value) {
    // Небольшая задержка для гарантии, что токен аутентификации уже доступен
    setTimeout(() => {
      fetchStats(false) // Не показываем ошибки при автоматической загрузке
    }, 100)
  }
})
</script>

<style scoped>
.user-stats {
  margin-bottom: 16px;
}

.stats-card {
  background: var(--tg-theme-secondary-bg-color, white);
  padding: 20px;
  border-radius: 16px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  border: 1px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
  transition: all 0.3s ease;
}

.stats-card h3 {
  margin-top: 0;
  margin-bottom: 16px;
  color: var(--tg-theme-text-color, #000000);
  text-align: center;
  font-size: 1.25rem;
  font-weight: 600;
  transition: color 0.3s ease;
}

.stats-content {
  margin: 16px 0;
}

.stat-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  padding: 16px;
  background: var(--tg-theme-bg-color, #ffffff);
  border-radius: 12px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}

.stat-item:last-child {
  margin-bottom: 0;
}

.stat-label {
  color: var(--tg-theme-text-color, #000000);
  font-size: 1rem;
  transition: color 0.3s ease;
}

.stat-value {
  font-weight: 700;
  font-size: 1.25rem;
  color: var(--tg-theme-button-color, #667eea);
  background: var(--tg-theme-secondary-bg-color, #f0f0f0);
  padding: 6px 16px;
  border-radius: 20px;
  transition: all 0.3s ease;
}

.stats-actions {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-top: 20px;
}

.refresh-btn,
.view-all-btn {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

/* Dark theme adjustments */
:global(.tg-theme-dark) .stats-card {
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
}

:global(.tg-theme-dark) .stat-item {
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.2);
}

:global(.tg-theme-dark) .view-all-btn {
  /* Dark theme styles now handled by ThemedButton component */
}

/* Responsive adjustments */
@media (max-width: 768px) {


  .stat-item {
    padding: 12px;
  }

  .refresh-btn,
  .view-all-btn {
    padding: 10px 16px;
  }
}
</style>