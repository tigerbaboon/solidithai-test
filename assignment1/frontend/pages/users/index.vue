<template>
  <Loading :loading="loading" />
  <Breadcrumb :items="title" />
  <div class="content">
    <div class="flex justify-end mb-4">
      <NuxtLink to="/users/create">
        <button class="btn primary md">
          <Icon name="mdi:plus" class="text-lg" /> Create User
        </button>
      </NuxtLink>
    </div>
    <div class="">
      <div class="overflow-auto">
        <table>
          <thead>
            <tr>
              <th class="">#</th>
              <th class="">Name</th>
              <th class="">Email</th>
              <th class="">Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(user, i) in users.datas" :key="i">
              <td class="">
                {{ i + 1 + (users.query.page - 1) * users.query.size }}
              </td>
              <td class="">{{ user.firstName }} {{ user.lastName }}</td>
              <td class="">
                {{ user.email }}
              </td>
              <td class="">
                <NuxtLink :to="`/users/${user.id}`">
                  <button class="btn primary sm">
                    <Icon name="mdi:magnify" class="text-lg" />
                    Detail
                  </button>
                </NuxtLink>
              </td>
            </tr>
          </tbody>
        </table>
        <Paginate
          :data="users.paginate"
          @setPage="setQuery('page', $event)"
          @setLimit="setQuery('limit', $event)"
          @reload="getUsers()"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Breadcrumb } from '~/models/app.model'
import type { UserList } from '~/models/user.model'
import service from '~/service'

definePageMeta({
  middleware: 'auth',
})

const title = ref<Breadcrumb[]>([])

const loading = ref(false)

const users = ref<UserList>({
  datas: [],
  query: {
    page: 1,
    size: 10,
    search: '',
  },
  paginate: null,
})

const getUsers = async () => {
  loading.value = true
  await service.user
    .List(users.value.query)
    .then((resp: any) => {
      const { data, paginate } = resp.data

      users.value.datas = data
      users.value.paginate = paginate

      title.value.push({
        text: 'users',
        to: '/users',
      })
    })
    .catch((err: any) => {
      errorResp(err.response)
    })
    .finally(() => {
      loading.value = false
    })
}

const setQuery = (prefix: any, event: any) => {
  switch (prefix) {
    case 'page':
      users.value.query.page = event
      break
    case 'limit':
      users.value.query.size = event
      break
    default:
      break
  }
}

onMounted(() => {
  getUsers()
})
</script>

<style scoped></style>
