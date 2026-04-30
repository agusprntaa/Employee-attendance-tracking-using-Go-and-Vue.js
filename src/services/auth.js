import API from './api'

export function loginAPI(data) {
  return API.post('/login', data)
}

export function refreshToken() {
  return API.post('/refresh')
}

export function logoutAPI() {
  return API.post('/logout')
}