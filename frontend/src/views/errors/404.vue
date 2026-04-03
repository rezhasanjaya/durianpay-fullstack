<script setup lang="ts">
import { useRouter } from 'vue-router'
import Button from '../../components/Button.vue'

const router = useRouter()
</script>

<template>
  <div class="">
    <div class="mx-auto flex min-h-screen max-w-6xl items-center justify-center px-6 py-10">
      <div class="w-full max-w-xl text-center">
        <div class="mx-auto mb-6 flex h-20 w-20 items-center justify-center rounded-2xl border border-neutral-700 bg-neutral-900 text-3xl font-bold text-white">
          404
        </div>
        <h1 class="text-3xl font-semibold text-white">Halaman Tidak Ditemukan</h1>
        <p class="mt-3 text-sm text-neutral-400">
          URL yang kamu tuju tidak ada.
        </p>
        <div class="mt-6 flex items-center justify-center gap-3">
          <Button variant="light" size="md" type="button" @click="router.back()">
            Kembali
          </Button>
          <Button variant="secondary" size="md" type="button" @click="router.push('/')">
            Ke Login
          </Button>
        </div>
      </div>
    </div>
  </div>
</template>
