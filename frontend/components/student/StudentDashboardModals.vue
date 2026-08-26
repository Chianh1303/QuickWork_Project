<template>
  <div v-if="selectedJobForApply" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
    <div class="bg-slate-900/95 rounded-3xl border border-cyan-500/20 shadow-2xl shadow-slate-950 max-w-lg w-full overflow-hidden animate-in fade-in zoom-in-95 duration-200 flex flex-col max-h-[90vh]">
      
      <!-- Modal Header -->
      <div class="p-6 pb-4 border-b border-cyan-500/15 flex items-start justify-between gap-3 bg-slate-950/60">
        <div>
          <div class="flex flex-wrap items-center gap-2 mb-1.5">
            <span class="inline-flex items-center px-3 py-1 rounded-lg text-xs font-bold bg-cyan-500/10 text-cyan-200 border border-cyan-500/20">
              {{ selectedJobForApply.company || companyNameLookup(selectedJobForApply) }}
            </span>
            <span v-if="selectedJobForApply.match_score !== undefined" class="inline-flex items-center px-3 py-1 rounded-full text-xs font-black bg-emerald-400/10 text-emerald-300 ring-1 ring-emerald-400/30">
              ✨ {{ selectedJobForApply.match_score }}% Phù hợp AI
            </span>
          </div>
          <h3 class="text-lg font-extrabold text-white leading-snug">
            {{ selectedJobForApply.title }}
          </h3>
        </div>
        <button @click="selectedJobForApply = null" class="text-slate-400 hover:text-white p-1.5 rounded-xl hover:bg-white/10 flex-shrink-0 transition-colors">
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <!-- Modal Body (Scrollable Details & Description) -->
      <div class="p-6 overflow-y-auto space-y-4 flex-1 custom-scrollbar">
        <!-- Metadata (Salary & Location) -->
        <div v-if="selectedJobForApply.salary || selectedJobForApply.location" class="flex flex-wrap items-center gap-3 text-xs font-semibold text-slate-300 bg-slate-950/70 border border-cyan-500/15 p-3.5 rounded-2xl">
          <div v-if="selectedJobForApply.salary" class="flex items-center gap-1.5 text-emerald-300 font-extrabold">
            <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <span>{{ Number(selectedJobForApply.salary || 0).toLocaleString('vi-VN') }} VNĐ</span>
          </div>
          <div v-if="selectedJobForApply.location" class="flex items-center gap-1.5 text-slate-300">
            <svg class="h-4 w-4 text-cyan-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
            </svg>
            <span>{{ selectedJobForApply.location }}</span>
          </div>
        </div>

        <!-- AI Skills Breakdown -->
        <div v-if="(selectedJobForApply.matching_skills && selectedJobForApply.matching_skills.length) || (selectedJobForApply.missing_skills && selectedJobForApply.missing_skills.length)" class="rounded-2xl border border-cyan-500/20 bg-cyan-500/5 p-4 space-y-2.5">
          <div class="text-xs font-black uppercase tracking-wider text-cyan-300 flex items-center gap-1.5">
            <span>✨ Phân Tích Độ Tương Thích AI</span>
          </div>
          <div v-if="selectedJobForApply.matching_skills && selectedJobForApply.matching_skills.length" class="space-y-1">
            <span class="text-[11px] font-extrabold text-emerald-400">Kỹ năng đáp ứng:</span>
            <div class="flex flex-wrap gap-1.5">
              <span v-for="skill in selectedJobForApply.matching_skills" :key="skill" class="px-2.5 py-0.5 rounded-md text-[11px] font-bold bg-emerald-500/10 text-emerald-300 border border-emerald-500/20">
                {{ skill }}
              </span>
            </div>
          </div>
          <div v-if="selectedJobForApply.missing_skills && selectedJobForApply.missing_skills.length" class="space-y-1 pt-1">
            <span class="text-[11px] font-extrabold text-slate-400">Kỹ năng cần bổ sung:</span>
            <div class="flex flex-wrap gap-1.5">
              <span v-for="skill in selectedJobForApply.missing_skills" :key="skill" class="px-2.5 py-0.5 rounded-md text-[11px] font-bold bg-slate-800 text-slate-400 border border-slate-700">
                {{ skill }}
              </span>
            </div>
          </div>
        </div>

        <!-- Description Box -->
        <div class="rounded-2xl border border-cyan-500/15 bg-slate-950/70 p-4 space-y-2">
          <h4 class="text-xs font-black uppercase tracking-wider text-cyan-300">Mô tả chi tiết công việc</h4>
          <p class="text-xs font-medium text-slate-300 leading-relaxed whitespace-pre-line max-h-44 overflow-y-auto pr-1">
            {{ selectedJobForApply.description || 'Chưa có thông tin mô tả chi tiết công việc.' }}
          </p>
        </div>

        <!-- Cover Note Input -->
        <form id="apply-opportunity-form" @submit.prevent="submitApplication" class="space-y-2 pt-1">
          <label class="block text-xs font-black uppercase tracking-wider text-cyan-300">
            Thư giới thiệu gửi Nhà tuyển dụng (Cover Note)
          </label>
          <textarea
            v-model="coverNoteText"
            rows="3"
            maxlength="500"
            class="block w-full px-3.5 py-2.5 border border-cyan-500/20 rounded-2xl text-xs bg-slate-950/90 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 resize-none transition-all font-medium"
            placeholder="Giới thiệu ngắn gọn về thế mạnh của bạn hoặc lý do bạn mong muốn ứng tuyển vị trí này..."
          ></textarea>
          <div class="text-right text-[11px] font-semibold text-slate-400">
            {{ coverNoteText.length }}/500 ký tự
          </div>
        </form>
      </div>

      <!-- Modal Footer -->
      <div class="flex items-center gap-3 justify-end p-4 border-t border-cyan-500/15 bg-slate-950/60">
        <button
          type="button"
          @click="selectedJobForApply = null"
          class="px-4 py-2 border border-white/10 text-xs font-bold rounded-xl text-slate-300 bg-white/5 hover:bg-white/10 hover:text-white transition-all"
        >
          Hủy bỏ
        </button>
        <button
          type="submit"
          form="apply-opportunity-form"
          :disabled="isSubmittingApply"
          class="px-5 py-2 text-xs font-extrabold rounded-xl text-white bg-gradient-to-r from-cyan-500 via-blue-600 to-emerald-500 hover:from-cyan-400 hover:to-emerald-400 shadow-md shadow-cyan-500/25 transition-all flex items-center space-x-2 disabled:opacity-50"
        >
          <span v-if="isSubmittingApply" class="animate-spin h-3.5 w-3.5 border-2 border-white border-t-transparent rounded-full"></span>
          <span>Gửi đơn ứng tuyển chính thức</span>
        </button>
      </div>

    </div>
  </div>
  <!-- Modal Xác nhận Hủy Ứng tuyển (Dark Theme đồng bộ) -->
  <div v-if="appIdToCancel !== null" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
    <div class="bg-slate-900/95 rounded-3xl border border-rose-500/25 shadow-2xl shadow-slate-950 max-w-sm w-full p-6 animate-in fade-in zoom-in-95 duration-150 text-center">
      <div class="mx-auto flex items-center justify-center h-14 w-14 rounded-2xl bg-rose-500/10 border border-rose-500/20 text-rose-400 mb-4 shadow-lg shadow-rose-500/10">
        <svg class="h-7 w-7" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
        </svg>
      </div>
      
      <h3 class="text-lg font-extrabold text-white mb-2">Xác nhận hủy ứng tuyển</h3>
      <p class="text-xs font-semibold text-slate-300 mb-6 leading-relaxed">
        Bạn có chắc chắn muốn rút hồ sơ khỏi vị trí này không? Hành động này không thể hoàn tác.
      </p>

      <div class="flex items-center gap-3 justify-center">
        <button
          type="button"
          @click="appIdToCancel = null"
          class="flex-1 px-4 py-2.5 border border-white/10 text-xs font-bold rounded-xl text-slate-300 bg-white/5 hover:bg-white/10 hover:text-white transition-all"
        >
          Không, quay lại
        </button>
        <button
          type="button"
          @click="confirmCancelApplication"
          :disabled="isCancellingApp"
          class="flex-1 px-4 py-2.5 border border-transparent text-xs font-extrabold rounded-xl text-white bg-gradient-to-r from-rose-600 to-red-600 hover:from-rose-500 hover:to-red-500 shadow-md shadow-rose-600/25 transition-all disabled:opacity-50"
        >
          {{ isCancellingApp ? 'Đang hủy...' : 'Vâng, hủy đơn' }}
        </button>
      </div>
    </div>
  </div>
  <div v-if="selectedOffer" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm">
  <div class="flex max-h-[90vh] w-full max-w-md flex-col overflow-hidden rounded-2xl border border-slate-800 bg-slate-950 text-left shadow-2xl animate-in fade-in zoom-in-95 duration-150">
    
    <div class="bg-gradient-to-br from-slate-950 to-cyan-950 p-6 text-white text-center relative">
      <span class="text-4xl block mb-1">💼</span>
      <h3 class="text-lg font-semibold">Lời Mời Nhận Việc (Job Offer)</h3>
      <p class="text-xs text-cyan-100 mt-1">Hồ sơ ứng tuyển của bạn đã được doanh nghiệp phê duyệt</p>
    </div>
