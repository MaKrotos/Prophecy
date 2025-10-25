<template>
  <div v-if="loading" class="loading-overlay">
    <div class="loading-content">
      <div class="spinner"></div>
      <p class="loading-text">{{ text || t('scan_view.camera_access_requested') }}</p>
    </div>
  </div>
</template>

<script setup>
import { useLocalization } from '@/locales/index.js'

const { t } = useLocalization()

defineProps({
  loading: {
    type: Boolean,
    required: true
  },
  text: {
    type: String,
    default: ''
  }
})
</script>

<style scoped>
.loading-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.loading-content {
  background: var(--tg-theme-bg-color, white);
  padding: 32px;
  border-radius: 16px;
  text-align: center;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
}

.spinner {
  width: 48px;
  height: 48px;
  border: 4px solid rgba(102, 126, 234, 0.3);
  border-top: 4px solid var(--tg-theme-button-color, #667eea);
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.loading-text {
  margin: 0;
  font-size: 1.1rem;
  color: var(--tg-theme-text-color, #333333);
  font-weight: 500;
}
</style>