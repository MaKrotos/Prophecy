<template>
  <div>
    <slot></slot>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useTelegramWebApp } from '../telegram/composables/useTelegramWebApp'
import { useLocalization } from '../locales/index.js'

const { t } = useLocalization()

const { 
  telegramUser,
  authHash,
  isTelegramReady,
  sendAuthToServer,
  hasValidToken,
  checkTelegramIdConsistency
} = useTelegramWebApp()

// Инициализация аутентификации
const initializeAuth = async () => {
  try {
    // Проверяем соответствие Telegram ID
    const isConsistent = checkTelegramIdConsistency()
    if (!isConsistent) {
      console.log('🔄 Необходима повторная авторизация из-за несоответствия Telegram ID')
    }

    // Авторизация только если Telegram готов и есть хэш
    if (isTelegramReady?.value && authHash.value) {
      console.log('📡 Отправка данных аутентификации на сервер...')
      await sendAuthToServer('auth/telegram', 3)
    } else if (!authHash.value) {
      console.warn('⚠️ Нет хэша аутентификации')
      throw new Error(t('auth.error.message'))
    }
  } catch (error) {
    console.error('❌ Ошибка инициализации аутентификации:', error)
    throw error
  }
}

// Повторная аутентификация
const retryAuth = async (endpoint = 'auth/telegram', maxRetries = 3) => {
  try {
    const result = await sendAuthToServer(endpoint, maxRetries)
    console.log('✅ Повторная авторизация успешна')
    return result
  } catch (error) {
    console.error('❌ Повторная авторизация не удалась после всех попыток:', error)
    throw error
  }
}

// Проверка состояния аутентификации
const isUserAuthenticated = () => {
  return hasValidToken() && telegramUser.value
}

// Следим за пользователем Telegram
onMounted(() => {
  if (telegramUser.value) {
    console.log('👤 Telegram user detected in AuthManager:', telegramUser.value)
  }
})

// Экспортируем методы для использования в родительских компонентах
defineExpose({
  initializeAuth,
  retryAuth,
  isUserAuthenticated
})
</script>