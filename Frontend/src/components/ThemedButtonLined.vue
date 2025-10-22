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
  border: 3px solid var(--tg-theme-button-color, #667eea);
  border-radius: 12px;
  padding: 12px 20px;
  font-size: 1rem;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.2s ease;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  background: transparent;
  color: var(--tg-theme-button-color, #667eea);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.themed-button:hover:not(.disabled) {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.themed-button:active:not(.disabled) {
  transform: translateY(0);
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.themed-button.disabled {
  opacity: 0.5;
  cursor: not-allowed;
  box-shadow: none;
}

/* Primary button */
.themed-button.primary {
  border-color: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-color, #667eea);
  background: transparent;
}

.themed-button.primary:hover:not(.disabled) {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
}

/* Secondary button */
.themed-button.secondary {
  border-color: var(--tg-theme-hint-color, #999999);
  color: var(--tg-theme-hint-color, #999999);
  background: transparent;
}

.themed-button.secondary:hover:not(.disabled) {
  background: var(--tg-theme-hint-color, #999999);
  color: white;
}

/* Danger button */
.themed-button.danger {
  border-color: #ff4757;
  color: #ff4757;
  background: transparent;
}

.themed-button.danger:hover:not(.disabled) {
  background: #ff4757;
  color: white;
}

/* Success button */
.themed-button.success {
  border-color: #2ed573;
  color: #2ed573;
  background: transparent;
}

.themed-button.success:hover:not(.disabled) {
  background: #2ed573;
  color: white;
}

/* Icon button */
.themed-button.icon {
  border: 3px solid var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-color, #667eea);
  background: transparent;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  padding: 0;
  font-size: 18px;
}

.themed-button.icon:hover:not(.disabled) {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
}

/* Small button */
.themed-button.small {
  padding: 8px 12px;
  font-size: 0.85rem;
  border-width: 2px;
}

.themed-button.small.icon {
  width: 32px;
  height: 32px;
  font-size: 14px;
  border-width: 2px;
}

/* Large button */
.themed-button.large {
  padding: 16px 28px;
  font-size: 1.2rem;
  border-width: 4px;
}

.themed-button.large.icon {
  width: 48px;
  height: 48px;
  font-size: 20px;
  border-width: 4px;
}

/* Full width button */
.themed-button.full-width {
  width: 100%;
}
</style>