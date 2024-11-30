import type { CreateMessage } from '~/models/message.model'
import { client } from './httpClient'

export const List = (payload: any) => {
  return client({
    url: '/messages',
    method: 'get',
    params: {
      ...payload,
    },
  })
}

export const Create = (payload: CreateMessage) => {
  return client({
    url: '/messages',
    method: 'post',
    data: payload,
  })
}
