<template>
  <div
    v-for="(user, i) in items"
    :key="i"
    class="bg-white rounded-lg px-3 py-2 shadow-lg flex gap-2 items-center"
  >
    <div class="relative">
      <img :src="user.avatar" alt="" class="w-10 h-10 rounded-full" />
      <span
        class="bottom-0 left-7 absolute w-3.5 h-3.5 bg-green-500 border-2 border-white rounded-full"
      ></span>
    </div>
    <div>{{ user.name }}</div>
  </div>
</template>

<script setup lang="ts">
import type { UserActive } from '~/models/user.model'

const props = defineProps({
  items: {
    type: Array as PropType<UserActive[]>,
  },
})
</script>

<style scoped></style>
