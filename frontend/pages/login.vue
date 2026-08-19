<template>
  <div class="relative min-h-[calc(100vh-4rem)] overflow-hidden bg-slate-950 px-4 py-10 sm:px-6 lg:px-8">
    <div class="absolute inset-0 bg-[radial-gradient(circle_at_20%_15%,rgba(34,211,238,0.18),transparent_28rem),radial-gradient(circle_at_85%_80%,rgba(14,116,144,0.16),transparent_30rem),linear-gradient(180deg,#020617_0%,#0f172a_50%,#020617_100%)]"></div>

    <div class="relative mx-auto grid min-h-[calc(100vh-5rem)] max-w-6xl items-center gap-8 lg:grid-cols-[1.05fr_0.95fr]">
      <section class="hidden lg:block">
        <div class="max-w-xl space-y-7">
          <NuxtLink to="/" class="inline-flex items-center gap-3">
            <span class="h-11 w-11 rounded-xl bg-gradient-to-tr from-cyan-400 to-slate-700 flex items-center justify-center text-slate-950 font-extrabold text-xl shadow-lg shadow-cyan-500/20">
              Q
            </span>
            <span class="text-2xl font-extrabold text-white">QuickWork</span>
          </NuxtLink>

          <div class="space-y-4">
            <p class="inline-flex rounded-full border border-cyan-400/20 bg-cyan-400/10 px-3 py-1 text-xs font-bold uppercase tracking-wide text-cyan-200">
              Truy cập không gian làm việc an toàn
            </p>
            <h1 class="text-5xl font-extrabold leading-none tracking-tight text-white">
              Tiếp tục công việc & kết nối tuyển dụng của bạn.
            </h1>
            <p class="text-base font-medium leading-8 text-slate-300">
              Sinh viên có thể quản lý danh sách ứng tuyển và ca làm việc. Nhà tuyển dụng quản lý hồ sơ ứng viên, đăng bài và trao đổi trực tiếp qua chat.
            </p>
          </div>

          <div class="grid grid-cols-3 gap-3">
            <div class="rounded-xl border border-white/10 bg-white/[0.06] p-4">
              <p class="text-2xl font-extrabold text-white">24/7</p>
              <p class="mt-1 text-xs font-semibold text-slate-400">Truy cập</p>
            </div>
            <div class="rounded-xl border border-white/10 bg-white/[0.06] p-4">
              <p class="text-2xl font-extrabold text-white">Nhanh chóng</p>
              <p class="mt-1 text-xs font-semibold text-slate-400">Nộp đơn</p>
            </div>
            <div class="rounded-xl border border-white/10 bg-white/[0.06] p-4">
              <p class="text-2xl font-extrabold text-white">Trực tiếp</p>
              <p class="mt-1 text-xs font-semibold text-slate-400">Nhắn tin</p>
            </div>
          </div>
        </div>
      </section>

      <section class="mx-auto w-full max-w-md">
        <div class="rounded-3xl border border-white/10 bg-slate-900/86 p-7 shadow-2xl shadow-slate-950/50 backdrop-blur">
          <div>
            <div class="flex justify-center lg:hidden">
              <div class="h-12 w-12 rounded-xl bg-gradient-to-tr from-cyan-400 to-slate-700 flex items-center justify-center text-slate-950 font-extrabold text-2xl shadow-lg shadow-cyan-500/20">
                Q
              </div>
            </div>
            <h2 class="mt-2 text-center text-3xl font-extrabold text-white tracking-tight lg:mt-0">
              Đăng Nhập QuickWork
            </h2>
            <p class="mt-2 text-center text-sm text-slate-300">
              Nhập thông tin tài khoản để truy cập vào hệ thống QuickWork.
            </p>
          </div>

          <!-- Alert for error messages -->
          <div v-if="errorMessage" class="mt-4 bg-rose-400/10 border border-rose-400/30 text-rose-100 px-4 py-3 rounded-xl text-xs flex items-center space-x-3" role="alert">
            <svg class="h-5 w-5 text-red-400 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <span class="font-medium">{{ errorMessage }}</span>
          </div>

          <form class="mt-6 space-y-5" @submit.prevent="handleLogin">
            <div class="space-y-4">
              <!-- Email Address -->
              <div>
                <label for="email-address" class="block text-xs font-extrabold text-cyan-300 uppercase tracking-wider mb-1.5">Địa chỉ Email</label>
                <div class="relative">
                  <div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none">
                    <svg class="h-4 w-4 text-slate-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                    </svg>
                  </div>
                  <input
                    id="email-address"
                    name="email"
                    type="email"
                    autocomplete="email"
                    required
                    v-model="form.email"
                    class="appearance-none relative block w-full pl-10 pr-3.5 py-2.5 border border-white/10 placeholder-slate-500 text-white rounded-xl bg-slate-950/70 hover:bg-slate-950 focus:outline-none focus:ring-2 focus:ring-cyan-400 focus:border-transparent transition-all duration-200 text-xs font-medium"
                    placeholder="nhap.email@domain.com"
                  />
                </div>
              </div>

              <!-- Password -->
              <div>
                <label for="password" class="block text-xs font-extrabold text-cyan-300 uppercase tracking-wider mb-1.5">Mật khẩu</label>
                <div class="relative">
                  <div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none">
                    <svg class="h-4 w-4 text-slate-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 7a2 2 0 012 2m-2 2a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2V9a2 2 0 00-2-2h-2zm0 0V5.5A2.5 2.5 0 0012.5 3h-5A2.5 2.5 0 005 5.5V7m10 4V9" />
                    </svg>
                  </div>
                  <input
                    id="password"
                    name="password"
                    type="password"
                    autocomplete="current-password"
                    required
                    v-model="form.password"
                    class="appearance-none relative block w-full pl-10 pr-3.5 py-2.5 border border-white/10 placeholder-slate-500 text-white rounded-xl bg-slate-950/70 hover:bg-slate-950 focus:outline-none focus:ring-2 focus:ring-cyan-400 focus:border-transparent transition-all duration-200 text-xs font-medium"
                    placeholder="••••••••"
                  />
                </div>
              </div>
            </div>

            <div class="flex items-center justify-between">
              <div class="flex items-center">
                <input
                  id="remember-me"
                  name="remember-me"
                  type="checkbox"
                  class="h-4 w-4 text-cyan-400 focus:ring-cyan-400 border-slate-700 rounded bg-slate-900 cursor-pointer"
                />
                <label for="remember-me" class="ml-2 block text-sm text-slate-300 font-medium cursor-pointer">
                  Ghi nhớ đăng nhập
                </label>
              </div>

              <div class="text-sm">
                <button @click.prevent="isForgotPasswordOpen = true" class="font-bold text-cyan-300 hover:text-cyan-200 transition-colors">
                  Quên mật khẩu?
                </button>
              </div>
            </div>

            <div>
              <button
                type="submit"
                :disabled="isLoading || lockSeconds > 0"
                class="group relative w-full flex justify-center py-3 px-4 border border-transparent text-sm font-bold rounded-lg text-slate-950 bg-cyan-400 hover:bg-cyan-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-400 transition-all duration-150 disabled:opacity-50 disabled:cursor-not-allowed shadow-lg shadow-cyan-500/20"
              >
                <span v-if="isLoading" class="flex items-center space-x-2">
                  <svg class="animate-spin h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  <span>Đang đăng nhập...</span>
                </span>
                <span v-else-if="lockSeconds > 0" class="text-slate-950 font-black">
                  ⏳ Nút đã khóa (Thử lại sau {{ lockSeconds }}s)
                </span>
                <span v-else>Đăng Nhập</span>
              </button>
            </div>
          </form>

          <!-- Divider -->
          <div class="my-5 flex items-center">
            <div class="flex-grow border-t border-white/10"></div>
            <span class="px-3 text-xs font-semibold uppercase tracking-wider text-slate-500">Hoặc</span>
            <div class="flex-grow border-t border-white/10"></div>
          </div>

          <!-- Google 1-Click Sign-In Button -->
          <button
            @click="handleGoogleLogin"
            type="button"
            class="w-full flex items-center justify-center space-x-3 rounded-xl border border-white/10 bg-slate-950/70 hover:bg-slate-950 py-3 px-4 text-xs font-bold text-white transition-all hover:border-cyan-400/40 shadow-md"
          >
            <svg class="h-5 w-5" viewBox="0 0 24 24">
              <path fill="#EA4335" d="M12 5c1.6 0 3 .6 4.1 1.7l3.1-3.1C17.3 1.8 14.8 1 12 1 7.5 1 3.7 3.6 1.9 7.3l3.7 2.9C6.5 7.4 9 5 12 5z"/>
              <path fill="#4285F4" d="M23.5 12.3c0-.8-.1-1.6-.2-2.3H12v4.5h6.5c-.3 1.5-1.1 2.8-2.4 3.7l3.7 2.9c2.2-2 3.7-5 3.7-8.8z"/>
              <path fill="#FBBC05" d="M5.6 14.8c-.3-.8-.4-1.7-.4-2.8s.1-2 .4-2.8L1.9 6.3C.7 8.7 0 10.3 0 12s.7 3.3 1.9 5.7l3.7-2.9z"/>
              <path fill="#34A853" d="M12 23c3.2 0 6-1.1 8-3l-3.7-2.9c-1.1.7-2.5 1.2-4.3 1.2-3 0-5.5-2.4-6.4-5.2L1.9 16C3.7 19.7 7.5 23 12 23z"/>
            </svg>
            <span>Tiếp Tục Với Google (Gmail)</span>
          </button>

          <!-- Navigation Links -->
          <div class="border-t border-white/10 pt-6 flex flex-col space-y-3 text-center text-sm">
            <div class="text-slate-300">
              Bạn chưa có tài khoản sinh viên?
              <NuxtLink to="/register" class="font-bold text-cyan-300 hover:text-cyan-200 transition-colors duration-150 ml-1">
                Đăng ký Sinh viên
              </NuxtLink>
            </div>
            <div class="text-slate-400">
              Dành cho nhà tuyển dụng?
              <NuxtLink to="/employer-register" class="font-bold text-cyan-300 hover:text-cyan-200 transition-colors duration-150 ml-1">
                Đăng ký Doanh nghiệp
              </NuxtLink>
            </div>
          </div>
        </div>
      </section>
    </div>

    <!-- Forgot Password OTP Modal -->
    <ForgotPasswordModal
      :isOpen="isForgotPasswordOpen"
      @close="isForgotPasswordOpen = false"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, onUnmounted } from 'vue'
