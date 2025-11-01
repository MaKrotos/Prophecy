<template>
  <div class="page">
    <div v-if="loading" class="loading">
      {{ t('session_join_view.loading') }}
    </div>
    <div v-else-if="success" class="success-container">
      <div class="success-card">
        <div class="success-icon">✅</div>
        <h2 class="success-title">{{ t('session_join_view.success') }}</h2>
        <p class="success-message">{{ t('session_join_view.description') }}</p>
        <div class="progress-bar">
          <div class="progress-fill"></div>
        </div>
      </div>
    </div>
    <div v-else-if="error" class="error-container">
      <div class="error-card">
        <div class="error-icon">⚠️</div>
        <h2 class="error-title">{{ t('session_join_view.join_error') }}</h2>
        <p class="error-message">{{ error }}</p>
        <ThemedButton button-type="primary" @click="retryLoad">
          {{ t('auth.error.retry') }}
        </ThemedButton>
      </div>
    </div>
    <div v-else class="join-content">
      <h2 class="page-title">🎮 {{ t('session_join_view.title') }}</h2>
      <p class="page-description">{{ t('session_join_view.you_invite_join') }}</p>

      <SessionInfo :session="{ name: sessionName, description: sessionDescription, architect_name: '', created_at: '' }"
        :show-name="true" :show-description="true" />

      <div class="button-group">
        <ThemedButton button-type="primary" class="join-button confirm-button" @click="joinSession" :disabled="joining">
          {{ joining ? t('session_join_view.joining') : t('session_join_view.yes') }}
        </ThemedButton>

        <ThemedButton button-type="secondary" @click="cancelJoin">
          {{ t('session_join_view.no') }}
        </ThemedButton>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useApi } from '../telegram/composables/useApi'
import ThemedButton from '../components/ThemedButton.vue'
import SessionInfo from '../components/SessionInfo.vue'
import { useLocalization } from '@/locales/index.js'
import { useTelegramWebAppSingleton } from '../telegram/composables/useTelegramWebAppSingleton'

const { t } = useLocalization()
const route = useRoute()
const router = useRouter()
const { apiPost, apiGet } = useApi()
const { startParam } = useTelegramWebAppSingleton()

const loading = ref(true)
const joining = ref(false)
const success = ref(false)
const error = ref('')
const sessionName = ref('')
const sessionDescription = ref('')

// Функция для повторной загрузки информации о сессии
const retryLoad = () => {
  error.value = ''
  loadSessionInfo()
}

// Присоединение к сессии
const joinSession = async () => {
  console.log("🔍 Присоединение к сессии по реферальной ссылке:", route.params.referral_link);
  console.log("🔍 Все параметры маршрута в joinSession:", route.params);
  if (joining.value) return

  try {
    joining.value = true
    const response = await apiPost(`joinSession/${route.params.referral_link}`)
    console.log("🔍 Ответ от сервера при присоединении к сессии:", response.status);

    if (response.ok) {
      success.value = true
      // Переход к списку сессий через 2 секунды
      setTimeout(() => {
        // Обнуляем startParam, чтобы окно подтверждения не открывалось повторно
        startParam.value = null
        router.push('/')
      }, 2000)
    } else {
      const errorData = await response.json()
      error.value = errorData.error || t('session_join_view.join_error')
      console.log("⚠️ Ошибка при присоединении к сессии:", error.value);
    }
  } catch (err) {
    console.error('Ошибка при присоединении к сессии:', err)
    error.value = t('session_join_view.join_error_general')
  } finally {
    joining.value = false
  }
}

// Отмена присоединения и возврат на главную страницу
const cancelJoin = () => {
  console.log("🔍 Отмена присоединения к сессии");
  console.log("🔍 Значение startParam перед обнулением:", startParam.value);
  // Обнуляем startParam, чтобы окно подтверждения не открывалось повторно
  startParam.value = null;
  console.log("🔍 Значение startParam после обнуления:", startParam.value);
  router.push('/')
}

