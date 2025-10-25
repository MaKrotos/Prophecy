<template>
  <div class="page">
    <div class="header-section">
      <h2 class="page-title">📱 {{ t('session_my_qr_view.title') }}</h2>
      <p class="page-description">{{ t('session_my_qr_view.description') }}</p>
    </div>
    
    <div class="qr-content">
      <ThemedCard card-type="default" class="qr-card">
        <div class="qr-header">
          <h3>{{ t('session_my_qr_view.your_qr_code') }}</h3>
          <p class="qr-subtitle">{{ t('session_my_qr_view.share_qr') }}</p>
        </div>
        
        <div class="qr-display">
          <div class="qr-placeholder">
            <div class="qr-grid"></div>
          </div>
          <div class="qr-info">
            <p class="qr-text">prophecy://session/player/12345</p>
          </div>
        </div>
        
        <div class="qr-actions">
          <ThemedButton button-type="primary" class="action-button">
            📋 {{ t('session_my_qr_view.copy_link') }}
          </ThemedButton>
          <ThemedButton button-type="secondary" class="action-button">
            📤 {{ t('session_my_qr_view.share') }}
          </ThemedButton>
        </div>
      </ThemedCard>
      
      <ThemedCard card-type="default" class="instructions-card">
        <h3>{{ t('session_my_qr_view.how_to_use') }}</h3>
        <ol class="instructions-list">
          <li>{{ t('session_my_qr_view.instruction_1') }}</li>
          <li>{{ t('session_my_qr_view.instruction_2') }}</li>
          <li>{{ t('session_my_qr_view.instruction_3') }}</li>
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

.qr-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.qr-card,
.instructions-card {
  background: var(--tg-theme-secondary-bg-color, white);
  border-radius: 12px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}

.qr-card:hover,
.instructions-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.qr-header {
  text-align: center;
  margin-bottom: 24px;
}

.qr-header h3 {
  color: var(--tg-theme-text-color, #000000);
  font-size: 1.25rem;
  font-weight: 600;
  margin: 0 0 8px 0;
  transition: color 0.3s ease;
}

.qr-subtitle {
  color: var(--tg-theme-hint-color, #666666);
  font-size: 0.9rem;
  margin: 0;
  transition: color 0.3s ease;
}

.qr-display {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
  margin-bottom: 24px;
}

.qr-placeholder {
  width: 200px;
  height: 200px;
  border: 2px solid var(--tg-theme-hint-color, #666666);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: white;
}

.qr-grid {
  width: 180px;
  height: 180px;
  background: 
    linear-gradient(to right, #000 2px, transparent 2px) 0 0,
    linear-gradient(to right, #000 2px, transparent 2px) 0 100%,
    linear-gradient(to left, #000 2px, transparent 2px) 100% 0,
    linear-gradient(to left, #000 2px, transparent 2px) 100% 100%,
    linear-gradient(to bottom, #000 2px, transparent 2px) 0 0,
    linear-gradient(to bottom, #000 2px, transparent 2px) 100% 0,
    linear-gradient(to top, #000 2px, transparent 2px) 0 100%,
    linear-gradient(to top, #000 2px, transparent 2px) 100% 100%;
  background-repeat: no-repeat;
  background-size: 20px 20px;
}

.qr-info {
  text-align: center;
}

.qr-text {
  color: var(--tg-theme-text-color, #333333);
  font-size: 0.875rem;
  font-family: monospace;
  word-break: break-all;
  margin: 0;
  transition: color 0.3s ease;
}

.qr-actions {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.action-button {
  width: 100%;
  justify-content: center;
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
.qr-header h3,
.qr-subtitle,
.qr-text,
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
  
  .qr-placeholder {
    width: 160px;
    height: 160px;
  }
  
  .qr-grid {
    width: 140px;
    height: 140px;
    background-size: 16px 16px;
  }
  
  .qr-actions {
    flex-direction: column;
  }
}
</style>