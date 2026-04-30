<script setup>
import { useRouter } from 'vue-router'

defineProps({
  user: Object
})

const router = useRouter()

function getInitials(name) {
  return name
    ?.split(' ')
    .map(n => n[0])
    .join('')
    .slice(0, 2)
    .toUpperCase()
}

const todayDate = new Date().toLocaleDateString('id-ID', {
  weekday: 'long',
  day: 'numeric',
  month: 'long',
  year: 'numeric'
})

function goToChangePassword() {
  router.push('/employee/change-password')
}
</script>

<template>
  <div class="profile-card" v-if="user">

    <div class="top">
      <div class="avatar">
        {{ getInitials(user.name) }}
      </div>

      <div class="info">
        <span class="badge">{{ user.division }}</span>
        <h3>{{ user.name }}</h3>
        <p>{{ todayDate }}</p>
      </div>
    </div>

    <div class="divider"></div>

    <button class="btn-change" @click="goToChangePassword">
      UBAH PASSWORD
    </button>

  </div>
</template>

<style scoped>
.profile-card {
  background: white;
  padding: 20px;
  border-radius: 20px; 
  margin-bottom: 16px;
  box-shadow: 0 10px 25px rgba(0,0,0,0.08);
}

.top {
  display: flex;
  gap: 14px;
  align-items: center;
}

.avatar {
  width: 56px;
  height: 56px;
  border-radius: 50%;
  background: linear-gradient(135deg, #6b8cff, #4f46e5);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: 18px;
  flex-shrink: 0;
}

.info h3 {
  margin: 4px 0;
  font-size: 15px;
  font-weight: 600;
}

.info p {
  font-size: 12px;
  color: #6b7280; 
}

.badge {
  background: #4f46e5;
  color: white;
  padding: 4px 10px;
  border-radius: 20px;
  font-size: 11px;
  font-weight: 500;
}

.divider {
  border-top: 1px dashed #e5e7eb;
  margin: 16px 0;
}

.btn-change {
  width: 100%;
  padding: 12px;
  border-radius: 14px; 
  background: #4f46e5;
  color: white;
  border: none;
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
}

@media (min-width: 1024px) {
  .profile-card {
    padding: 24px;
  }

  .avatar {
    width: 64px;
    height: 64px;
    font-size: 20px;
  }

  .info h3 {
    font-size: 16px;
  }

  .btn-change {
    padding: 14px;
  }
}
</style>