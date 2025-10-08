<template>
  <div class="page">
    <div class="profile-header">
      <div class="avatar">👤</div>
      <h2>Ваш Профиль</h2>
      <p>user@example.com</p>
    </div>
    
    <div class="profile-actions">
      <button class="action-btn">Редактировать профиль</button>
      <button class="action-btn">Настройки приватности</button>
      <button class="action-btn">Выйти</button>
    </div>
  </div>
</template>

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
  background: white;
  border: 1px solid #ddd;
  padding: 16px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s;
}

.action-btn:hover {
  background: #f5f5f5;
  border-color: #667eea;
}
</style>