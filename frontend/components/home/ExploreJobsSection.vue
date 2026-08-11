<template>
  <section :id="sectionId" class="relative">
    <div :class="standalone ? 'dashboard-body px-4 sm:px-6 lg:px-8 py-10' : 'px-4 sm:px-6 lg:px-8 py-16 lg:py-20'">
      <div class="mx-auto max-w-7xl">
        <div v-if="showHero" class="mx-auto mb-10 max-w-3xl rounded-2xl bg-gradient-to-br from-slate-950 via-slate-900 to-cyan-950 px-6 py-10 text-center shadow-xl shadow-slate-950/10">
          <h2 class="text-sm font-bold uppercase tracking-wide text-cyan-300">Vị Trí Mở Tuyển</h2>
          <p class="mt-2 text-3xl font-extrabold tracking-tight text-white sm:text-4xl">
            Việc làm bán thời gian phù hợp với sinh viên
          </p>
          <p class="mt-4 text-base text-slate-300 sm:text-lg">
            Khám phá danh sách việc làm uy tín được đăng tuyển bởi các doanh nghiệp đã xác thực.
          </p>
        </div>

        <div v-else class="mb-8 flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
          <div>
            <p class="text-sm font-bold uppercase tracking-[0.24em] text-cyan-300">Khám Phá Việc Làm</p>
            <h2 class="mt-2 text-3xl font-extrabold tracking-tight text-white sm:text-4xl">
              Ứng tuyển việc làm sinh viên đã xác thực
            </h2>
            <p class="mt-3 max-w-2xl text-sm font-medium leading-6 text-slate-400 sm:text-base">
              Tìm kiếm việc làm nhanh chóng với bộ lọc từ khóa, địa điểm và phân trang mượt mà.
            </p>
          </div>
          <span class="inline-flex w-fit items-center rounded-full border border-cyan-400/20 bg-cyan-400/10 px-4 py-2 text-sm font-bold text-cyan-100">
            {{ filteredJobs.length }} kết quả khả dụng
          </span>
        </div>

        <div class="mx-auto mb-10 grid max-w-5xl grid-cols-1 gap-4 rounded-xl border border-cyan-400/15 bg-slate-900/82 p-4 shadow-lg shadow-slate-950/30 backdrop-blur sm:grid-cols-2">
          <div class="relative">
            <span class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3 text-slate-500">
              <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </span>
            <input
              v-model="searchQuery"
              type="text"
              class="block w-full rounded-lg border border-white/10 bg-slate-950/70 py-2.5 pl-10 pr-3 text-sm text-slate-100 placeholder-slate-500 transition-all duration-200 focus:border-cyan-300 focus:outline-none focus:ring-2 focus:ring-cyan-400"
              placeholder="Tìm theo tiêu đề hoặc từ khóa công việc..."
            />
          </div>
          <div class="relative">
            <span class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3 text-slate-500">
              <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a2 2 0 01-2.828 0l-4.243-4.243a8 8 0 1111.314 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
            </span>
            <input
              v-model="locationQuery"
              type="text"
              class="block w-full rounded-lg border border-white/10 bg-slate-950/70 py-2.5 pl-10 pr-3 text-sm text-slate-100 placeholder-slate-500 transition-all duration-200 focus:border-cyan-300 focus:outline-none focus:ring-2 focus:ring-cyan-400"
              placeholder="Nhập địa điểm, thành phố..."
            />
          </div>
        </div>

        <div v-if="feedback" :class="[
          feedback.type === 'success' ? 'border-emerald-300 bg-emerald-500/10 text-emerald-300' : 'border-rose-300 bg-rose-500/10 text-rose-300',
          'mx-auto mb-8 flex max-w-4xl items-start justify-between rounded-r-lg border-l-4 p-4 transition-all duration-300'
        ]">
          <span class="text-sm font-medium">{{ feedback.message }}</span>
          <button @click="feedback = null" class="text-slate-400 hover:text-slate-200">✕</button>
        </div>

        <!-- Skeleton Loading (Mục 4) -->
        <JobSkeleton v-if="isLoadingJobs" :count="6" />

        <div v-else-if="filteredJobs.length === 0" class="rounded-xl border border-white/10 bg-slate-900/82 px-4 py-12 text-center text-slate-400">
          <p class="font-semibold text-slate-200">Hiện chưa có bài đăng việc làm nào phù hợp.</p>
        </div>

        <div v-else class="mx-auto grid max-w-6xl grid-cols-1 gap-5 md:grid-cols-2 xl:grid-cols-3">
          <div
            v-for="job in paginatedJobs"
            :key="job.id"
            class="flex min-h-[340px] flex-col rounded-lg border border-white/10 bg-slate-900/82 p-5 shadow-lg shadow-slate-950/25 backdrop-blur transition-all duration-200 hover:border-cyan-300/50 hover:shadow-cyan-950/30"
          >
            <div class="flex-1">
              <div class="mb-3 flex items-start justify-between gap-3">
                <span class="inline-flex max-w-[58%] items-center truncate rounded-full bg-cyan-400/10 px-2.5 py-1 text-xs font-semibold text-cyan-200 ring-1 ring-cyan-400/20">
                  {{ companyName(job) }}
                </span>
                <span class="whitespace-nowrap rounded-lg bg-emerald-400/10 px-2.5 py-1 text-sm font-extrabold text-emerald-300">
                  {{ formatCurrency(job.salary) }}
                </span>
              </div>
              <h3 class="min-h-[3.5rem] text-lg font-bold text-white line-clamp-2">{{ displayJobTitle(job.title) }}</h3>
              <p class="mt-1 flex items-center gap-1 text-sm font-medium text-slate-400">
                <svg class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a2 2 0 01-2.828 0l-4.243-4.243a8 8 0 1111.314 0z" />
                </svg>
                {{ job.location }}
              </p>
              <p v-if="job.working_date" class="mt-1 flex items-center gap-1 text-sm font-medium text-slate-400">
                <svg class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 002 2z" />
                </svg>
                {{ job.working_date }}
              </p>
              <p class="mt-4 text-sm leading-6 text-slate-300 line-clamp-4">
                {{ job.description }}
              </p>
            </div>

            <div class="mt-6 border-t border-white/10 pt-4">
              <template v-if="isAuthenticated && userRole === 'student'">
                <button
                  @click="handleApply(job.id)"
                  :disabled="isApplying === job.id"
                  class="focus-ring flex w-full justify-center rounded-lg border border-transparent bg-cyan-400 px-4 py-2 text-sm font-semibold text-slate-950 shadow-sm transition-colors hover:bg-cyan-300 disabled:cursor-not-allowed disabled:opacity-50"
                >
                  {{ isApplying === job.id ? 'Đang ứng tuyển...' : 'Ứng Tuyển Ngay' }}
                </button>
              </template>
              <template v-else-if="isAuthenticated && userRole === 'business'">
                <div class="py-2 text-center text-xs font-medium italic text-slate-400">Đang xem dưới quyền Doanh nghiệp</div>
              </template>
              <template v-else>
                <NuxtLink
                  to="/login"
                  class="flex w-full justify-center rounded-lg border border-white/10 px-4 py-2 text-center text-sm font-semibold text-cyan-200 transition-all hover:bg-cyan-400/10 hover:text-cyan-100"
                >
                  Đăng Nhập Để Ứng Tuyển
                </NuxtLink>
              </template>
            </div>
          </div>
        </div>

        <div class="mx-auto mt-6 max-w-6xl">
          <PaginationControls
            v-if="!isLoadingJobs && filteredJobs.length > 0"
            :page="jobsPage"
            :page-size="jobsPageSize"
            :total-items="filteredJobs.length"
            @update:page="jobsPage = $event"
          />
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useApi } from '~/composables/useApi'
import { useAuth } from '~/composables/useAuth'
import PaginationControls from '~/components/common/PaginationControls.vue'
import JobSkeleton from '~/components/common/JobSkeleton.vue'

