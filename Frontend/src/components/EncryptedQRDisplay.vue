<template>
  <ThemedCard v-if="encryptedValue" class="qr-display-card">
    <h3>{{ t('scan_view.scanned_qr') }}</h3>
    <p class="scanned-value">{{ encryptedValue }}</p>
    <div class="qr-code-container">
      <qrcode-vue :value="encryptedValue" :size="200" />
    </div>
  </ThemedCard>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useLocalization } from '@/locales/index.js'
import ThemedCard from '@/components/ThemedCard.vue'
import QrcodeVue from 'qrcode.vue'
import { encryptData, generateEncryptionKey } from '@/telegram/utils/crypto.js'

const { t } = useLocalization()

const props = defineProps({
  playerId: {
    type: Number,
    required: true
  },
  sessionId: {
    type: Number,
    required: true
  },
  sessionName: {
    type: String,
    required: true
  }
})

const encryptedValue = ref(null)

onMounted(() => {
  // Генерируем временную метку на стороне фронта
  const timestamp = Math.floor(Date.now() / 1000)
  
  // Генерируем ключ шифрования
  const key = generateEncryptionKey(props.playerId, props.sessionId, timestamp)
  
  if (key) {
    // Формируем данные для шифрования
    const data = JSON.stringify({
      player_id: props.playerId,
      session_id: props.sessionId,
      timestamp: timestamp,
      session_name: props.sessionName
    })
    
    // Шифруем данные
    const encrypted = encryptData(data, key)
    if (encrypted) {
      encryptedValue.value = encrypted
    }
  }
})
</script>

<style scoped>
.qr-display-card {
  margin-top: 24px;
  text-align: center;
  width: 100%;
  max-width: 300px;
}

.scanned-value {
  word-break: break-all;
  margin: 16px 0;
  padding: 12px;
  background-color: var(--tg-theme-secondary-bg-color, #f5f5f5);
  border-radius: 8px;
  font-family: monospace;
}

.qr-code-container {
  display: flex;
  justify-content: center;
  margin-top: 16px;
}
</style>