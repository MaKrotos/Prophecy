<template>
  <div class="page">
    <div class="header-section">
      <h2 class="page-title">🔍 {{ t('session_qr_scanner_view.title') }}</h2>
      <p class="page-description">{{ t('session_qr_scanner_view.description') }}</p>
    </div>
    
    <div class="scanner-content">
      <ThemedCard card-type="default" class="scanner-card">
        <div class="scanner-placeholder">
          <div class="scanner-icon">📷</div>
          <p class="scanner-text">{{ t('session_qr_scanner_view.camera_access') }}</p>
          <ThemedButton button-type="primary" class="scanner-button">
            {{ t('session_qr_scanner_view.enable_camera') }}
          </ThemedButton>
        </div>
      </ThemedCard>
      
      <ThemedCard card-type="default" class="instructions-card">
        <h3>{{ t('session_qr_scanner_view.how_to_scan') }}</h3>
        <ol class="instructions-list">
          <li>{{ t('session_qr_scanner_view.step_1') }}</li>
          <li>{{ t('session_qr_scanner_view.step_2') }}</li>
          <li>{{ t('session_qr_scanner_view.step_3') }}</li>
          <li>{{ t('session_qr_scanner_view.step_4') }}</li>
        </ol>
      </ThemedCard>
    </div>
  </div>
</template>

<script setup>
import { useLocalization } from '@/locales/index.js'
import ThemedCard from '../components/ThemedCard.vue'
import ThemedButton from '../components/ThemedButton.vue'

const { t } = useLocalization()
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

.scanner-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.scanner-card,
.instructions-card {
  background: var(--tg-theme-secondary-bg-color, white);
  border-radius: 12px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}

.scanner-card:hover,
.instructions-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.scanner-placeholder {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 20px;
  padding: 40px 20px;
  text-align: center;
}

.scanner-icon {
  font-size: 3rem;
}

.scanner-text {
  color: var(--tg-theme-text-color, #333333);
  font-size: 1rem;
  margin: 0;
  transition: color 0.3s ease;
}

.scanner-button {
  margin-top: 10px;
}

.instructions-card h3 {
  color: var(--tg-theme-text-color, #000000);
  font-size: 1.1rem;
  font-weight: 600;
  margin: 0 0 16px 0;
  transition: color 0.3s ease;
}

.instructions-list {
  color: var(--tg-theme-text-color, #333333);
  font-size: 0.95rem;
  line-height: 1.6;
  margin: 0;
  padding-left: 20px;
  transition: color 0.3s ease;
}

.instructions-list li {
  margin-bottom: 8px;
}

.instructions-list li:last-child {
  margin-bottom: 0;
}

/* Плавные переходы для всех элементов */
.page,
.page-title,
.page-description,
.scanner-text,
.instructions-card h3,
.instructions-list {
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
  
  .scanner-placeholder {
    padding: 30px 15px;
  }
}
</style>