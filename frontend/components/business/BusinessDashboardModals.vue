<template>
  <div v-if="selectedApp" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
  <div class="bg-slate-900 rounded-3xl border border-cyan-500/20 shadow-2xl shadow-slate-950 max-w-lg w-full overflow-hidden flex flex-col max-h-[90vh] text-left animate-in fade-in zoom-in-95 duration-150">
    
    <div class="p-6 border-b border-cyan-500/15 flex justify-between items-center bg-slate-950/80">
      <div>
        <h3 class="text-base font-extrabold text-white">Đánh giá hồ sơ & Gửi Offer</h3>
        <p class="text-xs text-slate-400 font-semibold mt-0.5">Ứng viên: <span class="text-cyan-300 font-extrabold">{{ selectedApp.student?.full_name || 'N/A' }}</span></p>
      </div>
      <button @click="closeModal" class="text-slate-400 hover:text-white font-bold text-xl p-1.5 rounded-xl hover:bg-white/10 transition-colors">&times;</button>
    </div>

    <div class="p-6 space-y-4 overflow-y-auto flex-1 custom-scrollbar">
      <div>
        <label class="text-xs font-bold text-slate-400 uppercase tracking-wider block mb-1.5">Lời nhắn từ ứng viên (Cover Note)</label>
        <div class="p-3.5 bg-cyan-500/10 border border-cyan-500/20 rounded-2xl text-xs font-medium text-slate-200 whitespace-pre-line">
          {{ selectedApp.cover_note || 'Ứng viên không để lại lời nhắn.' }}
        </div>
      </div>

      <div>
        <label class="text-xs font-bold text-slate-400 uppercase tracking-wider block mb-2">Quyết định</label>
        <div class="grid grid-cols-2 gap-3">
          <button 
            type="button"
            @click="reviewStatus = 'approved'"
            :class="[reviewStatus === 'approved' ? 'border-emerald-400 bg-emerald-500/20 text-emerald-300 ring-2 ring-emerald-400/30' : 'border-white/10 bg-slate-950 text-slate-300 hover:bg-white/5']"
            class="py-2.5 px-4 text-xs font-extrabold border rounded-2xl transition-all flex items-center justify-center space-x-1 cursor-pointer"
          >
            <span>👍 Chấp nhận & Gửi Offer</span>
          </button>
          <button 
            type="button"
            @click="reviewStatus = 'rejected'"
            :class="[reviewStatus === 'rejected' ? 'border-rose-400 bg-rose-500/20 text-rose-300 ring-2 ring-rose-400/30' : 'border-white/10 bg-slate-950 text-slate-300 hover:bg-white/5']"
            class="py-2.5 px-4 text-xs font-extrabold border rounded-2xl transition-all flex items-center justify-center space-x-1 cursor-pointer"
          >
            <span>👎 Từ chối đơn</span>
          </button>
        </div>
      </div>

      <div v-if="reviewStatus === 'approved'" class="p-4 bg-emerald-500/10 border border-emerald-400/20 rounded-2xl space-y-3 animate-in fade-in duration-150">
        <h4 class="text-xs font-extrabold text-emerald-300">✉️ Chi tiết Offer gửi Sinh viên</h4>
        
        <div class="grid grid-cols-2 gap-3">
          <div>
            <label class="text-[11px] font-bold text-slate-300 block mb-1">Mức lương Offer</label>
            <input 
              v-model="offerForm.salary"
              type="text" 
              placeholder="Ví dụ: 15,000,000 VND"
              class="w-full text-xs px-3 py-2 border border-white/15 rounded-xl bg-slate-950 text-white focus:outline-none focus:border-emerald-400"
            />
          </div>
          <div>
            <label class="text-[11px] font-bold text-slate-300 block mb-1">Ngày đi làm dự kiến</label>
            <input 
              v-model="offerForm.startDate"
              type="date" 
              class="w-full text-xs px-3 py-2 border border-white/15 rounded-xl bg-slate-950 text-white focus:outline-none focus:border-emerald-400"
            />
          </div>
        </div>

        <div>
          <label class="text-[11px] font-bold text-slate-300 block mb-1">Lời nhắn chào mừng từ HR</label>
          <textarea 
            v-model="offerForm.message"
            rows="2"
            placeholder="Chào mừng bạn gia nhập đội ngũ..."
            class="w-full text-xs p-2.5 border border-white/15 rounded-xl bg-slate-950 text-white focus:outline-none focus:border-emerald-400"
          ></textarea>
        </div>
      </div>
    </div>

    <div class="p-4 bg-slate-950/80 border-t border-cyan-500/15 flex items-center justify-end space-x-2">
      <button 
        @click="closeModal" 
        class="px-4 py-2 text-xs font-bold text-slate-300 hover:bg-white/10 rounded-xl transition-all"
      >
        Đóng
      </button>
      <button 
        @click="submitReview"
        :disabled="!reviewStatus || isSubmitting"
        class="px-5 py-2 text-xs font-extrabold text-slate-950 bg-emerald-400 hover:bg-emerald-300 disabled:opacity-50 rounded-xl shadow-lg shadow-emerald-500/20 transition-all cursor-pointer"
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
            <p class="text-xs font-bold uppercase tracking-wider text-cyan-200">Mức lương Offer</p>
            <p class="mt-1 text-sm font-extrabold text-white">{{ selectedCompletionApp.offer_salary || 'Thỏa thuận' }}</p>
          </div>
          <div>
            <p class="text-xs font-bold uppercase tracking-wider text-cyan-200">Ngày bắt đầu</p>
            <p class="mt-1 text-sm font-extrabold text-white">{{ selectedCompletionApp.offer_start_date || 'N/A' }}</p>
          </div>
          <div>
            <p class="text-xs font-bold uppercase tracking-wider text-cyan-200">Thanh toán Escrow</p>
            <p class="mt-1 text-sm font-extrabold text-white">Giả lập giải ngân</p>
          </div>
        </div>
      </div>

      <!-- Section: Work Report & Proof Submission from Student -->
      <div class="rounded-xl border border-cyan-400/30 bg-slate-900 p-4 space-y-3">
        <div class="flex items-center justify-between">
          <span class="text-[10px] font-black uppercase tracking-wider text-cyan-300 bg-cyan-400/10 px-2.5 py-0.5 rounded-full border border-cyan-400/20">
            Báo Cáo Hoàn Thành Từ Sinh Viên
          </span>
          <span v-if="selectedCompletionApp.submitted_at" class="text-[11px] font-semibold text-slate-400">
            Nộp lúc: {{ formatDate(selectedCompletionApp.submitted_at) }}
          </span>
        </div>

        <div>
          <p class="text-[10px] font-bold uppercase text-slate-400 mb-1">Nội dung báo cáo công việc</p>
          <div class="p-3 bg-slate-950 border border-white/10 rounded-xl text-xs font-medium text-slate-200 whitespace-pre-line leading-relaxed">
            {{ selectedCompletionApp.completion_note || 'Sinh viên không để lại ghi chú báo cáo.' }}
          </div>
        </div>

        <div v-if="selectedCompletionApp.completion_proof_url">
          <p class="text-[10px] font-bold uppercase text-slate-400 mb-1">Bằng chứng / File sản phẩm đính kèm</p>
          <a
            :href="getMediaUrl(selectedCompletionApp.completion_proof_url)"
            target="_blank"
            rel="noopener noreferrer"
            class="inline-flex items-center gap-1.5 rounded-xl border border-cyan-400/30 bg-cyan-400/10 px-3.5 py-2 text-xs font-extrabold text-cyan-300 hover:bg-cyan-400/20 transition-all shadow-sm"
          >
            <span>📄</span>
            <span>Xem Sản Phẩm / Bằng Chứng Hoàn Thành</span>
            <span>↗</span>
          </a>
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
    aria-label="Đóng chi tiết ứng viên"
    @click="closeManagedApplicantModal"
  ></button>

  <aside class="relative flex h-full w-full max-w-2xl flex-col overflow-hidden border-l border-white/10 bg-slate-950 shadow-2xl shadow-slate-950/70">
    <div class="border-b border-white/10 bg-slate-900/90 p-5">
      <div class="flex items-start justify-between gap-4">
        <div class="min-w-0">
          <p class="text-xs font-bold uppercase tracking-wider text-cyan-300">Quản lý Ứng viên</p>
          <h3 class="mt-1 truncate text-2xl font-extrabold text-white">
            {{ selectedManagedApplicant.student?.full_name || 'Ứng viên' }}
          </h3>
          <p class="mt-1 text-sm font-medium text-slate-400">
            Mã đơn #{{ selectedManagedApplicant.id }} · {{ selectedManagedApplicant.job?.title || jobTitleLookup(selectedManagedApplicant.job_id) }}
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
                {{ selectedManagedApplicant.status ? selectedManagedApplicant.status.replace('_', ' ') : 'Chờ duyệt' }}
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
              Hồ sơ CV
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
            Chưa cập nhật kỹ năng
          </span>
        </div>
      </section>

      <section class="grid gap-4 sm:grid-cols-2">
        <div class="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
          <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Thông tin Công việc</p>
          <h4 class="mt-2 text-base font-extrabold text-white">
            {{ selectedManagedApplicant.job?.title || jobTitleLookup(selectedManagedApplicant.job_id) }}
          </h4>
          <div class="mt-3 space-y-2 text-sm font-medium text-slate-400">
            <p>Địa điểm: <span class="font-bold text-slate-200">{{ selectedManagedApplicant.job?.location || 'N/A' }}</span></p>
            <p>Mức lương: <span class="font-bold text-emerald-300">{{ selectedManagedApplicant.job?.salary ? Number(selectedManagedApplicant.job.salary).toLocaleString() : 'Thỏa thuận' }}</span></p>
            <p>Ngày nộp đơn: <span class="font-bold text-slate-200">{{ formatDate(selectedManagedApplicant.applied_at || selectedManagedApplicant.id) }}</span></p>
          </div>
        </div>

        <div class="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
          <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Trạng thái Thanh toán</p>
          <div
            v-if="selectedManagedApplicant.status?.toLowerCase() === 'paid'"
            class="mt-3 rounded-xl border border-cyan-400/20 bg-cyan-400/10 p-3"
          >
            <p class="text-sm font-extrabold text-cyan-200">Đã thanh toán Escrow</p>
            <p class="mt-1 text-xs font-medium text-slate-400">Lương ca làm đã được giải ngân cho ứng viên này.</p>
          </div>
          <div
            v-else-if="selectedManagedApplicant.status?.toLowerCase() === 'student_completed'"
            class="mt-3 rounded-xl border border-amber-400/20 bg-amber-400/10 p-3"
          >
            <p class="text-sm font-extrabold text-amber-200">Chờ Doanh nghiệp duyệt</p>
            <p class="mt-1 text-xs font-medium text-slate-400">Sinh viên đã báo hoàn thành công việc. Cần xác nhận của doanh nghiệp.</p>
          </div>
          <div v-else class="mt-3 rounded-xl border border-white/10 bg-slate-900 p-3">
            <p class="text-sm font-extrabold text-slate-300">Chưa giải ngân</p>
            <p class="mt-1 text-xs font-medium text-slate-500">Nút thanh toán sẽ hiển thị khi ca làm được sinh viên báo hoàn tất.</p>
          </div>
        </div>
      </section>

      <section class="rounded-2xl border border-rose-500/20 bg-rose-500/5 p-4">
        <div class="flex items-center justify-between gap-3">
          <div>
            <p class="text-xs font-extrabold uppercase tracking-wider text-rose-300">Hỗ trợ & Khiếu nại (Dispute)</p>
            <p v-if="state.myBusinessTickets?.find((t: any) => Number(t.application_id) === Number(selectedManagedApplicant.id))" class="mt-1 text-xs font-semibold text-emerald-300">
              Đã có khiếu nại (Trạng thái: {{ state.myBusinessTickets.find((t: any) => Number(t.application_id) === Number(selectedManagedApplicant.id)).status === 'resolved' ? 'Đã duyệt' : state.myBusinessTickets.find((t: any) => Number(t.application_id) === Number(selectedManagedApplicant.id)).status === 'rejected' ? 'Bị bác bỏ' : 'Đang xử lý' }})
            </p>
            <p v-else class="mt-1 text-xs font-medium text-slate-400">Ứng viên vi phạm quy định? Gửi khiếu nại tới Ban quản trị.</p>
          </div>

          <button
            v-if="state.myBusinessTickets?.find((t: any) => Number(t.application_id) === Number(selectedManagedApplicant.id))"
            @click="state.openBusinessReappealModal(state.myBusinessTickets.find((t: any) => Number(t.application_id) === Number(selectedManagedApplicant.id)))"
            class="rounded-xl border border-emerald-400/30 bg-emerald-500/10 px-3.5 py-2 text-xs font-extrabold text-emerald-300 hover:bg-emerald-500/20 hover:text-white transition-all cursor-pointer flex-shrink-0"
          >
            ⚖️ Xem Phán Quyết / Tái Khiếu Nại
          </button>

          <button
            v-else
            @click="openBusinessDisputeModal(selectedManagedApplicant)"
            class="rounded-xl border border-rose-400/30 bg-rose-500/10 px-3.5 py-2 text-xs font-extrabold text-rose-300 hover:bg-rose-500/20 hover:text-white transition-all cursor-pointer flex-shrink-0"
          >
            ⚠️ Khiếu Nại Ứng Viên
          </button>
        </div>
      </section>

      <section class="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
        <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Thông tin Offer</p>
        <div
          v-if="selectedManagedApplicant.offer_salary || selectedManagedApplicant.offer_start_date || selectedManagedApplicant.offer_message"
          class="mt-3 grid gap-3 sm:grid-cols-3"
        >
          <div class="rounded-xl border border-emerald-400/20 bg-emerald-400/10 p-3">
            <p class="text-xs font-bold uppercase tracking-wider text-emerald-200">Mức lương Offer</p>
            <p class="mt-1 text-sm font-extrabold text-white">{{ selectedManagedApplicant.offer_salary || 'Thỏa thuận' }}</p>
          </div>
          <div class="rounded-xl border border-cyan-400/20 bg-cyan-400/10 p-3">
            <p class="text-xs font-bold uppercase tracking-wider text-cyan-200">Ngày bắt đầu</p>
            <p class="mt-1 text-sm font-extrabold text-white">{{ selectedManagedApplicant.offer_start_date || 'N/A' }}</p>
          </div>
          <div class="rounded-xl border border-white/10 bg-slate-900 p-3 sm:col-span-3">
            <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Lời nhắn</p>
            <p class="mt-1 whitespace-pre-line text-sm font-medium text-slate-300">{{ selectedManagedApplicant.offer_message || 'Không có lời nhắn offer.' }}</p>
          </div>
        </div>
        <p v-else class="mt-3 text-sm font-medium text-slate-500">Chưa có thông tin offer nào được ghi nhận.</p>
      </section>

      <!-- Section: Work Report & Proof Submission from Student -->
      <section v-if="selectedManagedApplicant.completion_note || selectedManagedApplicant.completion_proof_url" class="rounded-2xl border border-cyan-400/30 bg-slate-900 p-4 space-y-3">
        <div class="flex items-center justify-between">
          <span class="text-[10px] font-black uppercase tracking-wider text-cyan-300 bg-cyan-400/10 px-2.5 py-0.5 rounded-full border border-cyan-400/20">
            Báo Cáo Hoàn Thành Từ Sinh Viên
          </span>
          <span v-if="selectedManagedApplicant.submitted_at" class="text-[11px] font-semibold text-slate-400">
            Nộp lúc: {{ formatDate(selectedManagedApplicant.submitted_at) }}
          </span>
        </div>

        <div>
          <p class="text-[10px] font-bold uppercase text-slate-400 mb-1">Nội dung báo cáo công việc</p>
          <div class="p-3 bg-slate-950 border border-white/10 rounded-xl text-xs font-medium text-slate-200 whitespace-pre-line leading-relaxed">
            {{ selectedManagedApplicant.completion_note || 'Sinh viên không để lại ghi chú báo cáo.' }}
          </div>
        </div>

        <div v-if="selectedManagedApplicant.completion_proof_url">
          <p class="text-[10px] font-bold uppercase text-slate-400 mb-1">Bằng chứng / File sản phẩm đính kèm</p>
          <a
            :href="getMediaUrl(selectedManagedApplicant.completion_proof_url)"
            target="_blank"
            rel="noopener noreferrer"
            class="inline-flex items-center gap-1.5 rounded-xl border border-cyan-400/30 bg-cyan-400/10 px-3.5 py-2 text-xs font-extrabold text-cyan-300 hover:bg-cyan-400/20 transition-all shadow-sm"
          >
            <span>📄</span>
            <span>Xem Sản Phẩm / Bằng Chứng Hoàn Thành</span>
            <span>↗</span>
          </a>
        </div>
      </section>

      <section class="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
        <div class="flex flex-wrap items-center justify-between gap-3">
          <div>
            <p class="text-xs font-bold uppercase tracking-wider text-slate-500">Đánh giá & Nhận xét</p>
            <p class="mt-1 text-sm font-extrabold text-white">
              {{ reviewSummary.average_rating.toFixed(1) }} / 5 · {{ reviewSummary.total_reviews }} lượt đánh giá
            </p>
          </div>
          <button
            @click="openReviewsModal(selectedManagedApplicant)"
            class="rounded-lg border border-white/10 bg-white/5 px-3 py-1.5 text-xs font-bold text-slate-200 hover:bg-white/10 hover:text-white"
          >
            Xem đánh giá
          </button>
        </div>
        <div class="mt-3 space-y-2">
          <div v-if="isLoadingReviews" class="text-sm font-medium text-slate-500">Đang tải đánh giá...</div>
          <div
            v-for="review in reviewsList.slice(0, 3)"
            :key="review.id || review.ID"
            class="rounded-xl border border-white/10 bg-slate-900 p-3"
          >
            <div class="flex items-center justify-between gap-2">
              <p class="text-sm font-extrabold text-amber-300">{{ review.rating || review.Rating }} / 5</p>
              <p class="text-xs font-bold uppercase tracking-wider text-slate-500">{{ review.review_type || review.ReviewType || 'đánh giá' }}</p>
            </div>
            <p class="mt-1 text-sm font-medium text-slate-300">{{ review.comment || review.Comment || 'Không có nhận xét.' }}</p>
          </div>
          <p v-if="!isLoadingReviews && reviewsList.length === 0" class="text-sm font-medium text-slate-500">Chưa có đánh giá nào cho ứng viên này.</p>
        </div>
      </section>
    </div>

    <div class="border-t border-white/10 bg-slate-900/90 p-5">
      <div class="grid gap-2 sm:grid-cols-2">
        <button
          @click="openChatModal(selectedManagedApplicant); closeManagedApplicantModal()"
          class="rounded-lg border border-white/10 bg-white/5 px-4 py-2.5 text-sm font-bold text-slate-200 hover:bg-white/10 hover:text-white"
        >
          Trò chuyện cùng ứng viên
        </button>

        <button
          v-if="['pending', 'applied'].includes(selectedManagedApplicant.status?.toLowerCase())"
          @click="openReviewModal(selectedManagedApplicant); closeManagedApplicantModal()"
          class="rounded-lg bg-cyan-400 px-4 py-2.5 text-sm font-extrabold text-slate-950 shadow-lg shadow-cyan-500/20 hover:bg-cyan-300"
        >
          Đánh giá & Gửi Offer
        </button>

        <button
          v-if="!['pending', 'applied'].includes(selectedManagedApplicant.status?.toLowerCase())"
          @click="openReviewModal(selectedManagedApplicant); closeManagedApplicantModal()"
          class="rounded-lg border border-cyan-400/30 bg-cyan-400/10 px-4 py-2.5 text-sm font-bold text-cyan-200 hover:bg-cyan-400/15"
        >
          Xem chi tiết
        </button>

        <button
          v-if="selectedManagedApplicant.status?.toLowerCase() === 'student_completed'"
          @click="openCompletionModal(selectedManagedApplicant); closeManagedApplicantModal()"
          class="rounded-lg bg-emerald-400 px-4 py-2.5 text-sm font-extrabold text-slate-950 shadow-lg shadow-emerald-500/20 hover:bg-emerald-300"
        >
          Xác nhận hoàn thành
        </button>

        <div
          v-if="selectedManagedApplicant.status?.toLowerCase() === 'paid'"
          class="rounded-lg border border-cyan-400/20 bg-cyan-400/10 px-4 py-2.5 text-center text-sm font-extrabold text-cyan-200"
        >
          Đã thanh toán Escrow
        </div>

        <button
          v-if="selectedManagedApplicant.status?.toLowerCase() === 'paid' && !reviewsList.some(review => Number(review.reviewer_id || review.ReviewerID) === Number(currentBusinessUserId))"
          @click="openBusinessReviewModal(selectedManagedApplicant); closeManagedApplicantModal()"
          class="rounded-lg bg-amber-400 px-4 py-2.5 text-sm font-extrabold text-slate-950 shadow-lg shadow-amber-500/20 hover:bg-amber-300"
        >
          Đánh giá sinh viên
        </button>

        <button
          @click="openReviewsModal(selectedManagedApplicant)"
          class="rounded-lg border border-white/10 bg-white/5 px-4 py-2.5 text-sm font-bold text-slate-200 hover:bg-white/10 hover:text-white"
        >
          Xem đánh giá
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

  <!-- Business Dispute Ticket Modal -->
  <div v-if="isBusinessDisputeModalOpen && businessDisputeTargetApp" class="fixed inset-0 z-[10001] flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
    <div class="bg-slate-900 rounded-3xl border border-rose-500/30 shadow-2xl shadow-slate-950 max-w-lg w-full overflow-hidden animate-in fade-in zoom-in-95 duration-200">
      <div class="p-5 border-b border-indigo-500/15 bg-slate-950/80 flex justify-between items-center">
        <div>
          <h3 class="text-base font-extrabold text-white">⚠️ Gửi Khiếu Nại Ứng Viên</h3>
          <p class="text-xs text-rose-300 font-semibold mt-0.5">Ứng viên: {{ businessDisputeTargetApp.student?.full_name || 'N/A' }} · Mã đơn #{{ businessDisputeTargetApp.id }}</p>
        </div>
        <button @click="isBusinessDisputeModalOpen = false" class="text-slate-400 hover:text-white text-xl font-bold p-1 rounded-lg hover:bg-white/10 cursor-pointer">&times;</button>
      </div>

      <form @submit.prevent="submitBusinessDisputeTicket" class="p-6 space-y-4">
        <div>
          <label class="text-xs font-bold text-slate-300 uppercase tracking-wider block mb-1.5">Lý do khiếu nại <span class="text-rose-400">*</span></label>
          <select
            v-model="businessDisputeForm.reason"
            required
            class="w-full text-xs px-3.5 py-2.5 rounded-xl border border-white/15 bg-slate-950 text-white focus:outline-none focus:border-rose-400"
          >
            <option value="" disabled>-- Chọn lý do khiếu nại --</option>
            <option value="Ứng viên vắng mặt không báo trước">Ứng viên vắng mặt không báo trước</option>
            <option value="Nộp báo cáo công việc không đúng thực tế">Nộp báo cáo công việc không đúng thực tế</option>
            <option value="Thái độ cư xử vi phạm quy định">Thái độ cư xử vi phạm quy định</option>
            <option value="Khác">Lý do khác</option>
          </select>
        </div>

        <div>
          <label class="text-xs font-bold text-slate-300 uppercase tracking-wider block mb-1.5">Mô tả chi tiết sự việc <span class="text-rose-400">*</span></label>
          <textarea
            v-model="businessDisputeForm.description"
            rows="3"
            required
            placeholder="Mô tả cụ thể vi phạm của ứng viên, thời gian diễn ra và yêu cầu Ban quản trị giải quyết..."
            class="w-full text-xs p-3 rounded-xl border border-white/15 bg-slate-950 text-white placeholder-slate-500 focus:outline-none focus:border-rose-400"
          ></textarea>
        </div>

        <div>
          <label class="text-xs font-bold text-slate-300 uppercase tracking-wider block mb-1.5">Đề xuất xử lý với Admin (Tùy chọn)</label>
          <input
            v-model="businessDisputeForm.requested_action"
            type="text"
            placeholder="VD: Khóa tài khoản ứng viên / Hủy hoàn trả lương..."
            class="w-full text-xs px-3.5 py-2.5 rounded-xl border border-white/15 bg-slate-950 text-white placeholder-slate-500 focus:outline-none focus:border-rose-400 font-medium"
          />
        </div>

        <div>
          <label class="text-xs font-bold text-cyan-300 uppercase tracking-wider block mb-1.5">File/Ảnh Bằng Chứng (Tùy chọn)</label>
          <div class="space-y-2">
            <div class="flex items-center gap-2">
              <input
                type="file"
                ref="businessEvidenceFileInput"
                @change="$emit('upload-evidence', $event)"
                accept="image/*,.pdf,.doc,.docx,.zip"
                class="hidden"
              />
              <button
                type="button"
                @click="businessEvidenceFileInput?.click()"
                :disabled="state.isUploadingBusinessEvidence"
                class="flex-1 px-4 py-2.5 rounded-xl border border-dashed border-cyan-400/40 bg-cyan-950/20 hover:bg-cyan-900/40 text-cyan-300 text-xs font-bold transition-all flex items-center justify-center gap-2 cursor-pointer disabled:opacity-50"
              >
                <span>{{ state.isUploadingBusinessEvidence ? '⏳ Đang tải file lên hệ thống...' : '📁 Tải File Bằng Chứng Từ Máy Tính (Ảnh/PDF/Doc/Zip)' }}</span>
              </button>
            </div>

            <div v-if="businessDisputeForm.evidence_url" class="flex items-center justify-between p-2.5 rounded-xl bg-cyan-950/60 border border-cyan-400/30 text-xs">
              <div class="flex items-center gap-2 overflow-hidden pr-2">
                <span class="text-cyan-400 font-bold">✅ Bằng chứng:</span>
                <a :href="businessDisputeForm.evidence_url" target="_blank" class="text-cyan-200 underline truncate hover:text-white">{{ businessDisputeForm.evidence_url }}</a>
              </div>
              <button type="button" @click="businessDisputeForm.evidence_url = ''" class="text-rose-400 hover:text-rose-300 font-bold px-2 py-1 bg-rose-950/50 rounded-lg shrink-0">
                ✕ Xóa
              </button>
            </div>
          </div>
        </div>

        <div class="pt-3 border-t border-indigo-500/15 flex justify-end gap-3">
          <button
            type="button"
            @click="isBusinessDisputeModalOpen = false"
            class="px-4 py-2 text-xs font-bold text-slate-300 hover:bg-white/10 border border-white/10 rounded-xl transition-all cursor-pointer"
          >
            Hủy bỏ
          </button>
          <button
            type="submit"
            :disabled="isSubmittingBusinessDispute"
            class="px-5 py-2 text-xs font-extrabold text-white bg-rose-500 hover:bg-rose-400 disabled:opacity-50 rounded-xl shadow-lg shadow-rose-500/20 transition-all cursor-pointer uppercase tracking-wider"
          >
            {{ isSubmittingBusinessDispute ? 'Đang gửi...' : '⚠️ Xác Nhận Gửi Khiếu Nại' }}
          </button>
        </div>
      </form>
    </div>
  </div>

  <!-- Business Dispute Success Modal -->
  <div v-if="showDisputeSuccessModal" class="fixed inset-0 z-[10002] flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
    <div class="bg-slate-900 rounded-3xl border border-indigo-500/30 shadow-2xl shadow-slate-950 max-w-md w-full overflow-hidden animate-in fade-in zoom-in-95 duration-200">
      <div class="p-6 text-center space-y-4">
        <div class="mx-auto flex h-16 w-16 items-center justify-center rounded-2xl bg-indigo-500/10 border border-indigo-500/30 text-indigo-400 text-3xl shadow-lg shadow-indigo-500/20">
          🎉
        </div>
        <div>
          <h3 class="text-lg font-extrabold text-white">Gửi Khiếu Nại Thành Công</h3>
          <p class="mt-2 text-xs font-semibold text-slate-300 leading-relaxed">
            {{ disputeSuccessMsg }}
          </p>
        </div>
        <div class="pt-2">
          <button
            @click="showDisputeSuccessModal = false"
            class="w-full py-2.5 rounded-xl bg-gradient-to-r from-indigo-500 to-emerald-500 hover:brightness-110 text-white text-xs font-extrabold shadow-lg shadow-indigo-500/30 transition-all cursor-pointer uppercase tracking-wider"
          >
            Đã Hiểu / Đóng
          </button>
        </div>
      </div>
    </div>
  </div>

  <!-- Business Dispute Error/Warning Modal -->
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

  <!-- Business Reappeal / View Verdict Modal -->
  <div v-if="state.showBusinessReappealModal" class="fixed inset-0 z-[10002] flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
    <div class="bg-slate-900 rounded-3xl border border-white/10 shadow-2xl shadow-slate-950 max-w-lg w-full overflow-hidden animate-in fade-in zoom-in-95 duration-200">
      <div class="p-6 space-y-4">
        <div class="flex items-center justify-between border-b border-white/10 pb-3">
          <h3 class="text-sm font-extrabold text-white flex items-center gap-2">
            <span>⚖️ KẾT QUẢ KHIẾU NẠI & PHÁN QUYẾT ADMIN</span>
          </h3>
          <button @click="state.showBusinessReappealModal = false" class="text-slate-400 hover:text-white font-bold text-lg">✕</button>
        </div>

        <div v-if="state.selectedBusinessTicketForView" class="space-y-3">
          <div class="bg-slate-950/60 p-3 rounded-xl border border-white/10 space-y-1.5 text-xs">
            <div class="flex justify-between text-slate-400">
              <span>Đơn ứng tuyển: <strong class="text-white">#{{ state.selectedBusinessTicketForView.application_id }}</strong></span>
              <span>Trạng thái: 
                <span :class="state.selectedBusinessTicketForView.status === 'resolved' ? 'text-emerald-400 font-bold' : state.selectedBusinessTicketForView.status === 'rejected' ? 'text-rose-400 font-bold' : 'text-amber-400 font-bold'">
                  {{ state.selectedBusinessTicketForView.status === 'resolved' ? 'Đã duyệt' : state.selectedBusinessTicketForView.status === 'rejected' ? 'Bị bác bỏ' : 'Đang xử lý' }}
                </span>
              </span>
            </div>
            <div><strong class="text-slate-300">Lý do ban đầu:</strong> {{ state.selectedBusinessTicketForView.reason }}</div>
            <div><strong class="text-slate-300">Mô tả:</strong> {{ state.selectedBusinessTicketForView.description }}</div>
          </div>

          <div v-if="state.selectedBusinessTicketForView.verdict" class="bg-emerald-950/40 p-4 rounded-xl border border-emerald-500/30 space-y-1">
            <div class="text-xs font-bold text-emerald-400 uppercase tracking-wider">Phán quyết chính thức từ Admin:</div>
            <div class="text-sm font-semibold text-white leading-relaxed">{{ state.selectedBusinessTicketForView.verdict }}</div>
          </div>

          <div v-if="state.selectedBusinessTicketForView.is_reappealed" class="bg-amber-950/40 p-3 rounded-xl border border-amber-500/30 text-xs text-amber-200">
            <div class="font-bold text-amber-300">⚠️ Đã gửi yêu cầu Tái xem xét phán quyết</div>
            <div class="mt-1">Lý do phản hồi: {{ state.selectedBusinessTicketForView.reappeal_reason }}</div>
          </div>

          <!-- Form Reappeal if verdict exists and not reappealed -->
          <div v-else-if="state.selectedBusinessTicketForView.verdict" class="pt-2 space-y-3 border-t border-white/10">
            <div class="text-xs font-bold text-rose-300">Bạn thấy phán quyết chưa thỏa đáng? Nộp yêu cầu Tái xem xét:</div>
            <textarea
              v-model="state.businessReappealForm.reason"
              rows="3"
              placeholder="Nhập lý do chi tiết giải thích vì sao phán quyết chưa công bằng/thỏa đáng..."
              class="w-full text-xs p-3 rounded-xl border border-white/15 bg-slate-950 text-white placeholder-slate-500 focus:outline-none focus:border-rose-400"
            ></textarea>

            <div class="flex justify-end gap-3 pt-2">
              <button
                type="button"
                @click="state.showBusinessReappealModal = false"
                class="px-4 py-2 text-xs font-bold text-slate-300 hover:bg-white/10 border border-white/10 rounded-xl"
              >
                Đóng
              </button>
              <button
                type="button"
                @click="state.submitBusinessReappealTicket()"
                :disabled="state.isSubmittingBusinessReappeal"
                class="px-4 py-2 text-xs font-extrabold text-white bg-rose-500 hover:bg-rose-400 disabled:opacity-50 rounded-xl shadow-lg shadow-rose-500/20 uppercase tracking-wider cursor-pointer"
              >
                {{ state.isSubmittingBusinessReappeal ? 'Đang gửi...' : '⚠️ Nộp Yêu Cầu Tái Xem Xét' }}
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
import { useMedia } from '~/composables/useMedia'

const { getMediaUrl } = useMedia()
const props = defineProps<{ state: Record<string, any> }>()
const businessEvidenceFileInput = ref<HTMLInputElement | null>(null)
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
  currentBusinessUserId,
  isBusinessDisputeModalOpen,
  businessDisputeTargetApp,
  isSubmittingBusinessDispute,
  businessDisputeForm,
  openBusinessDisputeModal,
  submitBusinessDisputeTicket,
  showDisputeSuccessModal,
  disputeSuccessMsg,
  showDisputeErrorModal,
  disputeErrorMsg,
  myBusinessTickets,
  showBusinessReappealModal,
  openBusinessReappealModal,
  submitBusinessReappealTicket
} = toRefs(props.state)
</script>
