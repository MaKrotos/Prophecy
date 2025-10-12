<template>
  <div class="scan-view">
    <h1>{{ t('scan_view.title') }}</h1>
    <qrcode-stream @decode="onDecode" @init="onInit"></qrcode-stream>
    <div v-if="scanned" class="result">
      <h2>{{ t('scan_view.scanned_qr') }}</h2>
      <p>{{ scanned }}</p>
      <qrcode-vue :value="scanned" :size="200" />
    </div>
    <div v-if="error" class="error">
      <p>{{ error }}</p>
    </div>
  </div>
</template>

<script>
import { QrcodeStream } from 'vue-qrcode-reader'
import QrcodeVue from 'qrcode.vue'
import { useLocalization } from '@/locales/index.js'

export default {
  name: 'ScanView',
  components: {
    QrcodeStream,
    QrcodeVue
  },
  setup() {
    const { t } = useLocalization()
    return { t }
  },
  data() {
    return {
      scanned: '',
      error: ''
    }
  },
  methods: {
    onDecode(result) {
      this.scanned = result
      this.error = ''
    },
    onInit(promise) {
      promise.catch(e => {
        this.error = t('scan_view.camera_error') + e.message
      })
    }
  }
}
</script>

<style scoped>
.scan-view {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 40px;
}
.result {
  margin-top: 32px;
  text-align: center;
}
.error {
  margin-top: 24px;
  color: #c00;
}
</style>