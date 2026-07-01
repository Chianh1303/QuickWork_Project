<template>
      <!-- Section 3: Student Application History (C6) -->
      <div v-if="activeSection === 'applications'" class="space-y-6">
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between pb-6 border-b border-slate-200">
          <div>
            <h2 class="text-3xl font-extrabold text-slate-900 tracking-tight">My Applications</h2>
            <p class="mt-1 text-sm text-slate-500 font-medium">Track every opportunity you've applied for.</p>
          </div>
          <div class="mt-4 sm:mt-0 bg-white border border-slate-200 px-4 py-2 rounded-xl shadow-sm text-sm font-semibold text-slate-700">
            Total Applications: {{ filteredApps.length }}
          </div>
        </div>

        <!-- Search & Filter Controls -->
        <div class="bg-white p-4 rounded-xl border border-slate-200 shadow-sm grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div class="relative">
            <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
              <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </span>
            <input
              v-model="appSearchQuery"
              type="text"
              class="block w-full pl-10 pr-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 placeholder-slate-400 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white transition-all duration-200"
              placeholder="Search by job title or company..."
            />
          </div>
          <div>
            <select
              v-model="appStatusFilter"
              class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white transition-all duration-200"
            >
              <option value="all">All Application Statuses</option>
              <option value="applied">Applied (Pending)</option>
              <option value="approved">Approved</option>
              <option value="rejected">Rejected</option>
            </select>
          </div>
        </div>

        <!-- Loading / Skeleton -->
        <div v-if="isLoadingApps" class="bg-white rounded-xl border border-slate-200 p-6 space-y-4 animate-pulse">
          <div class="h-6 bg-slate-200 rounded w-1/4"></div>
          <div class="space-y-3">
            <div class="h-10 bg-slate-100 rounded"></div>
            <div class="h-10 bg-slate-100 rounded"></div>
            <div class="h-10 bg-slate-100 rounded"></div>
          </div>
        </div>

        <!-- Empty State -->
        <div v-else-if="filteredApps.length === 0" class="bg-white text-center py-16 px-4 rounded-xl border border-slate-200 shadow-sm">
          <div class="inline-flex p-4 rounded-full bg-slate-50 border border-slate-100 text-slate-400 mb-4">
            <svg class="h-12 w-12" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
          </div>
          <h3 class="text-lg font-bold text-slate-900">No applications matched</h3>
          <p class="text-sm text-slate-500 mt-1 max-w-sm mx-auto font-medium">
            You haven't applied to any roles matching these filters. Go to the jobs listing tab to apply!
          </p>
        </div>

        <!-- Desktop Applications Table -->
<div
  v-if="filteredApps.length > 0"
  class="hidden md:block bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden"
>
  <table class="min-w-full divide-y divide-slate-200">
    <thead class="bg-slate-50">
      <tr>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Position & Company</th>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Location</th>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Compensation</th>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Applied Date</th>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Status</th>
        <th class="px-6 py-4 text-right text-xs font-bold text-slate-500 uppercase tracking-wider">Actions</th>
      </tr>
    </thead>
   <tbody class="bg-white divide-y divide-slate-100">
  <template v-for="app in filteredApps" :key="app.id">
    <tr>
      <td class="px-6 py-4 whitespace-nowrap">
        <div class="text-sm font-semibold text-slate-900">{{ app.job?.title || 'Unknown Position' }}</div>
        <div class="text-xs text-slate-500 font-medium">{{ companyNameLookup(app.job?.business_id) }}</div>
      </td>
      <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-600 font-medium">
        {{ app.job?.location || 'Location N/A' }}
      </td>
      <td class="px-6 py-4 whitespace-nowrap text-sm font-extrabold text-emerald-600">
        ${{ Number(app.job?.salary || 0).toLocaleString() }}
      </td>
      <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500 font-medium">
        {{ formatDate(app.applied_at || app.id) }}
      </td>
      <td class="px-6 py-4 whitespace-nowrap">
        <span :class="[
          statusBadgeClass(app.status),
          'inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold border'
        ]">
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
            class="inline-flex items-center space-x-1 px-3 py-1.5 text-xs font-bold text-red-600 hover:text-red-700 hover:bg-red-50 rounded-lg border border-transparent hover:border-red-200 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-150"
          >
            <svg v-if="isCancellingApp === app.id" class="animate-spin h-3 w-3 text-red-600" fill="none" viewBox="0 0 24 24">
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

          <span 
            v-else-if="(app.status || '').toLowerCase() === 'offer_accepted'" 
            class="inline-flex items-center px-2.5 py-1 text-xs font-bold text-emerald-600 bg-emerald-50 border border-emerald-200 rounded-lg select-none"
          >
            🎉 Đã nhận việc
          </span>

          <span v-else class="text-xs font-medium text-slate-400 select-none">-</span>
          
        </div>
      </td>
    </tr>

    <tr v-if="(app.status || '').toLowerCase().replace('_', ' ') === 'offer accepted'" class="bg-slate-50/60">
  <td colspan="6" class="px-6 py-3 border-t border-slate-100">
    <div class="flex items-center justify-between">
      <div class="flex items-center space-x-3">
        <span class="flex h-2 w-2 relative">
          <span v-if="isWorking(app.job_id)" class="animate-ping absolute inline-flex h-full w-full rounded-full bg-amber-400 opacity-75"></span>
          <span :class="isWorking(app.job_id) ? 'bg-amber-500' : 'bg-slate-300'" class="relative inline-flex rounded-full h-2 w-2"></span>
        </span>
        <span class="text-xs font-bold text-slate-500 uppercase tracking-wider">Ca làm việc hôm nay:</span>
        <span class="text-xs font-semibold" :class="isWorking(app.job_id) ? 'text-amber-600 font-bold animate-pulse' : 'text-slate-500'">
          {{ isWorking(app.job_id) ? `⏱️ Đang làm việc (${getTimer(app.job_id)})` : 'Chưa vào ca' }}
        </span>
      </div>

      <div>
        <button
          v-if="!isWorking(app.job_id)"
          @click="handleCheckIn(app.job_id)"
          class="inline-flex items-center space-x-1 px-3 py-1 bg-emerald-600 hover:bg-emerald-500 text-white text-[11px] font-bold rounded-md shadow-sm transition-all duration-150"
        >
          <span>⚡ Check-in đi làm</span>
        </button>

        <button
          v-else
          @click="handleCheckOut(app.job_id)"
          class="inline-flex items-center space-x-1 px-3 py-1 bg-rose-600 hover:bg-rose-500 text-white text-[11px] font-bold rounded-md shadow-sm transition-all duration-150"
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
            v-for="app in filteredApps"
            :key="app.id"
            class="bg-white border border-slate-200 rounded-xl p-5 shadow-sm space-y-3"
          >
            <div class="flex justify-between items-start">
              <div>
                <h4 class="text-base font-bold text-slate-900">{{ app.job?.title || 'Unknown Position' }}</h4>
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
              <div>Location: <span class="text-slate-900">{{ app.job?.location || 'Location N/A' }}</span></div>
              <div>Salary: <span class="text-emerald-600 font-bold">${{ Number(app.job?.salary || 0).toLocaleString() }}</span></div>
              <div class="col-span-2">Applied Date: <span class="text-slate-900">{{ formatDate(app.applied_at || app.id) }}</span></div>
            </div>
          </div>
        </div>
      </div>
</template>

<script setup lang="ts">
import { toRefs } from 'vue'

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
</script>
