<script setup lang="ts">
import { ref } from 'vue'
import InputField from '../components/InputField.vue'
import Card from '../components/Card.vue'
import Button from '../components/Button.vue'
import { useAuthStore } from '../stores/auth'
import { useRouter } from 'vue-router'

const email = ref('')
const password = ref('')
const showPassword = ref(false)
const authStore = useAuthStore()
const route = useRouter()

const handleSubmit = async () => {
    try {
        await authStore.login(email.value, password.value)
        route.push('/dashboard')
    } catch (error) {
        //
    }
}
</script>

<template>
  <div>
    <div class="mx-auto flex min-h-screen max-w-6xl items-center justify-center px-6 py-10">
      <div class="w-full max-w-md">

        <div v-if="authStore.error" class="mb-3 text-sm text-red-400">
          {{ authStore.error }}
        </div>

        <Card
          variant="grey"
          :padded=true
        >
          <template #header>
            <h1 class="text-xl font-semibold text-slate-100">Login Page</h1>
          </template>

          <form class="space-y-4" 
            @submit.prevent="handleSubmit"
          >
            <InputField 
              label="Email" 
              type="email" 
              placeholder="Enter your email" 
              v-model="email" 
            />
            <InputField 
              label="Password" 
              :type="showPassword ? 'text' : 'password'" 
              placeholder="Enter your password" 
              v-model="password" 
            />
            <div class="flex justify-end">
                <button
                type="button"
                size="sm"
                class="justify-end text-xs"
                @click="showPassword = !showPassword"
                >
                {{ showPassword ? 'Hide Password' : 'Show Password' }}
                </button>
            </div>
           
            <Button v-if="!authStore.loading" class="mt-3" type="submit" variant="light" size="md" block>
              Login
            </Button>
            <Button v-else class="mt-3" type="button" variant="light" size="md" block disabled>
              Loading...
            </Button>
          </form>

        </Card>
      </div>
    </div>
  </div>
</template>
