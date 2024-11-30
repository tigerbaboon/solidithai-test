<template>
  <div class="w-[400px] shadow p-5 rounded-lg bg-white">
    <h1 class="text-2xl font-bold mb-5">Assignment 2</h1>
    <div class="flex flex-col gap-6">
      <div class="form-input">
        <div class="absolute flex justify-center items-center h-full pl-3">
          <Icon name="ph:user-fill" class="text-gray-500" />
        </div>
        <input v-model="email" type="text" placeholder="Email" class="group" />
      </div>
      <div class="form-input">
        <div class="absolute flex justify-center items-center h-full pl-3">
          <Icon name="teenyicons:password-solid" class="text-gray-500" />
        </div>
        <div
          class="absolute right-0 flex justify-center items-center h-full pr-3"
          @click="showPassword = !showPassword"
        >
          <Icon
            :name="
              (!showPassword && 'heroicons-solid:eye') ||
              'flowbite:eye-slash-solid'
            "
            class="text-gray-500"
          />
        </div>
        <input
          v-model="password"
          :type="(!showPassword && 'password') || 'text'"
          id="password"
          placeholder="Password"
          class="group"
          @keydown.enter="login"
        />
      </div>

      <button class="btn primary md w-full" :disabled="loading" @click="login">
        <template v-if="!loading"> Login </template>
        <template v-else>
          <Icon
            v-if="loading"
            name="uiw:loading"
            class="animate-spin text-2xl"
          />
        </template>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { swalToast } from '~/composables/useAlertDialog'
import { errorResp } from '~/composables/useErrorResp'
import service from '~/service'

definePageMeta({
  layout: 'auth',
  middleware: 'login',
})

const cookie = useStatefulCookie('token')
const router = useRouter()

const email = ref<string>('')
const password = ref<string>('')
const showPassword = ref<boolean>(false)
const loading = ref<boolean>(false)

const login = async () => {
  if (email.value == '' || password.value == '') {
    return swalToast({
      icon: 'warning',
      title: 'Email or Password is empty',
    })
  }

  loading.value = true

  await service.auth
    .login({
      email: email.value,
      password: password.value,
    })
    .then((resp: any) => {
      const { data } = resp.data

      cookie.value = data

      router.push('/')
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
