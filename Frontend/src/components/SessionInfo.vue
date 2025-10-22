<template>
  <div class="session-info">
    <!-- Session name (for SessionJoinView) -->
    <div v-if="showName" class="info-item">
      <span class="info-label">{{ t('session_join_view.session') }}:</span>
      <span class="info-value">{{ session.name }}</span>
    </div>

    <!-- Architect name -->
    <div v-if="session.architect_name" class="info-item">
      <span class="info-label">{{ t('session_detail_view.architect') }}:</span>
      <span class="info-value">{{ session.architect_name }}</span>
    </div>

    <!-- Created at -->
    <div v-if="session.created_at" class="info-item">
      <span class="info-label">{{ t('session_detail_view.created') }}:</span>
      <span class="info-value">{{ formatDate(session.created_at) }}</span>
    </div>

    <!-- Updated at -->
    <div v-if="showUpdatedAt && session.updated_at" class="info-item">
      <span class="info-label">{{ t('session_detail_view.updated') }}:</span>
      <span class="info-value">{{ formatDate(session.updated_at) }}</span>
    </div>

    <!-- Description -->
    <div v-if="showDescription" class="info-item">
      <span class="info-label">{{ t('sessions_view.description') }}:</span>
      <span class="info-value">{{ session.description || t('sessions_view.no_description') }}</span>
    </div>

    <!-- Player count -->
    <div v-if="showPlayerCount && session.player_count !== undefined" class="info-item">
      <span class="info-label">{{ t('sessions_view.players') }}:</span>
      <span class="info-value">{{ session.player_count }}</span>
    </div>
  </div>
</template>

<script setup>
import { useLocalization } from '@/locales/index.js'

const { t } = useLocalization()

const props = defineProps({
  session: {
    type: Object,
    required: true
  },
  showName: {
    type: Boolean,
    default: false
  },
  showDescription: {
    type: Boolean,
    default: false
  },
  showUpdatedAt: {
    type: Boolean,
    default: false
  },
  showPlayerCount: {
    type: Boolean,
    default: false
  }
})

// Форматирование даты
const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('ru-RU', {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}
</script>

<style scoped>
.session-info {
  display: flex;
  flex-direction: column;
  gap: 12px;
  background: var(--tg-theme-secondary-bg-color, white);
  padding: 16px;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.info-label {
  color: var(--tg-theme-hint-color, #666666);
  font-weight: 500;
  transition: color 0.3s ease;
}

.info-value {
  color: var(--tg-theme-text-color, #333333);
  font-weight: 500;
  transition: color 0.3s ease;
}

/* Плавные переходы для всех элементов */
.session-info,
.info-label,
.info-value {
  transition: all 0.3s ease;
}
</style>