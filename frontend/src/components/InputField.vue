<script setup>
defineProps({
  label: String,
  modelValue: [String, Number],
  type: { type: String, default: 'text' },
  size: { type: String, default: 'md' },
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
      :type="type === 'number' ? 'text' : type"
      :inputmode="type === 'number' ? 'numeric' : undefined"
      :pattern="type === 'number' ? '[0-9]*' : undefined"
      :value="modelValue"
      @beforeinput="(e) => {
        if (type !== 'number') return
        const inputType = e.inputType || ''
        if (inputType.startsWith('delete')) return
        const data = e.data ?? ''
        if (!/^[0-9]*$/.test(data)) e.preventDefault()
      }"
      @input="(e) => {
        const el = e.target
        if (type === 'number') {
          const cleaned = el.value.replace(/\\D+/g, '')
          if (el.value !== cleaned) el.value = cleaned
          $emit('update:modelValue', cleaned)
        } else {
          $emit('update:modelValue', el.value)
        }
      }"
      :placeholder="placeholder"
      :class="[
        'block w-full rounded-md border border-neutral-700 bg-neutral-900 text-neutral-100 placeholder:text-neutral-400 shadow-sm outline-none focus:border-zinc-400 focus:ring-1 focus:ring-zinc-400',
        size === 'sm' ? 'px-3 py-1.5 text-sm' : size === 'lg' ? 'px-4 py-2.5 text-base' : 'px-4 py-2 text-sm'
      ]"
    />
  </div>
</template>
