<template>
  <div class="page">
    <!-- Компонент статистики пользователей (только для админов) -->
    <UserStats v-if="isAdmin" />

    <!-- User Info Section -->
    <div class="user-info-section" v-if="userInfo">
      <div class="user-card">
        <div class="user-header">
          <h3 class="user-name">{{ userInfo.generatedName }}</h3>
          <div class="user-id">ID: {{ userInfo.id }}</div>
        </div>

        <div class="user-details">
          <div class="detail-item" v-if="userInfo.isAdmin">
            <span class="admin-badge">{{ t('home_view.admin') }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Create Session Button (only for architects) -->
    <ThemedButton
      v-if="isArchitect"
      buttonType="primary"
      class="create-session-button"
      @click="goToCreateSession"
    >
      + {{ t('home_view.sessions.create_session') }}
    </ThemedButton>

    <!-- Session List Component -->
    <SessionList />
  </div>
</template>

<script setup>
import { getUserInfoFromToken } from '../telegram/auth/user'

import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import UserStats from '../components/UserStats.vue'
import SessionList from '../components/SessionList.vue'
import ThemedButton from '../components/ThemedButton.vue'
import { useLocalization } from '@/locales/index.js'

const { t } = useLocalization()
const router = useRouter()

const userInfo = ref(null)
const isAdmin = computed(() => userInfo.value?.isAdmin || false)
const isArchitect = computed(() => {
  if (!userInfo.value) return false
  // Проверяем как числовое, так и строковое значение роли
  return userInfo.value.role && (userInfo.value.role === 'Архитектор' || userInfo.value.role === 1)
})

const goToCreateSession = () => {
  router.push('/sessions/create')
}

onMounted(() => {
  // Получаем информацию о пользователе из JWT токена
  userInfo.value = getUserInfoFromToken()
  console.log('User Info:', userInfo.value)
})
</script>

<style scoped>
.page {
  padding: 16px;

  background-color: var(--tg-theme-bg-color, #f5f5f5);
  transition: background-color 0.3s ease;
}

/* User Info Section Styles */
.user-info-section {
  margin-bottom: 24px;
}

.user-card {
  background: var(--tg-theme-secondary-bg-color, white);
  border-radius: 16px;
  padding: 20px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  border: 1px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
  transition: all 0.3s ease;
}

.user-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  flex-wrap: wrap;
  gap: 8px;
}

.user-name {
  color: var(--tg-theme-text-color, #333333);
  font-size: 1.5rem;
  font-weight: 700;
  margin: 0;
  transition: color 0.3s ease;
}

.user-details {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.detail-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.detail-label {
  color: var(--tg-theme-hint-color, #666666);
  font-size: 0.9rem;
  transition: color 0.3s ease;
}

.detail-value {
  color: var(--tg-theme-text-color, #333333);
  font-weight: 500;
  transition: color 0.3s ease;
}

.admin-badge {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
  padding: 6px 12px;
  border-radius: 16px;
  font-size: 0.9rem;
  font-weight: 600;
}

/* Dark theme adjustments */
:global(.tg-theme-dark) .user-card {
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .user-header {
    flex-direction: column;
    align-items: flex-start;
  }

  .user-name {
    font-size: 1.3rem;
  }

  .user-card {
    padding: 16px;
  }
}

</style>
