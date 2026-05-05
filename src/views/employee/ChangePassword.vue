<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import API from '@/services/api'

const router = useRouter()

const oldPassword = ref('')
const newPassword = ref('')
const confirmPassword = ref('')

const error = ref('')
const loading = ref(false)

async function handleSubmit() {
  if (loading.value) return

  const oldPass = oldPassword.value.trim()
  const newPass = newPassword.value.trim()
  const confirmPass = confirmPassword.value.trim()

  if (!oldPass || !newPass || !confirmPass) {
    error.value = 'Semua field wajib diisi'
    return
  }

  if (newPass.length < 6) {
    error.value = 'Password minimal 6 karakter'
    return
  }

  if (newPass !== confirmPass) {
    error.value = 'Konfirmasi password tidak cocok'
    return
  }

  error.value = ''
  loading.value = true

  try {
    await API.patch('/employee/change-password', {
      old_password: oldPass,
      new_password: newPass
    })

    alert('Password berhasil diubah')
    router.back()

  } catch (err) {
    error.value =
      err.response?.data?.message ||
      'Gagal mengubah password'
  } finally {
    loading.value = false
  }
}

function goBack() {
  router.back()
}
</script>

<template>
  <div class="wrapper">

    <div class="header">
      <img src="/goBack.png" class="back" @click="goBack">
    </div>

    <div class="content">
      <div class="card">

        <h2>UBAH PASSWORD</h2>

        <div class="field">
          <label>Password lama</label>
          <input type="password" v-model="oldPassword" />
        </div>

        <div class="field">
          <label>Password baru</label>
          <input type="password" v-model="newPassword" />
        </div>

        <div class="field">
          <label>Konfirmasi password</label>
          <input type="password" v-model="confirmPassword" />
        </div>

        <p v-if="error" class="error">{{ error }}</p>

        <button
          class="btn"
          @click="handleSubmit"
          :disabled="loading"
        >
          {{ loading ? 'Menyimpan...' : 'Ubah password' }}
        </button>

      </div>
    </div>

  </div>
</template>

<style scoped>
.wrapper {
  min-height: 100vh;
  background: #ffffff;
}

.header {
  height: 60px;
  background: #4f46e5;
  display: flex;
  align-items: center;
  padding: 0 30px;
}

.back {
  width: 24px;
  cursor: pointer;
}

.content {
  display: flex;
  justify-content: center;
  padding: 40px 20px;
}

.card {
  width: 100%;
  max-width: 360px;
  background: white;
  padding: 28px 22px;
  border-radius: 24px;
  box-shadow: 0 20px 40px rgba(0,0,0,0.15);
}

h2 {
  text-align: center;
  margin-bottom: 24px;
  font-weight: 700;
}

.field {
  margin-bottom: 16px;
}

label {
  font-size: 14px;
  font-weight: 500;
}

input {
  width: 100%;
  padding: 14px;
  border-radius: 14px;
  border: none;
  margin-top: 6px;
  background: #eef2ff;
  outline: none;
}

.btn {
  width: 100%;
  padding: 14px;
  border-radius: 16px;
  background: linear-gradient(135deg, #5b7cfa, #4f46e5);
  color: white;
  border: none;
  margin-top: 16px;
  font-weight: 600;
}

.btn:disabled {
  opacity: 0.6;
}

.error {
  color: #dc2626;
  font-size: 12px;
  margin-top: 8px;
  text-align: center;
}
</style>