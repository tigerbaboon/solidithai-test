import { useIndexStore } from '~/store/main'

import service from '~/service'
export const verify = async () => {
  var refToken: any = useStatefulCookie('token')

  if (refToken.value) {
    await service.user
      .GetInfo()
      .then(async (resp: any) => {
        const { data } = resp.data
        if (data != null) {
          const store: any = useIndexStore()
          store.$state.token = refToken.value
          store.$state.userId = data.id
          store.$state.name = data.firstName + ' ' + data.lastName
        }
      })
      .catch((err: any) => {
        errorResp(err.response)
      })
  }
}

export const logoutWeb = async () => {
  const router = useRouter()
  const token = useStatefulCookie('token')
  token.value = null
  router.push(`/login`)
}
