<template>
  <div class="page">
    <h2 class="page-title">👥 {{ t('users_view.title') }}</h2>
    <p class="page-description">{{ t('users_view.description') }}</p>

    <AnimatedCardList :items="users" :loading="loading" :no-more-items="noMoreUsers" key-field="id"
      card-class="user-card" :animation-delay="0.1" :loading-text="t('users_view.loading')"
      :no-more-items-text="t('users_view.no_more')">
      <template #card="{ item: user }">
        <div class="user-header">
          <div class="user-info">
            <h3 class="user-name">{{ user.generated_name }}</h3>
          </div>
          <div class="user-controls">
            <div class="user-badge"
              :class="{ 'admin-badge': user.is_admin, 'architect-badge': user.role && (user.role.String === 'Архитектор' || user.role === 1) }">
              <span v-if="user.role && (user.role.String === 'Архитектор' || user.role === 1)">{{ t('users_view.architect') }}</span>
              <span v-else-if="user.is_admin">{{ t('users_view.admin') }}</span>
              <span v-else>{{ t('users_view.user') }}</span>
            </div>
            <ThemedButton button-type="icon" @click="showRoleMenu(user)">
              ⚙️
            </ThemedButton>
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
            <span class="detail-label">{{ t('users_view.registration') }}:</span>
            <span class="detail-value">{{ formatDate(user.created_at) }}</span>
          </div>
        </div>
      </template>
    </AnimatedCardList>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useApi } from '../telegram/composables/useApi'
import AnimatedCardList from '../components/AnimatedCardList.vue'
import ThemedButton from '../components/ThemedButton.vue'
import { useLocalization } from '@/locales/index.js'
const { t } = useLocalization()

const { apiGet, apiPut } = useApi()

const users = ref([])
const loading = ref(false)
const offset = ref(0)
const limit = ref(20)
const noMoreUsers = ref(false)

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
      // Проверяем, что data - массив
      if (Array.isArray(data)) {
        if (data.length > 0) {
          // Нормализуем данные пользователей, чтобы убедиться, что role имеет правильную структуру
          const normalizedData = data.map(user => {
            // Убедимся, что у пользователя есть корректная структура role
            if (typeof user.role === 'number') {
              // Если role - число, преобразуем в строку для отображения
              const roleString = user.role === 1 ? 'Архитектор' : ''
              return {
                ...user,
                role: {
                  String: roleString,
                  Valid: user.role === 1
                }
              }
            } else if (user.role && typeof user.role === 'string') {
              // Если role - строка, преобразуем в ожидаемую структуру
              return {
                ...user,
                role: {
                  String: user.role,
                  Valid: user.role !== ''
                }
              }
            } else if (user.role && typeof user.role === 'object') {
              // Если role уже объект, убедимся, что он имеет правильную структуру
              return {
                ...user,
                role: {
                  String: user.role.String || '',
                  Valid: user.role.Valid !== undefined ? user.role.Valid : (user.role.String !== undefined && user.role.String !== '')
                }
              }
            } else {
              // Если role отсутствует или null/undefined, создаем пустую структуру
              return {
                ...user,
                role: {
                  String: '',
                  Valid: false
                }
              }
            }
          })

          users.value = [...users.value, ...normalizedData]
          offset.value += data.length

          // Если получено меньше пользователей, чем запрашивали, значит больше нет
          if (data.length < limit.value) {
            noMoreUsers.value = true
          }
        } else {
          noMoreUsers.value = true
        }
      } else {
        // Если data не массив, значит это ошибка или пустой ответ
        console.warn('Получен неожиданный формат данных:', data)
        noMoreUsers.value = true
      }
    } else {
      console.error('Ошибка при загрузке пользователей:', response.status)
      alert(t('users_view.load_error'))
    }
  } catch (error) {
    console.error('Ошибка при запросе пользователей:', error)
    alert(t('users_view.load_error_general'))
  } finally {
    loading.value = false
    // После загрузки проверяем, нужно ли загрузить еще
    checkIfNeedMoreUsers()
  }
}

