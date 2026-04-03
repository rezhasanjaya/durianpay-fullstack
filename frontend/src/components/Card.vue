<script setup lang="ts">
import { computed } from 'vue'

type Variant = 'grey' | 'black'

const props = withDefaults(
  defineProps<{
    padded?: boolean
    variant?: Variant
  }>(),
  {
    padded: true,
    variant: 'grey',
  }
)

const base ='rounded-xl w-full rounded-md border border-neutral-700 shadow-sm'

const variantClass = computed(() =>
  props.variant === 'black' ? 'bg-neutral-900' : 'bg-neutral-800/95'
)
</script>

<template>
  <div :class="[base, variantClass, padded === false ? '' : 'p-10']">
    <div v-if="$slots.header" class="mb-2">
      <slot name="header" />
    </div>

    <div :class="padded === false ? '' : 'py-2'">
      <slot />
    </div>

    <div v-if="$slots.footer" class="mt-4">
      <slot name="footer" />
    </div>
  </div>
</template>
