<template>
  <AdminShell>
    <div class="space-y-6">
      <!-- Hero Banner -->
      <section class="overflow-hidden rounded-3xl border border-cyan-500/20 bg-slate-950/80 p-6 sm:p-8 shadow-2xl shadow-slate-950/40 backdrop-blur-xl">
        <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-6">
          <div class="space-y-2">
            <span class="inline-flex rounded-full bg-amber-400/10 px-3.5 py-1 text-xs font-black uppercase tracking-wider text-amber-300 ring-1 ring-amber-400/30">
              Hàng Chờ Duyệt Ca Làm Việc
            </span>
            <h1 class="text-2xl sm:text-3xl font-black tracking-tight text-white">
              Duyệt Bài Đăng Tuyển Dụng (Job Approval)
            </h1>
            <p class="text-xs sm:text-sm font-semibold text-slate-300 max-w-2xl leading-relaxed">
              Các ca làm việc do Doanh nghiệp đăng tải cần được Admin kiểm duyệt nội dung, mức lương và thông tin trước khi phát hành lên sàn việc làm cho Sinh viên.
            </p>
          </div>

          <div class="flex items-center gap-3">
            <div class="rounded-2xl border border-cyan-500/20 bg-cyan-500/10 px-5 py-3.5 text-center">
              <p class="text-[10px] font-black uppercase tracking-wider text-cyan-300">Đang chờ duyệt</p>
              <p class="text-3xl font-black text-white mt-0.5">{{ pagination.total }} ca</p>
            </div>
          </div>
        </div>
      </section>

      <!-- Search & Filter Controls -->
      <section class="rounded-3xl border border-cyan-500/20 bg-slate-900/90 p-4 shadow-xl backdrop-blur-xl">
        <div class="relative">
          <span class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-slate-500">
            <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </span>
          <input
            v-model="searchQuery"
            type="text"
            @input="handleSearch"
            class="block w-full pl-10 pr-4 py-2.5 border border-cyan-500/20 rounded-xl text-xs bg-slate-950/80 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium transition-all"
            placeholder="Tìm theo tên bài tuyển dụng, tên công ty hoặc địa điểm làm việc..."
          />
        </div>
      </section>

      <!-- Loading State -->
      <div v-if="isLoading" class="p-12 text-center text-xs font-black text-slate-400 animate-pulse bg-slate-900/90 rounded-3xl border border-cyan-500/15">
        ⏳ Đang tải danh sách bài tuyển dụng chờ duyệt...
      </div>

      <!-- Empty State -->
      <div v-else-if="jobs.length === 0" class="rounded-3xl border border-cyan-500/15 bg-slate-900/90 p-12 text-center shadow-xl space-y-3">
        <div class="inline-flex p-4 rounded-2xl bg-cyan-500/10 border border-cyan-500/20 text-cyan-300">
          <svg class="h-10 w-10" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </div>
        <h3 class="text-base font-extrabold text-white">Không Có Ca Làm Nào Đang Chờ Duyệt</h3>
        <p class="text-xs text-slate-400 max-w-sm mx-auto font-medium">
          Tất cả các tin tuyển dụng đăng tải gần đây đều đã được duyệt hoặc chưa có bài đăng mới từ Doanh nghiệp.
        </p>
      </div>

      <!-- Job List Cards -->
      <div v-else class="space-y-4">
        <div
          v-for="job in jobs"
          :key="job.id"
          class="rounded-3xl border border-cyan-500/20 bg-slate-900/90 p-6 shadow-xl space-y-4 hover:border-cyan-500/40 transition-all backdrop-blur-xl"
        >
          <div class="flex flex-col sm:flex-row sm:items-start justify-between gap-4 border-b border-cyan-500/10 pb-4">
            <div class="flex items-start space-x-3.5">
              <div class="h-12 w-12 rounded-2xl bg-gradient-to-tr from-cyan-500 via-blue-600 to-emerald-400 flex items-center justify-center text-white font-black text-lg shadow-md shadow-cyan-500/20 flex-shrink-0">
                {{ job.company_name ? job.company_name.substring(0, 2).toUpperCase() : 'DN' }}
              </div>
              <div>
                <div class="flex items-center gap-2">
                  <h3 class="text-base font-black text-white tracking-tight">{{ job.title }}</h3>
                  <span class="rounded-full bg-amber-400/10 px-2.5 py-0.5 text-[10px] font-black uppercase tracking-wider text-amber-300 ring-1 ring-amber-400/30">
                    Chờ Admin Duyệt
                  </span>
                </div>
                <p class="text-xs font-bold text-cyan-300 mt-0.5">
                  Công ty: <strong class="text-white">{{ job.company_name || 'Doanh nghiệp chưa xác định' }}</strong> (MST: {{ job.tax_code || 'N/A' }})
                </p>
              </div>
            </div>

            <div class="text-left sm:text-right flex-shrink-0">
              <p class="text-lg font-black text-emerald-300 tracking-tight">
                {{ Number(job.salary || 0).toLocaleString('vi-VN') }} VNĐ
              </p>
              <p class="text-[11px] font-semibold text-slate-400 mt-0.5">
                Tuyển {{ job.slots }} vị trí • {{ job.job_type || 'Part-time' }}
              </p>
            </div>
          </div>

          <!-- Description & Details -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-4 text-xs font-medium text-slate-300">
            <div>
              <span class="text-slate-400 block text-[10px] uppercase font-black">Địa điểm làm việc:</span>
              <span class="text-white font-semibold">{{ job.location || 'Chưa cập nhật' }}</span>
            </div>
            <div>
              <span class="text-slate-400 block text-[10px] uppercase font-black">Thời gian ca làm:</span>
              <span class="text-white font-semibold">{{ job.working_date || 'Thỏa thuận' }}</span>
            </div>
            <div>
              <span class="text-slate-400 block text-[10px] uppercase font-black">Ngành nghề:</span>
              <span class="text-cyan-200 font-bold bg-cyan-500/10 px-2 py-0.5 rounded-lg border border-cyan-500/20 inline-block mt-0.5">
                {{ job.category || 'Công nghệ thông tin' }}
              </span>
            </div>
          </div>

          <div class="bg-slate-950/60 p-3.5 rounded-2xl border border-cyan-500/10 text-xs text-slate-300 leading-relaxed">
            <span class="text-[10px] font-black uppercase text-cyan-400 block mb-1">Mô tả công việc:</span>
            {{ job.description }}
          </div>

          <!-- Action Buttons -->
          <div class="flex items-center justify-end gap-3 pt-2">
            <button
              @click="handleReview(job.id, 'rejected')"
              :disabled="processingId === job.id"
              class="px-4 py-2 rounded-xl text-xs font-extrabold text-rose-300 border border-rose-500/30 bg-rose-500/10 hover:bg-rose-500/20 transition-all disabled:opacity-50"
            >
              🔴 Từ Chối Đăng Bài
            </button>

            <button
              @click="handleReview(job.id, 'approved')"
              :disabled="processingId === job.id"
              class="px-5 py-2.5 rounded-xl text-xs font-black text-white bg-gradient-to-r from-cyan-500 via-blue-600 to-emerald-500 hover:from-cyan-400 hover:to-emerald-400 shadow-md shadow-cyan-500/20 transition-all disabled:opacity-50"
            >
              🟢 Phê Duyệt Bài Đăng
            </button>
          </div>
        </div>
      </div>

      <!-- Pagination -->
      <PaginationControls
        v-if="jobs.length > 0"
        :page="pagination.page"
        :page-size="pagination.limit"
        :total-items="pagination.total"
        @update:page="handlePageChange"
      />
    </div>
  </AdminShell>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import AdminShell from '~/components/admin/AdminShell.vue'
