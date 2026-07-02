<template>
  <div v-if="selectedJobForApply" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm">
  <div class="bg-white rounded-2xl border border-slate-200 shadow-xl max-w-md w-full overflow-hidden animate-in fade-in zoom-in-95 duration-200">
    
    <div class="p-6 pb-4 border-b border-slate-100">
      <h3 class="text-xl font-bold text-slate-900">Apply Opportunity</h3>
      <p class="text-sm font-medium text-slate-500 mt-1">
        You are applying for: <span class="text-cyan-300 font-semibold">{{ selectedJobForApply.title }}</span>
      </p>
    </div>

    <form @submit.prevent="submitApplication" class="p-6 space-y-4">
      <div>
        <label class="block text-sm font-bold text-slate-700 mb-2">
          Cover Note (Lời nhắn gửi nhà tuyển dụng)
        </label>
        <textarea
          v-model="coverNoteText"
          rows="4"
          maxlength="500"
          class="block w-full px-3 py-2 border border-white/10 rounded-xl text-sm bg-slate-950/70 text-slate-100 focus:outline-none focus:ring-2 focus:ring-cyan-400 placeholder-slate-500 resize-none transition-all"
          placeholder="Giới thiệu ngắn gọn về thế mạnh của bạn hoặc lý do bạn mong muốn ứng tuyển vị trí này..."
        ></textarea>
        <div class="text-right text-xs font-medium text-slate-400 mt-1">
          {{ coverNoteText.length }}/500 characters
        </div>
      </div>

      <div class="flex items-center gap-3 justify-end pt-2 border-t border-slate-100">
        <button
          type="button"
          @click="selectedJobForApply = null"
          class="px-4 py-2 border border-slate-200 text-sm font-semibold rounded-lg text-slate-700 bg-white hover:bg-slate-50 transition-all"
        >
          Cancel
        </button>
        <button
          type="submit"
          :disabled="isSubmittingApply"
          class="px-5 py-2 border border-transparent text-sm font-semibold rounded-lg text-slate-950 bg-cyan-400 hover:bg-cyan-300 shadow-sm transition-all flex items-center space-x-2"
        >
          <span v-if="isSubmittingApply" class="animate-spin h-4 w-4 border-2 border-white border-t-transparent rounded-full"></span>
          <span>Submit Application</span>
        </button>
      </div>
    </form>

  </div>
</div>
<div v-if="appIdToCancel !== null" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm">
    <div class="bg-white rounded-2xl border border-slate-200 shadow-xl max-w-sm w-full p-6 animate-in fade-in zoom-in-95 duration-150 text-center">
      <div class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-red-50 text-red-600 mb-4">
        <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
        </svg>
      </div>
      
      <h3 class="text-lg font-bold text-slate-900 mb-2">Xác nhận hủy ứng tuyển</h3>
      <p class="text-sm text-slate-500 mb-6">
        Bạn có chắc chắn muốn rút hồ sơ khỏi vị trí này không? Hành động này không thể hoàn tác.
      </p>

      <div class="flex items-center gap-3 justify-center">
        <button
          type="button"
          @click="appIdToCancel = null"
          class="flex-1 px-4 py-2 border border-slate-200 text-sm font-semibold rounded-lg text-slate-700 bg-white hover:bg-slate-50 transition-all"
        >
          Không, quay lại
        </button>
        <button
          type="button"
          @click="confirmCancelApplication"
          class="flex-1 px-4 py-2 border border-transparent text-sm font-semibold rounded-lg text-white bg-red-600 hover:bg-red-500 shadow-sm transition-all"
        >
          Vâng, hủy đơn
        </button>
      </div>
    </div>
  </div>
  <div v-if="selectedOffer" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm">
  <div class="bg-white rounded-2xl border border-slate-200 shadow-2xl max-w-md w-full overflow-hidden text-left animate-in fade-in zoom-in-95 duration-150">
    
    <div class="bg-gradient-to-br from-slate-950 to-cyan-950 p-6 text-white text-center relative">
      <span class="text-4xl block mb-1">💼</span>
      <h3 class="text-lg font-bold">Lời Mời Nhận Việc (Job Offer)</h3>
      <p class="text-xs text-cyan-100 mt-1">Hồ sơ ứng tuyển của bạn đã được doanh nghiệp phê duyệt</p>
    </div>
