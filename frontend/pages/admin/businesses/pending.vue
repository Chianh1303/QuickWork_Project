<template>
  <AdminShell>
    <div class="space-y-6">
      <section class="rounded-3xl border border-white/10 bg-slate-950/80 p-6 shadow-2xl shadow-slate-950/40 ring-1 ring-cyan-400/10 lg:p-7">
        <div class="flex flex-col gap-6">
          <div class="flex flex-col gap-5 lg:flex-row lg:items-center lg:justify-between">
            <div>
              <h1 class="text-3xl font-extrabold tracking-tight text-white sm:text-4xl">
                Doanh nghiệp Chờ Xét Duyệt
              </h1>
              <p class="mt-2 max-w-2xl text-sm font-semibold leading-6 text-slate-300 sm:text-base">
                Tìm kiếm, xem hồ sơ KYB, duyệt hoặc từ chối doanh nghiệp đang chờ xác minh.
              </p>
            </div>

            <div class="grid min-w-0 grid-cols-2 gap-3 sm:min-w-[360px]">
              <div class="rounded-2xl border border-cyan-400/20 bg-cyan-400/10 px-5 py-4">
                <p class="text-xs font-extrabold uppercase tracking-wide text-cyan-200">Tổng số kết quả</p>
                <p class="mt-1 text-2xl font-extrabold text-white">{{ pagination.total }}</p>
              </div>
              <div class="rounded-2xl border border-white/10 bg-white/[0.06] px-5 py-4">
                <p class="text-xs font-extrabold uppercase tracking-wide text-slate-400">Trang hiện tại</p>
                <p class="mt-1 text-2xl font-extrabold text-white">{{ pagination.page }}</p>
              </div>
            </div>
          </div>
        </div>
      </section>

      <section class="rounded-2xl border border-white/10 bg-slate-950/70 p-4 shadow-xl shadow-slate-950/30 ring-1 ring-cyan-400/10 sm:p-5">
        <div class="grid gap-3 xl:grid-cols-[1fr_auto_auto] xl:items-end">
          <label class="block">
            <span class="mb-2 block text-xs font-extrabold uppercase tracking-wide text-slate-400">
              Tìm kiếm theo Tên công ty, Mã số thuế hoặc Email
            </span>
            <input
              v-model="searchInput"
              type="search"
              class="w-full rounded-xl border border-white/10 bg-slate-950/80 px-4 py-3 text-sm font-semibold text-white outline-none transition focus:border-cyan-300 focus:ring-2 focus:ring-cyan-400/30"
              placeholder="Nhập tên công ty, mã số thuế hoặc email..."
              @keydown.enter.prevent="applySearchNow"
            >
          </label>

          <button
            type="button"
            class="inline-flex h-12 items-center justify-center rounded-xl bg-cyan-400 px-7 text-sm font-extrabold text-slate-950 transition-colors hover:bg-cyan-300 disabled:cursor-not-allowed disabled:opacity-60"
            :disabled="isLoading"
            @click="applySearchNow"
          >
            Tìm kiếm
          </button>

          <button
            type="button"
            class="inline-flex h-12 items-center justify-center rounded-xl border border-white/10 bg-white/5 px-5 text-sm font-extrabold text-slate-200 transition-colors hover:bg-white/10 hover:text-white disabled:cursor-not-allowed disabled:opacity-60"
            :disabled="isLoading || (!searchInput && !activeSearch)"
            @click="clearSearch"
          >
            Clear
          </button>
        </div>
      </section>

      <section class="overflow-hidden rounded-2xl border border-white/10 bg-slate-950/70 shadow-xl shadow-slate-950/30 ring-1 ring-cyan-400/10">
        <div class="flex flex-col gap-3 border-b border-white/10 px-5 py-4 sm:flex-row sm:items-center sm:justify-between">
          <div>
            <h2 class="text-base font-extrabold text-white">Business KYB registrations</h2>
            <p class="mt-1 text-sm font-semibold text-slate-400">
              {{ activeSearch ? `Filtered by "${activeSearch}"` : 'Showing all pending businesses' }}
            </p>
          </div>
          <span class="inline-flex w-fit rounded-full bg-white/[0.06] px-3 py-1 text-xs font-extrabold uppercase tracking-wide text-slate-300 ring-1 ring-white/10">
            {{ pagination.limit }} per page
          </span>
        </div>

        <div v-if="errorMessage" class="p-6">
          <div class="rounded-2xl border border-rose-400/20 bg-rose-400/10 p-5">
            <h2 class="text-base font-extrabold text-rose-100">Unable to load pending businesses</h2>
            <p class="mt-2 text-sm font-semibold text-rose-200">{{ errorMessage }}</p>
            <button
              type="button"
              class="mt-4 rounded-xl bg-rose-300 px-4 py-2 text-sm font-extrabold text-slate-950 transition-colors hover:bg-rose-200"
              @click="fetchPendingBusinesses"
            >
              Retry
            </button>
          </div>
        </div>

        <div v-else-if="isLoading" class="space-y-3 p-5">
          <div
            v-for="index in 5"
            :key="index"
            class="h-16 animate-pulse rounded-xl border border-white/10 bg-white/[0.06]"
          />
        </div>

        <div v-else-if="items.length === 0" class="p-10 text-center">
          <div class="mx-auto flex h-14 w-14 items-center justify-center rounded-2xl bg-white/5 text-xl font-extrabold text-cyan-200 ring-1 ring-white/10">
            Q
          </div>
          <h2 class="mt-4 text-xl font-extrabold text-white">
            {{ activeSearch ? 'No pending businesses match your search.' : 'No pending businesses found.' }}
          </h2>
          <p class="mt-2 text-sm font-semibold text-slate-400">
            Hồ sơ business pending mới sẽ xuất hiện tại đây.
          </p>
        </div>

        <div v-else>
          <div class="hidden overflow-x-auto lg:block">
            <table class="min-w-full divide-y divide-white/10">
              <thead class="bg-white/[0.04]">
                <tr>
                  <th class="px-5 py-3 text-left text-xs font-extrabold uppercase tracking-wide text-slate-400">Tên Công ty</th>
                  <th class="px-5 py-3 text-left text-xs font-extrabold uppercase tracking-wide text-slate-400">Mã số thuế</th>
                  <th class="px-5 py-3 text-left text-xs font-extrabold uppercase tracking-wide text-slate-400">Email liên hệ</th>
                  <th class="px-5 py-3 text-left text-xs font-extrabold uppercase tracking-wide text-slate-400">Thời gian nộp</th>
                  <th class="px-5 py-3 text-left text-xs font-extrabold uppercase tracking-wide text-slate-400">Trạng thái</th>
                  <th class="px-5 py-3 text-right text-xs font-extrabold uppercase tracking-wide text-slate-400">Thao tác</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-white/10">
                <tr
                  v-for="item in items"
                  :key="item.business_id"
                  class="transition-colors hover:bg-cyan-400/10"
                >
                  <td class="px-5 py-4 align-middle">
                    <p class="text-sm font-extrabold text-white">{{ item.company_name }}</p>
                    <p class="text-xs font-semibold text-slate-400">ID {{ item.business_id }}</p>
                  </td>
                  <td class="px-5 py-4 align-middle text-sm font-semibold text-slate-300">{{ item.tax_code }}</td>
                  <td class="px-5 py-4 align-middle text-sm font-semibold text-slate-300">{{ item.email }}</td>
                  <td class="px-5 py-4 align-middle text-sm font-semibold text-slate-300">{{ formatDate(item.created_at) }}</td>
                  <td class="px-5 py-4 align-middle">
                    <span class="inline-flex rounded-full bg-amber-400/10 px-3 py-1 text-xs font-extrabold uppercase tracking-wide text-amber-200 ring-1 ring-amber-400/25">
                      Chờ duyệt
                    </span>
                  </td>
                  <td class="px-5 py-4 text-right align-middle">
                    <button
                      type="button"
                      class="rounded-xl bg-cyan-400 px-4 py-2 text-sm font-extrabold text-slate-950 transition-colors hover:bg-cyan-300"
                      @click="openKYBDetail(item.business_id)"
                    >
                      Xem hồ sơ
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <div class="space-y-3 p-4 lg:hidden">
            <article
              v-for="item in items"
              :key="item.business_id"
              class="rounded-2xl border border-white/10 bg-white/[0.04] p-4"
            >
              <div class="flex items-start justify-between gap-3">
                <div>
                  <h2 class="text-base font-extrabold text-white">{{ item.company_name }}</h2>
                  <p class="mt-1 text-sm font-semibold text-slate-400">{{ item.email }}</p>
                </div>
                <span class="rounded-full bg-amber-400/10 px-3 py-1 text-xs font-extrabold uppercase tracking-wide text-amber-200 ring-1 ring-amber-400/25">
                  Chờ duyệt
                </span>
              </div>
              <dl class="mt-4 grid gap-3 text-sm">
                <div>
                  <dt class="text-xs font-extrabold uppercase tracking-wide text-slate-500">Mã số thuế</dt>
                  <dd class="mt-1 font-semibold text-slate-200">{{ item.tax_code }}</dd>
                </div>
                <div>
                  <dt class="text-xs font-extrabold uppercase tracking-wide text-slate-500">Thời gian nộp</dt>
                  <dd class="mt-1 font-semibold text-slate-200">{{ formatDate(item.created_at) }}</dd>
                </div>
              </dl>
              <button
                type="button"
                class="mt-4 w-full rounded-xl bg-cyan-400 px-4 py-2 text-sm font-extrabold text-slate-950 transition-colors hover:bg-cyan-300"
                @click="openKYBDetail(item.business_id)"
              >
                Xem hồ sơ
              </button>
            </article>
          </div>
        </div>

        <div
          v-if="!errorMessage && !isLoading"
          class="flex flex-col gap-3 border-t border-white/10 px-5 py-4 sm:flex-row sm:items-center sm:justify-between"
        >
          <p class="text-sm font-semibold text-slate-400">
            Page {{ pagination.page }} of {{ pagination.total_pages || 1 }} · {{ pagination.total }} result{{ pagination.total === 1 ? '' : 's' }}
          </p>
          <div class="flex gap-2">
            <button
              type="button"
              class="rounded-xl border border-white/10 bg-white/5 px-4 py-2 text-sm font-extrabold text-slate-200 transition-colors hover:bg-white/10 disabled:cursor-not-allowed disabled:opacity-40"
              :disabled="pagination.page <= 1"
              @click="goToPage(pagination.page - 1)"
            >
              Previous
            </button>
            <button
              type="button"
              class="rounded-xl bg-cyan-400 px-4 py-2 text-sm font-extrabold text-slate-950 transition-colors hover:bg-cyan-300 disabled:cursor-not-allowed disabled:opacity-40"
              :disabled="pagination.total_pages === 0 || pagination.page >= pagination.total_pages"
              @click="goToPage(pagination.page + 1)"
            >
              Next
            </button>
          </div>
        </div>
      </section>
    </div>

    <div
      v-if="showDetailModal"
      class="fixed inset-0 z-50 flex items-center justify-center bg-slate-950/80 px-4 py-6 backdrop-blur-sm"
      @click.self="closeDetailModal"
    >
      <section class="max-h-[92vh] w-full max-w-3xl overflow-y-auto rounded-3xl border border-white/10 bg-slate-950 p-5 shadow-2xl shadow-slate-950 ring-1 ring-cyan-400/20 sm:p-6">
        <div class="flex items-start justify-between gap-4 border-b border-white/10 pb-4">
          <div>
            <p class="text-xs font-extrabold uppercase tracking-wide text-cyan-200">Business KYB Detail</p>
            <h2 class="mt-2 text-2xl font-extrabold text-white">
              {{ detail?.company_name || 'Loading...' }}
            </h2>
            <p class="mt-1 text-sm font-semibold text-slate-400">{{ detail?.email }}</p>
          </div>
          <button
            type="button"
            class="rounded-xl border border-white/10 bg-white/5 px-3 py-2 text-sm font-extrabold text-slate-300 hover:bg-white/10 hover:text-white"
            @click="closeDetailModal"
          >
            Close
          </button>
        </div>

        <div v-if="detailLoading" class="mt-5 space-y-3">
          <div v-for="index in 5" :key="index" class="h-14 animate-pulse rounded-xl bg-white/[0.06]" />
        </div>

        <div v-else-if="detailError" class="mt-5 rounded-2xl border border-rose-400/20 bg-rose-400/10 p-4">
          <p class="text-sm font-bold text-rose-100">{{ detailError }}</p>
        </div>

        <div v-else-if="detail" class="mt-5 space-y-5">
          <dl class="grid gap-3 sm:grid-cols-2">
            <div v-for="field in detailFields" :key="field.label" class="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
              <dt class="text-xs font-extrabold uppercase tracking-wide text-slate-500">{{ field.label }}</dt>
              <dd class="mt-2 break-words text-sm font-bold text-slate-100">{{ field.value || 'N/A' }}</dd>
            </div>
          </dl>

          <div class="rounded-2xl border border-amber-400/20 bg-amber-400/10 p-4">
            <p class="text-xs font-extrabold uppercase tracking-wide text-amber-200">Status</p>
            <p class="mt-2 text-sm font-bold text-white">
              {{ detail.status }} · Verified: {{ detail.is_verified ? 'Yes' : 'No' }}
            </p>
          </div>

          <div v-if="actionMessage" class="rounded-2xl border p-4" :class="actionMessage.type === 'error' ? 'border-rose-400/20 bg-rose-400/10 text-rose-100' : 'border-emerald-400/20 bg-emerald-400/10 text-emerald-100'">
            <p class="text-sm font-bold">{{ actionMessage.text }}</p>
          </div>

          <div class="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
            <label class="block">
              <span class="text-xs font-extrabold uppercase tracking-wide text-slate-400">
                Lý do từ chối
              </span>
              <textarea
                v-model="rejectReason"
                rows="3"
                class="mt-2 w-full rounded-xl border border-white/10 bg-slate-950/80 px-4 py-3 text-sm font-semibold text-white outline-none transition focus:border-cyan-300 focus:ring-2 focus:ring-cyan-400/30"
                placeholder="Nhập lý do từ chối tối thiểu 10 ký tự"
              />
            </label>
            <p class="mt-2 text-xs font-semibold text-slate-400">
              Frontend sẽ kiểm tra lý do, backend vẫn validate lại để đảm bảo an toàn.
            </p>
          </div>

          <div class="flex flex-col gap-3 sm:flex-row sm:justify-end">
            <button
              type="button"
              class="rounded-xl border border-rose-400/25 bg-rose-400/10 px-5 py-3 text-sm font-extrabold text-rose-100 transition-colors hover:bg-rose-300 hover:text-slate-950 disabled:cursor-not-allowed disabled:opacity-60"
              :disabled="isReviewing"
              @click="submitReview('rejected')"
            >
              Từ chối hồ sơ
            </button>
            <button
              type="button"
              class="rounded-xl bg-cyan-400 px-5 py-3 text-sm font-extrabold text-slate-950 transition-colors hover:bg-cyan-300 disabled:cursor-not-allowed disabled:opacity-60"
              :disabled="isReviewing"
              @click="submitReview('approved')"
            >
              Duyệt doanh nghiệp
            </button>
          </div>
        </div>
      </section>
    </div>
  </AdminShell>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import AdminShell from '~/components/admin/AdminShell.vue'
