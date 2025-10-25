<template>
  <ThemedCard v-if="value" class="qr-display-card">
    <h3>{{ t('scan_view.scanned_qr') }}</h3>
    <p class="scanned-value">{{ value }}</p>
    <div class="qr-code-container">
      <qrcode-vue :value="value" :size="200" />
    </div>
  </ThemedCard>
</template>

<script setup>
import { useLocalization } from '@/locales/index.js'
import ThemedCard from '@/components/ThemedCard.vue'
import QrcodeVue from 'qrcode.vue'

const { t } = useLocalization()

defineProps({
  value: {
    type: String,
    required: true
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