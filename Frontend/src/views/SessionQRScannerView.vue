<template>
  <div class="page">
    <div class="header-section">
      <h2 class="page-title">🔍 {{ t('session_qr_scanner_view.title') }}</h2>
      <p class="page-description">{{ t('session_qr_scanner_view.description') }}</p>
    </div>
    
    <div class="scanner-content">
      <!-- Отображение дешифрованных данных -->
      <DecryptedQRDisplay 
        v-if="scannedData" 
        :encrypted-value="scannedData" 
      />
      
      <!-- Сканер QR-кода -->
      <ThemedCard v-if="!scannedData" card-type="default" class="scanner-card">
        <div class="scanner-placeholder">
          <div class="scanner-icon">📷</div>
          <p class="scanner-text">{{ t('session_qr_scanner_view.camera_access') }}</p>
          <ThemedButton button-type="primary" class="scanner-button" @click="startScanning">
            {{ t('session_qr_scanner_view.enable_camera') }}
          </ThemedButton>
        </div>
      </ThemedCard>
      
      <!-- Кнопка для сброса сканирования -->
      <ThemedCard v-if="scannedData" card-type="default" class="reset-card">
        <ThemedButton button-type="secondary" class="reset-button" @click="resetScanning">
          🔄 {{ t('session_qr_scanner_view.scan_again') }}
        </ThemedButton>
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
import { ref } from 'vue'
import { useLocalization } from '@/locales/index.js'
import ThemedCard from '../components/ThemedCard.vue'
import ThemedButton from '../components/ThemedButton.vue'
import DecryptedQRDisplay from '../components/DecryptedQRDisplay.vue'

const { t } = useLocalization()

const scannedData = ref(null)

// Запуск сканирования (имитация)
const startScanning = () => {
  // В реальной реализации здесь будет код для доступа к камере
  // и сканирования QR-кода
  
  // Для демонстрации имитируем сканирование
  setTimeout(() => {
    // Здесь будет зашифрованное значение из QR-кода
    // В реальной реализации это значение будет получено от сканера
    scannedData.value = "зашифрованные_данные_из_qr_кода" // Заглушка
  }, 1000)
}

// Сброс сканирования
const resetScanning = () => {
  scannedData.value = null
}
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
.instructions-card,
.reset-card {
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

.reset-button {
  width: 100%;
  justify-content: center;
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