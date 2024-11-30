import type { Paginate, Query } from './app.model'

export interface User {
  id: string
  email: string
  password?: string
  firstName: string
  lastName: string
  avatar: string
}

export interface UserList {
  datas: User[]
  query: Query
  loading?: boolean
  paginate: Paginate | null
}

export interface UserActive {
  id: string
  name: string
  avatar: string
}
export interface UserChat {
  id: string
  name: string
  avatar: string
  message: string
  isOwner: boolean
}
