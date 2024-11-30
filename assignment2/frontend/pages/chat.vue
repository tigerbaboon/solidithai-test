<template>
  <Loading :loading="loading" />
  <div class="content w-full">
    <div class="flex flex-col gap-4">
      <div ref="chatContainer" class="h-[calc(80vh-6rem)] overflow-y-auto">
        <div v-for="(message, i) in message.datas" :key="i" class="">
          <UserChat :item="message"></UserChat>
        </div>
      </div>
      <div class="flex gap-4">
        <div class="form-input flex-auto">
          <input v-model="upMessage" type="text" @keydown.enter="sendMessage" />
        </div>
        <button class="btn primary md" @click="sendMessage">Send</button>
      </div>
    </div>
  </div>
  <!-- <div class="text-gray-500 text-4xl mt-[350px]">ASSIGNMENT 1</div> -->
</template>

<script setup lang="ts">
import type {
  CreateMessage,
  Message,
  MessageList,
} from '~/models/message.model'
import type { UserChat } from '~/models/user.model'
import service from '~/service'
import { useIndexStore } from '~/store/main'

definePageMeta({
  middleware: 'auth',
})

const store = useIndexStore()
const loading = ref(false)
const upMessage = ref<string>('')
const userChats = ref<UserChat[]>([])
const chatContainer = ref<HTMLElement | null>(null)
const message = ref<MessageList>({
  datas: [],
  query: {
    page: 1,
    size: 999,
    search: '',
    order_by: 'desc',
  },
  paginate: null,
})

const doWs = () => {
  const { $ws }: any = useNuxtApp()

  //   $ws.onopen = (e: any) => {
  //     console.log('WebSocket connected')
  //   }

  $ws.onmessage = (e: any) => {
    const data = JSON.parse(e.data)

    if (data.type === 'message') {
      const temp: Message = {
        userId: data.message.userId,
        firstName: data.message.firstName,
        lastName: data.message.lastName,
        avatar: data.message.avatar,
        text: data.message.text,
        isOwner: store.userId == data.message.userId ? true : false,
        createdAt: Date.now() / 1000,
      }

      message.value.datas.push(temp)

      if (store.userId == data.message.userId) {
        nextTick(() => {
          scrollToBottom()
        })
      }
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

const getMessage = async () => {
  loading.value = true

  await service.message
    .List(message.value.query)
    .then((resp: any) => {
      const { data, paginate } = resp.data

      let messages: Message[] = []

      for (let i = 0; i < data.length; i++) {
        const e = data[i]
        const temp: Message = {
          id: e.id,
          userId: e.userId,
          firstName: e.firstName,
          lastName: e.lastName,
          avatar: e.avatar,
          text: e.text,
          createdAt: e.createdAt,
          isOwner: store.userId == e.userId ? true : false,
        }

        messages.push(temp)
      }

      messages.sort((a, b) => a.createdAt - b.createdAt)

      message.value.datas = messages
      message.value.paginate = paginate

      //   nextTick(() => {
      //     scrollToBottom()
      //   })
    })
    .catch((err: any) => {
      errorResp(err.response)
    })
    .finally(() => {
      loading.value = false
    })
}

const sendMessage = async () => {
  if (upMessage.value == '') {
    return
  }

  loading.value = true
  const temp: CreateMessage = {
    text: upMessage.value,
    userId: store.userId,
  }

  await service.message
    .Create(temp)
    .then((resp: any) => {
      const { data } = resp.data

      const { $ws }: any = useNuxtApp()

      $ws.send(upMessage.value)
      upMessage.value = ''

      nextTick(() => {
        scrollToBottom()
      })
    })
    .catch((err: any) => {
      errorResp(err.response)
    })
    .finally(() => {
      loading.value = false
    })
}

const scrollToBottom = () => {
  if (chatContainer.value) {
    chatContainer.value.scrollTop = chatContainer.value.scrollHeight
  }
}

onMounted(async () => {
  await getMessage()

  if (store.isWsConnected) {
    doWs()
  }

  await scrollToBottom()
})
</script>

<style scoped></style>
