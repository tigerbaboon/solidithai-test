import { defineStore } from 'pinia'

interface State {
  token: string
  userId: string
  openSideBar: boolean
  name: string
  isWsConnected: boolean
}

export const useIndexStore = defineStore('index', {
  state: (): State => ({
    token: '',
    userId: '',
    openSideBar: true,
    name: '',
    isWsConnected: false,
  }),
})

// driver
// checker
