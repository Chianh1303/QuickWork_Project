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
  <div class="flex max-h-[90vh] w-full max-w-md flex-col overflow-hidden rounded-2xl border border-slate-800 bg-slate-950 text-left shadow-2xl animate-in fade-in zoom-in-95 duration-150">
    
    <div class="bg-gradient-to-br from-slate-950 to-cyan-950 p-6 text-white text-center relative">
      <span class="text-4xl block mb-1">💼</span>
      <h3 class="text-lg font-bold">Lời Mời Nhận Việc (Job Offer)</h3>
      <p class="text-xs text-cyan-100 mt-1">Hồ sơ ứng tuyển của bạn đã được doanh nghiệp phê duyệt</p>
    </div>
<div class="flex-1 overflow-y-auto p-6 space-y-5 bg-slate-900">
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
    </div>

    <div class="flex flex-shrink-0 items-center justify-between space-x-3 border-t border-slate-800/80 bg-slate-950 p-4">
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
  <div v-if="selectedManagedApplication" class="fixed inset-0 z-[49] flex justify-end bg-slate-950/80 backdrop-blur-md">
    <button
      type="button"
      class="absolute inset-0 cursor-default"
      aria-label="Close application details"
      @click="closeManagedApplicationModal"
    ></button>

    <aside class="relative flex h-full w-full max-w-2xl flex-col overflow-hidden border-l border-white/10 bg-slate-950 shadow-2xl shadow-slate-950/70">
      <div class="border-b border-white/10 bg-slate-900/90 p-5">
        <div class="flex items-start justify-between gap-4">
          <div class="min-w-0">
            <p class="text-xs font-bold uppercase tracking-wider text-cyan-300">Application Management</p>
            <h3 class="mt-1 truncate text-2xl font-extrabold text-white">
              {{ selectedManagedApplication.job?.title || 'Application details' }}
            </h3>
            <p class="mt-1 text-sm font-medium text-slate-400">
              {{ companyNameLookup(selectedManagedApplication.job) }} · Application #{{ selectedManagedApplication.id }}
            </p>
          </div>
          <button
            type="button"
            @click="closeManagedApplicationModal"
            class="rounded-lg bg-white/5 px-3 py-1.5 text-sm font-bold text-slate-300 hover:bg-white/10 hover:text-white"
          >
            Đóng
          </button>
        </div>
      </div>

      <div class="flex-1 space-y-5 overflow-y-auto p-5">
        <section class="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
          <div class="flex flex-wrap items-start justify-between gap-3">
            <div>
              <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Application Summary</p>
              <h4 class="mt-2 text-lg font-extrabold text-white">{{ selectedManagedApplication.job?.title || 'Unknown Position' }}</h4>
              <p class="mt-1 text-sm font-bold text-cyan-200">{{ companyNameLookup(selectedManagedApplication.job) }}</p>
            </div>
            <span :class="[
              statusBadgeClass(selectedManagedApplication.status),
              'inline-flex items-center rounded-full border px-2.5 py-1 text-xs font-bold capitalize'
            ]">
              {{ selectedManagedApplication.status ? selectedManagedApplication.status.replace('_', ' ') : 'Pending' }}
            </span>
          </div>

          <div class="mt-4 grid gap-3 sm:grid-cols-3">
            <div class="rounded-xl border border-white/10 bg-slate-900 p-3">
              <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Location</p>
              <p class="mt-1 text-sm font-extrabold text-slate-200">{{ selectedManagedApplication.job?.location || 'N/A' }}</p>
            </div>
            <div class="rounded-xl border border-emerald-400/20 bg-emerald-400/10 p-3">
              <p class="text-xs font-bold uppercase tracking-wider text-emerald-200">Salary</p>
              <p class="mt-1 text-sm font-extrabold text-white">${{ Number(selectedManagedApplication.job?.salary || 0).toLocaleString() }}</p>
            </div>
            <div class="rounded-xl border border-white/10 bg-slate-900 p-3">
              <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Applied</p>
              <p class="mt-1 text-sm font-extrabold text-slate-200">{{ formatDate(selectedManagedApplication.applied_at || selectedManagedApplication.id) }}</p>
            </div>
          </div>
        </section>

        <section class="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
          <div class="flex items-center justify-between gap-3">
            <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Offer</p>
            <button
              v-if="(selectedManagedApplication.status || '').toLowerCase() === 'approved'"
              @click="openOfferModal(selectedManagedApplication); closeManagedApplicationModal()"
              class="rounded-lg bg-emerald-500 px-3 py-1.5 text-xs font-extrabold text-white hover:bg-emerald-400"
            >
              View Offer
            </button>
          </div>
          <div
            v-if="selectedManagedApplication.offer_salary || selectedManagedApplication.offer_start_date || selectedManagedApplication.offer_message"
            class="mt-3 grid gap-3 sm:grid-cols-2"
          >
            <div class="rounded-xl border border-emerald-400/20 bg-emerald-400/10 p-3">
              <p class="text-xs font-bold uppercase tracking-wider text-emerald-200">Offer Salary</p>
              <p class="mt-1 text-sm font-extrabold text-white">{{ selectedManagedApplication.offer_salary || 'Thỏa thuận' }}</p>
            </div>
            <div class="rounded-xl border border-cyan-400/20 bg-cyan-400/10 p-3">
              <p class="text-xs font-bold uppercase tracking-wider text-cyan-200">Start Date</p>
              <p class="mt-1 text-sm font-extrabold text-white">{{ selectedManagedApplication.offer_start_date || 'Trao đổi sau' }}</p>
            </div>
            <div class="rounded-xl border border-white/10 bg-slate-900 p-3 sm:col-span-2">
              <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Message</p>
              <p class="mt-1 whitespace-pre-line text-sm font-medium text-slate-300">{{ selectedManagedApplication.offer_message || 'Không có lời nhắn offer.' }}</p>
            </div>
          </div>
          <p v-else class="mt-3 text-sm font-medium text-slate-500">Offer information appears when the employer approves your application.</p>
        </section>

        <section
          v-if="(selectedManagedApplication.status || '').toLowerCase() === 'offer_accepted'"
          class="rounded-2xl border border-amber-400/20 bg-amber-400/10 p-4"
        >
          <p class="text-xs font-bold uppercase tracking-wider text-amber-200">Work / Attendance</p>
          <div class="mt-3 flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
            <div>
              <p class="text-sm font-extrabold text-white">
                {{ isWorking(selectedManagedApplication.job_id) ? `Đang làm việc (${getTimer(selectedManagedApplication.job_id)})` : 'Chưa vào ca' }}
              </p>
              <p class="mt-1 text-xs font-medium text-slate-400">Theo dõi ca làm việc hiện tại của bạn.</p>
            </div>
            <button
              v-if="!isWorking(selectedManagedApplication.job_id)"
              @click="handleCheckIn(selectedManagedApplication.job_id)"
              class="rounded-lg bg-emerald-500 px-4 py-2 text-sm font-extrabold text-white hover:bg-emerald-400"
            >
              Check-in
            </button>
            <button
              v-else
              @click="handleCheckOut(selectedManagedApplication.job_id)"
              class="rounded-lg bg-rose-500 px-4 py-2 text-sm font-extrabold text-white hover:bg-rose-400"
            >
              Check-out
            </button>
          </div>
        </section>

        <section class="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
          <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Completion</p>
          <div class="mt-3">
            <button
              v-if="(selectedManagedApplication.status || '').toLowerCase() === 'offer_accepted'"
              @click="handleStudentComplete(selectedManagedApplication.id); closeManagedApplicationModal()"
              class="w-full rounded-lg bg-cyan-400 px-4 py-2.5 text-sm font-extrabold text-slate-950 hover:bg-cyan-300"
            >
              Xác nhận hoàn thành
            </button>
            <div
              v-else-if="(selectedManagedApplication.status || '').toLowerCase() === 'student_completed'"
              class="rounded-lg border border-amber-400/20 bg-amber-400/10 px-4 py-2.5 text-sm font-extrabold text-amber-200"
            >
              Đang chờ doanh nghiệp xác nhận
            </div>
            <div
              v-else-if="(selectedManagedApplication.status || '').toLowerCase() === 'paid'"
              class="rounded-lg border border-cyan-400/20 bg-cyan-400/10 px-4 py-2.5 text-sm font-extrabold text-cyan-200"
            >
              Đã giải ngân
            </div>
            <p v-else class="text-sm font-medium text-slate-500">Completion actions appear after you accept an offer.</p>
          </div>
        </section>

        <section class="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
          <div class="flex flex-wrap items-center justify-between gap-3">
            <div>
              <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Reviews</p>
              <p class="mt-1 text-sm font-extrabold text-white">
                {{ reviewSummary.average_rating.toFixed(1) }} / 5 · {{ reviewSummary.total_reviews }} reviews
              </p>
            </div>
            <div class="flex flex-wrap gap-2">
              <button
                v-if="(selectedManagedApplication.status || '').toLowerCase() === 'paid' && !reviewsList.some(review => Number(review.reviewer_id || review.ReviewerID) === Number(currentUserId))"
                @click="openReviewModal(selectedManagedApplication); closeManagedApplicationModal()"
                class="rounded-lg bg-amber-400 px-3 py-1.5 text-xs font-extrabold text-slate-950 hover:bg-amber-300"
              >
                Đánh giá doanh nghiệp
              </button>
              <button
                @click="openReviewsModal(selectedManagedApplication)"
                class="rounded-lg border border-white/10 bg-white/5 px-3 py-1.5 text-xs font-bold text-slate-200 hover:bg-white/10 hover:text-white"
              >
                Xem đánh giá
              </button>
            </div>
          </div>
          <div class="mt-3 space-y-2">
            <div v-if="isLoadingReviews" class="text-sm font-medium text-slate-500">Đang tải đánh giá...</div>
            <div
              v-for="review in reviewsList.slice(0, 3)"
              :key="review.id || review.ID"
              class="rounded-xl border border-white/10 bg-slate-900 p-3"
            >
              <p class="text-sm font-extrabold text-amber-300">{{ review.rating || review.Rating }} / 5</p>
              <p class="mt-1 text-sm font-medium text-slate-300">{{ review.comment || review.Comment || 'Không có nhận xét.' }}</p>
            </div>
            <p v-if="!isLoadingReviews && reviewsList.length === 0" class="text-sm font-medium text-slate-500">Chưa có đánh giá cho đơn này.</p>
          </div>
        </section>
      </div>

      <div class="border-t border-white/10 bg-slate-900/90 p-5">
        <div class="grid gap-2 sm:grid-cols-2">
          <button
            @click="openChatModal(selectedManagedApplication); closeManagedApplicationModal()"
            class="rounded-lg border border-white/10 bg-white/5 px-4 py-2.5 text-sm font-bold text-slate-200 hover:bg-white/10 hover:text-white"
          >
            Chat with employer
          </button>
          <button
            v-if="(selectedManagedApplication.status || '').toLowerCase() === 'pending'"
            :disabled="isCancellingApp === selectedManagedApplication.id"
            @click="triggerCancelConfirm(selectedManagedApplication.id); closeManagedApplicationModal()"
            class="rounded-lg border border-rose-400/20 bg-rose-500/10 px-4 py-2.5 text-sm font-bold text-rose-300 hover:bg-rose-500/20 disabled:opacity-50"
          >
            Hủy ứng tuyển
          </button>
        </div>
      </div>
    </aside>
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
  <div
  v-if="selectedReviewApp"
  class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/70 backdrop-blur-sm"