import { useAdminApi } from '~/composables/useAdminApi'
import type { BusinessKYBDetail, PendingBusinessItem, PaginationMeta, ReviewBusinessPayload } from '~/types/admin'

definePageMeta({
  middleware: 'auth'
})

const adminApi = useAdminApi()

const items = ref<PendingBusinessItem[]>([])
const searchInput = ref('')
const activeSearch = ref('')
const isLoading = ref(false)
const errorMessage = ref('')
const pagination = ref<PaginationMeta>({
  page: 1,
  limit: 10,
  total: 0,
  total_pages: 0
})

const showDetailModal = ref(false)
const detail = ref<BusinessKYBDetail | null>(null)
const detailLoading = ref(false)
const detailError = ref('')
const rejectReason = ref('')
const isReviewing = ref(false)
const actionMessage = ref<{ type: 'success' | 'error'; text: string } | null>(null)

let debounceTimer: ReturnType<typeof setTimeout> | null = null

const getApiError = (error: unknown) => {
  return error as { response?: { status?: number; _data?: { message?: string } } }
}

const extractErrorMessage = (error: unknown, fallback: string) => {
  return getApiError(error).response?._data?.message || fallback
}

const fetchPendingBusinesses = async () => {
  isLoading.value = true
  errorMessage.value = ''

  try {
    const response = await adminApi.getPendingBusinesses({
      page: pagination.value.page,
      limit: pagination.value.limit,
      search: activeSearch.value
    })

    items.value = response.items || []
    pagination.value = response.pagination
  } catch (error: unknown) {
    const status = getApiError(error).response?.status
    if (status === 403) {
      errorMessage.value = 'You do not have permission to view pending businesses.'
    } else {
      errorMessage.value = extractErrorMessage(error, 'Failed to load pending businesses.')
    }
    items.value = []
  } finally {
    isLoading.value = false
  }
}

