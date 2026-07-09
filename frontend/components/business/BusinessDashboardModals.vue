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

<div v-if="selectedManagedApplicant" class="fixed inset-0 z-[9998] flex justify-end bg-slate-950/80 backdrop-blur-md">
  <button
    type="button"
    class="absolute inset-0 cursor-default"
    aria-label="Close applicant details"
    @click="closeManagedApplicantModal"
  ></button>

  <aside class="relative flex h-full w-full max-w-2xl flex-col overflow-hidden border-l border-white/10 bg-slate-950 shadow-2xl shadow-slate-950/70">
    <div class="border-b border-white/10 bg-slate-900/90 p-5">
      <div class="flex items-start justify-between gap-4">
        <div class="min-w-0">
          <p class="text-xs font-bold uppercase tracking-wider text-cyan-300">Applicant Management</p>
          <h3 class="mt-1 truncate text-2xl font-extrabold text-white">
            {{ selectedManagedApplicant.student?.full_name || 'Ứng viên' }}
          </h3>
          <p class="mt-1 text-sm font-medium text-slate-400">
            Application #{{ selectedManagedApplicant.id }} · {{ selectedManagedApplicant.job?.title || jobTitleLookup(selectedManagedApplicant.job_id) }}
          </p>
        </div>
        <button
          type="button"
          @click="closeManagedApplicantModal"
          class="rounded-lg bg-white/5 px-3 py-1.5 text-sm font-bold text-slate-300 hover:bg-white/10 hover:text-white"
        >
          Đóng
        </button>
      </div>
    </div>

    <div class="flex-1 space-y-5 overflow-y-auto p-5">
      <section class="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
        <div class="flex items-start gap-4">
          <img
            v-if="selectedManagedApplicant.student?.avatar_url"
            :src="selectedManagedApplicant.student.avatar_url"
            class="h-16 w-16 rounded-2xl border border-white/10 object-cover"
          />
          <div
            v-else
            class="flex h-16 w-16 flex-shrink-0 items-center justify-center rounded-2xl border border-white/10 bg-slate-800 text-xl font-extrabold text-cyan-200"
          >
            {{ selectedManagedApplicant.student?.full_name?.charAt(0) || 'S' }}
          </div>

          <div class="min-w-0 flex-1">
            <div class="flex flex-wrap items-center gap-2">
              <h4 class="text-lg font-extrabold text-white">{{ selectedManagedApplicant.student?.full_name || 'N/A' }}</h4>
              <span :class="[
                statusBadgeClass(selectedManagedApplicant.status),
                'inline-flex items-center rounded-full border px-2.5 py-1 text-xs font-bold capitalize'
              ]">
                {{ selectedManagedApplicant.status ? selectedManagedApplicant.status.replace('_', ' ') : 'Pending' }}
              </span>
            </div>
            <p class="mt-1 text-sm font-medium text-slate-400">
              {{ selectedManagedApplicant.student?.phone || 'Chưa có số điện thoại' }}
            </p>
            <a
              v-if="selectedManagedApplicant.student?.cv_url"
              :href="selectedManagedApplicant.student.cv_url"
              target="_blank"
              class="mt-3 inline-flex items-center rounded-lg border border-rose-400/20 bg-rose-500/10 px-3 py-1.5 text-xs font-bold text-rose-200 hover:bg-rose-500/20 hover:text-white"
            >
              CV Profile
            </a>
          </div>
        </div>

        <div class="mt-4 flex flex-wrap gap-2">
          <span
            v-for="(skill, index) in parseSkills(selectedManagedApplicant.student?.skills)"
            :key="index"
            class="rounded-md border border-white/10 bg-slate-900 px-2 py-1 text-xs font-bold text-slate-300"
          >
            {{ skill }}
          </span>
          <span
            v-if="!selectedManagedApplicant.student?.skills || parseSkills(selectedManagedApplicant.student?.skills).length === 0"
            class="text-xs font-semibold text-slate-500"
          >
            No skills specified
          </span>
        </div>
      </section>

      <section class="grid gap-4 sm:grid-cols-2">
        <div class="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
          <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Job Information</p>
          <h4 class="mt-2 text-base font-extrabold text-white">
            {{ selectedManagedApplicant.job?.title || jobTitleLookup(selectedManagedApplicant.job_id) }}
          </h4>
          <div class="mt-3 space-y-2 text-sm font-medium text-slate-400">
            <p>Location: <span class="font-bold text-slate-200">{{ selectedManagedApplicant.job?.location || 'N/A' }}</span></p>
            <p>Salary: <span class="font-bold text-emerald-300">{{ selectedManagedApplicant.job?.salary ? Number(selectedManagedApplicant.job.salary).toLocaleString() : 'Thỏa thuận' }}</span></p>
            <p>Applied: <span class="font-bold text-slate-200">{{ formatDate(selectedManagedApplicant.applied_at || selectedManagedApplicant.id) }}</span></p>
          </div>
        </div>

        <div class="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
          <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Payment Status</p>
          <div
            v-if="selectedManagedApplicant.status?.toLowerCase() === 'paid'"
            class="mt-3 rounded-xl border border-cyan-400/20 bg-cyan-400/10 p-3"
          >
            <p class="text-sm font-extrabold text-cyan-200">Paid / Released</p>
            <p class="mt-1 text-xs font-medium text-slate-400">Payment has been released for this application.</p>
          </div>
          <div
            v-else-if="selectedManagedApplicant.status?.toLowerCase() === 'student_completed'"
            class="mt-3 rounded-xl border border-amber-400/20 bg-amber-400/10 p-3"
          >
            <p class="text-sm font-extrabold text-amber-200">Waiting Confirmation</p>
            <p class="mt-1 text-xs font-medium text-slate-400">Student marked the work complete. Employer confirmation is required.</p>
          </div>
          <div v-else class="mt-3 rounded-xl border border-white/10 bg-slate-900 p-3">
            <p class="text-sm font-extrabold text-slate-300">Not Released</p>
            <p class="mt-1 text-xs font-medium text-slate-500">Payment actions appear after work completion.</p>
          </div>
        </div>
      </section>

      <section class="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
        <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Offer Information</p>
        <div
          v-if="selectedManagedApplicant.offer_salary || selectedManagedApplicant.offer_start_date || selectedManagedApplicant.offer_message"
          class="mt-3 grid gap-3 sm:grid-cols-3"
        >
          <div class="rounded-xl border border-emerald-400/20 bg-emerald-400/10 p-3">
            <p class="text-xs font-bold uppercase tracking-wider text-emerald-200">Salary</p>
            <p class="mt-1 text-sm font-extrabold text-white">{{ selectedManagedApplicant.offer_salary || 'Thỏa thuận' }}</p>
          </div>
          <div class="rounded-xl border border-cyan-400/20 bg-cyan-400/10 p-3">
            <p class="text-xs font-bold uppercase tracking-wider text-cyan-200">Start Date</p>
            <p class="mt-1 text-sm font-extrabold text-white">{{ selectedManagedApplicant.offer_start_date || 'N/A' }}</p>
          </div>
          <div class="rounded-xl border border-white/10 bg-slate-900 p-3 sm:col-span-3">
            <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Message</p>
            <p class="mt-1 whitespace-pre-line text-sm font-medium text-slate-300">{{ selectedManagedApplicant.offer_message || 'Không có lời nhắn offer.' }}</p>
          </div>
        </div>
        <p v-else class="mt-3 text-sm font-medium text-slate-500">No offer information recorded yet.</p>
      </section>

      <section class="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
        <div class="flex flex-wrap items-center justify-between gap-3">
          <div>
            <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Reviews</p>
            <p class="mt-1 text-sm font-extrabold text-white">
              {{ reviewSummary.average_rating.toFixed(1) }} / 5 · {{ reviewSummary.total_reviews }} reviews
            </p>
          </div>
          <button
            @click="openReviewsModal(selectedManagedApplicant)"
            class="rounded-lg border border-white/10 bg-white/5 px-3 py-1.5 text-xs font-bold text-slate-200 hover:bg-white/10 hover:text-white"
          >
            View Reviews
          </button>
        </div>
        <div class="mt-3 space-y-2">
          <div v-if="isLoadingReviews" class="text-sm font-medium text-slate-500">Loading reviews...</div>
          <div
            v-for="review in reviewsList.slice(0, 3)"
            :key="review.id || review.ID"
            class="rounded-xl border border-white/10 bg-slate-900 p-3"
          >
            <div class="flex items-center justify-between gap-2">
              <p class="text-sm font-extrabold text-amber-300">{{ review.rating || review.Rating }} / 5</p>
              <p class="text-xs font-bold uppercase tracking-wider text-slate-500">{{ review.review_type || review.ReviewType || 'review' }}</p>
            </div>
            <p class="mt-1 text-sm font-medium text-slate-300">{{ review.comment || review.Comment || 'No comment.' }}</p>
          </div>
          <p v-if="!isLoadingReviews && reviewsList.length === 0" class="text-sm font-medium text-slate-500">No reviews for this application yet.</p>
        </div>
      </section>
    </div>

    <div class="border-t border-white/10 bg-slate-900/90 p-5">
      <div class="grid gap-2 sm:grid-cols-2">
        <button
          @click="openChatModal(selectedManagedApplicant); closeManagedApplicantModal()"
          class="rounded-lg border border-white/10 bg-white/5 px-4 py-2.5 text-sm font-bold text-slate-200 hover:bg-white/10 hover:text-white"
        >
          Chat with candidate
        </button>

        <button
          v-if="['pending', 'applied'].includes(selectedManagedApplicant.status?.toLowerCase())"
          @click="openReviewModal(selectedManagedApplicant); closeManagedApplicantModal()"
          class="rounded-lg bg-cyan-400 px-4 py-2.5 text-sm font-extrabold text-slate-950 shadow-lg shadow-cyan-500/20 hover:bg-cyan-300"
        >
          Review & Offer
        </button>

        <button
          v-if="!['pending', 'applied'].includes(selectedManagedApplicant.status?.toLowerCase())"
          @click="openReviewModal(selectedManagedApplicant); closeManagedApplicantModal()"
          class="rounded-lg border border-cyan-400/30 bg-cyan-400/10 px-4 py-2.5 text-sm font-bold text-cyan-200 hover:bg-cyan-400/15"
        >
          View Details
        </button>

        <button
          v-if="selectedManagedApplicant.status?.toLowerCase() === 'student_completed'"
          @click="openCompletionModal(selectedManagedApplicant); closeManagedApplicantModal()"
          class="rounded-lg bg-emerald-400 px-4 py-2.5 text-sm font-extrabold text-slate-950 shadow-lg shadow-emerald-500/20 hover:bg-emerald-300"
        >
          Confirm Complete
        </button>

        <div
          v-if="selectedManagedApplicant.status?.toLowerCase() === 'paid'"
          class="rounded-lg border border-cyan-400/20 bg-cyan-400/10 px-4 py-2.5 text-center text-sm font-extrabold text-cyan-200"
        >
          Paid / Released
        </div>

        <button
          v-if="selectedManagedApplicant.status?.toLowerCase() === 'paid' && !reviewsList.some(review => Number(review.reviewer_id || review.ReviewerID) === Number(currentBusinessUserId))"
          @click="openBusinessReviewModal(selectedManagedApplicant); closeManagedApplicantModal()"
          class="rounded-lg bg-amber-400 px-4 py-2.5 text-sm font-extrabold text-slate-950 shadow-lg shadow-amber-500/20 hover:bg-amber-300"
        >
          Review Student
        </button>

        <button
          @click="openReviewsModal(selectedManagedApplicant)"
          class="rounded-lg border border-white/10 bg-white/5 px-4 py-2.5 text-sm font-bold text-slate-200 hover:bg-white/10 hover:text-white"
        >
          View Reviews
        </button>
      </div>
    </div>
  </aside>
