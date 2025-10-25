<template>
  <div class="session-card-minimal" @click="onClick">
    <div class="session-header-minimal">
      <h3 class="session-name-minimal">{{ session.name }}</h3>
      <div class="session-meta">
        <span class="session-architect">{{ session.architect_name || 'Архитектор' }}</span>
        <span class="session-date">{{ formatDate(session.created_at) }}</span>
      </div>
    </div>
    
    <div class="session-description" v-if="session.description">
      {{ session.description }}
    </div>
  </div>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue'

const props = defineProps({
  session: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['click'])

// Форматирование даты
const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('ru-RU', {
    day: '2-digit',
    month: '2-digit',
    year: '2-digit'
  })
}

// Обработчик клика
const onClick = () => {
  emit('click', props.session)
}
</script>

<style scoped>
.session-card-minimal {
  background: var(--tg-theme-secondary-bg-color, #ffffff);
  border-radius: 12px;
  padding: 16px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
  cursor: pointer;
  border: 1px solid var(--tg-theme-hint-color, #e0e0e0);
}

.session-card-minimal:hover {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  transform: translateY(-2px);
}

.session-header-minimal {
  margin-bottom: 12px;
}

.session-name-minimal {
  color: var(--tg-theme-text-color, #333333);
  font-size: 1.1rem;
  font-weight: 600;
  margin: 0 0 8px 0;
  transition: color 0.3s ease;
}

.session-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 0.8rem;
}

.session-architect {
  color: var(--tg-theme-hint-color, #666666);
  font-weight: 500;
}

.session-date {
  color: var(--tg-theme-hint-color, #999999);
  font-weight: 400;
}

.session-description {
  color: var(--tg-theme-hint-color, #666666);
  font-size: 0.9rem;
  line-height: 1.4;
  margin-top: 8px;
}

/* Плавные переходы для всех элементов */
.session-name-minimal {
  transition: all 0.3s ease;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .session-card-minimal {
    padding: 12px;
  }
  
  .session-name-minimal {
    font-size: 1rem;
  }
  
  .session-meta {
    flex-direction: column;
    align-items: flex-start;
    gap: 4px;
  }
}

/* Dark theme adjustments */
:global(.tg-theme-dark) .session-card-minimal {
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.15);
  border-color: var(--tg-theme-hint-color, #444444);
}

:global(.tg-theme-dark) .session-card-minimal:hover {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.25);
}
</style>