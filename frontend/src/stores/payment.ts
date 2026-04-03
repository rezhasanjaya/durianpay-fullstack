import { defineStore } from "pinia";
import { ref } from "vue";
import { createApi } from "../lib/api";

export const usePaymentStore = defineStore("payment", () => {
    const api = createApi({ auth: true })

    const payments = ref<any[]>([])
    const loading = ref(false)
    const error = ref<string | null>(null)

    async function fetch(status?: string, id?: string) {
        loading.value = true
        error.value = null

        try {
            const params: Record<string, string> = {}

            if (status && status !== 'all') params.status = status
            if (id && id.trim()) params.id = id.trim()

            const response = await api.get('/dashboard/v1/payments', { params })
            payments.value = response.data.payments ?? []

        } catch (err) {
            console.error('Failed to fetch payments:', err)
            error.value = 'Failed to load payments'
        } finally {
            loading.value = false
        }
    }

  return { payments, loading, error, fetch }
})  
