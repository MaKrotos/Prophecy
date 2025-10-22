<template>
  <div class="page">
    <h2 class="page-title">🎮 {{ t('create_session_view.title') }}</h2>
    <p class="page-description">{{ t('create_session_view.description') }}</p>
    
    <form @submit.prevent="createSession" class="session-form">
      <div class="form-group">
        <ThemedInput
          v-model="sessionData.name"
          :label="t('create_session_view.name_label')"
          type="text"
          :placeholder="t('create_session_view.name_placeholder')"
        />
      </div>
      
      <div class="form-group">
        <ThemedInput
          v-model="sessionData.description"
          :label="t('create_session_view.description_label')"
          type="textarea"
          :placeholder="t('create_session_view.description_placeholder')"
          inputType="large"
        />
      </div>
      
      <div class="form-actions">
        <ThemedButton
          buttonType="secondary"
          @click="goBack"
        >
          {{ t('create_session_view.cancel') }}
        </ThemedButton>
        <ThemedButton
          buttonType="primary"
          :disabled="isSubmitting"
          type="submit"
        >
          {{ isSubmitting ? t('create_session_view.creating') : t('create_session_view.create') }}
        </ThemedButton>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useApi } from '../telegram/composables/useApi'
import { useLocalization } from '@/locales/index.js'
import ThemedInput from '@/components/ThemedInput.vue'
import ThemedButton from '@/components/ThemedButton.vue'

const { t } = useLocalization()
const router = useRouter()
const { apiPost, testUrlTransformation } = useApi()

// Тест преобразования URL
testUrlTransformation('sessions');

const sessionData = ref({
  name: '',
  description: ''
})

const isSubmitting = ref(false)

// Создание сессии
const createSession = async () => {
  if (isSubmitting.value) return
  
  try {
    isSubmitting.value = true
    
    const response = await apiPost('sessions', sessionData.value)
    
    if (response.ok) {
      const data = await response.json()
      if (window.Telegram && window.Telegram.WebApp) {
        window.Telegram.WebApp.showAlert(t('create_session_view.success'))
      }
      // Переход к списку сессий
      router.push('/sessions')
    } else {
      const errorData = await response.json()
      const errorMessage = errorData.error || t('create_session_view.create_error')
      if (window.Telegram && window.Telegram.WebApp) {
        window.Telegram.WebApp.showAlert(errorMessage)
      }
    }
  } catch (error) {
    console.error('Ошибка при создании сессии:', error)
    if (window.Telegram && window.Telegram.WebApp) {
      window.Telegram.WebApp.showAlert(t('create_session_view.create_error_general'))
    }
  } finally {
    isSubmitting.value = false
  }
}

// Возврат назад
const goBack = () => {
  router.go(-1)
}
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

.session-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

/* Плавные переходы для всех элементов */
.page,
.page-title,
.page-description {
  transition: all 0.3s ease;
}

.form-actions {
  display: flex;
  gap: 12px;
  margin-top: 8px;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .page {
    padding: 12px;
  }
  
  .page-title {
    font-size: 1.3rem;
  }
  
  .form-actions {
    flex-direction: column;
  }
}
</style>