defineProps({
  showHero: { type: Boolean, default: true },
  standalone: { type: Boolean, default: true },
  sectionId: { type: String, default: 'explore-jobs' }
})

const api = useApi()
const { isAuthenticated, userRole } = useAuth()

const jobs = ref<any[]>([])
const isLoadingJobs = ref(false)
const isApplying = ref<number | null>(null)
const feedback = ref<{ type: 'success' | 'error'; message: string } | null>(null)
const searchQuery = ref('')
const locationQuery = ref('')
const jobsPage = ref(1)
const jobsPageSize = 6

const filteredJobs = computed(() => {
  return jobs.value.filter(job => {
    const title = (job.title || '').toLowerCase()
    const description = (job.description || '').toLowerCase()
    const location = (job.location || '').toLowerCase()
    const search = searchQuery.value.toLowerCase()
    const locationSearch = locationQuery.value.toLowerCase()

    const matchesSearch = !search || title.includes(search) || description.includes(search)
    const matchesLocation = !locationSearch || location.includes(locationSearch)
    return matchesSearch && matchesLocation
  })
})

const paginatedJobs = computed(() => {
  const start = (jobsPage.value - 1) * jobsPageSize
  return filteredJobs.value.slice(start, start + jobsPageSize)
})

watch(filteredJobs, () => {
  jobsPage.value = 1
})

const fetchJobs = async () => {
  isLoadingJobs.value = true
  try {
    const res = await api.get('/api/jobs')
    jobs.value = res.data || []
  } catch (err) {
    console.error('Error fetching jobs:', err)
  } finally {
    isLoadingJobs.value = false
  }
}

const handleApply = async (jobId: number) => {
  isApplying.value = jobId
  feedback.value = null
  try {
    const res = await api.post('/api/jobs/apply', { job_id: jobId })
    feedback.value = { type: 'success', message: res.message || '🎉 Nộp đơn ứng tuyển thành công!' }
  } catch (err: any) {
    feedback.value = { type: 'error', message: err.response?._data?.message || 'Nộp đơn thất bại. Vui lòng thử lại.' }
  } finally {
    isApplying.value = null
  }
}

const companyName = (job: any) => {
  return job?.business?.company_name || `Doanh nghiệp #${job?.business_id || ''}`
}

const formatCurrency = (value: number | string | null | undefined) => {
  const amount = Number(value || 0)
  return `${amount.toLocaleString('vi-VN')} VNĐ`
}

const displayJobTitle = (title: string | null | undefined) => {
  return (title || 'Công việc chưa đặt tên').replace(/\bMarketting\b/gi, 'Marketing')
}

onMounted(fetchJobs)
</script>

