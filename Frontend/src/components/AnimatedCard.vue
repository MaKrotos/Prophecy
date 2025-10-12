<template>
  <div
    class="animated-card"
    :class="[customClass, { 'animated': !animationPlayed }]"
    :style="{ animationDelay: `${delay}s` }"
  >
    <slot></slot>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'

const props = defineProps({
  index: {
    type: Number,
    default: 0
  },
  animationDelay: {
    type: Number,
    default: 0.1
  },
  customClass: {
    type: String,
    default: ''
  },
  itemId: {
    type: [String, Number],
    default: null
  }
})

// Храним состояние анимации для каждого элемента
const animationPlayed = ref(false)

// Вычисляем задержку анимации: базовая задержка + индекс * шаг задержки
const delay = computed(() => {
  return props.animationDelay + (props.index * 0.1)
})

// При монтировании компонента устанавливаем флаг анимации
// В Vue 3 Composition API нет хука onMounted в <script setup> по умолчанию,
// но мы можем импортировать его
import { onMounted } from 'vue'

onMounted(() => {
  // Для новых элементов запускаем анимацию
  if (!animationPlayed.value) {
    // Устанавливаем флаг анимации в true после небольшой задержки,
    // чтобы дать время на проигрывание анимации
    setTimeout(() => {
      animationPlayed.value = true
    }, (props.animationDelay + (props.index * 0.1) + 0.5) * 1000) // 0.5s - длительность анимации
  }
})
</script>

<style scoped>
.animated-card {
  background: var(--tg-theme-secondary-bg-color, white);
  border-radius: 12px;
  padding: 16px;
  box-shadow: 0 1px 6px rgba(0, 0, 0, 0.08);
  border: 1px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
  transition: all 0.3s ease;
  
  /* Начальное состояние для карточек */
  opacity: 1;
  transform: translateY(0);
}

/* Начальное состояние для анимированных карточек */
.animated-card.animated {
  opacity: 0;
  transform: translateY(20px);
  animation: fadeInUp 0.5s ease forwards;
}

/* Эффекты наведения */
.animated-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.12);
}

@keyframes fadeInUp {
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Специфические стили для темной темы */
:global(.tg-theme-dark) .animated-card {
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
}

:global(.tg-theme-dark) .animated-card:hover {
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.4);
}

/* Плавные переходы для всех элементов */
.animated-card,
.animated-card * {
  transition: all 0.3s ease;
}
</style>