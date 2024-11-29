export const errorResp = (error: any) => {
  if (error) {
    var httpCode = error.status
    var message = error.data?.message || ''

    swalToast({
      icon: 'error',
      title: message,
    })

    if (httpCode == 403 || httpCode == 401) {
      setTimeout(() => {
        logoutWeb()
      }, 1000)
      return false
    }
  }
}
