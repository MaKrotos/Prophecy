<template>
  <div v-if="cameras.length > 1" class="camera-selector">
    <label for="camera-select">{{ t('scan_view.select_camera') }}</label>
    <select 
      id="camera-select" 
      :value="modelValue" 
      @change="$emit('update:modelValue', $event.target.value)"
      class="camera-select"
    >
      <option v-for="camera in cameras" :key="camera.deviceId" :value="camera.deviceId">
        {{ camera.label || `${t('scan_view.camera')} ${cameras.indexOf(camera) + 1}` }}
      </option>
    </select>
  </div>
</template>

<script setup>
import { useLocalization } from '@/locales/index.js'

const { t } = useLocalization()

defineProps({
  cameras: {
    type: Array,
    required: true
  },
  modelValue: {
    type: String,
    required: true
  }
})

defineEmits(['update:modelValue'])
</script>

<style scoped>
.camera-selector {
  margin-bottom: 20px;
  text-align: center;
  width: 100%;
  max-width: 400px;
}

.camera-selector label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
  color: var(--tg-theme-text-color, #333333);
}

.camera-select {
  width: 100%;
  padding: 12px 16px;
  border: 2px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
  border-radius: 8px;
  background-color: var(--tg-theme-bg-color, white);
  color: var(--tg-theme-text-color, #333333);
  font-size: 1rem;
  transition: border-color 0.3s ease;
}

.camera-select:focus {
  outline: none;
  border-color: var(--tg-theme-button-color, #667eea);
  box-shadow: 0 0 0 2px rgba(102, 126, 234, 0.2);
}
</style>