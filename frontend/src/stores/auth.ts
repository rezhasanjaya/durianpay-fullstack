import { defineStore } from "pinia";
import { ref,computed } from "vue";
import { createApi } from "../lib/api";


export const useAuthStore = defineStore("auth", () => {
    const api = createApi()

    const token = ref<string | null>(localStorage.getItem('token'))
    const role = ref<string | null>(localStorage.getItem('role'))
    const loading = ref(false)
    const error = ref<string | null>(null)

    const isLogin = computed(() => !!token.value)

    async function login(email: string, password: string) {
    loading.value = true
    error.value = null

    try {
        const response = await api.post('/dashboard/v1/auth/login', { email, password })

        token.value = response.data.token
        role.value = response.data.role

        if (response.data.token) {
            token.value = response.data.token
            localStorage.setItem('token', token.value ?? '')
        }

        if (response.data.role) {
            role.value = response.data.role
            localStorage.setItem('role', role.value ?? '')
        }

    } catch (error) {
        console.error('Login failed:', error)
        throw error

    } finally {
        loading.value = false
    }
    }

    async function logout() {
        token.value = null
        role.value = null
        localStorage.removeItem('token')
        localStorage.removeItem('role')
    }

    return { token, role, isLogin, error, loading, login, logout }
});
