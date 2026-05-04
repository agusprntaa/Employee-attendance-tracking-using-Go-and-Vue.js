import axios from 'axios'

// Semua request pakai prefix /api
// Vite proxy akan rewrite /api/login → /login di backend
const API = axios.create({
  baseURL: '/api'
})

// Interceptor: otomatis pasang token di setiap request
API.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// Interceptor: kalau 401 → paksa logout
API.interceptors.response.use(
  (res) => res,
  (err) => {
    if (err.response?.status === 401) {
      localStorage.removeItem('token')
      localStorage.removeItem('refresh_token')
      localStorage.removeItem('user')
      window.location.href = '/'
    }
    return Promise.reject(err)
  }
)

export function loginAPI(data) {
  return API.post('/login', data)
}

export function refreshAPI(refreshToken) {
  return API.post('/refresh', { refresh_token: refreshToken })
}

export function logoutAPI(refreshToken) {
  return API.post('/logout', { refresh_token: refreshToken })
}

export function getProfileAPI() {
  return API.get('/employee/profile')
}

export function getTodayAttendanceAPI() {
  return API.get('/attendance/today')
}

export function checkInAPI(data) {
  return API.post('/attendance/checkin', data)
}

export function checkOutAPI(data) {
  return API.patch('/attendance/checkout', data)
}

export function getHistoryAPI(page = 1, limit = 10) {
  return API.get(`/attendance/history?page=${page}&limit=${limit}`)
}

export function changePasswordAPI(data) {
  return API.patch('/employee/change-password', data)
}

export default API