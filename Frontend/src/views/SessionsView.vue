<template>
  <div class="page">
    <div class="header-section">
      <h2 class="page-title">🎮 {{ t('sessions_view.title') }}</h2>
      <p class="page-description">{{ t('sessions_view.description') }}</p>

      <button v-if="isArchitect" class="create-session-button" @click="goToCreateSession">
        + {{ t('sessions_view.create_session') }}
      </button>
    </div>

    <AnimatedCardList :items="sessions" :loading="loading" :no-more-items="noMoreSessions" key-field="id"
      card-class="session-card" :animation-delay="0.1" :loading-text="t('sessions_view.loading')"
      :no-more-items-text="t('sessions_view.no_more')">
      <template #card="{ item: session }">
        <div class="session-header">
          <h3 class="session-name">{{ session.name }}</h3>
          <div class="session-controls">
            <button class="session-button" @click="viewSession(session)">
              👁️
            </button>
            <button v-if="canManageSession(session)" class="session-button delete-button"
              @click="deleteSession(session)" :title="t('session_detail_view.delete')">
              🗑️
            </button>
          </div>
        </div>

        <div class="session-details">
          <div class="detail-item">
            <span class="detail-label">{{ t('sessions_view.description') }}:</span>
            <span class="detail-value">{{ session.description || t('sessions_view.no_description') }}</span>
          </div>

          <div class="detail-item">
            <span class="detail-label">{{ t('sessions_view.architect') }}:</span>
            <span class="detail-value">{{ session.architect_name }}</span>
          </div>

          <div class="detail-item">
            <span class="detail-label">{{ t('sessions_view.created') }}:</span>
            <span class="detail-value">{{ formatDate(session.created_at) }}</span>
          </div>

          <div class="detail-item">
            <span class="detail-label">{{ t('sessions_view.players') }}:</span>
            <span class="detail-value">{{ session.player_count }}</span>
          </div>

          <!-- Реферальная ссылка для всех пользователей -->
          <div class="detail-item">
            <span class="detail-label">{{ t('sessions_view.referral_link') }}:</span>
            <div class="referral-link-container">
              <span class="detail-value referral-link">{{ getReferralLink(session) }}</span>
              <button class="copy-button" @click="copyReferralLink(session)" :title="t('sessions_view.copy_link')">
                📋
              </button>
            </div>
          </div>
        </div>
      </template>
    </AnimatedCardList>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useApi } from '../telegram/composables/useApi'
import AnimatedCardList from '../components/AnimatedCardList.vue'
import { useLocalization } from '@/locales/index.js'
import { getUserInfoFromToken } from '../telegram/auth/user'

const { t } = useLocalization()
const router = useRouter()
const { apiGet, apiDelete } = useApi()

const sessions = ref([])
const loading = ref(false)
const noMoreSessions = ref(false)
const userInfo = ref(null)

// Проверяем, является ли пользователь архитектором
const isArchitect = computed(() => {
  if (!userInfo.value) return false
  return userInfo.value.role && userInfo.value.role.String === 'Архитектор'
})

// Проверяем, может ли пользователь управлять сессией
const canManageSession = (session) => {
  if (!userInfo.value) return false
  // Администраторы и архитектор, создавший сессию, могут управлять сессией
  return userInfo.value.is_admin ||
    (userInfo.value.role && userInfo.value.role.String === 'Архитектор' &&
      userInfo.value.id === session.architect_id)
}

// Получение реферальной ссылки для сессии
const getReferralLink = (session) => {
  // Используем имя бота из переменных окружения или значение по умолчанию
  const botUsername = import.meta.env.VITE_TELEGRAM_BOT_USERNAME || 'your_bot_username'
  // Формируем ссылку в формате Telegram Mini App
  return `https://t.me/${botUsername}?startapp=${session.referral_link}`
}

// Копирование реферальной ссылки в буфер обмена
const copyReferralLink = (session) => {
  const referralLink = getReferralLink(session)

  // Используем Clipboard API, если доступно
  if (navigator.clipboard) {
    navigator.clipboard.writeText(referralLink).then(() => {
      if (window.Telegram && window.Telegram.WebApp) {
        window.Telegram.WebApp.showAlert(t('sessions_view.link_copied'))
      }
    }).catch(err => {
      console.error('Не удалось скопировать ссылку: ', err)
      // Альтернативный метод копирования
      fallbackCopyTextToClipboard(referralLink)
    })
  } else {
    // Альтернативный метод копирования для старых браузеров
    fallbackCopyTextToClipboard(referralLink)
  }
}

