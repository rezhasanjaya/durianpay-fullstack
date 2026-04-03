<script setup lang="ts">
import { computed } from 'vue'

type Variant =
  | 'primary'
  | 'secondary'
  | 'success'
  | 'warning'
  | 'danger'
  | 'info'
  | 'light'
  | 'dark'

type Size = 'sm' | 'md' | 'lg'

const props = withDefaults(
  defineProps<{
    variant?: Variant
    size?: Size
    type?: 'button' | 'submit' | 'reset'
    block?: boolean
  }>(),
  {
    variant: 'primary',
    size: 'md',
    type: 'button',
    block: false,
  }
)

const base =
  'inline-flex items-center justify-center rounded-md font-medium transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2'

const sizes: Record<Size, string> = {
  sm: 'px-3 py-1.5 text-sm',
  md: 'px-4 py-2 text-sm',
  lg: 'px-5 py-2.5 text-base',
}

const variants: Record<Variant, string> = {
  primary: 'bg-indigo-600 text-white hover:bg-indigo-500 focus:ring-indigo-400',
  secondary: 'bg-slate-600 text-white hover:bg-slate-500 focus:ring-slate-400',
  success: 'bg-emerald-600 text-white hover:bg-emerald-500 focus:ring-emerald-400',
  warning: 'bg-amber-500 text-slate-900 hover:bg-amber-400 focus:ring-amber-300',
  danger: 'bg-rose-600 text-white hover:bg-rose-500 focus:ring-rose-400',
  info: 'bg-sky-600 text-white hover:bg-sky-500 focus:ring-sky-400',
  light: 'bg-white text-slate-900 hover:bg-slate-100 focus:ring-slate-200',
  dark: 'bg-slate-900 text-white hover:bg-slate-800 focus:ring-slate-700',
}

const classes = computed(() => [
  base,
  sizes[props.size],
  variants[props.variant],
  props.block ? 'w-full' : '',
])
</script>

<template>
  <button :type="type" :class="classes">
    <slot />
  </button>
</template>
