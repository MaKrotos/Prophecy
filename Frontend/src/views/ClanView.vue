<template>
  <div class="page">
    <div class="header-section">
      <h2 class="page-title">🏰 {{ t('clan_view.title') }}</h2>
      <p class="page-description">{{ t('clan_view.description') }}</p>
    </div>
    
    <div class="clan-content">
      <ThemedCard card-type="default" class="clan-card">
        <div class="clan-info">
          <div class="clan-icon">🏰</div>
          <h3>{{ t('clan_view.not_in_clan') }}</h3>
          <p class="clan-description">{{ t('clan_view.not_in_clan_description') }}</p>
        </div>
      </ThemedCard>
      
      <ThemedCard card-type="default" class="instructions-card">
        <h3>{{ t('clan_view.how_to_join_title') }}</h3>
        <p class="how-to-join-description">{{ t('clan_view.how_to_join_description') }}</p>
        <ol class="how-to-join-steps">
          <li>{{ t('clan_view.how_to_join_step1') }}</li>
          <li>{{ t('clan_view.how_to_join_step2') }}</li>
          <li>{{ t('clan_view.how_to_join_step3') }}</li>
        </ol>
      </ThemedCard>
      
      <ThemedCard card-type="default" class="instructions-card">
        <h3>{{ t('clan_view.benefits_title') }}</h3>
        <ul class="benefits-list">
          <li>{{ t('clan_view.benefit_1') }}</li>
          <li>{{ t('clan_view.benefit_2') }}</li>
          <li>{{ t('clan_view.benefit_3') }}</li>
          <li>{{ t('clan_view.benefit_4') }}</li>
        </ul>
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

.clan-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.clan-card,
.instructions-card {
  background: var(--tg-theme-secondary-bg-color, white);
  border-radius: 12px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}

.clan-card:hover,
.instructions-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.clan-info {
  text-align: center;
  margin-bottom: 24px;
}

.clan-icon {
  font-size: 3rem;
  margin-bottom: 16px;
}

.clan-info h3 {
  color: var(--tg-theme-text-color, #000000);
  font-size: 1.25rem;
  font-weight: 600;
  margin: 0 0 8px 0;
  transition: color 0.3s ease;
}

.clan-description {
  color: var(--tg-theme-hint-color, #666666);
  font-size: 0.9rem;
  margin: 0 0 16px 0;
  transition: color 0.3s ease;
}

.clan-actions {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.action-button {
  width: 100%;
  justify-content: center;
}

.instructions-card h3 {
  color: var(--tg-theme-text-color, #000000);
  font-size: 1.1rem;
  font-weight: 600;
  margin: 0 0 16px 0;
  transition: color 0.3s ease;
}

.benefits-list {
  color: var(--tg-theme-text-color, #333333);
  font-size: 0.95rem;
  line-height: 1.6;
  margin: 0;
  padding-left: 20px;
  transition: color 0.3s ease;
}

.benefits-list li {
  margin-bottom: 8px;
}

.benefits-list li:last-child {
  margin-bottom: 0;
}

/* Плавные переходы для всех элементов */
.page,
.page-title,
.page-description,
.clan-info h3,
.clan-description,
.instructions-card h3,
.benefits-list {
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
  
  .clan-actions {
    flex-direction: column;
  }
}
</style>