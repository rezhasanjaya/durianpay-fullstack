import axios from 'axios'
import { useAuthStore } from '../stores/auth'

type ApiOptions = {
  auth?: boolean
}

export function createApi(options: ApiOptions = {}) {
  const instance = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
    headers: {
      'Content-Type': 'application/json',
      Accept: 'application/json',
    },
  })

  if (options.auth) {
    const auth = useAuthStore()
    const token = auth.token
    if (token) {
      instance.defaults.headers.common.Authorization = `Bearer ${token}`
    }
  }

  return instance
}
