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
      
      <transition name="nav" mode="out-in">
        <BottomNav v-if="!isInSession" key="default-nav" />
        <BottomNav v-else key="session-nav" :nav-items="sessionNavItemsWithId" />
      </transition>
    </div>
  </ThemeManager>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import HeaderNav from '../components/HeaderNav.vue'
import BottomNav from '../components/BottomNav.vue'
import ThemeManager from '../components/ThemeManager.vue'
import { useTelegramWebAppSingleton } from '../telegram/composables/useTelegramWebAppSingleton'
import { getUserInfoFromToken } from '../telegram/auth/user'
import { useLocalization } from '../locales/index.js'

const route = useRoute()
const router = useRouter()
const transitionName = ref('slide-left')
const { t } = useLocalization()

// Используем рефакторизованный композабл
const {
  isTelegram,
  telegramUser,
  sendAuthToServer
} = useTelegramWebAppSingleton()

// Проверяем, является ли пользователь администратором
const userInfo = computed(() => getUserInfoFromToken())
const isAdmin = computed(() => userInfo.value?.isAdmin || false)

// Определяем, находится ли пользователь в контексте сессии
const isInSession = computed(() => {
  // Проверяем, что это маршрут сессии, связанный с конкретной сессией
  return (route.path.startsWith('/session/') && !route.path.startsWith('/session/qr-scanner')) ||
         (route.path.startsWith('/sessions/') && route.params.id) ||
         route.path.startsWith('/sessions/') && route.path.endsWith('/clan')
})

// Получаем ID сессии из текущего маршрута
const sessionId = computed(() => {
  // Проверяем, есть ли ID в параметрах маршрута
  if (route.params && route.params.id) {
    return route.params.id
  }
  
  // Если нет, пытаемся получить из пути /session/my-qr/:id
  const sessionPathMatch = route.path.match(/\/session\/my-qr\/([^\/]+)/)
  if (sessionPathMatch && sessionPathMatch[1]) {
    return sessionPathMatch[1]
  }
  
  // Также проверяем путь /sessions/:id
  const sessionDetailMatch = route.path.match(/\/sessions\/([^\/]+)/)
  if (sessionDetailMatch && sessionDetailMatch[1]) {
    return sessionDetailMatch[1]
  }
  
  // Проверяем путь /sessions/:id/clan
  const clanMatch = route.path.match(/\/sessions\/([^\/]+)\/clan/)
  if (clanMatch && clanMatch[1]) {
    return clanMatch[1]
  }
  
  // Если не удалось определить, возвращаем null
  return null
})

// Навигационные элементы для сессии с передачей ID
const sessionNavItemsWithId = computed(() => {
  // Базовые элементы навигации для сессии
  const baseItems = [
    { path: sessionId.value ? `/session/role/${sessionId.value}` : '/session/role', label: t('bottom_nav.my_role'), icon: 'M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z' },
    { path: sessionId.value ? `/session/friends/${sessionId.value}` : '/session/friends', label: t('bottom_nav.friends_list'), icon: 'M16 11c1.66 0 2.99-1.34 2.99-3S17.66 5 16 5c-1.66 0-3 1.34-3 3s1.34 3 3 3zm-8 0c1.66 0 2.99-1.34 2.99-3S9.66 5 8 5C6.34 5 5 6.34 5 8s1.34 3 3 3zm0 2c-2.33 0-7 1.17-7 3.5V19h14v-2.5c0-2.33-4.67-3.5-7-3.5zm8 0c-.29 0-.62.02-.97.05 1.16.84 1.97 1.97 1.97 3.45V19h6v-2.5c0-2.33-4.67-3.5-7-3.5z' },
    { path: sessionId.value ? `/session/my-qr/${sessionId.value}` : '/session/my-qr', label: t('bottom_nav.my_qr'), icon: 'M3 11h8V3H3v8zm2-6h4v4H5V5zm8 14h8v-8h-8v8zm2-6h4v4h-4v-4zm-8 2v4h8v-4H3zm10 0v4h8v-4h-8z' },
    { path: '/session/qr-scanner', label: t('bottom_nav.qr_scanner'), icon: 'M4 4h16v16H4V4zm2 2v12h12V6H6zm3 3h6v2H9V9zm0 4h6v2H9v-2z' }
  ]
  
  // Добавляем кнопку клана в конец
  return [
    ...baseItems,
    { path: sessionId.value ? `/sessions/${sessionId.value}/clan` : '/clan', label: t('bottom_nav.clan'), icon: 'M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z' }
  ]
})

onMounted(() => {

  
})

/**
 * Отмечает, что пользователь уже посещал приложение
 */
const markAsVisited = () => {
  localStorage.setItem('hasVisitedApp', 'true');
}


/**
 * Проверяет, первый ли раз пользователь открывает приложение
 * @returns {boolean} true, если первый раз, false если нет
 */
const isFirstTimeUser = () => {
  const hasVisited = localStorage.getItem('hasVisitedApp');
  return !hasVisited;
}

// Получаем пути маршрутов из meta данных
// Параметр order больше не используется
const getRoutesOrder = () => {
  return router.getRoutes()
    .filter(route => route.meta?.title && !route.meta?.isNested)
    .map(route => route.path)
}

// Следим за изменениями маршрута с определением направления анимации
// Маршруты больше не требуют строгой сортировки благодаря изменению путей
// Параметр order был удален из всех маршрутов
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
  return route.meta?.title || 'Prophecy'
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