>
  <div class="w-full max-w-md overflow-hidden rounded-2xl border border-white/10 bg-slate-950 shadow-2xl">
    <div class="p-6 border-b border-white/10">
      <h3 class="text-xl font-black text-white">Đánh giá doanh nghiệp</h3>
      <p class="mt-1 text-sm text-slate-400">
        {{ selectedReviewApp.job?.title || 'Công việc đã hoàn thành' }}
      </p>
    </div>

    <form @submit.prevent="submitReview" class="p-6 space-y-5">
      <div>
        <label class="block text-xs font-bold uppercase tracking-widest text-slate-400 mb-3">
          Số sao
        </label>

        <div class="flex gap-2">
          <button
            v-for="star in 5"
            :key="star"
            type="button"
            @click="reviewRating = star"
            class="text-3xl transition"
            :class="star <= reviewRating ? 'text-amber-400' : 'text-slate-700'"
          >
            ★
          </button>
        </div>
      </div>

      <div>
        <label class="block text-xs font-bold uppercase tracking-widest text-slate-400 mb-2">
          Nhận xét
        </label>

        <textarea
          v-model="reviewComment"
          rows="4"
          class="w-full rounded-xl border border-white/10 bg-slate-900 px-4 py-3 text-sm text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-amber-300 resize-none"
          placeholder="Chia sẻ trải nghiệm làm việc với doanh nghiệp..."
        ></textarea>
      </div>

      <div class="flex justify-end gap-3 pt-4 border-t border-white/10">
        <button
          type="button"
          @click="selectedReviewApp = null"
          class="px-4 py-2 rounded-lg bg-slate-800 text-slate-300 text-sm font-bold hover:bg-slate-700"
        >
          Hủy
        </button>

        <button
          type="submit"
          :disabled="isSubmittingReview"
          class="px-5 py-2 rounded-lg bg-amber-400 text-slate-950 text-sm font-extrabold hover:bg-amber-300 disabled:opacity-50"
        >
          {{ isSubmittingReview ? 'Đang gửi...' : 'Gửi đánh giá' }}
        </button>
      </div>
    </form>
  </div>
