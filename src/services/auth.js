import API from './api'

export function loginAPI(data) {
  return API.post('/login', data)
}