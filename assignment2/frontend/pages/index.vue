<template>
  <Loading :loading="loading" />
  <Breadcrumb :items="title" />
  <div class="flex flex-col flex-wrap md:flex-row gap-4">
    <UserActive :items="userAcives"></UserActive>
  </div>
</template>

<script setup lang="ts">
import type { Breadcrumb } from '~/models/app.model'
import type { UserActive } from '~/models/user.model'
import service from '~/service'
import { useIndexStore } from '~/store/main'

definePageMeta({
  middleware: 'auth',
})

const title = ref<Breadcrumb[]>([
  {
    text: 'online users',
    to: '/',
  },
])

const store = useIndexStore()
const loading = ref(false)
let userAcives = ref<UserActive[]>([])

const doWs = () => {
  const { $ws }: any = useNuxtApp()

  // $ws.onopen = (e: any) => {
  //   console.log('WebSocket connected')
  // }

  $ws.onmessage = (e: any) => {
    const data = JSON.parse(e.data)

    if (data.type == 'users') {
      let users: UserActive[] = []
      for (let i = 0; i < data.users.length; i++) {
        const e = data.users[i]
        const temp: UserActive = {
          id: e.id,
          name: e.firstName + ' ' + e.lastName,
          avatar: e.avatar,
        }

        users.push(temp)
      }
      userAcives.value = users
    }
  }
}

watch(
  () => store.isWsConnected,
  () => {
    doWs()
  },
  {
    deep: true,
  },
)

const getUserActive = async () => {
  loading.value = true

  await service.user
    .GetActive()
    .then((resp: any) => {
      const { data } = resp.data

      let users: UserActive[] = []

      for (let i = 0; i < data.length; i++) {
        const e = data[i]

        const temp: UserActive = {
          id: e.id,
          name: e.firstName + ' ' + e.lastName,
          avatar: e.avatar,
        }

        users.push(temp)
      }

      userAcives.value = users
    })
    .catch((err: any) => {
      errorResp(err.response)
    })
    .finally(() => {
      loading.value = false
    })
}

onMounted(() => {
  getUserActive()

  if (store.isWsConnected) {
    doWs()
  }
})
</script>

<style scoped></style>
