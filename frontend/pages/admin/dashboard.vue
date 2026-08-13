<template>
  <AdminShell>
    <div class="space-y-6">
      <section class="overflow-hidden rounded-3xl border border-white/10 bg-slate-950/80 shadow-2xl shadow-slate-950/40 ring-1 ring-cyan-400/10">
        <div class="grid gap-6 p-6 lg:grid-cols-[1.1fr_0.9fr] lg:p-8">
          <div class="flex flex-col justify-between gap-8">
            <div>
              <span class="inline-flex rounded-full bg-cyan-400/10 px-3 py-1 text-xs font-extrabold uppercase tracking-wide text-cyan-200 ring-1 ring-cyan-400/25">
                Bảng Quản trị Admin
              </span>
              <h1 class="mt-4 max-w-3xl text-3xl font-extrabold tracking-tight text-white sm:text-4xl">
                Trung tâm điều hành Quản trị QuickWork
              </h1>
              <p class="mt-3 max-w-2xl text-sm font-semibold leading-6 text-slate-300 sm:text-base">
                Theo dõi số liệu marketplace thực tế từ database và xử lý danh sách Doanh nghiệp đang chờ duyệt KYB.
              </p>
            </div>

            <div class="flex flex-col gap-3 sm:flex-row">
              <NuxtLink
                to="/admin/businesses/pending"
                class="inline-flex items-center justify-center rounded-xl bg-cyan-400 px-5 py-3 text-sm font-extrabold text-slate-950 shadow-lg shadow-cyan-500/20 transition-colors hover:bg-cyan-300"
              >
                Duyệt Doanh Nghiệp Pending
              </NuxtLink>
              <NuxtLink
                to="/admin/businesses"
                class="inline-flex items-center justify-center rounded-xl border border-cyan-400/30 bg-cyan-400/10 px-5 py-3 text-sm font-extrabold text-cyan-200 transition-colors hover:bg-cyan-400 hover:text-slate-950"
              >
                🏢 Quản lý Doanh nghiệp (A7)
              </NuxtLink>
              <NuxtLink
                to="/admin/students"
                class="inline-flex items-center justify-center rounded-xl border border-cyan-400/30 bg-cyan-400/10 px-5 py-3 text-sm font-extrabold text-cyan-200 transition-colors hover:bg-cyan-400 hover:text-slate-950"
              >
                🎓 Quản lý Sinh viên (A6)
              </NuxtLink>
              <button
                type="button"
                class="inline-flex items-center justify-center rounded-xl border border-white/10 bg-white/[0.06] px-5 py-3 text-sm font-bold text-slate-200 transition-colors hover:bg-white/10 disabled:cursor-not-allowed disabled:opacity-60"
                :disabled="isLoading"
                @click="fetchStats"
              >
                Cập nhật số liệu
              </button>
            </div>
          </div>

          <div class="rounded-2xl border border-cyan-400/20 bg-cyan-400/10 p-5">
            <p class="text-xs font-extrabold uppercase tracking-wide text-cyan-200">Hàng chờ duyệt KYB</p>
            <p class="mt-3 text-5xl font-extrabold text-white">
              {{ isLoading ? '...' : formatNumber(stats?.pending_businesses || 0) }}
            </p>
            <p class="mt-3 text-sm font-semibold leading-6 text-cyan-100">
              Doanh nghiệp đang chờ Admin duyệt trước khi được truy cập Business Dashboard.
            </p>
          </div>
        </div>
      </section>

      <section v-if="errorMessage" class="rounded-2xl border border-rose-400/20 bg-rose-400/10 p-5">
        <h2 class="text-base font-extrabold text-rose-100">Không tải được số liệu Dashboard</h2>
        <p class="mt-2 text-sm font-semibold text-rose-200">{{ errorMessage }}</p>
        <button
          type="button"
          class="mt-4 rounded-xl bg-rose-300 px-4 py-2 text-sm font-extrabold text-slate-950 transition-colors hover:bg-rose-200"
          @click="fetchStats"
        >
          Thử lại
        </button>
      </section>

      <section class="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
        <article
          v-for="card in statCards"
          :key="card.label"
          class="rounded-2xl border border-white/10 bg-slate-950/70 p-5 shadow-xl shadow-slate-950/25 ring-1 ring-cyan-400/10"
        >
          <div class="flex items-start justify-between gap-4">
            <div>
              <p class="text-xs font-extrabold uppercase tracking-wide text-slate-400">{{ card.label }}</p>
              <div v-if="isLoading" class="mt-3 h-9 w-28 animate-pulse rounded-lg bg-white/[0.08]" />
              <p v-else class="mt-3 text-3xl font-extrabold text-white">{{ card.value }}</p>
            </div>
            <span class="flex h-11 w-11 items-center justify-center rounded-xl bg-cyan-400/10 text-sm font-extrabold text-cyan-200 ring-1 ring-cyan-400/20">
              {{ card.short }}
            </span>
          </div>
          <p class="mt-4 text-sm font-semibold leading-6 text-slate-300">{{ card.caption }}</p>
        </article>
      </section>

      <!-- Data Visualization Section (Mục 3) -->
      <AdminDisbursementChart
        :total-disbursed="stats?.total_disbursed || 0"
        :student-count="stats?.total_students || 0"
        :business-count="stats?.total_businesses || 0"
      />

      <section class="rounded-2xl border border-white/10 bg-slate-950/70 p-5 shadow-xl shadow-slate-950/25 ring-1 ring-cyan-400/10">
        <div class="flex flex-col gap-4 md:flex-row md:items-center md:justify-between">
          <div>
            <p class="text-xs font-extrabold uppercase tracking-wide text-cyan-200">Quy trình chính</p>
            <h2 class="mt-2 text-xl font-extrabold text-white">Duyệt hồ sơ KYB Doanh nghiệp</h2>
            <p class="mt-2 text-sm font-semibold leading-6 text-slate-300">
              Mở danh sách pending để xem hồ sơ doanh nghiệp, duyệt hoặc từ chối kèm lý do. Hồ sơ đã xử lý sẽ biến mất khỏi queue.
            </p>
          </div>
          <NuxtLink
            to="/admin/businesses/pending"
            class="inline-flex items-center justify-center rounded-xl border border-cyan-400/25 bg-cyan-400/10 px-5 py-3 text-sm font-extrabold text-cyan-100 transition-colors hover:bg-cyan-400 hover:text-slate-950"
          >
            Đến Queue Chờ Duyệt
          </NuxtLink>
        </div>
      </section>
    </div>
  </AdminShell>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import AdminShell from '~/components/admin/AdminShell.vue'
