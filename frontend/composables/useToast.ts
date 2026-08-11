import { ref } from 'vue'

export interface ToastItem {
  id: string
  type: 'success' | 'error' | 'info' | 'warning'
  title?: string
  message: string
  duration?: number
}

const toasts = ref<ToastItem[]>([])

export const useToast = () => {
  const show = (options: Omit<ToastItem, 'id'>) => {
    const id = Math.random().toString(36).substring(2, 9)
    const toast: ToastItem = {
      id,
      type: options.type || 'info',
      title: options.title,
      message: options.message,
      duration: options.duration || 4000
    }

    toasts.value.push(toast)

    if (toast.duration > 0) {
      setTimeout(() => {
        remove(id)
      }, toast.duration)
    }
  }

  const success = (message: string, title: string = 'Thành công') => {
    show({ type: 'success', title, message })
  }

  const error = (message: string, title: string = 'Lỗi hệ thống') => {
    show({ type: 'error', title, message })
  }

  const info = (message: string, title: string = 'Thông báo') => {
    show({ type: 'info', title, message })
  }

  const warning = (message: string, title: string = 'Cảnh báo') => {
    show({ type: 'warning', title, message })
  }

  const remove = (id: string) => {
    toasts.value = toasts.value.filter(t => t.id !== id)
  }

  return {
    toasts,
    show,
    success,
    error,
    info,
    warning,
    remove
  }
}
