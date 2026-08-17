<template>
      <!-- Section 3: Jobs (List and Post Form) -->
      <div v-show="activeSection === 'jobs'" class="space-y-6">
        <div class="flex justify-end">
          <button
            @click="showCreateForm = !showCreateForm"
            class="px-5 py-2.5 bg-gradient-to-r from-cyan-400 to-emerald-400 hover:from-cyan-300 hover:to-emerald-300 text-slate-950 font-extrabold text-xs uppercase tracking-wider rounded-xl shadow-lg shadow-cyan-500/20 focus-ring transition-all"
          >
            {{ showCreateForm ? '← Quay lại Danh Sách' : '+ Đăng Tin Tuyển Dụng Mới' }}
          </button>
        </div>

        <!-- Post job form with Live Preview -->
        <div v-if="showCreateForm" class="grid grid-cols-1 lg:grid-cols-12 gap-6">
          <!-- Left: Input Form -->
          <div class="lg:col-span-7 bg-slate-900/90 rounded-3xl border border-cyan-400/20 shadow-2xl p-6 sm:p-8 backdrop-blur-xl space-y-6">
            <div class="flex items-center justify-between border-b border-white/10 pb-4">
              <div>
                <span class="text-[10px] font-black uppercase tracking-wider text-cyan-300 bg-cyan-400/10 px-2.5 py-0.5 rounded-full border border-cyan-400/20">
                  Tuyển Dụng Nhanh
                </span>
                <h3 class="mt-1 text-lg font-extrabold text-white">Soạn Thảo Thông Tin Tuyển Dụng</h3>
              </div>
              <button
                type="button"
                @click="handleAiGenerateJob"
                :disabled="isGeneratingAi"
                class="inline-flex items-center gap-1.5 px-3.5 py-2 bg-gradient-to-r from-indigo-500 to-cyan-500 hover:from-indigo-400 hover:to-cyan-400 text-white text-xs font-black rounded-xl shadow-md shadow-indigo-500/20 transition-all disabled:opacity-50"
              >
                <span>{{ isGeneratingAi ? '✨ Đang tạo...' : '✨ AI Soạn Mô Tả' }}</span>
              </button>
            </div>

            <form @submit.prevent="handleCreateJob" class="space-y-4">
              <div>
                <label for="job_title" class="block text-xs font-extrabold text-cyan-300 uppercase tracking-wider mb-1.5">Tiêu đề công việc</label>
                <input
                  id="job_title"
                  type="text"
                  v-model="jobForm.title"
                  required
                  class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-bold"
                  placeholder="VD: Nhân viên Phục vụ Bán thời gian / Thực tập sinh Marketing"
                />
              </div>

              <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                <div>
                  <label for="job_location" class="block text-xs font-extrabold text-cyan-300 uppercase tracking-wider mb-1.5">Địa điểm làm việc</label>
                  <input
                    id="job_location"
                    type="text"
                    v-model="jobForm.location"
                    required
                    class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium"
                    placeholder="VD: Quận 1, TP. Hồ Chí Minh"
                  />
                </div>

                <div>
                  <label for="job_working_date" class="block text-xs font-extrabold text-cyan-300 uppercase tracking-wider mb-1.5">Lịch làm việc</label>
                  <input
                    id="job_working_date"
                    type="text"
                    v-model="jobForm.working_date"
                    class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium"
                    placeholder="VD: T2 - T6 (Ca 4 tiếng)"
                  />
                </div>

                <div>
                  <label for="job_salary" class="block text-xs font-extrabold text-cyan-300 uppercase tracking-wider mb-1.5">Mức lương chi trả (VNĐ)</label>
                  <input
                    id="job_salary"
                    type="number"
                    v-model="jobForm.salary"
                    required
                    class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-bold"
                    placeholder="VD: 5000000"
                  />
                </div>

                <div>
                  <label for="job_slots" class="block text-xs font-extrabold text-cyan-300 uppercase tracking-wider mb-1.5">Số lượng tuyển dụng</label>
                  <input
                    id="job_slots"
                    type="number"
                    v-model="jobForm.slots"
                    required
                    min="1"
                    class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium"
                    placeholder="1"
                  />
                </div>
              </div>

              <div>
                <label for="job_description" class="block text-xs font-extrabold text-cyan-300 uppercase tracking-wider mb-1.5">Mô tả công việc & Yêu cầu</label>
                <textarea
                  id="job_description"
                  rows="6"
                  v-model="jobForm.description"
                  required
                  class="block w-full px-4 py-3 border border-white/10 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium leading-relaxed"
                  placeholder="Mô tả công việc chi tiết, trách nhiệm, quyền lợi..."
                ></textarea>
              </div>

              <div class="pt-4 border-t border-white/10 flex items-center justify-end gap-3">
                <button
                  type="button"
                  @click="showCreateForm = false"
                  class="px-4 py-2.5 border border-white/10 text-xs font-bold rounded-xl text-slate-300 bg-white/5 hover:bg-white/10 transition-all"
                >
                  Hủy bỏ
                </button>
                <button
                  type="submit"
                  :disabled="isCreatingJob"
                  class="px-6 py-2.5 rounded-xl text-xs font-black text-slate-950 bg-gradient-to-r from-cyan-400 to-emerald-400 hover:from-cyan-300 hover:to-emerald-300 shadow-lg shadow-cyan-500/25 transition-all disabled:opacity-50"
                >
                  <span v-if="isCreatingJob" class="flex items-center space-x-2">
                    <svg class="animate-spin h-4 w-4 text-slate-950" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    <span>Đang đăng tin...</span>
                  </span>
                  <span v-else>Đăng Tin Tuyển Dụng</span>
                </button>
              </div>
            </form>
          </div>

          <!-- Right: Live Preview Card (TopCV Style) -->
          <div class="lg:col-span-5 space-y-4">
            <div class="bg-slate-900/90 rounded-3xl border border-white/10 p-6 shadow-2xl backdrop-blur-xl">
              <div class="flex items-center justify-between border-b border-white/10 pb-3 mb-4">
                <span class="text-[10px] font-black uppercase tracking-wider text-emerald-400 bg-emerald-500/10 px-2.5 py-0.5 rounded-full border border-emerald-500/20">
                  Xem Trước Trực Quan
                </span>
                <span class="text-xs font-medium text-slate-400">Hiển thị với Ứng viên</span>
              </div>

              <div class="rounded-2xl border border-cyan-400/30 bg-slate-950 p-4 shadow-xl space-y-3">
                <div class="flex items-start gap-3">
                  <div class="h-12 w-12 rounded-xl bg-gradient-to-br from-indigo-600 to-cyan-500 flex items-center justify-center text-sm font-black text-white flex-shrink-0 shadow-md">
                    {{ (jobForm.title || 'Q').slice(0, 1).toUpperCase() }}
                  </div>
                  <div class="min-w-0 flex-1">
                    <h4 class="text-sm font-extrabold text-white truncate">
                      {{ jobForm.title || 'Tiêu đề công việc sẽ hiển thị tại đây' }}
                    </h4>
                    <p class="text-xs font-semibold text-cyan-300 mt-0.5">Doanh nghiệp của bạn</p>
                  </div>
                </div>

                <div class="flex flex-wrap items-center gap-1.5 pt-1">
                  <span class="text-xs font-black text-emerald-300 bg-emerald-500/10 border border-emerald-500/20 px-2 py-0.5 rounded-md">
                    {{ Number(jobForm.salary || 0).toLocaleString('vi-VN') }} VNĐ
                  </span>
                  <span class="text-[11px] font-medium text-slate-400 bg-slate-800/80 px-2 py-0.5 rounded-md truncate">
                    📍 {{ jobForm.location || 'Địa điểm' }}
                  </span>
                  <span v-if="jobForm.working_date" class="text-[11px] font-medium text-slate-400 bg-slate-800/80 px-2 py-0.5 rounded-md truncate">
                    📅 {{ jobForm.working_date }}
                  </span>
                </div>

                <div class="pt-2 border-t border-white/5 text-xs text-slate-400 line-clamp-3">
                  {{ jobForm.description || 'Nội dung mô tả công việc, yêu cầu kỹ năng và quyền lợi sẽ hiển thị tóm tắt tại đây.' }}
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Job listings table -->
        <div v-else class="bg-slate-900/90 rounded-3xl border border-white/10 shadow-2xl overflow-hidden backdrop-blur-xl">
          <div class="border-b border-white/10 px-6 py-5 flex items-center justify-between">
            <div>
              <h3 class="text-lg font-extrabold text-white">Danh Sách Tin Tuyển Dụng Đã Đăng</h3>
              <p class="text-xs font-semibold text-slate-400 mt-0.5">Theo dõi trạng thái duyệt và số lượng ứng viên ứng tuyển</p>
            </div>
            <span class="text-xs font-bold text-cyan-300 bg-cyan-400/10 px-3 py-1 rounded-full border border-cyan-400/20">
              {{ jobs.length }} tin tuyển dụng
            </span>
          </div>

          <div v-if="jobs.length === 0" class="text-center py-16 text-slate-400 space-y-2">
            <span class="text-4xl block">📋</span>
            <p class="font-extrabold text-white text-base">Chưa có tin tuyển dụng nào</p>
            <p class="text-xs text-slate-400 mt-1">Bấm "+ Đăng Tin Tuyển Dụng Mới" để tìm kiếm ứng viên tài năng.</p>
          </div>
          <div v-else class="overflow-x-auto">
            <table class="min-w-full divide-y divide-white/10">
              <thead class="bg-slate-950">
                <tr>
                  <th class="px-6 py-4 text-left text-xs font-black text-cyan-300 uppercase tracking-wider">Vị trí tuyển dụng</th>
                  <th class="px-6 py-4 text-left text-xs font-black text-cyan-300 uppercase tracking-wider">Địa điểm</th>
                  <th class="px-6 py-4 text-left text-xs font-black text-cyan-300 uppercase tracking-wider">Mức lương</th>
                  <th class="px-6 py-4 text-left text-xs font-black text-cyan-300 uppercase tracking-wider">Chỉ tiêu</th>
                  <th class="px-6 py-4 text-left text-xs font-black text-cyan-300 uppercase tracking-wider">Trạng thái</th>
                </tr>
              </thead>
              <tbody class="bg-slate-900/80 divide-y divide-white/10">
                <tr v-for="job in paginatedBusinessJobs" :key="job.id" class="hover:bg-white/5 transition-colors">
                  <td class="px-6 py-4 whitespace-nowrap text-xs font-extrabold text-white">{{ displayJobTitle(job.title) }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-xs text-slate-300 font-medium">{{ job.location }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-xs font-black text-emerald-300">{{ formatCurrency(job.salary) }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-xs text-slate-300 font-bold">{{ job.slots }} người</td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span :class="[
                      job.status === 'approved' ? 'bg-emerald-500/10 text-emerald-300 border-emerald-500/20' : 'bg-amber-500/10 text-amber-300 border-amber-500/20',
                      'inline-flex items-center px-2.5 py-0.5 rounded-full text-[11px] font-extrabold border uppercase'
                    ]">
                      {{ job.status || 'pending' }}
                    </span>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <PaginationControls
          v-if="!showCreateForm && jobs.length > 0"
          :page="businessJobsPage"
          :page-size="businessJobsPageSize"
          :total-items="jobs.length"
          @update:page="businessJobsPage = $event"
        />
      </div>
</template>

<script setup lang="ts">
import { ref, toRefs } from 'vue'
import { useAi } from '~/composables/useAi'
import { useToast } from '~/composables/useToast'
import PaginationControls from '~/components/common/PaginationControls.vue'

const props = defineProps<{ state: Record<string, any> }>()
const { generateJobDescription } = useAi()
const { success, warning, error } = useToast()
const isGeneratingAi = ref(false)

const handleAiGenerateJob = async () => {
  if (!jobForm.value?.title || !jobForm.value?.title.trim()) {
    warning('Vui lòng nhập Tiêu đề công việc trước (ví dụ: Pha chế ca tối).')
    return
  }

  isGeneratingAi.value = true
  try {
    const res = await generateJobDescription(jobForm.value.title)
    jobForm.value.description = `${res.description}\n\n📌 YÊU CẦU:\n${res.requirements}\n\n🎁 QUYỀN LỢI:\n${res.benefits}`
    success('AI đã tạo tự động bản mô tả công việc thành công!')
  } catch (err: any) {
    error('Không thể tạo mô tả bằng AI. Vui lòng thử lại.')
  } finally {
    isGeneratingAi.value = false
  }
}

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
  paginatedBusinessJobs,
  businessJobsPage,
  businessJobsPageSize,
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
  closeModal,
  submitReview,
  isChatModalOpen,
  selectedChatApp,
  currentBusinessUserId
} = toRefs(props.state)

const formatCurrency = (value: number | string | null | undefined) => {
  const amount = Number(value || 0)
  return `${amount.toLocaleString('vi-VN')} VND`
}

const displayJobTitle = (title: string | null | undefined) => {
  return (title || 'Untitled Job').replace(/\bMarketting\b/gi, 'Marketing')
}
</script>
