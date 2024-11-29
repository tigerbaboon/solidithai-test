import type { Login } from '~/models/auth.model'
import { client } from './httpClient'

export const login = (payload: Login) => {
  return client({
    url: '/auth/login',
    method: 'post',
    data: payload,
  })
}

export const logout = () => {
  return client({
    url: '/auth/logout',
    method: 'post',
  })
}
