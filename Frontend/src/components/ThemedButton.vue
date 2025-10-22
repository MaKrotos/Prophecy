<template>
  <button 
    :class="['themed-button', buttonType, { disabled: disabled }]"
    :disabled="disabled"
    @click="handleClick"
  >
    <slot></slot>
  </button>
</template>

<script setup>
const props = defineProps({
  buttonType: {
    type: String,
    default: 'primary'
  },
  disabled: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['click'])

const handleClick = (event) => {
  if (!props.disabled) {
    emit('click', event)
  }
}
</script>

<style scoped>
.themed-button {
  border: none;
  border-radius: 8px;
  padding: 12px 20px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.themed-button:hover:not(.disabled) {
  opacity: 0.9;
  transform: translateY(-1px);
}

.themed-button:active:not(.disabled) {
  transform: translateY(0);
}

.themed-button.disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

/* Primary button */
.themed-button.primary {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
}

/* Secondary button */
.themed-button.secondary {
  background: var(--tg-theme-secondary-bg-color, #f0f0f0);
  color: var(--tg-theme-text-color, #333333);
}

/* Danger button */
.themed-button.danger {
  background: #ff4757;
  color: white;
}

/* Success button */
.themed-button.success {
  background: #2ed573;
  color: white;
}

/* Icon button */
.themed-button.icon {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
  border-radius: 50%;
  width: 32px;
  height: 32px;
  padding: 0;
  font-size: 16px;
}

/* Small button */
.themed-button.small {
  padding: 8px 12px;
  font-size: 0.9rem;
}

/* Large button */
.themed-button.large {
  padding: 16px 24px;
  font-size: 1.1rem;
}

/* Full width button */
.themed-button.full-width {
  width: 100%;
}
</style>