<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center bg-slate-950/80 p-4 backdrop-blur-md">
    <div class="w-full max-w-md rounded-3xl border border-cyan-500/30 bg-slate-900 p-6 shadow-2xl shadow-cyan-950/80">
      
      <!-- Modal Header -->
      <div class="flex items-center justify-between border-b border-white/10 pb-4">
        <div class="flex items-center space-x-2">
          <div class="flex h-9 w-9 items-center justify-center rounded-xl bg-cyan-500/20 text-cyan-400">
            <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 7a2 2 0 012 2m-2 2a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2V9a2 2 0 00-2-2h-2zm0 0V5.5A2.5 2.5 0 0012.5 3h-5A2.5 2.5 0 005 5.5V7m10 4V9" />
            </svg>
          </div>
          <h3 class="text-lg font-extrabold text-white">Khôi Phục Mật Khẩu OTP</h3>
        </div>
        <button @click="closeModal" class="rounded-lg p-1 text-slate-400 hover:bg-white/10 hover:text-white transition-colors cursor-pointer">
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <!-- Messages Alert -->
      <div v-if="errorMessage" class="mt-4 rounded-xl border border-rose-500/30 bg-rose-500/10 p-3 text-xs font-medium text-rose-200">
        ⚠️ {{ errorMessage }}
      </div>
      <div v-if="successMessage" class="mt-4 rounded-xl border border-emerald-500/30 bg-emerald-500/10 p-3 text-xs font-medium text-emerald-200">
        ✅ {{ successMessage }}
      </div>

      <!-- STEP 1: Enter Email -->
      <div v-if="step === 1" class="mt-5 space-y-4">
        <p class="text-xs text-slate-300 leading-relaxed">
          Nhập địa chỉ Email đã đăng ký tài khoản QuickWork của bạn. Chúng tôi sẽ gửi mã xác thực <strong>OTP gồm 6 chữ số</strong> đến Email của bạn.
        </p>

        <div>
          <label class="block text-xs font-bold uppercase tracking-wider text-cyan-300 mb-1.5">Địa chỉ Email</label>
          <input
            type="email"
            v-model="email"
            placeholder="nhap.email@domain.com"
            @keyup.enter="handleSendOTP"
            class="w-full rounded-xl border border-white/10 bg-slate-950/80 px-4 py-2.5 text-xs text-white placeholder-slate-500 focus:border-cyan-400 focus:outline-none focus:ring-2 focus:ring-cyan-400/20"
          />
        </div>

        <button
          @click="handleSendOTP"
          :disabled="isLoading || !email"
          class="w-full rounded-xl bg-cyan-400 py-3 text-xs font-bold text-slate-950 hover:bg-cyan-300 disabled:opacity-50 transition-all shadow-lg shadow-cyan-500/20 cursor-pointer"
        >
          <span v-if="isLoading">📧 Đang gửi mã OTP...</span>
          <span v-else>Gửi Mã OTP Đến Gmail ➔</span>
        </button>
      </div>

      <!-- STEP 2: Enter OTP Code & New Password -->
      <div v-else-if="step === 2" class="mt-5 space-y-4">
        <p class="text-xs text-slate-300 leading-relaxed">
          Mã OTP đã được gửi đến <strong class="text-cyan-300">{{ email }}</strong>. Vui lòng nhập mã OTP 6 số và mật khẩu mới:
        </p>

        <div>
          <label class="block text-xs font-bold uppercase tracking-wider text-cyan-300 mb-1.5">Mã OTP (6 Chữ số)</label>
          <input
            type="text"
            v-model="otpCode"
            maxlength="6"
            placeholder="Mã 6 số (VD: 482915)"
            class="w-full rounded-xl border border-cyan-400/30 bg-slate-950/80 px-4 py-2.5 text-center text-lg font-black tracking-[8px] text-cyan-300 placeholder-slate-600 focus:border-cyan-400 focus:outline-none focus:ring-2 focus:ring-cyan-400/20"
          />
        </div>

        <div>
          <label class="block text-xs font-bold uppercase tracking-wider text-cyan-300 mb-1.5">Mật Khẩu Mới</label>
          <input
            type="password"
            v-model="newPassword"
            placeholder="••••••••"
            @keyup.enter="handleResetPassword"
            class="w-full rounded-xl border border-white/10 bg-slate-950/80 px-4 py-2.5 text-xs text-white placeholder-slate-500 focus:border-cyan-400 focus:outline-none focus:ring-2 focus:ring-cyan-400/20"
          />
        </div>

        <div class="flex items-center space-x-3 pt-2">
          <button
            @click="step = 1"
            class="w-1/3 rounded-xl border border-white/10 bg-slate-800 py-2.5 text-xs font-bold text-slate-300 hover:bg-slate-700 transition-colors cursor-pointer"
          >
            ← Quay lại
          </button>
          <button
            @click="handleResetPassword"
            :disabled="isLoading || !otpCode || !newPassword"
            class="w-2/3 rounded-xl bg-cyan-400 py-2.5 text-xs font-bold text-slate-950 hover:bg-cyan-300 disabled:opacity-50 transition-all shadow-lg shadow-cyan-500/20 cursor-pointer"
          >
            <span v-if="isLoading">⏳ Đang xử lý...</span>
            <span v-else>Xác Nhận Đổi Mật Khẩu 🔒</span>
          </button>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useApi } from '~/composables/useApi'

const props = defineProps({
  isOpen: Boolean
})

const emit = defineEmits(['close'])
const api = useApi()

const step = ref(1)
const email = ref('')
const otpCode = ref('')
const newPassword = ref('')
const isLoading = ref(false)
const errorMessage = ref('')
const successMessage = ref('')

const closeModal = () => {
  step.value = 1
  errorMessage.value = ''
  successMessage.value = ''
  emit('close')
}

const handleSendOTP = async () => {
  if (!email.value) return
  isLoading.value = true
  errorMessage.value = ''
  successMessage.value = ''

  try {
    const res: any = await api.post('/api/auth/forgot-password', {
      email: email.value
    }, { skipAutoLogout: true })
    successMessage.value = res?.message || 'Mã OTP đã được gửi thành công đến Gmail của bạn!'
    step.value = 2
  } catch (err: any) {
    console.error('Lỗi gửi OTP:', err)
    errorMessage.value = err?.data?.message || err?.message || 'Không thể gửi mã OTP. Vui lòng kiểm tra lại Email!'
  } finally {
    isLoading.value = false
  }
}

const handleResetPassword = async () => {
  if (!otpCode.value || !newPassword.value) return
  isLoading.value = true
  errorMessage.value = ''
  successMessage.value = ''

  try {
    const res: any = await api.post('/api/auth/reset-password', {
      email: email.value,
      otp_code: otpCode.value,
      new_password: newPassword.value
    }, { skipAutoLogout: true })
    successMessage.value = res?.message || 'Đặt lại mật khẩu thành công!'
    setTimeout(() => {
      closeModal()
    }, 2000)
  } catch (err: any) {
    console.error('Lỗi reset password:', err)
    errorMessage.value = err?.data?.message || err?.message || 'Mã OTP không đúng hoặc đã hết hạn!'
  } finally {
    isLoading.value = false
  }
}
</script>
