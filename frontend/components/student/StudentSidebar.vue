<template>
    <!-- Sidebar Navigation -->
    <aside class="w-full md:w-64 bg-white border-r border-slate-200 flex-shrink-0">
      <div class="p-6 border-b border-slate-100 flex items-center space-x-3">
        <span class="h-8 w-8 rounded-lg bg-gradient-to-tr from-blue-600 to-emerald-500 flex items-center justify-center text-white font-bold text-sm">
          S
        </span>
        <div>
          <h1 class="font-bold text-slate-900 text-sm">Student Portal</h1>
          <p class="text-xs text-slate-500 font-medium">QuickWork Platform</p>
        </div>
      </div>
      <nav class="p-4 space-y-1">
        <button
          v-for="item in navItems"
          :key="item.id"
          @click="activeSection = item.id"
          :class="[
            activeSection === item.id
              ? 'bg-blue-50 text-blue-700 font-semibold'
              : 'text-slate-600 hover:text-slate-900 hover:bg-slate-50',
            'w-full flex items-center space-x-3 px-4 py-2.5 text-sm rounded-xl transition-all duration-150'
          ]"
        >
          <!-- Render icons inline to avoid runtime compilation errors in Nuxt -->
          <svg v-if="item.id === 'jobs'" class="h-5 w-5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
          </svg>
          <svg v-else-if="item.id === 'profile'" class="h-5 w-5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
          </svg>
          <svg v-else-if="item.id === 'applications'" class="h-5 w-5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
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
