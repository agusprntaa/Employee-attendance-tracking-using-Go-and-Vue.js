import API from './api'

export function getDashboardSummary() {
  return API.get('/dashboard/summary')
}

export function getTodayAttendance(params) {
  return API.get('/dashboard/today', { params })
}

export function getQRCode() {
  return API.get('/dashboard/qr-code')
}

export function refreshQRCode() {
  return API.post('/dashboard/qr-code/refresh')
}

export function getBranchSettings() {
  return API.get('/settings')
}

export function getEmployees(params) {
  return API.get('/employees', { params })
}

export function addEmployee(data) {
  return API.post('/employees', data)
}

export function updateEmployee(id, data) {
  return API.put(`/employees/${id}`, data)
}

export function toggleEmployeeStatus(id) {
  return API.patch(`/employees/${id}/toggle-status`)
}

export function deleteEmployee(id) {
  return API.delete(`/employees/${id}`)
}