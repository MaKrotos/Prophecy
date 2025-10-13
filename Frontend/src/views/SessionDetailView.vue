<template>
  <div class="page">
    <div v-if="session" class="session-detail">
      <div class="session-header">
        <h2 class="session-title">{{ session.name }}</h2>
        <div class="session-actions">
          <button 
            v-if="canManageSession" 
            class="edit-button"
            @click="editSession"
          >
            ✏️ {{ t('session_detail_view.edit') }}
          </button>
          <button 
            v-if="canManageSession" 
            class="delete-button"
            @click="deleteSession"
          >
            🗑️ {{ t('session_detail_view.delete') }}
          </button>
        </div>
      </div>
      
      <p class="session-description">{{ session.description || t('session_detail_view.no_description') }}</p>
      
      <div class="session-info">
        <div class="info-item">
          <span class="info-label">{{ t('session_detail_view.architect') }}:</span>
          <span class="info-value">{{ session.architect_name }}</span>
        </div>
        
        <div class="info-item">
          <span class="info-label">{{ t('session_detail_view.created') }}:</span>
          <span class="info-value">{{ formatDate(session.created_at) }}</span>
        </div>
        
        <div class="info-item">
          <span class="info-label">{{ t('session_detail_view.updated') }}:</span>
          <span class="info-value">{{ formatDate(session.updated_at) }}</span>
        </div>
      </div>
      
      <div class="players-section">
        <div class="section-header">
          <h3 class="section-title">👥 {{ t('session_detail_view.players') }}</h3>
          <button 
            v-if="canManageSession" 
            class="add-player-button"
            @click="showAddPlayerDialog"
          >
            + {{ t('session_detail_view.add_player') }}
          </button>
        </div>
        
        <div v-if="loadingPlayers" class="loading">{{ t('session_detail_view.loading_players') }}</div>
        <div v-else-if="players.length === 0" class="no-players">{{ t('session_detail_view.no_players') }}</div>
        <div v-else class="players-list">
          <div 
            v-for="player in players" 
            :key="player.id" 
            class="player-card"
          >
            <div class="player-info">
              <span class="player-name">{{ player.generated_name }}</span>
              <span v-if="player.username" class="player-username">@{{ player.username }}</span>
            </div>
            <button 
              v-if="canManageSession" 
              class="remove-player-button"
              @click="removePlayer(player)"
            >
              🗑️
            </button>
          </div>
        </div>
      </div>
      
    </div>
    
    <div v-else-if="loading" class="loading">{{ t('session_detail_view.loading') }}</div>
    <div v-else class="error">{{ t('session_detail_view.load_error') }}</div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useApi } from '../telegram/composables/useApi'
import { useLocalization } from '@/locales/index.js'
import { getUserInfoFromToken } from '../telegram/auth/user'

const { t } = useLocalization()
const route = useRoute()
const router = useRouter()
const { apiGet, apiPut, apiDelete, apiPost } = useApi()

const session = ref(null)
const players = ref([])
const loading = ref(false)
const loadingPlayers = ref(false)
const userInfo = ref(null)

// Проверяем, может ли пользователь управлять сессией
const canManageSession = computed(() => {
  if (!userInfo.value || !session.value) return false
  // Администраторы и архитектор, создавший сессию, могут управлять сессией
  return userInfo.value.is_admin || 
         (userInfo.value.role && userInfo.value.role.String === 'Архитектор' && 
          userInfo.value.id === session.value.architect_id)
})

