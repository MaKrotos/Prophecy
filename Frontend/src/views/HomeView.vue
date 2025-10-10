<template>
  <div class="page">
    <h2 class="page-title">Добро пожаловать! 👋</h2>
    <p class="page-description">Это главная страница приложения.</p>

    <!-- User Info Section -->
    <div class="user-info-section" v-if="userInfo">
      <div class="user-card">
        <div class="user-header">
          <h3 class="user-name">{{ userInfo.generatedName }}</h3>
          <div class="user-id">ID: {{ userInfo.id }}</div>
        </div>

        <div class="user-details">
          <div class="detail-item">
            <span class="detail-label">Telegram ID:</span>
            <span class="detail-value">{{ userInfo.telegramId }}</span>
          </div>

          <div class="detail-item" v-if="userInfo.isAdmin">
            <span class="admin-badge">Администратор</span>
          </div>
        </div>
      </div>
    </div>

    <div class="cards-grid">
      <div class="card" v-for="i in 6" :key="i">
        <div class="card-icon">📦</div>
        <h3 class="card-title">Карточка {{ i }}</h3>
        <p class="card-description">Пример контента карточки</p>
        <div class="card-footer">
          <span class="card-badge">Новое</span>
        </div>
      </div>
    </div>

    <div class="stats-section">
      <div class="stat-item">
        <div class="stat-value">42</div>
        <div class="stat-label">Элементов</div>
      </div>
      <div class="stat-item">
        <div class="stat-value">7</div>
        <div class="stat-label">Задач</div>
      </div>
      <div class="stat-item">
        <div class="stat-value">99%</div>
        <div class="stat-label">Готово</div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.page {
  padding: 16px;
  min-height: 100vh;
  background-color: var(--tg-theme-bg-color, #f5f5f5);
  transition: background-color 0.3s ease;
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
  margin-bottom: 24px;
  transition: color 0.3s ease;
}

.cards-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
  margin-bottom: 24px;
}

.card {
  background: var(--tg-theme-secondary-bg-color, white);
  padding: 20px;
  border-radius: 16px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  border: 1px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
  transition: all 0.3s ease;
  cursor: pointer;
}

.card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.12);
}

.card-icon {
  font-size: 2rem;
  margin-bottom: 12px;
}

.card-title {
  color: var(--tg-theme-text-color, #333333);
  font-size: 1.1rem;
  font-weight: 600;
  margin-bottom: 8px;
  transition: color 0.3s ease;
}

.card-description {
  color: var(--tg-theme-hint-color, #666666);
  font-size: 0.9rem;
  line-height: 1.4;
  margin-bottom: 12px;
  transition: color 0.3s ease;
}

.card-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: auto;
}

.card-badge {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 0.75rem;
  font-weight: 500;
}

.stats-section {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 12px;
  margin-top: 24px;
}

.stat-item {
  background: var(--tg-theme-secondary-bg-color, white);
  padding: 16px;
  border-radius: 12px;
  text-align: center;
  border: 1px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
  transition: all 0.3s ease;
}

.stat-value {
  color: var(--tg-theme-button-color, #667eea);
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 4px;
  transition: color 0.3s ease;
}

.stat-label {
  color: var(--tg-theme-hint-color, #666666);
  font-size: 0.8rem;
  transition: color 0.3s ease;
}

/* Адаптивность для мобильных устройств */
@media (max-width: 768px) {
  .cards-grid {
    grid-template-columns: 1fr;
    gap: 12px;
  }

  .card {
    padding: 16px;
  }

  .stats-section {
    grid-template-columns: repeat(3, 1fr);
    gap: 8px;
  }

  .stat-item {
    padding: 12px;
  }

  .stat-value {
    font-size: 1.25rem;
  }
}

@media (max-width: 480px) {
  .page {
    padding: 12px;
  }

  .page-title {
    font-size: 1.3rem;
  }

  .cards-grid {
    gap: 10px;
  }

  .stats-section {
    grid-template-columns: 1fr;
    gap: 8px;
  }
}

/* Специфические стили для темной темы */
:global(.tg-theme-dark) .card {
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
}

:global(.tg-theme-dark) .card:hover {
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.4);
}

:global(.tg-theme-dark) .stat-item {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
}

/* Плавные переходы для всех элементов */
.page,
.page-title,
.page-description,
.card,
.card-title,
.card-description,
.stat-item,
.stat-value,
.stat-label {
  transition: all 0.3s ease;
}

/* Анимация появления карточек */
.card {
  animation: fadeInUp 0.5s ease forwards;
  opacity: 0;
  transform: translateY(20px);
}

.card:nth-child(1) {
  animation-delay: 0.1s;
}

.card:nth-child(2) {
  animation-delay: 0.2s;
}

.card:nth-child(3) {
  animation-delay: 0.3s;
}

.card:nth-child(4) {
  animation-delay: 0.4s;
}

.card:nth-child(5) {
  animation-delay: 0.5s;
}

.card:nth-child(6) {
  animation-delay: 0.6s;
}

@keyframes fadeInUp {
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* User Info Section Styles */
.user-info-section {
  margin-bottom: 24px;
}

.user-card {
  background: var(--tg-theme-secondary-bg-color, white);
  border-radius: 16px;
  padding: 20px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  border: 1px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
  transition: all 0.3s ease;
}

.user-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  flex-wrap: wrap;
  gap: 8px;
}

.user-name {
  color: var(--tg-theme-text-color, #333333);
  font-size: 1.5rem;
  font-weight: 700;
  margin: 0;
  transition: color 0.3s ease;
}

.user-id {
  color: var(--tg-theme-hint-color, #666666);
  font-size: 0.9rem;
  font-weight: 500;
  transition: color 0.3s ease;
}

.user-details {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.detail-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.detail-label {
  color: var(--tg-theme-hint-color, #666666);
  font-size: 0.9rem;
  transition: color 0.3s ease;
}

.detail-value {
  color: var(--tg-theme-text-color, #333333);
  font-weight: 500;
  transition: color 0.3s ease;
}

.admin-badge {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
  padding: 6px 12px;
  border-radius: 16px;
  font-size: 0.9rem;
  font-weight: 600;
}

/* Dark theme adjustments */
:global(.tg-theme-dark) .user-card {
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .user-header {
    flex-direction: column;
    align-items: flex-start;
  }

  .user-name {
    font-size: 1.3rem;
  }

  .user-card {
    padding: 16px;
  }
}
</style>

<script setup>
import { ref, onMounted } from 'vue'
import { getUserInfoFromToken } from '/src/telegram/index.js'

const userInfo = ref(null)

onMounted(() => {
  // Получаем информацию о пользователе из JWT токена
  userInfo.value = getUserInfoFromToken()

  console.log('User Info:', userInfo.value)
})
</script>