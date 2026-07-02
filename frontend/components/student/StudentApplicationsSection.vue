<template>
      <!-- Section 3: Student Application History (C6) -->
      <div v-if="activeSection === 'applications'" class="space-y-6">
        <!-- Search & Filter Controls -->
        <div class="bg-slate-900/82 p-4 rounded-xl border border-white/10 shadow-lg shadow-slate-950/25 grid grid-cols-1 sm:grid-cols-2 gap-4 backdrop-blur">
          <div class="relative">
            <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
              <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </span>
            <input
              v-model="appSearchQuery"
              type="text"
              class="block w-full pl-10 pr-3 py-2 border border-white/10 rounded-lg text-sm bg-slate-950/70 placeholder-slate-500 text-slate-100 focus:outline-none focus:ring-2 focus:ring-cyan-400 transition-all duration-200"
              placeholder="Search by job title or company..."
            />
          </div>
          <div>
            <select
              v-model="appStatusFilter"
              class="block w-full px-3 py-2 border border-white/10 rounded-lg text-sm bg-slate-950/70 text-slate-100 focus:outline-none focus:ring-2 focus:ring-cyan-400 transition-all duration-200"
            >
              <option value="all">All Application Statuses</option>
              <option value="pending">Applied (Pending)</option>
              <option value="approved">Approved</option>
              <option value="offer_accepted">Offer Accepted</option>
              <option value="student_completed">Waiting Business Confirmation</option>
              <option value="paid">Completed & Paid</option>
              <option value="rejected">Rejected</option>
            </select>
          </div>
        </div>

        <!-- Loading / Skeleton -->
        <div v-if="isLoadingApps" class="bg-slate-900/82 rounded-xl border border-white/10 p-6 space-y-4 animate-pulse">
          <div class="h-6 bg-slate-200 rounded w-1/4"></div>
          <div class="space-y-3">
            <div class="h-10 bg-slate-100 rounded"></div>
            <div class="h-10 bg-slate-100 rounded"></div>
            <div class="h-10 bg-slate-100 rounded"></div>
          </div>
        </div>

        <!-- Empty State -->
        <div v-else-if="filteredApps.length === 0" class="bg-slate-900/82 text-center py-16 px-4 rounded-xl border border-white/10 shadow-lg shadow-slate-950/25 backdrop-blur">
          <div class="inline-flex p-4 rounded-full bg-slate-50 border border-slate-100 text-slate-400 mb-4">
            <svg class="h-12 w-12" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
          </div>
          <h3 class="text-lg font-bold text-white">No applications matched</h3>
          <p class="text-sm text-slate-500 mt-1 max-w-sm mx-auto font-medium">
            You haven't applied to any roles matching these filters. Go to the jobs listing tab to apply!
          </p>
        </div>

     <!-- Desktop Applications Table -->
<div
  v-if="filteredApps.length > 0"
  class="hidden md:block bg-slate-900/82 rounded-xl border border-white/10 shadow-lg shadow-slate-950/25 overflow-hidden backdrop-blur"
