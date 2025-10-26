<template>
  <div class="scan-view">
    <h1>{{ t('scan_view.title') }}</h1>
    
    <!-- Индикатор загрузки камеры -->
    <CameraLoadingIndicator :loading="cameraInitializing" />
    
    <!-- Селектор камеры -->
    <CameraSelector
      v-model="selectedCamera"
      :cameras="cameras"
      @update:modelValue="onCameraChange"
    />
    
    <!-- Поток QR-кода с привязкой к выбранной камере -->
    <div class="scanner-container">
      <qrcode-stream
        :camera="selectedCamera"
        @decode="onDecode"
        @init="onInit"
        @error="onCameraError"
        class="scanner-stream"
        v-show="!cameraInitializing"
      ></qrcode-stream>
    </div>
    
    <!-- Сообщения -->
    <ScanMessages
      :success-message="successMessage"
      :error="error"
      :show-retry-button="cameraAccessDenied"
      @retry="retryCameraAccess"
    />
    
    <!-- Отображение отсканированного QR-кода -->
    <QRDisplay :value="scanned" />
  </div>
</template>

<script setup>
import { ref, onUnmounted } from 'vue'
import { QrcodeStream } from 'vue-qrcode-reader'
import { useLocalization } from '@/locales/index.js'
import CameraSelector from '@/components/CameraSelector.vue'
import QRDisplay from '@/components/QRDisplay.vue'
import ScanMessages from '@/components/ScanMessages.vue'
import CameraLoadingIndicator from '@/components/CameraLoadingIndicator.vue'

const { t } = useLocalization()

const scanned = ref('')
const error = ref('')
const successMessage = ref('')
const cameras = ref([])
const selectedCamera = ref('')
const cameraAccessRequested = ref(false)
const cameraAccessDenied = ref(false)
const cameraInitializing = ref(false)

const onDecode = (result) => {
  console.log('QR Code decoded:', result)
  if (!result) {
    console.warn('Empty result received')
    return
  }
  
  // Проверяем, является ли результат строкой
  if (typeof result !== 'string') {
    console.warn('Result is not a string:', result)
    // Если это объект, пытаемся получить текст
    if (result && typeof result === 'object' && result.text) {
      scanned.value = result.text
    } else {
      scanned.value = JSON.stringify(result)
    }
  } else {
    scanned.value = result
  }
  
  error.value = ''
  successMessage.value = t('scan_view.success_message')
  console.log('QR Code successfully scanned:', scanned.value)
  
  setTimeout(() => {
    successMessage.value = ''
  }, 3000)
}

const onCameraChange = () => {
  console.log('Camera changed')
  // При смене камеры сбрасываем состояние
  scanned.value = ''
  error.value = ''
  cameraAccessRequested.value = false
  cameraInitializing.value = false
}

const onCameraError = (error) => {
  console.error('Camera stream error:', error)
  console.error('Error details:', {
    name: error.name,
    message: error.message,
    stack: error.stack
  })
  
  // Проверяем тип ошибки и устанавливаем соответствующее сообщение
  if (error.name === 'NotAllowedError') {
    error.value = t('scan_view.camera_access_denied')
  } else if (error.name === 'NotFoundError') {
    error.value = t('scan_view.camera_not_found')
  } else if (error.name === 'NotReadableError') {
    error.value = t('scan_view.camera_not_readable')
  } else if (error.name === 'OverconstrainedError') {
    error.value = t('scan_view.camera_overconstrained')
  } else if (error.name === 'StreamApiNotSupportedError') {
    error.value = t('scan_view.stream_api_not_supported')
  } else {
    error.value = `${t('scan_view.camera_stream_error')}: ${error.message}`
  }
}

const retryCameraAccess = () => {
  console.log('Retrying camera access')
  cameraAccessDenied.value = false
  cameraAccessRequested.value = false
  cameraInitializing.value = false
  error.value = ''
}

const onInit = async (promise) => {
  try {
    if (cameraAccessRequested.value) {
      return
    }
    
    cameraAccessRequested.value = true
    cameraInitializing.value = true
    error.value = ''
    
    await promise
    console.log('Camera initialized successfully')
    
    // Получаем список камер
    const devices = await navigator.mediaDevices.enumerateDevices()
    console.log('Available devices:', devices)
    cameras.value = devices.filter(device => device.kind === 'videoinput')
    console.log('Video input devices:', cameras.value)
    
    if (cameras.value.length > 0) {
      selectedCamera.value = cameras.value[0].deviceId
      console.log('Selected camera:', cameras.value[0])
    } else {
      error.value = t('scan_view.no_cameras_found')
      console.warn('No cameras found')
      // Принудительно скрываем индикатор загрузки, если нет камер
      cameraInitializing.value = false
    }
    
  } catch (e) {
    console.error('Camera init error:', e)
    console.error('Error details:', {
      name: e.name,
      message: e.message,
      stack: e.stack
    })
    cameraAccessRequested.value = false
    
    if (e.name === 'NotAllowedError') {
      cameraAccessDenied.value = true
      error.value = t('scan_view.camera_access_denied')
      // Принудительно скрываем индикатор загрузки при отсутствии доступа
      cameraInitializing.value = false
    } else if (e.name === 'NotFoundError') {
      error.value = t('scan_view.camera_not_found')
      // Принудительно скрываем индикатор загрузки при отсутствии камеры
      cameraInitializing.value = false
    } else if (e.name === 'NotReadableError') {
      error.value = t('scan_view.camera_not_readable')
      // Принудительно скрываем индикатор загрузки при недоступности камеры
      cameraInitializing.value = false
    } else if (e.name === 'OverconstrainedError') {
      error.value = t('scan_view.camera_overconstrained')
      // Принудительно скрываем индикатор загрузки при неподходящей камере
      cameraInitializing.value = false
    } else if (e.name === 'StreamApiNotSupportedError') {
      error.value = t('scan_view.stream_api_not_supported')
      // Принудительно скрываем индикатор загрузки при отсутствии поддержки
      cameraInitializing.value = false
    } else {
      error.value = `${t('scan_view.camera_error')}: ${e.message}`
      // Принудительно скрываем индикатор загрузки при других ошибках
      cameraInitializing.value = false
    }
  } finally {
    // Добавляем небольшую задержку перед скрытием индикатора загрузки
    setTimeout(() => {
      cameraInitializing.value = false
    }, 100)
  }
}

onUnmounted(() => {
  console.log('ScanView component unmounted')
  scanned.value = ''
  error.value = ''
  successMessage.value = ''
  cameraAccessRequested.value = false
  cameraAccessDenied.value = false
  cameraInitializing.value = false
})
</script>

<style scoped>
.scan-view {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 20px;
  padding: 20px;
  width: 100%;
}

.scanner-container {
  width: 100%;
  max-width: 500px;
  margin: 20px 0;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.scanner-stream {
  width: 100%;
  height: auto;
  aspect-ratio: 4/3;
}
</style>