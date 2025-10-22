<template>
  <div class="page">
    <div v-if="loading" class="loading">
      {{ t('session_join_view.loading') }}
    </div>
    <div v-else-if="success" class="success">
      {{ t('session_join_view.success') }}
    </div>
    <div v-else-if="error" class="error">
      {{ error }}
    </div>
    <div v-else class="join-content">
      <h2 class="page-title">🎮 {{ t('session_join_view.title') }}</h2>
      <p class="page-description">{{ t('session_join_view.description') }}</p>

      <div class="session-info">
        <h3 class="session-name">{{ sessionName }}</h3>
        <p class="session-description">{{ sessionDescription || t('session_join_view.no_description') }}</p>
      </div>

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
import { useLocalization } from '@/locales/index.js'
import { useTelegramWebApp } from '../telegram/composables/useTelegramWebApp'

const { t } = useLocalization()
const route = useRoute()
const router = useRouter()
const { apiPost, apiGet } = useApi()
const { startParam } = useTelegramWebApp()

const loading = ref(true)
const joining = ref(false)
const success = ref(false)
const error = ref('')
const sessionName = ref('')
const sessionDescription = ref('')

// Присоединение к сессии
const joinSession = async () => {
  console.log("🔍 Присоединение к сессии по реферальной ссылке:", route.params.referral_link);
  console.log("🔍 Все параметры маршрута в joinSession:", route.params);
  if (joining.value) return

  try {
    joining.value = true
    const response = await apiPost(`sessions/join/${route.params.referral_link}`)
    console.log("🔍 Ответ от сервера при присоединении к сессии:", response.status);

    if (response.ok) {
      success.value = true
      if (window.Telegram && window.Telegram.WebApp) {
        window.Telegram.WebApp.showAlert(t('session_join_view.success'))
      }
      // Переход к списку сессий через 2 секунды
      setTimeout(() => {
        // Обнуляем startParam, чтобы окно подтверждения не открывалось повторно
        startParam.value = null
        router.push('/sessions')
      }, 2000)
    } else {
      const errorData = await response.json()
      error.value = errorData.error || t('session_join_view.join_error')
      console.log("⚠️ Ошибка при присоединении к сессии:", error.value);
      if (window.Telegram && window.Telegram.WebApp) {
        window.Telegram.WebApp.showAlert(error.value)
      }
    }
  } catch (err) {
    console.error('Ошибка при присоединении к сессии:', err)
    error.value = t('session_join_view.join_error_general')
    if (window.Telegram && window.Telegram.WebApp) {
      window.Telegram.WebApp.showAlert(error.value)
    }
  } finally {
    joining.value = false
  }
}

// Отмена присоединения и возврат к списку сессий
const cancelJoin = () => {
  console.log("🔍 Отмена присоединения к сессии");
  console.log("🔍 Значение startParam перед обнулением:", startParam.value);
  // Обнуляем startParam, чтобы окно подтверждения не открывалось повторно
  startParam.value = null;
  console.log("🔍 Значение startParam после обнуления:", startParam.value);
  router.push('/sessions')
}

// Получение информации о сессии
const loadSessionInfo = async () => {
  try {
    console.log("🔍 Загрузка информации о сессии по реферальной ссылке:", route.params.referral_link);
    console.log("🔍 Все параметры маршрута в loadSessionInfo:", route.params);
    loading.value = true
    // Сначала получаем информацию о сессии по реферальной ссылке
    const sessionResponse = await apiGet(`sessions/join/${route.params.referral_link}`)

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

.session-info {
  background: var(--tg-theme-secondary-bg-color, white);
  border-radius: 12px;
  padding: 20px;
  margin-bottom: 24px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
}

.session-name {
  color: var(--tg-theme-text-color, #333333);
  font-size: 1.25rem;
  font-weight: 600;
  margin: 0 0 12px 0;
  transition: color 0.3s ease;
}

.session-description {
  color: var(--tg-theme-text-color, #333333);
  font-size: 1rem;
  line-height: 1.5;
  margin: 0;
  transition: color 0.3s ease;
}

.button-group {
  display: flex;
  gap: 12px;
  margin-top: 24px;
}

.join-button,
.cancel-button {
  flex: 1;
  border: none;
  border-radius: 12px;
  padding: 14px 20px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.join-button {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
}

.join-button:hover:not(:disabled) {
  opacity: 0.9;
  transform: translateY(-1px);
}

.join-button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.cancel-button {
  background: var(--tg-theme-secondary-bg-color, #f0f0f0);
  color: var(--tg-theme-text-color, #333333);
}

.cancel-button:hover {
  opacity: 0.9;
  transform: translateY(-1px);
}

.loading,
.success,
.error {
  text-align: center;
  padding: 24px;
  color: var(--tg-theme-hint-color, #666666);
  transition: color 0.3s ease;
}

.success {
  color: #2ed573;
}

.error {
  color: #ff4757;
}

/* Плавные переходы для всех элементов */
.page,
.page-title,
.page-description,
.session-name,
.session-description {
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

  .session-info {
    padding: 16px;
  }
}
</style>