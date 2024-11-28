import Swal, { type SweetAlertIcon } from 'sweetalert2'

interface Toast {
  icon: string
  title: string
}

export const swalToast = (payload: Toast) => {
  Swal.fire({
    toast: true,
    position: 'top-end',
    icon: payload.icon as SweetAlertIcon,
    title: payload.title,
    showConfirmButton: false,
    timer: 2500,
    timerProgressBar: true,
    didOpen: (toast) => {
      toast.addEventListener('mouseenter', Swal.stopTimer)
      toast.addEventListener('mouseleave', Swal.resumeTimer)
    },
  })

  if (payload.title === 'Unauthorized') {
    setTimeout(() => {
      logoutWeb()
    }, 1000)
  }
}
