<template>
  <!-- Section 2: Profile Settings -->
  <div v-show="activeSection === 'profile'" class="max-w-6xl mx-auto space-y-6">
    <!-- ======================================================== -->
    <!-- TRẠNG THÁI 1: HIỂN THỊ DỮ LIỆU CŨ (VIEW MODE) -->
    <!-- ======================================================== -->
    <div v-if="!isEditing" class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      
      <!-- CỘT TRÁI (1/3): Thẻ Thông tin Sinh viên & Badge Xác thực -->
      <div class="lg:col-span-1 space-y-6">
        <div class="rounded-3xl border border-indigo-500/20 bg-slate-900/90 p-6 flex flex-col items-center text-center space-y-4 shadow-xl shadow-indigo-950/30 backdrop-blur-xl">
          <div class="relative">
            <img 
              :src="profileForm.avatar_url ? getMediaUrl(profileForm.avatar_url) : 'https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?auto=format&fit=crop&w=150&h=150&q=80'" 
              class="w-28 h-28 rounded-full object-cover border-4 border-indigo-500/30 shadow-xl shadow-indigo-500/20"
            />
            <span class="absolute bottom-1 right-1 h-5 w-5 rounded-full bg-emerald-400 ring-4 ring-slate-900 flex items-center justify-center text-[10px] text-slate-950 font-black">
              ✓
            </span>
          </div>

          <div class="space-y-1.5 w-full">
            <div class="inline-flex items-center gap-1.5 rounded-full bg-emerald-400/10 px-3 py-1 text-xs font-black text-emerald-300 ring-1 ring-emerald-400/30">
              <span>Tài Khoản Sinh Viên Xác Thực</span>
            </div>
            <h3 class="text-2xl font-extrabold text-white tracking-tight pt-1">{{ profileForm.full_name || 'Hồ sơ Sinh viên' }}</h3>
            <p class="text-xs font-semibold text-slate-400">Ứng viên Tìm việc Nhanh QuickWork</p>
          </div>

          <div class="w-full pt-3 border-t border-indigo-500/10 space-y-2.5 text-xs text-slate-300 font-semibold text-left">
            <div class="flex items-center justify-between bg-slate-950/60 px-3.5 py-2.5 rounded-xl border border-indigo-500/10">
              <span class="text-slate-400">Số điện thoại:</span>
              <span class="text-white font-bold">{{ profileForm.phone || 'Chưa cập nhật' }}</span>
            </div>
            <div class="flex items-center justify-between bg-slate-950/60 px-3.5 py-2.5 rounded-xl border border-indigo-500/10">
              <span class="text-slate-400">Giới tính:</span>
              <span class="text-white font-bold">{{ profileForm.gender || 'Chưa chọn' }}</span>
            </div>
          </div>

          <!-- Nút bấm Chuyển sang Edit Mode -->
          <button
            @click="isEditing = true"
            class="w-full py-2.5 px-4 rounded-xl text-xs font-extrabold text-indigo-200 border border-indigo-500/30 bg-indigo-500/10 hover:bg-indigo-500/20 hover:text-white transition-all shadow-md"
          >
            ✏️ Chỉnh sửa Hồ sơ & Tải CV
          </button>
        </div>

        <!-- Card Chỉ số Chuẩn hóa ATS -->
        <div class="rounded-3xl border border-indigo-500/20 bg-gradient-to-b from-indigo-950/30 to-slate-900/90 p-5 shadow-xl space-y-3">
          <div class="flex items-center justify-between">
            <span class="text-xs font-black uppercase tracking-wider text-indigo-300">Điểm Đạt Chuẩn ATS CV</span>
            <span class="rounded-full bg-emerald-400/10 px-2.5 py-0.5 text-xs font-black text-emerald-300 ring-1 ring-emerald-400/30">
              85/100 Điểm
            </span>
          </div>
          <div class="h-2 w-full overflow-hidden rounded-full bg-slate-800">
            <div class="h-full rounded-full bg-gradient-to-r from-indigo-500 to-emerald-400 w-[85%]"></div>
          </div>
          <p class="text-[11px] font-medium text-slate-400 leading-relaxed">
            Hồ sơ đạt tiêu chuẩn cao giúp tăng 3.5x tỷ lệ Doanh nghiệp duyệt đơn ứng tuyển ca làm.
          </p>
        </div>
      </div>

      <!-- CỘT PHẢI (2/3): Skills Stack, CV Attachment & AI CV Evaluator Banner -->
      <div class="lg:col-span-2 space-y-6">
        
        <!-- Khối AI Đánh giá & Chấm điểm CV (Nổi bật nhất) -->
        <div class="rounded-3xl border border-indigo-500/30 bg-gradient-to-r from-slate-950 via-indigo-950/60 to-slate-950 p-6 shadow-2xl shadow-indigo-950/40 ring-1 ring-indigo-500/20">
          <div class="flex flex-col sm:flex-row items-center justify-between gap-5">
            <div class="flex items-center space-x-4">
              <div class="h-14 w-14 rounded-2xl bg-gradient-to-tr from-indigo-500 via-blue-600 to-emerald-400 flex items-center justify-center text-white font-black text-2xl shadow-lg shadow-indigo-500/30 flex-shrink-0">
                🤖
              </div>
              <div>
                <div class="flex items-center gap-2">
                  <h4 class="text-lg font-extrabold text-white tracking-tight">AI CV Evaluator & ATS Audit</h4>
                  <span class="inline-flex rounded-full bg-emerald-400/10 px-2.5 py-0.5 text-[10px] font-black uppercase text-emerald-300 ring-1 ring-emerald-400/30">Gemini AI Engine</span>
                </div>
                <p class="text-xs text-slate-300 font-semibold mt-1 leading-relaxed">Phân tích CV tự động, phát hiện lỗi trình bày & tạo bản tóm tắt ấn tượng cho Nhà tuyển dụng</p>
              </div>
            </div>
            <div class="flex flex-col sm:flex-row gap-2 w-full sm:w-auto flex-shrink-0">
              <button
                v-if="aiResult"
                type="button"
                @click="openOrRunAiEvaluation(false)"
                class="inline-flex items-center justify-center rounded-xl bg-emerald-500/20 text-emerald-300 border border-emerald-500/30 px-5 py-3 text-xs font-black hover:bg-emerald-500/30 transition-all active:scale-95 shadow-lg"
              >
                📊 Xem Kết Quả Phân Tích Đã Lưu
              </button>

              <button
                type="button"
                @click="openOrRunAiEvaluation(true)"
                :disabled="isEvaluating"
                class="inline-flex items-center justify-center rounded-xl bg-gradient-to-r from-indigo-500 via-blue-600 to-emerald-500 px-6 py-3 text-xs font-black text-white shadow-xl shadow-indigo-500/25 transition-all hover:from-indigo-400 hover:to-emerald-400 active:scale-95 disabled:opacity-60"
              >
                <svg v-if="isEvaluating" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                <span>{{ isEvaluating ? 'Đang phân tích AI...' : (aiResult ? '✨ Phân Tích Lại CV' : '✨ Phân Tích CV Bằng AI') }}</span>
              </button>
            </div>
          </div>
        </div>

        <!-- Khối File CV Đính Kèm -->
        <div class="rounded-3xl border border-indigo-500/20 bg-slate-900/90 p-6 space-y-4 shadow-xl">
          <div class="flex items-center justify-between">
            <h4 class="text-xs font-black text-indigo-300 uppercase tracking-wider">File Hồ Sơ CV Đính Kèm</h4>
            <span class="text-[11px] font-bold text-slate-400">Định dạng PDF</span>
          </div>
          
          <div v-if="profileForm.cv_url" class="flex items-center justify-between bg-emerald-500/10 border border-emerald-500/20 p-4 rounded-2xl">
            <div class="flex items-center space-x-3.5">
              <span class="text-3xl">📄</span>
              <div>
                <p class="text-sm font-extrabold text-emerald-200">CV Đang Hoạt Động Trên Hệ Thống</p>
                <p class="text-xs font-medium text-emerald-400/90 mt-0.5">Sẵn sàng ứng tuyển tất cả các công việc trên QuickWork</p>
              </div>
            </div>
            <a 
              :href="getMediaUrl(profileForm.cv_url)" 
              target="_blank" 
              rel="noopener noreferrer"
              class="px-4 py-2 bg-emerald-400 text-slate-950 font-black text-xs rounded-xl hover:bg-emerald-300 transition-all shadow-md shadow-emerald-500/20 flex items-center gap-1.5"
            >
              <span>👁️ Xem trước file CV</span>
            </a>
          </div>

          <div v-else class="text-center py-8 border-2 border-dashed border-indigo-500/20 rounded-2xl bg-slate-950/60">
            <p class="text-sm text-slate-300 font-bold">Chưa có File CV nào được đính kèm</p>
            <p class="text-xs text-slate-400 mt-1">Bấm "Chỉnh sửa hồ sơ" để tải lên file CV bản mềm (PDF) của bạn.</p>
          </div>
        </div>

        <!-- Khối Kỹ năng Chuyên môn -->
        <div class="rounded-3xl border border-indigo-500/20 bg-slate-900/90 p-6 space-y-4 shadow-xl">
          <h4 class="text-xs font-black text-indigo-300 uppercase tracking-wider">Kỹ Năng & Công Nghệ Chuyên Môn (Core Skills)</h4>
          <div v-if="skillsArray.length > 0" class="flex flex-wrap gap-2">
            <span 
              v-for="skill in skillsArray" 
              :key="skill"
              class="px-3.5 py-1.5 bg-indigo-500/10 text-indigo-200 text-xs font-extrabold rounded-xl border border-indigo-500/20 shadow-sm"
            >
              {{ skill }}
            </span>
          </div>
          <p v-else class="text-xs text-slate-400 italic">Chưa bổ sung danh sách kỹ năng. Bấm "Chỉnh sửa hồ sơ" để điền stack công nghệ của bạn.</p>
        </div>

      </div>
    </div>

    <!-- ======================================================== -->
    <!-- TRẠNG THÁI 2: FORM CHỈNH SỬA (EDIT MODE) -->
    <!-- ======================================================== -->
    <div v-else class="rounded-3xl border border-indigo-500/20 bg-slate-900/95 p-6 sm:p-8 shadow-2xl backdrop-blur-xl">
      <div class="flex items-center justify-between border-b border-indigo-500/15 pb-4 mb-6">
        <h3 class="text-lg font-extrabold text-white">Chỉnh Sửa Thông Tin Hồ Sơ Sinh Viên</h3>
        <button @click="isEditing = false" class="text-xs font-bold text-slate-400 hover:text-white">✕ Đóng</button>
      </div>

      <form @submit.prevent="handleUpdateProfile" class="space-y-6">
        
        <!-- Upload Ảnh Avatar -->
        <div class="flex items-center space-x-6 bg-slate-950/80 p-4 rounded-2xl border border-indigo-500/15">
          <img 
            :src="avatarPreview || profileForm.avatar_url || 'https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?auto=format&fit=crop&w=150&h=150&q=80'" 
            class="w-20 h-20 rounded-full object-cover border-2 border-indigo-500/30 shadow-md"
          />
          <div>
            <input type="file" id="profile_avatar_file" accept="image/*" class="hidden" @change="onAvatarFileChange" />
            <label for="profile_avatar_file" class="px-4 py-2 text-xs font-extrabold text-indigo-200 bg-indigo-500/15 border border-indigo-500/30 rounded-xl shadow-sm hover:bg-indigo-500/25 cursor-pointer transition-all">
              Chọn ảnh đại diện mới
            </label>
            <p class="text-[11px] font-medium text-slate-400 mt-2">Hỗ trợ JPG hoặc PNG. Kích thước tối đa 2MB.</p>
          </div>
        </div>

        <!-- Các input chữ -->
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
          <div>
            <label for="profile_fullname" class="block text-xs font-extrabold text-indigo-300 uppercase tracking-wider mb-2">Họ và Tên</label>
            <input id="profile_fullname" type="text" v-model="profileForm.full_name" required class="block w-full px-4 py-2.5 border border-indigo-500/20 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-indigo-400 font-medium" />
          </div>
          <div>
            <label for="profile_phone" class="block text-xs font-extrabold text-indigo-300 uppercase tracking-wider mb-2">Số Điện Thoại</label>
            <input id="profile_phone" type="tel" v-model="profileForm.phone" class="block w-full px-4 py-2.5 border border-indigo-500/20 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-indigo-400 font-medium" />
          </div>
          <div>
            <label for="profile_gender" class="block text-xs font-extrabold text-indigo-300 uppercase tracking-wider mb-2">Giới Tính</label>
            <select id="profile_gender" v-model="profileForm.gender" class="block w-full px-4 py-2.5 border border-indigo-500/20 rounded-xl text-xs bg-slate-950 text-slate-100 focus:outline-none focus:ring-2 focus:ring-indigo-400 font-medium">
              <option value="">Chọn giới tính</option>
              <option value="Male">Nam</option>
              <option value="Female">Nữ</option>
              <option value="Other">Khác</option>
            </select>
          </div>
        </div>

        <!-- Kỹ năng -->
        <div>
          <label for="profile_skills" class="block text-xs font-extrabold text-indigo-300 uppercase tracking-wider mb-2">Kỹ Năng Chuyên Môn (Phân cách bằng dấu phẩy)</label>
          <input id="profile_skills" type="text" v-model="skillsText" class="block w-full px-4 py-2.5 border border-indigo-500/20 rounded-xl text-xs bg-slate-950 text-slate-100 focus:outline-none focus:ring-2 focus:ring-indigo-400 font-medium" placeholder="Golang, VueJS, ReactJS, Python, SQL" />
        </div>

        <!-- Upload File CV -->
        <div class="space-y-2">
          <label class="block text-xs font-extrabold text-indigo-300 uppercase tracking-wider">File Hồ Sơ CV (Định dạng PDF)</label>
          <div class="border-2 border-dashed border-indigo-500/25 hover:border-indigo-400 rounded-2xl p-6 transition-colors text-center relative bg-slate-950/80">
            <input type="file" id="profile_cv_file" accept=".pdf" class="absolute inset-0 w-full h-full opacity-0 cursor-pointer" @change="onCvFileChange" />
            <div class="space-y-1.5">
              <span class="text-2xl">📄</span>
              <p class="text-xs font-bold text-slate-200">
                {{ cvFileSelected ? cvFileSelected.name : 'Bấm vào đây để tải lên file CV PDF mới' }}
              </p>
            </div>
          </div>
        </div>

        <!-- Các nút Lưu / Hủy -->
        <div class="pt-4 border-t border-indigo-500/15 flex justify-end space-x-3">
          <button 
            type="button" 
            @click="isEditing = false"
            class="px-5 py-2.5 border border-white/10 text-xs font-bold rounded-xl text-slate-300 bg-white/5 hover:bg-white/10 transition-all"
          >
            Hủy bỏ
          </button>
          <button
            type="submit"
            :disabled="isSavingProfile"
            class="px-6 py-2.5 text-xs font-extrabold rounded-xl text-white bg-gradient-to-r from-indigo-500 via-blue-600 to-emerald-500 hover:from-indigo-400 hover:to-emerald-400 shadow-md shadow-indigo-500/20 disabled:opacity-50 transition-all"
          >
            <span v-if="isSavingProfile">Đang lưu...</span>
            <span v-else>Lưu Hồ Sơ Mới</span>
          </button>
        </div>
      </form>
    </div>

    <!-- Modal AI CV Review -->
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
import { useMedia } from '~/composables/useMedia'

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

import { onMounted } from 'vue'

const { evaluateCv, getLatestCvEvaluation } = useAi()
const { error } = useToast()
const { getMediaUrl } = useMedia()

const isEvaluating = ref(false)
const isAiModalOpen = ref(false)
const aiResult = ref<EvaluateCvResult | null>(null)

// Tự động nạp báo cáo đánh giá CV mới nhất từ MySQL Database ngay khi mở trang
onMounted(async () => {
  try {
    const latest = await getLatestCvEvaluation()
    if (latest) {
      aiResult.value = latest
    }
  } catch (err) {
    // Chưa có dữ liệu cũ
  }
})

const openOrRunAiEvaluation = async (forceReevaluate = false) => {
  // Nếu đã có báo cáo đã lưu trong DB và người dùng bấm "Xem Kết Quả Đã Lưu" -> Mở ngay tức thì (1-2ms)
  if (aiResult.value && !forceReevaluate) {
    isAiModalOpen.value = true
    return
  }

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
