<script setup>
defineProps({
  label: String,
  modelValue: [String, Number],
  type: { type: String, default: 'text' },
  placeholder: String
});

defineEmits(['update:modelValue']);
</script>

<template>
  <div class="flex flex-col gap-2">
    <label v-if="label" class="text-sm font-medium text-neutral-300">
      {{ label }}
    </label>
    
    <input
      :type="type"
      :value="modelValue"
      @input="$emit('update:modelValue', $event.target.value)"
      :placeholder="placeholder"
      class="block w-full rounded-md border border-neutral-700 bg-neutral-900 px-4 py-2 text-neutral-100 placeholder:text-neutral-400 shadow-sm outline-none focus:border-zinc-400 focus:ring-1 focus:ring-zinc-400 sm:text-sm"
    />
  </div>
</template>
