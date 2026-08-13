<template>
  <AdminShell>
    <div class="space-y-6">
      <!-- Header Banner -->
      <section class="rounded-3xl border border-white/10 bg-slate-950/80 p-6 shadow-2xl shadow-slate-950/40 ring-1 ring-cyan-400/10 lg:p-8">
        <div class="flex flex-col gap-6">
          <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
            <div>
              <h1 class="text-3xl font-extrabold tracking-tight text-white sm:text-4xl">
                Quản lý Doanh nghiệp Hệ thống
              </h1>
              <p class="mt-2 text-sm font-semibold text-slate-300">
                Quản lý danh sách doanh nghiệp, xem thông tin KYB chi tiết, khóa hoặc mở khóa quyền hoạt động của tài khoản.
              </p>
            </div>

            <NuxtLink
              to="/admin/dashboard"
              class="inline-flex items-center justify-center rounded-xl border border-white/10 bg-white/10 px-4 py-2.5 text-sm font-bold text-slate-200 transition-colors hover:bg-white/15"
            >
              Quay lại Dashboard
            </NuxtLink>
          </div>

          <AdminHeaderNav />
        </div>
      </section>

      <!-- Filters & Search Bar -->
      <section class="rounded-2xl border border-white/10 bg-slate-950/70 p-5 shadow-xl shadow-slate-950/25 ring-1 ring-cyan-400/10">
        <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
          <div class="relative flex-1">
            <input
              v-model="searchQuery"
              type="text"
              placeholder="Tìm kiếm theo Tên công ty, Mã số thuế, Email..."
              class="w-full rounded-xl border border-white/10 bg-slate-900/90 px-4 py-2.5 text-sm text-white placeholder-slate-400 focus:border-cyan-400 focus:outline-none focus:ring-1 focus:ring-cyan-400"
              @keyup.enter="fetchBusinesses"
            />
          </div>

          <div class="flex items-center gap-3">
            <select
              v-model="statusFilter"
              class="rounded-xl border border-white/10 bg-slate-900/90 px-3.5 py-2.5 text-sm text-white focus:border-cyan-400 focus:outline-none"
              @change="fetchBusinesses"
            >
              <option value="">Tất cả trạng thái</option>
              <option value="approved">🟢 Đã phê duyệt (Approved)</option>
              <option value="pending">🟡 Chờ duyệt (Pending)</option>
              <option value="locked">🔴 Đã bị khóa (Locked)</option>
            </select>

            <button
              @click="fetchBusinesses"
              class="inline-flex items-center justify-center rounded-xl bg-cyan-400 px-4 py-2.5 text-sm font-extrabold text-slate-950 transition-colors hover:bg-cyan-300"
            >
              Tìm kiếm
            </button>
          </div>
        </div>
      </section>

      <!-- Business Data Table -->
      <section class="overflow-hidden rounded-2xl border border-white/10 bg-slate-950/70 shadow-xl shadow-slate-950/25 ring-1 ring-cyan-400/10">
        <div v-if="isLoading" class="p-8 text-center text-slate-400">
          <span class="animate-pulse text-sm font-bold">Đang tải danh sách doanh nghiệp...</span>
        </div>

        <div v-else-if="businesses.length === 0" class="p-8 text-center text-slate-400">
          <p class="text-base font-semibold">Không tìm thấy doanh nghiệp nào phù hợp.</p>
        </div>

        <div v-else class="overflow-x-auto">
          <table class="w-full text-left text-sm text-slate-200">
            <thead class="bg-slate-900/90 text-xs uppercase text-slate-400">
              <tr>
                <th class="px-5 py-3.5 font-extrabold">Tên Công ty & Email</th>
                <th class="px-5 py-3.5 font-extrabold">Mã số thuế & SĐT</th>
                <th class="px-5 py-3.5 font-extrabold">Địa chỉ</th>
                <th class="px-5 py-3.5 font-extrabold">Số Job</th>
                <th class="px-5 py-3.5 font-extrabold">Trạng thái</th>
                <th class="px-5 py-3.5 font-extrabold text-right">Thao tác</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-white/5">
              <tr v-for="biz in businesses" :key="biz.business_id" class="hover:bg-white/[0.02]">
                <td class="px-5 py-4">
                  <div class="flex items-center gap-3">
                    <img :src="biz.logo_url || 'https://placehold.co/100x100/0f172a/fff?text=' + biz.company_name" class="h-10 w-10 rounded-lg object-cover bg-slate-800 border border-white/10" />
                    <div>
                      <p class="font-extrabold text-white flex items-center gap-1.5">
                        {{ biz.company_name }}
                        <span v-if="biz.is_verified" class="text-cyan-400" title="Đã xác thực KYB">✓</span>
                      </p>
                      <p class="text-xs text-slate-400">{{ biz.email }}</p>
                    </div>
                  </div>
                </td>

                <td class="px-5 py-4">
                  <p class="font-bold text-slate-200">MST: {{ biz.tax_code || 'N/A' }}</p>
                  <p class="text-xs text-slate-400">SĐT: {{ biz.phone || 'Chưa cập nhật' }}</p>
                </td>

                <td class="px-5 py-4 max-w-xs truncate text-xs text-slate-300">
                  {{ biz.address || 'Chưa có thông tin địa chỉ' }}
                </td>

                <td class="px-5 py-4 font-extrabold text-cyan-300">
                  {{ biz.job_count }} tin
                </td>

                <td class="px-5 py-4">
                  <span
                    class="inline-flex items-center gap-1.5 rounded-full px-2.5 py-1 text-xs font-bold border"
                    :class="{
                      'bg-emerald-400/10 text-emerald-300 border-emerald-400/30': biz.status === 'approved',
                      'bg-amber-400/10 text-amber-300 border-amber-400/30': biz.status === 'pending',
                      'bg-rose-400/10 text-rose-300 border-rose-400/30': biz.status === 'locked' || biz.status === 'rejected'
                    }"
                  >
                    <span>{{ getStatusLabel(biz.status) }}</span>
                  </span>
                </td>

                <td class="px-5 py-4 text-right space-x-2">
                  <button
                    @click="viewKYBDetail(biz.business_id)"
                    class="rounded-lg bg-white/10 px-3 py-1.5 text-xs font-bold text-slate-200 hover:bg-white/20 transition-colors"
                  >
                    Hồ sơ KYB
                  </button>

                  <button
                    v-if="biz.status === 'locked'"
                    @click="toggleStatus(biz, 'approved')"
                    class="rounded-lg bg-emerald-400/20 px-3 py-1.5 text-xs font-extrabold text-emerald-300 hover:bg-emerald-400/30 border border-emerald-400/30 transition-colors"
                  >
                    Mở khóa
                  </button>

                  <button
                    v-else
                    @click="toggleStatus(biz, 'locked')"
                    class="rounded-lg bg-rose-400/20 px-3 py-1.5 text-xs font-extrabold text-rose-300 hover:bg-rose-400/30 border border-rose-400/30 transition-colors"
                  >
                    Khóa doanh nghiệp
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </section>
    </div>

    <!-- KYB Detail Modal -->
    <div v-if="selectedKYB" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-sm">
      <div class="w-full max-w-lg rounded-2xl border border-white/10 bg-slate-900 p-6 shadow-2xl space-y-4">
        <div class="flex items-center justify-between border-b border-white/10 pb-3">
          <h3 class="text-lg font-extrabold text-white">Chi tiết Hồ sơ KYB Doanh nghiệp</h3>
          <button @click="selectedKYB = null" class="text-slate-400 hover:text-white text-lg font-bold">✕</button>
        </div>

        <div class="space-y-3 text-sm text-slate-300">
          <div class="flex items-center gap-3">
            <img :src="selectedKYB.logo_url" class="h-14 w-14 rounded-xl object-cover bg-slate-800 border border-cyan-400/30" />
            <div>
              <p class="text-base font-extrabold text-white">{{ selectedKYB.company_name }}</p>
              <p class="text-xs text-cyan-300 font-semibold">MST: {{ selectedKYB.tax_code }}</p>
            </div>
          </div>

          <div class="grid grid-cols-2 gap-2 bg-slate-950/60 p-3 rounded-xl border border-white/5">
            <div>
              <span class="text-xs text-slate-400 block font-bold">EMAIL ĐĂNG KÝ</span>
              <span class="font-semibold text-white">{{ selectedKYB.email }}</span>
            </div>
            <div>
              <span class="text-xs text-slate-400 block font-bold">SỐ ĐIỆN THOẠI</span>
              <span class="font-semibold text-white">{{ selectedKYB.phone || 'N/A' }}</span>
            </div>
          </div>

          <div>
            <span class="text-xs text-slate-400 block font-bold mb-1">ĐỊA CHỈ TRỤ SỞ</span>
            <p class="bg-slate-950/60 p-3 rounded-xl border border-white/5 text-xs font-medium text-slate-200">
              {{ selectedKYB.address || 'Chưa cập nhật' }}
            </p>
          </div>
        </div>

        <div class="pt-4 border-t border-white/10 flex justify-end">
          <button @click="selectedKYB = null" class="rounded-xl bg-cyan-400 px-5 py-2 text-xs font-extrabold text-slate-950">
            Đóng cửa sổ
          </button>
        </div>
      </div>
    </div>
  </AdminShell>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import AdminShell from '~/components/admin/AdminShell.vue'
