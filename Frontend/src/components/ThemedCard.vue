<template>
  <div :class="['themed-card', cardType, { clickable: clickable }]">
    <slot></slot>
  </div>
</template>

<script setup>
const props = defineProps({
  cardType: {
    type: String,
    default: 'default'
  },
  clickable: {
    type: Boolean,
    default: false
  }
})
</script>

<style scoped>
.themed-card {
  background: var(--tg-theme-secondary-bg-color, white);
  border-radius: 12px;
  padding: 16px;
  box-shadow: 0 1px 6px rgba(0, 0, 0, 0.08);
  border: 1px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
  transition: all 0.3s ease;
}

.themed-card.clickable {
  cursor: pointer;
}

.themed-card.clickable:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.12);
}

.themed-card.clickable:active {
  transform: translateY(0);
}

/* Default card */
.themed-card.default {
  /* Uses default styles */
}

/* Primary card */
.themed-card.primary {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
  border: none;
}

/* Secondary card */
.themed-card.secondary {
  background: var(--tg-theme-secondary-bg-color, #f5f5f5);
}

/* Success card */
.themed-card.success {
  background: #d2ffe4;
  border-color: #2ed573;
}

/* Danger card */
.themed-card.danger {
  background: #ffeaea;
  border-color: #ff4757;
}

/* Info card */
.themed-card.info {
  background: #e3f2fd;
  border-color: #2196f3;
}

/* Large padding */
.themed-card.large-padding {
  padding: 24px;
}

/* Small padding */
.themed-card.small-padding {
  padding: 12px;
}

/* No padding */
.themed-card.no-padding {
  padding: 0;
}

/* Rounded */
.themed-card.rounded {
  border-radius: 24px;
}

/* Flat */
.themed-card.flat {
  box-shadow: none;
  border: none;
}

/* Dark theme adjustments */
:global(.tg-theme-dark) .themed-card {
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
}

:global(.tg-theme-dark) .themed-card.clickable:hover {
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.4);
}
</style>