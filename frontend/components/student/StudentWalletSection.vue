<template>
  <div v-if="activeSection === 'wallet'" class="space-y-6">
    <div class="rounded-2xl border border-white/10 bg-slate-900/80 p-6 shadow-lg shadow-slate-950/25">
      <div class="flex items-center justify-between gap-4">
        <div>
          <p class="text-xs font-bold uppercase tracking-widest text-cyan-300">
            QuickWork Wallet
          </p>
          <h2 class="mt-2 text-3xl font-extrabold text-white">
            Ví thu nhập của bạn
          </h2>
          <p class="mt-1 text-sm font-medium text-slate-400">
            Theo dõi số dư và các khoản lương đã được doanh nghiệp giải ngân.
          </p>
        </div>

        <button
          @click="fetchWallet()"
          class="rounded-xl bg-cyan-400 px-4 py-2 text-sm font-bold text-slate-950 hover:bg-cyan-300 transition"
        >
          Làm mới
        </button>
      </div>

      <div class="mt-6 grid gap-4 md:grid-cols-3">
        <div class="rounded-2xl border border-cyan-400/20 bg-cyan-400/10 p-5">
          <p class="text-xs font-bold uppercase tracking-widest text-cyan-200">
            Số dư hiện tại
          </p>
          <p class="mt-3 text-3xl font-black text-white">
            {{ formatMoney(wallet?.balance || 0) }}
          </p>
        </div>

        <div class="rounded-2xl border border-emerald-400/20 bg-emerald-400/10 p-5">
          <p class="text-xs font-bold uppercase tracking-widest text-emerald-200">
            Tổng giao dịch
          </p>
          <p class="mt-3 text-3xl font-black text-white">
            {{ walletTransactions.length }}
          </p>
        </div>

        <div class="rounded-2xl border border-violet-400/20 bg-violet-400/10 p-5">
          <p class="text-xs font-bold uppercase tracking-widest text-violet-200">
            Trạng thái
          </p>
          <p class="mt-3 text-xl font-black text-white">
            Active
          </p>
        </div>
      </div>
    </div>

    <div class="rounded-2xl border border-white/10 bg-slate-900/80 shadow-lg shadow-slate-950/25 overflow-hidden">
      <div class="border-b border-white/10 px-6 py-4">
        <h3 class="text-lg font-bold text-white">Lịch sử giao dịch</h3>
        <p class="text-sm text-slate-400">Các khoản lương đã được ghi nhận.</p>
      </div>

      <div v-if="isLoadingWallet" class="p-6 text-sm font-semibold text-slate-400">
        Đang tải ví...
      </div>

      <div v-else-if="walletTransactions.length === 0" class="p-10 text-center">
        <p class="text-sm font-semibold text-slate-400">
          Chưa có giao dịch nào.
        </p>
      </div>

      <div v-else class="divide-y divide-white/10">
        <div
          v-for="tx in walletTransactions"
          :key="tx.id"
          class="flex items-center justify-between px-6 py-4 hover:bg-cyan-400/5 transition"
        >
          <div>
            <p class="text-sm font-bold text-white">
              {{ tx.description || 'Giao dịch ví' }}
            </p>
            <p class="mt-1 text-xs font-medium text-slate-500">
              {{ formatDate(tx.created_at) }}
            </p>
          </div>

          <div class="text-right">
            <p
              class="text-sm font-black"
              :class="tx.amount >= 0 ? 'text-emerald-400' : 'text-rose-400'"
            >
              {{ tx.amount >= 0 ? '+' : '-' }}{{ formatMoney(Math.abs(tx.amount || 0)) }}
            </p>
            <p class="mt-1 text-[11px] font-bold uppercase tracking-wider text-slate-500">
              {{ tx.type }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { toRefs } from 'vue'

const props = defineProps<{ state: Record<string, any> }>()

const {
  activeSection,
  wallet,
  walletTransactions,
  isLoadingWallet,
  fetchWallet
} = toRefs(props.state)

const formatMoney = (value: number) => {
  return Number(value || 0).toLocaleString('vi-VN') + 'đ'
}

const formatDate = (dateVal: any) => {
  if (!dateVal) return '-'
  return new Date(dateVal).toLocaleString('vi-VN')
}
</script>