<script setup lang="ts">
import { useAuthStore } from '../stores/auth'
import { usePaymentStore } from '../stores/payment'
import { useRouter } from 'vue-router'
import { onMounted, ref, watch } from 'vue'

import Button from '../components/Button.vue'
import Table from '../components/Table.vue'
import InputField from '../components/InputField.vue'

const authStore = useAuthStore()
const paymentStore = usePaymentStore()
const route = useRouter()
const selectedFilter = ref('all')
const searchValue = ref('')
const data = ref<any[]>([])

const handleLogout = () => {
    authStore.logout()
    route.push('/')
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

onMounted(() => {
    fetchPayments()
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
                </div>
             
                <Button @click="handleLogout" variant="danger" size="sm">
                    Logout
                </Button>
            </div>
            <div class="mt-6">
                <div class="grid grid-cols-1 gap-6 lg:grid-cols-12">
                    <section class="lg:col-span-9">
                        <Table title="Payments" subtitle="Incoming transactions" :searchable="true" :filterable="true">
                            <template #search>
                                <InputField
                                  type="text"
                                  v-model="searchValue"
                                  placeholder="Search by ID"
                                />
                            </template>
                            <template #filters>
                                <select
                                  class="w-full rounded-md border border-neutral-700 bg-neutral-900 px-3 py-2 text-sm text-neutral-100 outline-none focus:border-zinc-400 focus:ring-1 focus:ring-zinc-400"
                                  v-model="selectedFilter"
                                >
                                  <option value="all">All</option>
                                  <option value="completed">Completed</option>
                                  <option value="processing">Processing</option>
                                  <option value="failed">Failed</option>
                                </select>
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

                            <tr v-if="!data.length">
                              <td class="px-3 py-3 text-neutral-500" colspan="5">
                                No data
                              </td>
                            </tr>

                            <tr v-if="authStore.loading">
                              <td class="px-3 py-3 text-neutral-500" colspan="5">
                                Loading....
                              </td>
                            </tr>

                            <tr v-else v-for="item in data" :key="item.id">
                              <td class="px-3 py-3">{{ item.id }}</td>
                              <td class="px-3 py-3">{{ item.merchant }}</td>
                              <td class="px-3 py-3">{{ item.created_at || item.date }}</td>
                              <td class="px-3 py-3">{{ item.amount }}</td>
                              <td class="px-3 py-3">{{ item.status }}</td>
                            </tr>
                        </Table>
                    </section>

                    <aside class="lg:col-span-3 space-y-6">
                        <div class="rounded-xl border border-neutral-700 bg-neutral-900 p-6">
                            <h3 class="text-white">Total</h3>
                            <p class="text-neutral-400 text-sm">Total / Success / Failed</p>
                        </div>
                        <div class="rounded-xl border border-neutral-700 bg-neutral-900 p-6">
                            <h3 class="text-white">Success</h3>
                            <p class="text-neutral-400 text-sm">Total / Success / Failed</p>
                        </div>
                        <div class="rounded-xl border border-neutral-700 bg-neutral-900 p-6">
                            <h3 class="text-white">Failed</h3>
                            <h1 class="text-neutral-400">70</h1>
                        </div>
                    </aside>
                </div>
            </div>
            
           
        </div> 
    </div>
</template>
