<template>
    <!-- Section 2: Company Profile -->
      <div v-show="activeSection === 'profile'" class="max-w-2xl mx-auto space-y-6">
        <!-- ======================================================== -->
        <!-- TRẠNG THÁI 1: HIỂN THỊ DỮ LIỆU CŨ (VIEW MODE) -->
        <!-- ======================================================== -->
        <div v-if="!isEditing" class="space-y-6">
          <!-- Card thông tin công ty chính -->
          <div class="bg-white/95 rounded-2xl border border-cyan-100 shadow-sm shadow-slate-950/5 p-6 flex flex-col sm:flex-row items-center sm:items-start space-y-4 sm:space-y-0 sm:space-x-6 backdrop-blur">
            <img 
              :src="profileForm.logo_url || 'https://images.unsplash.com/photo-1486406146926-c627a92ad1ab?auto=format&fit=crop&w=150&h=150&q=80'" 
              class="w-24 h-24 rounded-xl object-cover border-2 border-slate-100 shadow-sm bg-slate-50 p-1"
            />
            <div class="flex-1 text-center sm:text-left space-y-3">
              <h3 class="text-2xl font-bold text-slate-900">{{ profileForm.company_name || 'Your Company Name' }}</h3>
              
              <div class="space-y-2 text-sm text-slate-600 font-medium">
                <p class="flex items-center justify-center sm:justify-start">
                  <svg class="mr-2 h-4 w-4 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.75 21h16.5M4.5 3h15l-.75 18H5.25L4.5 3zM9 7.5h6M9 11.25h6M9 15h6" />
                  </svg>
                  {{ profileForm.address || 'Address not updated yet' }}
                </p>
                <p class="flex items-center justify-center sm:justify-start">
                  <svg class="mr-2 h-4 w-4 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 002.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106a1.125 1.125 0 00-1.173.417l-.97 1.293a1.125 1.125 0 01-1.21.38 12.035 12.035 0 01-7.143-7.143 1.125 1.125 0 01.38-1.21l1.293-.97c.36-.27.527-.734.417-1.173L6.963 3.102A1.125 1.125 0 005.872 2.25H4.5A2.25 2.25 0 002.25 4.5v2.25z" />
                  </svg>
                  {{ profileForm.phone || 'Corporate phone not updated yet' }}
                </p>
              </div>
            </div>
          </div>
        </div>

        <!-- ======================================================== -->
        <!-- TRẠNG THÁI 2: FORM CHỈNH SỬA (EDIT MODE) -->
        <!-- ======================================================== -->
        <div v-else class="bg-white/95 rounded-2xl border border-cyan-100 shadow-sm shadow-slate-950/5 p-6 backdrop-blur">
          <form @submit.prevent="handleUpdateProfile" class="space-y-6">
            
            <!-- Upload Logo thực tế -->
            <div class="flex items-center space-x-6 bg-slate-50 p-4 rounded-xl border border-slate-100">
              <img 
                :src="logoPreview || profileForm.logo_url || 'https://images.unsplash.com/photo-1486406146926-c627a92ad1ab?auto=format&fit=crop&w=150&h=150&q=80'" 
                class="w-20 h-20 rounded-xl object-cover border border-slate-200 shadow-sm bg-white p-1"
              />
              <div>
                <input type="file" id="profile_logo_file" accept="image/*" class="hidden" @change="onLogoFileChange" />
                <label for="profile_logo_file" class="px-4 py-2 text-xs font-bold text-slate-700 bg-white border border-slate-300 rounded-lg shadow-sm hover:bg-slate-50 cursor-pointer transition-all">
                  Change Company Logo
                </label>
                <p class="text-xs text-slate-400 mt-1.5">JPG or PNG. Max 2MB.</p>
              </div>
            </div>

            <!-- Các ô nhập dữ liệu chữ -->
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
              <div>
                <label for="profile_company" class="block text-sm font-semibold text-slate-700 mb-1">Company Name</label>
                <input
                  id="profile_company"
                  type="text"
                  v-model="profileForm.company_name"
                  required
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-white text-slate-900 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-brand-300"
                />
              </div>
              <div>
                <label for="profile_phone" class="block text-sm font-semibold text-slate-700 mb-1">Corporate Phone</label>
                <input
                  id="profile_phone"
                  type="tel"
                  v-model="profileForm.phone"
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-white text-slate-900 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-brand-300"
                />
              </div>
              <div class="sm:col-span-2">
                <label for="profile_address" class="block text-sm font-semibold text-slate-700 mb-1">Company Address</label>
                <input
                  id="profile_address"
                  type="text"
                  v-model="profileForm.address"
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-white text-slate-900 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-brand-300"
                  placeholder="e.g. 123 Corporate Blvd, Hanoi"
                />
              </div>
            </div>

            <!-- Nút điều hướng Lưu / Hủy -->
            <div class="pt-4 border-t border-slate-100 flex justify-end space-x-3">
              <button 
                type="button" 
                @click="isEditing = false"
                class="px-5 py-2.5 border border-slate-200 text-sm font-semibold rounded-lg text-slate-700 bg-white hover:bg-slate-50 transition-all"
              >
                Cancel
              </button>
              <button
                type="submit"
                :disabled="isSavingProfile"
                class="px-6 py-2.5 border border-transparent text-sm font-semibold rounded-lg text-white bg-brand-600 hover:bg-brand-500 focus:ring-2 focus:ring-brand-500 disabled:opacity-50 disabled:cursor-not-allowed shadow-sm transition-all duration-150"
              >
                <span v-if="isSavingProfile" class="flex items-center space-x-2">
                  <svg class="animate-spin h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  <span>Saving Profile...</span>
                </span>
                <span v-else>Save Changes</span>
              </button>
            </div>
          </form>
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
