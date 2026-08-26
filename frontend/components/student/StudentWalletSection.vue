<template>
  <div v-show="activeSection === 'wallet' || activeSection?.value === 'wallet'" class="max-w-6xl mx-auto space-y-6">
    <!-- Virtual Escrow Card & Balance Banner -->
    <div class="relative overflow-hidden rounded-3xl border border-cyan-500/30 bg-gradient-to-br from-slate-950 via-cyan-950/80 to-slate-950 p-6 sm:p-8 shadow-2xl shadow-cyan-950/50 backdrop-blur-xl ring-1 ring-cyan-500/20">
      
      <!-- Ambient Glow Effects -->
      <div class="absolute -top-24 -right-24 h-64 w-64 rounded-full bg-cyan-500/20 blur-3xl pointer-events-none"></div>
      <div class="absolute -bottom-24 -left-24 h-64 w-64 rounded-full bg-emerald-500/15 blur-3xl pointer-events-none"></div>

      <div class="relative z-10 flex flex-col lg:flex-row lg:items-center lg:justify-between gap-6">
        <!-- Left Balance Info -->
        <div class="space-y-4">
          <div class="flex items-center gap-2">
            <span class="inline-flex items-center gap-1.5 rounded-full bg-emerald-400/10 px-3 py-1 text-xs font-black text-emerald-300 ring-1 ring-emerald-400/30">
              <span class="h-2 w-2 rounded-full bg-emerald-400 animate-pulse"></span>
              Smart Escrow Payment Guaranteed
            </span>
          </div>

          <div>
            <p class="text-xs font-black uppercase tracking-widest text-cyan-300">Tổng Số Dư Khả Dụng</p>
            <h2 class="mt-1 text-4xl sm:text-5xl font-black tracking-tight bg-gradient-to-r from-white via-slate-100 to-cyan-200 bg-clip-text text-transparent">
              {{ formatMoney(wallet?.balance || 0) }}
            </h2>
          </div>

          <p class="text-xs font-semibold text-slate-300 max-w-md leading-relaxed">
            Tiền lương ca làm việc được tạm giữ an toàn qua hợp đồng thông minh Escrow và tự động chuyển về ví khi Doanh nghiệp hoàn tất duyệt ca.
          </p>
        </div>

        <!-- Right Digital VIP Card Card Visual -->
        <div class="w-full lg:w-80 rounded-2xl border border-cyan-400/30 bg-gradient-to-tr from-cyan-900/90 via-slate-900/90 to-cyan-950/90 p-5 shadow-2xl shadow-cyan-950/60 space-y-4 flex-shrink-0 backdrop-blur-md">
          <div class="flex justify-between items-start">
            <div>
              <span class="text-[10px] font-black uppercase tracking-widest text-cyan-300">QuickWork Pay</span>
              <p class="text-xs font-black text-white">STUDENT VIP CARD</p>
            </div>
            <!-- Hologram Chip Mockup -->
            <div class="h-7 w-10 rounded-lg bg-gradient-to-tr from-amber-300 via-amber-400 to-amber-200 shadow-inner flex items-center justify-center">
              <div class="h-4 w-6 border border-amber-600/40 rounded-sm"></div>
            </div>
          </div>

          <div class="pt-2">
            <p class="text-[10px] font-mono tracking-widest text-slate-400">ACCOUNT ID</p>
            <p class="text-xs font-mono font-bold text-cyan-200 tracking-wider">**** **** **** {{ String(wallet?.id || '8888').slice(-4) }}</p>
          </div>

          <div class="flex justify-between items-center pt-2 border-t border-cyan-500/20 text-xs">
            <button
              @click="fetchWallet()"
              class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-xl bg-cyan-500/20 text-cyan-200 hover:bg-cyan-500/30 font-bold transition-all text-[11px]"
            >
              <span>🔄 Làm mới ví</span>
            </button>
            <button
              @click="openWithdrawModal"
              class="inline-flex items-center gap-1 px-3 py-1.5 rounded-xl bg-gradient-to-r from-emerald-500 to-teal-400 text-slate-950 hover:from-emerald-400 hover:to-teal-300 font-extrabold transition-all text-[11px] shadow-md shadow-emerald-500/20"
            >
              <span>⚡ Rút tiền ngân hàng</span>
            </button>
          </div>
        </div>
      </div>

      <!-- Quick Metrics Bar -->
      <div class="mt-8 grid gap-4 grid-cols-1 sm:grid-cols-3 pt-6 border-t border-cyan-500/15">
        <div class="rounded-2xl border border-cyan-500/20 bg-slate-950/60 p-4 flex items-center justify-between">
          <div>
            <p class="text-[11px] font-black uppercase tracking-wider text-cyan-300">Số Dư Hiện Tại</p>
            <p class="mt-1 text-xl font-black text-white">{{ formatMoney(wallet?.balance || 0) }}</p>
          </div>
          <button
            @click="openWithdrawModal"
            class="px-3 py-1.5 rounded-xl text-xs font-black bg-emerald-500/10 text-emerald-300 border border-emerald-500/20 hover:bg-emerald-500/20 transition-all"
          >
            Rút tiền
          </button>
        </div>

        <div class="rounded-2xl border border-emerald-500/20 bg-slate-950/60 p-4">
          <p class="text-[11px] font-black uppercase tracking-wider text-emerald-300">Tổng Số Giao Dịch</p>
          <p class="mt-1 text-xl font-black text-white">{{ (walletTransactions || []).length }} lượt</p>
        </div>

        <div class="rounded-2xl border border-cyan-500/20 bg-slate-950/60 p-4">
          <p class="text-[11px] font-black uppercase tracking-wider text-cyan-300">Chuẩn Hóa Escrow</p>
          <p class="mt-1 text-xl font-black text-emerald-300">Bảo Đảm 100%</p>
        </div>
      </div>
    </div>

    <!-- Transaction History Table -->
    <div class="rounded-3xl border border-cyan-500/20 bg-slate-900/90 shadow-xl overflow-hidden backdrop-blur-xl">
      <div class="border-b border-cyan-500/15 px-6 py-5 flex items-center justify-between">
        <div>
          <h3 class="text-lg font-extrabold text-white">Lịch Sử Giao Dịch Thu Nhập</h3>
          <p class="text-xs font-semibold text-slate-400 mt-0.5">Biến động số dư tài khoản từ việc ứng tuyển & hoàn thành ca làm</p>
        </div>
        <span class="text-xs font-bold text-cyan-300 bg-cyan-500/10 px-3 py-1 rounded-full border border-cyan-500/20">
          Cập nhật tự động
        </span>
      </div>

      <div v-if="isLoadingWallet" class="p-8 text-center text-xs font-extrabold text-slate-400 animate-pulse">
        ⏳ Đang tải dữ liệu biến động ví...
      </div>

      <div v-else-if="!walletTransactions || walletTransactions.length === 0" class="p-12 text-center space-y-2">
        <p class="text-sm font-extrabold text-white">Chưa có lịch sử giao dịch nào</p>
        <p class="text-xs text-slate-400 max-w-sm mx-auto">
          Hoàn thành ca làm việc đầu tiên để nhận giải ngân lương trực tiếp vào ví!
        </p>
      </div>

      <div v-else class="divide-y divide-cyan-500/10">
        <div
          v-for="tx in (walletTransactions || [])"
          :key="tx.id"
          class="flex items-center justify-between px-6 py-4 hover:bg-cyan-500/5 transition-colors"
        >
          <div class="flex items-center space-x-3.5">
            <div
              :class="tx.amount >= 0 ? 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20' : 'bg-rose-500/10 text-rose-400 border-rose-500/20'"
              class="h-10 w-10 rounded-2xl border flex items-center justify-center font-black text-base flex-shrink-0"
            >
              {{ tx.amount >= 0 ? '↓' : '↑' }}
            </div>
            <div>
              <p class="text-xs font-extrabold text-white">
                {{ tx.description || 'Giải ngân lương ca làm việc' }}
              </p>
              <p class="mt-0.5 text-[11px] font-semibold text-slate-400">
                Mã GD: #TX-{{ String(tx.id || '000').slice(-6) }} • {{ formatDate(tx.created_at) }}
              </p>
            </div>
          </div>

          <div class="text-right">
            <p
              class="text-sm font-black tracking-tight"
              :class="tx.amount >= 0 ? 'text-emerald-300' : 'text-rose-300'"
            >
              {{ tx.amount >= 0 ? '+' : '-' }}{{ formatMoney(Math.abs(tx.amount || 0)) }}
            </p>
            <p class="mt-0.5 text-[10px] font-black uppercase tracking-wider text-cyan-300">
              {{ tx.type || 'LƯƠNG CA LÀM' }}
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal Yêu cầu Rút tiền về Ngân hàng (Bank Withdrawal Modal) -->
    <div v-if="isWithdrawModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
      <div class="bg-slate-900/95 rounded-3xl border border-cyan-500/20 shadow-2xl shadow-slate-950 max-w-md w-full overflow-hidden animate-in fade-in zoom-in-95 duration-200">
        <div class="p-6 pb-4 border-b border-cyan-500/15 flex items-center justify-between bg-slate-950/60">
          <div>
            <span class="text-[10px] font-black uppercase tracking-wider text-emerald-400 bg-emerald-500/10 px-2.5 py-0.5 rounded-full border border-emerald-500/20">
              Chuyển khoản 24/7
            </span>
            <h3 class="mt-1 text-lg font-extrabold text-white">Rút Tiền Về Tài Khoản Ngân Hàng</h3>
          </div>
          <button @click="isWithdrawModalOpen = false" class="text-slate-400 hover:text-white p-1 rounded-lg">✕</button>
        </div>

        <form @submit.prevent="submitWithdrawRequest" class="p-6 space-y-4">
          <div>
            <label class="block text-xs font-extrabold text-cyan-300 uppercase tracking-wider mb-1.5">Ngân hàng thụ hưởng</label>
            <select v-model="withdrawForm.bank" required class="block w-full px-3.5 py-2.5 border border-cyan-500/20 rounded-xl text-xs bg-slate-950 text-slate-100 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium">
              <option value="">-- Chọn ngân hàng --</option>
              <option value="Vietcombank">Vietcombank (VCB)</option>
              <option value="MBBank">MB Bank (Quân Đội)</option>
              <option value="Techcombank">Techcombank (TCB)</option>
              <option value="VPBank">VPBank</option>
              <option value="ACB">ACB (Á Châu)</option>
              <option value="BIDV">BIDV</option>
              <option value="VietinBank">VietinBank</option>
              <option value="TPBank">TPBank</option>
              <option value="Agribank">Agribank</option>
            </select>
          </div>

          <div>
            <label class="block text-xs font-extrabold text-cyan-300 uppercase tracking-wider mb-1.5">Số tài khoản nhận tiền</label>
            <input
              v-model="withdrawForm.accountNumber"
              type="text"
              required
              class="block w-full px-3.5 py-2.5 border border-cyan-500/20 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-mono font-bold"
              placeholder="VD: 19038888888888"
            />
          </div>

          <div>
            <label class="block text-xs font-extrabold text-cyan-300 uppercase tracking-wider mb-1.5">Tên chủ tài khoản (In hoa không dấu)</label>
            <input
              v-model="withdrawForm.accountHolder"
              type="text"
              required
              class="block w-full px-3.5 py-2.5 border border-cyan-500/20 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-bold uppercase"
              placeholder="VD: NGUYEN VAN A"
            />
          </div>

          <div>
            <div class="flex justify-between items-center mb-1.5">
              <label class="text-xs font-extrabold text-cyan-300 uppercase tracking-wider">Số tiền rút (VNĐ)</label>
              <span class="text-[11px] font-bold text-slate-400">Khả dụng: {{ formatMoney(wallet?.balance || 0) }}</span>
            </div>
            <div class="relative">
              <input
                v-model="withdrawForm.amount"
                type="number"
                :max="wallet?.balance || 0"
                min="50000"
                step="10000"
                required
                class="block w-full px-3.5 py-2.5 border border-cyan-500/20 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-bold"
                placeholder="Tối thiểu 50.000 VNĐ"
              />
              <button
                type="button"
                @click="withdrawForm.amount = String(wallet?.balance || 0)"
                class="absolute right-2.5 top-2 px-2 py-0.5 rounded-lg bg-cyan-500/20 text-cyan-300 text-[10px] font-black hover:bg-cyan-500/30 transition-all"
              >
                Tất cả
              </button>
            </div>
          </div>

          <div class="p-3 rounded-xl bg-slate-950/60 border border-cyan-500/10 text-[11px] text-slate-400 leading-relaxed">
            💡 Tiền sẽ được chuyển khoản tự động về tài khoản ngân hàng của bạn trong vòng 1-24 giờ làm việc. Miễn phí rút tiền.
          </div>

          <div class="flex items-center gap-3 pt-2">
            <button
              type="button"
              @click="isWithdrawModalOpen = false"
              class="flex-1 px-4 py-2.5 border border-white/10 text-xs font-bold rounded-xl text-slate-300 bg-white/5 hover:bg-white/10 hover:text-white transition-all"
            >
              Hủy bỏ
            </button>
            <button
              type="submit"
              :disabled="isSubmittingWithdraw || Number(wallet?.balance || 0) < 50000"
              class="flex-1 px-4 py-2.5 rounded-xl text-xs font-extrabold text-slate-950 bg-gradient-to-r from-emerald-400 to-teal-300 hover:from-emerald-300 hover:to-teal-200 shadow-md shadow-emerald-500/25 transition-all disabled:opacity-50"
            >
              {{ isSubmittingWithdraw ? 'Đang tạo lệnh...' : 'Xác nhận rút tiền' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, toRefs } from 'vue'
import { useToast } from '~/composables/useToast'

const props = defineProps<{ state: Record<string, any> }>()

const {
  activeSection,
  wallet,
  walletTransactions,
  isLoadingWallet,
  fetchWallet
} = toRefs(props.state)

const { success, error } = useToast()

const isWithdrawModalOpen = ref(false)
const isSubmittingWithdraw = ref(false)
const withdrawForm = ref({
  bank: '',
  accountNumber: '',
  accountHolder: '',
  amount: ''
})

const openWithdrawModal = () => {
  withdrawForm.value = {
    bank: '',
    accountNumber: '',
    accountHolder: '',
    amount: String(wallet?.value?.balance || '')
  }
  isWithdrawModalOpen.value = true
}

const submitWithdrawRequest = async () => {
  const amt = Number(withdrawForm.value.amount || 0)
  const currentBal = Number(wallet?.value?.balance || 0)

  if (amt <= 0 || amt > currentBal) {
    error('Số tiền rút không hợp lệ hoặc vượt quá số dư khả dụng!')
    return
  }

  if (amt < 50000) {
    error('Số tiền rút tối thiểu là 50.000 VNĐ!')
    return
  }

  isSubmittingWithdraw.value = true
  // Mock request processing
  setTimeout(() => {
    isSubmittingWithdraw.value = false
    isWithdrawModalOpen.value = false
    success(`🎉 Đã gửi lệnh rút ${amt.toLocaleString('vi-VN')} VNĐ về ngân hàng ${withdrawForm.value.bank}!`)
  }, 1000)
}

const formatMoney = (value: number) => {
  return Number(value || 0).toLocaleString('vi-VN') + 'đ'
}

const formatDate = (dateVal: any) => {
  if (!dateVal) return '-'
  return new Date(dateVal).toLocaleString('vi-VN')
}
</script>