// Альтернативный метод копирования текста в буфер обмена
const fallbackCopyTextToClipboard = (text) => {
  const textArea = document.createElement("textarea")
  textArea.value = text

  // Избегаем прокрутки страницы
  textArea.style.top = "0"
  textArea.style.left = "0"
  textArea.style.position = "fixed"
  textArea.style.opacity = "0"

  document.body.appendChild(textArea)
  textArea.focus()
  textArea.select()

  try {
    const successful = document.execCommand('copy')
    if (successful) {
      if (window.Telegram && window.Telegram.WebApp) {
        window.Telegram.WebApp.showAlert(t('sessions_view.link_copied'))
      }
    } else {
      console.error('Не удалось скопировать ссылку')
      if (window.Telegram && window.Telegram.WebApp) {
        window.Telegram.WebApp.showAlert(t('sessions_view.copy_error'))
      }
    }
  } catch (err) {
    console.error('Ошибка при копировании ссылки: ', err)
    if (window.Telegram && window.Telegram.WebApp) {
      window.Telegram.WebApp.showAlert(t('sessions_view.copy_error'))
    }
  }

  document.body.removeChild(textArea)
}

// Форматирование даты
const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('ru-RU', {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric'
  })
}

// Загрузка сессий
const loadSessions = async () => {
  if (loading.value || noMoreSessions.value) return

  try {
    loading.value = true
    const response = await apiGet('sessions')

    if (response.ok) {
      const data = await response.json()
      // Проверяем, что data - массив
      if (Array.isArray(data)) {
        sessions.value = data
        noMoreSessions.value = true
      } else {
        console.warn('Получен неожиданный формат данных:', data)
        noMoreSessions.value = true
      }
    } else {
      console.error('Ошибка при загрузке сессий:', response.status)
      if (window.Telegram && window.Telegram.WebApp) {
        window.Telegram.WebApp.showAlert(t('sessions_view.load_error'))
      }
    }
  } catch (error) {
    console.error('Ошибка при запросе сессий:', error)
    if (window.Telegram && window.Telegram.WebApp) {
      window.Telegram.WebApp.showAlert(t('sessions_view.load_error_general'))
    }
  } finally {
    loading.value = false
  }
}

// Переход к созданию сессии
const goToCreateSession = () => {
  router.push('/sessions/create')
}

// Просмотр сессии
const viewSession = (session) => {
  router.push(`/sessions/${session.id}`)
}

// Удаление сессии с подтверждением
const deleteSession = (session) => {
  if (window.Telegram && window.Telegram.WebApp) {
    window.Telegram.WebApp.showConfirm(
      t('session_detail_view.delete_confirm') + ' "' + session.name + '"?',
      async (confirmed) => {
        if (confirmed) {
          try {
            const response = await apiDelete(`sessions/${session.id}`)

            if (response.ok) {
              // Успешно удалено, обновляем список сессий
              window.Telegram.WebApp.showAlert(t('session_detail_view.delete_success'))
              loadSessions() // Перезагружаем список сессий
            } else {
              console.error('Ошибка при удалении сессии:', response.status)
              window.Telegram.WebApp.showAlert(t('session_detail_view.delete_error'))
            }
          } catch (error) {
            console.error('Ошибка при запросе удаления сессии:', error)
            window.Telegram.WebApp.showAlert(t('session_detail_view.delete_error_general'))
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
  loadSessions()
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

.session-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  flex-wrap: wrap;
  gap: 8px;
}

.session-name {
  color: var(--tg-theme-text-color, #333333);
  font-size: 1.1rem;
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

.session-controls {
  display: flex;
  align-items: center;
  gap: 8px;
}

.session-button {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
  border: none;
  border-radius: 50%;
  width: 22px;
  height: 22px;
  font-size: 16px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
}

.session-button:hover {
  opacity: 0.8;
  transform: scale(1.1);
}

.referral-link-container {
  display: flex;
  align-items: center;
  gap: 8px;
  width: 100%;
}

.referral-link {
  flex: 1;
  word-break: break-all;
  font-size: 0.8rem;
}

.copy-button {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
  border: none;
  border-radius: 4px;
  padding: 4px 8px;
  font-size: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
  flex-shrink: 0;
}

.copy-button:hover {
  opacity: 0.8;
}

.delete-button {
  background: #ff4757 !important;
}

/* Плавные переходы для всех элементов */
.page,
.page-title,
.page-description,
.session-name,
.detail-label,
.detail-value {
  transition: all 0.3s ease;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .session-header {
    flex-direction: column;
    align-items: flex-start;
  }

  .session-name {
    font-size: 1rem;
  }

  .session-card {
    padding: 12px;
  }
}

/* Dark theme adjustments */
:global(.tg-theme-dark) .session-card {
  box-shadow: 0 1px 6px rgba(0, 0, 0, 0.2);
}
</style>