<div class="flex-1 overflow-y-auto p-6 space-y-5 bg-slate-900">
      <div>
        <label class="text-[10px] font-semibold text-slate-500 uppercase tracking-widest block mb-1">Vị trí & Công ty</label>
        <div class="text-sm font-semibold text-white uppercase tracking-wide">{{ selectedOffer.job?.title || 'Unknown Position' }}</div>
      </div>

      <div class="grid grid-cols-2 gap-4 border-t border-b border-slate-800/60 py-3.5 my-2">
        <div>
          <label class="text-[10px] font-semibold text-slate-500 uppercase tracking-widest block mb-1">💰 Mức lương Offer</label>
          <div class="text-sm font-semibold text-brand-400 tracking-wider">{{ selectedOffer.offer_salary || 'Thỏa thuận' }}</div>
        </div>
        <div>
          <label class="text-[10px] font-semibold text-slate-500 uppercase tracking-widest block mb-1">📅 Ngày bắt đầu</label>
          <div class="text-sm font-semibold text-slate-300 tracking-wide">{{ selectedOffer.offer_start_date || 'Trao đổi sau' }}</div>
        </div>
      </div>

      <div>
        <label class="text-[10px] font-semibold text-slate-500 uppercase tracking-widest block mb-1.5">✉️ Thư chào mời từ phía HR</label>
        <div class="p-4 bg-slate-950 border border-slate-800 rounded-xl text-xs font-medium text-slate-300 whitespace-pre-line leading-relaxed">
          {{ selectedOffer.offer_message || 'Chào mừng bạn đến với công ty!' }}
        </div>
      </div>
    </div>

    <div class="flex flex-shrink-0 items-center justify-between space-x-3 border-t border-slate-800/80 bg-slate-950 p-4">
      <button 
        @click="handleOfferResponse('decline')"
        :disabled="isResponding"
        class="w-1/2 px-4 py-2.5 text-xs font-semibold uppercase tracking-wider text-rose-400 bg-rose-500/10 hover:bg-rose-500/20 rounded-xl border border-rose-500/10 transition-all text-center"
      >
        Từ chối Offer
      </button>
      <button 
        @click="handleOfferResponse('accept')"
        :disabled="isResponding"
        class="w-1/2 px-4 py-2.5 text-xs font-semibold uppercase tracking-wider text-white bg-emerald-600 hover:bg-emerald-500 rounded-xl transition-all text-center shadow-lg shadow-emerald-600/10"
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
            <h3 class="text-xs font-semibold text-white uppercase tracking-wider">
              Trò chuyện cùng {{ selectedChatApp.job?.business?.company_name || 'Nhà Tuyển Dụng' }}
            </h3>
            <p class="text-[10px] text-slate-500 font-semibold uppercase mt-0.5 tracking-wide">
              Vị trí: {{ selectedChatApp.job?.title }} — Mã đơn: #{{ selectedChatApp.id }}
            </p>
          </div>
        </div>
        <button 
          @click="isChatModalOpen = false" 
          class="text-slate-400 hover:text-white text-xs font-semibold px-2.5 py-1.5 bg-slate-800 hover:bg-slate-700 rounded-xl transition-all"
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
      <div class="border-b border-cyan-500/15 bg-slate-900/90 p-5">
        <div class="flex items-start justify-between gap-4">
          <div class="min-w-0">
            <span class="inline-flex items-center rounded-full bg-cyan-500/10 px-2.5 py-0.5 text-xs font-black uppercase tracking-wider text-cyan-300 ring-1 ring-cyan-500/20">
              Tiến Trình Ứng Tuyển
            </span>
            <h3 class="mt-1.5 truncate text-xl font-extrabold text-white">
              {{ selectedManagedApplication.job?.title || 'Chi tiết đơn ứng tuyển' }}
            </h3>
            <p class="mt-1 text-xs font-semibold text-slate-400">
              {{ companyNameLookup(selectedManagedApplication.job) }} · Mã đơn #{{ selectedManagedApplication.id }}
            </p>
          </div>
          <button
            type="button"
            @click="closeManagedApplicationModal"
            class="rounded-xl bg-white/5 px-3 py-1.5 text-xs font-bold text-slate-300 hover:bg-white/10 hover:text-white transition-colors"
          >
            ✕ Đóng
          </button>
        </div>
      </div>

      <div class="flex-1 space-y-5 overflow-y-auto p-5">
        <!-- 5-Step Visual Stepper -->
        <section class="rounded-2xl border border-cyan-500/20 bg-slate-900/80 p-4 shadow-md">
          <p class="text-[11px] font-black uppercase tracking-wider text-cyan-300 mb-3">Trạng Thái Tiến Độ Ca Làm</p>
          <div class="relative flex items-center justify-between gap-1">
            <div class="absolute left-0 top-3.5 h-0.5 w-full bg-slate-800 -z-0"></div>
            
            <!-- Step 1: Applied -->
            <div class="relative z-10 flex flex-col items-center text-center">
              <div class="h-7 w-7 rounded-full flex items-center justify-center text-xs font-bold ring-4 ring-slate-950 bg-emerald-500 text-slate-950">
                ✓
              </div>
              <span class="mt-1.5 text-[10px] font-extrabold text-emerald-300">Đã nộp</span>
            </div>

            <!-- Step 2: Approved / Offer -->
            <div class="relative z-10 flex flex-col items-center text-center">
              <div
                :class="[
                  ['approved', 'offer_accepted', 'student_completed', 'paid'].includes((selectedManagedApplication.status || '').toLowerCase())
                    ? 'bg-emerald-500 text-slate-950'
                    : (selectedManagedApplication.status || '').toLowerCase() === 'rejected'
                    ? 'bg-rose-500 text-white'
                    : 'bg-slate-800 text-slate-400',
                  'h-7 w-7 rounded-full flex items-center justify-center text-xs font-bold ring-4 ring-slate-950'
                ]"
              >
                {{ ['approved', 'offer_accepted', 'student_completed', 'paid'].includes((selectedManagedApplication.status || '').toLowerCase()) ? '✓' : '2' }}
              </div>
              <span class="mt-1.5 text-[10px] font-bold" :class="['approved', 'offer_accepted', 'student_completed', 'paid'].includes((selectedManagedApplication.status || '').toLowerCase()) ? 'text-emerald-300' : 'text-slate-500'">
                DN duyệt
              </span>
            </div>

            <!-- Step 3: Offer Accepted -->
            <div class="relative z-10 flex flex-col items-center text-center">
              <div
                :class="[
                  ['offer_accepted', 'student_completed', 'paid'].includes((selectedManagedApplication.status || '').toLowerCase())
                    ? 'bg-emerald-500 text-slate-950'
                    : 'bg-slate-800 text-slate-400',
                  'h-7 w-7 rounded-full flex items-center justify-center text-xs font-bold ring-4 ring-slate-950'
                ]"
              >
                {{ ['offer_accepted', 'student_completed', 'paid'].includes((selectedManagedApplication.status || '').toLowerCase()) ? '✓' : '3' }}
              </div>
              <span class="mt-1.5 text-[10px] font-bold" :class="['offer_accepted', 'student_completed', 'paid'].includes((selectedManagedApplication.status || '').toLowerCase()) ? 'text-emerald-300' : 'text-slate-500'">
                Nhận việc
              </span>
            </div>

            <!-- Step 4: Completed Shift -->
            <div class="relative z-10 flex flex-col items-center text-center">
              <div
                :class="[
                  ['student_completed', 'paid'].includes((selectedManagedApplication.status || '').toLowerCase())
                    ? 'bg-emerald-500 text-slate-950'
                    : 'bg-slate-800 text-slate-400',
                  'h-7 w-7 rounded-full flex items-center justify-center text-xs font-bold ring-4 ring-slate-950'
                ]"
              >
                {{ ['student_completed', 'paid'].includes((selectedManagedApplication.status || '').toLowerCase()) ? '✓' : '4' }}
              </div>
              <span class="mt-1.5 text-[10px] font-bold" :class="['student_completed', 'paid'].includes((selectedManagedApplication.status || '').toLowerCase()) ? 'text-emerald-300' : 'text-slate-500'">
                Hoàn tất ca
              </span>
            </div>

            <!-- Step 5: Paid / Disbursed -->
            <div class="relative z-10 flex flex-col items-center text-center">
              <div
                :class="[
                  (selectedManagedApplication.status || '').toLowerCase() === 'paid'
                    ? 'bg-emerald-400 text-slate-950 shadow-lg shadow-emerald-500/30'
                    : 'bg-slate-800 text-slate-400',
                  'h-7 w-7 rounded-full flex items-center justify-center text-xs font-bold ring-4 ring-slate-950'
                ]"
              >
                {{ (selectedManagedApplication.status || '').toLowerCase() === 'paid' ? '💰' : '5' }}
              </div>
              <span class="mt-1.5 text-[10px] font-bold" :class="(selectedManagedApplication.status || '').toLowerCase() === 'paid' ? 'text-emerald-300' : 'text-slate-500'">
                Nhận lương
              </span>
            </div>
          </div>
        </section>

        <!-- Application Summary Card -->
        <section class="rounded-2xl border border-white/10 bg-slate-900/60 p-4">
          <div class="flex flex-wrap items-start justify-between gap-3">
            <div>
              <p class="text-[11px] font-black uppercase tracking-wider text-slate-400">Thông tin công việc</p>
              <h4 class="mt-1.5 text-base font-extrabold text-white">{{ selectedManagedApplication.job?.title || 'Vị trí chưa xác định' }}</h4>
              <p class="mt-0.5 text-xs font-bold text-cyan-300">{{ companyNameLookup(selectedManagedApplication.job) }}</p>
            </div>
            <span :class="[
              statusBadgeClass(selectedManagedApplication.status),
              'inline-flex items-center rounded-full border px-3 py-1 text-xs font-black capitalize'
            ]">
              {{ selectedManagedApplication.status ? selectedManagedApplication.status.replace('_', ' ') : 'Chờ duyệt' }}
            </span>
          </div>

          <div class="mt-4 grid gap-3 sm:grid-cols-3">
            <div class="rounded-xl border border-white/10 bg-slate-950/80 p-3">
              <p class="text-[10px] font-black uppercase tracking-wider text-slate-400">Địa điểm</p>
              <p class="mt-1 text-xs font-bold text-slate-200">{{ selectedManagedApplication.job?.location || 'N/A' }}</p>
            </div>
            <div class="rounded-xl border border-emerald-400/20 bg-emerald-400/10 p-3">
              <p class="text-[10px] font-black uppercase tracking-wider text-emerald-300">Mức lương</p>
              <p class="mt-1 text-xs font-black text-emerald-300">{{ Number(selectedManagedApplication.job?.salary || 0).toLocaleString('vi-VN') }} VNĐ</p>
            </div>
            <div class="rounded-xl border border-white/10 bg-slate-950/80 p-3">
              <p class="text-[10px] font-black uppercase tracking-wider text-slate-400">Ngày nộp</p>
              <p class="mt-1 text-xs font-bold text-slate-200">{{ formatDate(selectedManagedApplication.applied_at || selectedManagedApplication.id) }}</p>
            </div>
          </div>
        </section>

        <section class="rounded-2xl border border-white/10 bg-slate-900/60 p-4">
          <div class="flex items-center justify-between gap-3">
            <p class="text-xs font-black uppercase tracking-wider text-cyan-300">Chi Tiết Lời Mời Nhận Việc (Offer)</p>
            <button
              v-if="(selectedManagedApplication.status || '').toLowerCase() === 'approved'"
              @click="openOfferModal(selectedManagedApplication); closeManagedApplicationModal()"
              class="rounded-xl bg-emerald-500 px-3.5 py-1.5 text-xs font-black text-slate-950 hover:bg-emerald-400 transition-all shadow-md shadow-emerald-500/20"
            >
              Xem & Phản Hồi Offer
            </button>
          </div>
          <div
            v-if="selectedManagedApplication.offer_salary || selectedManagedApplication.offer_start_date || selectedManagedApplication.offer_message"
            class="mt-3 grid gap-3 sm:grid-cols-2"
          >
            <div class="rounded-xl border border-emerald-400/20 bg-emerald-400/10 p-3">
              <p class="text-xs font-semibold uppercase tracking-wider text-emerald-200">Offer Salary</p>
              <p class="mt-1 text-sm font-semibold text-white">{{ selectedManagedApplication.offer_salary || 'Thỏa thuận' }}</p>
            </div>
            <div class="rounded-xl border border-cyan-400/20 bg-cyan-400/10 p-3">
              <p class="text-xs font-semibold uppercase tracking-wider text-cyan-200">Start Date</p>
              <p class="mt-1 text-sm font-semibold text-white">{{ selectedManagedApplication.offer_start_date || 'Trao đổi sau' }}</p>
            </div>
            <div class="rounded-xl border border-white/10 bg-slate-900 p-3 sm:col-span-2">
              <p class="text-xs font-semibold uppercase tracking-wider text-slate-500">Message</p>
              <p class="mt-1 whitespace-pre-line text-sm font-medium text-slate-300">{{ selectedManagedApplication.offer_message || 'Không có lời nhắn offer.' }}</p>
            </div>
          </div>
          <p v-else class="mt-3 text-sm font-medium text-slate-500">Offer information appears when the employer approves your application.</p>
        </section>

        <section
          v-if="(selectedManagedApplication.status || '').toLowerCase() === 'offer_accepted'"
          class="rounded-2xl border border-amber-400/20 bg-amber-400/10 p-4"
        >
          <p class="text-xs font-semibold uppercase tracking-wider text-amber-200">Work / Attendance</p>
          <div class="mt-3 flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
            <div>
              <p class="text-sm font-semibold text-white">
                {{ isWorking(selectedManagedApplication.job_id) ? `Đang làm việc (${getTimer(selectedManagedApplication.job_id)})` : 'Chưa vào ca' }}
              </p>
              <p class="mt-1 text-xs font-medium text-slate-400">Theo dõi ca làm việc hiện tại của bạn.</p>
            </div>
            <button
              v-if="!isWorking(selectedManagedApplication.job_id)"
              @click="handleCheckIn(selectedManagedApplication.job_id)"
              class="rounded-lg bg-emerald-500 px-4 py-2 text-sm font-semibold text-white hover:bg-emerald-400"
            >
              Check-in
            </button>
            <button
              v-else
              @click="triggerCheckOutConfirm(selectedManagedApplication.job_id)"
              class="rounded-lg bg-rose-500 px-4 py-2 text-sm font-semibold text-white hover:bg-rose-400 cursor-pointer"
            >
              Check-out
            </button>
          </div>
        </section>

        <section class="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
          <p class="text-xs font-semibold uppercase tracking-wider text-slate-500">Completion</p>
          <div class="mt-3">
            <button
              v-if="(selectedManagedApplication.status || '').toLowerCase() === 'offer_accepted'"
              @click="handleStudentComplete(selectedManagedApplication.id); closeManagedApplicationModal()"
              class="w-full rounded-lg bg-cyan-400 px-4 py-2.5 text-sm font-semibold text-slate-950 hover:bg-cyan-300"
            >
              Xác nhận hoàn thành
            </button>
            <div
              v-else-if="(selectedManagedApplication.status || '').toLowerCase() === 'student_completed'"
              class="space-y-2"
            >
              <div class="rounded-lg border border-amber-400/20 bg-amber-400/10 px-4 py-2 text-xs font-bold text-amber-300 text-center">
                🟡 Đang chờ doanh nghiệp duyệt giải ngân
              </div>
              <button
                @click="openStudentCompletionModal(selectedManagedApplication); closeManagedApplicationModal()"
                class="w-full rounded-xl border border-cyan-400/30 bg-cyan-950/40 hover:bg-cyan-900/60 px-4 py-2.5 text-xs font-extrabold text-cyan-300 hover:text-white transition-all cursor-pointer shadow-md flex items-center justify-center gap-2"
              >
                <span>✏️ Chỉnh Sửa / Nộp Bổ Sung Bài Báo Cáo</span>
              </button>
            </div>
            <div
              v-else-if="(selectedManagedApplication.status || '').toLowerCase() === 'paid'"
              class="rounded-lg border border-cyan-400/20 bg-cyan-400/10 px-4 py-2.5 text-sm font-semibold text-cyan-200"
            >
              Đã giải ngân
            </div>
            <p v-else class="text-sm font-medium text-slate-500">Completion actions appear after you accept an offer.</p>
          </div>
        </section>

        <section class="rounded-2xl border border-rose-500/20 bg-rose-500/5 p-4">
          <div class="flex items-center justify-between gap-3">
            <div>
              <p class="text-xs font-extrabold uppercase tracking-wider text-rose-300">Hỗ trợ & Khiếu nại (Dispute)</p>
              <p v-if="state.myTickets?.find((t: any) => Number(t.application_id) === Number(selectedManagedApplication.id))" class="mt-1 text-xs font-semibold text-emerald-300">
                Đã có khiếu nại (Trạng thái: {{ state.myTickets.find((t: any) => Number(t.application_id) === Number(selectedManagedApplication.id)).status === 'resolved' ? 'Đã duyệt' : state.myTickets.find((t: any) => Number(t.application_id) === Number(selectedManagedApplication.id)).status === 'rejected' ? 'Bị bác bỏ' : 'Đang xử lý' }})
              </p>
              <p v-else class="mt-1 text-xs font-medium text-slate-400">Gặp sự cố với đơn ứng tuyển này? Gửi khiếu nại tới Ban quản trị.</p>
            </div>

            <button
              v-if="state.myTickets?.find((t: any) => Number(t.application_id) === Number(selectedManagedApplication.id))"
              @click="state.openReappealModal(state.myTickets.find((t: any) => Number(t.application_id) === Number(selectedManagedApplication.id)))"
              class="rounded-xl border border-emerald-400/30 bg-emerald-500/10 px-3.5 py-2 text-xs font-extrabold text-emerald-300 hover:bg-emerald-500/20 hover:text-white transition-all cursor-pointer flex-shrink-0"
            >
              ⚖️ Xem Phán Quyết / Tái Khiếu Nại
            </button>

            <button
              v-else
              @click="openDisputeModal(selectedManagedApplication)"
              class="rounded-xl border border-rose-400/30 bg-rose-500/10 px-3.5 py-2 text-xs font-extrabold text-rose-300 hover:bg-rose-500/20 hover:text-white transition-all cursor-pointer flex-shrink-0"
            >
              ⚠️ Gửi Khiếu Nại
            </button>
          </div>
        </section>

        <section class="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
          <div class="flex flex-wrap items-center justify-between gap-3">
            <div>
              <p class="text-xs font-semibold uppercase tracking-wider text-slate-500">Reviews</p>
              <p class="mt-1 text-sm font-semibold text-white">
                {{ reviewSummary.average_rating.toFixed(1) }} / 5 · {{ reviewSummary.total_reviews }} reviews
              </p>
            </div>
            <div class="flex flex-wrap gap-2">
              <button
                v-if="(selectedManagedApplication.status || '').toLowerCase() === 'paid' && !reviewsList.some(review => Number(review.reviewer_id || review.ReviewerID) === Number(currentUserId))"
                @click="openReviewModal(selectedManagedApplication); closeManagedApplicationModal()"
                class="rounded-lg bg-amber-400 px-3 py-1.5 text-xs font-semibold text-slate-950 hover:bg-amber-300"
              >
                Đánh giá doanh nghiệp
              </button>
              <button
                @click="openReviewsModal(selectedManagedApplication)"
                class="rounded-lg border border-white/10 bg-white/5 px-3 py-1.5 text-xs font-semibold text-slate-200 hover:bg-white/10 hover:text-white"
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
              <p class="text-sm font-semibold text-amber-300">{{ review.rating || review.Rating }} / 5</p>
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
            class="rounded-lg border border-white/10 bg-white/5 px-4 py-2.5 text-sm font-semibold text-slate-200 hover:bg-white/10 hover:text-white"
          >
            Chat with employer
          </button>
          <button
            v-if="(selectedManagedApplication.status || '').toLowerCase() === 'pending'"
            :disabled="isCancellingApp === selectedManagedApplication.id"
            @click="triggerCancelConfirm(selectedManagedApplication.id); closeManagedApplicationModal()"
            class="rounded-lg border border-rose-400/20 bg-rose-500/10 px-4 py-2.5 text-sm font-semibold text-rose-300 hover:bg-rose-500/20 disabled:opacity-50"
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
      class="fixed top-5 right-5 z-[100] max-w-sm w-full bg-slate-900/95 shadow-2xl shadow-slate-950/60 rounded-xl pointer-events-auto ring-1 ring-white/10 overflow-hidden border backdrop-blur animate-in slide-in-from-top-3 duration-200"
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
            <p class="text-sm font-semibold text-white">
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
      <h3 class="text-xl font-semibold text-white">Đánh giá doanh nghiệp</h3>
      <p class="mt-1 text-sm text-slate-400">
        {{ selectedReviewApp.job?.title || 'Công việc đã hoàn thành' }}
      </p>
    </div>

    <form @submit.prevent="submitReview" class="p-6 space-y-5">
      <div>
        <label class="block text-xs font-semibold uppercase tracking-widest text-slate-400 mb-3">
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
        <label class="block text-xs font-semibold uppercase tracking-widest text-slate-400 mb-2">
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
          class="px-4 py-2 rounded-lg bg-slate-800 text-slate-300 text-sm font-semibold hover:bg-slate-700"
        >
          Hủy
        </button>

        <button
          type="submit"
          :disabled="isSubmittingReview"
          class="px-5 py-2 rounded-lg bg-amber-400 text-slate-950 text-sm font-semibold hover:bg-amber-300 disabled:opacity-50"
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
        <p class="text-xs font-semibold uppercase tracking-wider text-cyan-300">Reviews</p>
        <h3 class="mt-1 text-xl font-semibold text-white">{{ selectedReviewsApp?.job?.title || 'Application reviews' }}</h3>
        <p class="mt-1 text-sm font-medium text-slate-400">
          {{ reviewSummary.average_rating.toFixed(1) }} / 5 · {{ reviewSummary.total_reviews }} reviews
        </p>
      </div>
      <button
        type="button"
        @click="closeReviewsModal"
        class="rounded-lg bg-white/5 px-3 py-1.5 text-sm font-semibold text-slate-300 hover:bg-white/10 hover:text-white"
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
          <p class="text-sm font-semibold text-amber-300">{{ review.rating || review.Rating }} / 5</p>
          <p class="text-xs font-semibold uppercase tracking-wider text-slate-500">{{ review.review_type || review.ReviewType || 'review' }}</p>
        </div>
        <p class="mt-2 text-sm font-medium text-slate-300">{{ review.comment || review.Comment || 'Không có nhận xét.' }}</p>
      </div>
      <p v-if="!isLoadingReviews && reviewsList.length === 0" class="rounded-xl border border-white/10 bg-white/[0.04] p-4 text-sm font-medium text-slate-500">
        Chưa có đánh giá cho đơn này.
      </p>
    </div>
  </div>