// Форматирование даты
const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('ru-RU', {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

// Загрузка сессии
const loadSession = async () => {
  try {
    loading.value = true
    const response = await apiGet(`sessions/${route.params.id}`)
    
    if (response.ok) {
      const data = await response.json()
      session.value = data
      // Загружаем связанные данные
      loadPlayers()
    } else {
      console.error('Ошибка при загрузке сессии:', response.status)
      if (window.Telegram && window.Telegram.WebApp) {
        window.Telegram.WebApp.showAlert(t('session_detail_view.load_error'))
      }
    }
  } catch (error) {
    console.error('Ошибка при запросе сессии:', error)
    if (window.Telegram && window.Telegram.WebApp) {
      window.Telegram.WebApp.showAlert(t('session_detail_view.load_error_general'))
    }
  } finally {
    loading.value = false
  }
}

// Загрузка игроков
const loadPlayers = async () => {
  try {
    loadingPlayers.value = true
    const response = await apiGet(`sessions/${route.params.id}/players`)
    
    if (response.ok) {
      const data = await response.json()
      players.value = Array.isArray(data) ? data : []
    } else {
      console.error('Ошибка при загрузке игроков:', response.status)
    }
  } catch (error) {
    console.error('Ошибка при запросе игроков:', error)
  } finally {
    loadingPlayers.value = false
  }
}

// Редактирование сессии
const editSession = () => {
  // Пока просто показываем сообщение
  if (window.Telegram && window.Telegram.WebApp) {
    window.Telegram.WebApp.showAlert(t('session_detail_view.edit_not_implemented'))
  }
}

// Удаление сессии
const deleteSession = () => {
  if (window.Telegram && window.Telegram.WebApp) {
    window.Telegram.WebApp.showConfirm(
      t('session_detail_view.delete_confirm'),
      async (confirmed) => {
        if (confirmed) {
          try {
            const response = await apiDelete(`sessions/${route.params.id}`)
            
            if (response.ok) {
              if (window.Telegram && window.Telegram.WebApp) {
                window.Telegram.WebApp.showAlert(t('session_detail_view.delete_success'))
              }
              router.push('/sessions')
            } else {
              if (window.Telegram && window.Telegram.WebApp) {
                window.Telegram.WebApp.showAlert(t('session_detail_view.delete_error'))
              }
            }
          } catch (error) {
            console.error('Ошибка при удалении сессии:', error)
            if (window.Telegram && window.Telegram.WebApp) {
              window.Telegram.WebApp.showAlert(t('session_detail_view.delete_error_general'))
            }
          }
        }
      }
    )
  }
}

// Показ диалога добавления игрока
const showAddPlayerDialog = () => {
  if (window.Telegram && window.Telegram.WebApp) {
    window.Telegram.WebApp.showAlert(t('session_detail_view.add_player_not_implemented'))
  }
}

// Удаление игрока из сессии
const removePlayer = (player) => {
  if (window.Telegram && window.Telegram.WebApp) {
    window.Telegram.WebApp.showConfirm(
      `${t('session_detail_view.remove_player_confirm')} ${player.generated_name}?`,
      async (confirmed) => {
        if (confirmed) {
          try {
            const response = await apiDelete(`sessions/${route.params.id}/players`, {
              body: JSON.stringify({ user_id: player.id })
            })
            
            if (response.ok) {
              // Обновляем список игроков
              loadPlayers()
              if (window.Telegram && window.Telegram.WebApp) {
                window.Telegram.WebApp.showAlert(t('session_detail_view.remove_player_success'))
              }
            } else {
              if (window.Telegram && window.Telegram.WebApp) {
                window.Telegram.WebApp.showAlert(t('session_detail_view.remove_player_error'))
              }
            }
          } catch (error) {
            console.error('Ошибка при удалении игрока:', error)
            if (window.Telegram && window.Telegram.WebApp) {
              window.Telegram.WebApp.showAlert(t('session_detail_view.remove_player_error_general'))
            }
          }
        }
      }
    )
  }
}

// Первоначальная загрузка
onMounted(() => {
  // Получаем информацию о пользователе из JWT токена
  userInfo.value = getUserInfoFromToken()
  loadSession()
})
</script>

<style scoped>
.page {
  padding: 16px;
  background-color: var(--tg-theme-bg-color, #f5f5f5);
  transition: background-color 0.3s ease;
}

.session-detail {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.session-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 12px;
  padding-bottom: 16px;
  border-bottom: 1px solid var(--tg-theme-hint-color, #e0e0e0);
}

.session-title {
  color: var(--tg-theme-text-color, #000000);
  font-size: 1.5rem;
  font-weight: 700;
  margin: 0;
  transition: color 0.3s ease;
}

.session-actions {
  display: flex;
  gap: 8px;
}

.edit-button,
.delete-button {
  padding: 8px 12px;
  border: none;
  border-radius: 8px;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  gap: 4px;
}

.edit-button {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
}

.delete-button {
  background: #ff4757;
  color: white;
}

.edit-button:hover,
.delete-button:hover {
  opacity: 0.9;
  transform: translateY(-1px);
}

.session-description {
  color: var(--tg-theme-text-color, #333333);
  font-size: 1rem;
  line-height: 1.5;
  transition: color 0.3s ease;
}

.session-info {
  display: flex;
  flex-direction: column;
  gap: 12px;
  background: var(--tg-theme-secondary-bg-color, white);
  padding: 16px;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.info-label {
  color: var(--tg-theme-hint-color, #666666);
  font-weight: 500;
  transition: color 0.3s ease;
}

.info-value {
  color: var(--tg-theme-text-color, #333333);
  font-weight: 500;
  transition: color 0.3s ease;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  flex-wrap: wrap;
  gap: 12px;
}

.section-title {
  color: var(--tg-theme-text-color, #000000);
  font-size: 1.25rem;
  font-weight: 600;
  margin: 0;
  transition: color 0.3s ease;
}

.add-player-button {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
  border: none;
  border-radius: 8px;
  padding: 8px 12px;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  gap: 4px;
}

.add-player-button:hover {
  opacity: 0.9;
  transform: translateY(-1px);
}

.players-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.player-card {
  background: var(--tg-theme-secondary-bg-color, white);
  border-radius: 12px;
  padding: 16px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  display: flex;
  justify-content: space-between;
  align-items: center;
  transition: all 0.3s ease;
}

.player-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.player-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.player-name {
  color: var(--tg-theme-text-color, #333333);
  font-weight: 600;
  transition: color 0.3s ease;
}

.player-username {
  color: var(--tg-theme-hint-color, #666666);
  font-size: 0.9rem;
  transition: color 0.3s ease;
}

.remove-player-button {
  background: none;
  border: none;
  font-size: 1rem;
  cursor: pointer;
  padding: 4px;
  border-radius: 4px;
  transition: background-color 0.3s ease;
}

.remove-player-button:hover {
  background-color: rgba(0, 0, 0, 0.05);
}

.loading,
.error,
.no-players {
  text-align: center;
  padding: 24px;
  color: var(--tg-theme-hint-color, #666666);
  transition: color 0.3s ease;
}

.error {
  color: #ff4757;
}

/* Плавные переходы для всех элементов */
.page,
.session-title,
.session-description,
.info-label,
.info-value,
.section-title,
.player-name,
.player-username {
  transition: all 0.3s ease;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .session-header {
    flex-direction: column;
    align-items: flex-start;
  }
  
  .session-title {
    font-size: 1.3rem;
  }
  
  .session-actions {
    width: 100%;
    justify-content: space-between;
  }
  
  .edit-button,
  .delete-button {
    flex: 1;
    justify-content: center;
  }
  
  .section-header {
    flex-direction: column;
    align-items: flex-start;
  }
  
  .add-player-button {
    width: 100%;
    justify-content: center;
  }
  
  .player-card {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }
  
  .player-info {
    width: 100%;
  }
  
  .remove-player-button {
    align-self: flex-end;
  }
}
</style>