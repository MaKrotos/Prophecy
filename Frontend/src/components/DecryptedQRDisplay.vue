<template>
  <ThemedCard v-if="decryptedData" class="qr-display-card">
    <h3>{{ t('scan_view.scanned_qr') }}</h3>
    <div class="decrypted-info">
      <p><strong>{{ t('session_detail_view.session') }}:</strong> {{ decryptedData.session_name }}</p>
      <p><strong>{{ t('session_detail_view.player_id') }}:</strong> {{ decryptedData.player_id }}</p>
      <p><strong>{{ t('session_detail_view.session_id') }}:</strong> {{ decryptedData.session_id }}</p>
      <p><strong>{{ t('session_detail_view.timestamp') }}:</strong> {{ formatDate(decryptedData.timestamp) }}</p>
    </div>
  </ThemedCard>
  
  <ThemedCard v-else-if="error" class="qr-display-card error-card">
    <h3>{{ t('scan_view.error') }}</h3>
    <p class="error-message">{{ error }}</p>
  </ThemedCard>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useLocalization } from '@/locales/index.js'
import ThemedCard from '@/components/ThemedCard.vue'
import { decryptData, generateEncryptionKey } from '@/telegram/utils/crypto.js'

const { t } = useLocalization()

const props = defineProps({
  encryptedValue: {
    type: String,
    required: true
  }
})

const decryptedData = ref(null)
const error = ref(null)

// Форматирование даты
const formatDate = (timestamp) => {
  const date = new Date(timestamp * 1000)
  return date.toLocaleString('ru-RU', {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit'
  })
}

onMounted(() => {
  try {
    // Пытаемся дешифровать данные с разными возможными ключами
    // В реальной реализации здесь должна быть более сложная логика
    // для определения правильного ключа
    
    // Для демонстрации попробуем несколько вариантов временных меток
    const currentTime = Math.floor(Date.now() / 1000)
    
    // Попробуем дешифровать с разными временными метками в пределах 10 минут
    for (let i = 0; i < 600; i++) {
      const timestamp = currentTime - i
      // Здесь нам нужно знать player_id и session_id для генерации ключа
      // В реальной реализации эти данные могут быть закодированы в самом QR-коде
      // или передаваться отдельно
      
      // Для демонстрации предположим, что у нас есть эти данные
      // В реальной реализации это должно быть реализовано по-другому
      const playerId = 1 // Заглушка
      const sessionId = 1 // Заглушка
      
      const key = generateEncryptionKey(playerId, sessionId, timestamp)
      if (key) {
        const decrypted = decryptData(props.encryptedValue, key)
        if (decrypted) {
          try {
            const parsedData = JSON.parse(decrypted)
            if (parsedData.player_id && parsedData.session_id) {
              // Добавляем временную метку, если её нет
              if (!parsedData.timestamp) {
                parsedData.timestamp = timestamp
              }
              decryptedData.value = parsedData
              return
            }
          } catch (e) {
            // Игнорируем ошибки парсинга, продолжаем попытки
          }
        }
      }
    }
    
    error.value = t('scan_view.decryption_failed')
  } catch (e) {
    console.error('Ошибка дешифрования:', e)
    error.value = t('scan_view.decryption_error')
  }
})
</script>

<style scoped>
.qr-display-card {
  margin-top: 24px;
  width: 100%;
  max-width: 300px;
}

.decrypted-info {
  text-align: left;
  padding: 16px;
}

.decrypted-info p {
  margin: 8px 0;
  word-break: break-word;
}

.error-card {
  background-color: #ffebee;
  border: 1px solid #f44336;
}

.error-message {
  color: #f44336;
  font-weight: bold;
  text-align: center;
}
</style>