<template>
    <div class="auth-error-layout">
        <div class="container">
            <div class="header">
                <div class="error-icon">
                    <svg width="64" height="64" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path
                            d="M12 2C6.48 2 2 6.48 2 12C2 17.52 6.48 22 12 22C17.52 22 22 17.52 22 12C22 6.48 17.52 2 12 2ZM13 17H11V15H13V17ZM13 13H11V7H13V13Z"
                            fill="#FF6B6B" />
                    </svg>
                </div>
                <h1 class="title">Ошибка авторизации</h1>
                <p class="subtitle">Не удалось выполнить вход через Telegram</p>
            </div>

            <div class="content">
                <div class="error-card">
                    <h2>Что произошло:</h2>
                    <p class="error-message">{{ errorMessage }}</p>
                </div>

                <div class="actions">
                    <button class="retry-button" @click="retryAuth">
                        Попробовать снова
                    </button>

                    <button class="later-button" @click="tryLater">
                        Попробовать позже
                    </button>
                </div>

                <div class="info">
                    <h3>Возможные причины:</h3>
                    <ul class="reasons-list">
                        <li>🔧 Временные технические работы на сервере</li>
                        <li>🔒 Проблемы с безопасностью Telegram WebApp</li>
                        <li>🌐 Нестабильное интернет-соединение</li>
                        <li>🔄 Устаревшая версия приложения</li>
                    </ul>
                </div>
            </div>

            <div class="footer">
                <p class="footer-text">
                    Если проблема сохраняется, попробуйте обновить Telegram или обратиться в поддержку
                </p>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
    errorMessage: {
        type: String,
        default: 'Не удалось получить токен авторизации от сервера'
    }
})

const emit = defineEmits(['retry', 'tryLater'])

const retryAuth = () => {
    emit('retry')
}

const tryLater = () => {
    emit('tryLater')
}
</script>

<style scoped>
.auth-error-layout {
    min-height: 100vh;
    background: linear-gradient(135deg, #ff6b6b 0%, #ee5a24 100%);
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
}

.header {
    margin-bottom: 30px;
}

.error-icon {
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

.error-card {
    background: #fff5f5;
    border-radius: 15px;
    padding: 25px;
    margin-bottom: 25px;
    text-align: left;
    border: 1px solid #fed7d7;
}

.error-card h2 {
    color: #333;
    margin-bottom: 15px;
    font-size: 1.3rem;
}

.error-message {
    color: #e53e3e;
    font-weight: 500;
    line-height: 1.6;
}

.actions {
    display: flex;
    flex-direction: column;
    gap: 15px;
    margin-bottom: 30px;
}

.retry-button {
    background: #4299e1;
    color: white;
    border: none;
    padding: 15px 30px;
    border-radius: 50px;
    font-weight: 600;
    font-size: 1.1rem;
    cursor: pointer;
    transition: all 0.3s ease;
    box-shadow: 0 4px 15px rgba(66, 153, 225, 0.3);
}

.retry-button:hover {
    background: #3182ce;
    transform: translateY(-2px);
    box-shadow: 0 6px 20px rgba(66, 153, 225, 0.4);
}

.later-button {
    background: #edf2f7;
    color: #4a5568;
    border: none;
    padding: 15px 30px;
    border-radius: 50px;
    font-weight: 600;
    font-size: 1.1rem;
    cursor: pointer;
    transition: all 0.3s ease;
}

.later-button:hover {
    background: #e2e8f0;
    transform: translateY(-2px);
}

.info {
    text-align: left;
    background: #f8f9fa;
    border-radius: 15px;
    padding: 25px;
}

.info h3 {
    color: #333;
    margin-bottom: 15px;
    font-size: 1.2rem;
}

.reasons-list {
    color: #555;
    line-height: 1.6;
    list-style: none;
    padding: 0;
}

.reasons-list li {
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

    .error-card {
        padding: 20px;
    }

    .actions {
        gap: 10px;
    }

    .retry-button,
    .later-button {
        padding: 12px 25px;
        font-size: 1rem;
    }
}
</style>