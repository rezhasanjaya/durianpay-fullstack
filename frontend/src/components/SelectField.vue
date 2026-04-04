<script setup>
defineProps({
  modelValue: [String, Number],
  size: { type: String, default: 'md' },
});

defineEmits(['update:modelValue']);
</script>

<template>
  <select
    :value="modelValue"
    @change="$emit('update:modelValue', $event.target.value)"
    :class="[
      'w-full rounded-md border border-neutral-700 bg-neutral-900 text-neutral-100 outline-none focus:border-zinc-400 focus:ring-1 focus:ring-zinc-400',
      size === 'sm' ? 'px-3 py-1.5 text-sm' : size === 'lg' ? 'px-4 py-2.5 text-base' : 'px-3 py-2 text-sm'
    ]"
  >
    <slot />
  </select>
</template>
