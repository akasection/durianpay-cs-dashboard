<template>
  <div>

  </div>
</template>

<script setup lang="ts">
definePageMeta({
  middleware: 'auth-logged-in'
})

await navigateTo('/dashboard/payment')
</script>

<style scoped>

</style>
