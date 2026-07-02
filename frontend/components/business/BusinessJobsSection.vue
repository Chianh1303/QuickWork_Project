<template>
      <!-- Section 3: Jobs (List and Post Form) -->
      <div v-show="activeSection === 'jobs'" class="space-y-6">
        <div class="flex justify-end">
          <button
            @click="showCreateForm = !showCreateForm"
            class="px-4 py-2 bg-slate-950 hover:bg-slate-800 text-white font-semibold text-sm rounded-lg shadow-sm focus-ring transition-colors"
          >
            {{ showCreateForm ? 'Back to Listings' : '+ Post a New Job' }}
          </button>
        </div>

        <!-- Post job form -->
        <div v-if="showCreateForm" class="bg-white/95 rounded-xl border border-cyan-100 shadow-sm shadow-slate-950/5 p-6 max-w-2xl mx-auto backdrop-blur">
          <h3 class="text-lg font-bold text-slate-900 border-b border-slate-100 pb-3 mb-6">Configure Job Parameters</h3>
          <form @submit.prevent="handleCreateJob" class="space-y-6">
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
              <div class="sm:col-span-2">
                <label for="job_title" class="block text-sm font-semibold text-slate-700 mb-1">Job Title</label>
                <input
                  id="job_title"
                  type="text"
                  v-model="jobForm.title"
                  required
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-white text-slate-900 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-brand-300"
                  placeholder="e.g. Golang Backend developer Trainee"
                />
              </div>

              <div>
                <label for="job_location" class="block text-sm font-semibold text-slate-700 mb-1">Location</label>
                <input
                  id="job_location"
                  type="text"
                  v-model="jobForm.location"
                  required
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-white text-slate-900 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-brand-300"
                  placeholder="e.g. Cau Giay, Hanoi"
                />
              </div>

              <div>
                <label for="job_working_date" class="block text-sm font-semibold text-slate-700 mb-1">Working Schedule</label>
                <input
                  id="job_working_date"
                  type="text"
                  v-model="jobForm.working_date"
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-white text-slate-900 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-brand-300"
                  placeholder="e.g. Monday - Friday"
                />
              </div>

              <div>
                <label for="job_salary" class="block text-sm font-semibold text-slate-700 mb-1">Salary (VND)</label>
                <input
                  id="job_salary"
                  type="number"
                  step="0.01"
                  v-model="jobForm.salary"
                  required
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-white text-slate-900 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-brand-300"
                  placeholder="e.g. 5000000"
                />
              </div>

              <div>
                <label for="job_slots" class="block text-sm font-semibold text-slate-700 mb-1">Slots Available</label>
                <input
                  id="job_slots"
                  type="number"
                  v-model="jobForm.slots"
                  required
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-white text-slate-900 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-brand-300"
                  placeholder="1"
                />
              </div>

              <div class="sm:col-span-2">
                <label for="job_description" class="block text-sm font-semibold text-slate-700 mb-1">Detailed Description</label>
                <textarea
                  id="job_description"
                  rows="4"
                  v-model="jobForm.description"
                  required
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-white text-slate-900 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-brand-300"
                  placeholder="Responsibilities, requirements, skills..."
                ></textarea>
              </div>
            </div>

            <div class="pt-4 border-t border-slate-100 flex justify-end">
              <button
                type="submit"
                :disabled="isCreatingJob"
                class="px-6 py-2.5 border border-transparent text-sm font-semibold rounded-lg text-white bg-brand-600 hover:bg-brand-500 focus-ring disabled:opacity-50 disabled:cursor-not-allowed shadow-sm transition-all duration-150"
              >
                <span v-if="isCreatingJob" class="flex items-center space-x-2">
                  <svg class="animate-spin h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  <span>Creating Job...</span>
                </span>
                <span v-else>Publish Job Listing</span>
              </button>
            </div>
          </form>
        </div>

        <!-- Job listings table -->
        <div v-else class="bg-white/95 rounded-xl border border-slate-200 shadow-sm shadow-slate-950/5 overflow-hidden backdrop-blur">
          <div v-if="jobs.length === 0" class="text-center py-16 text-slate-500">
            <p class="font-bold text-slate-700 text-lg">No active jobs posted</p>
            <p class="text-sm text-slate-400 mt-1">Publish job openings to discover talent.</p>
          </div>
          <div v-else class="overflow-x-auto">
            <table class="min-w-full divide-y divide-slate-200">
              <thead class="bg-slate-950">
                <tr>
                  <th class="px-6 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Job Title</th>
                  <th class="px-6 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Location</th>
                  <th class="px-6 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Salary</th>
                  <th class="px-6 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Slots</th>
                  <th class="px-6 py-4 text-left text-xs font-bold text-slate-300 uppercase tracking-wider">Status</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-slate-100">
                <tr v-for="job in paginatedBusinessJobs" :key="job.id" class="hover:bg-cyan-50/40 transition-colors">
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-semibold text-slate-900">{{ displayJobTitle(job.title) }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">{{ job.location }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-bold text-emerald-700">{{ formatCurrency(job.salary) }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">{{ job.slots }}</td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span :class="[
                      job.status === 'approved' ? 'bg-emerald-50 text-emerald-700 border-emerald-200' : 'bg-amber-50 text-amber-700 border-amber-200',
                      'inline-flex items-center px-2 py-0.5 rounded-full text-xs font-semibold border'
                    ]">
                      {{ job.status || 'pending' }}
                    </span>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <PaginationControls
          v-if="!showCreateForm && jobs.length > 0"
          :page="businessJobsPage"
          :page-size="businessJobsPageSize"
          :total-items="jobs.length"
          @update:page="businessJobsPage = $event"
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
  paginatedBusinessJobs,
  businessJobsPage,
  businessJobsPageSize,
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

const formatCurrency = (value: number | string | null | undefined) => {
  const amount = Number(value || 0)
  return `${amount.toLocaleString('vi-VN')} VND`
}

const displayJobTitle = (title: string | null | undefined) => {
  return (title || 'Untitled Job').replace(/\bMarketting\b/gi, 'Marketing')
}
</script>
