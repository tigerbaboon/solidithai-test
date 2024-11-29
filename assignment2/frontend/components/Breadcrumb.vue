<template>
  <div class="flex justify-between items-center p-4 mb-4">
    <div class="flex">
      <div
        v-for="(item, i) in props.items"
        :key="i"
        class="text-2xl font-bold uppercase"
      >
        <NuxtLink :to="item.to" class="hover:text-gray-500">
          {{ item.text }}
        </NuxtLink>
        <span v-if="i < props.items.length - 1" class="mx-2"> > </span>
      </div>
    </div>

    <NuxtLink v-if="back" :to="back">
      <button class="btn outline-secondary md">
        <Icon name="mdi:arrow-left" class="text-lg" /> Back
      </button>
    </NuxtLink>
  </div>
</template>

<script setup lang="ts">
import type { Breadcrumb } from '~/models/app.model'

const props = defineProps({
  items: {
    type: Array as PropType<Breadcrumb[]>,
    required: true,
  },
  back: {
    type: String,
  },
})
</script>

<style scoped></style>
