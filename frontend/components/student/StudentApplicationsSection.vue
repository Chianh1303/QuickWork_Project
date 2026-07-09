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
  class="hidden md:block overflow-hidden rounded-xl border border-white/10 bg-slate-900/82 shadow-lg shadow-slate-950/25 backdrop-blur"
>
  <table class="min-w-full table-fixed divide-y divide-white/10">
    <thead class="bg-slate-950">
      <tr>
        <th class="w-[27%] px-5 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Position & Company</th>
        <th class="w-[18%] px-5 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Location</th>
        <th class="w-[16%] px-5 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Compensation</th>
        <th class="w-[14%] px-5 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Applied Date</th>
        <th class="w-[13%] px-5 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Status</th>
        <th class="w-[12%] px-5 py-4 text-right text-xs font-bold text-slate-300 uppercase tracking-wider">Actions</th>
      </tr>
    </thead>

    <tbody class="bg-slate-900/82 divide-y divide-white/10">
        <tr v-for="app in paginatedApps" :key="app.id" class="hover:bg-cyan-400/10 transition-colors">
          <td class="px-5 py-4">
            <div class="truncate text-sm font-semibold text-white">
              {{ app.job?.title || 'Unknown Position' }}
            </div>
            <div class="truncate text-xs text-slate-500 font-medium">
              {{ companyNameLookup(app.job) }}
            </div>
          </td>

          <td class="px-5 py-4 text-sm text-slate-400 font-medium">
            <div class="truncate">{{ app.job?.location || 'Location N/A' }}</div>
          </td>

          <td class="px-5 py-4 whitespace-nowrap text-sm font-extrabold text-emerald-500">
            ${{ Number(app.job?.salary || 0).toLocaleString() }}
          </td>

          <td class="px-5 py-4 whitespace-nowrap text-sm text-slate-400 font-medium">
            {{ formatDate(app.applied_at || app.id) }}
          </td>

          <td class="px-5 py-4 whitespace-nowrap">
            <span
              :class="[
                statusBadgeClass(app.status),
                'inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold border capitalize'
              ]"
            >
              {{ app.status ? app.status.replace('_', ' ') : 'Pending' }}
            </span>
          </td>

          <td class="px-5 py-4 text-right text-sm font-medium">
            <div class="flex items-center justify-end gap-2">
              <button
                @click="openChatModal(app)"
                class="inline-flex h-9 items-center justify-center rounded-lg border border-white/10 bg-white/5 px-3 text-[11px] font-bold uppercase tracking-wider text-slate-200 transition-all hover:bg-white/10 hover:text-white"
              >
                <span>Chat</span>
              </button>

              <button
                @click="openManagedApplicationModal(app)"
                class="inline-flex h-9 items-center justify-center rounded-lg bg-cyan-400 px-3 text-[11px] font-extrabold uppercase tracking-wider text-slate-950 shadow-sm shadow-cyan-500/20 transition-all hover:bg-cyan-300"
              >
                Manage
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
            class="bg-slate-900/82 border border-white/10 rounded-xl p-5 shadow-lg shadow-slate-950/25 space-y-3 hover:border-cyan-300/50 hover:shadow-cyan-950/30 transition-all backdrop-blur"
          >
            <div class="flex justify-between items-start">
              <div>
                <h4 class="text-base font-bold text-white">{{ app.job?.title || 'Unknown Position' }}</h4>
                <p class="text-sm text-slate-500 font-medium">{{ companyNameLookup(app.job) }}</p>
              </div>
              <span :class="[
                statusBadgeClass(app.status),
                'inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold border'
              ]">
                {{ app.status }}
              </span>
            </div>

            <div class="grid grid-cols-2 gap-2 pt-2 text-xs border-t border-white/10 font-medium text-slate-500">
              <div>Location: <span class="text-slate-200">{{ app.job?.location || 'Location N/A' }}</span></div>
              <div>Salary: <span class="text-emerald-600 font-bold">${{ Number(app.job?.salary || 0).toLocaleString() }}</span></div>
              <div class="col-span-2">Applied Date: <span class="text-slate-200">{{ formatDate(app.applied_at || app.id) }}</span></div>
            </div>

            <div class="grid grid-cols-2 gap-2 pt-2">
              <button
                @click="openChatModal(app)"
                class="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-xs font-bold text-slate-200 hover:bg-white/10 hover:text-white"
              >
                Chat
              </button>

              <button
                @click="openManagedApplicationModal(app)"
                class="w-full rounded-lg bg-cyan-400 px-3 py-2 text-xs font-extrabold text-slate-950 hover:bg-cyan-300"
              >
                Manage
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
