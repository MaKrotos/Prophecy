<template>
  <div :class="themeClass">
    <slot></slot>
  </div>
</template>

<script setup>
import { computed, onMounted, onUnmounted } from 'vue'
import { useTelegramWebAppSingleton } from '../telegram/composables/useTelegramWebAppSingleton'

const props = defineProps({
  applyToBody: {
    type: Boolean,
    default: false
  }
})

const { 
  isTelegram, 
  themeParams, 
  isDarkTheme,
  applyThemeToApp 
} = useTelegramWebAppSingleton()

// Вычисляем класс темы
const themeClass = computed(() => {
  return {
    'telegram-env': isTelegram.value,
    'tg-theme-dark': isDarkTheme.value,
    'tg-theme-light': !isDarkTheme.value
  }
})

// Обработчик изменения темы
const handleThemeChange = (theme) => {
  if (theme) {
    console.log('🎨 Theme changed in ThemeManager:', theme)
    applyThemeToApp(theme)
    
    // Обновляем заголовок и фон в Telegram WebApp
    updateTelegramAppColors(theme)
  }
}

// Функция для обновления цветов в Telegram WebApp
const updateTelegramAppColors = (theme) => {
  if (!isTelegram.value || !window.Telegram?.WebApp) return
  
  const webApp = window.Telegram.WebApp
  
  try {
    // Устанавливаем цвет заголовка
    if (theme.bg_color) {
      webApp.setHeaderColor(theme.bg_color)
    }
    
    // Устанавливаем цвет фона
    if (theme.bg_color) {
      webApp.setBackgroundColor(theme.bg_color)
    }
    
    console.log('📱 Telegram WebApp colors updated')
  } catch (error) {
    console.error('❌ Error updating Telegram WebApp colors:', error)
  }
}

// Добавляем глобальные обработчики для тем
onMounted(() => {
  if (isTelegram.value && window.Telegram?.WebApp) {
    const webApp = window.Telegram.WebApp
    
    // Обработчик изменения темы
    webApp.onEvent('themeChanged', handleThemeChange)
    
    // Обработчик изменения viewport
    webApp.onEvent('viewportChanged', () => {
      console.log('📱 Viewport changed:', {
        height: webApp.viewportHeight,
        stableHeight: webApp.viewportStableHeight,
        isExpanded: webApp.isExpanded
      })
    })
  }
})

// Удаляем обработчики при размонтировании
onUnmounted(() => {
  if (isTelegram.value && window.Telegram?.WebApp) {
    const webApp = window.Telegram.WebApp
    
    // Удаляем обработчики событий
    webApp.offEvent('themeChanged', handleThemeChange)
  }
})
</script>

<style scoped>
/* Стили для Telegram среды */
.telegram-env {
  background-color: var(--tg-theme-bg-color);
  color: var(--tg-theme-text-color);
}
</style>