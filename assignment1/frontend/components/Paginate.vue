<template>
  <div
    class="mt-6 flex justify-center flex-col md:flex-row md:justify-between text-xs items-center"
  >
    <div>
      Showing {{ rows != 0 ? (currentPage - 1) * perPage + 1 : 0 }} to
      {{ currentPage * perPage > rows ? rows : currentPage * perPage }}
      of {{ rows }} entries
    </div>
    <div class="flex items-center mt-4 md:mt-0">
      <div class="mr-2">Display</div>
      <select
        v-model="perPage"
        class="cursor-pointer hover:bg-white/5 w-12 mr-2 text-xs h-7 rounded-lg"
        @change="setLimit()"
      >
        <option
          class=""
          v-for="(option, index) in options"
          :value="option"
          :key="index"
        >
          {{ option }}
        </option>
      </select>

      <button
        :class="` cursor-pointer hover:bg-white/5  p-2 rounded-l-lg ${
          currentPage <= 1
            ? ''
            : 'cursor-pointer hover:bg-white/5 rounded-l-lg '
        } `"
        :disabled="currentPage <= 1"
        @click="changeCurrentPage('first')"
      >
        <svg
          fill="none"
          stroke="currentColor"
          stroke-width="1.5"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
          aria-hidden="true"
          class="h-3 w-3"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M18.75 19.5l-7.5-7.5 7.5-7.5m-6 15L5.25 12l7.5-7.5"
          />
        </svg>
      </button>

      <button
        :class="` cursor-pointer hover:bg-white/5 p-2 ${
          currentPage <= 1 ? '' : 'cursor-pointer hover:bg-white/5 '
        } `"
        :disabled="currentPage <= 1"
        @click="increaseCurrentPage(false)"
      >
        <svg
          fill="none"
          stroke="currentColor"
          stroke-width="1.5"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
          aria-hidden="true"
          class="w-3 h-3"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M15.75 19.5L8.25 12l7.5-7.5"
          ></path>
        </svg>
      </button>

      <button class="cursor-pointer hover:bg-white/5 p-1.5">
        {{ currentPage }}/{{ maxPage }}
      </button>
      <button
        :class="` cursor-pointer hover:bg-white/5 p-2  ${
          currentPage >= maxPage ? '' : 'cursor-pointer hover:bg-white/5 '
        }`"
        :disabled="currentPage >= maxPage"
        @click="increaseCurrentPage(true)"
      >
        <svg
          fill="none"
          stroke="currentColor"
          stroke-width="1.5"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
          aria-hidden="true"
          class="w-3 h-3"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M8.25 4.5l7.5 7.5-7.5 7.5"
          ></path>
        </svg>
      </button>

      <button
        :class="` cursor-pointer hover:bg-white/5 p-2 rounded-r-lg ${
          currentPage >= maxPage
            ? ''
            : 'cursor-pointer hover:bg-white/5 rounded-r-lg'
        }`"
        :disabled="currentPage >= maxPage"
        @click="changeCurrentPage('last')"
      >
        <svg
          fill="none"
          stroke="currentColor"
          stroke-width="1.5"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
          aria-hidden="true"
          class="w-3 h-3"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M11.25 4.5l7.5 7.5-7.5 7.5m-6-15l7.5 7.5-7.5 7.5"
          ></path>
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
const props = defineProps({
  data: { type: Object as PropType<any>, default: {} },
})

const rows = ref<number>(0)
const perPage = ref<number>(10)
const currentPage = ref<number>(0)
const options = ref<number[]>([10, 20, 50, 100])

const emit = defineEmits(['set-page', 'set-limit', 'reload'])

const maxPage = computed(() => {
  try {
    const maxPage =
      Math.floor(rows.value / perPage.value) +
      (rows.value % perPage.value !== 0 ? 1 : 0)
    return maxPage
  } catch {
    return 1
  }
})

watch(
  () => props.data,
  () => {
    initData()
  },
)

const initData = () => {
  if (props.data != null && Object.keys(props.data).length > 0) {
    const val = JSON.parse(JSON.stringify(props.data))
    perPage.value = val.size
    rows.value = val.total

    let page = 0

    if (val.page) {
      page = val.page
      const maxPage = Math.floor(val.total / val.limit) + 1
      currentPage.value = page > maxPage ? maxPage : page
    }

    if (val.from) {
      page = val.from
      const maxPage = Math.floor(val.total / val.limit) + 1
      currentPage.value = page > maxPage ? maxPage : page
    }
  }
}

const increaseCurrentPage = (val: boolean) => {
  currentPage.value += val ? 1 : -1
  checkCurrentPage()
}

const checkCurrentPage = () => {
  currentPage.value = currentPage.value <= 0 ? 1 : currentPage.value
  currentPage.value =
    currentPage.value >= maxPage.value ? maxPage.value : currentPage.value
  setPage()
}
const changeCurrentPage = (type: string) => {
  if (type === 'first') {
    currentPage.value = 1
  } else if (type === 'last') {
    currentPage.value = maxPage.value
  }

  setPage()
}

const setLimit = () => {
  emit('set-page', 1)
  emit('set-limit', perPage.value)
  emit('reload')
}
const setPage = () => {
  emit('set-page', currentPage.value)
  emit('reload')
}

onMounted(() => {
  initData()
})
</script>