import { useApi } from '~/composables/useApi'
import { useToast } from '~/composables/useToast'

interface BusinessItem {
  business_id: number
  user_id: number
  company_name: string
  tax_code: string
  email: string
  phone: string
  address: string
  logo_url: string
  status: string
  is_verified: boolean
  job_count: number
  created_at: string
}

const api = useApi()
const { success, error } = useToast()

const businesses = ref<BusinessItem[]>([])
const isLoading = ref(false)
const searchQuery = ref('')
const statusFilter = ref('')
const selectedKYB = ref<any | null>(null)

const fetchBusinesses = async () => {
  isLoading.value = true
  try {
    const params = new URLSearchParams()
    if (searchQuery.value) params.append('search', searchQuery.value)
    if (statusFilter.value) params.append('status', statusFilter.value)

    const res = await api.get<{ items: BusinessItem[] }>(`/api/admin/businesses?${params.toString()}`)
    businesses.value = res.items || []
  } catch (err: any) {
    error('Không thể tải danh sách doanh nghiệp')
  } finally {
    isLoading.value = false
  }
}

const viewKYBDetail = async (id: number) => {
  try {
    const res = await api.get(`/api/admin/businesses/${id}`)
    selectedKYB.value = res
  } catch (err: any) {
    error('Không thể lấy thông tin KYB chi tiết')
  }
}

const toggleStatus = async (biz: BusinessItem, newStatus: string) => {
  const actionText = newStatus === 'locked' ? 'Khóa' : 'Mở khóa'
  if (!confirm(`Bạn có chắc chắn muốn ${actionText} tài khoản của doanh nghiệp "${biz.company_name}"?`)) {
    return
  }

  try {
    await api.put(`/api/admin/businesses/${biz.business_id}/status`, { status: newStatus })
    biz.status = newStatus
    success(`Đã ${actionText} tài khoản doanh nghiệp thành công!`)
  } catch (err: any) {
    error(`Không thể ${actionText} tài khoản doanh nghiệp`)
  }
}

const getStatusLabel = (status: string) => {
  switch (status) {
    case 'approved': return '🟢 Hoạt động'
    case 'pending': return '🟡 Chờ duyệt KYB'
    case 'locked': return '🔴 Đã bị khóa'
    case 'rejected': return '🔴 Từ chối KYB'
    default: return status
  }
}

onMounted(() => {
  fetchBusinesses()
})
</script>
