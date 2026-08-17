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

        <!-- Quick Pill Filters -->
        <div class="mx-auto mb-6 flex max-w-5xl flex-wrap items-center gap-2">
          <button
            v-for="pill in pillFilters"
            :key="pill.id"
            @click="activePill = pill.id"
            :class="[
              activePill === pill.id
                ? 'bg-cyan-400 text-slate-950 font-extrabold shadow-md shadow-cyan-500/20'
                : 'bg-slate-900/90 text-slate-300 hover:bg-slate-800 hover:text-white border border-white/10 font-semibold',
              'inline-flex items-center gap-1.5 rounded-full px-4 py-2 text-xs transition-all'
            ]"
          >
            <span>{{ pill.icon }}</span>
            <span>{{ pill.label }}</span>
            <span v-if="pill.id === 'saved' && savedJobIds.length > 0" class="rounded-full bg-rose-500 px-1.5 py-0.2 text-[10px] font-black text-white ml-0.5">
              {{ savedJobIds.length }}
            </span>
          </button>
        </div>

        <div class="mx-auto mb-10 grid max-w-5xl grid-cols-1 gap-4 rounded-2xl border border-cyan-400/15 bg-slate-900/82 p-4 shadow-lg shadow-slate-950/30 backdrop-blur sm:grid-cols-2">
          <div class="relative">
            <span class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3 text-slate-500">
              <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </span>
            <input
              v-model="searchQuery"
              type="text"
              class="block w-full rounded-xl border border-white/10 bg-slate-950/70 py-2.5 pl-10 pr-3 text-sm text-slate-100 placeholder-slate-500 transition-all duration-200 focus:border-cyan-300 focus:outline-none focus:ring-2 focus:ring-cyan-400"
              placeholder="Tìm theo tiêu đề hoặc từ khóa công việc..."
            />
          </div>
          <div class="relative">
            <span class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3 text-slate-500">
              <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.828 0l-4.243-4.243a8 8 0 1111.314 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
            </span>
            <input
              v-model="locationQuery"
              type="text"
              class="block w-full rounded-xl border border-white/10 bg-slate-950/70 py-2.5 pl-10 pr-3 text-sm text-slate-100 placeholder-slate-500 transition-all duration-200 focus:border-cyan-300 focus:outline-none focus:ring-2 focus:ring-cyan-400"
              placeholder="Nhập địa điểm, thành phố..."
            />
          </div>
        </div>

        <div v-if="feedback" :class="[
          feedback.type === 'success' ? 'border-emerald-300 bg-emerald-500/10 text-emerald-300' : 'border-rose-300 bg-rose-500/10 text-rose-300',
          'mx-auto mb-8 flex max-w-4xl items-start justify-between rounded-r-xl border-l-4 p-4 transition-all duration-300'
        ]">
          <span class="text-sm font-medium">{{ feedback.message }}</span>
          <button @click="feedback = null" class="text-slate-400 hover:text-slate-200">✕</button>
        </div>

        <!-- Skeleton Loading -->
        <JobSkeleton v-if="isLoadingJobs" :count="6" />

        <div v-else-if="filteredJobs.length === 0" class="rounded-2xl border border-white/10 bg-slate-900/82 px-4 py-16 text-center text-slate-400">
          <span class="text-4xl block mb-2">🔍</span>
          <p class="font-bold text-slate-100 text-lg">Hiện chưa có bài đăng việc làm nào phù hợp.</p>
          <p class="text-sm text-slate-400 mt-1">Hãy thử thay đổi từ khóa tìm kiếm hoặc chọn bộ lọc khác.</p>
        </div>

        <div v-else class="mx-auto grid max-w-6xl grid-cols-1 lg:grid-cols-2 gap-4">
          <div
            v-for="job in paginatedJobs"
            :key="job.id"
            class="group relative flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4 rounded-2xl border border-white/10 bg-slate-900/85 p-4 shadow-md shadow-slate-950/20 backdrop-blur transition-all duration-200 hover:border-cyan-400/50 hover:bg-slate-900 hover:shadow-xl hover:shadow-cyan-950/30"
          >
            <!-- Left & Middle: Logo + Job details (Clickable to open Detail Modal) -->
            <div @click="openJobDetail(job)" class="flex items-start sm:items-center gap-3.5 min-w-0 flex-1 w-full cursor-pointer">
              <!-- Business Logo -->
              <div class="relative h-14 w-14 sm:h-16 sm:w-16 flex-shrink-0 overflow-hidden rounded-xl border border-white/10 bg-slate-800 shadow-md flex items-center justify-center group-hover:scale-105 transition-transform">
                <img
                  v-if="job.business?.logo_url"
                  :src="getMediaUrl(job.business.logo_url)"
                  :alt="companyName(job)"
                  class="h-full w-full object-cover relative z-10"
                  @error="handleImgError"
                />
                <div class="absolute inset-0 flex items-center justify-center bg-gradient-to-br from-indigo-600 via-blue-600 to-cyan-500 text-sm font-black text-white">
                  {{ getCompanyInitial(companyName(job)) }}
                </div>
              </div>

              <!-- Main Details -->
              <div class="min-w-0 flex-1 space-y-1">
                <h3 class="text-sm sm:text-base font-extrabold text-white group-hover:text-cyan-300 transition-colors truncate" :title="displayJobTitle(job.title)">
                  {{ displayJobTitle(job.title) }}
                </h3>

                <p class="text-xs font-semibold text-slate-300 truncate" :title="companyName(job)">
                  {{ companyName(job) }}
                </p>

                <!-- Tags / Badges Row -->
                <div class="flex flex-wrap items-center gap-1.5 pt-0.5">
                  <span class="inline-flex items-center text-xs font-black text-emerald-300 bg-emerald-500/10 border border-emerald-500/20 px-2 py-0.5 rounded-md whitespace-nowrap">
                    {{ formatCurrency(job.salary) }}
                  </span>

                  <span v-if="job.location" class="inline-flex items-center gap-1 text-[11px] font-medium text-slate-400 bg-slate-800/80 border border-slate-700/60 px-2 py-0.5 rounded-md truncate max-w-[130px]" :title="job.location">
                    <svg class="h-3 w-3 flex-shrink-0 text-cyan-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.828 0l-4.243-4.243a8 8 0 1111.314 0z" />
                    </svg>
                    <span class="truncate">{{ job.location }}</span>
                  </span>

                  <span v-if="job.working_date" class="hidden sm:inline-flex items-center gap-1 text-[11px] font-medium text-slate-400 bg-slate-800/80 border border-slate-700/60 px-2 py-0.5 rounded-md truncate max-w-[130px]" :title="job.working_date">
                    <svg class="h-3 w-3 flex-shrink-0 text-cyan-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                    </svg>
                    <span class="truncate">{{ job.working_date }}</span>
                  </span>
                </div>
              </div>
            </div>

            <!-- Right Side: Bookmark, Share & Action Button -->
            <div class="w-full sm:w-auto flex-shrink-0 pt-2 sm:pt-0 border-t sm:border-t-0 border-white/5 flex items-center justify-between sm:justify-end gap-2">
              <div class="flex items-center gap-1.5">
                <!-- Bookmark Heart Button -->
                <button
                  type="button"
                  @click="toggleSaveJob(job)"
                  :title="isJobSaved(job.id) ? 'Bỏ lưu việc làm' : 'Lưu vào mục yêu thích'"
                  class="p-2 rounded-xl border border-white/10 bg-white/5 hover:bg-white/10 transition-colors"
                >
                  <svg
                    class="h-4 w-4 transition-transform active:scale-125"
                    :class="isJobSaved(job.id) ? 'text-rose-500 fill-rose-500' : 'text-slate-400 hover:text-rose-400'"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
                  </svg>
                </button>

                <!-- Share Link Button -->
                <button
                  type="button"
                  @click="shareJob(job)"
                  title="Sao chép liên kết chia sẻ"
                  class="p-2 rounded-xl border border-white/10 bg-white/5 hover:bg-white/10 text-slate-400 hover:text-cyan-300 transition-colors"
                >
                  <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
                  </svg>
                </button>
              </div>

              <template v-if="isAuthenticated && userRole === 'student'">
                <button
                  @click="handleApply(job.id)"
                  :disabled="isApplying === job.id"
                  class="focus-ring flex items-center justify-center gap-1.5 rounded-xl bg-cyan-400 px-4 py-2 text-xs font-extrabold text-slate-950 shadow-sm transition-all hover:bg-cyan-300 disabled:cursor-not-allowed disabled:opacity-50 whitespace-nowrap"
                >
                  <span>{{ isApplying === job.id ? 'Đang nộp...' : 'Ứng tuyển' }}</span>
                  <svg v-if="isApplying !== job.id" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
                  </svg>
                </button>
              </template>
              <template v-else-if="isAuthenticated && userRole === 'business'">
                <span class="text-[11px] font-medium italic text-slate-400">Doanh nghiệp</span>
              </template>
              <template v-else>
                <NuxtLink
                  to="/login"
                  class="flex items-center justify-center rounded-xl border border-white/10 px-4 py-2 text-center text-xs font-bold text-cyan-200 transition-all hover:bg-cyan-400/10 hover:text-cyan-100 whitespace-nowrap"
                >
                  Ứng tuyển
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

    <!-- Modal Chi tiết Công việc (Job Details Modal) -->
    <div v-if="selectedJobDetail" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
      <div class="bg-slate-900/95 rounded-3xl border border-cyan-400/20 shadow-2xl shadow-slate-950 max-w-lg w-full overflow-hidden animate-in fade-in zoom-in-95 duration-200 flex flex-col max-h-[90vh]">
        <!-- Header -->
        <div class="p-6 pb-4 border-b border-white/10 flex items-start justify-between gap-3 bg-slate-950/60">
          <div class="flex items-start gap-3.5">
            <div class="relative h-14 w-14 flex-shrink-0 overflow-hidden rounded-xl border border-white/10 bg-slate-800 shadow-md">
              <img
                v-if="selectedJobDetail.business?.logo_url"
                :src="getMediaUrl(selectedJobDetail.business.logo_url)"
                :alt="companyName(selectedJobDetail)"
                class="h-full w-full object-cover relative z-10"
                @error="handleImgError"
              />
              <div class="absolute inset-0 flex items-center justify-center bg-gradient-to-br from-indigo-600 to-cyan-500 text-sm font-black text-white">
                {{ getCompanyInitial(companyName(selectedJobDetail)) }}
              </div>
            </div>
            <div>
              <span class="inline-flex items-center px-2.5 py-0.5 rounded-md text-xs font-bold bg-cyan-400/10 text-cyan-300 border border-cyan-400/20">
                {{ companyName(selectedJobDetail) }}
              </span>
              <h3 class="text-base font-extrabold text-white mt-1 leading-snug">
                {{ displayJobTitle(selectedJobDetail.title) }}
              </h3>
            </div>
          </div>
          <button @click="selectedJobDetail = null" class="text-slate-400 hover:text-white p-1.5 rounded-xl hover:bg-white/10 flex-shrink-0">
            ✕
          </button>
        </div>

        <!-- Body -->
        <div class="p-6 overflow-y-auto space-y-4 flex-1">
          <div class="grid grid-cols-2 gap-3 p-3.5 rounded-2xl bg-slate-950/70 border border-white/10 text-xs">
            <div>
              <p class="text-[10px] font-black uppercase text-slate-400">Mức lương</p>
              <p class="text-emerald-300 font-extrabold text-sm mt-0.5">{{ formatCurrency(selectedJobDetail.salary) }}</p>
            </div>
            <div>
              <p class="text-[10px] font-black uppercase text-slate-400">Địa điểm</p>
              <p class="text-slate-200 font-bold mt-0.5 truncate">{{ selectedJobDetail.location || 'Toàn quốc' }}</p>
            </div>
            <div v-if="selectedJobDetail.working_date" class="col-span-2 pt-2 border-t border-white/5">
              <p class="text-[10px] font-black uppercase text-slate-400">Thời gian làm việc</p>
              <p class="text-slate-200 font-medium mt-0.5">{{ selectedJobDetail.working_date }}</p>
            </div>
          </div>

          <div class="rounded-2xl border border-white/10 bg-slate-950/70 p-4 space-y-2">
            <h4 class="text-xs font-black uppercase tracking-wider text-cyan-300">Mô tả công việc chi tiết</h4>
            <p class="text-xs font-medium text-slate-300 leading-relaxed whitespace-pre-line">
              {{ selectedJobDetail.description || 'Chưa có thông tin mô tả chi tiết công việc.' }}
            </p>
          </div>
        </div>

        <!-- Footer -->
        <div class="flex items-center gap-3 justify-end p-4 border-t border-white/10 bg-slate-950/60">
          <button
            type="button"
            @click="toggleSaveJob(selectedJobDetail)"
            class="px-3.5 py-2.5 rounded-xl border border-white/10 text-xs font-bold text-slate-300 hover:bg-white/10 flex items-center gap-1.5"
          >
            <span>{{ isJobSaved(selectedJobDetail.id) ? '❤️ Đã lưu' : '🤍 Lưu tin' }}</span>
          </button>
          <button
            type="button"
            @click="handleApply(selectedJobDetail.id); selectedJobDetail = null"
            class="px-5 py-2.5 rounded-xl text-xs font-extrabold text-slate-950 bg-cyan-400 hover:bg-cyan-300 shadow-md shadow-cyan-500/25 transition-all"
          >
            Ứng tuyển ngay
          </button>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useApi } from '~/composables/useApi'