import AdminDisbursementChart from '~/components/admin/AdminDisbursementChart.vue'
import { useAdminApi } from '~/composables/useAdminApi'
import type { AdminDashboardStats } from '~/types/admin'

definePageMeta({
  middleware: 'auth'
})

const adminApi = useAdminApi()
const stats = ref<AdminDashboardStats | null>(null)
const isLoading = ref(false)
const errorMessage = ref('')

const formatNumber = (value: number) => new Intl.NumberFormat('vi-VN').format(value)
const formatCurrency = (value: number) =>
  new Intl.NumberFormat('vi-VN', { style: 'currency', currency: 'VND', maximumFractionDigits: 0 }).format(value)
const extractErrorMessage = (error: unknown, fallback: string) => {
  const maybeError = error as { response?: { _data?: { message?: string } } }
  return maybeError.response?._data?.message || fallback
}

const statCards = computed(() => [
  {
    label: 'Tổng sinh viên',
    short: 'SV',
    value: formatNumber(stats.value?.total_students || 0),
    caption: 'Số tài khoản có role student.'
  },
  {
    label: 'Tổng doanh nghiệp',
    short: 'DN',
    value: formatNumber(stats.value?.total_businesses || 0),
    caption: 'Số tài khoản có role business.'
  },
  {
    label: 'Doanh nghiệp chờ duyệt',
    short: 'KYB',
    value: formatNumber(stats.value?.pending_businesses || 0),
    caption: 'Business có trạng thái pending cần Admin xử lý.'
  },
  {
    label: 'Tổng Job',
    short: 'JOB',
    value: formatNumber(stats.value?.total_jobs || 0),
    caption: 'Tổng tin tuyển dụng trong hệ thống.'
  },
  {
    label: 'Job pending',
    short: 'PEN',
    value: formatNumber(stats.value?.pending_jobs || 0),
    caption: 'Tin tuyển dụng đang chờ duyệt.'
  },
  {
    label: 'Tổng giải ngân',
    short: 'VND',
    value: formatCurrency(stats.value?.total_disbursed || 0),
    caption: 'Tổng lương đã giải ngân qua hệ thống, không phải doanh thu ròng QuickWork.'
  }
])

const fetchStats = async () => {
  isLoading.value = true
  errorMessage.value = ''

  try {
    stats.value = await adminApi.getDashboardStats()
  } catch (error: unknown) {
    errorMessage.value = extractErrorMessage(error, 'Không thể tải số liệu dashboard.')
  } finally {
    isLoading.value = false
  }
}

onMounted(fetchStats)
</script>
