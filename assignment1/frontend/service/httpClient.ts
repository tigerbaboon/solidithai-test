import type { Option } from '~/models/axios.model'
import axios from 'axios'

export const client = (ops: Option) => {
  const config: any = useRuntimeConfig()

  let api: any = axios.create()

  api.defaults.baseURL = config.public.WEB_API

  api.defaults.headers.common['Content-Type'] = 'application/json'
  api.interceptors.request.use(
    async function (config: any) {
      let token = await useCookie('token')
      config.headers['Authorization'] = 'Bearer ' + token.value
      return config
    },
    function (error: any) {
      // return Promise.reject(error)
    },
  )
  api.defaults.validaStatus = false
  ops.method = ops.method.toLowerCase()
  return api({
    ...ops,
  })
}
