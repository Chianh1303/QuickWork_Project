<template>
  <!-- Section 3: Student Application History (C6) -->
  <div v-if="activeSection === 'applications'" class="max-w-6xl mx-auto space-y-6">
    <!-- Header Banner -->
    <div class="rounded-3xl border border-cyan-500/20 bg-slate-900/90 p-6 shadow-xl backdrop-blur-xl flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
      <div>
        <span class="inline-flex rounded-full bg-cyan-500/10 px-3 py-1 text-xs font-black uppercase tracking-wider text-cyan-300 ring-1 ring-cyan-500/30">
          Cổng Theo Dõi Đơn Ứng Tuyển
        </span>
        <h2 class="mt-2 text-2xl font-extrabold text-white tracking-tight">Danh Sách Đơn Ứng Tuyển Của Tôi</h2>
        <p class="mt-1 text-xs font-semibold text-slate-300">Theo dõi tiến trình xét duyệt đơn, phản hồi offer & tiến độ ca làm việc với Doanh nghiệp</p>
      </div>

      <div class="flex items-center gap-2">
        <span class="rounded-2xl border border-cyan-500/20 bg-slate-950/80 px-4 py-2 text-xs font-extrabold text-cyan-200">
          Tổng số: <strong class="text-white text-sm font-black ml-1">{{ filteredApps.length }}</strong> đơn
        </span>
      </div>
    </div>

    <!-- Search & Filter Controls -->
    <div class="bg-slate-900/90 p-4 rounded-2xl border border-cyan-500/20 shadow-xl shadow-cyan-950/30 backdrop-blur-xl grid grid-cols-1 sm:grid-cols-2 gap-4">
      <div class="relative">
        <span class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-slate-500">
          <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </span>
        <input
          v-model="appSearchQuery"
          type="text"
          class="block w-full pl-10 pr-3.5 py-2.5 border border-cyan-500/20 rounded-xl text-xs bg-slate-950/80 placeholder-slate-500 text-slate-100 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium transition-all"
          placeholder="Tìm theo tên vị trí công việc hoặc tên công ty..."
        />
      </div>
      <div>
        <select
          v-model="appStatusFilter"
          class="block w-full px-3.5 py-2.5 border border-cyan-500/20 rounded-xl text-xs bg-slate-950/80 text-slate-100 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium transition-all"
        >
          <option value="all">Tất cả trạng thái đơn</option>
          <option value="pending">🟡 Đang chờ xét duyệt (Pending)</option>
          <option value="approved">🟢 Đã được Doanh nghiệp duyệt</option>
          <option value="offer_accepted">✨ Đã nhận việc / Chấp nhận Offer</option>
          <option value="student_completed">⏳ Đã xong ca - Chờ xác nhận lương</option>
          <option value="paid">💰 Đã hoàn thành & Giải ngân lương</option>
          <option value="rejected">🔴 Đã bị từ chối</option>
        </select>
      </div>
    </div>

    <!-- Loading / Skeleton -->
    <div v-if="isLoadingApps" class="bg-slate-900/90 rounded-3xl border border-cyan-500/10 p-6 space-y-4 animate-pulse">
      <div class="h-6 bg-slate-800 rounded-lg w-1/4"></div>
      <div class="space-y-3">
        <div class="h-12 bg-slate-800/60 rounded-xl"></div>
        <div class="h-12 bg-slate-800/60 rounded-xl"></div>
        <div class="h-12 bg-slate-800/60 rounded-xl"></div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else-if="filteredApps.length === 0" class="bg-slate-900/90 text-center py-16 px-4 rounded-3xl border border-cyan-500/15 shadow-xl">
      <div class="inline-flex p-4 rounded-2xl bg-cyan-500/10 border border-cyan-500/20 text-cyan-300 mb-4">
        <svg class="h-10 w-10" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
        </svg>
      </div>
      <h3 class="text-base font-extrabold text-white">Chưa có đơn ứng tuyển nào phù hợp</h3>
      <p class="text-xs text-slate-400 mt-1 max-w-sm mx-auto font-medium leading-relaxed">
        Bạn chưa ứng tuyển công việc nào theo bộ lọc này. Hãy sang mục "Tìm việc làm" để nộp đơn ngay!
      </p>
    </div>

    <!-- Desktop Applications Table -->
    <div
      v-if="filteredApps.length > 0"
      class="hidden md:block overflow-hidden rounded-3xl border border-cyan-500/20 bg-slate-900/90 shadow-xl"
    >
      <table class="min-w-full table-fixed divide-y divide-cyan-500/10">
        <thead class="bg-slate-950/90">
          <tr>
            <th class="w-[30%] px-5 py-4 text-left text-xs font-black uppercase tracking-wider text-cyan-300">Vị trí & Công ty</th>
            <th class="w-[18%] px-5 py-4 text-left text-xs font-black uppercase tracking-wider text-cyan-300">Địa điểm</th>
            <th class="w-[16%] px-5 py-4 text-left text-xs font-black uppercase tracking-wider text-cyan-300">Thỏa thuận Lương</th>
            <th class="w-[14%] px-5 py-4 text-left text-xs font-black uppercase tracking-wider text-cyan-300">Ngày nộp</th>
            <th class="w-[12%] px-5 py-4 text-left text-xs font-black uppercase tracking-wider text-cyan-300">Trạng thái</th>
            <th class="w-[10%] px-5 py-4 text-right text-xs font-black uppercase tracking-wider text-cyan-300">Thao tác</th>
          </tr>
        </thead>

        <tbody class="bg-slate-900/90 divide-y divide-cyan-500/10">
          <tr v-for="app in paginatedApps" :key="app.id" class="hover:bg-cyan-500/5 transition-colors">
            <td class="px-5 py-4">
              <div class="truncate text-xs font-extrabold text-white">
                {{ app.job?.title || 'Công việc chưa xác định' }}
              </div>
              <div class="truncate text-[11px] font-bold text-cyan-300 mt-0.5">
                {{ companyNameLookup(app.job) }}
              </div>
            </td>

            <td class="px-5 py-4 text-xs text-slate-300 font-medium">
              <div class="truncate">{{ app.job?.location || 'Chưa cập nhật' }}</div>
            </td>

            <td class="px-5 py-4 whitespace-nowrap text-xs font-black text-emerald-300">
              {{ Number(app.job?.salary || 0).toLocaleString('vi-VN') }} VNĐ
            </td>

            <td class="px-5 py-4 whitespace-nowrap text-xs text-slate-400 font-semibold">
              {{ formatDate(app.applied_at || app.id) }}
            </td>

            <td class="px-5 py-4 whitespace-nowrap">
              <span
                :class="[
                  statusBadgeClass(app.status),
                  'inline-flex items-center px-2.5 py-1 rounded-full text-[11px] font-extrabold border capitalize'
                ]"
              >
                {{ app.status ? app.status.replace('_', ' ') : 'Pending' }}
              </span>
            </td>

            <td class="px-5 py-4 text-right text-xs font-medium">
              <div class="flex items-center justify-end gap-2">
                <button
                  @click="openChatModal(app)"
                  class="inline-flex h-8 items-center justify-center rounded-xl border border-cyan-500/30 bg-cyan-500/10 px-3 text-[11px] font-extrabold uppercase tracking-wider text-cyan-200 transition-all hover:bg-cyan-500/20 hover:text-white"
                >
                  <span>Chat</span>
                </button>

                <button
                  @click="openManagedApplicationModal(app)"
                  class="inline-flex h-8 items-center justify-center rounded-xl bg-[#22D3EE] px-3.5 text-[11px] font-black uppercase tracking-wider text-slate-950 shadow-md shadow-cyan-400/20 transition-all hover:bg-[#67E8F9] hover:scale-105 active:scale-95"
                >
                  Chi tiết
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Mobile Applications Cards -->
    <div v-show="filteredApps.length > 0" class="md:hidden space-y-4">
      <div
        v-for="app in paginatedApps"
        :key="app.id"
        class="bg-slate-900/90 border border-cyan-500/20 rounded-2xl p-5 space-y-3 shadow-xl"
      >
        <div class="flex justify-between items-start">
          <div>
            <h4 class="text-sm font-extrabold text-white">{{ app.job?.title || 'Unknown Position' }}</h4>
            <p class="text-xs text-cyan-300 font-bold mt-0.5">{{ companyNameLookup(app.job) }}</p>
          </div>
          <span :class="[
            statusBadgeClass(app.status),
            'inline-flex items-center px-2.5 py-1 rounded-full text-[11px] font-extrabold border'
          ]">
            {{ app.status }}
          </span>
        </div>

        <div class="grid grid-cols-2 gap-2 pt-2 text-xs border-t border-cyan-500/10 font-semibold text-slate-400">
          <div>Địa điểm: <span class="text-slate-200">{{ app.job?.location || 'N/A' }}</span></div>
          <div>Lương: <span class="text-emerald-300 font-extrabold">{{ Number(app.job?.salary || 0).toLocaleString('vi-VN') }}đ</span></div>
          <div class="col-span-2">Ngày nộp: <span class="text-slate-200">{{ formatDate(app.applied_at || app.id) }}</span></div>
        </div>

        <div class="grid grid-cols-2 gap-2 pt-2">
          <button
            @click="openChatModal(app)"
            class="w-full rounded-xl border border-cyan-500/30 bg-cyan-500/10 px-3 py-2 text-xs font-extrabold text-cyan-200 hover:bg-cyan-500/20 hover:text-white"
          >
            Nhắn tin
          </button>

          <button
            @click="openManagedApplicationModal(app)"
            class="w-full rounded-xl bg-[#22D3EE] px-3 py-2 text-xs font-black uppercase tracking-wider text-slate-950 shadow-md shadow-cyan-400/20 hover:bg-[#67E8F9] transition-all"
          >
            Chi tiết
          </button>
        </div>
      </div>
    </div>

    <PaginationControls
      v-if="filteredApps.length > 0"
      :page="applicationsPage"
      :page-size="applicationsPageSize"
      :total-items="filteredApps.length"
      @update:page="applicationsPage = $event"
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
  paginatedApps,
  applicationsPage,
  applicationsPageSize,
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
  handleStudentComplete,
  openReviewModal,
  openManagedApplicationModal,
  toast
} = toRefs(props.state)
</script>
