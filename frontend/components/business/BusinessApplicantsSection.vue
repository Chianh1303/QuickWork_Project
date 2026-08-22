<template>
  <!-- Section 4: Employer Applicant Management (Synchronized SaaS Dark Theme) -->
  <div v-show="activeSection === 'applicants'" class="space-y-6">
    
    <!-- Header & Search/Filter Control Bar -->
    <div class="bg-slate-900/90 rounded-3xl border border-indigo-500/20 p-5 shadow-2xl backdrop-blur-xl space-y-4">
      <div class="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4 border-b border-white/10 pb-4">
        <div>
          <h2 class="text-xl font-extrabold text-white tracking-tight">Danh Sách Ứng Viên Nộp Đơn</h2>
          <p class="text-xs font-semibold text-slate-400 mt-0.5">Duyệt hồ sơ, phỏng vấn, trao đổi trực tiếp và quản lý hợp đồng làm việc</p>
        </div>

        <span class="inline-flex items-center gap-1.5 rounded-full border border-cyan-400/30 bg-cyan-400/10 px-3.5 py-1 text-xs font-black text-cyan-300">
          👥 Tổng đơn: {{ filteredApps.length }} ứng viên
        </span>
      </div>

      <!-- Search Input & Status Filter Grid -->
      <div class="grid grid-cols-1 sm:grid-cols-12 gap-3">
        <!-- Search Input -->
        <div class="sm:col-span-7 relative">
          <span class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-slate-400">
            <svg class="h-4 w-4 text-cyan-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </span>
          <input
            v-model="applicantSearchQuery"
            type="text"
            class="block w-full pl-10 pr-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-950 text-white placeholder-slate-400 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-semibold"
            placeholder="Tìm kiếm theo tên sinh viên, số điện thoại hoặc vị trí công việc..."
          />
        </div>

        <!-- Status Filter Select -->
        <div class="sm:col-span-5">
          <select
            v-model="applicantStatusFilter"
            class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-950 text-slate-200 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-bold cursor-pointer"
          >
            <option value="all">Tất cả trạng thái ứng tuyển</option>
            <option value="pending">Mới ứng tuyển (Chờ duyệt)</option>
            <option value="approved">Đã duyệt (Trúng tuyển)</option>
            <option value="offer_accepted">Đã chấp nhận Offer</option>
            <option value="student_completed">Chờ Doanh nghiệp xác nhận hoàn thành</option>
            <option value="paid">Đã hoàn thành & Giải ngân</option>
            <option value="rejected">Đã từ chối</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Loading state skeleton -->
    <div v-if="isLoadingApps" class="bg-slate-900/90 rounded-3xl border border-white/10 p-8 space-y-4 animate-pulse backdrop-blur-xl">
      <div class="h-6 bg-slate-800 rounded-xl w-1/4"></div>
      <div class="space-y-3">
        <div class="h-12 bg-slate-800/60 rounded-xl"></div>
        <div class="h-12 bg-slate-800/60 rounded-xl"></div>
        <div class="h-12 bg-slate-800/60 rounded-xl"></div>
      </div>
    </div>

    <!-- Empty state -->
    <div v-else-if="filteredApps.length === 0" class="bg-slate-900/90 text-center py-16 px-6 rounded-3xl border border-white/10 shadow-2xl backdrop-blur-xl space-y-3">
      <span class="text-5xl block">👥</span>
      <h3 class="text-lg font-extrabold text-white">Chưa Có Đơn Ứng Tuyển Nào</h3>
      <p class="text-xs text-slate-400 max-w-md mx-auto font-medium leading-relaxed">
        Hiện chưa có ứng viên nộp đơn phù hợp với điều kiện tìm kiếm hoặc bộ lọc trạng thái của bạn.
      </p>
    </div>

    <!-- DESKTOP SYNCHRONIZED DARK TABLE -->
    <div v-else class="hidden lg:block overflow-hidden rounded-3xl border border-white/10 bg-slate-900/90 shadow-2xl backdrop-blur-xl">
      <table class="min-w-full table-fixed divide-y divide-white/10">
        <thead class="bg-slate-950">
          <tr>
            <th class="w-[28%] px-5 py-4 text-left text-xs font-extrabold text-cyan-300 uppercase tracking-wider">Ứng viên & Kỹ năng</th>
            <th class="w-[24%] px-5 py-4 text-left text-xs font-extrabold text-cyan-300 uppercase tracking-wider">Vị trí ứng tuyển</th>
            <th class="w-[14%] px-5 py-4 text-left text-xs font-extrabold text-cyan-300 uppercase tracking-wider">Liên hệ</th>
            <th class="w-[12%] px-5 py-4 text-left text-xs font-extrabold text-cyan-300 uppercase tracking-wider">Ngày nộp</th>
            <th class="w-[12%] px-5 py-4 text-left text-xs font-extrabold text-cyan-300 uppercase tracking-wider">Trạng thái</th>
            <th class="w-[10%] px-5 py-4 text-right text-xs font-extrabold text-cyan-300 uppercase tracking-wider">Thao tác</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-white/5 bg-slate-900/60">
          <tr v-for="app in paginatedApplicants" :key="app.id" class="transition-colors hover:bg-white/5">
            
            <!-- Candidate Info Column -->
            <td class="px-5 py-4">
              <div class="flex items-start gap-3">
                <img 
                  v-if="app.student?.avatar_url" 
                  :src="app.student.avatar_url" 
                  class="h-10 w-10 rounded-xl object-cover border border-cyan-400/30 shadow-md flex-shrink-0" 
                />
                <div v-else class="h-10 w-10 rounded-xl bg-gradient-to-br from-indigo-600 to-cyan-500 border border-cyan-400/30 flex items-center justify-center text-white font-black text-sm flex-shrink-0 shadow-md">
                  {{ (app.student?.full_name || 'S').slice(0, 1).toUpperCase() }}
                </div>
                
                <div class="min-w-0 space-y-1">
                  <div class="flex items-center gap-2">
                    <span class="truncate text-sm font-extrabold text-white tracking-tight">{{ app.student?.full_name || 'Ứng viên chưa cập nhật' }}</span>
                    
                    <a 
                      v-if="app.student?.cv_url" 
                      :href="app.student.cv_url" 
                      target="_blank" 
                      class="inline-flex flex-shrink-0 items-center gap-1 text-[10px] text-rose-300 hover:text-white font-extrabold bg-rose-500/10 hover:bg-rose-500/20 px-2 py-0.5 rounded-md border border-rose-500/20 transition-all shadow-sm"
                    >
                      <svg class="h-3 w-3" fill="currentColor" viewBox="0 0 20 20">
                        <path d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4z" />
                      </svg>
                      <span>Xem CV</span>
                    </a>
                  </div>

                  <div class="flex flex-wrap gap-1">
                    <span 
                      v-for="(skill, sIdx) in parseSkills(app.student?.skills)" 
                      :key="sIdx"
                      class="inline-block bg-slate-950 text-indigo-300 border border-indigo-500/20 text-[10px] font-bold px-2 py-0.5 rounded-md"
                    >
                      {{ skill }}
                    </span>
                    <span v-if="!app.student?.skills || parseSkills(app.student?.skills).length === 0" class="text-[11px] text-slate-500 italic">
                      Chưa cập nhật kỹ năng
                    </span>
                  </div>
                </div>
              </div>
            </td>

            <!-- Job Title Column -->
            <td class="px-5 py-4 text-xs font-bold text-slate-200">
              <div class="truncate bg-slate-950/80 px-3 py-1.5 rounded-xl border border-white/5 inline-block max-w-full">
                {{ app.job?.title || jobTitleLookup(app.job_id) }}
              </div>
            </td>

            <!-- Phone Contact Column -->
            <td class="px-5 py-4 whitespace-nowrap text-xs text-slate-300 font-semibold">
              📞 {{ app.student?.phone || 'Chưa có SĐT' }}
            </td>

            <!-- Applied Date Column -->
            <td class="px-5 py-4 whitespace-nowrap text-xs text-slate-400 font-medium">
              📅 {{ formatDate(app.applied_at || app.id) }}
            </td>

            <!-- Status Badge Column -->
            <td class="px-5 py-4 whitespace-nowrap">
              <span :class="[
                statusBadgeClass(app.status),
                'inline-flex items-center px-2.5 py-1 rounded-full text-[11px] font-extrabold border uppercase tracking-wider'
              ]">
                {{ app.status ? app.status.replace(/_/g, ' ') : 'Chờ duyệt' }}
              </span>
            </td>

            <!-- Action Buttons Toolbar -->
            <td class="px-5 py-4 text-right whitespace-nowrap">
              <div class="flex items-center justify-end gap-2">
                <!-- Chat Button -->
                <button 
                  @click="openChatModal(app)" 
                  class="p-2 rounded-xl border border-indigo-500/20 bg-indigo-500/10 text-indigo-300 hover:bg-indigo-500 hover:text-white transition-all cursor-pointer"
                  title="Nhắn tin trực tiếp với ứng viên"
                >
                  <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
                  </svg>
                </button>

                <!-- Prominent Quản lý Button (Opens full Candidate Management Drawer) -->
                <button
                  @click="openManagedApplicantModal(app)"
                  class="px-4 py-2 rounded-xl bg-gradient-to-r from-cyan-400 via-blue-500 to-emerald-400 text-slate-950 font-black text-xs hover:brightness-110 shadow-lg shadow-cyan-500/25 transition-all cursor-pointer"
                  title="Quản lý chi tiết ứng viên & xét duyệt"
                >
                  Quản lý
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- MOBILE SYNCHRONIZED CARDS -->
    <div class="lg:hidden space-y-4">
      <div
        v-for="app in paginatedApplicants"
        :key="app.id"
        class="bg-slate-900/90 border border-white/10 rounded-2xl p-5 shadow-xl space-y-4 backdrop-blur-xl"
      >
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <img v-if="app.student?.avatar_url" :src="app.student.avatar_url" class="h-10 w-10 rounded-xl object-cover border border-cyan-400/30" />
            <div v-else class="h-10 w-10 bg-indigo-600 rounded-xl border border-cyan-400/30 flex items-center justify-center font-black text-white">
              {{ (app.student?.full_name || 'S').slice(0, 1).toUpperCase() }}
            </div>
            <div>
              <h4 class="text-sm font-extrabold text-white">{{ app.student?.full_name || 'N/A' }}</h4>
              <p class="text-xs text-slate-400">SĐT: {{ app.student?.phone || 'Chưa cập nhật' }}</p>
            </div>
          </div>

          <span :class="[
            statusBadgeClass(app.status),
            'inline-flex items-center px-2.5 py-1 rounded-full text-[10px] font-extrabold border uppercase'
          ]">
            {{ app.status }}
          </span>
        </div>

        <div class="pt-2 border-t border-white/10 text-xs text-slate-300 space-y-1">
          <div>Vị trí: <span class="text-cyan-300 font-extrabold">{{ jobTitleLookup(app.job_id) }}</span></div>
          <div>Ngày nộp: <span class="text-slate-400 font-medium">{{ formatDate(app.applied_at || app.id) }}</span></div>
        </div>

        <!-- Mobile actions -->
        <div class="grid grid-cols-2 gap-2 pt-2 border-t border-white/5">
          <button
            @click="openChatModal(app)"
            class="w-full rounded-xl border border-indigo-500/20 bg-indigo-500/10 py-2 text-center text-xs font-bold text-indigo-300 hover:bg-indigo-500 hover:text-white transition-all cursor-pointer"
          >
            💬 Nhắn tin
          </button>

          <button
            @click="openManagedApplicantModal(app)"
            class="w-full rounded-xl bg-gradient-to-r from-cyan-400 to-emerald-400 py-2 text-center text-xs font-black text-slate-950 shadow-md transition-all cursor-pointer"
          >
            Quản lý
          </button>
        </div>
      </div>
    </div>

    <!-- Pagination Controls -->
    <PaginationControls
      v-if="filteredApps.length > 0"
      :page="applicantsPage"
      :page-size="applicantsPageSize"
      :total-items="filteredApps.length"
      @update:page="applicantsPage = $event"
    />
  </div>
</template>

<script setup lang="ts">
import { toRefs } from 'vue'
import PaginationControls from '~/components/common/PaginationControls.vue'

const props = defineProps<{ state: Record<string, any> }>()
const {
  activeSection,
  isLoadingApps,
  applicantSearchQuery,
  applicantStatusFilter,
  filteredApps,
  paginatedApplicants,
  applicantsPage,
  applicantsPageSize,
  jobTitleLookup,
  formatDate,
  statusBadgeClass,
  parseSkills,
  openChatModal,
  openReviewModal,
  openCompletionModal,
  openBusinessReviewModal,
  openManagedApplicantModal
} = toRefs(props.state)
</script>
