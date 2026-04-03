<script setup lang="ts">
defineProps<{
  title?: string
  subtitle?: string
  padded?: boolean
}>()
</script>

<template>
  <div class="rounded-lg border border-neutral-700 bg-neutral-800/80 shadow-sm">
    <div v-if="title || subtitle || $slots.header" class="border-b border-neutral-700 px-4 py-3">
      <slot name="header">
        <div class="text-sm font-semibold text-neutral-100">{{ title }}</div>
        <div v-if="subtitle" class="text-xs text-neutral-400">{{ subtitle }}</div>
      </slot>
    </div>

    <div :class="padded === false ? '' : 'px-4 py-4'">
      <slot />
    </div>

    <div v-if="$slots.footer" class="border-t border-neutral-700 px-4 py-3">
      <slot name="footer" />
    </div>
  </div>
</template>
