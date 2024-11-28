import { useIndexStore } from '~/store/main'
import { verify } from '../composables/useAuth'
const store = useIndexStore()

export default defineNuxtRouteMiddleware(async (to, from) => {
  verify()
  const refToken = useStatefulCookie('token')

  if (!refToken.value) {
    return navigateTo('/login')
  }
})
