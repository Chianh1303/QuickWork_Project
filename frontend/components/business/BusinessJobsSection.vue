<template>
  <!-- Section: Business Job Management -->
  <div v-show="activeSection === 'jobs'" class="space-y-6">
    
    <!-- Top Bar: Quick Stats & Create Button -->
    <div class="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4 bg-slate-900/90 rounded-3xl border border-indigo-500/20 p-5 shadow-xl backdrop-blur-xl">
      <div>
        <h2 class="text-xl font-extrabold text-white tracking-tight">Quản Lý Bài Đăng Tuyển Dụng</h2>
        <p class="text-xs font-semibold text-slate-400 mt-0.5">Theo dõi tin tuyển dụng, lượt nộp đơn và quản lý trạng thái tuyển dụng</p>
      </div>

      <button
        @click="showCreateForm ? (showCreateForm = false) : openCreateJobModal()"
        class="flex items-center gap-2 px-5 py-2.5 bg-gradient-to-r from-indigo-500 via-blue-600 to-cyan-400 hover:brightness-110 text-white font-extrabold text-xs uppercase tracking-wider rounded-xl shadow-lg shadow-indigo-500/25 transition-all cursor-pointer"
      >
        <span class="text-base">{{ showCreateForm ? '←' : '+' }}</span>
        <span>{{ showCreateForm ? 'Quay lại Danh Sách' : 'Đăng Tin Tuyển Dụng Mới' }}</span>
      </button>
    </div>

    <!-- CREATE OR EDIT JOB FORM WITH LIVE PREVIEW (TOPCV STYLE 3-SECTIONS) -->
    <div v-if="showCreateForm" class="grid grid-cols-1 lg:grid-cols-12 gap-6">
      <!-- Left Column: TopCV Input Form -->
      <div class="lg:col-span-7 bg-slate-900/95 rounded-3xl border border-cyan-400/30 shadow-2xl p-6 sm:p-8 backdrop-blur-xl space-y-6">
        <div class="border-b border-white/10 pb-4">
          <span class="text-[10px] font-black uppercase tracking-wider text-cyan-300 bg-cyan-400/10 px-2.5 py-0.5 rounded-full border border-cyan-400/20">
            {{ isEditingJob ? 'Chỉnh Sửa Bài Đăng' : 'Tạo Tin Tuyển Dụng Mới' }}
          </span>
          <h3 class="mt-1 text-lg font-extrabold text-white">
            {{ isEditingJob ? 'Cập Nhật Thông Tin Tuyển Dụng' : 'Soạn Thảo Thông Tin Tuyển Dụng Chuẩn TopCV' }}
          </h3>
        </div>

        <form @submit.prevent="handleCreateOrUpdateJob" class="space-y-6">
          
          <!-- PHẦN 1: THÔNG TIN VỊ TRÍ TUYỂN DỤNG -->
          <div class="rounded-2xl border border-white/10 bg-slate-950/60 p-4.5 space-y-4">
            <div class="flex items-center gap-2 text-cyan-300 font-extrabold text-xs uppercase tracking-wider">
              <span>📍</span>
              <span>Phần 1: Thông Tin Vị Trí Tuyển Dụng</span>
            </div>

            <!-- Job Title -->
            <div>
              <label for="job_title" class="block text-xs font-bold text-slate-300 mb-1">Tiêu đề vị trí tuyển dụng <span class="text-rose-400">*</span></label>
              <input
                id="job_title"
                type="text"
                v-model="jobForm.title"
                required
                class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-900 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-bold"
                placeholder="VD: Thực tập sinh Backend Developer (Go / Node.js) hoặc Barista ca tối"
              />
            </div>

            <!-- Category & Job Type -->
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
              <div>
                <label for="job_category" class="block text-xs font-bold text-slate-300 mb-1">Ngành nghề tuyển dụng <span class="text-rose-400">*</span></label>
                <select
                  id="job_category"
                  v-model="jobForm.category"
                  class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-900 text-slate-100 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-semibold cursor-pointer"
                >
                  <option value="it">CNTT & Phần mềm</option>
                  <option value="marketing">Marketing & Media</option>
                  <option value="design">Thiết Kế Đồ Họa</option>
                  <option value="f&b">F&B & Phục Vụ</option>
                  <option value="giasu">Gia Sư & Giáo Dục</option>
                  <option value="sales">Bán Hàng & Sales</option>
                </select>
              </div>

              <div>
                <label for="job_type" class="block text-xs font-bold text-slate-300 mb-1">Hình thức làm việc <span class="text-rose-400">*</span></label>
                <select
                  id="job_type"
                  v-model="jobForm.job_type"
                  class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-900 text-slate-100 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-semibold cursor-pointer"
                >
                  <option value="part_time">Bán thời gian (Part-Time)</option>
                  <option value="internship">Thực tập (Internship)</option>
                  <option value="remote">Làm việc từ xa (Remote)</option>
                </select>
              </div>
            </div>

            <!-- Location -->
            <div>
              <label for="job_location" class="block text-xs font-bold text-slate-300 mb-1">Địa điểm làm việc cụ thể <span class="text-rose-400">*</span></label>
              <input
                id="job_location"
                type="text"
                v-model="jobForm.location"
                required
                class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-900 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium"
                placeholder="VD: Số 123 Đường Nam Kỳ Khởi Nghĩa, Quận 1, TP. Hồ Chí Minh"
              />
            </div>
          </div>

          <!-- PHẦN 2: LƯƠNG, CHỈ TIÊU & CA LÀM -->
          <div class="rounded-2xl border border-white/10 bg-slate-950/60 p-4.5 space-y-4">
            <div class="flex items-center gap-2 text-emerald-300 font-extrabold text-xs uppercase tracking-wider">
              <span>💰</span>
              <span>Phần 2: Lương, Chỉ Tiêu & Lịch Ca Làm</span>
            </div>

            <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
              <div>
                <label for="job_salary" class="block text-xs font-bold text-slate-300 mb-1">Mức lương (VNĐ) <span class="text-rose-400">*</span></label>
                <input
                  id="job_salary"
                  type="number"
                  v-model="jobForm.salary"
                  required
                  class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-900 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-bold"
                  placeholder="VD: 5000000"
                />
                <span class="text-[10px] text-emerald-400 font-semibold mt-1 block">
                  = {{ Number(jobForm.salary || 0).toLocaleString('vi-VN') }} VNĐ
                </span>
              </div>

              <div>
                <label for="job_slots" class="block text-xs font-bold text-slate-300 mb-1">Số lượng tuyển (Chỉ tiêu) <span class="text-rose-400">*</span></label>
                <input
                  id="job_slots"
                  type="number"
                  v-model="jobForm.slots"
                  required
                  min="1"
                  class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-900 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium"
                  placeholder="1"
                />
              </div>

              <div>
                <label for="job_working_date" class="block text-xs font-bold text-slate-300 mb-1">Lịch ca làm / Thời gian</label>
                <input
                  id="job_working_date"
                  type="text"
                  v-model="jobForm.working_date"
                  class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-900 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium"
                  placeholder="VD: T2 - T6 (Ca linh hoạt 4 tiếng)"
                />
              </div>
            </div>
          </div>

          <!-- PHẦN 3: NỘI DUNG CHI TIẾT (3 Ô TÁCH BIỆT RÕ RÀNG) -->
          <div class="rounded-2xl border border-white/10 bg-slate-950/60 p-4.5 space-y-4">
            <div class="flex items-center gap-2 text-indigo-300 font-extrabold text-xs uppercase tracking-wider">
              <span>Phần 3: Mô Tả Công Việc, Yêu Cầu & Quyền Lợi</span>
            </div>

            <!-- 1. Mô tả công việc -->
            <div>
              <label for="job_desc_field" class="block text-xs font-bold text-slate-200 mb-1">1. Mô tả công việc (Nhiệm vụ hàng ngày) <span class="text-rose-400">*</span></label>
              <textarea
                id="job_desc_field"
                rows="4"
                v-model="jobForm.description"
                required
                class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-900 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium leading-relaxed"
                placeholder="Nhập trách nhiệm và các công việc cụ thể sinh viên sẽ thực hiện..."
              ></textarea>
            </div>

            <!-- 2. Yêu cầu ứng viên -->
            <div>
              <label for="job_req_field" class="block text-xs font-bold text-slate-200 mb-1">2. Yêu cầu ứng viên (Kỹ năng, độ tuổi, tinh thần...)</label>
              <textarea
                id="job_req_field"
                rows="3"
                v-model="jobForm.requirements"
                class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-900 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium leading-relaxed"
                placeholder="Nhập kỹ năng cần thiết, tính cách, hoặc thiết bị bắt buộc (VD: Có laptop cá nhân)..."
              ></textarea>
            </div>

            <!-- 3. Quyền lợi & Chế độ -->
            <div>
              <label for="job_ben_field" class="block text-xs font-bold text-slate-200 mb-1">3. Quyền lợi & Chế độ đãi ngộ</label>
              <textarea
                id="job_ben_field"
                rows="3"
                v-model="jobForm.benefits"
                class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-900 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium leading-relaxed"
                placeholder="Nhập chế độ thưởng, phụ cấp gửi xe, bao ăn trưa, đào tạo kỹ năng..."
              ></textarea>
            </div>
          </div>

          <!-- Form Actions -->
          <div class="pt-4 border-t border-white/10 flex items-center justify-end gap-3">
            <button
              type="button"
              @click="showCreateForm = false"
              class="px-5 py-2.5 border border-white/10 text-xs font-bold rounded-xl text-slate-300 bg-white/5 hover:bg-white/10 transition-all cursor-pointer"
            >
              Hủy bỏ
            </button>

            <button
              type="submit"
              :disabled="isCreatingJob"
              class="px-7 py-2.5 rounded-xl text-xs font-black text-slate-950 bg-gradient-to-r from-cyan-400 to-emerald-400 hover:brightness-110 shadow-lg shadow-cyan-500/25 transition-all disabled:opacity-50 cursor-pointer"
            >
              <span v-if="isCreatingJob" class="flex items-center space-x-2">
                <svg class="animate-spin h-4 w-4 text-slate-950" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                <span>Đang lưu bài đăng...</span>
              </span>
              <span v-else>{{ isEditingJob ? 'Lưu Thay Đổi Bài Đăng' : 'Đăng Tin Tuyển Dụng' }}</span>
            </button>
          </div>
        </form>
      </div>

      <!-- Right Column: Live Preview Card (TopCV Style) -->
      <div class="lg:col-span-5 space-y-4">
        <div class="bg-slate-900/90 rounded-3xl border border-white/10 p-6 shadow-2xl backdrop-blur-xl">
          <div class="flex items-center justify-between border-b border-white/10 pb-3 mb-4">
            <span class="text-[10px] font-black uppercase tracking-wider text-emerald-400 bg-emerald-500/10 px-2.5 py-0.5 rounded-full border border-emerald-500/20">
              Xem Trước Giao Diện Với Sinh Viên
            </span>
            <span class="text-xs font-medium text-slate-400">Live Preview</span>
          </div>

          <div class="rounded-2xl border border-cyan-400/30 bg-slate-950 p-4.5 shadow-xl space-y-3">
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

            <!-- Preview Sections -->
            <div class="space-y-2 pt-2 border-t border-white/5 text-xs">
              <div>
                <p class="font-extrabold text-cyan-300">Mô tả công việc:</p>
                <p class="text-slate-300 line-clamp-3 leading-relaxed mt-0.5">{{ jobForm.description || 'Nội dung mô tả công việc...' }}</p>
              </div>

              <div v-if="jobForm.requirements">
                <p class="font-extrabold text-cyan-300">Yêu cầu ứng viên:</p>
                <p class="text-slate-300 line-clamp-2 leading-relaxed mt-0.5">{{ jobForm.requirements }}</p>
              </div>

              <div v-if="jobForm.benefits">
                <p class="font-extrabold text-emerald-300">Quyền lợi & Chế độ đãi ngộ:</p>
                <p class="text-slate-300 line-clamp-2 leading-relaxed mt-0.5">{{ jobForm.benefits }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- MAIN JOB LISTINGS CONTAINER WITH SEARCH, FILTERS & ACTION TOOLBAR -->
    <div v-else class="bg-slate-900/90 rounded-3xl border border-white/10 shadow-2xl overflow-hidden backdrop-blur-xl space-y-4 p-6">
      
      <!-- Search & Status Filter Bar -->
      <div class="flex flex-col sm:flex-row items-stretch sm:items-center justify-between gap-3 border-b border-white/10 pb-4">
        <!-- Search Input -->
        <div class="relative flex-1 max-w-md">
          <span class="absolute inset-y-0 left-0 flex items-center pl-3 text-slate-400">
            <svg class="h-4 w-4 text-cyan-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </span>
          <input
            v-model="localSearchQuery"
            type="text"
            placeholder="Tìm theo tên bài đăng..."
            class="w-full pl-9 pr-3 py-2 border border-white/10 rounded-xl bg-slate-950 text-xs font-semibold text-white placeholder-slate-400 focus:border-cyan-400 focus:outline-none"
          />
        </div>

        <!-- Status Filter Pills -->
        <div class="flex items-center gap-1.5 overflow-x-auto">
          <button
            v-for="filter in [
              { id: 'all', label: 'Tất cả' },
              { id: 'open', label: 'Đang mở tuyển 🟢' },
              { id: 'closed', label: 'Đã đóng 🔴' }
            ]"
            :key="filter.id"
            @click="localStatusFilter = filter.id"
            :class="[
              localStatusFilter === filter.id
                ? 'bg-cyan-400 text-slate-950 font-black'
                : 'bg-slate-950 text-slate-300 font-semibold border border-white/10 hover:bg-white/5',
              'px-3 py-1.5 rounded-xl text-xs whitespace-nowrap transition-all cursor-pointer'
            ]"
          >
            {{ filter.label }}
          </button>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="filteredMyJobs.length === 0" class="text-center py-16 text-slate-400 space-y-3">
        <span class="text-5xl block">📋</span>
        <p class="font-extrabold text-white text-base">Không tìm thấy bài tuyển dụng nào</p>
        <p class="text-xs text-slate-400">Bấm "+ Đăng Tin Tuyển Dụng Mới" để tạo bài viết thu hút sinh viên.</p>
      </div>

      <!-- RICH PRODUCTION JOB LIST CARDS -->
      <div v-else class="space-y-3">
        <div
          v-for="job in paginatedFilteredJobs"
          :key="job.id"
          class="group relative flex flex-col lg:flex-row items-start lg:items-center justify-between gap-6 rounded-2xl border border-white/10 bg-slate-950/80 p-5.5 sm:p-6 shadow-xl hover:border-cyan-400/60 hover:bg-slate-950 transition-all"
        >
          <!-- Left Info Column -->
          <div class="space-y-1.5 flex-1 min-w-0">
            <div class="flex flex-wrap items-center gap-2">
              <h3 class="text-base font-extrabold text-white group-hover:text-cyan-300 transition-colors truncate">
                {{ displayJobTitle(job.title) }}
              </h3>

              <!-- Status Badge -->
              <span :class="[
                job.status === 'closed'
                  ? 'bg-rose-500/10 text-rose-300 border-rose-500/20'
                  : 'bg-emerald-500/10 text-emerald-300 border-emerald-500/20',
                'inline-flex items-center gap-1 px-2.5 py-0.5 rounded-full text-[11px] font-black border uppercase'
              ]">
                <span :class="['h-1.5 w-1.5 rounded-full', job.status === 'closed' ? 'bg-rose-400' : 'bg-emerald-400 animate-ping']"></span>
                <span>{{ job.status === 'closed' ? 'Đã đóng tuyển' : 'Đang mở tuyển' }}</span>
              </span>
            </div>

            <!-- Meta badging row -->
            <div class="flex flex-wrap items-center gap-2 text-xs font-semibold text-slate-300">
              <span class="text-emerald-300 font-extrabold bg-emerald-500/10 border border-emerald-500/20 px-2 py-0.5 rounded-md">
                {{ formatCurrency(job.salary) }}
              </span>
              <span class="bg-slate-800 px-2 py-0.5 rounded-md text-slate-300">📍 {{ job.location }}</span>
              <span v-if="job.working_date" class="bg-slate-800 px-2 py-0.5 rounded-md text-slate-300">📅 {{ job.working_date }}</span>
              <span class="bg-indigo-500/10 border border-indigo-500/20 px-2 py-0.5 rounded-md text-indigo-300">Chỉ tiêu: {{ job.slots }} người</span>
            </div>
          </div>

          <!-- Right Column: REAL PRODUCTION ICON BUTTON TOOLBAR -->
          <div class="flex items-center gap-2 w-full lg:w-auto justify-end border-t lg:border-t-0 border-white/5 pt-3 lg:pt-0">
            <!-- Button 1: View Applicants for this job -->
            <button
              @click="filterApplicantsForJob(job.id)"
              class="flex items-center gap-1.5 px-3 py-2 rounded-xl bg-cyan-400/10 border border-cyan-400/30 text-cyan-300 hover:bg-cyan-400 hover:text-slate-950 font-extrabold text-xs transition-all cursor-pointer"
              title="Xem danh sách ứng viên nộp cho bài tuyển dụng này"
            >
              <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
              <span>Ứng viên ({{ getApplicantsCountForJob(job.id) }})</span>
            </button>

            <!-- Button 2: Toggle Status (Open / Closed) -->
            <button
              @click="toggleJobStatus(job)"
              :class="[
                job.status === 'closed'
                  ? 'bg-emerald-500/10 border-emerald-500/30 text-emerald-300 hover:bg-emerald-500/20'
                  : 'bg-amber-500/10 border-amber-500/30 text-amber-300 hover:bg-amber-500/20',
                'p-2 rounded-xl border font-bold text-xs transition-all cursor-pointer'
              ]"
              :title="job.status === 'closed' ? 'Mở tuyển lại' : 'Tạm dừng / Đóng tin tuyển dụng'"
            >
              <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9v6m4-6v6m7-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </button>

            <!-- Button 3: REAL PRODUCTION PENCIL EDIT ICON BUTTON -->
            <button
              @click="openEditJobModal(job)"
              class="p-2 rounded-xl border border-indigo-400/30 bg-indigo-500/10 text-indigo-300 hover:bg-indigo-500 hover:text-white transition-all cursor-pointer"
              title="Chỉnh sửa bài tuyển dụng"
            >
              <!-- Pencil / Edit Icon -->
              <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
              </svg>
            </button>

            <!-- Button 4: REAL PRODUCTION TRASH DELETE ICON BUTTON -->
            <button
              @click="openDeleteJobModal(job)"
              class="p-2 rounded-xl border border-rose-500/30 bg-rose-500/10 text-rose-400 hover:bg-rose-500 hover:text-white transition-all cursor-pointer"
              title="Xóa tin tuyển dụng"
            >
              <!-- Trash / Delete Icon -->
              <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
            </button>
          </div>
        </div>
      </div>

      <PaginationControls
        v-if="!showCreateForm && filteredMyJobs.length > 0"
        :page="businessJobsPage"
        :page-size="businessJobsPageSize"
        :total-items="filteredMyJobs.length"
        @update:page="businessJobsPage = $event"
      />
    </div>

    <!-- DELETE CONFIRMATION MODAL -->
    <div v-if="showDeleteConfirmModal" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-md">
      <div class="w-full max-w-md bg-slate-900 border border-rose-500/30 rounded-3xl p-6 shadow-2xl space-y-4">
        <div class="flex items-center gap-3">
          <div class="h-10 w-10 rounded-2xl bg-rose-500/20 text-rose-400 flex items-center justify-center text-xl font-bold border border-rose-500/30">
            🗑️
          </div>
          <div>
            <h3 class="text-base font-extrabold text-white">Xác Nhận Xóa Bài Đăng</h3>
            <p class="text-xs text-slate-400">Hành động này không thể hoàn tác</p>
          </div>
        </div>

        <p class="text-xs text-slate-300 leading-relaxed">
          Bạn có chắc chắn muốn xóa bài tuyển dụng <strong class="text-white">"{{ deletingJob?.title }}"</strong> không?
        </p>

        <div class="flex items-center justify-end gap-2 pt-2 border-t border-white/10">
          <button
            @click="showDeleteConfirmModal = false"
            class="px-4 py-2 border border-white/10 text-xs font-bold rounded-xl text-slate-300 hover:bg-white/5 cursor-pointer"
          >
            Hủy bỏ
          </button>
          <button
            @click="handleDeleteJob"
            :disabled="isDeletingJob"
            class="px-5 py-2 bg-rose-500 hover:bg-rose-400 text-white font-extrabold text-xs rounded-xl shadow-lg shadow-rose-500/20 cursor-pointer disabled:opacity-50"
          >
            {{ isDeletingJob ? 'Đang xóa...' : 'Xóa Bài Đăng' }}
          </button>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, computed, toRefs } from 'vue'
