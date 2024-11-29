import { verify } from '~/composables/useAuth'

export default defineNuxtRouteMiddleware(async (to, from) => {
  verify()
  const refToken = useStatefulCookie('token')
  if (refToken.value) {
    return navigateTo('/')
  }
})
