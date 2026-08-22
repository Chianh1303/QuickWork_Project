<template>
  <!-- Section 4: Employer Applicant Management (B6) -->
  <div v-show="activeSection === 'applicants'" class="space-y-6">
    <!-- Search & Filter Controls -->
    <div class="bg-white/90 p-4 rounded-xl border border-cyan-100 shadow-sm shadow-slate-950/5 grid grid-cols-1 sm:grid-cols-2 gap-4 backdrop-blur">
      <div class="relative">
        <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </span>
        <input
          v-model="applicantSearchQuery"
          type="text"
          class="block w-full pl-10 pr-3 py-2.5 border border-slate-200 rounded-lg text-sm bg-white placeholder-slate-400 text-slate-900 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-brand-300 transition-all duration-200"
          placeholder="Tìm kiếm theo tên sinh viên hoặc vị trí việc làm..."
        />
      </div>
      <div>
        <select
          v-model="applicantStatusFilter"
          class="block w-full px-3 py-2.5 border border-slate-200 rounded-lg text-sm bg-white text-slate-900 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-brand-300 transition-all duration-200 font-medium"
        >
          <option value="all">Tất cả trạng thái</option>
          <option value="pending">Mới ứng tuyển (Chờ duyệt)</option>
          <option value="approved">Đã duyệt (Trúng tuyển)</option>
          <option value="offer_accepted">Đã chấp nhận Offer</option>
          <option value="student_completed">Chờ Doanh nghiệp xác nhận hoàn thành</option>
          <option value="paid">Đã hoàn thành & Giải ngân</option>
          <option value="rejected">Đã từ chối</option>
        </select>
      </div>
    </div>

    <!-- Loading state skeleton -->
    <div v-if="isLoadingApps" class="bg-white/90 rounded-xl border border-slate-200 p-6 space-y-4 animate-pulse">
      <div class="h-6 bg-slate-200 rounded w-1/4"></div>
      <div class="space-y-3">
        <div class="h-10 bg-slate-100 rounded"></div>
        <div class="h-10 bg-slate-100 rounded"></div>
        <div class="h-10 bg-slate-100 rounded"></div>
      </div>
    </div>

    <!-- Empty state -->
    <div v-else-if="filteredApps.length === 0" class="bg-white/90 text-center py-16 px-4 rounded-xl border border-slate-200 shadow-sm backdrop-blur">
      <div class="inline-flex p-4 rounded-full bg-slate-50 border border-slate-100 text-slate-400 mb-4">
        <svg class="h-12 w-12" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
        </svg>
      </div>
      <h3 class="text-lg font-bold text-slate-900">Chưa Có Đơn Ứng Tuyển Nào</h3>
      <p class="text-sm text-slate-500 mt-1 max-w-sm mx-auto font-medium">
        Hiện chưa có ứng viên nộp đơn phù hợp với điều kiện lọc tìm kiếm của bạn.
      </p>
    </div>

    <div v-else class="hidden lg:block overflow-hidden rounded-xl border border-white/10 bg-slate-900/85 shadow-lg shadow-slate-950/25 backdrop-blur">
      <table class="min-w-full table-fixed divide-y divide-white/10">
        <thead class="bg-slate-950">
          <tr>
            <th class="w-[27%] px-5 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Ứng viên</th>
            <th class="w-[23%] px-5 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Vị trí ứng tuyển</th>
            <th class="w-[14%] px-5 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Số điện thoại</th>
            <th class="w-[13%] px-5 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Ngày nộp đơn</th>
            <th class="w-[12%] px-5 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Trạng thái</th>
            <th class="w-[11%] px-5 py-4 text-right text-xs font-bold text-slate-300 uppercase tracking-wider">Thao tác</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-white/10 bg-slate-900/70">
          <tr v-for="app in paginatedApplicants" :key="app.id" class="transition-colors hover:bg-cyan-400/10">
            <td class="px-5 py-4">
              <div class="flex items-start">
                <img v-if="app.student?.avatar_url" :src="app.student.avatar_url" class="h-10 w-10 rounded-full object-cover border border-white/10 mr-3 shadow-sm" />
                <div v-else class="h-10 w-10 rounded-full bg-slate-800 border border-white/10 flex items-center justify-center text-cyan-200 font-bold mr-3 shadow-sm">
                  {{ app.student?.full_name?.charAt(0) || 'S' }}
                </div>
                
                <div class="min-w-0 flex flex-col space-y-1.5">
                  <div class="flex items-center gap-2">
                    <span class="truncate text-sm font-bold text-white tracking-tight">{{ app.student?.full_name || 'N/A' }}</span>
                    
                    <a 
                      v-if="app.student?.cv_url" 
                      :href="app.student.cv_url" 
                      target="_blank" 
                      class="inline-flex flex-shrink-0 items-center space-x-1 text-[10px] text-rose-200 hover:text-white font-bold bg-rose-500/10 hover:bg-rose-500/20 px-2 py-0.5 rounded border border-rose-400/20 transition-all duration-150 shadow-sm"
                    >
                      <svg class="h-3 w-3" fill="currentColor" viewBox="0 0 20 20">
                        <path d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4z" />
                      </svg>
                      <span>Hồ sơ CV</span>
                    </a>
                  </div>

                  <div class="flex flex-wrap gap-1 max-w-[260px]">
                    <span 
                      v-for="(skill, sIdx) in parseSkills(app.student?.skills)" 
                      :key="sIdx"
                      class="inline-block bg-slate-800 text-slate-300 border border-white/10 text-[10px] font-bold px-1.5 py-0.5 rounded-md shadow-sm"
                    >
                      {{ skill }}
                    </span>
                    <span v-if="!app.student?.skills || parseSkills(app.student?.skills).length === 0" class="text-xs text-slate-400 italic">
                      Chưa cập nhật kỹ năng
                    </span>
                  </div>
                </div>
              </div>
            </td>

            <td class="px-5 py-4 text-sm font-semibold text-slate-200">
              <div class="truncate">{{ app.job?.title || jobTitleLookup(app.job_id) }}</div>
            </td>

            <td class="px-5 py-4 whitespace-nowrap text-sm text-slate-400 font-medium">
              {{ app.student?.phone || 'Chưa cập nhật' }}
            </td>

            <td class="px-5 py-4 whitespace-nowrap text-sm text-slate-400 font-medium">
              {{ formatDate(app.applied_at || app.id) }}
            </td>

            <td class="px-5 py-4 whitespace-nowrap">
              <span :class="[
                statusBadgeClass(app.status),
                'inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold border capitalize'
              ]">
                {{ app.status ? app.status.replace('_', ' ') : 'Chờ duyệt' }}
              </span>
            </td>

            <td class="px-5 py-4 text-right text-sm font-medium">
              <div class="flex items-center justify-end gap-2">
                <button 
                  @click="openChatModal(app)" 
                  class="inline-flex h-9 items-center justify-center rounded-lg border border-white/10 bg-white/5 px-3 text-[11px] font-bold uppercase tracking-wider text-slate-200 transition-all hover:bg-white/10 hover:text-white"
                >
                  <span>Nhắn tin</span>
                </button>

                <button
                  @click="openManagedApplicantModal(app)"
                  class="inline-flex h-9 items-center justify-center rounded-lg bg-cyan-400 px-3 text-[11px] font-extrabold uppercase tracking-wider text-slate-950 shadow-sm shadow-cyan-500/20 transition-all hover:bg-cyan-300"
                >
                  Quản lý
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Mobile Applicants Cards -->
    <div class="lg:hidden space-y-4">
      <div
        v-for="app in paginatedApplicants"
        :key="app.id"
        class="bg-slate-900/85 border border-white/10 rounded-xl p-5 shadow-lg shadow-slate-950/25 space-y-4 hover:border-cyan-300/40 hover:shadow-cyan-950/30 transition-all backdrop-blur"
      >
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <img v-if="app.student?.avatar_url" :src="app.student.avatar_url" class="h-10 w-10 rounded-full border border-white/10 mr-3" />
            <div v-else class="h-10 w-10 bg-slate-800 rounded-full border border-white/10 flex items-center justify-center font-bold text-cyan-200 mr-3">
              {{ app.student?.full_name?.charAt(0) || 'S' }}
            </div>
            <div>
              <h4 class="text-sm font-bold text-white">{{ app.student?.full_name || 'N/A' }}</h4>
              <p class="text-xs text-slate-500">Kỹ năng: {{ app.student?.skills || 'Chưa cập nhật' }}</p>
            </div>
          </div>
          <span :class="[
            statusBadgeClass(app.status),
            'inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold border'
          ]">
            {{ app.status }}
          </span>
        </div>

        <div class="pt-2 border-t border-white/10 text-xs font-medium text-slate-500 space-y-1">
          <div>Vị trí: <span class="text-slate-200 font-bold">{{ jobTitleLookup(app.job_id) }}</span></div>
          <div>Liên hệ: <span class="text-slate-200">{{ app.student?.phone }}</span></div>
          <div>Ngày nộp: <span class="text-slate-200">{{ formatDate(app.applied_at || app.id) }}</span></div>
        </div>

        <!-- Mobile actions -->
        <div class="grid grid-cols-2 gap-2 pt-2">
          <button
            @click="openChatModal(app)"
            class="w-full rounded-lg border border-white/10 bg-white/5 py-2 text-center text-xs font-bold text-slate-200 shadow-sm hover:bg-white/10 hover:text-white"
          >
            Nhắn tin
          </button>

          <button
            @click="openManagedApplicantModal(app)"
            class="w-full rounded-lg bg-cyan-400 py-2 text-center text-xs font-extrabold text-slate-950 shadow-sm shadow-cyan-500/20 hover:bg-cyan-300"
          >
            Quản lý
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
</script>
