<template>
  <div class="page">
    <h2 class="page-title">👥 Все пользователи</h2>
    <p class="page-description">Список всех пользователей системы</p>

    <div class="users-list" ref="scrollContainer" @scroll="handleScroll">
      <div
        v-for="user in users"
        :key="user.id"
        class="user-card"
      >
        <div class="user-header">
          <div class="user-info">
            <h3 class="user-name">{{ user.generated_name }}</h3>
          </div>
          <div class="user-badge" :class="{ 'admin-badge': user.is_admin }">
            {{ user.is_admin ? 'Админ' : 'Пользователь' }}
          </div>
        </div>

        <div class="user-details">
          <div class="detail-item">
            <span class="detail-label">ID:</span>
            <span class="detail-value">{{ user.id }}</span>
          </div>
          
          <div class="detail-item" v-if="user.username">
            <span class="detail-label">Username:</span>
            <span class="detail-value">@{{ user.username }}</span>
          </div>
          
          <div class="detail-item">
            <span class="detail-label">Регистрация:</span>
            <span class="detail-value">{{ formatDate(user.created_at) }}</span>
          </div>
        </div>
      </div>

      <div v-if="loading" class="loading-indicator">
        <div class="spinner"></div>
        <p>Загрузка пользователей...</p>
      </div>

      <div v-if="noMoreUsers && !loading" class="no-more-users">
        <p>Все пользователи загружены</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useApi } from '../telegram/composables/useApi'

const { apiGet } = useApi()

const users = ref([])
const loading = ref(false)
const offset = ref(0)
const limit = ref(20)
const noMoreUsers = ref(false)
const scrollContainer = ref(null)

// Форматирование даты
const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('ru-RU', {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric'
  })
}

// Загрузка пользователей
const loadUsers = async () => {
  if (loading.value || noMoreUsers.value) return
  
  try {
    loading.value = true
    const response = await apiGet(`users?limit=${limit.value}&offset=${offset.value}`)
    
    if (response.ok) {
      const data = await response.json()
      if (data.length > 0) {
        users.value = [...users.value, ...data]
        offset.value += data.length
      } else {
        noMoreUsers.value = true
      }
    } else {
      console.error('Ошибка при загрузке пользователей:', response.status)
      alert('Не удалось загрузить список пользователей')
    }
  } catch (error) {
    console.error('Ошибка при запросе пользователей:', error)
    alert('Произошла ошибка при загрузке списка пользователей')
  } finally {
    loading.value = false
  }
}

// Обработка скролла
const handleScroll = () => {
  const container = scrollContainer.value
  if (!container) return

  const { scrollTop, scrollHeight, clientHeight } = container
  const scrollPercentage = (scrollTop + clientHeight) / scrollHeight

  // Загружаем больше пользователей, когда пользователь прокрутил 80% контента
  if (scrollPercentage > 0.8) {
    loadUsers()
  }
}

// Первоначальная загрузка
onMounted(() => {
  loadUsers()
})
</script>

<style scoped>
.page {
  padding: 16px;
  min-height: 100vh;
  background-color: var(--tg-theme-bg-color, #f5f5f5);
  transition: background-color 0.3s ease;
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
  margin-bottom: 24px;
  transition: color 0.3s ease;
}

.users-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
  max-height: calc(100vh - 150px);
  overflow-y: auto;
  padding-right: 8px;
}

.user-card {
  background: var(--tg-theme-secondary-bg-color, white);
  border-radius: 12px;
  padding: 16px;
  box-shadow: 0 1px 6px rgba(0, 0, 0, 0.08);
  border: 1px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
  transition: all 0.3s ease;
}

.user-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  flex-wrap: wrap;
  gap: 8px;
}

.user-info {
  flex: 1;
}

.user-name {
  color: var(--tg-theme-text-color, #333333);
  font-size: 1.1rem;
  font-weight: 600;
  margin: 0;
  transition: color 0.3s ease;
}

.user-badge {
  background: var(--tg-theme-hint-color, #e0e0e0);
  color: var(--tg-theme-text-color, #333333);
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 0.75rem;
  font-weight: 500;
}

.user-badge.admin-badge {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
}

.user-details {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.detail-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 0.9rem;
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

.loading-indicator {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 16px;
  color: var(--tg-theme-hint-color, #666666);
}

.spinner {
  width: 24px;
  height: 24px;
  border: 2px solid rgba(0, 0, 0, 0.1);
  border-top: 2px solid var(--tg-theme-button-color, #667eea);
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 8px;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.no-more-users {
  text-align: center;
  padding: 16px;
  color: var(--tg-theme-hint-color, #666666);
  font-style: italic;
  font-size: 0.9rem;
}

/* Dark theme adjustments */
:global(.tg-theme-dark) .user-card {
  box-shadow: 0 1px 6px rgba(0, 0, 0, 0.2);
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .user-header {
    flex-direction: column;
    align-items: flex-start;
  }
  
  .user-name {
    font-size: 1rem;
  }
  
  .user-card {
    padding: 12px;
  }
  
  .users-list {
    max-height: calc(100vh - 130px);
  }
}
</style>