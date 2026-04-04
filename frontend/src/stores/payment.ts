import { defineStore } from "pinia";
import { ref } from "vue";
import { createApi } from "../lib/api";

export const usePaymentStore = defineStore("payment", () => {
    const api = createApi({ auth: true })

    const payments = ref<any[]>([])
    const widget = ref<any[]>([])
    const totalData = ref<number>(0)
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

    async function fetchWidget() {
        loading.value = true
        error.value = null
        
        try {   
            const response = await api.get('/dashboard/v1/widget')
            widget.value = response.data.widget ?? []
            totalData.value = response.data.total ?? 0
        } catch (err) {
            console.error('Failed to fetch widget data:', err)
            error.value = 'Failed to load widget data'
        } finally {
            loading.value = false
        }
    }

  return { payments, widget, totalData, loading, error, fetch, fetchWidget }
})  
