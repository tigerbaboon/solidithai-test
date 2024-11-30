<template>
  <div
    class="fixed bottom-0 top-0 z-30 md:z-20 flex flex-col border-zinc-300/25 bg-white py-6 ltr:border-r rtl:border-l transition-all duration-300 ease-in-out max-md:shadow-2xl ltr:max-md:-left-[20rem] rtl:max-md:-right-[20rem] overflow-hidden"
    :class="
      (store.openSideBar && 'w-[20rem] sm:w-[20rem] md:w-[20rem]') ||
      'w-0 sm:w-0 md:w-[6.225em]'
    "
  >
    <div class="flex flex-col">
      <div
        class="flex align-middle items-center px-6 pb-6"
        :class="(store.openSideBar && 'justify-between') || 'justify-center'"
      >
        <a :class="(store.openSideBar && 'block') || 'hidden'" href="/">
          <Icon name="logos:aws" class="text-4xl" />
        </a>

        <button
          type="button"
          class="flex h-12 w-12 items-center justify-center"
          @click="store.openSideBar = !store.openSideBar"
        >
          <Icon
            :name="
              (store.openSideBar && 'ant-design:bars-outlined') ||
              'gravity-ui:bars'
            "
            class="text-2xl"
          />
        </button>
      </div>

      <nav class="h-full px-6 no-scrollbar">
        <ul>
          <li
            v-for="(menu, menuIndex) in menus"
            :key="menuIndex"
            class="cursor-pointer flex list-none items-center overflow-hidden whitespace-nowrap pb-1"
          >
            <template v-if="!menu?.childs">
              <nuxt-link
                :to="menu.to"
                class="menu-bar"
                @click="isMobile && hiddenMenu()"
              >
                <Icon
                  :name="menu.icon"
                  class="w-6 flex-none me-2 ms-1 text-xl"
                />
                <div class="flex w-full items-center justify-between truncate">
                  <div class="overflow-hidden truncate whitespace-nowrap">
                    {{ menu.name }}
                  </div>
                </div>
              </nuxt-link>
            </template>
            <template v-else>
              <div class="flex flex-col w-full">
                <div
                  class="menu-bar grow"
                  :class="(menu.show && 'menu-active') || ''"
                  @click="menu.show = !menu.show"
                >
                  <Icon
                    :name="menu.icon"
                    class="w-6 flex-none me-2 ms-1 text-xl"
                  />
                  <div
                    class="flex w-full items-center justify-between truncate"
                  >
                    <div class="overflow-hidden truncate whitespace-nowrap">
                      {{ menu.name }}
                    </div>

                    <Icon
                      name="mingcute:down-fill"
                      class="transition ease-in-out delay-200"
                      :class="(menu.show && 'rotate-0') || 'rotate-90'"
                    />
                  </div>
                </div>
                <ul
                  class="transition ease-in-out delay-200 overflow-hidden"
                  :class="[
                    (store.openSideBar && 'ml-5') || 'ml-0',
                    (menu.show && 'h-auto') || 'h-0',
                  ]"
                >
                  <li
                    v-for="(child, childIndex) in menu.childs"
                    :key="childIndex"
                    class="cursor-pointer flex list-none items-center overflow-hidden whitespace-nowrap pb-1"
                  >
                    <nuxt-link
                      :to="child.to"
                      class="menu-bar"
                      @click="isMobile && hiddenMenu()"
                    >
                      <Icon
                        :name="child.icon"
                        class="w-6 flex-none me-2 ms-1"
                      />
                      <div
                        class="flex w-full items-center justify-between truncate"
                      >
                        <div class="overflow-hidden truncate whitespace-nowrap">
                          {{ child.name }}
                        </div>
                      </div>
                    </nuxt-link>
                  </li>
                </ul>
              </div>
            </template>
          </li>
        </ul>
      </nav>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { MenuList } from '~/models/app.model'
import { useIndexStore } from '~/store/main'

const store = useIndexStore()
const route = useRoute()
const router = useRouter()

const screenWidth = ref<number>(0)
const screenWidthMobile = ref<number>(1280)

const menus = ref<MenuList[]>([
  { name: 'Home', icon: 'mdi:home-variant', to: '/' },
  { name: 'Chat', icon: 'mdi:message-processing', to: '/chat' },
  { name: 'User', icon: 'mdi:account', to: '/users' },
])

const isMobile = computed(() => {
  return screenWidth.value <= screenWidthMobile.value
})

const initDate = () => {
  updateScreenWidth()
  window.addEventListener('resize', updateScreenWidth)
}

const updateScreenWidth = () => {
  screenWidth.value = window.innerWidth
  store.openSideBar = isMobile.value ? false : true
}

const hiddenMenu = () => {
  store.openSideBar = false
}

watch(
  () => route.path,
  () => {
    checkMenuActive()
  },
  {
    deep: true,
  },
)

const checkMenuActive = () => {
  for (let i = 0; i < menus.value.length; i++) {
    const menu = menus.value[i]

    if (menu?.childs?.length) {
      let showMenu = false
      for (let ii = 0; ii < menu.childs.length; ii++) {
        const child = menu.childs[ii]
        if (child.to === route.path) {
          showMenu = true
        }
      }

      menu.show = showMenu
    }
  }
}

onMounted(() => {
  if (typeof window !== 'undefined') {
    initDate()
  }

  checkMenuActive()
})
</script>

<style lang="scss" scoped>
.menu-bar {
  @apply p-2 flex justify-start align-middle items-center cursor-pointer overflow-hidden rounded-xl border text-zinc-500 hover:text-zinc-950 grow transition-all duration-300 ease-in-out border-transparent;

  &.router-link-exact-active {
    @apply text-zinc-950 border-zinc-300;
  }
  &.menu-active {
    @apply text-zinc-950;
  }
}
</style>