</div>

<!-- MODAL NỘP BÁO CÁO & BẰNG CHỨNG HOÀN THÀNH CÔNG VIỆC -->
<div v-if="selectedCompletionAppForReport" class="fixed inset-0 z-[9999] flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
  <div class="w-full max-w-lg overflow-hidden rounded-3xl border border-emerald-500/30 bg-slate-900 shadow-2xl shadow-slate-950 animate-in fade-in zoom-in-95 duration-150">
    <!-- Header -->
    <div class="border-b border-white/10 bg-slate-950/80 p-6 flex justify-between items-center">
      <div>
        <span class="text-[10px] font-black uppercase tracking-wider text-emerald-300 bg-emerald-400/10 px-2.5 py-0.5 rounded-full border border-emerald-400/20">
          Nộp Báo Cáo Hoàn Thành
        </span>
        <h3 class="mt-1 text-lg font-extrabold text-white">Báo Cáo & Nộp Sản Phẩm / Bằng Chứng</h3>
        <p class="mt-0.5 text-xs text-slate-400 font-medium">Vị trí: {{ selectedCompletionAppForReport.job?.title || 'Công việc' }}</p>
      </div>
      <button @click="selectedCompletionAppForReport = null" class="text-slate-400 hover:text-white font-bold text-lg cursor-pointer p-1">&times;</button>
    </div>

    <!-- Body -->
    <form @submit.prevent="submitStudentCompletionReport" class="p-6 space-y-4">
      <div>
        <label class="block text-xs font-extrabold text-emerald-300 uppercase tracking-wider mb-1.5">
          Báo cáo kết quả công việc đã làm *
        </label>
        <textarea
          v-model="completionForm.note"
          rows="4"
          required
          class="w-full resize-none rounded-2xl border border-white/10 bg-slate-950 px-4 py-3 text-xs text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-emerald-400 font-medium leading-relaxed"
          placeholder="Mô tả công việc bạn đã hoàn thành (số ca làm, số giờ làm việc, kết quả thực hiện, ghi chú gửi HR...)"
        ></textarea>
      </div>

      <div>
        <label class="block text-xs font-extrabold text-cyan-300 uppercase tracking-wider mb-1.5">
          File Bài Nộp / Bằng Chứng Hoàn Thành
        </label>
        <div class="space-y-2">
          <div class="flex items-center gap-2">
            <input
              type="file"
              ref="studentCompletionProofFileInput"
              @change="$emit('upload-completion-proof', $event)"
              accept="image/*,.pdf,.doc,.docx,.zip,.rar,.txt"
              class="hidden"
            />
            <button
              type="button"
              @click="studentCompletionProofFileInput?.click()"
              :disabled="state.isUploadingCompletionProof"
              class="flex-1 px-4 py-2.5 rounded-xl border border-dashed border-cyan-400/40 bg-cyan-950/20 hover:bg-cyan-900/40 text-cyan-300 text-xs font-bold transition-all flex items-center justify-center gap-2 cursor-pointer disabled:opacity-50"
            >
              <span>{{ state.isUploadingCompletionProof ? '⏳ Đang tải file lên hệ thống...' : '📁 Tải File Bài Nộp Từ Máy Tính (Ảnh/PDF/Doc/Zip)' }}</span>
            </button>
          </div>

          <div v-if="completionForm.proofUrl" class="flex items-center justify-between p-2.5 rounded-xl bg-cyan-950/60 border border-cyan-400/30 text-xs">
            <div class="flex items-center gap-2 overflow-hidden pr-2">
              <span class="text-cyan-400 font-bold">✅ File bài nộp:</span>
              <a :href="completionForm.proofUrl" target="_blank" class="text-cyan-200 underline truncate hover:text-white">{{ completionForm.proofUrl }}</a>
            </div>
            <button type="button" @click="completionForm.proofUrl = ''" class="text-rose-400 hover:text-rose-300 font-bold px-2 py-1 bg-rose-950/50 rounded-lg shrink-0">
              ✕ Xóa
            </button>
          </div>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="pt-4 border-t border-white/10 flex justify-end gap-3">
        <button
          type="button"
          @click="selectedCompletionAppForReport = null"
          class="px-5 py-2.5 text-xs font-bold text-slate-300 hover:bg-white/10 border border-white/10 rounded-xl transition-all cursor-pointer"
        >
          Hủy bỏ
        </button>
        <button
          type="submit"
          :disabled="isSubmittingCompletion"
          class="px-6 py-2.5 text-xs font-black text-slate-950 bg-gradient-to-r from-emerald-400 to-cyan-400 hover:brightness-110 disabled:opacity-50 rounded-xl shadow-lg shadow-emerald-500/20 transition-all cursor-pointer uppercase tracking-wider"
        >
          {{ isSubmittingCompletion ? 'Đang gửi...' : 'Gửi Báo Cáo & Báo Hoàn Thành' }}
        </button>
      </div>
    </form>
  </div>
