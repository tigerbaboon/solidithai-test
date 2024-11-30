<template>
  <Loading :loading="loading"></Loading>

  <Breadcrumb :items="title" :back="'/users'" />

  <div class="content">
    <div class="flex justify-end gap-2">
      <template v-if="isEdit">
        <button class="btn outline-secondary md" @click="isEdit = !isEdit">
          <Icon name="mdi:close" class="text-lg" />Cancle
        </button>
        <button class="btn success md" @click="update">
          <Icon name="mdi:check" class="text-lg" /> Submit
        </button>
      </template>
      <template v-else>
        <button
          class="btn primary md"
          @click="showModalResetPassword = !showModalResetPassword"
        >
          <Icon name="heroicons-solid:eye" class="text-lg" />Reset Password
        </button>

        <button class="btn warning md" @click="isEdit = !isEdit">
          <Icon name="mdi:pencil" class="text-lg" />Edit
        </button>

        <button class="btn danger md" @click="deleteUser">
          <Icon name="mdi:close" class="text-lg" />Delete
        </button>
      </template>
    </div>
    <div class="flex justify-center">
      <div class="w-auto flex flex-col gap-4">
        <template v-if="!isEdit">
          <img
            :src="user.avatar"
            alt=""
            class="h-[150px] w-[150px] rounded-full"
          />
        </template>
        <template v-else>
          <img
            :src="userEdit.avatar"
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
        </template>
      </div>
    </div>
    <div class="grid md:grid-cols-2 gap-4 mt-4">
      <div class="form-input">
        <label>Email <span class="text-red-500 text-sm">*</span></label>
        <input
          v-if="!isEdit"
          v-model="user.email"
          type="text"
          placeholder="Email"
          class="bg-gray-200"
          readonly
        />
        <input
          v-else
          v-model="userEdit.email"
          type="text"
          placeholder="Email"
        />
      </div>
      <div class="form-input">
        <label>Firstname <span class="text-red-500 text-sm">*</span></label>
        <input
          v-if="!isEdit"
          v-model="user.firstName"
          type="text"
          placeholder="Firstname"
          class="bg-gray-200"
          readonly
        />
        <input
          v-else
          v-model="userEdit.firstName"
          type="text"
          placeholder="Firstname"
        />
      </div>
      <div class="form-input">
        <label>Lastname <span class="text-red-500 text-sm">*</span></label>
        <input
          v-if="!isEdit"
          v-model="user.lastName"
          type="text"
          placeholder="Lastname"
          class="bg-gray-200"
          readonly
        />
        <input
          v-else
          v-model="userEdit.lastName"
          type="text"
          placeholder="Lastname"
        />
      </div>
    </div>
  </div>

  <ModalsResetPassword
    :show="showModalResetPassword"
    :user="user"
    @close="showModalResetPassword = false"
    @reload="getUser"
  ></ModalsResetPassword>
</template>

<script setup lang="ts">
import Swal from 'sweetalert2'
import type { Breadcrumb } from '~/models/app.model'
import type { User } from '~/models/user.model'
import service from '~/service'

definePageMeta({
  middleware: ['auth'],
})

const route = useRoute()
const router = useRouter()
const id = ref<any>(route.params.id)
const user = ref<User>({
  id: '',
  email: '',
  password: '',
  firstName: '',
  lastName: '',
  avatar: '',
})
const userEdit = ref<User>({
  id: '',
  email: '',
  password: '',
  firstName: '',
  lastName: '',
  avatar: '',
})

const loading = ref(true)
const isEdit = ref(false)
const showModalResetPassword = ref(false)
const title = ref<Breadcrumb[]>([])
const loadingRan = ref(false)
const getUser = async () => {
  loading.value = true
  await service.user
    .Get(id.value)
    .then((resp: any) => {
      const { data } = resp.data

      user.value.firstName = data.firstName
      user.value.lastName = data.lastName
      user.value.email = data.email
      user.value.id = data.id
      user.value.avatar = data.avatar

      userEdit.value = JSON.parse(JSON.stringify(user.value))

      title.value.push(
        {
          text: 'users',
          to: '/users',
        },
        {
          text: user.value.firstName + ' ' + user.value.lastName,
          to: `/users/${user.value.id}`,
        },
      )
    })
    .catch((err: any) => {
      errorResp(err.response)
    })
    .finally(() => {
      loading.value = false
    })
}

watch(
  () => isEdit,
  () => {
    if (!isEdit.value) {
      userEdit.value = JSON.parse(JSON.stringify(user.value))
    }
  },
  {
    deep: true,
  },
)

const update = async () => {
  if (
    !userEdit.value.email ||
    !userEdit.value.firstName ||
    !userEdit.value.lastName
  ) {
    return swalToast({
      icon: 'error',
      title: 'Please fill all fields',
    })
  }

  loading.value = true
  await service.user
    .Update(userEdit.value.id, userEdit.value)
    .then((resp: any) => {
      swalToast({
        icon: 'success',
        title: 'Success',
      })
      isEdit.value = false
      router.push('/users')
    })
    .catch((err: any) => {
      errorResp(err.response)
    })
    .finally(() => {
      loading.value = false
    })
}

const deleteUser = () => {
  Swal.fire({
    title: 'Are you sure?',
    text: "You won't be able to revert this!",
    icon: 'warning',
    showCancelButton: true,
    confirmButtonColor: '#F43F5E',
    confirmButtonText: 'Delete',
  }).then(async (result) => {
    if (result.isConfirmed) {
      loading.value = true
      await service.user
        .Delete(user.value.id)
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
  })
}

const random = () => {
  loadingRan.value = true
  setTimeout(() => {
    userEdit.value.avatar = randomAvatar()
    loadingRan.value = false
  }, 1500)
}

onMounted(() => {
  getUser()
})
</script>

<style scoped></style>
