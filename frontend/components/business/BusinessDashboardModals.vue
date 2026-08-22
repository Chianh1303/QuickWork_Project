<template>
  <div>
    <!-- MODAL 1: REVIEW & OFFER MODAL -->
    <div v-if="selectedApp" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
      <div class="bg-slate-950 rounded-3xl border border-cyan-400/30 shadow-2xl max-w-lg w-full overflow-hidden flex flex-col max-h-[90vh] text-left animate-in fade-in zoom-in-95 duration-150">
        
        <!-- Header -->
        <div class="p-6 border-b border-white/10 flex justify-between items-center bg-slate-900/90">
          <div>
            <span class="text-[10px] font-black uppercase tracking-wider text-cyan-300 bg-cyan-400/10 px-2.5 py-0.5 rounded-full border border-cyan-400/20">
              Đánh Giá Hồ Sơ & Gửi Offer
            </span>
            <h3 class="text-lg font-extrabold text-white mt-1">Chi Tiết Ứng Viên & Phản Hồi</h3>
            <p class="text-xs text-slate-400 font-medium mt-0.5">Ứng viên: <span class="text-cyan-300 font-extrabold">{{ selectedApp.student?.full_name || 'N/A' }}</span></p>
          </div>
          <button @click="closeModal" class="text-slate-400 hover:text-white font-bold text-xl cursor-pointer p-1">&times;</button>
        </div>

        <!-- Body -->
        <div class="p-6 space-y-5 overflow-y-auto flex-1">
          <!-- Cover Note -->
          <div>
            <label class="text-xs font-extrabold text-cyan-300 uppercase tracking-wider block mb-1.5">Lời nhắn từ ứng viên (Cover Note)</label>
            <div class="p-4 bg-slate-900 border border-white/10 rounded-2xl text-xs font-medium text-slate-200 whitespace-pre-line leading-relaxed">
              {{ selectedApp.cover_note || 'Ứng viên không để lại lời nhắn.' }}
            </div>
          </div>

          <!-- Decision Selection -->
          <div>
            <label class="text-xs font-extrabold text-cyan-300 uppercase tracking-wider block mb-2">Quyết định xét duyệt</label>
            <div class="grid grid-cols-2 gap-3">
              <button 
                type="button"
                @click="reviewStatus = 'approved'"
                :class="[reviewStatus === 'approved' ? 'border-emerald-400 bg-emerald-500/20 text-emerald-300 ring-2 ring-emerald-400/30' : 'border-white/10 bg-slate-900 text-slate-300 hover:bg-white/5']"
                class="py-3 px-4 text-xs font-extrabold border rounded-2xl transition-all flex items-center justify-center cursor-pointer"
              >
                Chấp nhận & Gửi Offer
              </button>
              <button 
                type="button"
                @click="reviewStatus = 'rejected'"
                :class="[reviewStatus === 'rejected' ? 'border-rose-400 bg-rose-500/20 text-rose-300 ring-2 ring-rose-400/30' : 'border-white/10 bg-slate-900 text-slate-300 hover:bg-white/5']"
                class="py-3 px-4 text-xs font-extrabold border rounded-2xl transition-all flex items-center justify-center cursor-pointer"
              >
                Từ chối đơn
              </button>
            </div>
          </div>

          <!-- Offer Input Section (Shown when Approved) -->
          <div v-if="reviewStatus === 'approved'" class="p-4.5 bg-emerald-500/10 border border-emerald-500/20 rounded-2xl space-y-4 animate-in fade-in duration-150">
            <h4 class="text-xs font-extrabold text-emerald-300">
              Chi Tiết Offer Gửi Sinh Viên
            </h4>
            
            <div class="grid grid-cols-2 gap-3">
              <div>
                <label class="text-[11px] font-bold text-slate-300 block mb-1">Mức lương Offer (VNĐ)</label>
                <input 
                  v-model="offerForm.salary"
                  type="text" 
                  placeholder="VD: 5,000,000 VNĐ"
                  class="w-full text-xs px-3.5 py-2.5 border border-white/10 rounded-xl bg-slate-900 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-emerald-400 font-bold"
                />
              </div>
              <div>
                <label class="text-[11px] font-bold text-slate-300 block mb-1">Ngày bắt đầu ca làm</label>
                <input 
                  v-model="offerForm.startDate"
                  type="date" 
                  class="w-full text-xs px-3.5 py-2.5 border border-white/10 rounded-xl bg-slate-900 text-white focus:outline-none focus:ring-2 focus:ring-emerald-400 font-semibold"
                />
              </div>
            </div>

            <div>
              <label class="text-[11px] font-bold text-slate-300 block mb-1">Lời nhắn chào mừng từ Doanh nghiệp</label>
              <textarea 
                v-model="offerForm.message"
                rows="3"
                placeholder="Chào mừng bạn gia nhập đội ngũ, vui lòng xem thông tin chi tiết ca làm..."
                class="w-full text-xs p-3 border border-white/10 rounded-xl bg-slate-900 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-emerald-400 leading-relaxed font-medium"
              ></textarea>
            </div>
          </div>
        </div>

        <!-- Footer Actions -->
        <div class="p-5 bg-slate-900/90 border-t border-white/10 flex items-center justify-end space-x-3">
          <button 
            @click="closeModal" 
            class="px-5 py-2.5 text-xs font-bold text-slate-300 hover:bg-white/10 border border-white/10 rounded-xl transition-all cursor-pointer"
          >
            Hủy bỏ
          </button>
          <button 
            @click="submitReview"
            :disabled="!reviewStatus || isSubmitting"
            class="px-6 py-2.5 text-xs font-black text-slate-950 bg-gradient-to-r from-cyan-400 to-emerald-400 hover:brightness-110 disabled:opacity-50 rounded-xl shadow-lg shadow-cyan-500/25 transition-all cursor-pointer"
          >
            {{ isSubmitting ? 'Đang gửi...' : 'Xác Nhận Phản Hồi' }}
          </button>
        </div>

      </div>
    </div>

    <!-- MODAL 2: COMPLETION & ESCROW RELEASE MODAL -->
    <div v-if="selectedCompletionApp" class="fixed inset-0 z-[9999] flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
      <div class="w-full max-w-xl overflow-hidden rounded-3xl border border-cyan-400/30 bg-slate-950 shadow-2xl animate-in fade-in zoom-in-95 duration-150">
        <!-- Header -->
        <div class="border-b border-white/10 bg-slate-900/90 p-6">
          <div class="flex items-start justify-between gap-4">
            <div>
              <span class="text-[10px] font-black uppercase tracking-wider text-cyan-300 bg-cyan-400/10 px-2.5 py-0.5 rounded-full border border-cyan-400/20">
                Xác Nhận Hoàn Thành
              </span>
              <h3 class="mt-1 text-xl font-extrabold text-white">Xác Nhận Hoàn Thành & Giải Ngân Escrow</h3>
              <p class="mt-1 text-xs font-semibold text-slate-400">
                Sinh viên đã hoàn tất ca làm. Vui lòng đối soát trước khi xác nhận thanh toán cuối cùng.
              </p>
            </div>
            <button
              @click="closeCompletionModal"
              class="rounded-xl bg-white/5 border border-white/10 px-3 py-1.5 text-xs font-bold text-slate-300 hover:bg-white/10 hover:text-white cursor-pointer"
            >
              Đóng
            </button>
          </div>
        </div>

        <!-- Body -->
        <div class="space-y-4 p-6">
          <div class="grid gap-3 sm:grid-cols-2">
            <div class="rounded-2xl border border-white/10 bg-slate-900 p-4">
              <p class="text-[10px] font-black uppercase tracking-wider text-slate-400">Ứng viên nhận việc</p>
              <p class="mt-1.5 text-base font-extrabold text-white">{{ selectedCompletionApp.student?.full_name || 'N/A' }}</p>
              <p class="mt-1 text-xs font-medium text-slate-400">SĐT: {{ selectedCompletionApp.student?.phone || 'Chưa có số điện thoại' }}</p>
            </div>
            <div class="rounded-2xl border border-white/10 bg-slate-900 p-4">
              <p class="text-[10px] font-black uppercase tracking-wider text-slate-400">Vị trí công việc</p>
              <p class="mt-1.5 text-base font-extrabold text-white">{{ selectedCompletionApp.job?.title || jobTitleLookup(selectedCompletionApp.job_id) }}</p>
              <p class="mt-1 text-xs font-medium text-slate-400">Mã đơn #{{ selectedCompletionApp.id }}</p>
            </div>
          </div>

          <div class="rounded-2xl border border-cyan-400/30 bg-cyan-400/10 p-4.5">
            <div class="grid gap-3 sm:grid-cols-3">
              <div>
                <p class="text-[10px] font-black uppercase tracking-wider text-cyan-300">Mức lương Offer</p>
                <p class="mt-1 text-sm font-black text-white">{{ selectedCompletionApp.offer_salary || 'Thỏa thuận' }}</p>
              </div>
              <div>
                <p class="text-[10px] font-black uppercase tracking-wider text-cyan-300">Ngày bắt đầu</p>
                <p class="mt-1 text-sm font-black text-white">{{ selectedCompletionApp.offer_start_date || 'N/A' }}</p>
              </div>
              <div>
                <p class="text-[10px] font-black uppercase tracking-wider text-cyan-300">Smart Escrow</p>
                <p class="mt-1 text-sm font-black text-emerald-300">Sẵn sàng giải ngân</p>
              </div>
            </div>
          </div>

          <div class="rounded-2xl border border-amber-400/30 bg-amber-400/10 p-4 text-xs font-semibold text-amber-200 leading-relaxed">
            Sau khi bấm xác nhận, trạng thái đơn sẽ chuyển sang <strong class="text-white">paid</strong>, hệ thống sẽ xác nhận hoàn tất ca làm và ghi nhận giải ngân cho sinh viên.
          </div>
        </div>

        <!-- Footer -->
        <div class="flex flex-col-reverse gap-3 border-t border-white/10 bg-slate-900/90 p-5 sm:flex-row sm:justify-end">
          <button
            @click="closeCompletionModal"
            class="rounded-xl border border-white/10 px-5 py-2.5 text-xs font-bold text-slate-300 hover:bg-white/10 cursor-pointer"
          >
            Kiểm tra lại
          </button>
          <button
            @click="submitBusinessCompletion"
            :disabled="isCompletingJob"
            class="rounded-xl bg-gradient-to-r from-cyan-400 to-emerald-400 px-6 py-2.5 text-xs font-black text-slate-950 shadow-lg shadow-cyan-500/25 hover:brightness-110 disabled:opacity-50 cursor-pointer"
          >
            {{ isCompletingJob ? 'Đang xác nhận...' : 'Xác Nhận & Giải Ngân Escrow' }}
          </button>
        </div>
      </div>
    </div>

    <!-- MODAL 3: MANAGED APPLICANT SLIDE-OVER DRAWER -->
    <div v-if="selectedManagedApplicant" class="fixed inset-0 z-[9998] flex justify-end bg-slate-950/80 backdrop-blur-md">
      <button
        type="button"
        class="absolute inset-0 cursor-default"
        aria-label="Đóng chi tiết ứng viên"
        @click="closeManagedApplicantModal"
      ></button>

      <aside class="relative flex h-full w-full max-w-2xl flex-col overflow-hidden border-l border-cyan-400/30 bg-slate-950 shadow-2xl">
        <!-- Drawer Header -->
        <div class="border-b border-white/10 bg-slate-900/90 p-6">
          <div class="flex items-start justify-between gap-4">
            <div class="min-w-0">
              <span class="text-[10px] font-black uppercase tracking-wider text-cyan-300 bg-cyan-400/10 px-2.5 py-0.5 rounded-full border border-cyan-400/20">
                Hồ Sơ Ứng Viên Chi Tiết
              </span>
              <h3 class="mt-1.5 truncate text-2xl font-extrabold text-white">
                {{ selectedManagedApplicant.student?.full_name || 'Ứng viên' }}
              </h3>
              <p class="mt-1 text-xs font-semibold text-slate-400">
                Mã đơn #{{ selectedManagedApplicant.id }} · {{ selectedManagedApplicant.job?.title || jobTitleLookup(selectedManagedApplicant.job_id) }}
              </p>
            </div>
            <button
              type="button"
              @click="closeManagedApplicantModal"
              class="rounded-xl bg-white/5 border border-white/10 px-3 py-1.5 text-xs font-bold text-slate-300 hover:bg-white/10 hover:text-white cursor-pointer"
            >
              Đóng
            </button>
          </div>
        </div>

        <!-- Drawer Content Scroll -->
        <div class="flex-1 space-y-5 overflow-y-auto p-6">
          <!-- Section A: Student Summary -->
          <section class="rounded-2xl border border-white/10 bg-slate-900/90 p-5 space-y-4">
            <div class="flex items-start gap-4">
              <img
                v-if="selectedManagedApplicant.student?.avatar_url"
                :src="selectedManagedApplicant.student.avatar_url"
                class="h-16 w-16 rounded-2xl border-2 border-cyan-400/30 object-cover shadow-md"
              />
              <div
                v-else
                class="flex h-16 w-16 flex-shrink-0 items-center justify-center rounded-2xl border-2 border-cyan-400/30 bg-gradient-to-br from-indigo-600 to-cyan-500 text-xl font-black text-white shadow-md"
              >
                {{ (selectedManagedApplicant.student?.full_name || 'S').slice(0, 1).toUpperCase() }}
              </div>

              <div class="min-w-0 flex-1 space-y-1">
                <div class="flex flex-wrap items-center gap-2">
                  <h4 class="text-lg font-extrabold text-white">{{ selectedManagedApplicant.student?.full_name || 'N/A' }}</h4>
                  <span :class="[
                    statusBadgeClass(selectedManagedApplicant.status),
                    'inline-flex items-center rounded-full border px-2.5 py-0.5 text-[10px] font-extrabold uppercase tracking-wider'
                  ]">
                    {{ selectedManagedApplicant.status ? selectedManagedApplicant.status.replace(/_/g, ' ') : 'Chờ duyệt' }}
                  </span>
                </div>
                <p class="text-xs font-semibold text-slate-300">
                  SĐT: {{ selectedManagedApplicant.student?.phone || 'Chưa có số điện thoại' }}
                </p>
                <a
                  v-if="selectedManagedApplicant.student?.cv_url"
                  :href="selectedManagedApplicant.student.cv_url"
                  target="_blank"
                  class="mt-2 inline-flex items-center gap-1 rounded-xl border border-rose-500/20 bg-rose-500/10 px-3 py-1 text-xs font-extrabold text-rose-300 hover:bg-rose-500/20 hover:text-white transition-all shadow-sm"
                >
                  Xem Hồ Sơ CV
                </a>
              </div>
            </div>

            <div class="mt-3 flex flex-wrap gap-1.5 pt-3 border-t border-white/5">
              <span
                v-for="(skill, index) in parseSkills(selectedManagedApplicant.student?.skills)"
                :key="index"
                class="rounded-lg border border-indigo-500/20 bg-slate-950 px-2.5 py-1 text-xs font-bold text-indigo-300"
              >
                {{ skill }}
              </span>
              <span
                v-if="!selectedManagedApplicant.student?.skills || parseSkills(selectedManagedApplicant.student?.skills).length === 0"
                class="text-xs font-medium text-slate-500 italic"
              >
                Chưa cập nhật kỹ năng
              </span>
            </div>
          </section>

          <!-- Section B: Job & Payment Info -->
          <section class="grid gap-4 sm:grid-cols-2">
            <div class="rounded-2xl border border-white/10 bg-slate-900/90 p-5 space-y-2">
              <p class="text-[10px] font-black uppercase tracking-wider text-cyan-300">Thông tin Vị trí công việc</p>
              <h4 class="text-sm font-extrabold text-white">
                {{ selectedManagedApplicant.job?.title || jobTitleLookup(selectedManagedApplicant.job_id) }}
              </h4>
              <div class="space-y-1 text-xs font-medium text-slate-300 pt-1">
                <p>Địa điểm: <span class="font-bold text-white">{{ selectedManagedApplicant.job?.location || 'N/A' }}</span></p>
                <p>Mức lương: <span class="font-bold text-emerald-300">{{ selectedManagedApplicant.job?.salary ? Number(selectedManagedApplicant.job.salary).toLocaleString('vi-VN') + ' VNĐ' : 'Thỏa thuận' }}</span></p>
                <p>Ngày nộp đơn: <span class="font-bold text-slate-300">{{ formatDate(selectedManagedApplicant.applied_at || selectedManagedApplicant.id) }}</span></p>
              </div>
            </div>

            <div class="rounded-2xl border border-white/10 bg-slate-900/90 p-5 space-y-2">
              <p class="text-[10px] font-black uppercase tracking-wider text-cyan-300">Trạng thái Thanh toán Escrow</p>
              <div
                v-if="selectedManagedApplicant.status?.toLowerCase() === 'paid'"
                class="rounded-xl border border-emerald-400/30 bg-emerald-400/10 p-3"
              >
                <p class="text-xs font-black text-emerald-300 uppercase">Đã giải ngân Escrow</p>
                <p class="mt-0.5 text-[11px] font-medium text-slate-300">Lương ca làm đã được thanh toán cho ứng viên.</p>
              </div>
              <div
                v-else-if="selectedManagedApplicant.status?.toLowerCase() === 'student_completed'"
                class="rounded-xl border border-amber-400/30 bg-amber-400/10 p-3"
              >
                <p class="text-xs font-black text-amber-300 uppercase">Chờ Doanh nghiệp xác nhận</p>
                <p class="mt-0.5 text-[11px] font-medium text-slate-300">Sinh viên đã báo hoàn thành công việc.</p>
              </div>
              <div v-else class="rounded-xl border border-white/10 bg-slate-950 p-3">
                <p class="text-xs font-bold text-slate-400">Chưa giải ngân</p>
                <p class="mt-0.5 text-[11px] font-medium text-slate-500">Nút giải ngân sẽ mở khi ca làm hoàn tất.</p>
              </div>
            </div>
          </section>

          <!-- Section C: Offer Info -->
          <section class="rounded-2xl border border-emerald-500/20 bg-slate-900/90 p-5 space-y-3">
            <p class="text-[10px] font-black uppercase tracking-wider text-emerald-300">Thông tin Offer Đã Gửi</p>
            <div
              v-if="selectedManagedApplicant.offer_salary || selectedManagedApplicant.offer_start_date || selectedManagedApplicant.offer_message"
              class="grid gap-3 sm:grid-cols-2"
            >
              <div class="rounded-xl border border-emerald-400/20 bg-emerald-500/10 p-3">
                <p class="text-[10px] font-bold uppercase text-emerald-300">Mức lương Offer</p>
                <p class="mt-0.5 text-xs font-black text-white">{{ selectedManagedApplicant.offer_salary || 'Thỏa thuận' }}</p>
              </div>
              <div class="rounded-xl border border-cyan-400/20 bg-cyan-500/10 p-3">
                <p class="text-[10px] font-bold uppercase text-cyan-300">Ngày bắt đầu</p>
                <p class="mt-0.5 text-xs font-black text-white">{{ selectedManagedApplicant.offer_start_date || 'N/A' }}</p>
              </div>
              <div class="rounded-xl border border-white/10 bg-slate-950 p-3 sm:col-span-2">
                <p class="text-[10px] font-bold uppercase text-slate-400">Lời nhắn HR</p>
                <p class="mt-1 whitespace-pre-line text-xs font-medium text-slate-300 leading-relaxed">{{ selectedManagedApplicant.offer_message || 'Không có lời nhắn offer.' }}</p>
              </div>
            </div>
            <p v-else class="text-xs font-medium text-slate-500 italic">Chưa có thông tin offer nào được ghi nhận.</p>
          </section>

          <!-- Section D: Reviews & Ratings -->
          <section class="rounded-2xl border border-amber-500/20 bg-slate-900/90 p-5 space-y-3">
            <div class="flex items-center justify-between gap-3">
              <div>
                <p class="text-[10px] font-black uppercase tracking-wider text-amber-300">Đánh giá & Nhận xét</p>
                <p class="mt-0.5 text-sm font-black text-white">
                  {{ reviewSummary.average_rating.toFixed(1) }} / 5 · {{ reviewSummary.total_reviews }} lượt đánh giá
                </p>
              </div>
              <button
                @click="openReviewsModal(selectedManagedApplicant)"
                class="rounded-xl border border-white/10 bg-white/5 px-3 py-1.5 text-xs font-bold text-slate-300 hover:bg-white/10 hover:text-white cursor-pointer"
              >
                Xem lịch sử
              </button>
            </div>
            <div class="space-y-2">
              <div v-if="isLoadingReviews" class="text-xs font-medium text-slate-500">Đang tải đánh giá...</div>
              <div
                v-for="review in reviewsList.slice(0, 3)"
                :key="review.id || review.ID"
                class="rounded-xl border border-white/10 bg-slate-950 p-3"
              >
                <div class="flex items-center justify-between gap-2">
                  <p class="text-xs font-black text-amber-300">{{ review.rating || review.Rating }} / 5</p>
                  <p class="text-[10px] font-bold uppercase text-slate-500">{{ review.review_type || review.ReviewType || 'Đánh giá' }}</p>
                </div>
                <p class="mt-1 text-xs font-medium text-slate-300 leading-relaxed">{{ review.comment || review.Comment || 'Không có nhận xét.' }}</p>
              </div>
              <p v-if="!isLoadingReviews && reviewsList.length === 0" class="text-xs font-medium text-slate-500 italic">Chưa có đánh giá nào cho ứng viên này.</p>
            </div>
          </section>
        </div>

        <!-- Drawer Footer Actions -->
        <div class="border-t border-white/10 bg-slate-900/90 p-5">
          <div class="grid gap-2.5 sm:grid-cols-2">
            <!-- 1. Open Chat Modal -->
            <button
              @click="openChatModal(selectedManagedApplicant); closeManagedApplicantModal()"
              class="rounded-xl border border-indigo-500/30 bg-indigo-500/10 px-4 py-2.5 text-xs font-extrabold text-indigo-200 hover:bg-indigo-500 hover:text-white transition-all cursor-pointer"
            >
              Nhắn tin với ứng viên
            </button>

            <!-- 2. Open Review & Offer Modal -->
            <button
              @click="openReviewModal(selectedManagedApplicant); closeManagedApplicantModal()"
              class="rounded-xl bg-gradient-to-r from-cyan-400 to-emerald-400 px-4 py-2.5 text-xs font-black text-slate-950 shadow-md hover:brightness-110 transition-all cursor-pointer"
            >
              Đánh giá & Gửi Offer
            </button>

            <!-- 3. Open Escrow Completion Modal -->
            <button
              v-if="['student_completed', 'offer_accepted', 'approved', 'paid'].includes(selectedManagedApplicant.status?.toLowerCase())"
              @click="openCompletionModal(selectedManagedApplicant); closeManagedApplicantModal()"
              class="rounded-xl bg-gradient-to-r from-emerald-400 to-cyan-400 px-4 py-2.5 text-xs font-black text-slate-950 shadow-md hover:brightness-110 transition-all cursor-pointer"
            >
              Xác nhận & Giải ngân Escrow
            </button>

            <!-- 4. Open Business Post-Job Review Modal -->
            <button
              @click="openBusinessReviewModal(selectedManagedApplicant); closeManagedApplicantModal()"
              class="rounded-xl bg-amber-400 px-4 py-2.5 text-xs font-black text-slate-950 shadow-md hover:bg-amber-300 transition-all cursor-pointer"
            >
              Đánh giá sinh viên
            </button>

            <!-- 5. Open Review History Modal -->
            <button
              @click="openReviewsModal(selectedManagedApplicant)"
              class="rounded-xl border border-white/10 bg-white/5 px-4 py-2.5 text-xs font-bold text-slate-300 hover:bg-white/10 hover:text-white transition-all cursor-pointer sm:col-span-2"
            >
              Xem lịch sử đánh giá
            </button>
          </div>
        </div>
      </aside>
    </div>

    <!-- MODAL 4: POST-JOB RATING MODAL -->
    <div v-if="selectedBusinessReviewApp" class="fixed inset-0 z-[9999] flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
      <div class="w-full max-w-lg overflow-hidden rounded-3xl border border-amber-400/30 bg-slate-950 shadow-2xl animate-in fade-in zoom-in-95 duration-150">
        <div class="border-b border-white/10 bg-slate-900/90 p-6">
          <div class="flex items-start justify-between gap-4">
            <div>
              <span class="text-[10px] font-black uppercase tracking-wider text-amber-300 bg-amber-400/10 px-2.5 py-0.5 rounded-full border border-amber-400/20">
                Đánh Giá Sau Công Việc
              </span>
              <h3 class="mt-1 text-xl font-extrabold text-white">Đánh Giá & Nhận Xét Ứng Viên</h3>
              <p class="mt-1 text-xs font-semibold text-slate-400">
                {{ selectedBusinessReviewApp.student?.full_name || 'Ứng viên' }} · {{ selectedBusinessReviewApp.job?.title || jobTitleLookup(selectedBusinessReviewApp.job_id) }}
              </p>
            </div>
            <button
              type="button"
              @click="closeBusinessReviewModal"
              class="rounded-xl bg-white/5 border border-white/10 px-3 py-1.5 text-xs font-bold text-slate-300 hover:bg-white/10 hover:text-white cursor-pointer"
            >
              Đóng
            </button>
          </div>
        </div>

        <form class="space-y-5 p-6" @submit.prevent="submitBusinessReview">
          <div class="rounded-2xl border border-amber-400/30 bg-amber-400/10 p-5">
            <p class="text-xs font-black uppercase tracking-wider text-amber-300">Mức độ hài lòng của Doanh nghiệp</p>
            <div class="mt-3 flex items-center gap-2">
              <button
                v-for="star in 5"
                :key="star"
                type="button"
                @click="businessReviewRating = star"
                class="text-3xl leading-none transition-transform duration-150 hover:scale-110 focus:outline-none cursor-pointer"
                :class="star <= businessReviewRating ? 'text-amber-300' : 'text-slate-700'"
                :aria-label="`Chọn ${star} sao`"
              >
                ★
              </button>
              <span class="ml-2 text-sm font-extrabold text-white">{{ businessReviewRating }}/5 sao</span>
            </div>
          </div>

          <div>
            <label class="mb-1.5 block text-xs font-extrabold uppercase tracking-wider text-cyan-300">
              Nhận xét chi tiết
            </label>
            <textarea
              v-model="businessReviewComment"
              rows="4"
              class="w-full resize-none rounded-2xl border border-white/10 bg-slate-900 px-4 py-3 text-xs text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-amber-300 font-medium leading-relaxed"
              placeholder="Chia sẻ về thái độ làm việc, kỹ năng chuyên môn và trách nhiệm của sinh viên..."
            ></textarea>
          </div>

          <div class="flex flex-col-reverse gap-3 border-t border-white/10 pt-4 sm:flex-row sm:justify-end">
            <button
              type="button"
              @click="closeBusinessReviewModal"
              class="rounded-xl border border-white/10 px-5 py-2.5 text-xs font-bold text-slate-300 hover:bg-white/10 cursor-pointer"
            >
              Hủy bỏ
            </button>
            <button
              type="submit"
              :disabled="isSubmittingBusinessReview"
              class="rounded-xl bg-amber-400 px-6 py-2.5 text-xs font-black text-slate-950 shadow-lg shadow-amber-500/25 hover:bg-amber-300 disabled:opacity-50 cursor-pointer"
            >
              {{ isSubmittingBusinessReview ? 'Đang gửi...' : 'Gửi Đánh Giá' }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- MODAL 5: APPLICATION REVIEWS HISTORY MODAL -->
    <div v-if="isReviewsModalOpen" class="fixed inset-0 z-[10000] flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
      <div class="w-full max-w-xl overflow-hidden rounded-3xl border border-cyan-400/30 bg-slate-950 shadow-2xl animate-in fade-in zoom-in-95 duration-150">
        <div class="flex items-start justify-between gap-4 border-b border-white/10 bg-slate-900/90 p-6">
          <div>
            <span class="text-[10px] font-black uppercase tracking-wider text-cyan-300 bg-cyan-400/10 px-2.5 py-0.5 rounded-full border border-cyan-400/20">
              Lịch Sử Đánh Giá
            </span>
            <h3 class="mt-1 text-xl font-extrabold text-white">
              {{ selectedReviewsApp?.student?.full_name || selectedReviewsApp?.job?.title || 'Đánh giá ứng viên' }}
            </h3>
            <p class="mt-1 text-xs font-semibold text-amber-300">
              {{ reviewSummary.average_rating.toFixed(1) }} / 5 · {{ reviewSummary.total_reviews }} lượt đánh giá
            </p>
          </div>
          <button
            type="button"
            @click="closeReviewsModal"
            class="rounded-xl bg-white/5 border border-white/10 px-3 py-1.5 text-xs font-bold text-slate-300 hover:bg-white/10 hover:text-white cursor-pointer"
          >
            Đóng
          </button>
        </div>

        <div class="max-h-[60vh] space-y-3 overflow-y-auto p-6">
          <div v-if="isLoadingReviews" class="rounded-2xl border border-white/10 bg-slate-900 p-4 text-xs font-medium text-slate-400">
            Đang tải lịch sử đánh giá...
          </div>
          <div
            v-for="review in reviewsList"
            :key="review.id || review.ID"
            class="rounded-2xl border border-white/10 bg-slate-900 p-4.5 space-y-2"
          >
            <div class="flex items-center justify-between gap-3">
              <p class="text-xs font-black text-amber-300">{{ review.rating || review.Rating }} / 5</p>
              <p class="text-[10px] font-bold uppercase text-slate-500">{{ review.review_type || review.ReviewType || 'Đánh giá' }}</p>
            </div>
            <p class="text-xs font-medium text-slate-300 leading-relaxed">{{ review.comment || review.Comment || 'Không có nhận xét.' }}</p>
          </div>
          <p v-if="!isLoadingReviews && reviewsList.length === 0" class="rounded-2xl border border-white/10 bg-slate-900 p-4 text-xs font-medium text-slate-500 italic">
            Chưa có đánh giá nào cho ứng viên này.
          </p>
        </div>
      </div>
    </div>

    <!-- MODAL 6: ENTERPRISE CHAT MODAL -->
    <div 
      v-if="isChatModalOpen && selectedChatApp" 
      class="fixed inset-0 z-[9999] flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md"
    >
      <div class="relative w-full max-w-2xl bg-slate-950 border border-cyan-400/30 rounded-3xl shadow-2xl overflow-hidden z-10 animate-in fade-in zoom-in-95 duration-150">
        
        <!-- Chat Header -->
        <div class="p-5 bg-slate-900/90 border-b border-white/10 flex items-center justify-between">
          <div class="flex items-center gap-3">
            <div class="relative flex-shrink-0">
              <img 
                v-if="selectedChatApp.student?.avatar_url" 
                :src="selectedChatApp.student.avatar_url" 
                class="h-10 w-10 rounded-xl object-cover border border-cyan-400/30"
              />
              <div v-else class="h-10 w-10 rounded-xl bg-gradient-to-br from-indigo-600 to-cyan-500 border border-cyan-400/30 flex items-center justify-center text-white font-black text-sm">
                {{ (selectedChatApp.student?.full_name || 'S').slice(0, 1).toUpperCase() }}
              </div>
              <span class="absolute -bottom-0.5 -right-0.5 h-3 w-3 rounded-full bg-emerald-400 border-2 border-slate-950 animate-ping"></span>
            </div>

            <div>
              <h3 class="text-sm font-extrabold text-white">
                {{ selectedChatApp.student?.full_name || 'Sinh viên' }}
              </h3>
              <p class="text-[11px] font-semibold text-slate-400 mt-0.5">
                Vị trí: <span class="text-cyan-300 font-bold">{{ selectedChatApp.job?.title || 'Công việc' }}</span> · Mã đơn #{{ selectedChatApp.id }}
              </p>
            </div>
          </div>

          <button 
            @click="isChatModalOpen = false" 
            class="text-slate-400 hover:text-white text-xs font-bold px-3 py-1.5 border border-white/10 bg-white/5 hover:bg-white/10 rounded-xl transition-all cursor-pointer"
          >
            Đóng
          </button>
        </div>

        <!-- Chat Container -->
        <div class="p-4 bg-slate-950">
          <ChatBox 
            v-if="currentBusinessUserId"
            :applicationId="selectedChatApp.id"
            :currentUserId="currentBusinessUserId"
            :targetId="selectedChatApp.student?.user_id || selectedChatApp.student_id"
          />
        </div>

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
  isEditing,
  profileForm,
  logoPreview,
  onLogoFileChange,
  isSavingProfile,
  handleUpdateProfile,
  showCreateForm,
  jobForm,
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
  openCompletionModal,
  openBusinessReviewModal,
  openManagedApplicantModal,
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
  selectedBusinessReviewApp,
  businessReviewRating,
  businessReviewComment,
  isSubmittingBusinessReview,
  closeBusinessReviewModal,
  submitBusinessReview,
  selectedManagedApplicant,
  closeManagedApplicantModal,
  selectedReviewsApp,
  reviewsList,
  reviewSummary,
  isLoadingReviews,
  isReviewsModalOpen,
  openReviewsModal,
  closeReviewsModal,
  closeModal,
  submitReview,
  isChatModalOpen,
  selectedChatApp,
  currentBusinessUserId
} = toRefs(props.state)
</script>
