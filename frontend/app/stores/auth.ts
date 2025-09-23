import { defineStore } from 'pinia'
import type { BaseResponse } from '~/../types/http'

export type User = {
  username: string
  roles: string[]
  iat: number // Date
  exp: number // Date
  iss: string
}

export const useAuthStore = defineStore('auth', () => {
  const user = useCookie<User | undefined>('user_info', {
    default: () => ref(undefined),
  })
  const token = useCookie<string>('auth_token', {
    expires: undefined,
    default: () => ref(''),
  })
  const loading = ref(false)
  const isLoggedIn = computed(() => !!token.value && !!user.value)
  const toast = useToast()

  const api = useApi()

  async function logout() {
    user.value = undefined
    token.value = ''
    await navigateTo('/dashboard/login')
    toast.add({
      title: 'Logged out',
      icon: 'i-material-symbols-check-circle-outline',
      color: 'info',
      description: 'You have been logged out.',
    })
  }

  async function login(username: string, password: string) {
    loading.value = true
    if(!loading.value) return
    try {
      const response = await api<BaseResponse<{ token: string; user: User }>>('/dashboard/v1/auth/login', {
        method: 'POST',
        body: {
          username,
          password,
        },
      })
      token.value = response.data.token
      user.value = response.data.user
    } catch (error) {
      user.value = undefined
      throw error
    } finally {
      loading.value = false
    }
  }
  return {
    user,
    token,
    loading,
    isLoggedIn,
    login,
    logout,
  }
})
