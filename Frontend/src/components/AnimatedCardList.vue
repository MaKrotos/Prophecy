<template>
  <div class="animated-card-list">
    <slot name="header"></slot>
    
    <div class="cards-container">
      <AnimatedCard
        v-for="(item, index) in items"
        :key="getItemKey(item, index)"
        :index="index"
        :animation-delay="animationDelay"
        :custom-class="cardClass"
        :item-id="getItemKey(item, index)"
      >
        <slot name="card" :item="item" :index="index"></slot>
      </AnimatedCard>
    </div>
    
    <div v-if="loading" class="loading-indicator">
      <div class="spinner"></div>
      <p>{{ loadingText }}</p>
    </div>
    
    <div v-if="noMoreItems && !loading" class="no-more-items">
      <p>{{ noMoreItemsText }}</p>
    </div>
    
    <slot name="footer"></slot>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import AnimatedCard from './AnimatedCard.vue'

const props = defineProps({
  items: {
    type: Array,
    default: () => []
  },
  loading: {
    type: Boolean,
    default: false
  },
  noMoreItems: {
    type: Boolean,
    default: false
  },
  keyField: {
    type: String,
    default: 'id'
  },
  cardClass: {
    type: String,
    default: ''
  },
  animationDelay: {
    type: Number,
    default: 0.1
  },
  loadingText: {
    type: String,
    default: 'Загрузка...'
  },
  noMoreItemsText: {
    type: String,
    default: 'Все элементы загружены'
  }
})

const getItemKey = (item, index) => {
  return item[props.keyField] || index
}
</script>

<style scoped>
.animated-card-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.cards-container {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.loading-indicator {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 16px;
  color: var(--tg-theme-hint-color, #666666);
}

.spinner {
  width: 24px;
  height: 24px;
  border: 2px solid rgba(0, 0, 0, 0.1);
  border-top: 2px solid var(--tg-theme-button-color, #667eea);
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 8px;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.no-more-items {
  text-align: center;
  padding: 16px;
  color: var(--tg-theme-hint-color, #666666);
  font-style: italic;
  font-size: 0.9rem;
}
</style>