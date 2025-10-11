<template>
  <div v-if="isAdmin" class="user-stats">
    <div class="stats-card">
      <h3>📊 Статистика пользователей</h3>
      <div class="stats-content">
        <div class="stat-item">
          <span class="stat-label">👥 Всего пользователей</span>
          <span class="stat-value">{{ totalUsers }}</span>
        </div>
        <div class="stat-item">
          <span class="stat-label">👑 Администраторов</span>
          <span class="stat-value">{{ adminUsers }}</span>
        </div>
      </div>
      <div class="stats-actions">
        <button @click="fetchStats" class="refresh-btn" :disabled="loading">
          <span v-if="loading">⏳ Загрузка...</span>
          <span v-else>🔄 Обновить статистику</span>
        </button>
        <button @click="viewAllUsers" class="view-all-btn">
          👥 Просмотр всех пользователей
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useTelegramWebApp } from '../telegram/composables/useTelegramWebApp'
import { getUserInfoFromToken } from '../telegram/auth/user'
import { useApi } from '../telegram/composables/useApi'

const router = useRouter()
const { sendAuthToServer, jwtToken } = useTelegramWebApp()
const { apiGet } = useApi()

// Получаем информацию о пользователе из токена
const userInfo = computed(() => getUserInfoFromToken())
const isAdmin = computed(() => userInfo.value?.isAdmin || false)

const totalUsers = ref(0)
const adminUsers = ref(0)
const loading = ref(false)

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
    } else {
      console.error('Ошибка при получении статистики:', response.status)
      // Показываем уведомление об ошибке только если это ручное обновление
      if (showErrors) {
        alert('Не удалось загрузить статистику пользователей')
      }
    }
  } catch (error) {
    console.error('Ошибка при запросе статистики:', error)
    // Показываем уведомление об ошибке только если это ручное обновление
    if (showErrors) {
      alert('Произошла ошибка при загрузке статистики пользователей')
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
  margin: 16px;
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

.refresh-btn, .view-all-btn {
  background-color: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
  border: none;
  padding: 12px 20px;
  border-radius: 12px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  width: 100%;
  transition: all 0.3s ease;
  display: flex;
  justify-content: center;
  align-items: center;
}

.view-all-btn {
  background-color: var(--tg-theme-secondary-bg-color, #f0f0f0);
  color: var(--tg-theme-text-color, #333333);
  border: 1px solid var(--tg-theme-hint-color, #cccccc);
}

.refresh-btn:hover:not(:disabled), .view-all-btn:hover {
  opacity: 0.9;
  transform: translateY(-1px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.refresh-btn:active:not(:disabled), .view-all-btn:active {
  transform: translateY(0);
}

.refresh-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
  transform: none;
}

/* Dark theme adjustments */
:global(.tg-theme-dark) .stats-card {
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
}

:global(.tg-theme-dark) .stat-item {
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.2);
}

:global(.tg-theme-dark) .view-all-btn {
  background: var(--tg-theme-bg-color, #2a3b4d);
  border-color: var(--tg-theme-hint-color, #4a5b6d);
}

/* Responsive adjustments */
@media (max-width: 768px) {
  
  
  .stats-card {
    padding: 16px;
  }
  
  .stat-item {
    padding: 12px;
  }
  
  .refresh-btn, .view-all-btn {
    padding: 10px 16px;
  }
}
</style>