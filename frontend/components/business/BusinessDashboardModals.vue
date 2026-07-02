<template>
  <div v-if="selectedApp" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm">
  <div class="bg-white rounded-2xl border border-slate-200 shadow-xl max-w-lg w-full overflow-hidden flex flex-col max-h-[90vh] text-left animate-in fade-in zoom-in-95 duration-150">
    
    <div class="p-6 border-b border-slate-100 flex justify-between items-center bg-slate-50">
      <div>
        <h3 class="text-base font-bold text-slate-900">Đánh giá hồ sơ & Gửi Offer</h3>
        <p class="text-xs text-slate-500 font-medium mt-0.5">Ứng viên: <span class="text-slate-700 font-bold">{{ selectedApp.student?.full_name || 'N/A' }}</span></p>
      </div>
      <button @click="closeModal" class="text-slate-400 hover:text-slate-600 font-bold text-xl">&times;</button>
    </div>

    <div class="p-6 space-y-4 overflow-y-auto flex-1">
      <div>
        <label class="text-xs font-bold text-slate-400 uppercase tracking-wider block mb-1">Lời nhắn từ ứng viên (Cover Note)</label>
        <div class="p-3 bg-cyan-400/10 border border-cyan-400/20 rounded-xl text-xs font-medium text-slate-200 whitespace-pre-line">
          {{ selectedApp.cover_note || 'Ứng viên không để lại lời nhắn.' }}
        </div>
      </div>

      <div>
        <label class="text-xs font-bold text-slate-400 uppercase tracking-wider block mb-2">Quyết định</label>
        <div class="grid grid-cols-2 gap-3">
          <button 
            type="button"
            @click="reviewStatus = 'approved'"
            :class="[reviewStatus === 'approved' ? 'border-emerald-500 bg-emerald-50 text-emerald-700 ring-2 ring-emerald-500/20' : 'border-slate-200 bg-white text-slate-700 hover:bg-slate-50']"
            class="py-2 px-4 text-xs font-bold border rounded-xl transition-all flex items-center justify-center space-x-1"
          >
            <span>👍 Chấp nhận & Gửi Offer</span>
          </button>
          <button 
            type="button"
            @click="reviewStatus = 'rejected'"
            :class="[reviewStatus === 'rejected' ? 'border-rose-500 bg-rose-50 text-rose-700 ring-2 ring-rose-500/20' : 'border-slate-200 bg-white text-slate-700 hover:bg-slate-50']"
            class="py-2 px-4 text-xs font-bold border rounded-xl transition-all flex items-center justify-center space-x-1"
          >
            <span>👎 Từ chối đơn</span>
          </button>
        </div>
      </div>

      <div v-if="reviewStatus === 'approved'" class="p-4 bg-emerald-50/30 border border-emerald-100 rounded-xl space-y-3 animate-in fade-in duration-150">
        <h4 class="text-xs font-bold text-emerald-800">✉️ Chi tiết Offer gửi Sinh viên</h4>
        
        <div class="grid grid-cols-2 gap-3">
          <div>
            <label class="text-[11px] font-bold text-slate-600 block mb-1">Mức lương Offer</label>
            <input 
              v-model="offerForm.salary"
              type="text" 
              placeholder="Ví dụ: 15,000,000 VND"
              class="w-full text-xs px-3 py-2 border border-slate-200 rounded-lg bg-white text-slate-800 focus:outline-none focus:border-emerald-500"
            />
          </div>
          <div>
            <label class="text-[11px] font-bold text-slate-600 block mb-1">Ngày đi làm dự kiến</label>
            <input 
              v-model="offerForm.startDate"
              type="date" 
              class="w-full text-xs px-3 py-2 border border-slate-200 rounded-lg bg-white text-slate-800 focus:outline-none focus:border-emerald-500"
            />
          </div>
        </div>

        <div>
          <label class="text-[11px] font-bold text-slate-600 block mb-1">Lời nhắn chào mừng từ HR</label>
          <textarea 
            v-model="offerForm.message"
            rows="2"
            placeholder="Chào mừng bạn gia nhập đội ngũ..."
            class="w-full text-xs p-2.5 border border-slate-200 rounded-lg bg-white text-slate-800 focus:outline-none focus:border-emerald-500"
          ></textarea>
        </div>
      </div>
    </div>

    <div class="p-4 bg-slate-50 border-t border-slate-100 flex items-center justify-end space-x-2">
      <button 
        @click="closeModal" 
        class="px-4 py-2 text-xs font-bold text-slate-600 hover:bg-slate-100 rounded-lg transition-all"
      >
        Đóng
      </button>
      <button 
        @click="submitReview"
        :disabled="!reviewStatus || isSubmitting"
        class="px-4 py-2 text-xs font-bold text-slate-950 bg-cyan-400 hover:bg-cyan-300 disabled:opacity-50 rounded-lg shadow-sm transition-all"
      >
        {{ isSubmitting ? 'Đang gửi...' : 'Xác nhận phản hồi' }}
      </button>
    </div>

  </div>
</div>

