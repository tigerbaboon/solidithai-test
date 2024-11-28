import type { Paginate, Query } from './app.model'

export interface User {
  id: string
  email: string
  password?: string
  firstName: string
  lastName: string
}

export interface UserList {
  datas: User[]
  query: Query
  loading?: boolean
  paginate: Paginate | null
}
