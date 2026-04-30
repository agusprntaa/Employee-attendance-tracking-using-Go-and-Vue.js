<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { loginAPI } from '../../services/auth'
import { useAuth } from '@/composables/useAuth'

const router = useRouter()
const { setUser } = useAuth()

const username = ref('')
const password = ref('')
const remember = ref(false)
const loading = ref(false)

const locationGranted = ref(false)
const locationError = ref('')

const errorUsername = ref('')
const errorPassword = ref('')
const errorGlobal = ref('')

// INIT
onMounted(() => {
  const isRemember = localStorage.getItem('remember')
  const savedUsername = localStorage.getItem('savedUsername')

  if (isRemember === 'true' && savedUsername) {
    username.value = savedUsername
    remember.value = true
  }
})

// VALIDASI
function validate() {
  let valid = true
  errorUsername.value = ''
  errorPassword.value = ''
  errorGlobal.value = ''

  if (!username.value.trim()) {
    errorUsername.value = 'Masukkan username'
    valid = false
  }

  if (!password.value.trim()) {
    errorPassword.value = 'Masukkan password'
    valid = false
  }

  if (!locationGranted.value) {
    locationError.value = 'Izin lokasi diperlukan'
    valid = false
  }

  return valid
}

// LOGIN
async function login() {
  if (loading.value) return

  if (!validate()) return

  loading.value = true
  errorGlobal.value = ''

  try {
    const res = await loginAPI({
      username: username.value.trim(),
      password: password.value.trim()
    })

    const { token, refresh_token, user } = res.data.data

    localStorage.setItem('token', token)
    localStorage.setItem('refresh_token', refresh_token)

    setUser(user)

    router.push('/')

    // remember
    if (remember.value) {
      localStorage.setItem('remember', 'true')
      localStorage.setItem('savedUsername', username.value)
    } else {
      localStorage.removeItem('remember')
      localStorage.removeItem('savedUsername')
    }

  } catch (err) {
    errorGlobal.value =
      err.response?.data?.message ||
      err.response?.data?.error ||
      'Login gagal'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="container">
    <div class="card">
      <h2>ABSENSI KARYAWAN</h2>

      <div
        class="location-box"
        :class="{ active: locationGranted }"
        @click="requestLocation"
      >
        <span>
          {{ locationGranted ? '✓ Lokasi aktif' : 'Klik untuk izin lokasi' }}
        </span>
      </div>
      <p v-if="locationError" class="error">{{ locationError }}</p>

      <label>Username</label>
      <input
        v-model="username"
        type="text"
        placeholder="Masukkan username"
        @keyup.enter="login"
      />
      <p v-if="errorUsername" class="error">{{ errorUsername }}</p>

      <label>Password</label>
      <input
        type="password"
        v-model="password"
        placeholder="Masukkan password"
        @keyup.enter="login"
      />
      <p v-if="errorPassword" class="error">{{ errorPassword }}</p>

      <label class="remember">
        <input type="checkbox" v-model="remember" />
        <span>Ingat saya</span>
      </label>

      <button @click="login" :disabled="loading">
        {{ loading ? 'Loading...' : 'Sign In' }}
      </button>

      <p v-if="errorGlobal" class="error global-error">{{ errorGlobal }}</p>
    </div>
  </div>
</template>

<style scoped>
.container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #f3f4f6;
  padding: 20px;
}

.card {
  width: 100%;
  max-width: 420px;
  background: #ffffff;
  padding: 32px 28px;
  border-radius: 24px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.08);
}

h2 {
  text-align: center;
  margin-bottom: 24px;
  font-weight: 700;
  font-size: 24px;
}

.location-box {
  background: #e5e7eb;
  padding: 10px;
  border-radius: 10px;
  margin-bottom: 8px;
  cursor: pointer;
  text-align: center;
  transition: background 0.2s;
}

.location-box.active {
  background: #d1fae5;
}

.location-box span {
  font-size: 12px;
  color: gray;
}

.location-box.active span {
  color: #065f46;
}

label {
  display: block;
  font-size: 14px;
  margin-bottom: 6px;
  margin-top: 12px;
}

input[type='text'],
input[type='email'],
input[type='password'] {
  width: 100%;
  padding: 14px;
  border-radius: 14px;
  border: 1px solid #ddd;
  font-size: 14px;
  box-sizing: border-box;
}

.remember {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 12px;
  cursor: pointer;
  font-size: 14px;
  color: #374151;
}

.remember input {
  width: 16px;
  height: 16px;
  cursor: pointer;
}

button {
  width: 100%;
  padding: 16px;
  margin-top: 24px;
  border: none;
  border-radius: 14px;
  background: #4f46e5;
  color: white;
  font-weight: 600;
  cursor: pointer;
}

button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.error {
  color: #dc2626;
  font-size: 12px;
  margin-top: 4px;
}

.global-error {
  text-align: center;
  margin-top: 12px;
}
</style>
