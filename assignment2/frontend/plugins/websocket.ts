import { useIndexStore } from '~/store/main'

export default defineNuxtPlugin((nuxtApp) => {
  onNuxtReady(() => {
    const store = useIndexStore()
    const checkUserId = setInterval(() => {
      if (store.userId !== '') {
        const ws = new WebSocket(
          `ws://localhost:8080/api/ws?userID=${store.userId}`,
        )
        // ws.onopen = () => {
        //   console.log('WebSocket connected')
        // }

        nuxtApp.provide('ws', ws)
        store.isWsConnected = true
        clearInterval(checkUserId)
      }
    }, 1000)
  })
})