</div>

<div v-if="selectedBusinessReviewApp" class="fixed inset-0 z-[9999] flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
  <div class="w-full max-w-lg overflow-hidden rounded-2xl border border-white/10 bg-slate-950 shadow-2xl shadow-slate-950/70">
    <div class="border-b border-white/10 bg-slate-900/90 p-5">
      <div class="flex items-start justify-between gap-4">
        <div>
          <p class="text-xs font-bold uppercase tracking-wider text-amber-300">Đánh giá sau công việc</p>
          <h3 class="mt-1 text-xl font-extrabold text-white">Đánh giá ứng viên</h3>
          <p class="mt-1 text-sm font-medium text-slate-400">
            {{ selectedBusinessReviewApp.student?.full_name || 'Ứng viên' }} · {{ selectedBusinessReviewApp.job?.title || jobTitleLookup(selectedBusinessReviewApp.job_id) }}
          </p>
        </div>
        <button
          type="button"
          @click="closeBusinessReviewModal"
          class="rounded-lg bg-white/5 px-3 py-1.5 text-sm font-bold text-slate-300 hover:bg-white/10 hover:text-white"
        >
          Đóng
        </button>
      </div>
    </div>

    <form class="space-y-5 p-5" @submit.prevent="submitBusinessReview">
      <div class="rounded-xl border border-amber-400/20 bg-amber-400/10 p-4">
        <p class="text-xs font-bold uppercase tracking-wider text-amber-200">Mức độ hài lòng</p>
        <div class="mt-3 flex items-center gap-2">
          <button
            v-for="star in 5"
            :key="star"
            type="button"
            @click="businessReviewRating = star"
            class="text-3xl leading-none transition-transform duration-150 hover:scale-110 focus:outline-none"
            :class="star <= businessReviewRating ? 'text-amber-300' : 'text-slate-700'"
            :aria-label="`Chọn ${star} sao`"
          >
            ★
          </button>
          <span class="ml-2 text-sm font-bold text-slate-300">{{ businessReviewRating }}/5</span>
        </div>
      </div>

      <div>
        <label class="mb-2 block text-xs font-bold uppercase tracking-wider text-slate-400">
          Nhận xét
        </label>
        <textarea
          v-model="businessReviewComment"
          rows="4"
          class="w-full resize-none rounded-xl border border-white/10 bg-slate-900 px-4 py-3 text-sm text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-amber-300"
          placeholder="Chia sẻ về thái độ, kỹ năng và mức độ hoàn thành công việc của ứng viên..."
        ></textarea>
      </div>

      <div class="flex flex-col-reverse gap-2 border-t border-white/10 pt-4 sm:flex-row sm:justify-end">
        <button
          type="button"
          @click="closeBusinessReviewModal"
          class="rounded-lg border border-white/10 px-4 py-2.5 text-sm font-bold text-slate-300 hover:bg-white/10 hover:text-white"
        >
          Hủy
        </button>
        <button
          type="submit"
          :disabled="isSubmittingBusinessReview"
          class="rounded-lg bg-amber-400 px-4 py-2.5 text-sm font-extrabold text-slate-950 shadow-lg shadow-amber-500/20 hover:bg-amber-300 disabled:cursor-not-allowed disabled:opacity-60"
        >
          {{ isSubmittingBusinessReview ? 'Đang gửi...' : 'Gửi đánh giá' }}
        </button>
      </div>
    </form>
  </div>
