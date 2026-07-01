<template>
      <!-- Section 1: Dashboard (Metrics) -->
      <div v-if="activeSection === 'dashboard'" class="space-y-6">
        <div class="pb-6 border-b border-slate-200">
          <h2 class="text-3xl font-extrabold text-slate-900 tracking-tight font-sans">Performance Overview</h2>
          <p class="mt-1 text-sm text-slate-500 font-medium">Real-time candidate metrics and application actions.</p>
        </div>

        <!-- Metrics Cards -->
        <div class="grid grid-cols-2 lg:grid-cols-4 gap-5">
          <div v-for="card in metricsCards" :key="card.title" class="bg-white p-6 rounded-2xl border border-slate-200 shadow-sm flex flex-col justify-between">
            <span class="text-sm font-semibold text-slate-500 uppercase tracking-wider">{{ card.title }}</span>
            <div class="flex items-baseline justify-between mt-4">
              <span class="text-3xl font-extrabold text-slate-900">{{ card.value }}</span>
              <span :class="[card.color, 'text-xs font-semibold px-2 py-0.5 rounded-full border']">{{ card.label }}</span>
            </div>
          </div>
        </div>

        <!-- Job Postings Summary -->
        <div class="bg-white rounded-xl border border-slate-200 p-6 shadow-sm">
          <h3 class="text-lg font-bold text-slate-900 mb-4">Postings Summary</h3>
          <div class="grid grid-cols-3 gap-4 text-center">
            <div class="p-4 bg-slate-50 rounded-xl">
              <p class="text-xs font-semibold text-slate-500 uppercase">Active Jobs</p>
              <p class="text-2xl font-extrabold text-slate-900 mt-1">{{ jobs.length }}</p>
            </div>
            <div class="p-4 bg-slate-50 rounded-xl">
              <p class="text-xs font-semibold text-slate-500 uppercase">Incoming Applicants</p>
              <p class="text-2xl font-extrabold text-slate-900 mt-1">{{ applications.length }}</p>
            </div>
            <div class="p-4 bg-slate-50 rounded-xl">
              <p class="text-xs font-semibold text-slate-500 uppercase">Fill Ratio</p>
              <p class="text-2xl font-extrabold text-slate-900 mt-1">{{ fillRatio }}%</p>
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
