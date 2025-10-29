<template>
  <div class="page">
    <div class="header-section">
      <h2 class="page-title">👥 {{ t('session_friends_view.title') }}</h2>
      <p class="page-description">{{ t('session_friends_view.description') }}</p>
      <button class="scan-qr-btn" @click="scanQRCode">
        {{ t('session_friends_view.scan_qr') }}
      </button>
    </div>
    
    <AnimatedCardList
      :items="friends"
      :loading="loading"
      :no-more-items="true"
      key-field="id"
      card-class="friend-card"
      :animation-delay="0.1"
      :loading-text="t('session_friends_view.loading')"
      :no-more-items-text="t('session_friends_view.no_friends')"
    >
      <template #card="{ item: friend }">
        <div class="friend-content">
          <div class="friend-avatar">
            <div class="avatar-placeholder">{{ friend.name.charAt(0) }}</div>
          </div>
          <div class="friend-info">
            <h3 class="friend-name">{{ friend.name }}</h3>
            <p class="friend-status" :class="{ online: friend.online }">
              {{ friend.online ? t('session_friends_view.online') : t('session_friends_view.offline') }}
            </p>
          </div>
          <div class="friend-actions">
            <button class="remove-friend-btn" @click="removeFriend(friend.id)">
              {{ t('session_friends_view.remove_friend') }}
            </button>
          </div>
        </div>
      </template>
    </AnimatedCardList>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useLocalization } from '@/locales/index.js'
import AnimatedCardList from '../components/AnimatedCardList.vue'
import { useApi } from '@/telegram/composables/useApi.js'
import { useRoute, useRouter } from 'vue-router'

const { t } = useLocalization()
const { apiGet, apiPost, apiDelete } = useApi()
const route = useRoute()
const router = useRouter()

const friends = ref([])
const loading = ref(false)

// Функция для сканирования QR кода
const scanQRCode = () => {
  // Получаем ID сессии из параметров маршрута
  const sessionId = route.params.id
  
  // Если есть ID сессии, переходим на страницу сканирования QR кода с этим ID
  if (sessionId) {
    router.push(`/session/qr-scanner`)
  } else {
    // Если нет ID сессии, переходим на общую страницу сканирования
    router.push('/scan')
  }
}

// Получение списка друзей с бэкенда
const loadFriends = async () => {
  loading.value = true
  
  try {
    // Получаем ID сессии из параметров маршрута
    const sessionId = route.params.id
    
    // Если нет ID сессии, очищаем список друзей и выходим
    if (!sessionId) {
      friends.value = []
      loading.value = false
      return
    }
    
    // Запрашиваем данные у бэкенда
    const response = await apiGet(`/sessions/${sessionId}/friends`)
    const data = await response.json()
    
    // Обновляем список друзей
    friends.value = data.map(friend => ({
      id: friend.id,
      name: friend.friend_name,
      online: friend.online || false
    }))
  } catch (error) {
    console.error('Failed to load friends:', error)
    // В случае ошибки очищаем список друзей
    friends.value = []
  } finally {
    loading.value = false
  }
}

// Добавление друга в сессию
const addFriend = async (friendId) => {
  try {
    const sessionId = route.params.id
    // Проверяем, что есть ID сессии
    if (!sessionId) {
      console.error('No session ID provided')
      return
    }
    
    await apiPost(`/sessions/${sessionId}/friends?friend_id=${friendId}`, {})
    // После добавления друга перезагружаем список
    await loadFriends()
  } catch (error) {
    console.error('Failed to add friend:', error)
  }
}

// Удаление друга из сессии
const removeFriend = async (friendId) => {
  // Подтверждение удаления
  if (!confirm(t('session_friends_view.confirm_remove_friend'))) {
    return
  }
  
  try {
    const sessionId = route.params.id
    // Проверяем, что есть ID сессии
    if (!sessionId) {
      console.error('No session ID provided')
      return
    }
    
    await apiDelete(`/sessions/${sessionId}/friends?friend_id=${friendId}`)
    // После успешного удаления удаляем друга из списка
    friends.value = friends.value.filter(friend => friend.id !== friendId)
  } catch (error) {
    console.error('Failed to remove friend:', error)
    // В случае ошибки перезагружаем список
    await loadFriends()
  }
}

// Загружаем список друзей при монтировании компонента
onMounted(() => {
  loadFriends()
})

// Следим за изменением параметров маршрута и перезагружаем список друзей
watch(() => route.params.id, (newId, oldId) => {
  if (newId !== oldId) {
    loadFriends()
  }
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

.friends-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.loading,
.no-friends {
  text-align: center;
  padding: 24px;
  color: var(--tg-theme-hint-color, #666666);
  transition: color 0.3s ease;
}

.friends-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 16px;
}

.friend-content {
  display: flex;
  align-items: center;
  gap: 16px;
}

.friend-avatar {
  flex-shrink: 0;
}

.avatar-placeholder {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  background: var(--tg-theme-button-color, #667eea);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.25rem;
  font-weight: 600;
}

.friend-info {
  flex: 1;
  min-width: 0;
}

.friend-name {
  color: var(--tg-theme-text-color, #000000);
  font-size: 1.1rem;
  font-weight: 600;
  margin: 0 0 4px 0;
  transition: color 0.3s ease;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.friend-status {
  color: var(--tg-theme-hint-color, #666666);
  font-size: 0.9rem;
  margin: 0;
  transition: color 0.3s ease;
}

.friend-status.online {
  color: var(--tg-theme-button-color, #667eea);
}

/* Плавные переходы для всех элементов */
.page,
.page-title,
.page-description,
.friend-name,
.friend-status {
  transition: all 0.3s ease;
}

.friend-actions {
  margin-left: auto;
  display: flex;
  align-items: center;
}

.remove-friend-btn {
  background-color: #ff4757;
  color: white;
  border: none;
  border-radius: 4px;
  padding: 6px 12px;
  font-size: 0.8rem;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.remove-friend-btn:hover {
  background-color: #ff2e43;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .page {
    padding: 12px;
  }
  
  .page-title {
    font-size: 1.3rem;
  }
  
  .friends-grid {
    grid-template-columns: 1fr;
  }
  
  .friend-card {
    padding: 12px;
  }
  
  .avatar-placeholder {
    width: 40px;
    height: 40px;
    font-size: 1rem;
  }
  
  .friend-name {
    font-size: 1rem;
  }
}
</style>