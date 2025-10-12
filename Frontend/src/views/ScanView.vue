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

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { QrcodeStream } from 'vue-qrcode-reader'
import QrcodeVue from 'qrcode.vue'
import { useLocalization } from '@/locales/index.js'

const { t } = useLocalization()

// Реактивные переменные
const scanned = ref('')
const error = ref('')
const successMessage = ref('')
const cameras = ref([])
const selectedCamera = ref('')
const cameraAccessRequested = ref(false)
const cameraAccessDenied = ref(false)

// Функция обработки декодирования QR-кода
const onDecode = (result) => {
  scanned.value = result
  error.value = ''
  successMessage.value = t('scan_view.success_message')
  
  // Скрываем сообщение об успехе через 3 секунды
  setTimeout(() => {
    successMessage.value = ''
  }, 3000)
}

// Функция инициализации камеры
const onInit = async (promise) => {
  try {
    // Проверяем, запрашивали ли мы уже доступ к камере
    if (cameraAccessRequested.value) {
      return
    }
    
    cameraAccessRequested.value = true
    
    await promise
    
    // Получаем список доступных камер
    const devices = await navigator.mediaDevices.enumerateDevices()
    cameras.value = devices.filter(device => device.kind === 'videoinput')
    
    // Если есть камеры, устанавливаем первую как выбранную по умолчанию
    if (cameras.value.length > 0) {
      selectedCamera.value = cameras.value[0].deviceId
    }
  } catch (e) {
    if (e.name === 'NotAllowedError' || e.name === 'PermissionDeniedError') {
      cameraAccessDenied.value = true
      error.value = t('scan_view.camera_access_denied')
    } else {
      error.value = t('scan_view.camera_error') + e.message
    }
    cameraAccessRequested.value = false
  }
}

// Функция изменения выбранной камеры
const onCameraChange = (deviceId) => {
  selectedCamera.value = deviceId
  // Здесь можно добавить логику переключения камеры, если библиотека поддерживает
}

// Очистка при размонтировании компонента
onUnmounted(() => {
  // Останавливаем камеру при уходе со страницы
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