import { useAuth } from '~/composables/useAuth'
import { useMedia } from '~/composables/useMedia'
import { useSavedJobs } from '~/composables/useSavedJobs'
import PaginationControls from '~/components/common/PaginationControls.vue'
import JobSkeleton from '~/components/common/JobSkeleton.vue'

defineProps({
  showHero: { type: Boolean, default: true },
  standalone: { type: Boolean, default: true },
  sectionId: { type: String, default: 'explore-jobs' }
})

const api = useApi()
const { isAuthenticated, userRole } = useAuth()
const { getMediaUrl, getCompanyInitial } = useMedia()
const { savedJobIds, isJobSaved, toggleSaveJob, shareJob } = useSavedJobs()

const handleImgError = (event: Event) => {
  const target = event.target as HTMLImageElement
  if (target) {
    target.style.display = 'none'
  }
}

const jobs = ref<any[]>([])
const isLoadingJobs = ref(false)
const isApplying = ref<number | null>(null)
const feedback = ref<{ type: 'success' | 'error'; message: string } | null>(null)
const searchQuery = ref('')
const locationQuery = ref('')
const activePill = ref('all')
const selectedJobDetail = ref<any | null>(null)
const jobsPage = ref(1)
const jobsPageSize = 6

const pillFilters = [
  { id: 'all', label: 'Tất cả việc làm', icon: '⚡' },
  { id: 'high-salary', label: 'Lương > 5 Triệu', icon: '💰' },
  { id: 'it', label: 'IT & Phần mềm', icon: '💻' },
  { id: 'marketing', label: 'Marketing / Media', icon: '📈' },
  { id: 'design', label: 'Thiết kế', icon: '🎨' },
  { id: 'saved', label: 'Việc đã lưu', icon: '❤️' }
]

