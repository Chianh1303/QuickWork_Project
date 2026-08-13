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
       <div class="bg-slate-900/82 p-3 rounded-xl border border-cyan-400/15 shadow-lg shadow-slate-950/30 backdrop-blur">
          <div class="mb-3 flex items-center justify-between gap-3">
            <span class="text-xs font-bold uppercase tracking-wide text-cyan-200/80">Bộ lọc việc làm</span>
            <span class="inline-flex items-center rounded-full bg-white/10 px-3 py-1 text-xs font-bold text-cyan-100 ring-1 ring-white/10">
              {{ filteredJobs.length }} kết quả
            </span>
          </div>
          <form @submit.prevent="fetchJobs" class="space-y-3">
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-3">
              <div class="relative">
                <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-500">
                  <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-5.2-5.2m0 0A7.2 7.2 0 105.6 5.6a7.2 7.2 0 0010.2 10.2z" />
                  </svg>
                </span>
                <input v-model="jobsSearchQuery" type="text" class="block w-full pl-9 pr-3 py-2.5 border border-white/10 rounded-lg text-sm bg-slate-950/70 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 focus:border-cyan-300" placeholder="Tìm theo tiêu đề, từ khóa hoặc kỹ năng..." />
              </div>

              <div class="relative">
                <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-500">
                  <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21s7-4.35 7-11a7 7 0 10-14 0c0 6.65 7 11 7 11z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10.5h.01" />
                  </svg>
                </span>
                <input v-model="jobsLocationQuery" type="text" class="block w-full pl-9 pr-3 py-2.5 border border-white/10 rounded-lg text-sm bg-slate-950/70 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 focus:border-cyan-300" placeholder="Nhập địa điểm, thành phố..." />
              </div>
            </div>

            <div class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-[1fr_1fr_1fr_auto] gap-3">
              <div>
                <select v-model="filterCategory" class="block w-full px-3 py-2.5 border border-white/10 rounded-lg text-sm bg-slate-950/70 text-slate-100 focus:outline-none focus:ring-2 focus:ring-cyan-400 focus:border-cyan-300">
                  <option value="all">Tất cả ngành nghề</option>
                  <option value="it">Công nghệ thông tin</option>
                  <option value="marketing">Marketing / Truyền thông</option>
                  <option value="design">Thiết kế đồ họa</option>
                </select>
              </div>

              <div>
                <select v-model="filterJobType" class="block w-full px-3 py-2.5 border border-white/10 rounded-lg text-sm bg-slate-950/70 text-slate-100 focus:outline-none focus:ring-2 focus:ring-cyan-400 focus:border-cyan-300">
                  <option value="all">Tất cả hình thức</option>
                  <option value="full-time">Toàn thời gian</option>
                  <option value="part-time">Bán thời gian</option>
                  <option value="intern">Thực tập</option>
                </select>
              </div>

              <div>
                <select v-model="filterMinSalary" class="block w-full px-3 py-2.5 border border-white/10 rounded-lg text-sm bg-slate-950/70 text-slate-100 focus:outline-none focus:ring-2 focus:ring-cyan-400 focus:border-cyan-300">
                  <option value="">Mức lương bất kỳ</option>
                  <option value="2000000">Từ 2.000.000 VNĐ</option>
                  <option value="5000000">Từ 5.000.000 VNĐ</option>
                  <option value="10000000">Từ 10.000.000 VNĐ</option>
                </select>
              </div>

              <div class="grid grid-cols-2 gap-2 sm:col-span-2 xl:col-span-1">
                <button type="button" @click="resetFilters" class="px-4 py-2.5 border border-white/10 text-sm font-semibold rounded-lg text-slate-200 bg-white/10 hover:bg-white/15 transition-all focus-ring">
                  Xóa lọc
                </button>
                <button type="submit" class="px-5 py-2.5 border border-transparent text-sm font-semibold rounded-lg text-slate-950 bg-cyan-400 hover:bg-cyan-300 shadow-sm transition-all focus-ring">
                  Lọc ngay
                </button>
              </div>
            </div>

          </form>
        </div>
        <!-- Skeleton / Loading -->
        <div v-if="isLoadingJobs" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div v-for="n in 6" :key="n" class="bg-slate-900/80 rounded-xl border border-white/10 p-6 space-y-4 animate-pulse">
            <div class="h-4 bg-slate-200 rounded w-2/3"></div>
            <div class="h-3 bg-slate-200 rounded w-1/2"></div>
            <div class="h-16 bg-slate-100 rounded"></div>
            <div class="h-8 bg-slate-200 rounded"></div>
          </div>
        </div>

        <!-- Jobs Grid -->
        <div v-else-if="filteredJobs.length === 0" class="bg-slate-900/82 text-center py-16 px-4 rounded-xl border border-white/10 shadow-lg shadow-slate-950/30 text-slate-400 backdrop-blur">
          <svg class="mx-auto h-12 w-12 text-slate-300 mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <p class="font-bold text-slate-100 text-lg">Không tìm thấy việc làm nào</p>
          <p class="text-sm text-slate-400 mt-1">Hãy thử thay đổi từ khóa hoặc xóa bộ lọc tìm kiếm.</p>
        </div>

        <div v-else class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-5">
          <div
            v-for="job in paginatedJobs"
            :key="job.id"
            class="bg-slate-900/82 rounded-lg border border-white/10 shadow-lg shadow-slate-950/25 p-5 flex flex-col min-h-[340px] hover:border-cyan-300/50 hover:shadow-cyan-950/30 transition-all duration-200 backdrop-blur"
          >
            <div class="flex-1">
              <div class="flex justify-between items-start mb-3">
                <span class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold bg-cyan-400/10 text-cyan-200 ring-1 ring-cyan-400/20 max-w-[58%] truncate">
                  {{ companyNameLookup(job) }}
                </span>
                <span class="text-sm font-extrabold text-emerald-300 bg-emerald-400/10 px-2.5 py-1 rounded-lg whitespace-nowrap">
                  {{ formatCurrency(job.salary) }}
                </span>
              </div>
              <h3 class="text-lg font-bold text-white line-clamp-2 min-h-[3.5rem]">{{ displayJobTitle(job.title) }}</h3>
              <p class="text-sm font-medium text-slate-400 mt-1 flex items-center gap-1">
                <svg class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                </svg>
                {{ job.location || 'Chưa cập nhật địa điểm' }}
              </p>
              <p class="text-sm font-medium text-slate-400 mt-1 flex items-center gap-1" v-if="job.working_date">
                <svg class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
                {{ job.working_date }}
              </p>
              <p class="text-sm text-slate-300 mt-4 line-clamp-4 leading-6 font-medium">
                {{ job.description || 'Chưa có mô tả công việc.' }}
              </p>
            </div>

            <div class="mt-5 pt-4 border-t border-white/10">
              <button
                :disabled="checkIfApplied(job.id)"
                @click="handleApply(job)"
                :class="checkIfApplied(job.id) ? 'bg-slate-800 text-slate-400 border-white/10' : 'bg-cyan-400 hover:bg-cyan-300 text-slate-950 border-transparent'"
                class="w-full flex justify-center py-2.5 px-4 border text-sm font-semibold rounded-lg transition-colors focus-ring disabled:cursor-not-allowed shadow-sm"
              >
                <span v-if="isApplying === job.id" class="flex items-center space-x-2">
                  <svg class="animate-spin h-4 w-4 text-slate-950" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  <span>Đang ứng tuyển...</span>
                </span>
                <span v-else>
                  {{ checkIfApplied(job.id) ? 'Đã nộp đơn' : 'Ứng Tuyển Ngay' }}
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
import PaginationControls from '~/components/common/PaginationControls.vue'
import StudentAiRecommendedJobs from '~/components/student/StudentAiRecommendedJobs.vue'

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
