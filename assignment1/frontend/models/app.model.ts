import { Breadcrumb } from './../.nuxt/components.d'
export interface MenuList {
  name: string
  icon: string
  to?: string
  show?: boolean
  childs?: Menu[]
}

export type Menu = Pick<MenuList, 'name' | 'icon' | 'to'>

export interface Breadcrumb {
  text: string
  to: string
}

export interface Query {
  page: number
  size: number
  search: string
  sort_by?: string
  order_by?: string
}

export interface Paginate {
  page: number
  size: number
  total: number
}
