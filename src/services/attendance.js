import axios from 'axios'

const API = axios.create({
  baseURL: '/'
})

export function checkInAPI() {
  return API.post('/attendance/check-in')
}

export function getAttendanceHistory() {
  return API.get('/attendance/history')
}