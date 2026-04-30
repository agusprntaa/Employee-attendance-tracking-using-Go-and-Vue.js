import axios from 'axios'

const API = axios.create({
    baseURL: '/'
})

export function loginAPI(data) {
    return API.post('/login', data)
}