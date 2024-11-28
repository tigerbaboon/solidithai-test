<template>
  <div class="flex grow flex-col h-screen w-full">
    <Sidebar />
    <div
      class="flex flex-col justify-between transition-all duration-300 ease-in-out h-screen w-full sm:pl-0 scrollbar-auto bg-zinc-100"
      :class="(store.openSideBar && 'md:pl-[20rem]') || 'md:pl-[6.225em]'"
    >
      <Header />
      <div class="grow p-5">
        <slot />
      </div>
      <Footer />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useIndexStore } from '~/store/main'

const store = useIndexStore()
</script>

<style scoped></style>
