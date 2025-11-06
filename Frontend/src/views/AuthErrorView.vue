<template>
  <AuthErrorLayout 
    :error-message="errorMessage" 
    @retry="handleRetry"
    @try-later="handleTryLater"
  />
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import AuthErrorLayout from '@/layouts/AuthErrorLayout.vue'

const router = useRouter()
const errorMessage = ref('')

// Получаем сообщение об ошибке из параметров маршрута или localStorage
onMounted(() => {
  // Попробуем получить ошибку из query параметров
  const errorParam = router.currentRoute.value.query.error
  if (errorParam) {
    errorMessage.value = decodeURIComponent(errorParam)
  } else {
    // Если нет параметра, попробуем получить из localStorage
    const storedError = localStorage.getItem('authErrorMessage')
    if (storedError) {
      errorMessage.value = storedError
      localStorage.removeItem('authErrorMessage') // Удаляем после использования
    }
  }
})

const handleRetry = () => {
  // Попытка повторной авторизации
  // Перенаправляем на главную страницу, где будет выполнена попытка авторизации
  router.push('/')
}

const handleTryLater = () => {
  // Отложить попытку авторизации
  // Можно перенаправить на главную страницу или другую подходящую
  router.push('/')
}
</script>