>
  <table class="min-w-full divide-y divide-white/10">
    <thead class="bg-slate-950">
      <tr>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Position & Company</th>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Location</th>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Compensation</th>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Applied Date</th>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Status</th>
        <th class="px-6 py-4 text-right text-xs font-bold text-slate-300 uppercase tracking-wider">Actions</th>
      </tr>
    </thead>

    <tbody class="bg-slate-900/82 divide-y divide-white/10">
      <template v-for="app in paginatedApps" :key="app.id">
        <tr class="hover:bg-cyan-400/10 transition-colors">
          <td class="px-6 py-4 whitespace-nowrap">
            <div class="text-sm font-semibold text-white">
              {{ app.job?.title || 'Unknown Position' }}
            </div>
            <div class="text-xs text-slate-500 font-medium">
              {{ companyNameLookup(app.job?.business_id) }}
            </div>
          </td>

          <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-400 font-medium">
            {{ app.job?.location || 'Location N/A' }}
          </td>

          <td class="px-6 py-4 whitespace-nowrap text-sm font-extrabold text-emerald-500">
            ${{ Number(app.job?.salary || 0).toLocaleString() }}
          </td>

          <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-400 font-medium">
            {{ formatDate(app.applied_at || app.id) }}
          </td>

          <td class="px-6 py-4 whitespace-nowrap">
            <span
              :class="[
                statusBadgeClass(app.status),
                'inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold border capitalize'
              ]"
            >
              {{ app.status ? app.status.replace('_', ' ') : 'Pending' }}
            </span>
          </td>

          <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
            <div class="flex items-center justify-end gap-2">
              <button
                @click="openChatModal(app)"
                class="px-3 py-1.5 bg-slate-800 hover:bg-slate-700 text-slate-200 text-[11px] font-bold uppercase tracking-wider rounded-lg border border-slate-700/60 transition-all flex items-center space-x-1.5"
              >
                <span>💬 Chat với HR</span>
              </button>

              <button
                v-if="(app.status || '').toLowerCase() === 'pending'"
                :disabled="isCancellingApp === app.id"
                @click="triggerCancelConfirm(app.id)"
                class="inline-flex items-center space-x-1 px-3 py-1.5 text-xs font-bold text-red-400 hover:text-red-300 hover:bg-red-500/10 rounded-lg border border-transparent hover:border-red-400/20 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-150"
              >
                <svg
                  v-if="isCancellingApp === app.id"
                  class="animate-spin h-3 w-3 text-red-400"
                  fill="none"
                  viewBox="0 0 24 24"
                >
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                <span>Hủy ứng tuyển</span>
              </button>

              <button
                v-else-if="(app.status || '').toLowerCase() === 'approved'"
                @click="openOfferModal(app)"
                class="inline-flex items-center px-3 py-1.5 text-xs font-bold text-white bg-gradient-to-r from-emerald-500 to-teal-600 hover:from-emerald-600 hover:to-teal-700 rounded-lg shadow-sm transition-all duration-150 animate-pulse"
              >
                <span>✨ Xem Offer</span>
              </button>

              <button
                v-else-if="(app.status || '').toLowerCase() === 'offer_accepted'"
                @click="handleStudentComplete(app.id)"
                class="inline-flex items-center px-3 py-1.5 text-xs font-bold text-white bg-cyan-600 hover:bg-cyan-500 rounded-lg shadow-sm transition-all duration-150"
              >
                ✅ Xác nhận hoàn thành
              </button>

              <span
                v-else-if="(app.status || '').toLowerCase() === 'student_completed'"
                class="inline-flex items-center px-3 py-1.5 text-xs font-bold text-amber-400 bg-amber-500/10 border border-amber-400/20 rounded-lg select-none"
              >
                ⏳ Chờ doanh nghiệp xác nhận
              </span>

              <span
                v-else-if="(app.status || '').toLowerCase() === 'paid'"
                class="inline-flex items-center px-3 py-1.5 text-xs font-bold text-emerald-400 bg-emerald-500/10 border border-emerald-400/20 rounded-lg select-none"
              >
                💰 Đã giải ngân
              </span>

              <span v-else class="text-xs font-medium text-slate-400 select-none">-</span>
            </div>
          </td>
        </tr>

        <tr
          v-if="(app.status || '').toLowerCase() === 'offer_accepted'"
          class="bg-slate-950/80"
        >
          <td colspan="6" class="px-6 py-3 border-t border-slate-800 bg-slate-950/80">
            <div class="flex items-center justify-between rounded-xl bg-slate-900/70 border border-slate-800 px-4 py-3">
              <div class="flex items-center space-x-3">
                <span class="flex h-2 w-2 relative">
                  <span
                    v-if="isWorking(app.job_id)"
                    class="animate-ping absolute inline-flex h-full w-full rounded-full bg-amber-400 opacity-75"
                  ></span>

                  <span
                    :class="isWorking(app.job_id) ? 'bg-amber-500' : 'bg-slate-500'"
                    class="relative inline-flex rounded-full h-2 w-2"
                  ></span>
                </span>

                <span class="text-xs font-bold text-slate-300 uppercase tracking-wider">
                  Ca làm việc hôm nay:
                </span>

                <span
                  class="text-xs font-semibold"
                  :class="isWorking(app.job_id) ? 'text-amber-400 font-bold animate-pulse' : 'text-slate-400'"
                >
                  {{ isWorking(app.job_id) ? `⏱️ Đang làm việc (${getTimer(app.job_id)})` : 'Chưa vào ca' }}
                </span>
              </div>

              <div class="flex items-center gap-2">
                <button
                  v-if="!isWorking(app.job_id)"
                  @click="handleCheckIn(app.job_id)"
                  class="inline-flex items-center space-x-1 px-4 py-2 bg-emerald-600 hover:bg-emerald-500 text-white text-[11px] font-bold rounded-lg shadow-sm transition-all duration-150"
                >
                  <span>⚡ Check-in đi làm</span>
                </button>

                <button
                  v-else
                  @click="handleCheckOut(app.job_id)"
                  class="inline-flex items-center space-x-1 px-4 py-2 bg-rose-600 hover:bg-rose-500 text-white text-[11px] font-bold rounded-lg shadow-sm transition-all duration-150"
                >
                  <span>🛑 Check-out ra ca</span>
                </button>
              </div>
            </div>
          </td>
        </tr>
      </template>
    </tbody>
  </table>