// Установка роли пользователя
const setUserRole = async (user, newRole) => {
  try {
    // Преобразуем строковое значение роли в числовое
    let roleValue
    if (newRole === 'Архитектор' || newRole === 1) {
      roleValue = 1
    } else {
      roleValue = 0
    }
    
    const response = await apiPut(`users/${user.id}/role`, { role: roleValue })

    if (response.ok) {
      // Обновляем роль пользователя в списке
      const updatedUsers = users.value.map(u =>
        u.id === user.id ? { ...u, role: { String: newRole, Valid: newRole !== '' } } : u
      )
      users.value = updatedUsers

      if (window.Telegram && window.Telegram.WebApp) {
        window.Telegram.WebApp.showAlert(t('users_view.role_updated'))
      }
    } else {
      if (window.Telegram && window.Telegram.WebApp) {
        window.Telegram.WebApp.showAlert(t('users_view.role_update_error'))
      }
    }
  } catch (error) {
    console.error('Ошибка при обновлении роли пользователя:', error)
    if (window.Telegram && window.Telegram.WebApp) {
      window.Telegram.WebApp.showAlert(t('users_view.role_update_error'))
    }
  }
}

// Показ меню управления ролью пользователя
const showRoleMenu = (user) => {
  // Определяем текущую роль пользователя
  let role = ''
  if (typeof user.role === 'number') {
    role = user.role === 1 ? 'Архитектор' : ''
  } else if (user.role && user.role.String) {
    role = user.role.String
  }

  if (role === 'Архитектор') {
    if (window.Telegram && window.Telegram.WebApp) {
      window.Telegram.WebApp.showConfirm(
        t('users_view.remove_role') + '?',
        (confirmed) => {
          if (confirmed) {
            setUserRole(user, 0) // 0 - обычная роль пользователя
          }
        }
      )
    }
  } else {
    // Для не-админов можно только назначить архитектором
    if (window.Telegram && window.Telegram.WebApp) {
      window.Telegram.WebApp.showConfirm(
        t('users_view.set_role') + ' ' + t('users_view.architect') + '?',
        (confirmed) => {
          if (confirmed) {
            setUserRole(user, 1) // 1 - роль архитектора
          }
        }
      )
    }
  }
}

// Проверка, нужно ли загрузить еще пользователей
const checkIfNeedMoreUsers = () => {
  const mainContent = document.querySelector('.main-content')
  if (mainContent) {
    const { scrollTop, scrollHeight, clientHeight } = mainContent
    const scrollPercentage = (scrollTop + clientHeight) / scrollHeight

    // Если скролл находится в пределах, где нужно подгружать, и еще есть что подгружать
    if (scrollPercentage > 0.8 && !noMoreUsers.value && !loading.value) {
      loadUsers()
    }
  }
}

// Обработка скролла главного контента
const handleScroll = () => {
  const mainContent = document.querySelector('.main-content')
  if (!mainContent) return

  const { scrollTop, scrollHeight, clientHeight } = mainContent
  const scrollPercentage = (scrollTop + clientHeight) / scrollHeight

  // Загружаем больше пользователей, когда пользователь прокрутил 80% контента
  if (scrollPercentage > 0.8 && !noMoreUsers.value && !loading.value) {
    loadUsers()
  }
}

// Первоначальная загрузка
onMounted(() => {
  loadUsers()
  // Добавляем обработчик скролла к main-content
  const mainContent = document.querySelector('.main-content')
  if (mainContent) {
    mainContent.addEventListener('scroll', handleScroll)
  }
})

// Удаляем обработчик скролла при размонтировании компонента
onUnmounted(() => {
  const mainContent = document.querySelector('.main-content')
  if (mainContent) {
    mainContent.removeEventListener('scroll', handleScroll)
  }
})
</script>

<style scoped>
.page {
  padding: 16px;

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
  padding-right: 8px;
}

.user-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  flex-wrap: wrap;
  gap: 8px;
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

.user-badge.architect-badge {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
  border: 2px solid #ffd700;
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

/* Плавные переходы для всех элементов */
.page,
.page-title,
.page-description,
.user-name,
.detail-label,
.detail-value {
  transition: all 0.3s ease;
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

.user-controls {
  display: flex;
  align-items: center;
  gap: 8px;
}
</style>