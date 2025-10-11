import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
    meta: {
      title: 'Главная',
      order: 1
    }
  },
  {
    path: '/messages',
    name: 'messages',
    component: () => import('../views/MessagesView.vue'),
    meta: {
      title: 'Сообщения',
      order: 2
    }
  },
  {
    path: '/profile',
    name: 'profile',
    component: () => import('../views/ProfileView.vue'),
    meta: {
      title: 'Профиль',
      order: 3
    }
  },
  {
    path: '/settings',
    name: 'settings',
    component: () => import('../views/SettingsView.vue'),
    meta: {
      title: 'Настройки',
      order: 4
    }
  },
  {
    path: '/scan',
    name: 'scan',
    component: () => import('../views/ScanView.vue'),
    meta: {
      title: 'Сканер',
      order: 5
    }
  },
  {
    path: '/users',
    name: 'users',
    component: () => import('../views/UsersView.vue'),
    meta: {
      title: 'Все пользователи',
      order: 6
    }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Глобальный обработчик для заголовков страниц
router.beforeEach((to, from, next) => {
  // Можно также установить заголовок документа здесь
  document.title = to.meta.title || 'Мое Приложение'
  next()
})

export default router