<div class="p-6 space-y-5 bg-slate-900">
      <div>
        <label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest block mb-1">Vị trí & Công ty</label>
        <div class="text-sm font-black text-white uppercase tracking-wide">{{ selectedOffer.job?.title || 'Unknown Position' }}</div>
      </div>

      <div class="grid grid-cols-2 gap-4 border-t border-b border-slate-800/60 py-3.5 my-2">
        <div>
          <label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest block mb-1">💰 Mức lương Offer</label>
          <div class="text-sm font-black text-brand-400 tracking-wider">{{ selectedOffer.offer_salary || 'Thỏa thuận' }}</div>
        </div>
        <div>
          <label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest block mb-1">📅 Ngày bắt đầu</label>
          <div class="text-sm font-bold text-slate-300 tracking-wide">{{ selectedOffer.offer_start_date || 'Trao đổi sau' }}</div>
        </div>
      </div>

      <div>
        <label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest block mb-1.5">✉️ Thư chào mời từ phía HR</label>
        <div class="p-4 bg-slate-950 border border-slate-800 rounded-xl text-xs font-medium text-slate-300 whitespace-pre-line leading-relaxed">
          {{ selectedOffer.offer_message || 'Chào mừng bạn đến với công ty!' }}
        </div>
      </div>

      <div class="pt-4 border-t border-slate-800/60">
        <label class="text-[10px] font-bold text-slate-500 uppercase tracking-widest block mb-3">💬 Thảo luận trực tuyến với HR</label>
        <ChatBox 
          v-if="currentUserId !== null"
          :applicationId="selectedOffer.id"
          :targetId="selectedOffer.job?.business_id"
          :currentUserId="currentUserId"
        />
      </div>
    </div>

    <div class="p-4 bg-slate-950/60 border-t border-slate-800/80 flex items-center justify-between space-x-3">
      <button 
        @click="handleOfferResponse('decline')"
        :disabled="isResponding"
        class="w-1/2 px-4 py-2.5 text-xs font-bold uppercase tracking-wider text-rose-400 bg-rose-500/10 hover:bg-rose-500/20 rounded-xl border border-rose-500/10 transition-all text-center"
      >
        Từ chối Offer
      </button>
      <button 
        @click="handleOfferResponse('accept')"
        :disabled="isResponding"
        class="w-1/2 px-4 py-2.5 text-xs font-bold uppercase tracking-wider text-white bg-emerald-600 hover:bg-emerald-500 rounded-xl transition-all text-center shadow-lg shadow-emerald-600/10"
      >
        {{ isResponding ? 'Đang gửi...' : 'Đồng ý nhận việc' }}
      </button>
    </div>
  </div>
  </div>
  <!-- ======================================================== -->
  <!-- 🌟 MODAL CỬA SỔ CHAT REAL-TIME (DÀNH CHO STUDENT) -->
  <!-- ======================================================== -->
  <div v-if="isChatModalOpen && selectedChatApp" class="fixed inset-0 z-50 flex items-center justify-center p-4">
    <!-- Lớp phủ mờ nền sau -->
    <div class="absolute inset-0 bg-slate-950/60 backdrop-blur-sm" @click="isChatModalOpen = false"></div>
    
    <!-- Thân Modal -->
    <div class="relative w-full max-w-2xl bg-slate-950 border border-slate-800 rounded-2xl shadow-2xl overflow-hidden animate-in fade-in zoom-in-95 duration-150">
      
      <!-- Header Modal -->
      <div class="p-4 bg-slate-900 border-b border-slate-800/80 flex items-center justify-between">
        <div class="flex items-center space-x-3">
          <div class="w-2 h-2 bg-emerald-500 rounded-full animate-pulse"></div>
          <div>
            <h3 class="text-xs font-black text-white uppercase tracking-wider">
              Trò chuyện cùng {{ selectedChatApp.job?.business?.company_name || 'Nhà Tuyển Dụng' }}
            </h3>
            <p class="text-[10px] text-slate-500 font-bold uppercase mt-0.5 tracking-wide">
              Vị trí: {{ selectedChatApp.job?.title }} — Mã đơn: #{{ selectedChatApp.id }}
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

      <!-- Ruột Modal: Gọi component ChatBox truyền các tham số tương ứng -->
<div class="p-4 bg-slate-900/40">
        <ChatBox 
          v-if="currentUserId !== null"
          :applicationId="selectedChatApp.id"
          :targetId="selectedChatApp.job?.business.user_id || selectedChatApp.business_id"
          :currentUserId="currentUserId"
        />
      </div>

    </div>
  </div>
  <Transition
    enter-active-class="transform ease-out duration-300 transition"
    enter-from-class="translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-2"
    enter-to-class="translate-y-0 opacity-100 sm:translate-x-0"
    leave-active-class="transition ease-in duration-200"
    leave-from-class="opacity-100"
    leave-to-class="opacity-0"
  >
    <div 
      v-if="toast.show" 
      class="fixed bottom-5 right-5 z-50 max-w-sm w-full bg-slate-900/95 shadow-2xl shadow-slate-950/40 rounded-xl pointer-events-auto ring-1 ring-white/10 overflow-hidden border backdrop-blur"
      :class="toast.type === 'success' ? 'border-emerald-400/40' : 'border-rose-400/40'"
    >
      <div class="p-4">
        <div class="flex items-start">
          <div class="flex-shrink-0">
            <svg v-if="toast.type === 'success'" class="h-6 w-6 text-emerald-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <svg v-else class="h-6 w-6 text-rose-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div class="ml-3 w-0 flex-1 pt-0.5">
            <p class="text-sm font-bold text-white">
              {{ toast.type === 'success' ? 'Thông báo hệ thống' : 'Đã xảy ra lỗi' }}
            </p>
            <p class="mt-1 text-xs font-medium text-slate-300 leading-relaxed">
              {{ toast.message }}
            </p>
          </div>
          <div class="ml-4 flex-shrink-0 flex">
            <button @click="toast.show = false" class="bg-white/10 rounded-md inline-flex text-slate-400 hover:text-white focus:outline-none">
              <span class="sr-only">Close</span>
              <svg class="h-4 w-4" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { toRefs } from 'vue'
import ChatBox from '~/components/ChatBox.vue'

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
  wallet,
  walletTransactions,
  isLoadingWallet,
  fetchWallet,
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