import PaginationControls from '~/components/common/PaginationControls.vue'
import { useApi } from '~/composables/useApi'
import { useToast } from '~/composables/useToast'

definePageMeta({
  middleware: 'auth'
})

const api = useApi()
const { success, error } = useToast()

const jobs = ref<any[]>([])
const isLoading = ref(false)
const processingId = ref<number | null>(null)
const searchQuery = ref('')

const pagination = reactive({
  page: 1,
  limit: 10,
  total: 0,
  totalPages: 0
})

const fetchPendingJobs = async () => {
  isLoading.value = true
  try {
    const res = await api.get('/api/admin/jobs/pending', {
      params: {
        page: pagination.page,
        limit: pagination.limit,
        search: searchQuery.value
      }
    })
    jobs.value = res.items || []
    if (res.pagination) {
      pagination.total = res.pagination.total
      pagination.totalPages = res.pagination.totalPages
    }
  } catch (err: any) {
    console.error('Fetch pending jobs error:', err)
    error('Không thể tải danh sách bài tuyển dụng chờ duyệt')
  } finally {
    isLoading.value = false
  }
}

const handleReview = async (jobId: number, status: 'approved' | 'rejected') => {
  processingId.value = jobId
  try {
    await api.put(`/api/admin/jobs/${jobId}/status`, { status })
    if (status === 'approved') {
      success('🎉 Đã duyệt phát hành bài đăng tuyển dụng thành công!')
    } else {
      success('Đã từ chối bài đăng tuyển dụng.')
    }
    await fetchPendingJobs()
  } catch (err: any) {
    console.error('Review job error:', err)
    error(err?.data?.message || 'Không thể cập nhật trạng thái bài tuyển dụng')
  } finally {
    processingId.value = null
  }
}

let searchTimer: any = null
const handleSearch = () => {
  if (searchTimer) clearTimeout(searchTimer)
  searchTimer = setTimeout(() => {
    pagination.page = 1
    fetchPendingJobs()
  }, 350)
}

const handlePageChange = (newPage: number) => {
  pagination.page = newPage
  fetchPendingJobs()
}

onMounted(() => {
  fetchPendingJobs()
})
</script>
