<template>
  <div class="page">
    <div class="header-section">
      <h2 class="page-title">🎭 {{ t('session_role_view.title') }}</h2>
      <p class="page-description">{{ t('session_role_view.description') }}</p>
    </div>
    
    <div class="role-content">
      <ThemedCard card-type="default" class="role-card">
        <div class="role-header">
          <h3>{{ t('session_role_view.your_role') }}</h3>
          <div class="role-badge">Игрок</div>
        </div>
        <div class="role-description">
          <p>{{ t('session_role_view.role_description') }}</p>
        </div>
      </ThemedCard>
      
      <ThemedCard card-type="default" class="character-card">
        <div class="character-header">
          <h3>{{ t('session_role_view.character') }}</h3>
          <div class="character-name">Загадочный незнакомец</div>
        </div>
        <div class="character-description">
          <p>{{ t('session_role_view.character_description') }}</p>
        </div>
      </ThemedCard>
    </div>
  </div>
</template>

<script setup>
import { useLocalization } from '@/locales/index.js'
import ThemedCard from '../components/ThemedCard.vue'

const { t } = useLocalization()
</script>

<style scoped>
.page {
  padding: 16px;
  background-color: var(--tg-theme-bg-color, #f5f5f5);
  transition: background-color 0.3s ease;
}

.header-section {
  margin-bottom: 24px;
}

.page-title {
  color: var(--tg-theme-text-color, #000000);
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 8px;
  transition: color 0.3s ease;
}

.page-description {
  color: var(--tg-theme-hint-color, #666666);
  font-size: 1rem;
  margin-bottom: 16px;
  transition: color 0.3s ease;
}

.role-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.role-card,
.character-card {
  background: var(--tg-theme-secondary-bg-color, white);
  border-radius: 12px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}

.role-card:hover,
.character-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.role-header,
.character-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  flex-wrap: wrap;
  gap: 12px;
}

.role-header h3,
.character-header h3 {
  color: var(--tg-theme-text-color, #000000);
  font-size: 1.25rem;
  font-weight: 600;
  margin: 0;
  transition: color 0.3s ease;
}

.role-badge {
  background: var(--tg-theme-button-color, #667eea);
  color: white;
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 0.875rem;
  font-weight: 500;
}

.character-name {
  background: var(--tg-theme-hint-color, #666666);
  color: white;
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 0.875rem;
  font-weight: 500;
}

.role-description p,
.character-description p {
  color: var(--tg-theme-text-color, #333333);
  font-size: 1rem;
  line-height: 1.5;
  margin: 0;
  transition: color 0.3s ease;
}

/* Плавные переходы для всех элементов */
.page,
.page-title,
.page-description,
.role-header h3,
.character-header h3,
.role-description p,
.character-description p {
  transition: all 0.3s ease;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .page {
    padding: 12px;
  }
  
  .page-title {
    font-size: 1.3rem;
  }
  
  .role-header h3,
  .character-header h3 {
    font-size: 1.1rem;
  }
}
</style>