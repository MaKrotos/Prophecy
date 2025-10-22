<template>
  <div class="page">
    <div class="profile-header">
      <div class="avatar">👤</div>
      <h2>{{ t('profile_view.title') }}</h2>
      <p>user@example.com</p>
    </div>
    
    <div class="profile-actions">
      <ThemedButton buttonType="secondary" class="action-btn">
        {{ t('profile_view.edit_profile') }}
      </ThemedButton>
      <ThemedButton buttonType="secondary" class="action-btn">
        {{ t('profile_view.privacy_settings') }}
      </ThemedButton>
      <ThemedButton buttonType="danger" class="action-btn">
        {{ t('profile_view.logout') }}
      </ThemedButton>
    </div>
  </div>
</template>

<script setup>
import { useLocalization } from '@/locales/index.js'
import ThemedButton from '@/components/ThemedButton.vue'

const { t } = useLocalization()
</script>

<style scoped>
.profile-header {
  text-align: center;
  margin-bottom: 32px;
}

.avatar {
  width: 80px;
  height: 80px;
  background: linear-gradient(135deg, #667eea, #764ba2);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2rem;
  margin: 0 auto 16px;
}

.profile-actions {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.action-btn {
  width: 100%;
}
</style>