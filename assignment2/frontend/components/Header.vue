<template>
  <header
    class="sticky top-0 z-10 flex flex-wrap justify-between gap-4 border-b border-zinc-300/25 bg-white/75 p-3 sm:p-3 md:p-6 backdrop-blur-md"
  >
    <div
      class="flex flex-wrap items-center gap-4 ltr:md:mr-auto rtl:md:ml-auto"
    >
      <button
        type="button"
        class="flex sm:flex md:hidden h-12 w-12 items-center justify-center"
        @click="store.openSideBar = !store.openSideBar"
      >
        <Icon name="fa:bars" class="text-xl" />
      </button>
    </div>

    <div
      class="flex flex-wrap items-center gap-4 ltr:md:ml-auto rtl:md:mr-auto"
    >
      <ClientOnly>
        <div>
          {{ store.name }}
        </div>
      </ClientOnly>

      <button
        v-if="!loading"
        type="button"
        class="btn-default-info flex items-center"
        @click="logout()"
      >
        <Icon name="material-symbols:logout" class="text-xl" />
      </button>
      <Icon v-else name="uiw:loading" class="animate-spin text-[1.6rem]" />
    </div>
  </header>
</template>

<script setup lang="ts">
import service from '~/service'
import { useIndexStore } from '~/store/main'

const store = useIndexStore()
const loading = ref(false)

const logout = async () => {
  loading.value = true
  await service.auth
    .logout()
    .then((resp: any) => {
      const { data } = resp.data

      logoutWeb()
    })
    .catch((err: any) => {
      errorResp(err.response)
    })
    .finally(() => {
      loading.value = false
    })
}
</script>

<style scoped></style>
