import type { Paginate, Query } from './app.model'

export interface Message {
  id?: string
  userId: string
  firstName: string
  lastName: string
  avatar: string
  text: string
  createdAt: number
  isOwner: boolean
}

export interface MessageList {
  datas: Message[]
  query: Query
  loading?: boolean
  paginate: Paginate | null
}

export interface CreateMessage {
  userId: string
  text: string
}