</div>
        <!-- Mobile Applications Cards -->
        <div v-show="filteredApps.length > 0" class="md:hidden space-y-4">
          <div
            v-for="app in paginatedApps"
            :key="app.id"
            class="bg-slate-900/82 border border-white/10 rounded-xl p-5 shadow-lg shadow-slate-950/25 space-y-3 hover:border-cyan-300/50 hover:shadow-cyan-950/30 transition-all backdrop-blur"
          >
            <div class="flex justify-between items-start">
              <div>
                <h4 class="text-base font-bold text-white">{{ app.job?.title || 'Unknown Position' }}</h4>
                <p class="text-sm text-slate-500 font-medium">{{ companyNameLookup(app.job?.business_id) }}</p>
              </div>
              <span :class="[
                statusBadgeClass(app.status),
                'inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold border'
              ]">
                {{ app.status }}
              </span>
            </div>

            <div class="grid grid-cols-2 gap-2 pt-2 text-xs border-t border-slate-100 font-medium text-slate-500">
              <div>Location: <span class="text-slate-200">{{ app.job?.location || 'Location N/A' }}</span></div>
              <div>Salary: <span class="text-emerald-600 font-bold">${{ Number(app.job?.salary || 0).toLocaleString() }}</span></div>
              <div class="col-span-2">Applied Date: <span class="text-slate-200">{{ formatDate(app.applied_at || app.id) }}</span></div>
            </div>

            <div class="grid grid-cols-1 gap-2 pt-2">
              <button
                @click="openChatModal(app)"
                class="w-full rounded-lg bg-slate-800 px-3 py-2 text-xs font-bold text-slate-200 hover:bg-slate-700"
              >
                Chat với HR
              </button>

              <button
                v-if="(app.status || '').toLowerCase() === 'pending'"
                :disabled="isCancellingApp === app.id"
                @click="triggerCancelConfirm(app.id)"
                class="w-full rounded-lg border border-rose-400/20 bg-rose-500/10 px-3 py-2 text-xs font-bold text-rose-300 disabled:cursor-not-allowed disabled:opacity-50"
              >
                Hủy ứng tuyển
              </button>

              <button
                v-else-if="(app.status || '').toLowerCase() === 'approved'"
                @click="openOfferModal(app)"
                class="w-full rounded-lg bg-emerald-600 px-3 py-2 text-xs font-bold text-white hover:bg-emerald-500"
              >
                Xem Offer
              </button>

              <button
                v-else-if="(app.status || '').toLowerCase() === 'offer_accepted'"
                @click="handleStudentComplete(app.id)"
                class="w-full rounded-lg bg-cyan-400 px-3 py-2 text-xs font-bold text-slate-950 hover:bg-cyan-300"
              >
                Xác nhận hoàn thành
              </button>

              <div
                v-else-if="(app.status || '').toLowerCase() === 'student_completed'"
                class="rounded-lg border border-amber-400/20 bg-amber-500/10 px-3 py-2 text-center text-xs font-bold text-amber-300"
              >
                Chờ doanh nghiệp xác nhận
              </div>

              <div
                v-else-if="(app.status || '').toLowerCase() === 'paid'"
                class="rounded-lg border border-emerald-400/20 bg-emerald-500/10 px-3 py-2 text-center text-xs font-bold text-emerald-300"
              >
                Đã giải ngân
              </div>
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
  toast
} = toRefs(props.state)
</script>
