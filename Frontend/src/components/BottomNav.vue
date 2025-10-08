<script setup>
import { ref, watch, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()
const activeTab = ref('/')

const navItems = [
  { path: '/', label: 'Главная', icon: 'M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z' },
  { path: '/messages', label: 'Сообщения', icon: 'M20 2H4c-1.1 0-1.99.9-1.99 2L2 22l4-4h14c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm-2 12H6v-2h12v2zm0-3H6V9h12v2zm0-3H6V6h12v2z' },
  { path: '/profile', label: 'Профиль', icon: 'M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z' },
  { path: '/settings', label: 'Настройки', icon: 'M19.14 12.94c.04-.3.06-.61.06-.94 0-.32-.02-.64-.07-.94l2.03-1.58c.18-.14.23-.41.12-.61l-1.92-3.32c-.12-.22-.37-.29-.59-.22l-2.39.96c-.5-.38-1.03-.7-1.62-.94l-.36-2.54c-.04-.24-.24-.41-.48-.41h-3.84c-.24 0-.43.17-.47.41l-.36 2.54c-.59.24-1.13.57-1.62.94l-2.39-.96c-.22-.08-.47 0-.59.22L2.74 8.87c-.12.21-.08.47.12.61l2.03 1.58c-.05.3-.09.63-.09.94s.02.64.07.94l-2.03 1.58c-.18.14-.23.41-.12.61l1.92 3.32c.12.22.37.29.59.22l2.39-.96c.5.38 1.03.7 1.62.94l.36 2.54c.05.24.24.41.48.41h3.84c.24 0 .44-.17.47-.41l.36-2.54c.59-.24 1.13-.56 1.62-.94l2.39.96c.22.08.47 0 .59-.22l1.92-3.32c.12-.22.07-.47-.12-.61l-2.01-1.58zM12 15.6c-1.98 0-3.6-1.62-3.6-3.6s1.62-3.6 3.6-3.6 3.6 1.62 3.6 3.6-1.62 3.6-3.6 3.6z' },
  { path: '/scan', label: 'Сканер', icon: 'M4 4h16v16H4V4zm2 2v12h12V6H6zm3 3h6v2H9V9zm0 4h6v2H9v-2z' }
]

// Инициализация при монтировании
onMounted(() => {
  if (route && route.path) {
    activeTab.value = route.path
  }
})

// Следим за изменениями маршрута
watch(() => route?.path, (newPath) => {
  if (newPath) {
    activeTab.value = newPath
  }
})

const navigateTo = (path) => {
  activeTab.value = path
  router.push(path)
}
</script>

<template>
  <nav class="bottom-nav">
    <div class="nav-container">
      <button
        v-for="item in navItems"
        :key="item.path"
        class="nav-item"
        :class="{ active: activeTab === item.path }"
        @click="navigateTo(item.path)"
      >
        <svg width="24" height="24" viewBox="0 0 24 24" fill="currentColor">
          <path :d="item.icon" />
        </svg>
        <span class="nav-label">{{ item.label }}</span>
      </button>
    </div>
  </nav>
</template>

<style scoped>
.bottom-nav {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background: var(--tg-theme-bg-color, white);
  border-top: 1px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
  padding: 8px 0;
  z-index: 100;
  transition: all 0.3s ease;
}

.nav-container {
  display: flex;
  justify-content: space-around;
  align-items: center;
  max-width: 480px;
  margin: 0 auto;
}

.nav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  background: none;
  border: none;
  padding: 8px 12px;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
  color: var(--tg-theme-hint-color, #666);
  min-width: 60px;
}

.nav-item:hover {
  background-color: var(--tg-theme-secondary-bg-color, #f0f0f0);
}

.nav-item.active {
  color: var(--tg-theme-button-color, #667eea);
  background-color: var(--tg-theme-secondary-bg-color, #f0f0f0);
}

.nav-item.active .nav-label {
  color: var(--tg-theme-button-color, #667eea);
  font-weight: 600;
}

.nav-label {
  font-size: 0.75rem;
  margin-top: 4px;
  font-weight: 500;
  color: var(--tg-theme-hint-color, #666);
  transition: all 0.3s ease;
}

/* Анимация для активного элемента */
.nav-item {
  position: relative;
  overflow: hidden;
}

.nav-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: var(--tg-theme-button-color, #667eea);
  opacity: 0.1;
  border-radius: 12px;
  transform: scale(0);
  transition: transform 0.3s ease;
}

.nav-item.active::before {
  transform: scale(1);
}

/* Улучшенные стили для темной темы */
:global(.tg-theme-dark) .bottom-nav {
  border-top-color: var(--tg-theme-secondary-bg-color, #2a3b4d);
}

:global(.tg-theme-dark) .nav-item:hover {
  background-color: var(--tg-theme-secondary-bg-color, #2a3b4d);
}

:global(.tg-theme-dark) .nav-item.active {
  background-color: var(--tg-theme-secondary-bg-color, #2a3b4d);
}

/* Адаптация для разных размеров экрана */
@media (max-width: 480px) {
  .nav-container {
    padding: 0 8px;
  }
  
  .nav-item {
    padding: 6px 8px;
    min-width: 50px;
  }
  
  .nav-label {
    font-size: 0.7rem;
  }
}

@media (max-width: 360px) {
  .nav-label {
    font-size: 0.65rem;
  }
  
  .nav-item {
    padding: 4px 6px;
  }
}

/* Плавные переходы для всех свойств */
.bottom-nav,
.nav-item,
.nav-label {
  transition: all 0.3s ease;
}
</style>