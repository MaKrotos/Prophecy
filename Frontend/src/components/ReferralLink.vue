<template>
  <div class="referral-link-container">
    <span class="referral-link">{{ referralLink }}</span>
    <button class="copy-button" @click="copyReferralLink" :title="t('sessions_view.copy_link')">
      📋
    </button>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useLocalization } from '@/locales/index.js'

const { t } = useLocalization()

const props = defineProps({
  session: {
    type: Object,
    required: true
  },
  actionPrefix: {
    type: String,
    default: ''
  }
})

// Получение реферальной ссылки для сессии
const referralLink = computed(() => {
  // Используем имя бота из переменных окружения или значение по умолчанию
  const botUsername = import.meta.env.VITE_TELEGRAM_BOT_USERNAME || 'your_bot_username'
  // Формируем ссылку в формате Telegram Mini App
  const prefix = props.actionPrefix ? `${props.actionPrefix}_` : ''
  return `https://t.me/${botUsername}?startapp=${prefix}${props.session.referral_link}`
})

// Копирование реферальной ссылки в буфер обмена
const copyReferralLink = () => {
  const link = referralLink.value

  // Используем Clipboard API, если доступно
  if (navigator.clipboard) {
    navigator.clipboard.writeText(link).then(() => {
      if (window.Telegram && window.Telegram.WebApp) {
        window.Telegram.WebApp.showAlert(t('sessions_view.link_copied'))
      }
    }).catch(err => {
      console.error('Не удалось скопировать ссылку: ', err)
      // Альтернативный метод копирования
      fallbackCopyTextToClipboard(link)
    })
  } else {
    // Альтернативный метод копирования для старых браузеров
    fallbackCopyTextToClipboard(link)
  }
}

// Альтернативный метод копирования текста в буфер обмена
const fallbackCopyTextToClipboard = (text) => {
  const textArea = document.createElement("textarea")
  textArea.value = text

  // Избегаем прокрутки страницы
  textArea.style.top = "0"
  textArea.style.left = "0"
  textArea.style.position = "fixed"
  textArea.style.opacity = "0"

  document.body.appendChild(textArea)
  textArea.focus()
  textArea.select()

  try {
    const successful = document.execCommand('copy')
    if (successful) {
      if (window.Telegram && window.Telegram.WebApp) {
        window.Telegram.WebApp.showAlert(t('sessions_view.link_copied'))
      }
    } else {
      console.error('Не удалось скопировать ссылку')
      if (window.Telegram && window.Telegram.WebApp) {
        window.Telegram.WebApp.showAlert(t('sessions_view.copy_error'))
      }
    }
  } catch (err) {
    console.error('Ошибка при копировании ссылки: ', err)
    if (window.Telegram && window.Telegram.WebApp) {
      window.Telegram.WebApp.showAlert(t('sessions_view.copy_error'))
    }
  }

  document.body.removeChild(textArea)
}
</script>

<style scoped>
.referral-link-container {
  display: flex;
  align-items: center;
  gap: 8px;
  width: 100%;
}

.referral-link {
  flex: 1;
  word-break: break-all;
  font-size: 0.9rem;
  color: var(--tg-theme-text-color, #333333);
  background: var(--tg-theme-bg-color, #f5f5f5);
  padding: 8px 12px;
  border-radius: 8px;
}

.copy-button {
  background: var(--tg-theme-button-color, #667eea);
  color: var(--tg-theme-button-text-color, white);
  border: none;
  border-radius: 4px;
  padding: 8px 12px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.3s ease;
  flex-shrink: 0;
}

.copy-button:hover {
  opacity: 0.8;
}
</style>