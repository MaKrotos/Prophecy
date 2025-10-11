<template>
  <div class="app" :class="{ 
    'telegram-env': isTelegram,
    'tg-theme-dark': isDarkTheme,
    'tg-theme-light': !isDarkTheme
  }">
    <HeaderNav :title="pageTitle" />
    
    <main class="main-content">
      <router-view v-slot="{ Component }">
        <transition :name="transitionName" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </main>
    
    <BottomNav />
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import HeaderNav from '../components/HeaderNav.vue'
import BottomNav from '../components/BottomNav.vue'
import UserStats from '../components/UserStats.vue'
import { useTelegramWebApp } from '../telegram/composables/useTelegramWebApp'
import { getUserInfoFromToken } from '../telegram/auth/user'

const route = useRoute()
const router = useRouter()
const transitionName = ref('slide-left')

// Используем рефакторизованный композабл
const { 
  isTelegram, 
  themeParams, 
  isDarkTheme,
  telegramUser,
  applyThemeToApp,
  refreshTheme,
  sendAuthToServer
} = useTelegramWebApp()

// Проверяем, является ли пользователь администратором
const userInfo = computed(() => getUserInfoFromToken())
const isAdmin = computed(() => userInfo.value?.isAdmin || false)

// Получаем порядок маршрутов из meta данных
const getRoutesOrder = () => {
  return router.getRoutes()
    .filter(route => route.meta?.order && !route.meta?.isNested)
    .sort((a, b) => a.meta.order - b.meta.order)
    .map(route => route.path)
}

// Следим за изменениями маршрута с определением направления анимации
watch(() => route.path, (newPath, oldPath) => {
  const routesOrder = getRoutesOrder()
  const oldIndex = routesOrder.indexOf(oldPath)
  const newIndex = routesOrder.indexOf(newPath)
  
  // Если оба маршрута найдены в порядке - определяем направление
  if (oldIndex !== -1 && newIndex !== -1) {
    transitionName.value = newIndex > oldIndex ? 'slide-left' : 'slide-right'
  } else {
    // Для маршрутов не в порядке используем fade анимацию
    transitionName.value = 'fade'
  }
})

// Следим за изменениями темы
watch(themeParams, (newTheme) => {
  if (newTheme) {
    console.log('🎨 Theme changed in MainLayout:', newTheme)
    applyThemeToApp(newTheme)
    
    // Обновляем заголовок и фон в Telegram WebApp
    updateTelegramAppColors(newTheme)
  }
})

// Следим за изменениями темы (темная/светлая)
watch(isDarkTheme, (newIsDark) => {
  console.log('🌓 Dark theme changed:', newIsDark)
  // Можно добавить дополнительную логику при смене темы
})

// Следим за пользователем Telegram
watch(telegramUser, (newUser) => {
  if (newUser) {
    console.log('👤 Telegram user detected in MainLayout:', newUser)
    // Можно отправить данные на сервер при обнаружении пользователя
    // sendAuthToServer().catch(console.error)
  }
})

// Заголовок страницы из meta данных текущего route
const pageTitle = computed(() => {
  return route.meta?.title || 'Мое Приложение'
})

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

// Дополнительная инициализация при монтировании
onMounted(async () => {
  console.log('🚀 MainLayout mounted - Telegram environment:', isTelegram.value)
  
  // Ждем следующего тика для гарантии применения стилей
  await nextTick()
  
  // Применяем тему сразу при загрузке, если она уже доступна
  if (themeParams.value) {
    console.log('🎨 Applying initial theme:', themeParams.value)
    applyThemeToApp(themeParams.value)
    updateTelegramAppColors(themeParams.value)
  } else {
    // Если тема еще не загружена, пробуем обновить
    refreshTheme()
  }
  
  // Добавляем глобальные обработчики для тем
  if (isTelegram.value && window.Telegram?.WebApp) {
    const webApp = window.Telegram.WebApp
    
    // Обработчик изменения темы
    webApp.onEvent('themeChanged', (themeParams) => {
      console.log('🎨 Global theme changed:', themeParams)
      applyThemeToApp(themeParams)
      updateTelegramAppColors(themeParams)
    })
    
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
</script>

<style scoped>
.app {
  height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: var(--tg-theme-bg-color, #ffffff);
  color: var(--tg-theme-text-color, #000000);
  transition: background-color 0.3s ease, color 0.3s ease;
  overflow: hidden;
}

.main-content {
  flex: 1;
  overflow-y: auto;
  padding: 16px;
  padding-bottom: 80px;
  background-color: var(--tg-theme-bg-color, #ffffff);
  color: var(--tg-theme-text-color, #000000);
  position: relative;
}

/* Стили для Telegram среды */
.app.telegram-env {
  background-color: var(--tg-theme-bg-color);
  color: var(--tg-theme-text-color);
}

/* Анимации переходов */
.slide-left-enter-active,
.slide-left-leave-active,
.slide-right-enter-active,
.slide-right-leave-active {
  transition: all 0.3s ease;
}

.slide-left-enter-from {
  opacity: 0;
  transform: translateX(30px);
}

.slide-left-leave-to {
  opacity: 0;
  transform: translateX(-30px);
}

.slide-right-enter-from {
  opacity: 0;
  transform: translateX(-30px);
}

.slide-right-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
