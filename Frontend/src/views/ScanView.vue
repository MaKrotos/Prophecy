<template>
  <div class="scan-view">
    <h1>{{ t('scan_view.title') }}</h1>
    
    <!-- Селектор камеры -->
    <div v-if="cameras.length > 1" class="camera-selector">
      <label for="camera-select">{{ t('scan_view.select_camera') }}</label>
      <select id="camera-select" v-model="selectedCamera" @change="onCameraChange">
        <option v-for="camera in cameras" :key="camera.deviceId" :value="camera.deviceId">
          {{ camera.label || `Camera ${cameras.indexOf(camera) + 1}` }}
        </option>
      </select>
    </div>
    
    <!-- Поток QR-кода с привязкой к выбранной камере -->
    <qrcode-stream 
      :camera="selectedCamera"
      @decode="onDecode" 
      @init="onInit"
      @error="onCameraError"
    ></qrcode-stream>
    
    <!-- Сообщения -->
    <div v-if="successMessage" class="success">
      <p>{{ successMessage }}</p>
    </div>
    
    <div v-if="scanned" class="result">
      <h2>{{ t('scan_view.scanned_qr') }}</h2>
      <p>{{ scanned }}</p>
      <qrcode-vue :value="scanned" :size="200" />
    </div>
    
    <div v-if="error" class="error">
      <p>{{ error }}</p>
      <button v-if="cameraAccessDenied" @click="retryCameraAccess">
        {{ t('scan_view.retry_camera') }}
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { QrcodeStream } from 'vue-qrcode-reader'
import QrcodeVue from 'qrcode.vue'
import { useLocalization } from '@/locales/index.js'

const { t } = useLocalization()

const scanned = ref('')
const error = ref('')
const successMessage = ref('')
const cameras = ref([])
const selectedCamera = ref('')
const cameraAccessRequested = ref(false)
const cameraAccessDenied = ref(false)

const onDecode = (result) => {
  scanned.value = result
  error.value = ''
  successMessage.value = t('scan_view.success_message')
  
  setTimeout(() => {
    successMessage.value = ''
  }, 3000)
}

const onCameraChange = () => {
  // При смене камеры сбрасываем состояние
  scanned.value = ''
  error.value = ''
  cameraAccessRequested.value = false
}

const onCameraError = (error) => {
  console.error('Camera stream error:', error)
  error.value = t('scan_view.camera_stream_error')
}

const retryCameraAccess = () => {
  cameraAccessDenied.value = false
  cameraAccessRequested.value = false
  error.value = ''
}

const onInit = async (promise) => {
  try {
    if (cameraAccessRequested.value) {
      return
    }
    
    cameraAccessRequested.value = true
    error.value = ''
    
    await promise
    
    // Получаем список камер
    const devices = await navigator.mediaDevices.enumerateDevices()
    cameras.value = devices.filter(device => device.kind === 'videoinput')
    
    if (cameras.value.length > 0) {
      selectedCamera.value = cameras.value[0].deviceId
    } else {
      error.value = t('scan_view.no_cameras_found')
    }
    
  } catch (e) {
    console.error('Camera init error:', e)
    cameraAccessRequested.value = false
    
    if (e.name === 'NotAllowedError') {
      cameraAccessDenied.value = true
      error.value = t('scan_view.camera_access_denied')
    } else if (e.name === 'NotFoundError') {
      error.value = t('scan_view.camera_not_found')
    } else {
      error.value = `${t('scan_view.camera_error')}: ${e.message}`
    }
  }
}

onUnmounted(() => {
  scanned.value = ''
  error.value = ''
  successMessage.value = ''
  cameraAccessRequested.value = false
  cameraAccessDenied.value = false
})
</script>

<style scoped>
.scan-view {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 40px;
  padding: 20px;
}

.camera-selector {
  margin-bottom: 20px;
  text-align: center;
}

.camera-selector label {
  display: block;
  margin-bottom: 8px;
}

.camera-selector select {
  padding: 8px 12px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.result {
  margin-top: 32px;
  text-align: center;
  padding: 20px;
  border: 1px solid #4CAF50;
  border-radius: 8px;
  background-color: #f8fff8;
}

.success {
  margin-top: 20px;
  padding: 10px 20px;
  background-color: #4CAF50;
  color: white;
  border-radius: 4px;
}

.error {
  margin-top: 24px;
  padding: 15px;
  color: #d32f2f;
  background-color: #ffebee;
  border: 1px solid #f44336;
  border-radius: 4px;
  text-align: center;
}

.error button {
  margin-top: 10px;
  padding: 8px 16px;
  background-color: #d32f2f;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.error button:hover {
  background-color: #c62828;
}
</style>