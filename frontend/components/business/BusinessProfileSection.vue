<template>
    <!-- Section 2: Company Profile -->
      <div v-show="activeSection === 'profile'" class="max-w-2xl mx-auto space-y-6">
        
        <!-- Header Tổng -->
        <div class="flex items-center justify-between pb-6 border-b border-slate-200">
          <div>
            <h2 class="text-3xl font-extrabold text-slate-900 tracking-tight">Company Profile</h2>
            <p class="mt-1 text-sm text-slate-500 font-medium">Update public info, logo, and address details.</p>
          </div>
          <!-- Nút chuyển sang chế độ sửa -->
          <button 
            v-if="!isEditing" 
            @click="isEditing = true"
            class="px-4 py-2 text-sm font-bold text-violet-600 bg-violet-50 hover:bg-violet-100 rounded-lg transition-all flex items-center space-x-1"
          >
            <span>✏️</span> <span>Edit Profile</span>
          </button>
        </div>

        <!-- ======================================================== -->
        <!-- TRẠNG THÁI 1: HIỂN THỊ DỮ LIỆU CŨ (VIEW MODE) -->
        <!-- ======================================================== -->
        <div v-if="!isEditing" class="space-y-6">
          <!-- Card thông tin công ty chính -->
          <div class="bg-white rounded-2xl border border-slate-200 shadow-sm p-6 flex flex-col sm:flex-row items-center sm:items-start space-y-4 sm:space-y-0 sm:space-x-6">
            <img 
              :src="profileForm.logo_url || 'https://images.unsplash.com/photo-1486406146926-c627a92ad1ab?auto=format&fit=crop&w=150&h=150&q=80'" 
              class="w-24 h-24 rounded-xl object-cover border-2 border-slate-100 shadow-sm bg-slate-50 p-1"
            />
            <div class="flex-1 text-center sm:text-left space-y-3">
              <h3 class="text-2xl font-bold text-slate-900">{{ profileForm.company_name || 'Your Company Name' }}</h3>
              
              <div class="space-y-2 text-sm text-slate-600 font-medium">
                <p class="flex items-center justify-center sm:justify-start">
                  <span class="mr-2">🏢</span> {{ profileForm.address || 'Address not updated yet' }}
                </p>
                <p class="flex items-center justify-center sm:justify-start">
                  <span class="mr-2">📞</span> {{ profileForm.phone || 'Corporate phone not updated yet' }}
                </p>
              </div>
            </div>
          </div>
        </div>

        <!-- ======================================================== -->
        <!-- TRẠNG THÁI 2: FORM CHỈNH SỬA (EDIT MODE) -->
        <!-- ======================================================== -->
        <div v-else class="bg-white rounded-2xl border border-slate-200 shadow-sm p-6">
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
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-violet-500 focus:bg-white"
                />
              </div>
              <div>
                <label for="profile_phone" class="block text-sm font-semibold text-slate-700 mb-1">Corporate Phone</label>
                <input
                  id="profile_phone"
                  type="tel"
                  v-model="profileForm.phone"
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-violet-500 focus:bg-white"
                />
              </div>
              <div class="sm:col-span-2">
                <label for="profile_address" class="block text-sm font-semibold text-slate-700 mb-1">Company Address</label>
                <input
                  id="profile_address"
                  type="text"
                  v-model="profileForm.address"
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-violet-500 focus:bg-white"
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
                class="px-6 py-2.5 border border-transparent text-sm font-semibold rounded-lg text-white bg-violet-600 hover:bg-violet-500 focus:ring-2 focus:ring-violet-500 disabled:opacity-50 disabled:cursor-not-allowed shadow-sm transition-all duration-150"
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