import PaginationControls from '~/components/common/PaginationControls.vue'

const props = defineProps<{ state: Record<string, any> }>()

const localSearchQuery = ref('')
const localStatusFilter = ref('all')

const {
  activeSection,
  jobs,
  applications,
  showCreateForm,
  jobForm,
  isCreatingJob,
  businessJobsPage,
  businessJobsPageSize,
  applicantSearchQuery,
  jobTitleLookup,
  openCreateJobModal,
  openEditJobModal,
  isEditingJob,
  handleCreateOrUpdateJob,
  toggleJobStatus,
  openDeleteJobModal,
  showDeleteConfirmModal,
  deletingJob,
  isDeletingJob,
  handleDeleteJob
} = toRefs(props.state)

const filteredMyJobs = computed(() => {
  return (jobs.value || []).filter((job: any) => {
    const title = (job.title || '').toLowerCase()
    const matchesSearch = !localSearchQuery.value || title.includes(localSearchQuery.value.toLowerCase())
    const matchesStatus = localStatusFilter.value === 'all' ||
      (localStatusFilter.value === 'open' && job.status !== 'closed') ||
      (localStatusFilter.value === 'closed' && job.status === 'closed')

    return matchesSearch && matchesStatus
  })
})

const paginatedFilteredJobs = computed(() => {
  const start = (businessJobsPage.value - 1) * businessJobsPageSize.value
  return filteredMyJobs.value.slice(start, start + businessJobsPageSize.value)
})

const getApplicantsCountForJob = (jobId: number): number => {
  return (applications.value || []).filter((app: any) => app.job_id === jobId).length
}

const filterApplicantsForJob = (jobId: number) => {
  activeSection.value = 'applicants'
  applicantSearchQuery.value = jobTitleLookup.value(jobId)
}

const formatCurrency = (value: number | string | null | undefined) => {
  const amount = Number(value || 0)
  return `${amount.toLocaleString('vi-VN')} VNĐ`
}

const displayJobTitle = (title: string | null | undefined) => {
  return (title || 'Untitled Job').replace(/\bMarketting\b/gi, 'Marketing')
}
</script>
