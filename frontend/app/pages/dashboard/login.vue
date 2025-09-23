<template>
  <div>
    <div>
      <h1 class="text-2xl font-bold mb-4">Login</h1>
      <p class="mb-4">Please enter your credentials to log in.</p>
    </div>
    <UForm :schema="schema" :state="loginData" class="space-y-4" @submit="onSubmit">
      <div class="mb-4">
        <UFormField label="Email" class="mb-2">
          <UInput placeholder="Enter your email" v-model="loginData.username" />
        </UFormField>

        <UFormField label="Password" class="mb-2">
          <UInput type="password" placeholder="Enter your password" v-model="loginData.password" />
        </UFormField>
      </div>
      <div class="mb-4">
        <UButton type="submit" :disabled="store.loading || !loginData.username || !loginData.password">Login</UButton>
      </div>
    </UForm>
  </div>
</template>

<script setup lang="ts">
import { object, string } from 'yup'
import type { InferType } from 'yup'
import type { FormSubmitEvent } from '@nuxt/ui'
import { useAuthStore } from '~/stores/auth'

const store = useAuthStore()
const toast = useToast()

const schema = object({
  username: string().required('Required'),
  password: string()
    .min(4, 'Must be at least 4 characters')
    .required('Required')
})

type Schema = InferType<typeof schema>

const loginData = reactive({
  username: undefined,
  password: undefined
})

async function onSubmit(event: FormSubmitEvent<Schema>) {
  event.preventDefault()
  try {
    const validatedData = await schema.validate(loginData, { abortEarly: false })
    console.log('Validated Data:', validatedData)
    await store.login(validatedData.username, validatedData.password)
    toast.add({
      title: 'Login successful',
      icon: 'i-material-symbols-check-circle-outline',
      color: 'success',
      description: `Welcome back, ${validatedData.username}!`
    })
    // Redirect to dashboard or another page if needed
    await navigateTo('/dashboard')
  } catch (err) {
    if (err instanceof Error) {
      console.error('Validation Error:', err)
    }
    toast.add({
        title: 'Login error',
        icon: 'i-material-symbols-error-outline',
        color: 'error',
        description: (err as any).data.message
      })
  }
}
</script>

<style scoped>

</style>
