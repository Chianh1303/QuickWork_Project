<template>
     <!-- Section 2: Profile Settings -->
      <div v-show="activeSection === 'profile'" class="max-w-2xl mx-auto space-y-6">
        
        <!-- Header Tổng -->
        <div class="flex items-center justify-between pb-6 border-b border-slate-200">
          <div>
            <h2 class="text-3xl font-extrabold text-slate-900 tracking-tight">My Professional Profile</h2>
            <p class="mt-1 text-sm text-slate-500 font-medium">View and manage your portfolio details, core skills, and active CV.</p>
          </div>
          <!-- Nút chuyển đổi trạng thái nếu không trong chế độ sửa -->
          <button 
            v-if="!isEditing" 
            @click="isEditing = true"
            class="px-4 py-2 text-sm font-bold text-blue-600 bg-blue-50 hover:bg-blue-100 rounded-lg transition-all flex items-center space-x-1"
          >
            <span>✏️</span> <span>Edit Profile</span>
          </button>
        </div>

        <!-- ======================================================== -->
        <!-- TRẠNG THÁI 1: HIỂN THỊ DỮ LIỆU CŨ (VIEW MODE) -->
        <!-- ======================================================== -->
        <div v-if="!isEditing" class="space-y-6">
          <!-- Thẻ Card thông tin chính -->
          <div class="bg-white rounded-2xl border border-slate-200 shadow-sm p-6 flex flex-col sm:flex-row items-center sm:items-start space-y-4 sm:space-y-0 sm:space-x-6">
            <img 
              :src="profileForm.avatar_url || 'https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?auto=format&fit=crop&w=150&h=150&q=80'" 
              class="w-24 h-24 rounded-full object-cover border-4 border-slate-50 shadow-md"
            />
            <div class="flex-1 text-center sm:text-left space-y-2">
              <h3 class="text-2xl font-bold text-slate-900">{{ profileForm.full_name || 'Your Full Name' }}</h3>
              <div class="flex flex-wrap justify-center sm:justify-start gap-4 text-sm text-slate-600 font-medium">
                <span class="flex items-center">📞 {{ profileForm.phone || 'Not updated yet' }}</span>
                <span class="flex items-center">🚻 {{ profileForm.gender || 'Not specified' }}</span>
              </div>
            </div>
          </div>

          <!-- Khối Kỹ năng -->
          <div class="bg-white rounded-2xl border border-slate-200 shadow-sm p-6 space-y-3">
            <h4 class="text-sm font-bold text-slate-400 uppercase tracking-wider">Core Skills & Technologies</h4>
            <div v-if="skillsArray.length > 0" class="flex flex-wrap gap-2">
              <span 
                v-for="skill in skillsArray" 
                :key="skill"
                class="px-3 py-1 bg-slate-100 text-slate-800 text-xs font-semibold rounded-full border border-slate-200"
              >
                {{ skill }}
              </span>
            </div>
            <p v-else class="text-sm text-slate-400 italic">No skills added yet. Click edit to add your stack.</p>
          </div>

          <!-- Khối File CV Đính Kèm -->
          <div class="bg-white rounded-2xl border border-slate-200 shadow-sm p-6 space-y-3">
            <h4 class="text-sm font-bold text-slate-400 uppercase tracking-wider">Attached Curriculum Vitae (CV)</h4>
            <div v-if="profileForm.cv_url" class="flex items-center justify-between bg-emerald-50 border border-emerald-100 p-4 rounded-xl">
              <div class="flex items-center space-x-3">
                <span class="text-3xl">📄</span>
                <div>
                  <p class="text-sm font-bold text-emerald-950">Your CV is Active</p>
                  <p class="text-xs text-emerald-700">Ready to apply for jobs on QuickWork</p>
                </div>
              </div>
              <a 
                :href="profileForm.cv_url" 
                target="_blank" 
                class="px-4 py-2 bg-emerald-600 hover:bg-emerald-700 text-white text-xs font-bold rounded-lg shadow-sm transition-colors"
              >
                View CV ↗
              </a>
            </div>
            <div v-else class="text-center p-6 border-2 border-dashed border-slate-200 rounded-xl bg-slate-50">
              <p class="text-sm text-slate-500 font-medium">You haven't uploaded any CV PDF file yet.</p>
            </div>
          </div>
        </div>

        <!-- ======================================================== -->
        <!-- TRẠNG THÁI 2: FORM CHỈNH SỬA (EDIT MODE) -->
        <!-- ======================================================== -->
        <div v-else class="bg-white rounded-2xl border border-slate-200 shadow-sm p-6">
          <form @submit.prevent="handleUpdateProfile" class="space-y-6">
            
            <!-- Upload Ảnh -->
            <div class="flex items-center space-x-6 bg-slate-50 p-4 rounded-xl border border-slate-100">
              <img 
                :src="avatarPreview || profileForm.avatar_url || 'https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?auto=format&fit=crop&w=150&h=150&q=80'" 
                class="w-20 h-20 rounded-full object-cover border-2 border-slate-200 shadow-sm"
              />
              <div>
                <input type="file" id="profile_avatar_file" accept="image/*" class="hidden" @change="onAvatarFileChange" />
                <label for="profile_avatar_file" class="px-4 py-2 text-xs font-bold text-slate-700 bg-white border border-slate-300 rounded-lg shadow-sm hover:bg-slate-50 cursor-pointer transition-all">
                  Change Avatar Image
                </label>
                <p class="text-xs text-slate-400 mt-1.5">JPG or PNG. Max 2MB.</p>
              </div>
            </div>

            <!-- Các input chữ -->
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
              <div>
                <label for="profile_fullname" class="block text-sm font-semibold text-slate-700 mb-1">Full Name</label>
                <input id="profile_fullname" type="text" v-model="profileForm.full_name" required class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white" />
              </div>
              <div>
                <label for="profile_phone" class="block text-sm font-semibold text-slate-700 mb-1">Phone Number</label>
                <input id="profile_phone" type="tel" v-model="profileForm.phone" class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white" />
              </div>
              <div>
                <label for="profile_gender" class="block text-sm font-semibold text-slate-700 mb-1">Gender</label>
                <select id="profile_gender" v-model="profileForm.gender" class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white">
                  <option value="">Select Gender</option>
                  <option value="Male">Male</option>
                  <option value="Female">Female</option>
                  <option value="Other">Other</option>
                </select>
              </div>
            </div>

            <!-- Kỹ năng -->
            <div>
              <label for="profile_skills" class="block text-sm font-semibold text-slate-700 mb-1">Skills (Comma-separated)</label>
              <input id="profile_skills" type="text" v-model="skillsText" class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white" placeholder="Go, Vue, JavaScript" />
            </div>

            <!-- Upload File CV -->
            <div class="space-y-2">
              <label class="block text-sm font-semibold text-slate-700">CV File (PDF Only)</label>
              <div class="border-2 border-dashed border-slate-200 hover:border-blue-400 rounded-xl p-5 transition-colors text-center relative bg-slate-50">
                <input type="file" id="profile_cv_file" accept=".pdf" class="absolute inset-0 w-full h-full opacity-0 cursor-pointer" @change="onCvFileChange" />
                <div class="space-y-1">
                  <span class="text-xl">📄</span>
                  <p class="text-sm font-medium text-slate-700">
                    {{ cvFileSelected ? cvFileSelected.name : 'Click to upload your new CV' }}
                  </p>
                </div>
              </div>
            </div>

            <!-- Các nút Lưu / Hủy -->
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
                class="px-6 py-2.5 border border-transparent text-sm font-semibold rounded-lg text-white bg-blue-600 hover:bg-blue-500 focus:ring-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed shadow-sm transition-all duration-150"
              >
                <span v-if="isSavingProfile">Saving Profile...</span>
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
