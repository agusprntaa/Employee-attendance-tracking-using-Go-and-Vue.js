<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { loginAPI } from '../../services/auth'

const router = useRouter()

// STATE
const username = ref('')
const password = ref('')
const remember = ref(false)
const loading = ref(false)

const locationGranted = ref(false)
const locationError = ref('')

const errorUsername = ref('')
const errorPassword = ref('')
const errorGlobal = ref('')

//init
onMounted(() => {
  // auto username
  const isRemember = localStorage.getItem('remember')
  const savedUsername = localStorage.getItem('savedUsername')

  if (isRemember === 'true' && savedUsername) {
    username.value = savedUsername
    remember.value = true
  }

  // auto login
  const token = localStorage.getItem('token')
  const user = JSON.parse(localStorage.getItem('user'))

  if (token && user) {
    redirectByRole(user)
  }
})

//redirect
function redirectByRole(user) {
  const { role, tipe } = user

  const routeMap = {
    admin: {
      cabang: '/admin-cabang/dashboard',
      pusat: '/admin-pusat/dashboard'
    },
    karyawan: '/employee/dashboard'
  }

  if (role === 'admin') {
    const path = routeMap.admin[tipe]
    if (!path) return console.error('Tipe admin tidak valid')
    return router.push(path)
  }

  if (role === 'karyawan') {
    return router.push(routeMap.karyawan)
  }

  console.error('Role tidak valid')
}

//location
function requestLocation() {
  locationError.value = ''

  navigator.geolocation.getCurrentPosition(
    () => {
      locationGranted.value = true
      locationError.value = ''
    },
    (err) => {
      if (err.code === 1) locationError.value = 'Izin lokasi ditolak'
      else if (err.code === 2) locationError.value = 'Lokasi tidak tersedia'
      else if (err.code === 3) locationError.value = 'Request lokasi timeout'
    },
    { timeout: 5000 }
  )
}

//validasi
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

//login
async function login() {
  if (loading.value) return

  // paksa minta lokasi dulu
  if (!locationGranted.value) {
    await new Promise((resolve) => {
      navigator.geolocation.getCurrentPosition(
        () => {
          locationGranted.value = true
          locationError.value = ''
          resolve()
        },
        () => {
          locationError.value = 'Izin lokasi diperlukan'
          resolve()
        }
      )
    })
  }

  if (!validate()) return

  loading.value = true
  errorGlobal.value = ''

  try {
    const res = await loginAPI({
      username: username.value.trim(),
      password: password.value.trim()
    })

    console.log('USER FROM API:', res.data.user)

    const token = res.data.token
    const user = res.data.user

    // simpan ke localStorage
    localStorage.setItem('token', token)
    localStorage.setItem('user', JSON.stringify(user))

    // remember username
    if (remember.value) {
      localStorage.setItem('remember', 'true')
      localStorage.setItem('savedUsername', username.value)
    } else {
      localStorage.removeItem('remember')
      localStorage.removeItem('savedUsername')
    }

    // redirect
    redirectByRole(user)

  } catch (err) {
    console.log('ERROR:', err)
    errorGlobal.value =
      err.response?.data?.message || 'Login gagal'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="container">
    <div class="card">
      <h2>ABSENSI KARYAWAN</h2>

      <div class="location-box" @click="requestLocation">
        <span>
          {{ locationGranted ? 'Lokasi aktif' : 'Izin lokasi diperlukan' }}
        </span>
      </div>
      <p v-if="locationError" class="error">{{ locationError }}</p>
      
      <label>Username</label>
      <input v-model="username" placeholder="Masukkan username" />
      <p v-if="errorUsername" class="error">{{ errorUsername }}</p>

      <label>Password</label>
      <input type="password" v-model="password" placeholder="Masukkan password" />
      <p v-if="errorPassword" class="error">{{ errorPassword }}</p>
      
    <label class="remember">
        <input type="checkbox" v-model="remember" />
        <span>Ingat saya</span>
    </label>

      <button @click="login" :disabled="loading">
        {{ loading ? 'Loading...' : 'Sign In' }}
      </button>

      <p v-if="errorGlobal" class="error">{{ errorGlobal }}</p>
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
  box-shadow: 0 20px 40px rgba(0,0,0,0.08);
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
  margin-bottom: 20px;
  cursor: pointer;
  text-align: center;
}

.location-box span {
  font-size: 12px;
  color: gray;
}

label {
  display: block;
  font-size: 14px;
  margin-bottom: 6px;
  margin-top: 12px;
}

input {
  width: 100%;
  padding: 14px;
  border-radius: 14px;
  border: 1px solid #ddd;
  font-size: 14px;
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
  background:#4f46e5;
  color: white;
  font-weight: 600;
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
</style>