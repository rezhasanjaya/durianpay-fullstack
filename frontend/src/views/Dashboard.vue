<script setup lang="ts">
import { useAuthStore } from '../stores/auth'
import { usePaymentStore } from '../stores/payment'
import { useRouter } from 'vue-router'
import { onMounted, ref, watch } from 'vue'
import { formatTimestamp, capitalize } from '../utils/format'

import Button from '../components/Button.vue'
import Table from '../components/Table.vue'
import InputField from '../components/InputField.vue'
import SelectField from '../components/SelectField.vue'

const authStore = useAuthStore()
const paymentStore = usePaymentStore()
const route = useRouter()
const selectedFilter = ref('all')
const searchValue = ref('')
const data = ref<any[]>([])
const widget = ref<any>(null)
const totalWidget = ref<number>(0)

const handleLogout = () => {
    authStore.logout()
    route.push('/')
}

const handleReset = () => {
    selectedFilter.value = 'all'
    searchValue.value = ''
}

let searchTimer: ReturnType<typeof setTimeout> | null = null

const fetchPayments = async () => {
    try {
        await paymentStore.fetch(selectedFilter.value, searchValue.value)
        data.value = paymentStore.payments
    } catch (error) {
        ///
    }
}

const fetchWidget = async () => {
    try {
        await paymentStore.fetchWidget()
        widget.value = paymentStore.widget
        totalWidget.value = paymentStore.totalData
    } catch (error) {
        ///
    }
}

onMounted(() => {
    fetchPayments()
    fetchWidget()
})

watch(selectedFilter, () => {
    fetchPayments()
})

watch(searchValue, () => {
    if (searchTimer) clearTimeout(searchTimer)
    searchTimer = setTimeout(() => {
        fetchPayments()
    }, 400)
})
</script>

<template>
    <div>
        <div class="px-6 py-10">
            <div class="flex items-center justify-between gap-3">
                <div>
                <h1 class="text-2xl font-bold text-slate-100">Dashboard</h1>
                    <p class="text-neutral-400 text-sm">Hai, {{ authStore.role }}</p>
                    <p v-if="authStore.role == 'operation'" class="text-neutral-400 text-sm">You are an Operations user. You can filter and search data.</p>
                    <p v-if="authStore.role == 'cs'" class="text-neutral-400 text-sm">You are a Customer Service user. You can view all data.</p>
                </div>
             
                <Button @click="handleLogout" variant="danger" size="sm">
                    Logout
                </Button>
            </div>
            <div class="mt-6">
                <div class="grid grid-cols-1 gap-6 lg:grid-cols-12">
                    <section class="lg:col-span-9">
                        <Table title="Payments" subtitle="Incoming transactions" :searchable="true" :filterable="true">
                            <template v-if="authStore.role == 'operation'" #search>
                                <InputField
                                    type="number"
                                    v-model="searchValue"
                                    placeholder="Search by ID"
                                    size="sm"
                                />
                            </template>
                            <template v-if="authStore.role == 'operation'" #filters>
                                <div class="flex d-flex gap-2">
                                    <SelectField v-model="selectedFilter" size="sm">
                                        <option value="all">All</option>
                                        <option value="completed">Completed</option>
                                        <option value="processing">Processing</option>
                                        <option value="failed">Failed</option>
                                    </SelectField>
                                    <Button @click="handleReset" variant="warning" size="xs">
                                         {{ paymentStore.loading ? "Loading..." : "Reset" }}
                                    </Button>
                                </div>
                               
                            </template>

                            <template #head>
                                <tr>
                                  <th class="px-3 py-2">Payment ID</th>
                                  <th class="px-3 py-2">Merchant</th>
                                  <th class="px-3 py-2">Date</th>
                                  <th class="px-3 py-2">Amount</th>
                                  <th class="px-3 py-2">Status</th>
                                </tr>
                            </template>

                            <tr v-if="paymentStore.loading">
                              <td class="px-3 py-3 text-neutral-500" colspan="5">
                                Loading...
                              </td>
                            </tr>
                            <tr v-else-if="!data.length">
                              <td class="px-3 py-3 text-neutral-500" colspan="5">
                                No data
                              </td>
                            </tr>
                            <tr v-else v-for="item in data" :key="item.id">
                              <td class="px-3 py-3">{{ item.id }}</td>
                              <td class="px-3 py-3">{{ item.merchant }}</td>
                              <td class="px-3 py-3">{{ formatTimestamp(item.created_at) }}</td>
                              <td class="px-3 py-3">{{ item.amount }}</td>
                              <td class="px-3 py-3">{{ item.status }}</td>
                            </tr>
                        </Table>
                    </section>

                    <aside class="lg:col-span-3 space-y-6">
                        <div class="rounded-xl border border-neutral-700 bg-neutral-900 p-6">
                            <h3 class="text-white">Total</h3>
                            <h1 class="text-neutral-400">{{ totalWidget }}</h1>
                        </div>
                        <div v-for="wid in widget" class="rounded-xl border border-neutral-700 bg-neutral-900 p-6">
                            <h3 class="text-white">{{ capitalize(wid.status) }}</h3>
                            <h1 class="text-neutral-400">{{ wid.total }}</h1>
                        </div>
                    </aside>
                </div>
            </div>
        </div> 
    </div>
</template>
