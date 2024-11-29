import type { CookieOptions } from 'nuxt/app'

export const useStatefulCookie = <T>(
  key: string,
  options?: CookieOptions<T>
) => {
  const cookie = useCookie<T | undefined>(key)

  const state = useState<T | undefined>(key, () => cookie.value)

  watch(state, () => {
    cookie.value = state.value
  })

  return state
}
