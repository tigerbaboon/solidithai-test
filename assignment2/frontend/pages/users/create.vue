<template>
  <Loading :loading="loading"></Loading>
  <Breadcrumb :items="title" :back="'/users'" />
  <div class="content">
    <div class="flex justify-end gap-2">
      <button class="btn success md" @click="submit">
        <Icon name="mdi:check" class="text-lg" /> Submit
      </button>
    </div>
    <div class="flex justify-center">
      <div class="w-auto flex flex-col gap-4">
        <img
          :src="user.avatar"
          alt=""
          class="h-[150px] w-[150px] rounded-full"
        />
        <button class="btn primary md" @click="random">
          <template v-if="!loadingRan">
            <Icon name="mdi:reload" class="text-lg" />
          </template>
          <template v-else>
            <Icon name="uiw:loading" class="animate-spin text-2xl" />
          </template>
          Random Avatar
        </button>
      </div>
    </div>
    <div class="grid md:grid-cols-2 gap-4 mt-4">
      <div class="form-input">
        <label>Email <span class="text-red-500 text-sm">*</span></label>
        <input v-model="user.email" type="text" placeholder="Email" />
      </div>
      <div class="form-input">
        <label>Firstname <span class="text-red-500 text-sm">*</span></label>
        <input v-model="user.firstName" type="text" placeholder="Firstname" />
      </div>
      <div class="form-input">
        <label>Lastname <span class="text-red-500 text-sm">*</span></label>
        <input v-model="user.lastName" type="text" placeholder="Lastname" />
      </div>
      <div>
        <label
          >Password

          <span class="text-red-500 text-sm"
            >* password lenght more than or equal 8
          </span>
        </label>
        <div class="form-input">
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
            v-model="user.password"
            :type="(!showPassword && 'password') || 'text'"
            placeholder="Password"
          />
        </div>
      </div>

      <div>
        <label
          >Confirm Password <span class="text-red-500 text-sm">*</span></label
        >
        <div class="form-input">
          <div
            class="absolute right-0 flex justify-center items-center h-full pr-3"
            @click="showConfirmPassword = !showConfirmPassword"
          >
            <Icon
              :name="
                (!showConfirmPassword && 'heroicons-solid:eye') ||
                'flowbite:eye-slash-solid'
              "
              class="text-gray-500"
            />
          </div>
          <input
            v-model="confirmPassword"
            :type="(!showConfirmPassword && 'password') || 'text'"
            placeholder="Confirm Password"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Breadcrumb } from '~/models/app.model'
import type { User } from '~/models/user.model'
import service from '~/service'

definePageMeta({
  middleware: ['auth'],
})

const title = ref<Breadcrumb[]>([
  {
    text: 'users',
    to: '/users',
  },
  {
    text: 'create',
    to: '/users/create',
  },
])

const router = useRouter()

const loading = ref(false)
const loadingRan = ref(false)
const showPassword = ref<boolean>(false)
const showConfirmPassword = ref<boolean>(false)

const user = ref<User>({
  id: '',
  email: '',
  password: '',
  firstName: '',
  lastName: '',
  avatar: '',
})

const confirmPassword = ref<string>('')

const submit = async () => {
  if (
    !user.value.email ||
    !user.value.password ||
    !user.value.firstName ||
    !user.value.lastName ||
    !confirmPassword.value
  ) {
    return swalToast({
      icon: 'error',
      title: 'Please fill all fields',
    })
  }

  if (user.value.password !== confirmPassword.value) {
    return swalToast({
      icon: 'error',
      title: 'Password and Confirm Password not match',
    })
  }

  loading.value = true
  await service.user
    .Create(user.value)
    .then((resp: any) => {
      swalToast({
        icon: 'success',
        title: 'Success',
      })

      router.push('/users')
    })
    .catch((err: any) => {
      errorResp(err.response)
    })
    .finally(() => {
      loading.value = false
    })
}

const random = () => {
  loadingRan.value = true
  setTimeout(() => {
    user.value.avatar = randomAvatar()
    loadingRan.value = false
  }, 1500)
}

onMounted(() => {
  random()
})
</script>

<style scoped></style>
