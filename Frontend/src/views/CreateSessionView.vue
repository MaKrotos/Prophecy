<template>
  <div class="page">
    <h2 class="page-title">🎮 {{ t('create_session_view.title') }}</h2>
    <p class="page-description">{{ t('create_session_view.description') }}</p>
    
    <form @submit.prevent="createSession" class="session-form">
      <div class="form-group">
        <label class="form-label">{{ t('create_session_view.name_label') }}</label>
        <input 
          v-model="sessionData.name" 
          type="text" 
          class="form-input"
          :placeholder="t('create_session_view.name_placeholder')"
          required
        />
      </div>
      
      <div class="form-group">
        <label class="form-label">{{ t('create_session_view.description_label') }}</label>
        <textarea 
          v-model="sessionData.description" 
          class="form-textarea"
          :placeholder="t('create_session_view.description_placeholder')"
          rows="4"
        ></textarea>
      </div>
      
      <div class="form-actions">
        <button 
          type="button" 
          class="cancel-button"
          @click="goBack"
        >
          {{ t('create_session_view.cancel') }}
        </button>
        <button 
          type="submit" 
          class="submit-button"
          :disabled="isSubmitting"
        >
          {{ isSubmitting ? t('create_session_view.creating') : t('create_session_view.create') }}
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useApi } from '../telegram/composables/useApi'
import { useLocalization } from '@/locales/index.js'

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

.form-label {
  color: var(--tg-theme-text-color, #333333);
  font-weight: 500;
  transition: color 0.3s ease;
}

.form-input,
.form-textarea {
  background: var(--tg-theme-secondary-bg-color, white);
  border: 1px solid var(--tg-theme-hint-color, #e0e0e0);
  border-radius: 12px;
  padding: 12px 16px;
  font-size: 1rem;
  color: var(--tg-theme-text-color, #333333);
  transition: all 0.3s ease;
}

.form-input:focus,
.form-textarea:focus {
  outline: none;
  border-color: var(--tg-theme-button-color, #667eea);
  box-shadow: 0 0 0 2px rgba(102, 126, 234, 0.2);
}

.form-textarea {
  resize: vertical;
  min-height: 100px;
}

.form-actions {
  display: flex;
  gap: 12px;
  margin-top: 8px;
}

.cancel-button,
.submit-button {
  flex: 1;
  padding: 14px 20px;
  border: none;
  border-radius: 12px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.cancel-button {
  background: var(--tg-theme-secondary-bg-color, #f0f0f0);
  color: var(--tg-theme-text-color, #333333);
}

.submit-button {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
}

.cancel-button:hover {
  opacity: 0.8;
}

.submit-button:hover:not(:disabled) {
  opacity: 0.9;
  transform: translateY(-1px);
}

.submit-button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

/* Плавные переходы для всех элементов */
.page,
.page-title,
.page-description,
.form-label,
.form-input,
.form-textarea {
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
  
  .form-actions {
    flex-direction: column;
  }
}
</style>