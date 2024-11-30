<template>
  <div class="flex gap-2 py-6 w-full" :class="item.isOwner && 'justify-end'">
    <template v-if="!item.isOwner">
      <img :src="item.avatar" alt="" class="w-10 h-10 rounded-full" />
      <div class="relative">
        <div class="-top-6 left-3 absolute text-sm text-slate-500 w-[250px]">
          {{ item.firstName }} {{ item.lastName }}
        </div>
        <div class="message">
          {{ item.text }}
        </div>
      </div>
    </template>
    <template v-else>
      <div class="relative">
        <div
          class="-top-6 right-3 absolute text-sm text-slate-500 w-[250px] text-end"
        >
          {{ item.firstName }} {{ item.lastName }}
        </div>
        <div class="message">
          {{ item.text }}
        </div>
      </div>
      <img :src="item.avatar" alt="" class="w-10 h-10 rounded-full" />
    </template>
  </div>
</template>

<script setup lang="ts">
import type { Message } from '~/models/message.model'

const props = defineProps({
  item: {
    type: Object as PropType<Message>,
    required: true,
  },
})
</script>

<style scoped>
.message {
  @apply px-3 py-2 bg-slate-200 rounded-xl break-words max-w-[250px];
}
</style>
