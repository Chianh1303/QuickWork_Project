<template>
      <!-- Section 1: Dashboard (Explore Jobs) -->
      <div v-if="activeSection === 'jobs'" class="space-y-6">
        <!-- AI Recommended Jobs Section -->
        <StudentAiRecommendedJobs @apply="handleAiApply" />

        <!-- Feedback Banner -->
        <div v-if="feedback" :class="[
          feedback.type === 'success' ? 'bg-emerald-50 border-emerald-300 text-emerald-800' : 'bg-red-50 border-red-300 text-red-800',
          'border-l-4 p-4 rounded-r-lg flex justify-between items-start transition-all duration-300 shadow-sm'
        ]">
          <div class="flex items-center space-x-3">
            <svg v-if="feedback.type === 'success'" class="h-5 w-5 text-emerald-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <svg v-else class="h-5 w-5 text-red-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <span class="text-sm font-semibold">{{ feedback.message }}</span>
          </div>
          <button @click="feedback = null" class="text-slate-400 hover:text-slate-600">
            <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

       <!-- Filters -->
       <div class="bg-slate-900/90 p-4 rounded-2xl border border-indigo-500/20 shadow-xl shadow-indigo-950/30 backdrop-blur-xl">
          <div class="mb-3.5 flex items-center justify-between gap-3">
            <span class="text-xs font-black uppercase tracking-wider text-indigo-300">Bộ Lọc Tìm Kiếm Việc Làm Enterprise</span>
            <span class="inline-flex items-center rounded-full bg-indigo-500/10 px-3 py-1 text-xs font-extrabold text-indigo-200 ring-1 ring-indigo-500/20">
              {{ filteredJobs.length }} công việc khả dụng
            </span>
          </div>
          <form @submit.prevent="fetchJobs" class="space-y-3">
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-3">
              <div class="relative">
                <span class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-slate-500">
                  <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-5.2-5.2m0 0A7.2 7.2 0 105.6 5.6a7.2 7.2 0 0010.2 10.2z" />
                  </svg>
                </span>
                <input v-model="jobsSearchQuery" type="text" class="block w-full pl-10 pr-3.5 py-2.5 border border-indigo-500/20 rounded-xl text-sm bg-slate-950/80 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-indigo-400 focus:border-indigo-400 font-medium" placeholder="Tìm theo vị trí tuyển dụng, từ khóa kỹ năng..." />
              </div>

              <div class="relative">
                <span class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-slate-500">
                  <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21s7-4.35 7-11a7 7 0 10-14 0c0 6.65 7 11 7 11z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10.5h.01" />
                  </svg>
                </span>
                <input v-model="jobsLocationQuery" type="text" class="block w-full pl-10 pr-3.5 py-2.5 border border-indigo-500/20 rounded-xl text-sm bg-slate-950/80 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-indigo-400 focus:border-indigo-400 font-medium" placeholder="Thành phố, địa điểm làm việc..." />
              </div>
            </div>

            <div class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-[1fr_1fr_1fr_1fr_auto] gap-3">
              <div>
                <select v-model="filterCategory" class="block w-full px-3 py-2.5 border border-indigo-500/20 rounded-xl text-xs bg-slate-950/80 text-slate-100 focus:outline-none focus:ring-2 focus:ring-indigo-400 font-medium">
                  <option value="all">Tất cả ngành nghề</option>
                  <option value="it">Công nghệ thông tin</option>
                  <option value="marketing">Marketing / Truyền thông</option>
                  <option value="design">Thiết kế đồ họa</option>
                </select>
              </div>

              <div>
                <select v-model="filterJobType" class="block w-full px-3 py-2.5 border border-indigo-500/20 rounded-xl text-xs bg-slate-950/80 text-slate-100 focus:outline-none focus:ring-2 focus:ring-indigo-400 font-medium">
                  <option value="all">Tất cả hình thức</option>
                  <option value="full-time">Toàn thời gian</option>
                  <option value="part-time">Bán thời gian</option>
                  <option value="intern">Thực tập</option>
                </select>
              </div>

              <div>
                <select v-model="filterMinSalary" class="block w-full px-3 py-2.5 border border-indigo-500/20 rounded-xl text-xs bg-slate-950/80 text-slate-100 focus:outline-none focus:ring-2 focus:ring-indigo-400 font-medium">
                  <option value="">Mức lương bất kỳ</option>
                  <option value="2000000">Từ 2.000.000 VNĐ</option>
                  <option value="5000000">Từ 5.000.000 VNĐ</option>
                  <option value="10000000">Từ 10.000.000 VNĐ</option>
                </select>
              </div>

              <div>
                <select v-model="filterApplyStatus" class="block w-full px-3 py-2.5 border border-indigo-500/20 rounded-xl text-xs bg-slate-950/80 text-slate-100 focus:outline-none focus:ring-2 focus:ring-indigo-400 font-bold">
                  <option value="all">Tất cả (Ưu tiên chưa nộp)</option>
                  <option value="unapplied">⚡ Chỉ việc CHƯA nộp</option>
                  <option value="applied">✓ Việc ĐÃ nộp đơn</option>
                </select>
              </div>

              <div class="grid grid-cols-2 gap-2 sm:col-span-2 xl:col-span-1">
                <button type="button" @click="resetFilters" class="px-3.5 py-2.5 border border-white/10 text-xs font-bold rounded-xl text-slate-300 bg-white/5 hover:bg-white/10 transition-all focus-ring">
                  Xóa lọc
                </button>
                <button type="submit" class="px-4 py-2.5 text-xs font-extrabold rounded-xl text-white bg-gradient-to-r from-indigo-500 via-blue-600 to-emerald-500 hover:from-indigo-400 hover:to-emerald-400 shadow-md shadow-indigo-500/20 transition-all focus-ring">
                  Lọc ngay
                </button>
              </div>
            </div>
          </form>
        </div>

        <!-- Skeleton / Loading -->
        <div v-if="isLoadingJobs" class="grid grid-cols-1 lg:grid-cols-2 gap-4">
          <div v-for="n in 6" :key="n" class="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-3.5 bg-slate-900/90 rounded-2xl border border-indigo-500/10 p-4 shadow-md animate-pulse">
            <div class="flex items-start sm:items-center gap-3.5 min-w-0 flex-1 w-full">
              <div class="h-14 w-14 sm:h-16 sm:w-16 rounded-xl bg-slate-800 flex-shrink-0"></div>
              <div class="min-w-0 flex-1 space-y-2">
                <div class="h-4 w-3/4 rounded bg-slate-800"></div>
                <div class="h-3 w-1/2 rounded bg-slate-800"></div>
                <div class="flex gap-2 pt-1">
                  <div class="h-5 w-20 rounded bg-slate-800"></div>
                  <div class="h-5 w-24 rounded bg-slate-800"></div>
                </div>
              </div>
            </div>
            <div class="w-full sm:w-auto flex justify-end">
              <div class="h-8 w-24 rounded-xl bg-slate-800"></div>
            </div>
          </div>
        </div>

        <!-- Jobs Grid -->
        <div v-else-if="filteredJobs.length === 0" class="bg-slate-900/90 text-center py-16 px-4 rounded-2xl border border-indigo-500/15 text-slate-400">
          <svg class="mx-auto h-12 w-12 text-slate-400 mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <p class="font-bold text-slate-100 text-lg">Không tìm thấy việc làm nào phù hợp</p>
          <p class="text-sm text-slate-400 mt-1">Hãy thử thay đổi từ khóa hoặc xóa bộ lọc tìm kiếm.</p>
        </div>

        <div v-else class="grid grid-cols-1 lg:grid-cols-2 gap-4">
          <div
            v-for="job in paginatedJobs"
            :key="job.id"
            class="group relative flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4 rounded-2xl border border-indigo-500/15 bg-slate-900/90 p-4 shadow-md shadow-indigo-950/20 backdrop-blur-xl transition-all duration-200 hover:border-indigo-400/50 hover:bg-slate-900 hover:shadow-xl hover:shadow-indigo-500/10"
          >
            <!-- Left & Middle: Logo + Job details (Clickable to view details) -->
            <div @click="handleApply(job)" class="flex items-start sm:items-center gap-3.5 min-w-0 flex-1 w-full cursor-pointer">
              <!-- Business Logo -->
              <div class="relative h-14 w-14 sm:h-16 sm:w-16 flex-shrink-0 overflow-hidden rounded-xl border border-indigo-500/20 bg-slate-800 shadow-md flex items-center justify-center group-hover:scale-105 transition-transform">
                <img
                  v-if="job.business?.logo_url"
                  :src="getMediaUrl(job.business.logo_url)"
                  :alt="companyNameLookup(job)"
                  class="h-full w-full object-cover relative z-10"
                  @error="handleImgError"
                />
                <div class="absolute inset-0 flex items-center justify-center bg-gradient-to-br from-indigo-600 via-blue-600 to-emerald-500 text-sm font-black text-white">
                  {{ getCompanyInitial(companyNameLookup(job)) }}
                </div>
              </div>

              <!-- Main Details -->
              <div class="min-w-0 flex-1 space-y-1">
                <h3 class="text-sm sm:text-base font-extrabold text-white group-hover:text-indigo-300 transition-colors truncate" :title="displayJobTitle(job.title)">
                  {{ displayJobTitle(job.title) }}
                </h3>

                <p class="text-xs font-semibold text-slate-300 truncate" :title="companyNameLookup(job)">
                  {{ companyNameLookup(job) }}
                </p>

                <!-- Tags / Badges Row -->
                <div class="flex flex-wrap items-center gap-1.5 pt-0.5">
                  <span class="inline-flex items-center text-xs font-black text-emerald-300 bg-emerald-500/10 border border-emerald-500/20 px-2 py-0.5 rounded-md whitespace-nowrap">
                    {{ formatCurrency(job.salary) }}
                  </span>

                  <span v-if="job.location" class="inline-flex items-center gap-1 text-[11px] font-medium text-slate-400 bg-slate-800/80 border border-slate-700/60 px-2 py-0.5 rounded-md truncate max-w-[130px]" :title="job.location">
                    <svg class="h-3 w-3 flex-shrink-0 text-indigo-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.243-4.243a8 8 0 1111.314 0z" />
                    </svg>
                    <span class="truncate">{{ job.location }}</span>
                  </span>

                  <span v-if="job.working_date" class="hidden sm:inline-flex items-center gap-1 text-[11px] font-medium text-slate-400 bg-slate-800/80 border border-slate-700/60 px-2 py-0.5 rounded-md truncate max-w-[130px]" :title="job.working_date">
                    <svg class="h-3 w-3 flex-shrink-0 text-indigo-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                    </svg>
                    <span class="truncate">{{ job.working_date }}</span>
                  </span>
                </div>
              </div>
            </div>

            <!-- Right Side: Bookmark, Share & Action Button -->
            <div class="w-full sm:w-auto flex-shrink-0 pt-2 sm:pt-0 border-t sm:border-t-0 border-indigo-500/10 flex items-center justify-between sm:justify-end gap-2">
              <div class="flex items-center gap-1.5">
                <!-- Bookmark Heart Button -->
                <button
                  type="button"
                  @click.stop="toggleSaveJob(job)"
                  :title="isJobSaved(job.id) ? 'Bỏ lưu việc làm' : 'Lưu vào mục yêu thích'"
                  class="p-2 rounded-xl border border-indigo-500/20 bg-slate-950/60 hover:bg-slate-950 transition-colors"
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
                  @click.stop="shareJob(job)"
                  title="Sao chép liên kết chia sẻ"
                  class="p-2 rounded-xl border border-indigo-500/20 bg-slate-950/60 hover:bg-slate-950 text-slate-400 hover:text-indigo-300 transition-colors"
                >
                  <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
                  </svg>
                </button>
              </div>

              <button
                :disabled="checkIfApplied(job.id)"
                @click="handleApply(job)"
                :class="checkIfApplied(job.id) ? 'bg-slate-800 text-slate-400 border border-slate-700/60' : 'bg-gradient-to-r from-indigo-500 to-emerald-500 hover:from-indigo-400 hover:to-emerald-400 text-white shadow-md shadow-indigo-500/20'"
                class="flex-1 sm:flex-initial flex items-center justify-center gap-1.5 py-2 px-4 rounded-xl text-xs font-extrabold transition-all focus-ring disabled:cursor-not-allowed whitespace-nowrap"
              >
                <span v-if="isApplying === job.id" class="flex items-center space-x-2">
                  <svg class="animate-spin h-3.5 w-3.5 text-white" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  <span>Đang nộp...</span>
                </span>
                <span v-else class="flex items-center justify-center gap-1.5">
                  <span>{{ checkIfApplied(job.id) ? '✓ Đã Nộp' : 'Ứng tuyển' }}</span>
                  <svg v-if="!checkIfApplied(job.id)" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
                  </svg>
                </span>
              </button>
            </div>
          </div>
        </div>

        <PaginationControls
          v-if="!isLoadingJobs && filteredJobs.length > 0"
          :page="jobsPage"
          :page-size="jobsPageSize"
          :total-items="filteredJobs.length"
          @update:page="jobsPage = $event"
        />
      </div>
