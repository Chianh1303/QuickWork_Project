<template>
     <!-- Section 2: Profile Settings -->
      <div v-show="activeSection === 'profile'" class="max-w-2xl mx-auto space-y-6">
        <!-- ======================================================== -->
        <!-- TRẠNG THÁI 1: HIỂN THỊ DỮ LIỆU CŨ (VIEW MODE) -->
        <!-- ======================================================== -->
        <div v-if="!isEditing" class="space-y-6">
          <!-- Thẻ Card thông tin chính -->
          <div class="rounded-2xl border border-white/10 bg-slate-900/82 shadow-lg shadow-slate-950/25 p-6 flex flex-col sm:flex-row items-center sm:items-start space-y-4 sm:space-y-0 sm:space-x-6 backdrop-blur">
            <img 
              :src="profileForm.avatar_url || 'https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?auto=format&fit=crop&w=150&h=150&q=80'" 
              class="w-24 h-24 rounded-full object-cover border-4 border-cyan-400/20 shadow-md"
            />
            <div class="flex-1 text-center sm:text-left space-y-2">
              <h3 class="text-2xl font-bold text-white">{{ profileForm.full_name || 'Your Full Name' }}</h3>
              <div class="flex flex-wrap justify-center sm:justify-start gap-4 text-sm text-slate-300 font-medium">
                <span class="flex items-center gap-1.5">
                  <svg class="h-4 w-4 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 002.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106a1.125 1.125 0 00-1.173.417l-.97 1.293a1.125 1.125 0 01-1.21.38 12.035 12.035 0 01-7.143-7.143 1.125 1.125 0 01.38-1.21l1.293-.97c.36-.27.527-.734.417-1.173L6.963 3.102A1.125 1.125 0 005.872 2.25H4.5A2.25 2.25 0 002.25 4.5v2.25z" />
                  </svg>
                  {{ profileForm.phone || 'Not updated yet' }}
                </span>
                <span class="flex items-center gap-1.5">
                  <svg class="h-4 w-4 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.75 6a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0zM4.5 20.25a8.25 8.25 0 0115 0" />
                  </svg>
                  {{ profileForm.gender || 'Not specified' }}
                </span>
              </div>
            </div>
          </div>

          <!-- Khối Kỹ năng -->
          <div class="rounded-2xl border border-white/10 bg-slate-900/82 shadow-lg shadow-slate-950/25 p-6 space-y-3 backdrop-blur">
            <h4 class="text-sm font-bold text-cyan-200/80 uppercase tracking-wider">Core Skills & Technologies</h4>
            <div v-if="skillsArray.length > 0" class="flex flex-wrap gap-2">
              <span 
                v-for="skill in skillsArray" 
                :key="skill"
                class="px-3 py-1 bg-cyan-400/10 text-cyan-100 text-xs font-semibold rounded-full border border-cyan-400/20"
              >
                {{ skill }}
              </span>
            </div>
            <p v-else class="text-sm text-slate-400 italic">No skills added yet. Click edit to add your stack.</p>
          </div>

          <!-- Khối File CV Đính Kèm -->
          <div class="rounded-2xl border border-white/10 bg-slate-900/82 shadow-lg shadow-slate-950/25 p-6 space-y-3 backdrop-blur">
            <h4 class="text-sm font-bold text-cyan-200/80 uppercase tracking-wider">File Hồ sơ CV Cá nhân</h4>
            <div v-if="profileForm.cv_url" class="flex items-center justify-between bg-emerald-400/10 border border-emerald-400/20 p-4 rounded-xl">
              <div class="flex items-center space-x-3">
                <span class="text-3xl">📄</span>
                <div>
                  <p class="text-sm font-bold text-emerald-100">CV đang hoạt động</p>
                  <p class="text-xs text-emerald-300">Đã sẵn sàng ứng tuyển các công việc trên QuickWork</p>
                </div>
              </div>
              <a 
                :href="profileForm.cv_url" 
                target="_blank" 
                class="px-4 py-2 bg-emerald-400 text-slate-950 font-bold text-xs rounded-lg hover:bg-emerald-300 transition-colors shadow-sm"
              >
                Xem CV
              </a>
            </div>
            <div v-else class="text-center py-6 border-2 border-dashed border-white/10 rounded-xl bg-slate-950/50">
              <p class="text-sm text-slate-400 font-semibold">Chưa có CV nào được đính kèm</p>
              <p class="text-xs text-slate-500 mt-1">Bấm "Chỉnh sửa hồ sơ" bên dưới để tải lên CV của bạn.</p>
            </div>
          </div>

          <!-- Khối AI Đánh giá & Chấm điểm CV (Mục AI mới) -->
          <div class="rounded-2xl border border-cyan-400/30 bg-gradient-to-br from-cyan-950/40 via-slate-900/80 to-slate-950/90 shadow-xl shadow-cyan-950/30 p-6 backdrop-blur ring-1 ring-cyan-400/20">
            <div class="flex flex-col sm:flex-row items-center justify-between gap-4">
              <div class="flex items-center space-x-4">
                <div class="h-12 w-12 rounded-2xl bg-gradient-to-tr from-cyan-400 to-emerald-400 flex items-center justify-center text-slate-950 font-black text-2xl shadow-lg shadow-cyan-500/30">
                  🤖
                </div>
                <div>
                  <h4 class="text-base font-extrabold text-white flex items-center gap-2">
                    AI CV Evaluator
                    <span class="inline-flex rounded-full bg-cyan-400/20 px-2.5 py-0.5 text-[10px] font-black uppercase text-cyan-300 ring-1 ring-cyan-400/30">Mới</span>
                  </h4>
                  <p class="text-xs text-slate-300 font-semibold mt-0.5">Phân tích CV, phát hiện điểm yếu & nhận tóm tắt AI ấn tượng</p>
                </div>
              </div>
              <button
                type="button"
                @click="runAiCvEvaluation"
                :disabled="isEvaluating"
                class="w-full sm:w-auto inline-flex items-center justify-center rounded-xl bg-cyan-400 px-5 py-3 text-xs font-black text-slate-950 shadow-lg shadow-cyan-500/25 transition-all hover:bg-cyan-300 active:scale-95 disabled:opacity-60"
              >
                <svg v-if="isEvaluating" class="animate-spin -ml-1 mr-2 h-4 w-4 text-slate-950" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                <span>{{ isEvaluating ? 'Đang phân tích AI...' : '✨ Phân Tích CV Bằng AI' }}</span>
              </button>
            </div>
          </div>
        </div>

        <!-- ======================================================== -->
        <!-- TRẠNG THÁI 2: FORM CHỈNH SỬA (EDIT MODE) -->
        <!-- ======================================================== -->
        <div v-else class="rounded-2xl border border-white/10 bg-slate-900/82 shadow-lg shadow-slate-950/25 p-6 backdrop-blur">
          <form @submit.prevent="handleUpdateProfile" class="space-y-6">
            
            <!-- Upload Ảnh -->
            <div class="flex items-center space-x-6 bg-slate-950/60 p-4 rounded-xl border border-white/10">
              <img 
                :src="avatarPreview || profileForm.avatar_url || 'https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?auto=format&fit=crop&w=150&h=150&q=80'" 
                class="w-20 h-20 rounded-full object-cover border-2 border-slate-200 shadow-sm"
              />
              <div>
                <input type="file" id="profile_avatar_file" accept="image/*" class="hidden" @change="onAvatarFileChange" />
                <label for="profile_avatar_file" class="px-4 py-2 text-xs font-bold text-cyan-100 bg-cyan-400/10 border border-cyan-400/20 rounded-lg shadow-sm hover:bg-cyan-400/15 cursor-pointer transition-all">
                  Change Avatar Image
                </label>
                <p class="text-xs text-slate-400 mt-1.5">JPG or PNG. Max 2MB.</p>
              </div>
            </div>

            <!-- Các input chữ -->
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
              <div>
                <label for="profile_fullname" class="block text-sm font-semibold text-slate-200 mb-1">Full Name</label>
                <input id="profile_fullname" type="text" v-model="profileForm.full_name" required class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-white text-slate-900 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-brand-300" />
              </div>
              <div>
                <label for="profile_phone" class="block text-sm font-semibold text-slate-200 mb-1">Phone Number</label>
                <input id="profile_phone" type="tel" v-model="profileForm.phone" class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-white text-slate-900 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-brand-300" />
              </div>
              <div>
                <label for="profile_gender" class="block text-sm font-semibold text-slate-200 mb-1">Gender</label>
                <select id="profile_gender" v-model="profileForm.gender" class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-white text-slate-900 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-brand-300">
                  <option value="">Select Gender</option>
                  <option value="Male">Male</option>
                  <option value="Female">Female</option>
                  <option value="Other">Other</option>
                </select>
              </div>
            </div>

            <!-- Kỹ năng -->
            <div>
              <label for="profile_skills" class="block text-sm font-semibold text-slate-200 mb-1">Skills (Comma-separated)</label>
              <input id="profile_skills" type="text" v-model="skillsText" class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-white text-slate-900 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-brand-300" placeholder="Go, Vue, JavaScript" />
            </div>

            <!-- Upload File CV -->
            <div class="space-y-2">
              <label class="block text-sm font-semibold text-slate-200">CV File (PDF Only)</label>
              <div class="border-2 border-dashed border-white/10 hover:border-cyan-400 rounded-xl p-5 transition-colors text-center relative bg-slate-950/60">
                <input type="file" id="profile_cv_file" accept=".pdf" class="absolute inset-0 w-full h-full opacity-0 cursor-pointer" @change="onCvFileChange" />
                <div class="space-y-1">
                  <span class="text-xl">📄</span>
                  <p class="text-sm font-medium text-slate-300">
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
                class="px-5 py-2.5 border border-white/10 text-sm font-semibold rounded-lg text-slate-200 bg-white/10 hover:bg-white/15 transition-all"
              >
                Hủy bỏ
              </button>
              <button
                type="submit"
                :disabled="isSavingProfile"
                class="px-6 py-2.5 border border-transparent text-sm font-semibold rounded-lg text-slate-950 bg-cyan-400 hover:bg-cyan-300 focus:ring-2 focus:ring-cyan-400 disabled:opacity-50 disabled:cursor-not-allowed shadow-sm transition-all duration-150"
              >
                <span v-if="isSavingProfile">Đang lưu...</span>
                <span v-else>Lưu thay đổi</span>
              </button>
            </div>
          </form>
        </div>

        <!-- Modal AI CV Review (Đưa ra ngoài khối v-if/v-else để hiển thị ở cả 2 trạng thái) -->
        <AiCvReviewModal
          :is-open="isAiModalOpen"
          :result="aiResult"
          @close="isAiModalOpen = false"
        />
      </div>
</template>

<script setup lang="ts">
import { ref, toRefs } from 'vue'
import AiCvReviewModal from '~/components/student/AiCvReviewModal.vue'
import { useAi, type EvaluateCvResult } from '~/composables/useAi'
import { useToast } from '~/composables/useToast'

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

const { evaluateCv } = useAi()
const { error } = useToast()

const isEvaluating = ref(false)
const isAiModalOpen = ref(false)
const aiResult = ref<EvaluateCvResult | null>(null)

const runAiCvEvaluation = async () => {
  isEvaluating.value = true
  try {
    const rawSkills = skillsArray?.value || []
    const formVal = profileForm?.value || {}
    const res = await evaluateCv({
      full_name: formVal.full_name,
      phone: formVal.phone,
      gender: formVal.gender,
      skills: Array.isArray(rawSkills) ? rawSkills : [],
      cv_url: formVal.cv_url
    })
    aiResult.value = res
    isAiModalOpen.value = true
  } catch (err: any) {
    console.error('AI CV Evaluation Error:', err)
    const msg = err?.data?.message || err?.message || 'Không thể kết nối đến hệ thống phân tích AI. Vui lòng thử lại sau.'
    error(msg)
  } finally {
    isEvaluating.value = false
  }
}
</script>
