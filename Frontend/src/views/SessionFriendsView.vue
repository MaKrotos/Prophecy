<template>
  <div class="page">
    <div class="header-section">
      <h2 class="page-title">👥 {{ t('session_friends_view.title') }}</h2>
      <p class="page-description">{{ t('session_friends_view.description') }}</p>
    </div>
    
    <div class="friends-list">
      <div v-if="loading" class="loading">{{ t('session_friends_view.loading') }}</div>
      <div v-else-if="friends.length === 0" class="no-friends">{{ t('session_friends_view.no_friends') }}</div>
      <div v-else class="friends-grid">
        <ThemedCard 
          v-for="friend in friends" 
          :key="friend.id" 
          card-type="default" 
          class="friend-card"
        >
          <div class="friend-avatar">
            <div class="avatar-placeholder">{{ friend.name.charAt(0) }}</div>
          </div>
          <div class="friend-info">
            <h3 class="friend-name">{{ friend.name }}</h3>
            <p class="friend-status" :class="{ online: friend.online }">
              {{ friend.online ? t('session_friends_view.online') : t('session_friends_view.offline') }}
            </p>
          </div>
        </ThemedCard>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useLocalization } from '@/locales/index.js'
import ThemedCard from '../components/ThemedCard.vue'

const { t } = useLocalization()

const friends = ref([])
const loading = ref(false)

// Заглушка для списка друзей
const loadFriends = () => {
  loading.value = true
  
  // Имитация загрузки данных
  setTimeout(() => {
    friends.value = [
      { id: 1, name: 'Алексей', online: true },
      { id: 2, name: 'Мария', online: true },
      { id: 3, name: 'Дмитрий', online: false },
      { id: 4, name: 'Елена', online: true },
      { id: 5, name: 'Иван', online: false },
      { id: 6, name: 'Ольга', online: true }
    ]
    loading.value = false
  }, 1000)
}

onMounted(() => {
  loadFriends()
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

.friend-card {
  background: var(--tg-theme-secondary-bg-color, white);
  border-radius: 12px;
  padding: 16px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  display: flex;
  align-items: center;
  gap: 16px;
  transition: all 0.3s ease;
}

.friend-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
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