</div>
<div v-if="isReviewsModalOpen" class="fixed inset-0 z-[60] flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
  <div class="w-full max-w-xl overflow-hidden rounded-2xl border border-white/10 bg-slate-950 shadow-2xl shadow-slate-950/70">
    <div class="flex items-start justify-between gap-4 border-b border-white/10 bg-slate-900/90 p-5">
      <div>
        <p class="text-xs font-bold uppercase tracking-wider text-cyan-300">Reviews</p>
        <h3 class="mt-1 text-xl font-extrabold text-white">{{ selectedReviewsApp?.job?.title || 'Application reviews' }}</h3>
        <p class="mt-1 text-sm font-medium text-slate-400">
          {{ reviewSummary.average_rating.toFixed(1) }} / 5 · {{ reviewSummary.total_reviews }} reviews
        </p>
      </div>
      <button
        type="button"
        @click="closeReviewsModal"
        class="rounded-lg bg-white/5 px-3 py-1.5 text-sm font-bold text-slate-300 hover:bg-white/10 hover:text-white"
      >
        Đóng
      </button>
    </div>

    <div class="max-h-[70vh] space-y-3 overflow-y-auto p-5">
      <div v-if="isLoadingReviews" class="rounded-xl border border-white/10 bg-white/[0.04] p-4 text-sm font-medium text-slate-400">
        Đang tải đánh giá...
      </div>
      <div
        v-for="review in reviewsList"
        :key="review.id || review.ID"
        class="rounded-xl border border-white/10 bg-white/[0.04] p-4"
      >
        <div class="flex items-center justify-between gap-3">
          <p class="text-sm font-extrabold text-amber-300">{{ review.rating || review.Rating }} / 5</p>
          <p class="text-xs font-bold uppercase tracking-wider text-slate-500">{{ review.review_type || review.ReviewType || 'review' }}</p>
        </div>
        <p class="mt-2 text-sm font-medium text-slate-300">{{ review.comment || review.Comment || 'Không có nhận xét.' }}</p>
      </div>
      <p v-if="!isLoadingReviews && reviewsList.length === 0" class="rounded-xl border border-white/10 bg-white/[0.04] p-4 text-sm font-medium text-slate-500">
        Chưa có đánh giá cho đơn này.
      </p>
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
  selectedManagedApplication,
  closeManagedApplicationModal,
  selectedReviewApp,
  reviewRating,
  reviewComment,
  isSubmittingReview,
  submitReview,
  reviewSummary,
  reviewsList,
  isLoadingReviews,
  selectedReviewsApp,
  isReviewsModalOpen,
  openReviewsModal,
  closeReviewsModal,
  openManagedApplicationModal,
  handleStudentComplete,
  openReviewModal,
  toast
} = toRefs(props.state)
</script>
