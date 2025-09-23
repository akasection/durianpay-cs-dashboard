<template>
  <li class="px-4 py-2 hover:bg-green-200 text-green-800" :class="[ disabled ? 'opacity-50 cursor-not-allowed' : 'cursor-pointer' ]">
    <UIcon :name="icon" class="size-5 inline-block mr-2 align-middle" />
    <p class="inline-block align-middle uppercase tracking-widest text-xs font-semibold">
      <slot />
    </p>
  </li>
</template>

<script setup lang="ts">
withDefaults(defineProps<{
  icon?: string
  disabled?: boolean
}>(), {
  icon: 'i-material-symbols-arrow-circle-right',
  disabled: false
})
</script>