<div v-if="selectedCompletionApp" class="fixed inset-0 z-[9999] flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
  <div class="w-full max-w-xl overflow-hidden rounded-2xl border border-white/10 bg-slate-950 shadow-2xl shadow-slate-950/70">
    <div class="border-b border-white/10 bg-slate-900/90 p-5">
      <div class="flex items-start justify-between gap-4">
        <div>
          <p class="text-xs font-bold uppercase tracking-wider text-cyan-300">Hoàn tất công việc</p>
          <h3 class="mt-1 text-xl font-extrabold text-white">Xác nhận hoàn thành & giải ngân</h3>
          <p class="mt-1 text-sm font-medium text-slate-400">
            Sinh viên đã báo hoàn thành. Vui lòng kiểm tra trước khi xác nhận cuối cùng.
          </p>
        </div>
        <button
          @click="closeCompletionModal"
          class="rounded-lg bg-white/5 px-3 py-1.5 text-sm font-bold text-slate-300 hover:bg-white/10 hover:text-white"
        >
          Đóng
        </button>
      </div>
    </div>

    <div class="space-y-4 p-5">
      <div class="grid gap-3 sm:grid-cols-2">
        <div class="rounded-xl border border-white/10 bg-white/[0.04] p-4">
          <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Ứng viên</p>
          <p class="mt-2 text-base font-extrabold text-white">{{ selectedCompletionApp.student?.full_name || 'N/A' }}</p>
          <p class="mt-1 text-sm font-medium text-slate-400">{{ selectedCompletionApp.student?.phone || 'Chưa có số điện thoại' }}</p>
        </div>
        <div class="rounded-xl border border-white/10 bg-white/[0.04] p-4">
          <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Vị trí</p>
          <p class="mt-2 text-base font-extrabold text-white">{{ selectedCompletionApp.job?.title || jobTitleLookup(selectedCompletionApp.job_id) }}</p>
          <p class="mt-1 text-sm font-medium text-slate-400">Mã đơn #{{ selectedCompletionApp.id }}</p>
        </div>
      </div>

      <div class="rounded-xl border border-cyan-400/20 bg-cyan-400/10 p-4">
        <div class="grid gap-3 sm:grid-cols-3">
          <div>
            <p class="text-xs font-bold uppercase tracking-wider text-cyan-200">Offer salary</p>
            <p class="mt-1 text-sm font-extrabold text-white">{{ selectedCompletionApp.offer_salary || 'Thỏa thuận' }}</p>
          </div>
          <div>
            <p class="text-xs font-bold uppercase tracking-wider text-cyan-200">Start date</p>
            <p class="mt-1 text-sm font-extrabold text-white">{{ selectedCompletionApp.offer_start_date || 'N/A' }}</p>
          </div>
          <div>
            <p class="text-xs font-bold uppercase tracking-wider text-cyan-200">Payment</p>
            <p class="mt-1 text-sm font-extrabold text-white">Giả lập giải ngân</p>
          </div>
        </div>
      </div>

      <div class="rounded-xl border border-amber-400/20 bg-amber-400/10 p-4 text-sm font-medium text-amber-100">
        Sau khi xác nhận, hệ thống sẽ chuyển trạng thái đơn sang <span class="font-extrabold">paid</span>, đánh dấu doanh nghiệp đã xác nhận và ghi nhận thời điểm giải ngân.
      </div>
    </div>

    <div class="flex flex-col-reverse gap-2 border-t border-white/10 bg-slate-900/80 p-5 sm:flex-row sm:justify-end">
      <button
        @click="closeCompletionModal"
        class="rounded-lg border border-white/10 px-4 py-2.5 text-sm font-bold text-slate-300 hover:bg-white/10 hover:text-white"
      >
        Kiểm tra lại
      </button>
      <button
        @click="submitBusinessCompletion"
        :disabled="isCompletingJob"
        class="rounded-lg bg-cyan-400 px-4 py-2.5 text-sm font-extrabold text-slate-950 shadow-lg shadow-cyan-500/20 hover:bg-cyan-300 disabled:cursor-not-allowed disabled:opacity-60"
      >
        {{ isCompletingJob ? 'Đang xác nhận...' : 'Xác nhận & giải ngân' }}
      </button>
    </div>
  </div>
</div>

<div 
    v-if="isChatModalOpen && selectedChatApp" 
    class="fixed inset-0 z-[9999] flex items-center justify-center p-4"
  >
    <div 
      class="fixed inset-0 bg-slate-950/80 backdrop-blur-md" 
      @click="isChatModalOpen = false"
    ></div>
    
    <div class="relative w-full max-w-2xl bg-slate-950 border border-slate-800 rounded-2xl shadow-2xl overflow-hidden z-10">
      
      <div class="p-4 bg-slate-900 border-b border-slate-800/80 flex items-center justify-between">
        <div class="flex items-center space-x-3">
          <div class="w-2 h-2 bg-emerald-500 rounded-full animate-pulse"></div>
          <div>
            <h3 class="text-xs font-black text-white uppercase tracking-wider">
              Trò chuyện cùng Ứng viên: {{ selectedChatApp.student?.full_name || 'Sinh viên' }}
            </h3>
            <p class="text-[10px] text-slate-500 font-bold uppercase mt-0.5 tracking-wide">
              Vị trí ứng tuyển: {{ selectedChatApp.job?.title }} — Mã đơn: #{{ selectedChatApp.id }}
            </p>
          </div>
        </div>
        <button 
          @click="isChatModalOpen = false" 
          class="text-slate-400 hover:text-white text-xs font-bold px-2.5 py-1.5 bg-slate-800 hover:bg-slate-700 rounded-xl transition-all"
        >
          Đóng ✕
        </button>
      </div>

      <div class="p-4 bg-slate-900/40 min-h-[400px]">
        <ChatBox 
          v-if="currentBusinessUserId"
          :applicationId="selectedChatApp.id"
          :currentUserId="currentBusinessUserId"
          :targetId="selectedChatApp.student?.user_id || selectedChatApp.student_id"
        />
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { toRefs } from 'vue'
import ChatBox from '~/components/ChatBox.vue'

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
  selectedCompletionApp,
  isCompletingJob,
  closeCompletionModal,
  submitBusinessCompletion,
  closeModal,
  submitReview,
  isChatModalOpen,
  selectedChatApp,
  currentBusinessUserId
} = toRefs(props.state)
</script>