const openJobDetail = (job: any) => {
  selectedJobDetail.value = job
}

const filteredJobs = computed(() => {
  return jobs.value.filter(job => {
    const title = (job.title || '').toLowerCase()
    const description = (job.description || '').toLowerCase()
    const location = (job.location || '').toLowerCase()
    const category = (job.category || '').toLowerCase()
    const salary = Number(job.salary || 0)
    const search = searchQuery.value.toLowerCase()
    const locationSearch = locationQuery.value.toLowerCase()

    const matchesSearch = !search || title.includes(search) || description.includes(search)
    const matchesLocation = !locationSearch || location.includes(locationSearch)

    let matchesPill = true
    if (activePill.value === 'high-salary') {
      matchesPill = salary >= 5000000
    } else if (activePill.value === 'it') {
      matchesPill = category.includes('it') || title.includes('developer') || title.includes('backend') || title.includes('frontend') || title.includes('software')
    } else if (activePill.value === 'marketing') {
      matchesPill = category.includes('marketing') || title.includes('marketing') || title.includes('content') || title.includes('media')
    } else if (activePill.value === 'design') {
      matchesPill = category.includes('design') || title.includes('thiết kế') || title.includes('design') || title.includes('ui')
    } else if (activePill.value === 'saved') {
      matchesPill = savedJobIds.value.includes(job.id)
    }

    return matchesSearch && matchesLocation && matchesPill
  })
})

const paginatedJobs = computed(() => {
  const start = (jobsPage.value - 1) * jobsPageSize
  return filteredJobs.value.slice(start, start + jobsPageSize)
})

watch([filteredJobs, activePill], () => {
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

