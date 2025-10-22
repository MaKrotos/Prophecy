<template>
  <ThemeManager>
    <div class="app">
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
  </ThemeManager>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import HeaderNav from '../components/HeaderNav.vue'
import BottomNav from '../components/BottomNav.vue'
import UserStats from '../components/UserStats.vue'
import ThemeManager from '../components/ThemeManager.vue'
import { useTelegramWebApp } from '../telegram/composables/useTelegramWebApp'
import { getUserInfoFromToken } from '../telegram/auth/user'

const route = useRoute()
const router = useRouter()
const transitionName = ref('slide-left')

// Используем рефакторизованный композабл
const {
  isTelegram,
  telegramUser,
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
