<template>
  <div
    v-if="props.show"
    class="flex justify-end items-center transition-all duration-1000 ease-in-out"
  >
    <div class="fixed inset-0 flex items-center justify-center z-40">
      <div
        class="modal-overlay absolute w-full h-full bg-black/50"
        @click="emit('close')"
      />

      <div
        class="rounded-md bg-white max-sm:w-[90%] max-md:w-[70%] w-[50%] z-40 px-4 py-2 overflow-y-auto"
      >
        <div class="flex justify-between items-center me-2 pt-1">
          <div class="font-bold text-xl mb-3">Reset Password</div>
          <button @click="emit('close')">
            <Icon name="ant-design:close" />
          </button>
        </div>

        <hr class="border my-2" />

        <div class="grid gap-4">
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
                v-model="password"
                :type="(!showPassword && 'password') || 'text'"
                placeholder="Password"
              />
            </div>
          </div>

          <div>
            <label
              >Confirm Password
              <span class="text-red-500 text-sm">*</span></label
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
          <button
            class="btn success md w-full"
            :disabled="loading"
            @click="reset"
          >
            <template v-if="!loading"> Submit </template>
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
    </div>
  </div>
  <!-- </Transition> -->
</template>

<script setup lang="ts">
import type { User } from '~/models/user.model'
import service from '~/service'

const props = defineProps({
  show: {
    type: Boolean,
    default: false,
  },
  user: {
    type: Object as PropType<User>,
    required: true,
  },
})

const loading = ref(false)
const showPassword = ref<boolean>(false)
const showConfirmPassword = ref<boolean>(false)

const password = ref<string>('')
const confirmPassword = ref<string>('')

const emit = defineEmits(['close', 'reload'])

const reset = async () => {
  if (password.value !== confirmPassword.value) {
    swalToast({
      icon: 'warning',
      title: 'Password not match',
    })
    return
  }

  loading.value = true
  await service.user
    .ResetPassword(props.user.id, password.value)
    .then((resp: any) => {
      swalToast({
        icon: 'success',
        title: 'Password has been reset',
      })
      emit('reload')
      emit('close')
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
