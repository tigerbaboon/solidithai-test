import { defineStore } from 'pinia'

interface State {
  token: string
  userId: string
  openSideBar: boolean
  name: string
}

export const useIndexStore = defineStore('index', {
  state: (): State => ({
    token: '',
    userId: '',
    openSideBar: true,
    name: '',
  }),
})

// driver
// checker
