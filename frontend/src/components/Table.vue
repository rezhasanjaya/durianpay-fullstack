<script setup lang="ts">
defineProps<{
  title?: string
  subtitle?: string
  searchable?: boolean
  filterable?: boolean
}>()
</script>

<template>
  <div class="w-full rounded-md border border-neutral-700 bg-neutral-800/95 shadow-md shadow-black/20 ring-1 ring-neutral-700">
    <div class="border-b border-neutral-700/70 px-6 py-4">
      <div class="flex flex-col gap-3 md:flex-row md:items-center md:justify-between">
        <div>
          <div v-if="title" class="text-sm font-semibold text-neutral-100">{{ title }}</div>
          <div v-if="subtitle" class="text-xs text-neutral-400">{{ subtitle }}</div>
        </div>
        <div class="flex flex-col gap-2 sm:flex-row sm:items-center">
          <div v-if="searchable" class="w-full sm:w-64">
            <slot name="search" />
          </div>
          <div v-if="filterable" class="w-full sm:w-auto">
            <slot name="filters" />
          </div>
        </div>
      </div>
    </div>

    <div class="px-6 py-4">
      <div class="overflow-x-auto">
        <table class="min-w-full border-separate border-spacing-0 text-left text-sm text-neutral-200">
          <thead class="text-xs uppercase text-neutral-400">
            <slot name="head" />
          </thead>
          <tbody class="divide-y divide-neutral-800">
            <slot />
          </tbody>
        </table>
      </div>

      <div v-if="$slots.empty" class="py-8 text-center text-sm text-neutral-400">
        <slot name="empty" />
      </div>
    </div>
  </div>
</template>
