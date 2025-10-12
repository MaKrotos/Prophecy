<template>
  <div class="page">
    <h2>{{ t('settings_view.title') }}</h2>
    <p>{{ t('settings_view.description') }}</p>
  </div>
</template>

<script setup>
import { useLocalization } from '@/locales/index.js'
const { t } = useLocalization()
</script>