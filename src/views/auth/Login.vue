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

// INIT
onMounted(() => {
  // Auto-isi email jika remember diaktifkan sebelumnya
  const isRemember = localStorage.getItem('remember')
  const savedUsername = localStorage.getItem('savedUsername')
  if (isRemember === 'true' && savedUsername) {
    username.value = savedUsername
    remember.value = true
  }

  // Auto-login jika token & user masih ada
  const token = localStorage.getItem('token')
  const userRaw = localStorage.getItem('user')
  if (token && userRaw) {
    const user = JSON.parse(userRaw)
    redirectByRole(user)
  }
})

// REDIRECT berdasarkan role & tipe
// role dari backend: "super_admin" / "admin_cabang" / "karyawan"
// tipe dari backend: "pusat" / "cabang"
function redirectByRole(user) {
  const { role } = user

  if (role === 'super_admin') {
    return router.push('/admin-pusat/dashboard')
  }

  if (role === 'admin_cabang') {
    return router.push('/admin-cabang/dashboard')
  }

  if (role === 'karyawan') {
    return router.push('/employee/dashboard')
  }

  console.error('Role tidak dikenali:', role)
}

// MINTA IZIN LOKASI
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

// VALIDASI FORM
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

  // Minta izin lokasi dulu jika belum
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

    // Response dari backend sekarang: { status, data: { token, refresh_token, user } }
    const { token, refresh_token, user } = res.data.data

    // Simpan ke localStorage
    localStorage.setItem('token', token)
    localStorage.setItem('refresh_token', refresh_token)
    localStorage.setItem('user', JSON.stringify(user))

    // Remember username
    if (remember.value) {
      localStorage.setItem('remember', 'true')
      localStorage.setItem('savedUsername', username.value)
    } else {
      localStorage.removeItem('remember')
      localStorage.removeItem('savedUsername')
    }

    redirectByRole(user)

  } catch (err) {
    console.error('Login error:', err)
    // Ambil pesan error dari response backend
    errorGlobal.value =
      err.response?.data?.message ||
      err.response?.data?.error ||
      'Login gagal, coba lagi'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="container">
    <div class="card">
      <h2>ABSENSI KARYAWAN</h2>

      <!-- Tombol izin lokasi -->
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

      <!-- Input username -->
      <label>Username</label>
      <input
        v-model="username"
        type="text"
        placeholder="Masukkan username"
        @keyup.enter="login"
      />
      <p v-if="errorUsername" class="error">{{ errorUsername }}</p>

      <!-- Input password -->
      <label>Password</label>
      <input
        type="password"
        v-model="password"
        placeholder="Masukkan password"
        @keyup.enter="login"
      />
      <p v-if="errorPassword" class="error">{{ errorPassword }}</p>

      <!-- Remember me -->
      <label class="remember">
        <input type="checkbox" v-model="remember" />
        <span>Ingat saya</span>
      </label>

      <!-- Tombol login -->
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
