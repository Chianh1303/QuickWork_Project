<template>
  <!-- Section 2: Profile Settings -->
  <div v-show="activeSection === 'profile'" class="max-w-6xl mx-auto space-y-6">
    <!-- ======================================================== -->
    <!-- TRẠNG THÁI 1: HIỂN THỊ DỮ LIỆU CŨ (VIEW MODE) -->
    <!-- ======================================================== -->
    <div v-if="!isEditing" class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      
      <!-- CỘT TRÁI (1/3): Thẻ Thông tin Sinh viên & Badge Xác thực -->
      <div class="lg:col-span-1 space-y-6">
        <div class="rounded-3xl border border-cyan-500/20 bg-slate-900/90 p-6 flex flex-col items-center text-center space-y-4 shadow-xl shadow-cyan-950/30 backdrop-blur-xl">
          <div class="relative">
            <img 
              :src="profileForm.avatar_url ? getMediaUrl(profileForm.avatar_url) : 'https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?auto=format&fit=crop&w=150&h=150&q=80'" 
              class="w-28 h-28 rounded-full object-cover border-4 border-cyan-500/30 shadow-xl shadow-cyan-500/20"
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

          <div class="w-full pt-3 border-t border-cyan-500/10 space-y-2.5 text-xs text-slate-300 font-semibold text-left">
            <div class="flex items-center justify-between bg-slate-950/60 px-3.5 py-2.5 rounded-xl border border-cyan-500/10">
              <span class="text-slate-400">Số điện thoại:</span>
              <span class="text-white font-bold">{{ profileForm.phone || 'Chưa cập nhật' }}</span>
            </div>
            <div class="flex items-center justify-between bg-slate-950/60 px-3.5 py-2.5 rounded-xl border border-cyan-500/10">
              <span class="text-slate-400">Giới tính:</span>
              <span class="text-white font-bold">{{ profileForm.gender || 'Chưa chọn' }}</span>
            </div>
          </div>

          <!-- Nút bấm Chuyển sang Edit Mode -->
          <button
            @click="isEditing = true"
            class="w-full py-2.5 px-4 rounded-xl text-xs font-extrabold text-cyan-200 border border-cyan-500/30 bg-cyan-500/10 hover:bg-cyan-500/20 hover:text-white transition-all shadow-md cursor-pointer"
          >
            Chỉnh sửa Hồ sơ & Tải CV
          </button>

          <!-- Nút Đăng xuất nổi bật trên Mobile & Desktop -->
          <button
            v-if="handleLogout"
            type="button"
            @click="handleLogout"
            class="w-full py-2.5 px-4 rounded-xl text-xs font-extrabold text-rose-300 border border-rose-500/30 bg-rose-500/10 hover:bg-rose-500/20 hover:text-white transition-all shadow-md flex items-center justify-center gap-2 cursor-pointer mt-2"
          >
            <svg class="h-4 w-4 text-rose-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
            </svg>
            <span>Đăng xuất tài khoản</span>
          </button>
        </div>

        <!-- Card Chỉ số Chuẩn hóa ATS -->
        <div class="rounded-3xl border border-cyan-500/20 bg-slate-900/90 p-6 space-y-3.5 shadow-xl shadow-cyan-950/30 backdrop-blur-xl">
          <div class="flex items-center justify-between">
            <span class="text-xs font-black uppercase tracking-wider text-cyan-300">Độ Hoàn Thiện Hồ Sơ (ATS)</span>
            <span class="text-sm font-black text-cyan-400">{{ profileReadiness }}%</span>
          </div>
          <div class="w-full h-3 rounded-full bg-slate-950 p-0.5 border border-cyan-500/20">
            <div 
              class="h-full rounded-full bg-gradient-to-r from-cyan-500 via-blue-500 to-emerald-400 transition-all duration-500 shadow-md shadow-cyan-500/30"
              :style="{ width: `${profileReadiness}%` }"
            ></div>
          </div>
          <p class="text-[11px] font-semibold text-slate-400 leading-relaxed">
            Hồ sơ đầy đủ thông tin kỹ năng & file CV giúp tăng <strong class="text-cyan-300">85% cơ hội</strong> trúng tuyển nhanh.
          </p>
        </div>
      </div>

      <!-- CỘT PHẢI (2/3): Xem thông tin CV, Kỹ năng & Nút bấm xem báo cáo AI -->
      <div class="lg:col-span-2 space-y-6">
        
        <!-- Khối File CV Đã Lưu -->
        <div class="rounded-3xl border border-cyan-500/20 bg-slate-900/90 p-6 space-y-4 shadow-xl shadow-cyan-950/30 backdrop-blur-xl">
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-3">
              <div class="flex h-10 w-10 items-center justify-center rounded-xl bg-cyan-500/10 text-cyan-300 border border-cyan-500/20">
                📄
              </div>
              <div>
                <h4 class="text-sm font-extrabold text-white">File Hồ Sơ CV Ứng Tuyển</h4>
                <p class="text-xs font-medium text-slate-400">Bản CV được sử dụng để ứng tuyển công việc hàng ngày</p>
              </div>
            </div>

            <button
              type="button"
              @click="runNewAiEvaluation"
              :disabled="isEvaluating"
              class="px-4 py-2.5 bg-gradient-to-r from-cyan-400 to-emerald-400 hover:from-cyan-300 hover:to-emerald-300 text-slate-950 font-black text-xs rounded-xl transition-all shadow-md shadow-cyan-500/20 flex items-center gap-2 cursor-pointer disabled:opacity-50"
            >
              <svg v-if="isEvaluating" class="animate-spin h-3.5 w-3.5" viewBox="0 0 24 24" fill="none">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8H4z"></path>
              </svg>
              <span>{{ isEvaluating ? '🤖 Đang Phân Tích Lượt Mới...' : '🤖 Phân Tích CV Bằng AI' }}</span>
            </button>
          </div>

          <!-- Chi tiết file CV -->
          <div v-if="profileForm.cv_url" class="flex flex-col sm:flex-row items-start sm:items-center justify-between p-4 rounded-2xl bg-slate-950/80 border border-cyan-500/20 gap-3">
            <div class="flex items-center space-x-3 min-w-0">
              <svg class="h-8 w-8 text-rose-400 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                <path d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4z" />
              </svg>
              <div class="min-w-0">
                <p class="text-xs font-extrabold text-white truncate">Hoso_CV_SinhVien.pdf</p>
                <p class="text-[11px] font-semibold text-emerald-300 mt-0.5">✓ Đã tải lên thành công</p>
              </div>
            </div>
            <a 
              :href="getMediaUrl(profileForm.cv_url)" 
              target="_blank" 
              class="w-full sm:w-auto px-4 py-2 bg-slate-800 hover:bg-slate-700 text-cyan-200 text-xs font-bold rounded-xl border border-cyan-500/20 transition-all text-center"
            >
              <span>Xem trước file CV</span>
            </a>
          </div>

          <div v-else class="text-center py-8 border-2 border-dashed border-cyan-500/20 rounded-2xl bg-slate-950/60">
            <p class="text-sm text-slate-300 font-bold">Chưa có File CV nào được đính kèm</p>
            <p class="text-xs text-slate-400 mt-1">Bấm "Chỉnh sửa hồ sơ" để tải lên file CV bản mềm (PDF) của bạn.</p>
          </div>
        </div>

        <!-- Khối Kỹ năng Chuyên môn -->
        <div class="rounded-3xl border border-cyan-500/20 bg-slate-900/90 p-6 space-y-4 shadow-xl">
          <h4 class="text-xs font-black text-cyan-300 uppercase tracking-wider">Kỹ Năng & Công Nghệ Chuyên Môn (Core Skills)</h4>
          <div v-if="skillsArray.length > 0" class="flex flex-wrap gap-2">
            <span 
              v-for="skill in skillsArray" 
              :key="skill"
              class="px-3.5 py-1.5 bg-cyan-500/10 text-cyan-200 text-xs font-extrabold rounded-xl border border-cyan-500/20 shadow-sm"
            >
              {{ skill }}
            </span>
          </div>
          <p v-else class="text-xs text-slate-400 italic">Chưa bổ sung danh sách kỹ năng. Bấm "Chỉnh sửa hồ sơ" để điền stack công nghệ của bạn.</p>
        </div>

        <!-- Khối Báo Cáo Phân Tích CV Đã Lưu (Nằm ngay dưới phần Kỹ Năng) -->
        <div v-if="aiResult" class="rounded-3xl border border-emerald-500/30 bg-emerald-950/30 p-6 shadow-xl flex flex-col sm:flex-row items-center justify-between gap-4">
          <div class="flex items-center space-x-3.5">
            <span class="text-3xl">📊</span>
            <div>
              <h4 class="text-sm font-extrabold text-emerald-200">Kết Quả Phân Tích CV Đã Lưu</h4>
              <p class="text-xs text-emerald-400/90 mt-0.5">Báo cáo đánh giá điểm số ATS & gợi ý cải thiện của AI sẵn sàng xem lại tức thì</p>
            </div>
          </div>
          <button
            type="button"
            @click="openSavedAiEvaluation"
            class="w-full sm:w-auto px-6 py-3 bg-emerald-400 hover:bg-emerald-300 text-slate-950 font-black text-xs rounded-xl transition-all shadow-lg shadow-emerald-500/20 flex items-center justify-center gap-2 flex-shrink-0"
          >
            <span>Xem Kết Quả Chi Tiết</span>
          </button>
        </div>

      </div>
    </div>

    <!-- ======================================================== -->
    <!-- TRẠNG THÁI 2: FORM CHỈNH SỬA (EDIT MODE) -->
    <!-- ======================================================== -->
    <div v-else class="rounded-3xl border border-cyan-500/20 bg-slate-900/95 p-6 sm:p-8 shadow-2xl backdrop-blur-xl">
      <div class="flex items-center justify-between border-b border-cyan-500/15 pb-4 mb-6">
        <h3 class="text-lg font-extrabold text-white">Chỉnh Sửa Thông Tin Hồ Sơ Sinh Viên</h3>
        <button @click="isEditing = false" class="text-xs font-bold text-slate-400 hover:text-white">✕ Đóng</button>
      </div>

      <form @submit.prevent="handleUpdateProfile" class="space-y-6">
        <!-- Upload Ảnh Đại Diện -->
        <div class="flex items-center space-x-6 bg-slate-950/60 p-4 rounded-2xl border border-cyan-500/10">
          <img 
            :src="avatarPreview || (profileForm.avatar_url ? getMediaUrl(profileForm.avatar_url) : 'https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?auto=format&fit=crop&w=150&h=150&q=80')" 
            class="w-20 h-20 rounded-full object-cover border-2 border-cyan-500/30 shadow-md"
          />
          <div>
            <input type="file" id="profile_avatar_file" accept="image/*" class="hidden" @change="onAvatarFileChange" />
            <label for="profile_avatar_file" class="px-4 py-2 text-xs font-extrabold text-cyan-200 bg-cyan-500/15 border border-cyan-500/30 rounded-xl shadow-sm hover:bg-cyan-500/25 cursor-pointer transition-all">
              Chọn ảnh đại diện mới
            </label>
            <p class="text-[11px] font-medium text-slate-400 mt-2">Hỗ trợ JPG hoặc PNG. Kích thước tối đa 2MB.</p>
          </div>
        </div>

        <!-- Các input chữ -->
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
          <div>
            <label for="profile_fullname" class="block text-xs font-extrabold text-cyan-300 uppercase tracking-wider mb-2">Họ và Tên</label>
            <input id="profile_fullname" type="text" v-model="profileForm.full_name" required class="block w-full px-4 py-2.5 border border-cyan-500/20 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium" />
          </div>
          <div>
            <label for="profile_phone" class="block text-xs font-extrabold text-cyan-300 uppercase tracking-wider mb-2">Số Điện Thoại</label>
            <input id="profile_phone" type="tel" v-model="profileForm.phone" class="block w-full px-4 py-2.5 border border-cyan-500/20 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium" />
          </div>
          <div>
            <label for="profile_gender" class="block text-xs font-extrabold text-cyan-300 uppercase tracking-wider mb-2">Giới Tính</label>
            <select id="profile_gender" v-model="profileForm.gender" class="block w-full px-4 py-2.5 border border-cyan-500/20 rounded-xl text-xs bg-slate-950 text-slate-100 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium">
              <option value="">Chọn giới tính</option>
              <option value="Male">Nam</option>
              <option value="Female">Nữ</option>
              <option value="Other">Khác</option>
            </select>
          </div>
        </div>

        <!-- Kỹ năng -->
        <div>
          <label for="profile_skills" class="block text-xs font-extrabold text-cyan-300 uppercase tracking-wider mb-2">Kỹ Năng Chuyên Môn (Phân cách bằng dấu phẩy)</label>
          <input id="profile_skills" type="text" v-model="skillsText" class="block w-full px-4 py-2.5 border border-cyan-500/20 rounded-xl text-xs bg-slate-950 text-slate-100 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium" placeholder="Golang, VueJS, ReactJS, Python, SQL" />
        </div>

        <!-- Upload File CV -->
        <div class="bg-slate-950/60 p-5 rounded-2xl border border-cyan-500/20 space-y-3">
          <label for="profile_cv_file" class="block text-xs font-extrabold text-cyan-300 uppercase tracking-wider">File Hồ Sơ CV (Bản mềm PDF/DOCX)</label>
          <div class="flex items-center space-x-4">
            <input type="file" id="profile_cv_file" accept=".pdf,.doc,.docx" class="hidden" @change="onCvFileChange" />
            <label for="profile_cv_file" class="px-4 py-2.5 text-xs font-extrabold text-cyan-200 bg-cyan-500/15 border border-cyan-500/30 rounded-xl shadow-sm hover:bg-cyan-500/25 cursor-pointer transition-all flex items-center space-x-2">
              <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12" />
              </svg>
              <span>Tải file CV mới từ máy tính</span>
            </label>
            <span v-if="profileForm.cv_url" class="text-xs font-bold text-emerald-300">✓ Đã có file CV</span>
            <span v-else class="text-xs font-medium text-slate-400">Chưa chọn file mới</span>
          </div>
          <p class="text-[11px] font-medium text-slate-400">Định dạng hỗ trợ: .PDF, .DOC, .DOCX. Dung lượng tối đa 5MB.</p>
        </div>

        <!-- Nút bấm Submit -->
        <div class="flex items-center justify-end space-x-3 pt-4 border-t border-cyan-500/15">
          <button
            type="button"
            @click="isEditing = false"
            class="px-5 py-2.5 text-xs font-bold rounded-xl text-slate-300 bg-slate-800 hover:bg-slate-700 transition-colors"
          >
            Hủy bỏ
          </button>
          <button
            type="submit"
            :disabled="isSavingProfile"
            class="px-6 py-2.5 text-xs font-extrabold rounded-xl text-white bg-gradient-to-r from-cyan-500 via-blue-600 to-emerald-500 hover:from-cyan-400 hover:to-emerald-400 shadow-md shadow-cyan-500/20 disabled:opacity-50 transition-all"
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
      :is-evaluating="isEvaluating"
      @reevaluate="runNewAiEvaluation"
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
  handleStudentComplete,
  openReviewModal,
  openManagedApplicationModal,
  handleLogout,
  toast
} = toRefs(props.state)

const { evaluateCv, getLatestCvEvaluation } = useAi()
const { getMediaUrl } = useMedia()
const { error } = useToast()

const isAiModalOpen = ref(false)
const isEvaluating = ref(false)
const aiResult = ref<EvaluateCvResult | null>(null)

// Tải báo cáo AI đã lưu gần nhất khi vừa khởi tạo trang
if (import.meta.client) {
  getLatestCvEvaluation().then((latest) => {
    if (latest) aiResult.value = latest
  }).catch(() => {})
}

// Mở xem báo cáo phân tích CV cũ đã lưu
const openSavedAiEvaluation = () => {
  if (aiResult.value) {
    isAiModalOpen.value = true
  }
}

// Luôn thực hiện phân tích lượt mới bằng AI
const runNewAiEvaluation = async () => {
  const formVal = profileForm?.value || {}
  if (!formVal.cv_url) {
    error('Vui lòng tải lên file CV trước khi thực hiện phân tích bằng AI.')
    return
  }

  isEvaluating.value = true
  try {
    const rawSkills = skillsArray?.value || []
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
