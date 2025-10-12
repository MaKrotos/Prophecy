<template>
  <div class="page">
    <h2>{{ t('messages_view.title') }}</h2>
    <p>{{ t('messages_view.description') }}</p>
  </div>
</template>

<script setup>
import { useLocalization } from '/src/locales/index.js'

const { t } = useLocalization()
</script>