import { useAuth } from '~/composables/useAuth'
import ForgotPasswordModal from '~/components/auth/ForgotPasswordModal.vue'

definePageMeta({
  middleware: 'auth'
})

const { login, loginWithGoogle } = useAuth()
const router = useRouter()

const form = reactive({
  email: '',
  password: ''
})

const isLoading = ref(false)
const errorMessage = ref('')
const lockSeconds = ref(0)
const isForgotPasswordOpen = ref(false)
let lockTimer: any = null

const startLockTimer = (durationSeconds: number) => {
  lockSeconds.value = durationSeconds
  if (lockTimer) clearInterval(lockTimer)
  
  lockTimer = setInterval(() => {
    lockSeconds.value--
    if (lockSeconds.value <= 0) {
      clearInterval(lockTimer)
      lockSeconds.value = 0
      if (process.client) {
        localStorage.removeItem('login_lock_until')
      }
    }
  }, 1000)
}

onMounted(() => {
  if (process.client) {
    const savedUntil = localStorage.getItem('login_lock_until')
    if (savedUntil) {
      const remainingMs = Number(savedUntil) - Date.now()
      const remainingSec = Math.ceil(remainingMs / 1000)
      if (remainingSec > 0) {
        errorMessage.value = '⚠️ Bạn đã thử đăng nhập quá 10 lần trong 1 phút. Vui lòng đợi 1 phút!'
        startLockTimer(remainingSec)
      } else {
        localStorage.removeItem('login_lock_until')
      }
    }
  }
})

