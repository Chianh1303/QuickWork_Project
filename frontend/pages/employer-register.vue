<template>
  <div class="relative min-h-[calc(100vh-4rem)] overflow-hidden bg-slate-950 px-4 py-10 sm:px-6 lg:px-8">
    <div class="absolute inset-0 bg-[radial-gradient(circle_at_18%_18%,rgba(34,211,238,0.18),transparent_28rem),radial-gradient(circle_at_82%_78%,rgba(59,130,246,0.12),transparent_30rem),linear-gradient(180deg,#020617_0%,#0f172a_50%,#020617_100%)]"></div>

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
              Giao diện Nhà tuyển dụng
            </p>
            <h1 class="text-5xl font-extrabold leading-none tracking-tight text-white">
              Đăng tin tuyển dụng & kết nối ứng viên tài năng.
            </h1>
            <p class="text-base font-medium leading-8 text-slate-300">
              Tạo tài khoản Doanh nghiệp để đăng tin tuyển dụng, xem hồ sơ sinh viên, nhắn tin trực tiếp và quản lý tiến trình tuyển dụng tập trung.
            </p>
          </div>

          <div class="grid grid-cols-3 gap-3">
            <div class="rounded-xl border border-white/10 bg-white/[0.06] p-4">
              <p class="text-2xl font-extrabold text-white">Nhanh chóng</p>
              <p class="mt-1 text-xs font-semibold text-slate-400">Đăng tin bài</p>
            </div>
            <div class="rounded-xl border border-white/10 bg-white/[0.06] p-4">
              <p class="text-2xl font-extrabold text-white">Trực tiếp</p>
              <p class="mt-1 text-xs font-semibold text-slate-400">Nhắn tin HR</p>
            </div>
            <div class="rounded-xl border border-white/10 bg-white/[0.06] p-4">
              <p class="text-2xl font-extrabold text-white">Quản lý</p>
              <p class="mt-1 text-xs font-semibold text-slate-400">Đơn ứng tuyển</p>
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
              Đăng Ký Doanh Nghiệp
            </h2>
            <p class="mt-2 text-center text-sm text-slate-300">
              Tuyển dụng nhân sự sinh viên và quản lý các công việc đã đăng.
            </p>
          </div>

          <!-- Role Switcher Pills -->
          <div class="mt-6 flex rounded-2xl bg-slate-950/80 p-1 border border-white/10">
            <NuxtLink
              to="/register"
              class="flex-1 py-2 text-xs rounded-xl text-slate-400 hover:text-white font-bold text-center flex items-center justify-center gap-1.5 transition-all"
            >
              <span>🎓</span>
              <span>Sinh Viên</span>
            </NuxtLink>
            <span class="flex-1 py-2 text-xs rounded-xl bg-cyan-400 text-slate-950 font-black shadow-md shadow-cyan-500/20 text-center flex items-center justify-center gap-1.5 cursor-default">
              <span>🏢</span>
              <span>Doanh Nghiệp</span>
            </span>
          </div>

          <!-- Success / Error Alert -->
          <div v-if="successMessage" class="mt-6 bg-emerald-400/10 border border-emerald-400/30 text-emerald-100 px-4 py-3 rounded-lg text-sm flex items-center space-x-3" role="alert">
            <svg class="h-5 w-5 text-emerald-300 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <span class="font-medium">{{ successMessage }}</span>
          </div>

          <div v-if="errorMessage" class="mt-6 bg-rose-400/10 border border-rose-400/30 text-rose-100 px-4 py-3 rounded-lg text-sm flex items-center space-x-3" role="alert">
            <svg class="h-5 w-5 text-rose-300 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <span class="font-medium">{{ errorMessage }}</span>
          </div>

          <form class="mt-8 space-y-5" @submit.prevent="handleRegister" v-if="!registrationSuccess">
            <div class="space-y-4">
              <!-- Full Name / HR Name -->
              <div>
                <label for="full_name" class="block text-sm font-semibold text-slate-200 mb-1">Họ tên Người đại diện / HR</label>
                <input
                  id="full_name"
                  type="text"
                  required
                  v-model="form.full_name"
                  class="appearance-none block w-full px-3 py-3 border border-white/10 placeholder-slate-500 text-white rounded-lg bg-slate-950/70 hover:bg-slate-950 focus:outline-none focus:ring-2 focus:ring-cyan-400 focus:border-transparent transition-all duration-200 sm:text-sm"
                  placeholder="Ví dụ: Nguyễn Thị B"
                />
              </div>

              <!-- Corporate Email -->
              <div>
                <label for="email" class="block text-sm font-semibold text-slate-200 mb-1">Email Doanh nghiệp</label>
                <input
                  id="email"
                  type="email"
                  required
                  v-model="form.email"
                  class="appearance-none block w-full px-3 py-3 border border-white/10 placeholder-slate-500 text-white rounded-lg bg-slate-950/70 hover:bg-slate-950 focus:outline-none focus:ring-2 focus:ring-cyan-400 focus:border-transparent transition-all duration-200 sm:text-sm"
                  placeholder="Ví dụ: hr@fpt.com"
                />
              </div>

              <!-- Password -->
              <div>
                <label for="password" class="block text-sm font-semibold text-slate-200 mb-1">Mật khẩu</label>
                <input
                  id="password"
                  type="password"
                  required
                  v-model="form.password"
                  class="appearance-none block w-full px-3 py-3 border border-white/10 placeholder-slate-500 text-white rounded-lg bg-slate-950/70 hover:bg-slate-950 focus:outline-none focus:ring-2 focus:ring-cyan-400 focus:border-transparent transition-all duration-200 sm:text-sm"
                  placeholder="Tối thiểu 6 ký tự"
                />
              </div>

              <!-- Company Name -->
              <div>
                <label for="company_name" class="block text-sm font-semibold text-slate-200 mb-1">Tên Công ty / Doanh nghiệp</label>
                <input
                  id="company_name"
                  type="text"
                  required
                  v-model="form.company_name"
                  class="appearance-none block w-full px-3 py-3 border border-white/10 placeholder-slate-500 text-white rounded-lg bg-slate-950/70 hover:bg-slate-950 focus:outline-none focus:ring-2 focus:ring-cyan-400 focus:border-transparent transition-all duration-200 sm:text-sm"
                  placeholder="Ví dụ: Công ty TNHH QuickWork"
                />
              </div>

              <!-- Tax Code -->
              <div>
                <label for="tax_code" class="block text-sm font-semibold text-slate-200 mb-1">Mã số thuế</label>
                <input
                  id="tax_code"
                  type="text"
                  required
                  v-model="form.tax_code"
                  class="appearance-none block w-full px-3 py-3 border border-white/10 placeholder-slate-500 text-white rounded-lg bg-slate-950/70 hover:bg-slate-950 focus:outline-none focus:ring-2 focus:ring-cyan-400 focus:border-transparent transition-all duration-200 sm:text-sm"
                  placeholder="Ví dụ: 0312345678"
                />
              </div>
            </div>

            <div class="pt-2">
              <button
                type="submit"
                :disabled="isLoading"
                class="group relative w-full flex justify-center py-3 px-4 border border-transparent text-sm font-bold rounded-lg text-slate-950 bg-cyan-400 hover:bg-cyan-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-400 transition-all duration-150 disabled:opacity-50 disabled:cursor-not-allowed shadow-lg shadow-cyan-500/20"
              >
                <span v-if="isLoading" class="flex items-center space-x-2">
                  <svg class="animate-spin h-5 w-5 text-slate-950" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  <span>Đang tạo tài khoản...</span>
                </span>
                <span v-else>Đăng Ký Doanh Nghiệp</span>
              </button>
            </div>
          </form>

          <!-- Success Screen -->
          <div v-else class="text-center py-8 space-y-4">
            <div class="inline-flex items-center justify-center p-3 bg-emerald-400/10 border border-emerald-400/30 rounded-full text-emerald-300 mb-2">
              <svg class="h-12 w-12" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <h3 class="text-xl font-bold text-white">Đăng Ký Doanh Nghiệp Thành Công!</h3>
            <p class="text-slate-300 text-sm max-w-sm mx-auto font-medium">
              Hồ sơ của bạn đã được lưu (chờ xét duyệt KYB từ Admin). Bạn có thể đăng nhập ngay bây giờ.
            </p>
            <div class="pt-4">
              <NuxtLink
                to="/login"
                class="inline-flex justify-center items-center py-2.5 px-6 font-bold text-sm text-slate-950 bg-cyan-400 rounded-lg hover:bg-cyan-300 shadow-md shadow-cyan-500/10 focus-ring"
              >
                Đến Trang Đăng Nhập
              </NuxtLink>
            </div>
          </div>

          <!-- Navigation Links -->
          <div class="border-t border-white/10 pt-6 flex flex-col space-y-3 text-center text-sm">
            <div class="text-slate-300">
              Bạn đã có tài khoản?
              <NuxtLink to="/login" class="font-bold text-cyan-300 hover:text-cyan-200 transition-colors duration-150 ml-1">
                Đăng nhập ngay
              </NuxtLink>
            </div>
            <div class="text-slate-400">
              Bạn là sinh viên tìm việc?
              <NuxtLink to="/register" class="font-bold text-cyan-300 hover:text-cyan-200 transition-colors duration-150 ml-1">
                Đăng ký Sinh viên
              </NuxtLink>
            </div>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useAuth } from '~/composables/useAuth'

definePageMeta({
  middleware: 'auth'
})

const { register } = useAuth()

const form = reactive({
  full_name: '',
  email: '',
  password: '',
  company_name: '',
  tax_code: ''
})

const isLoading = ref(false)
const errorMessage = ref('')
const successMessage = ref('')
const registrationSuccess = ref(false)

const handleRegister = async () => {
  isLoading.value = true
  errorMessage.value = ''
  successMessage.value = ''

  try {
    const res = await register({
      email: form.email,
      password: form.password,
      role: 'business',
      full_name: form.full_name,
      company_name: form.company_name,
      tax_code: form.tax_code
    })
    successMessage.value = res?.message || '🎉 Đăng ký doanh nghiệp thành công!'
    registrationSuccess.value = true
  } catch (err: any) {
    errorMessage.value = err?.data?.message || err?.response?._data?.message || err?.message || 'Đăng ký thất bại. Vui lòng kiểm tra lại thông tin công ty.'
  } finally {
    isLoading.value = false
  }
}
</script>

