<template>
    <!-- Sidebar Navigation -->
    <aside class="w-full md:w-64 bg-white border-r border-slate-200 flex-shrink-0">
      <div class="p-6 border-b border-slate-100 flex items-center space-x-3">
        <span class="h-8 w-8 rounded-lg bg-gradient-to-tr from-violet-600 to-indigo-500 flex items-center justify-center text-white font-bold text-sm">
          B
        </span>
        <div>
          <h1 class="font-bold text-slate-900 text-sm">Employer Console</h1>
          <p class="text-xs text-slate-500 font-medium">QuickWork Business</p>
        </div>
      </div>
      <nav class="p-4 space-y-1">
        <button
          v-for="item in navItems"
          :key="item.id"
          @click="activeSection = item.id"
          :class="[
            activeSection === item.id
              ? 'bg-violet-50 text-violet-700 font-semibold'
              : 'text-slate-600 hover:text-slate-900 hover:bg-slate-50',
            'w-full flex items-center space-x-3 px-4 py-2.5 text-sm rounded-xl transition-all duration-150'
          ]"
        >
          <!-- Render icons inline to avoid runtime compilation errors in Nuxt -->
          <svg v-if="item.id === 'dashboard'" class="h-5 w-5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 002 2h2a2 2 0 002-2z"/>
          </svg>
          <svg v-else-if="item.id === 'profile'" class="h-5 w-5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"/>
          </svg>
          <svg v-else-if="item.id === 'jobs'" class="h-5 w-5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
          </svg>
          <svg v-else-if="item.id === 'applicants'" class="h-5 w-5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"/>
          </svg>
          <span>{{ item.name }}</span>
        </button>
      </nav>
    </aside>
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