// Получение информации о сессии
const loadSessionInfo = async () => {
  try {
    console.log("🔍 Загрузка информации о сессии по реферальной ссылке:", route.params.referral_link);
    console.log("🔍 Все параметры маршрута в loadSessionInfo:", route.params);
    loading.value = true
    // Сначала получаем информацию о сессии по реферальной ссылке
    const sessionResponse = await apiGet(`joinSession/${route.params.referral_link}`)

    if (sessionResponse.ok) {
      const sessionData = await sessionResponse.json()
      sessionName.value = sessionData.name
      sessionDescription.value = sessionData.description
      console.log("✅ Информация о сессии загружена:", sessionData);
    } else {
      error.value = t('session_join_view.load_error')
      console.log("⚠️ Ошибка загрузки информации о сессии:", sessionResponse.status);
    }
  } catch (err) {
    console.error('Ошибка при загрузке информации о сессии:', err)
    error.value = t('session_join_view.load_error_general')
  } finally {
    loading.value = false
  }
}

// Первоначальная загрузка
onMounted(() => {
  console.log("🔍 SessionJoinView mounted, referral_link:", route.params.referral_link);
  console.log("🔍 SessionJoinView mounted, все параметры маршрута:", route.params);
  loadSessionInfo()
})
</script>

<style scoped>
.page {
  padding: 16px;
  background-color: var(--tg-theme-bg-color, #f5f5f5);
  transition: background-color 0.3s ease;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
}

.page-title {
  color: var(--tg-theme-text-color, #000000);
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 8px;
  transition: color 0.3s ease;
  text-align: center;
}

.page-description {
  color: var(--tg-theme-hint-color, #666666);
  font-size: 1rem;
  margin-bottom: 24px;
  transition: color 0.3s ease;
  text-align: center;
}

.join-content {
  width: 100%;
  max-width: 500px;
  text-align: center;
}


.loading {
  text-align: center;
  padding: 24px;
  color: var(--tg-theme-hint-color, #666666);
  transition: color 0.3s ease;
}

/* Success Container Styles */
.success-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  padding: 24px;
}

.success-card {
  background: var(--tg-theme-secondary-bg-color, white);
  border-radius: 16px;
  padding: 32px 24px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  text-align: center;
  max-width: 500px;
  width: 100%;
  animation: fadeIn 0.5s ease-out;
}

.success-icon {
  font-size: 3rem;
  margin-bottom: 16px;
}

.success-title {
  color: var(--tg-theme-text-color, #333333);
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 12px;
}

.success-message {
  color: var(--tg-theme-hint-color, #666666);
  font-size: 1rem;
  margin-bottom: 24px;
}

.progress-bar {
  height: 4px;
  background: var(--tg-theme-secondary-bg-color, #f0f0f0);
  border-radius: 2px;
  overflow: hidden;
  margin-top: 24px;
}

.progress-fill {
  height: 100%;
  width: 100%;
  background: var(--tg-theme-button-color, #667eea);
  animation: progress 2s linear forwards;
  transform-origin: left;
}

/* Error Container Styles */
.error-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  padding: 24px;
}

.error-card {
  background: var(--tg-theme-secondary-bg-color, white);
  border-radius: 16px;
  padding: 32px 24px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  text-align: center;
  max-width: 500px;
  width: 100%;
  animation: fadeIn 0.5s ease-out;
}

.error-icon {
  font-size: 3rem;
  margin-bottom: 16px;
}

.error-title {
  color: var(--tg-theme-text-color, #333333);
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 12px;
}

.error-message {
  color: #ff4757;
  font-size: 1rem;
  margin-bottom: 24px;
}

/* Button Group Styles */
.button-group {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-top: 24px;
}

.join-button,
.confirm-button {
  width: 100%;
}

/* Animations */
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes progress {
  from {
    transform: scaleX(0);
  }

  to {
    transform: scaleX(1);
  }
}

/* Плавные переходы для всех элементов */
.page,
.page-title,
.page-description {
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

  .success-card,
  .error-card {
    padding: 24px 16px;
  }

  .success-title,
  .error-title {
    font-size: 1.3rem;
  }
}
</style>