</div>

  <!-- Check-Out Confirmation Modal -->
  <div v-if="showCheckOutConfirmModal" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
    <div class="bg-slate-900 rounded-3xl border border-rose-500/30 shadow-2xl shadow-slate-950 max-w-md w-full overflow-hidden animate-in fade-in zoom-in-95 duration-200">
      <div class="p-6 text-center space-y-4">
        <div class="mx-auto flex h-14 w-14 items-center justify-center rounded-2xl bg-rose-500/10 border border-rose-500/30 text-rose-400 text-2xl shadow-lg shadow-rose-500/20">
          🛑
        </div>
        <div>
          <h3 class="text-lg font-extrabold text-white">Xác Nhận Kết Thúc Ca Làm</h3>
          <p class="mt-2 text-xs font-semibold text-slate-300 leading-relaxed">
            Bạn có chắc chắn muốn kết thúc ca làm việc hiện tại không? Hệ thống sẽ ghi nhận tổng thời gian làm việc và đồng bộ báo cáo.
          </p>
        </div>
        <div class="flex items-center justify-center gap-3 pt-2">
          <button
            @click="showCheckOutConfirmModal = false"
            class="w-full py-2.5 rounded-xl border border-white/10 bg-slate-950 text-slate-300 text-xs font-extrabold hover:bg-white/10 transition-all cursor-pointer"
          >
            Hủy Bỏ
          </button>
          <button
            @click="executeCheckOut"
            :disabled="isSubmittingCheckOut"
            class="w-full py-2.5 rounded-xl bg-gradient-to-r from-rose-500 to-red-600 hover:from-rose-400 hover:to-red-500 text-white text-xs font-extrabold shadow-lg shadow-rose-500/30 transition-all disabled:opacity-50 cursor-pointer"
          >
            {{ isSubmittingCheckOut ? 'Đang xử lý...' : '🛑 Kết Thúc Ca' }}
          </button>
        </div>
      </div>
    </div>
  </div>

  <!-- Student Dispute Ticket Modal -->
  <div v-if="isDisputeModalOpen && disputeTargetApp" class="fixed inset-0 z-[10001] flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
    <div class="bg-slate-900 rounded-3xl border border-rose-500/30 shadow-2xl shadow-slate-950 max-w-lg w-full overflow-hidden animate-in fade-in zoom-in-95 duration-200">
      <div class="p-5 border-b border-white/10 bg-slate-950/80 flex justify-between items-center">
        <div>
          <h3 class="text-base font-extrabold text-white">⚠️ Gửi Khiếu Nại Đơn Ứng Tuyển</h3>
          <p class="text-xs text-rose-300 font-semibold mt-0.5">Mã đơn #{{ disputeTargetApp.id }} · {{ disputeTargetApp.job?.title || 'Công việc' }}</p>
        </div>
        <button @click="isDisputeModalOpen = false" class="text-slate-400 hover:text-white text-xl font-bold p-1 rounded-lg hover:bg-white/10 cursor-pointer">&times;</button>
      </div>

      <form @submit.prevent="submitDisputeTicket" class="p-6 space-y-4">
        <div>
          <label class="text-xs font-bold text-slate-300 uppercase tracking-wider block mb-1.5">Lý do khiếu nại <span class="text-rose-400">*</span></label>
          <select
            v-model="disputeForm.reason"
            required
            class="w-full text-xs px-3.5 py-2.5 rounded-xl border border-white/15 bg-slate-950 text-white focus:outline-none focus:border-rose-400"
          >
            <option value="" disabled>-- Chọn lý do khiếu nại --</option>
            <option value="Không thanh toán lương đúng hẹn">Không thanh toán lương đúng hẹn</option>
            <option value="Yêu cầu công việc không đúng mô tả">Yêu cầu công việc không đúng mô tả</option>
            <option value="Thái độ cư xử vi phạm quy định">Thái độ cư xử vi phạm quy định</option>
            <option value="Khác">Lý do khác</option>
          </select>
        </div>

        <div>
          <label class="text-xs font-bold text-slate-300 uppercase tracking-wider block mb-1.5">Mô tả chi tiết sự việc <span class="text-rose-400">*</span></label>
          <textarea
            v-model="disputeForm.description"
            rows="3"
            required
            placeholder="Mô tả cụ thể sự việc, thời gian diễn ra và thông tin liên quan để Ban quản trị đối soát..."
            class="w-full text-xs p-3 rounded-xl border border-white/15 bg-slate-950 text-white placeholder-slate-500 focus:outline-none focus:border-rose-400"
          ></textarea>
        </div>

        <div>
          <label class="text-xs font-bold text-slate-300 uppercase tracking-wider block mb-1.5">Đề xuất xử lý với Admin (Tùy chọn)</label>
          <input
            v-model="disputeForm.requested_action"
            type="text"
            placeholder="VD: Yêu cầu hoàn tiền Escrow / Cấp lại chứng nhận..."
            class="w-full text-xs px-3.5 py-2.5 rounded-xl border border-white/15 bg-slate-950 text-white placeholder-slate-500 focus:outline-none focus:border-rose-400 font-medium"
          />
        </div>

        <div>
          <label class="text-xs font-bold text-cyan-300 uppercase tracking-wider block mb-1.5">File/Ảnh Bằng Chứng (Tùy chọn)</label>
          <div class="space-y-2">
            <div class="flex items-center gap-2">
              <input
                type="file"
                ref="studentEvidenceFileInput"
                @change="$emit('upload-evidence', $event)"
                accept="image/*,.pdf,.doc,.docx,.zip"
                class="hidden"
              />
              <button
                type="button"
                @click="studentEvidenceFileInput?.click()"
                :disabled="state.isUploadingEvidence"
                class="flex-1 px-4 py-2.5 rounded-xl border border-dashed border-cyan-400/40 bg-cyan-950/20 hover:bg-cyan-900/40 text-cyan-300 text-xs font-bold transition-all flex items-center justify-center gap-2 cursor-pointer disabled:opacity-50"
              >
                <span>{{ state.isUploadingEvidence ? '⏳ Đang tải file lên hệ thống...' : '📁 Tải File Bằng Chứng Từ Máy Tính (Ảnh/PDF/Doc/Zip)' }}</span>
              </button>
            </div>

            <div v-if="disputeForm.evidence_url" class="flex items-center justify-between p-2.5 rounded-xl bg-cyan-950/60 border border-cyan-400/30 text-xs">
              <div class="flex items-center gap-2 overflow-hidden pr-2">
                <span class="text-cyan-400 font-bold">✅ Bằng chứng:</span>
                <a :href="disputeForm.evidence_url" target="_blank" class="text-cyan-200 underline truncate hover:text-white">{{ disputeForm.evidence_url }}</a>
              </div>
              <button type="button" @click="disputeForm.evidence_url = ''" class="text-rose-400 hover:text-rose-300 font-bold px-2 py-1 bg-rose-950/50 rounded-lg shrink-0">
                ✕ Xóa
              </button>
            </div>
          </div>
        </div>

        <div class="pt-3 border-t border-white/10 flex justify-end gap-3">
          <button
            type="button"
            @click="isDisputeModalOpen = false"
            class="px-4 py-2 text-xs font-bold text-slate-300 hover:bg-white/10 border border-white/10 rounded-xl transition-all cursor-pointer"
          >
            Hủy bỏ
          </button>
          <button
            type="submit"
            :disabled="isSubmittingDispute"
            class="px-5 py-2 text-xs font-extrabold text-white bg-rose-500 hover:bg-rose-400 disabled:opacity-50 rounded-xl shadow-lg shadow-rose-500/20 transition-all cursor-pointer uppercase tracking-wider"
          >
            {{ isSubmittingDispute ? 'Đang gửi...' : '⚠️ Xác Nhận Gửi Khiếu Nại' }}
          </button>
        </div>
      </form>
    </div>
  </div>

  <!-- Student Dispute Error/Warning Modal -->
  <div v-if="showDisputeErrorModal" class="fixed inset-0 z-[10003] flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
    <div class="bg-slate-900 rounded-3xl border border-rose-500/30 shadow-2xl shadow-slate-950 max-w-md w-full overflow-hidden animate-in fade-in zoom-in-95 duration-200">
      <div class="p-6 text-center space-y-4">
        <div class="mx-auto flex h-16 w-16 items-center justify-center rounded-2xl bg-amber-500/10 border border-amber-500/30 text-amber-400 text-3xl shadow-lg shadow-amber-500/20">
          ⚠️
        </div>
        <div>
          <h3 class="text-lg font-extrabold text-white">Thông Báo Khiếu Nại</h3>
          <p class="mt-2 text-xs font-semibold text-slate-300 leading-relaxed">
            {{ disputeErrorMsg }}
          </p>
        </div>
        <div class="pt-2">
          <button
            @click="showDisputeErrorModal = false"
            class="w-full py-2.5 rounded-xl bg-slate-800 hover:bg-slate-700 text-white text-xs font-extrabold shadow-lg transition-all cursor-pointer uppercase tracking-wider"
          >
            Đã Hiểu / Đóng
          </button>
        </div>
      </div>
    </div>
  </div>

  <!-- Student Reappeal / View Verdict Modal -->
  <div v-if="state.showReappealModal" class="fixed inset-0 z-[10002] flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
    <div class="bg-slate-900 rounded-3xl border border-white/10 shadow-2xl shadow-slate-950 max-w-lg w-full overflow-hidden animate-in fade-in zoom-in-95 duration-200">
      <div class="p-6 space-y-4">
        <div class="flex items-center justify-between border-b border-white/10 pb-3">
          <h3 class="text-sm font-extrabold text-white flex items-center gap-2">
            <span>⚖️ KẾT QUẢ KHIẾU NẠI & PHÁN QUYẾT ADMIN</span>
          </h3>
          <button @click="state.showReappealModal = false" class="text-slate-400 hover:text-white font-bold text-lg">✕</button>
        </div>

        <div v-if="state.selectedTicketForView" class="space-y-3">
          <div class="bg-slate-950/60 p-3 rounded-xl border border-white/10 space-y-1.5 text-xs">
            <div class="flex justify-between text-slate-400">
              <span>Đơn ứng tuyển: <strong class="text-white">#{{ state.selectedTicketForView.application_id }}</strong></span>
              <span>Trạng thái: 
                <span :class="state.selectedTicketForView.status === 'resolved' ? 'text-emerald-400 font-bold' : state.selectedTicketForView.status === 'rejected' ? 'text-rose-400 font-bold' : 'text-amber-400 font-bold'">
                  {{ state.selectedTicketForView.status === 'resolved' ? 'Đã duyệt' : state.selectedTicketForView.status === 'rejected' ? 'Bị bác bỏ' : 'Đang xử lý' }}
                </span>
              </span>
            </div>
            <div><strong class="text-slate-300">Lý do ban đầu:</strong> {{ state.selectedTicketForView.reason }}</div>
            <div><strong class="text-slate-300">Mô tả:</strong> {{ state.selectedTicketForView.description }}</div>
          </div>

          <div v-if="state.selectedTicketForView.verdict" class="bg-emerald-950/40 p-4 rounded-xl border border-emerald-500/30 space-y-1">
            <div class="text-xs font-bold text-emerald-400 uppercase tracking-wider">Phán quyết chính thức từ Admin:</div>
            <div class="text-sm font-semibold text-white leading-relaxed">{{ state.selectedTicketForView.verdict }}</div>
          </div>

          <div v-if="state.selectedTicketForView.is_reappealed" class="bg-amber-950/40 p-3 rounded-xl border border-amber-500/30 text-xs text-amber-200">
            <div class="font-bold text-amber-300">⚠️ Đã gửi yêu cầu Tái xem xét phán quyết</div>
            <div class="mt-1">Lý do phản hồi: {{ state.selectedTicketForView.reappeal_reason }}</div>
          </div>

          <!-- Form Reappeal if verdict exists and not reappealed -->
          <div v-else-if="state.selectedTicketForView.verdict" class="pt-2 space-y-3 border-t border-white/10">
            <div class="text-xs font-bold text-rose-300">Bạn thấy phán quyết chưa thỏa đáng? Nộp yêu cầu Tái xem xét:</div>
            <textarea
              v-model="state.reappealForm.reason"
              rows="3"
              placeholder="Nhập lý do chi tiết giải thích vì sao phán quyết chưa công bằng/thỏa đáng..."
              class="w-full text-xs p-3 rounded-xl border border-white/15 bg-slate-950 text-white placeholder-slate-500 focus:outline-none focus:border-rose-400"
            ></textarea>

            <div class="flex justify-end gap-3 pt-2">
              <button
                type="button"
                @click="state.showReappealModal = false"
                class="px-4 py-2 text-xs font-bold text-slate-300 hover:bg-white/10 border border-white/10 rounded-xl"
              >
                Đóng
              </button>
              <button
                type="button"
                @click="state.submitReappealTicket()"
                :disabled="state.isSubmittingReappeal"
                class="px-4 py-2 text-xs font-extrabold text-white bg-rose-500 hover:bg-rose-400 disabled:opacity-50 rounded-xl shadow-lg shadow-rose-500/20 uppercase tracking-wider cursor-pointer"
              >
                {{ state.isSubmittingReappeal ? 'Đang gửi...' : '⚠️ Nộp Yêu Cầu Tái Xem Xét' }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, toRefs } from 'vue'
import ChatBox from '~/components/ChatBox.vue'

const props = defineProps<{ state: Record<string, any> }>()
const studentEvidenceFileInput = ref<HTMLInputElement | null>(null)
const studentCompletionProofFileInput = ref<HTMLInputElement | null>(null)
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
  triggerCheckOutConfirm,
  executeCheckOut,
  showCheckOutConfirmModal,
  isSubmittingCheckOut,
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
  selectedCompletionAppForReport,
  completionForm,
  isSubmittingCompletion,
  isUploadingCompletionProof,
  openStudentCompletionModal,
  submitStudentCompletionReport,
  isDisputeModalOpen,
  disputeTargetApp,
  isSubmittingDispute,
  disputeForm,
  openDisputeModal,
  submitDisputeTicket,
  showDisputeErrorModal,
  disputeErrorMsg,
  myTickets,
  showReappealModal,
  openReappealModal,
  submitReappealTicket,
  openReviewModal,
  toast
} = toRefs(props.state)
</script>
