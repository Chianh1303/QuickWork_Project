<template>
  <!-- Section 4: Employer Applicant Management (B6) -->
  <div v-show="activeSection === 'applicants'" class="space-y-6 max-w-7xl mx-auto">
    <!-- Option B Sub-Tabs Bar: Active Applicants vs Archived / Closed History -->
    <div class="flex flex-wrap items-center justify-between gap-4 border-b border-white/10 pb-4">
      <div class="flex flex-wrap items-center gap-3">
        <button
          @click="applicantTab = 'active'; applicantStatusFilter = 'all'"
          :class="[
            applicantTab === 'active'
              ? 'bg-cyan-400 text-slate-950 font-black shadow-lg shadow-cyan-500/25 ring-2 ring-cyan-400/30'
              : 'bg-slate-900 text-slate-400 hover:text-white font-bold border border-white/10 hover:bg-slate-800/80',
            'px-3.5 py-2.5 sm:px-4 text-xs rounded-xl transition-all cursor-pointer flex items-center gap-2'
          ]"
        >
          <span>⚡ Ứng Viên Đang Xử Lý</span>
          <span :class="[applicantTab === 'active' ? 'bg-slate-950/20 text-slate-950' : 'bg-slate-800 text-cyan-300 border border-cyan-400/20', 'px-2 py-0.5 rounded-full text-[11px] font-black']">
            {{ activeAppsCount }}
          </span>
        </button>

        <button
          @click="applicantTab = 'closed'; applicantStatusFilter = 'all'"
          :class="[
            applicantTab === 'closed'
              ? 'bg-cyan-400 text-slate-950 font-black shadow-lg shadow-cyan-500/25 ring-2 ring-cyan-400/30'
              : 'bg-slate-900 text-slate-400 hover:text-white font-bold border border-white/10 hover:bg-slate-800/80',
            'px-3.5 py-2.5 sm:px-4 text-xs rounded-xl transition-all cursor-pointer flex items-center gap-2'
          ]"
        >
          <span>📜 Lịch Sử & Đơn Đã Đóng</span>
          <span :class="[applicantTab === 'closed' ? 'bg-slate-950/20 text-slate-950' : 'bg-slate-800 text-slate-400 border border-white/10', 'px-2 py-0.5 rounded-full text-[11px] font-black']">
            {{ closedAppsCount }}
          </span>
        </button>
      </div>
    </div>

    <!-- Search & Filter Controls -->
    <div class="bg-slate-900/90 p-4 rounded-2xl border border-cyan-500/20 shadow-xl shadow-cyan-950/30 grid grid-cols-1 sm:grid-cols-2 gap-4 backdrop-blur-xl">
      <div class="relative">
        <span class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-slate-500">
          <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </span>
        <input
          v-model="applicantSearchQuery"
          type="text"
          class="block w-full pl-10 pr-3.5 py-2.5 border border-cyan-500/20 rounded-xl text-xs bg-slate-950/80 placeholder-slate-500 text-white focus:outline-none focus:ring-2 focus:ring-cyan-400 transition-all font-medium"
          placeholder="Tìm kiếm theo tên sinh viên hoặc vị trí việc làm..."
        />
      </div>
      <div>
        <select
          v-model="applicantStatusFilter"
          class="block w-full px-3.5 py-2.5 border border-cyan-500/20 rounded-xl text-xs bg-slate-950/80 text-white focus:outline-none focus:ring-2 focus:ring-cyan-400 transition-all font-semibold cursor-pointer"
        >
          <option value="all">{{ applicantTab === 'active' ? 'Tất cả đơn đang xử lý' : 'Tất cả đơn đã đóng' }}</option>
          <template v-if="applicantTab === 'active'">
            <option value="pending">Mới ứng tuyển (Chờ duyệt)</option>
            <option value="approved">Đã duyệt (Trúng tuyển)</option>
            <option value="offer_accepted">Đã chấp nhận Offer</option>
            <option value="student_completed">Chờ xác nhận hoàn thành</option>
          </template>
          <template v-else>
            <option value="paid">Đã hoàn thành & Giải ngân</option>
            <option value="rejected">Đã từ chối</option>
          </template>
        </select>
      </div>
    </div>

    <!-- Loading state skeleton -->
    <div v-if="isLoadingApps" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="n in 3" :key="n" class="rounded-2xl border border-cyan-500/15 bg-slate-900/90 p-4 sm:p-5 space-y-4 animate-pulse">
        <div class="flex items-center gap-3">
          <div class="h-10 w-10 sm:h-12 sm:w-12 bg-slate-800 rounded-2xl"></div>
          <div class="space-y-2 flex-1">
            <div class="h-4 bg-slate-800 rounded w-3/4"></div>
            <div class="h-3 bg-slate-800 rounded w-1/2"></div>
          </div>
        </div>
        <div class="h-10 bg-slate-800/60 rounded-xl"></div>
      </div>
    </div>

    <!-- Empty state -->
    <div v-else-if="filteredApps.length === 0" class="bg-slate-900/90 text-center py-12 sm:py-16 px-4 rounded-3xl border border-cyan-500/15 shadow-xl backdrop-blur">
      <div class="inline-flex p-3.5 sm:p-4 rounded-2xl bg-cyan-500/10 border border-cyan-500/20 text-cyan-300 mb-3 sm:mb-4">
        <svg class="h-8 w-8 sm:h-10 sm:w-10" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
        </svg>
      </div>
      <h3 class="text-base font-extrabold text-white">Chưa Có Đơn Ứng Tuyển Nào</h3>
      <p class="text-xs text-slate-400 mt-1 max-w-sm mx-auto font-medium">
        Hiện chưa có ứng viên nộp đơn phù hợp với điều kiện lọc tìm kiếm của bạn.
      </p>
    </div>

    <!-- Modern Compact Candidate Grid (Responsive Mobile -> Tablet -> Desktop) -->
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 font-sans">
      <div
        v-for="app in paginatedApplicants"
        :key="app.id"
        @click="handleManagedDetail(app)"
        class="group relative rounded-2xl border border-cyan-500/20 bg-slate-900/90 p-4 sm:p-5 shadow-xl hover:border-cyan-400/50 hover:bg-slate-900 hover:shadow-2xl hover:shadow-cyan-500/10 transition-all duration-300 flex flex-col justify-between cursor-pointer backdrop-blur-xl"
      >
        <div class="space-y-3 sm:space-y-4">
          <!-- Top Row: Student Avatar + Name & Status Badge -->
          <div class="flex items-start justify-between gap-2.5">
            <div class="flex items-center gap-2.5 sm:gap-3 min-w-0 flex-1">
              <div class="relative h-11 w-11 sm:h-12 sm:w-12 flex-shrink-0 overflow-hidden rounded-2xl border border-cyan-500/30 bg-slate-800 shadow-md">
                <img
                  v-if="app.student?.avatar_url"
                  :src="app.student.avatar_url"
                  :alt="app.student?.full_name"
                  class="h-full w-full object-cover"
                />
                <div v-else class="h-full w-full bg-gradient-to-br from-cyan-600 via-blue-600 to-emerald-500 flex items-center justify-center text-xs font-black text-white">
                  {{ app.student?.full_name?.charAt(0) || 'S' }}
                </div>
              </div>

              <div class="min-w-0 flex-1">
                <h4 class="text-sm sm:text-base font-extrabold text-white group-hover:text-cyan-300 transition-colors truncate">
                  {{ app.student?.full_name || 'N/A' }}
                </h4>
                <p class="text-[11px] sm:text-xs font-semibold text-cyan-300 truncate mt-0.5">
                  🎯 {{ app.job?.title || jobTitleLookup(app.job_id) || 'Công việc chưa xác định' }}
                </p>
              </div>
            </div>

            <span
              :class="[
                statusBadgeClass(app.status),
                'inline-flex items-center px-2 py-0.5 sm:px-2.5 sm:py-1 rounded-full text-[10px] sm:text-[11px] font-black border uppercase tracking-wider flex-shrink-0 shadow-sm whitespace-nowrap'
              ]"
            >
              {{ app.status ? app.status.replace('_', ' ') : 'Chờ duyệt' }}
            </span>
          </div>

          <!-- Candidate Meta Info & CV Link -->
          <div class="space-y-1.5 sm:space-y-2 pt-2.5 sm:pt-3 border-t border-cyan-500/10 text-xs">
            <div class="flex items-center justify-between text-slate-300 font-medium">
              <span class="text-slate-400">📞 Điện thoại:</span>
              <span class="font-extrabold text-slate-100">{{ app.student?.phone || 'Chưa cập nhật' }}</span>
            </div>

            <div class="flex items-center justify-between text-slate-300 font-medium">
              <span class="text-slate-400">📅 Ngày nộp:</span>
              <span class="font-bold text-slate-300">{{ formatDate(app.applied_at || app.id) }}</span>
            </div>

            <!-- Skills & CV Link Row -->
            <div class="pt-1.5 sm:pt-2 flex flex-wrap items-center justify-between gap-2">
              <div class="flex flex-wrap gap-1 max-w-[65%]">
                <span
                  v-for="(skill, sIdx) in parseSkills(app.student?.skills).slice(0, 3)"
                  :key="sIdx"
                  class="inline-block bg-cyan-500/10 text-cyan-200 border border-cyan-500/20 text-[10px] font-extrabold px-1.5 py-0.5 rounded-md"
                >
                  {{ skill }}
                </span>
                <span v-if="!app.student?.skills || parseSkills(app.student?.skills).length === 0" class="text-[11px] text-slate-500 italic">
                  Chưa có kỹ năng
                </span>
              </div>

              <a
                v-if="app.student?.cv_url"
                :href="app.student.cv_url"
                target="_blank"
                @click.stop
                class="inline-flex items-center gap-1 text-[10px] sm:text-[11px] font-black text-rose-300 bg-rose-500/15 hover:bg-rose-500/25 px-2 py-0.5 sm:px-2.5 sm:py-1 rounded-lg border border-rose-500/30 transition-all shadow-sm"
              >
                📄 Xem CV
              </a>
            </div>
          </div>
        </div>

        <!-- Bottom Action Buttons: Nhắn tin & Quản lý -->
        <div class="mt-3.5 sm:mt-4 pt-2.5 sm:pt-3 border-t border-cyan-500/10 grid grid-cols-2 gap-2 sm:gap-2.5">
          <button
            type="button"
            @click.stop="handleOpenChat(app)"
            class="py-2 px-2.5 sm:py-2.5 sm:px-3 rounded-xl border border-cyan-500/30 bg-cyan-500/10 hover:bg-cyan-500/20 text-cyan-200 hover:text-white font-extrabold text-xs transition-all flex items-center justify-center gap-1.5 shadow-md cursor-pointer"
          >
            <svg class="h-3.5 w-3.5 sm:h-4 sm:w-4 text-cyan-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
            </svg>
            <span>Nhắn tin</span>
          </button>

          <button
            type="button"
            @click.stop="handleManagedDetail(app)"
            class="py-2 px-2.5 sm:py-2.5 sm:px-3 rounded-xl bg-gradient-to-r from-cyan-400 to-emerald-400 hover:from-cyan-300 hover:to-emerald-300 text-slate-950 font-black text-xs shadow-lg shadow-cyan-500/20 transition-all hover:scale-[1.02] active:scale-98 flex items-center justify-center gap-1.5 cursor-pointer"
          >
            <svg class="h-3.5 w-3.5 sm:h-4 sm:w-4 text-slate-950" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4" />
            </svg>
            <span>Quản lý</span>
          </button>
        </div>
      </div>
    </div>

    <PaginationControls
      v-if="filteredApps.length > 0"
      :page="applicantsPage"
      :page-size="applicantsPageSize"
      :total-items="filteredApps.length"
      @update:page="applicantsPage = $event"
    />
  </div>
