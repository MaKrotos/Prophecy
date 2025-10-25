<template>
  <!-- Сообщение об успехе -->
  <ThemedCard v-if="successMessage" class="message-card success">
    <div class="message-content">
      <span class="icon">✓</span>
      <p>{{ successMessage }}</p>
    </div>
  </ThemedCard>

  <!-- Сообщение об ошибке -->
  <ThemedCard v-if="error" class="message-card error">
    <div class="message-content">
      <span class="icon">⚠</span>
      <p>{{ error }}</p>
    </div>
    <ThemedButton 
      v-if="showRetryButton" 
      buttonType="danger" 
      class="retry-button"
      @click="$emit('retry')"
    >
      {{ t('scan_view.retry_camera') }}
    </ThemedButton>
  </ThemedCard>
</template>

<script setup>
import { useLocalization } from '@/locales/index.js'
import ThemedCard from '@/components/ThemedCard.vue'
import ThemedButton from '@/components/ThemedButton.vue'

const { t } = useLocalization()

defineProps({
  successMessage: {
    type: String,
    default: ''
  },
  error: {
    type: String,
    default: ''
  },
  showRetryButton: {
    type: Boolean,
    default: false
  }
})

defineEmits(['retry'])
</script>

<style scoped>
.message-card {
  margin-top: 20px;
  width: 100%;
  max-width: 400px;
}

.message-content {
  display: flex;
  align-items: center;
  gap: 12px;
}

.message-content .icon {
  font-size: 1.5rem;
  font-weight: bold;
}

.message-card.success {
  border-left: 4px solid #4CAF50;
  background-color: #f8fff8;
}

.message-card.success .icon {
  color: #4CAF50;
}

.message-card.error {
  border-left: 4px solid #f44336;
  background-color: #ffebee;
}

.message-card.error .icon {
  color: #f44336;
}

.retry-button {
  margin-top: 16px;
  width: 100%;
}
</style>