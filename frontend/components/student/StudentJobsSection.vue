<template>
      <!-- Section 1: Dashboard (Explore Jobs) -->
      <div v-if="activeSection === 'jobs'" class="space-y-6">
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between pb-6 border-b border-slate-200">
          <div>
            <h2 class="text-3xl font-extrabold text-slate-900 tracking-tight">Available Gigs & Jobs</h2>
            <p class="mt-1 text-sm text-slate-500 font-medium">Explore current work opportunities in real-time.</p>
          </div>
        </div>

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
       <div class="bg-white p-6 rounded-xl border border-slate-200 shadow-sm space-y-4">
          <form @submit.prevent="fetchJobs" class="space-y-4">
            
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="relative">
                <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">🔍</span>
                <input v-model="jobsSearchQuery" type="text" class="block w-full pl-10 pr-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white" placeholder="Search by title, keyword, skills..." />
              </div>

              <div class="relative">
                <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">📍</span>
                <input v-model="jobsLocationQuery" type="text" class="block w-full pl-10 pr-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white" placeholder="Filter by location (e.g. Hanoi, Ho Chi Minh)..." />
              </div>
            </div>

            <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 pt-2">
              <div>
                <select v-model="filterCategory" class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white">
                  <option value="all">📁 All Categories (Ngành nghề)</option>
                  <option value="it">Information Technology</option>
                  <option value="marketing">Marketing</option>
                  <option value="design">Graphic Design</option>
                </select>
              </div>

              <div>
                <select v-model="filterJobType" class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white">
                  <option value="all">⏱️ All Job Types (Hình thức)</option>
                  <option value="full-time">Full-time</option>
                  <option value="part-time">Part-time</option>
                  <option value="intern">Internship</option>
                </select>
              </div>

              <div>
                <select v-model="filterMinSalary" class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white">
                  <option value="">💵 Any Salary (Mức lương)</option>
                  <option value="2000000">> 2,000,000 VND</option>
                  <option value="5000000">> 5,000,000 VND</option>
                  <option value="10000000">> 10,000,000 VND</option>
                </select>
              </div>
            </div>

            <div class="flex items-center gap-2 justify-end pt-2 border-t border-slate-100">
              <button type="button" @click="resetFilters" class="px-4 py-2 border border-slate-200 text-sm font-semibold rounded-lg text-slate-700 bg-white hover:bg-slate-50 transition-all">
                Clear Filters
              </button>
              <button type="submit" class="px-6 py-2 border border-transparent text-sm font-semibold rounded-lg text-white bg-blue-600 hover:bg-blue-500 shadow-sm transition-all">
                Apply Filters
              </button>
            </div>

          </form>
        </div>
        <!-- Skeleton / Loading -->
        <div v-if="isLoadingJobs" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div v-for="n in 6" :key="n" class="bg-white rounded-xl border border-slate-200 p-6 space-y-4 animate-pulse">
            <div class="h-4 bg-slate-200 rounded w-2/3"></div>
            <div class="h-3 bg-slate-200 rounded w-1/2"></div>
            <div class="h-16 bg-slate-100 rounded"></div>
            <div class="h-8 bg-slate-200 rounded"></div>
          </div>
        </div>

        <!-- Jobs Grid -->
        <div v-else-if="filteredJobs.length === 0" class="bg-white text-center py-16 px-4 rounded-xl border border-slate-200 shadow-sm text-slate-500">
          <svg class="mx-auto h-12 w-12 text-slate-300 mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <p class="font-bold text-slate-700 text-lg">No work listings found</p>
          <p class="text-sm text-slate-400 mt-1">Try modifying your search or filters.</p>
        </div>

        <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div
            v-for="job in filteredJobs"
            :key="job.id"
            class="bg-white rounded-xl border border-slate-200 shadow-sm p-6 flex flex-col justify-between hover:shadow-md transition-shadow duration-200"
          >
            <div>
              <div class="flex justify-between items-start mb-3">
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-semibold bg-blue-50 text-blue-700">
                  {{ companyNameLookup(job.business_id) }}
                </span>
                <span class="text-sm font-extrabold text-emerald-600 bg-emerald-50 px-2.5 py-1 rounded-lg">
                  ${{ Number(job.salary || 0).toLocaleString() }}
                </span>
              </div>
              <h3 class="text-lg font-bold text-slate-900 line-clamp-1">{{ job.title || 'Untitled Job' }}</h3>
              <p class="text-sm font-medium text-slate-500 mt-1 flex items-center gap-1">
                <svg class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                </svg>
                {{ job.location || 'Location N/A' }}
              </p>
              <p class="text-sm font-medium text-slate-500 mt-1 flex items-center gap-1" v-if="job.working_date">
                <svg class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
                {{ job.working_date }}
              </p>
              <p class="text-sm text-slate-600 mt-4 line-clamp-4 bg-slate-50 p-3 rounded-lg border border-slate-100 font-medium">
                {{ job.description || 'No description provided.' }}
              </p>
            </div>

       <div class="mt-6 pt-4 border-t border-slate-100">
  <button
    :disabled="checkIfApplied(job.id)"
    @click="handleApply(job)"
    :class="checkIfApplied(job.id) ? 'bg-slate-400 text-slate-100' : 'bg-blue-600 hover:bg-blue-500 text-white'"
    class="w-full flex justify-center py-2.5 px-4 border border-transparent text-sm font-semibold rounded-lg transition-colors focus-ring disabled:cursor-not-allowed shadow-sm"
  >
    <span v-if="isApplying === job.id" class="flex items-center space-x-2">
      <svg class="animate-spin h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      <span>Applying...</span>
    </span>
    
    <span v-else>
      {{ checkIfApplied(job.id) ? '✓ Applied' : 'Apply Instantly' }}
    </span>
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
