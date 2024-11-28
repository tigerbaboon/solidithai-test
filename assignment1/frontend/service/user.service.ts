import type { User } from '~/models/user.model'
import { client } from './httpClient'

export const GetInfo = () => {
  return client({
    url: '/users/info',
    method: 'get',
  })
}

export const List = (payload: any) => {
  return client({
    url: '/users',
    method: 'get',
    params: {
      ...payload,
    },
  })
}

export const Create = (payload: User) => {
  return client({
    url: '/users',
    method: 'post',
    data: payload,
  })
}

export const Update = (id: any, payload: User) => {
  return client({
    url: `/users/${id}`,
    method: 'put',
    data: payload,
  })
}

export const Delete = (id: any) => {
  return client({
    url: `/users/${id}`,
    method: 'delete',
  })
}

export const Get = (id: any) => {
  return client({
    url: `/users/${id}`,
    method: 'get',
  })
}

export const ResetPassword = (id: any, password: string) => {
  return client({
    url: `/users/${id}/password`,
    method: 'patch',
    data: {
      password: password,
    },
  })
}