const openKYBDetail = async (businessId: number) => {
  showDetailModal.value = true
  detail.value = null
  detailError.value = ''
  actionMessage.value = null
  rejectReason.value = ''
  detailLoading.value = true

  try {
    detail.value = await adminApi.getBusinessDetail(businessId)
    rejectReason.value = detail.value.reject_reason || ''
  } catch (error: unknown) {
    detailError.value = extractErrorMessage(error, 'Không thể tải hồ sơ doanh nghiệp.')
  } finally {
    detailLoading.value = false
  }
}

const closeDetailModal = () => {
  if (isReviewing.value) return
  resetDetailModal()
}

const resetDetailModal = () => {
  showDetailModal.value = false
  detail.value = null
  actionMessage.value = null
  rejectReason.value = ''
}

const submitReview = async (decision: ReviewBusinessPayload['decision']) => {
  if (!detail.value) return

  const reason = rejectReason.value.trim()
  if (decision === 'rejected' && reason.length < 10) {
    actionMessage.value = { type: 'error', text: 'Lý do từ chối phải có ít nhất 10 ký tự.' }
    return
  }

  isReviewing.value = true
  actionMessage.value = null

  try {
    await adminApi.reviewBusiness(detail.value.business_id, {
      decision,
      reject_reason: decision === 'rejected' ? reason : ''
    })

    // Sau khi duyệt/từ chối phải tải lại queue để hồ sơ đã xử lý biến mất khỏi danh sách pending.
    resetDetailModal()
    await fetchPendingBusinesses()
  } catch (error: unknown) {
    const message = extractErrorMessage(error, 'Không thể cập nhật kết quả KYB.')
    actionMessage.value = { type: 'error', text: message }
  } finally {
    isReviewing.value = false
  }
}

