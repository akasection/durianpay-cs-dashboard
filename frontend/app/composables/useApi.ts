import { useAuthStore } from '../stores/auth'

export function useApi() {
  const config = useRuntimeConfig().public
  const store = useAuthStore()

  const fetcher = $fetch.create({
    baseURL: config.apiBaseUrl ?? 'http://localhost:8081',
    onRequest({ options }) {
      if(options.method === 'GET') {
        options.params ??= {}
        options.params.token = store.token
      } else {
        options.headers.set('Authorization', `Bearer ${store.token}`)
      }
    },
  })

  return fetcher
}