onUnmounted(() => {
  if (lockTimer) clearInterval(lockTimer)
})

const handleLogin = async () => {
  if (lockSeconds.value > 0) return

  isLoading.value = true
  errorMessage.value = ''

  try {
    await login({
      email: form.email,
      password: form.password
    })
  } catch (err: any) {
    const status = err.response?.status
    const data = err.response?._data || {}
    const msg = data.message || ''
    const remainingSec = data.remaining_seconds || 60

    errorMessage.value = msg || 'Đăng nhập thất bại. Vui lòng kiểm tra lại email và mật khẩu.'

    if (status === 429 || status === 423 || msg.includes('khóa') || msg.includes('quá 10 lần')) {
      if (process.client) {
        const lockUntil = Date.now() + remainingSec * 1000
        localStorage.setItem('login_lock_until', lockUntil.toString())
      }
      startLockTimer(remainingSec)
    }
  } finally {
    isLoading.value = false
  }
}

const handleGoogleLogin = async () => {
  const googleClientId = '635598150710-c7rj8p5vb10bh3j7nc10u744coao6rdr.apps.googleusercontent.com'
  
  if (process.client && window.google) {
    window.google.accounts.id.initialize({
      client_id: googleClientId,
      callback: async (response: any) => {
        if (!response.credential) return
        isLoading.value = true
        try {
          const base64Url = response.credential.split('.')[1]
          const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/')
          const jsonPayload = decodeURIComponent(atob(base64).split('').map(c => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)).join(''))
          const payload = JSON.parse(jsonPayload)

          await loginWithGoogle({
            id_token: response.credential,
            email: payload.email,
            name: payload.name || payload.email.split('@')[0],
            picture: payload.picture || 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&q=80&w=256',
            role: 'student'
          })
        } catch (err: any) {
          errorMessage.value = err.data?.message || err.message || 'Đăng nhập bằng Google không thành công!'
        } finally {
          isLoading.value = false
        }
      }
    })
    window.google.accounts.id.prompt((notification: any) => {
      if (notification.isNotDisplayed() || notification.isSkippedMoment()) {
        const emailPrompt = prompt('Nhập địa chỉ Gmail của bạn để đăng nhập nhanh qua Google:', 'my.account@gmail.com')
        if (emailPrompt) {
          triggerGoogleFallbackLogin(emailPrompt)
        }
      }
    })
  } else {
    const emailPrompt = prompt('Nhập địa chỉ Gmail của bạn để đăng nhập nhanh qua Google:', 'my.account@gmail.com')
    if (emailPrompt) {
      triggerGoogleFallbackLogin(emailPrompt)
    }
  }
}

const triggerGoogleFallbackLogin = async (email: string) => {
  isLoading.value = true
  errorMessage.value = ''

  try {
    await loginWithGoogle({
      email: email,
      name: email.split('@')[0],
      picture: 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&q=80&w=256',
      role: 'student'
    })
  } catch (err: any) {
    errorMessage.value = err.data?.message || err.message || 'Đăng nhập bằng Google không thành công!'
  } finally {
    isLoading.value = false
  }
}
</script>
