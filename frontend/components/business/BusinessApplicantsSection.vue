<template>
      <!-- Section 4: Employer Applicant Management (B6) -->
      <div v-show="activeSection === 'applicants'" class="space-y-6">
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between pb-6 border-b border-slate-200">
          <div>
            <h2 class="text-3xl font-extrabold text-slate-900 tracking-tight">Applicant Management</h2>
            <p class="mt-1 text-sm text-slate-500 font-medium">Review and process student applications.</p>
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
              v-model="applicantSearchQuery"
              type="text"
              class="block w-full pl-10 pr-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 placeholder-slate-400 text-slate-900 focus:outline-none focus:ring-2 focus:ring-violet-500 focus:bg-white transition-all duration-200"
              placeholder="Search by student name or job..."
            />
          </div>
          <div>
            <select
              v-model="applicantStatusFilter"
              class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-violet-500 focus:bg-white transition-all duration-200"
            >
              <option value="all">All Statuses</option>
              <option value="applied">Applied (Pending)</option>
              <option value="approved">Approved</option>
              <option value="rejected">Rejected</option>
            </select>
          </div>
        </div>

        <!-- Loading state skeleton -->
        <div v-if="isLoadingApps" class="bg-white rounded-xl border border-slate-200 p-6 space-y-4 animate-pulse">
          <div class="h-6 bg-slate-200 rounded w-1/4"></div>
          <div class="space-y-3">
            <div class="h-10 bg-slate-100 rounded"></div>
            <div class="h-10 bg-slate-100 rounded"></div>
            <div class="h-10 bg-slate-100 rounded"></div>
          </div>
        </div>

        <!-- Empty state -->
        <div v-else-if="filteredApps.length === 0" class="bg-white text-center py-16 px-4 rounded-xl border border-slate-200 shadow-sm">
          <div class="inline-flex p-4 rounded-full bg-slate-50 border border-slate-100 text-slate-400 mb-4">
            <svg class="h-12 w-12" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
          </div>
          <h3 class="text-lg font-bold text-slate-900">No applicants found</h3>
          <p class="text-sm text-slate-500 mt-1 max-w-sm mx-auto font-medium">
            There are currently no incoming applications matching these parameters.
          </p>
        </div>

      <div v-else class="hidden lg:block bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden">
  <table class="min-w-full divide-y divide-slate-200">
    <thead class="bg-slate-50">
      <tr>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Candidate</th>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Position Applied</th>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Contact Phone</th>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Applied Date</th>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Status</th>
        <th class="px-6 py-4 text-center text-xs font-bold text-slate-500 uppercase tracking-wider">Actions</th>
      </tr>
    </thead>
    <tbody class="bg-white divide-y divide-slate-100">
  <tr v-for="app in filteredApps" :key="app.id">
    
    <td class="px-6 py-4">
      <div class="flex items-start">
        <img v-if="app.student?.avatar_url" :src="app.student.avatar_url" class="h-10 w-10 rounded-full object-cover border border-slate-200 mr-3 shadow-sm" />
        <div v-else class="h-10 w-10 rounded-full bg-slate-100 border border-slate-200 flex items-center justify-center text-slate-500 font-bold mr-3 shadow-sm">
          {{ app.student?.full_name?.charAt(0) || 'S' }}
        </div>
        
        <div class="flex flex-col space-y-1.5">
          <div class="flex items-center gap-2">
            <span class="text-sm font-bold text-slate-950 tracking-tight">{{ app.student?.full_name || 'N/A' }}</span>
            
            <a 
              v-if="app.student?.cv_url" 
              :href="app.student.cv_url" 
              target="_blank" 
              class="inline-flex items-center space-x-1 text-[10px] text-rose-600 hover:text-rose-800 font-bold bg-rose-50 hover:bg-rose-100 px-2 py-0.5 rounded border border-rose-200/60 transition-all duration-150 shadow-sm"
            >
              <svg class="h-3 w-3" fill="currentColor" viewBox="0 0 20 20">
                <path d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4z" />
              </svg>
              <span>CV Profile</span>
            </a>
          </div>

          <div class="flex flex-wrap gap-1 max-w-[260px]">
            <span 
              v-for="(skill, sIdx) in parseSkills(app.student?.skills)" 
              :key="sIdx"
              class="inline-block bg-slate-50 text-slate-600 border border-slate-200/80 text-[10px] font-bold px-1.5 py-0.5 rounded-md shadow-sm"
            >
              {{ skill }}
            </span>
            <span v-if="!app.student?.skills || parseSkills(app.student?.skills).length === 0" class="text-xs text-slate-400 italic">
              No skills specified
            </span>
          </div>
        </div>
      </div>
    </td>

    <td class="px-6 py-4 whitespace-nowrap text-sm font-semibold text-slate-800">
      {{ app.job?.title || jobTitleLookup(app.job_id) }}
    </td>

    <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-600 font-medium">
      {{ app.student?.phone || 'N/A' }}
    </td>

    <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500 font-medium">
      {{ formatDate(app.applied_at || app.id) }}
    </td>

    <td class="px-6 py-4 whitespace-nowrap">
      <span :class="[
        statusBadgeClass(app.status),
        'inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold border capitalize'
      ]">
        {{ app.status ? app.status.replace('_', ' ') : 'Pending' }}
      </span>
    </td>

    <td class="px-6 py-4 whitespace-nowrap text-center text-sm font-medium">
      <div class="flex justify-center items-center gap-2">
        <button 
          @click="openChatModal(app)" 
          class="px-3 py-1.5 bg-slate-800 hover:bg-slate-700 text-slate-200 text-[11px] font-bold uppercase tracking-wider rounded-lg border border-slate-700/60 transition-all flex items-center space-x-1.5"
        >
          <span>💬 Chat với Ứng viên</span>
        </button>

        <button
          v-if="app.status?.toLowerCase() === 'pending'"
          @click="openReviewModal(app)"
          class="px-3 py-1.5 text-xs font-bold text-indigo-600 bg-indigo-50 hover:bg-indigo-100 rounded-lg transition-all duration-150"
        >
          Review & Offer
        </button>
        
        <button
          v-else
          @click="openReviewModal(app)"
          class="px-3 py-1.5 text-xs font-medium text-slate-500 bg-slate-100 hover:bg-slate-200 rounded-lg transition-all duration-150"
        >
          View Details
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
            v-for="app in filteredApps"
            :key="app.id"
            class="bg-white border border-slate-200 rounded-xl p-5 shadow-sm space-y-4"
          >
            <div class="flex items-center justify-between">
              <div class="flex items-center">
                <img v-if="app.student?.avatar_url" :src="app.student.avatar_url" class="h-10 w-10 rounded-full border border-slate-200 mr-3" />
                <div v-else class="h-10 w-10 bg-slate-100 rounded-full border border-slate-200 flex items-center justify-center font-bold text-slate-500 mr-3">
                  {{ app.student?.full_name?.charAt(0) || 'S' }}
                </div>
                <div>
                  <h4 class="text-sm font-bold text-slate-900">{{ app.student?.full_name || 'N/A' }}</h4>
                  <p class="text-xs text-slate-500">Skills: {{ app.student?.skills || 'N/A' }}</p>
                </div>
              </div>
              <span :class="[
                statusBadgeClass(app.status),
                'inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold border'
              ]">
                {{ app.status }}
              </span>
            </div>

            <div class="pt-2 border-t border-slate-100 text-xs font-medium text-slate-600 space-y-1">
              <div>Position: <span class="text-slate-900 font-bold">{{ jobTitleLookup(app.job_id) }}</span></div>
              <div>Contact: <span class="text-slate-900">{{ app.student?.phone }}</span></div>
              <div>Applied: <span class="text-slate-900">{{ formatDate(app.applied_at || app.id) }}</span></div>
            </div>

            <!-- Mobile actions -->
            <div v-if="app.status?.toLowerCase() === 'applied'" class="flex space-x-2 pt-2">
              <button
                @click="triggerConfirmModal(app, 'approved')"
                class="flex-1 text-center py-2 text-xs font-bold text-white bg-emerald-600 hover:bg-emerald-500 rounded-lg shadow-sm"
              >
                Approve
              </button>
              <button
                @click="triggerConfirmModal(app, 'rejected')"
                class="flex-1 text-center py-2 text-xs font-bold text-white bg-rose-600 hover:bg-rose-500 rounded-lg shadow-sm"
              >
                Reject
              </button>
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
  jobTitleLookup,
  formatDate,
  statusBadgeClass,
  parseSkills,
  openChatModal,
  openReviewModal,
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
