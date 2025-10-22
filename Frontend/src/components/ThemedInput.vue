<template>
  <div :class="['themed-input-container', { 'has-error': hasError }]">
    <label v-if="label" :for="inputId" class="input-label">{{ label }}</label>
    <input
      :id="inputId"
      :class="['themed-input', inputType]"
      :type="type"
      :placeholder="placeholder"
      :value="modelValue"
      :disabled="disabled"
      @input="handleInput"
      @blur="handleBlur"
      @focus="handleFocus"
    />
    <div v-if="hasError" class="error-message">{{ errorMessage }}</div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
  modelValue: {
    type: [String, Number],
    default: ''
  },
  label: {
    type: String,
    default: ''
  },
  type: {
    type: String,
    default: 'text'
  },
  placeholder: {
    type: String,
    default: ''
  },
  disabled: {
    type: Boolean,
    default: false
  },
  hasError: {
    type: Boolean,
    default: false
  },
  errorMessage: {
    type: String,
    default: ''
  },
  inputType: {
    type: String,
    default: 'default'
  }
})

const emit = defineEmits(['update:modelValue', 'blur', 'focus'])

const inputId = ref(`input-${Math.random().toString(36).substr(2, 9)}`)

const handleInput = (event) => {
  emit('update:modelValue', event.target.value)
}

const handleBlur = (event) => {
  emit('blur', event)
}

const handleFocus = (event) => {
  emit('focus', event)
}
</script>

<style scoped>
.themed-input-container {
  display: flex;
  flex-direction: column;
  gap: 8px;
  width: 100%;
}

.input-label {
  color: var(--tg-theme-text-color, #333333);
  font-weight: 500;
  font-size: 0.9rem;
  transition: color 0.3s ease;
}

.themed-input {
  padding: 12px 16px;
  border: 1px solid var(--tg-theme-hint-color, #cccccc);
  border-radius: 8px;
  background: var(--tg-theme-secondary-bg-color, white);
  color: var(--tg-theme-text-color, #333333);
  font-size: 1rem;
  transition: all 0.3s ease;
  outline: none;
}

.themed-input:focus {
  border-color: var(--tg-theme-button-color, #667eea);
  box-shadow: 0 0 0 2px rgba(102, 126, 234, 0.2);
}

.themed-input:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.themed-input.has-error,
.themed-input:invalid {
  border-color: #ff4757;
}

.error-message {
  color: #ff4757;
  font-size: 0.85rem;
  font-weight: 500;
}

/* Default input */
.themed-input.default {
  /* Uses default styles */
}

/* Large input */
.themed-input.large {
  padding: 16px 20px;
  font-size: 1.1rem;
}

/* Small input */
.themed-input.small {
  padding: 8px 12px;
  font-size: 0.9rem;
}

/* Rounded input */
.themed-input.rounded {
  border-radius: 24px;
}

/* Flat input */
.themed-input.flat {
  border: none;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.themed-input.flat:focus {
  box-shadow: 0 0 0 2px rgba(102, 126, 234, 0.2);
}
</style>