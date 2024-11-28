<template>
  <div class="flex justify-center items-center absolute" v-if="props.loading">
    <div
      class="fixed inset-0 flex items-center justify-center z-50 bg-black/40"
    >
      <div
        class="lg:pl-20 modal-overlay w-full h-full absolute flex justify-center items-center"
      >
        <div class="absolute w-full h-full"></div>

        <!-- Modern Circular Loader -->
        <div
          class="w-16 h-16 border-4 border-dotted rounded-full animate-spin border-t-blue-500 border-r-transparent"
        ></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  loading: { type: Boolean, default: false },
})
</script>

<style scoped></style>