</template>

<script setup lang="ts">
import { toRefs } from 'vue'
import PaginationControls from '~/components/common/PaginationControls.vue'

const props = defineProps<{ state: Record<string, any> }>()
const {
  activeSection,
  navItems,
  jobs,
  applications,
  metricsCards,
  fillRatio,
  isEditing,
  profileForm,
  logoPreview,
  onLogoFileChange,
  isSavingProfile,
  handleUpdateProfile,
  showCreateForm,
  jobForm,
  handleCreateJob,
  isCreatingJob,
  isLoadingJobs,
  isLoadingApps,
  applicantSearchQuery,
  applicantStatusFilter,
  applicantTab,
  activeAppsCount,
  closedAppsCount,
  filteredApps,
  paginatedApplicants,
  applicantsPage,
  applicantsPageSize,
  jobTitleLookup,
  formatDate,
  statusBadgeClass,
  parseSkills,
  openChatModal,
  openReviewModal,
  openCompletionModal,
  openBusinessReviewModal,
  openManagedApplicantModal,
  triggerConfirmModal,
  showConfirmModal,
  confirmTarget,
  confirmAction,
  isReviewing,
  handleReviewApplication,
  selectedApp,
  reviewStatus,
  offerForm,
  isSubmitting,
  closeModal,
  submitReview,
  isChatModalOpen,
  selectedChatApp,
  currentBusinessUserId
} = toRefs(props.state)

const handleOpenChat = (app: any) => {
  if (!app) return
  if (typeof openChatModal?.value === 'function') {
    openChatModal.value(app)
  } else if (typeof (openChatModal as any) === 'function') {
    (openChatModal as any)(app)
  }
}

const handleManagedDetail = (app: any) => {
  if (!app) return
  if (typeof openManagedApplicantModal?.value === 'function') {
    openManagedApplicantModal.value(app)
  } else if (typeof (openManagedApplicantModal as any) === 'function') {
    (openManagedApplicantModal as any)(app)
  }
}
</script>