</template>

<script setup lang="ts">
import { toRefs } from 'vue'
import { useMedia } from '~/composables/useMedia'
import { useSavedJobs } from '~/composables/useSavedJobs'
import PaginationControls from '~/components/common/PaginationControls.vue'
import StudentAiRecommendedJobs from '~/components/student/StudentAiRecommendedJobs.vue'

const { getMediaUrl, getCompanyInitial } = useMedia()
const { savedJobIds, isJobSaved, toggleSaveJob, shareJob } = useSavedJobs()

const handleImgError = (event: Event) => {
  const target = event.target as HTMLImageElement
  if (target) {
    target.style.display = 'none'
  }
}

const props = defineProps<{ state: Record<string, any> }>()
const {
  activeSection,
  navItems,
  feedback,
  jobsSearchQuery,
  jobsLocationQuery,
  filterCategory,
  filterJobType,
  filterMinSalary,
  filterApplyStatus,
  resetFilters,
  fetchJobs,
  isLoadingJobs,
  filteredJobs,
  paginatedJobs,
  jobsPage,
  jobsPageSize,
  companyNameLookup,
  checkIfApplied,
  handleApply,
  isApplying,
  isEditing,
  profileForm,
  skillsArray,
  avatarPreview,
  onAvatarFileChange,
  skillsText,
  onCvFileChange,
  isSavingProfile,
  handleUpdateProfile,
  isLoadingApps,
  filteredApps,
  appSearchQuery,
  appStatusFilter,
  openChatModal,
  openOfferModal,
  triggerCancelConfirm,
  isCancellingApp,
  formatDate,
  statusBadgeClass,
  isWorking,
  getTimer,
  handleCheckIn,
  handleCheckOut,
  selectedJobForApply,
  coverNoteText,
  submitApplication,
  isSubmittingApply,
  appIdToCancel,
  confirmCancelApplication,
  selectedOffer,
  currentUserId,
  handleOfferResponse,
  isResponding,
  isChatModalOpen,
  selectedChatApp,
  toast
} = toRefs(props.state)

const formatCurrency = (value: number | string | null | undefined) => {
  const amount = Number(value || 0)
  return `${amount.toLocaleString('vi-VN')} VND`
}

const displayJobTitle = (title: string | null | undefined) => {
  return (title || 'Untitled Job').replace(/\bMarketting\b/gi, 'Marketing')
}

const handleAiApply = (aiJob: any) => {
  if (!aiJob) return
  const mappedJob = {
    id: aiJob.job_id,
    title: aiJob.job_title,
    company: aiJob.company,
    description: aiJob.description,
    salary: aiJob.salary,
    location: aiJob.location,
    ...aiJob
  }
  if (typeof handleApply.value === 'function') {
    handleApply.value(mappedJob)
  }
}
</script>