</div>

<div v-if="isReviewsModalOpen" class="fixed inset-0 z-[10000] flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
  <div class="w-full max-w-xl overflow-hidden rounded-2xl border border-white/10 bg-slate-950 shadow-2xl shadow-slate-950/70">
    <div class="flex items-start justify-between gap-4 border-b border-white/10 bg-slate-900/90 p-5">
      <div>
        <p class="text-xs font-bold uppercase tracking-wider text-cyan-300">Reviews</p>
        <h3 class="mt-1 text-xl font-extrabold text-white">
          {{ selectedReviewsApp?.student?.full_name || selectedReviewsApp?.job?.title || 'Application reviews' }}
        </h3>
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
        Loading reviews...
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
        <p class="mt-2 text-sm font-medium text-slate-300">{{ review.comment || review.Comment || 'No comment.' }}</p>
      </div>
      <p v-if="!isLoadingReviews && reviewsList.length === 0" class="rounded-xl border border-white/10 bg-white/[0.04] p-4 text-sm font-medium text-slate-500">
        No reviews for this application yet.
      </p>
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
  openCompletionModal,
  closeCompletionModal,
  submitBusinessCompletion,
  selectedBusinessReviewApp,
  businessReviewRating,
  businessReviewComment,
  isSubmittingBusinessReview,
  openBusinessReviewModal,
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
