<template>
  <div class="telegram-only-layout">
    <div class="container">
      <div class="header">
        <div class="telegram-icon">
          <svg width="64" height="64" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M12 2C6.48 2 2 6.48 2 12C2 17.52 6.48 22 12 22C17.52 22 22 17.52 22 12C22 6.48 17.52 2 12 2Z"
              fill="#0088CC" />
            <path
              d="M16.64 8.08C16.49 8.03 16.33 8 16.17 8C15.86 8 15.56 8.1 15.31 8.28L8.94 12.65C8.24 13.11 7.74 13.85 7.56 14.69C7.38 15.53 7.53 16.41 7.99 17.11C8.45 17.81 9.19 18.31 10.03 18.49C10.87 18.67 11.75 18.52 12.45 18.06L18.82 13.69C19.32 13.35 19.64 12.82 19.69 12.23C19.74 11.64 19.5 11.06 19.05 10.65L16.64 8.08Z"
              fill="white" />
            <path
              d="M9.75 14.5C9.61 14.5 9.47 14.47 9.34 14.41C9.21 14.35 9.09 14.27 9 14.16C8.91 14.05 8.84 13.92 8.8 13.78C8.76 13.64 8.75 13.49 8.77 13.35C8.79 13.21 8.84 13.07 8.92 12.95C9 12.83 9.1 12.73 9.22 12.65L12.45 10.45C12.67 10.3 12.94 10.25 13.19 10.31C13.44 10.37 13.65 10.53 13.77 10.75C13.89 10.97 13.91 11.23 13.82 11.46C13.73 11.69 13.54 11.87 13.3 11.95L10.07 14.15C9.98 14.2 9.87 14.23 9.75 14.23V14.5Z"
              fill="#0088CC" />
          </svg>
        </div>
        <h1 class="title">Telegram WebApp</h1>
        <p class="subtitle">Это приложение работает только внутри Telegram</p>
      </div>

      <div class="content">
        <div class="info-card">
          <h2>Как использовать:</h2>
          <ol class="instructions">
            <li>Откройте приложение в Telegram</li>
            <li>Найдите нашего бота: <strong>@MyExampleBot</strong></li>
            <li>Запустите WebApp через команду /start или меню бота</li>
          </ol>
        </div>

        <div class="bot-link">
          <a href="https://t.me/MyExampleBot" target="_blank" class="telegram-button" @click="openTelegram">
            <span class="button-icon">📱</span>
            Перейти к боту
          </a>
        </div>

        <div class="features">
          <h3>Возможности приложения:</h3>
          <ul class="features-list">
            <li>✅ Удобный интерфейс оптимизирован для Telegram</li>
            <li>✅ Интеграция с Telegram Mini Apps</li>
            <li>✅ Безопасная авторизация</li>
            <li>✅ Мгновенная загрузка</li>
          </ul>
        </div>
      </div>

      <div class="footer">
        <p class="footer-text">
          Если у вас возникли проблемы с доступом, убедитесь что используете последнюю версию Telegram
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
const openTelegram = () => {
  // Можно добавить аналитику или дополнительную логику
  console.log('Redirecting to Telegram bot...');
};
</script>

<style scoped>
.telegram-only-layout {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  overflow-y: auto;
}

.container {
  max-width: 500px;
  width: 100%;
  background: white;
  border-radius: 20px;
  padding: 40px 30px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
  text-align: center;
  max-height: 90vh;
  overflow-y: auto;
}

.header {
  margin-bottom: 30px;
}

.telegram-icon {
  margin-bottom: 20px;
}

.title {
  color: #333;
  font-size: 2.5rem;
  font-weight: 700;
  margin-bottom: 10px;
}

.subtitle {
  color: #666;
  font-size: 1.1rem;
  margin-bottom: 0;
}

.content {
  margin: 30px 0;
}

.info-card {
  background: #f8f9fa;
  border-radius: 15px;
  padding: 25px;
  margin-bottom: 25px;
  text-align: left;
}

.info-card h2 {
  color: #333;
  margin-bottom: 15px;
  font-size: 1.3rem;
}

.instructions {
  color: #555;
  line-height: 1.6;
  padding-left: 20px;
}

.instructions li {
  margin-bottom: 8px;
}

.telegram-button {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  background: #0088cc;
  color: white;
  text-decoration: none;
  padding: 15px 30px;
  border-radius: 50px;
  font-weight: 600;
  font-size: 1.1rem;
  transition: all 0.3s ease;
  box-shadow: 0 4px 15px rgba(0, 136, 204, 0.3);
}

.telegram-button:hover {
  background: #0077b3;
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(0, 136, 204, 0.4);
}

.button-icon {
  font-size: 1.2rem;
}

.features {
  margin-top: 30px;
  text-align: left;
}

.features h3 {
  color: #333;
  margin-bottom: 15px;
  font-size: 1.2rem;
}

.features-list {
  color: #555;
  line-height: 1.6;
  list-style: none;
  padding: 0;
}

.features-list li {
  margin-bottom: 8px;
  padding-left: 0;
}

.footer {
  margin-top: 30px;
  padding-top: 20px;
  border-top: 1px solid #eee;
}

.footer-text {
  color: #888;
  font-size: 0.9rem;
  line-height: 1.4;
}

/* Адаптивность */
@media (max-width: 480px) {
  .container {
    padding: 30px 20px;
    margin: 10px;
  }

  .title {
    font-size: 2rem;
  }

  .info-card {
    padding: 20px;
  }

  .telegram-button {
    padding: 12px 25px;
    font-size: 1rem;
  }
}
</style>