const applySearchNow = () => {
  activeSearch.value = searchInput.value.trim()
  pagination.value.page = 1
  fetchPendingBusinesses()
}

const clearSearch = () => {
  searchInput.value = ''
  activeSearch.value = ''
  pagination.value.page = 1
  fetchPendingBusinesses()
}

const goToPage = (page: number) => {
  if (page < 1 || (pagination.value.total_pages > 0 && page > pagination.value.total_pages)) return
  pagination.value.page = page
  fetchPendingBusinesses()
}

const formatDate = (value: string | null) => {
  if (!value) return 'N/A'
  return new Intl.DateTimeFormat('vi-VN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  }).format(new Date(value))
}

const detailFields = computed(() => {
  if (!detail.value) return []

  return [
    { label: 'Tên doanh nghiệp', value: detail.value.company_name },
    { label: 'Email', value: detail.value.email },
    { label: 'Mã số thuế', value: detail.value.tax_code },
    { label: 'Điện thoại', value: detail.value.phone },
    { label: 'Địa chỉ', value: detail.value.address },
    { label: 'Ngày đăng ký', value: formatDate(detail.value.created_at) },
    { label: 'Reviewed At', value: formatDate(detail.value.reviewed_at) },
    { label: 'Logo URL', value: detail.value.logo_url }
  ]
})

watch(searchInput, () => {
  if (debounceTimer) clearTimeout(debounceTimer)
  debounceTimer = setTimeout(() => {
    if (searchInput.value.trim() !== activeSearch.value) {
      applySearchNow()
    }
  }, 400)
})

onMounted(fetchPendingBusinesses)
</script>
