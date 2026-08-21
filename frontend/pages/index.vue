<template>
  <div id="home" class="overflow-hidden bg-slate-950 text-white font-sans selection:bg-cyan-500 selection:text-slate-950 relative">
    
    <!-- 1. INTERACTIVE MOUSE SPOTLIGHT AURA -->
    <div
      class="pointer-events-none fixed inset-0 z-30 transition-opacity duration-300"
      :style="{
        background: `radial-gradient(600px circle at ${mousePos.x}px ${mousePos.y}px, rgba(34, 211, 238, 0.12), transparent 80%)`
      }"
    ></div>

    <!-- TOP SCROLL PROGRESS AURA GLOW -->
    <div class="fixed top-0 left-0 right-0 h-1.5 bg-gradient-to-r from-cyan-400 via-emerald-400 via-indigo-500 to-rose-500 z-50 shadow-lg shadow-cyan-500/40"></div>

    <!-- STUNNING SPLIT-SCREEN HERO SECTION WITH DYNAMIC ROTATING SHOWCASE & REALTIME FEED -->
    <section class="relative border-b border-white/10 pt-10 pb-16 lg:pt-14 lg:pb-24">
      <!-- Ambient Background Glow & Mesh Gradients -->
      <div class="absolute inset-0 bg-[radial-gradient(circle_at_12%_18%,rgba(34,211,238,0.22),transparent_35rem),radial-gradient(circle_at_88%_14%,rgba(16,185,129,0.18),transparent_38rem),linear-gradient(180deg,#020617_0%,#0f172a_50%,#062f3a_100%)]"></div>
      <div class="absolute left-1/3 top-8 h-96 w-96 -translate-x-1/2 rounded-full bg-cyan-400/10 blur-[120px] pointer-events-none animate-pulse"></div>

      <div class="relative mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div class="grid grid-cols-1 lg:grid-cols-12 gap-12 lg:gap-8 items-center">
          
          <!-- LEFT COLUMN: TEXT & LIVE NEON SEARCH CAPSULE & AUTO-ROTATING HEADLINE -->
          <div class="lg:col-span-7 space-y-6 text-left">
            <!-- Pulsing Status Badge -->
            <div class="inline-flex items-center gap-2.5 rounded-full border border-cyan-400/30 bg-cyan-400/10 px-4 py-1.5 text-xs sm:text-sm font-extrabold text-cyan-200 backdrop-blur-md shadow-lg shadow-cyan-950/40">
              <span class="relative flex h-2.5 w-2.5">
                <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-cyan-400 opacity-75"></span>
                <span class="relative inline-flex rounded-full h-2.5 w-2.5 bg-cyan-300"></span>
              </span>
              <span>Nền tảng Tuyển dụng & Việc làm Sinh viên Số 1 Việt Nam</span>
            </div>

            <!-- Headline with Auto-Flipping Keyword -->
            <h1 class="text-4xl font-extrabold leading-[1.08] tracking-tight sm:text-5xl lg:text-6xl min-h-[120px] sm:min-h-[140px]">
              Tìm Việc Bán Thời Gian & <br class="hidden sm:block" />
              Thực Tập
              <transition name="flip-fade" mode="out-in">
                <span :key="activeWordIndex" class="inline-block bg-gradient-to-r from-cyan-300 via-emerald-300 to-teal-200 bg-clip-text text-transparent underline decoration-cyan-400/40 underline-offset-8">
                  {{ titleWords[activeWordIndex] }}
                </span>
              </transition>
            </h1>

            <p class="text-base font-medium leading-relaxed text-slate-300 sm:text-lg max-w-2xl">
              QuickWork kết nối sinh viên với hàng ngàn cơ hội việc làm & thực tập uy tín từ các Doanh nghiệp đã xác thực KYB. Tích hợp AI chấm điểm CV và Ví thanh toán tức thì.
            </p>

            <!-- MAIN INTERACTIVE JOB SEARCH BAR -->
            <div class="pt-2">
              <div class="rounded-3xl border border-cyan-400/30 bg-slate-900/90 p-3 shadow-2xl shadow-cyan-950/60 backdrop-blur-xl hover:border-cyan-400/50 transition-all">
                <form @submit.prevent="handleSearch" class="grid grid-cols-1 gap-2.5 sm:grid-cols-12 items-center">
                  <!-- Input: Từ khóa -->
                  <div class="relative sm:col-span-6">
                    <div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3.5 text-slate-400">
                      <svg class="h-5 w-5 text-cyan-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                      </svg>
                    </div>
                    <input
                      v-model="heroSearchKeyword"
                      type="text"
                      class="w-full rounded-2xl border border-white/10 bg-slate-950/80 py-3 pl-11 pr-3 text-xs sm:text-sm font-semibold text-white placeholder-slate-400 transition-all focus:border-cyan-400 focus:outline-none focus:ring-2 focus:ring-cyan-400/30"
                      placeholder="Vị trí tuyển dụng, kỹ năng..."
                    />
                  </div>

                  <!-- Select: Địa điểm & Submit -->
                  <div class="relative sm:col-span-6 flex gap-2">
                    <div class="relative flex-1">
                      <div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3 text-slate-400">
                        <svg class="h-4 w-4 text-emerald-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.828 0l-4.243-4.243a8 8 0 1111.314 0z" />
                        </svg>
                      </div>
                      <select
                        v-model="heroSelectedLocation"
                        class="w-full appearance-none rounded-2xl border border-white/10 bg-slate-950/80 py-3 pl-9 pr-7 text-xs sm:text-sm font-semibold text-white transition-all focus:border-cyan-400 focus:outline-none focus:ring-2 focus:ring-cyan-400/30 cursor-pointer"
                      >
                        <option value="">Tất cả địa điểm</option>
                        <option value="Hà Nội">Hà Nội</option>
                        <option value="Hồ Chí Minh">TP. Hồ Chí Minh</option>
                        <option value="Đà Nẵng">Đà Nẵng</option>
                        <option value="Remote">Làm việc từ xa (Remote)</option>
                      </select>
                    </div>

                    <button
                      type="submit"
                      class="flex items-center justify-center gap-1.5 rounded-2xl bg-cyan-400 hover:bg-cyan-300 px-5 py-3 text-xs sm:text-sm font-extrabold text-slate-950 shadow-md shadow-cyan-500/25 transition-all hover:scale-[1.02] active:scale-95 whitespace-nowrap cursor-pointer"
                    >
                      <span>Tìm Việc Ngay</span>
                      <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M14 5l7 7m0 0l-7 7m7-7H3" />
                      </svg>
                    </button>
                  </div>
                </form>

                <!-- Popular Search Tag Pills -->
                <div class="mt-3 flex flex-wrap items-center gap-1.5 px-1 text-xs font-semibold text-slate-400">
                  <span class="text-slate-500 font-bold">Gợi ý từ khóa:</span>
                  <button
                    v-for="tag in popularTags"
                    :key="tag"
                    @click="applyTagSearch(tag)"
                    class="rounded-lg bg-white/5 border border-white/10 px-2 py-0.5 text-slate-300 hover:bg-cyan-400/20 hover:text-cyan-200 hover:border-cyan-400/40 transition-colors cursor-pointer"
                  >
                    #{{ tag }}
                  </button>
                </div>
              </div>
            </div>

            <!-- Quick Stats Row with Animated Increment Numbers -->
            <div class="flex flex-wrap items-center gap-6 pt-2 text-xs sm:text-sm font-bold text-slate-300">
              <div class="flex items-center gap-2">
                <span class="text-cyan-400 text-base">⚡</span>
                <span><strong>{{ countJobs }}+</strong> Việc làm mở mới</span>
              </div>
              <div class="flex items-center gap-2">
                <span class="text-emerald-400 text-base">🛡️</span>
                <span><strong>{{ countBusinesses }}+</strong> Doanh nghiệp KYB</span>
              </div>
              <div class="flex items-center gap-2">
                <span class="text-amber-400 text-base">🎓</span>
                <span><strong>{{ countStudents }}+</strong> Sinh viên nhận việc</span>
              </div>
            </div>
          </div>

          <!-- RIGHT COLUMN: INTERACTIVE PRODUCT SHOWCASE WIDGET CARD WITH AUTO-SLIDING JOBS -->
          <div class="lg:col-span-5 relative">
            <div class="absolute -inset-4 rounded-3xl bg-gradient-to-tr from-cyan-400/20 via-emerald-400/15 to-transparent blur-3xl pointer-events-none"></div>

            <!-- FLOATING BADGE 1: AI GEMINI CV SCORE -->
            <div class="absolute -top-6 -left-6 z-20 hidden sm:flex items-center gap-3 rounded-2xl border border-cyan-400/30 bg-slate-900/95 px-4 py-3 shadow-2xl backdrop-blur-md animate-bounce duration-[3000ms]">
              <div class="flex h-10 w-10 items-center justify-center rounded-xl bg-cyan-400/20 text-xl text-cyan-300 border border-cyan-400/30">
                🤖
              </div>
              <div>
                <p class="text-[10px] font-black uppercase tracking-wider text-cyan-300">AI Gemini Evaluation</p>
                <p class="text-xs font-extrabold text-white">Chấm Điểm CV: <span class="text-emerald-400">95/100 ATS</span></p>
              </div>
            </div>

            <!-- FLOATING BADGE 2: INSTANT SALARY PAYOUT -->
            <div class="absolute -bottom-6 -right-6 z-20 hidden sm:flex items-center gap-3 rounded-2xl border border-emerald-400/30 bg-slate-900/95 px-4 py-3 shadow-2xl backdrop-blur-md">
              <div class="flex h-10 w-10 items-center justify-center rounded-xl bg-emerald-400/20 text-xl text-emerald-300 border border-emerald-400/30">
                💳
              </div>
              <div>
                <p class="text-[10px] font-black uppercase tracking-wider text-emerald-300">Ví QuickWork</p>
                <p class="text-xs font-extrabold text-white">Nhận Lương: <span class="text-emerald-400">+1.200.000 VNĐ</span></p>
              </div>
            </div>

            <!-- MAIN CENTRAL DYNAMIC SHOWCASE TICKET -->
            <div class="relative overflow-hidden rounded-3xl border border-cyan-400/25 bg-gradient-to-b from-slate-900/95 to-slate-950 p-6 shadow-2xl shadow-slate-950/80 backdrop-blur-xl min-h-[320px] flex flex-col justify-between">
              
              <transition name="slide-up" mode="out-in">
                <div :key="activeShowcaseIndex" class="space-y-4">
                  <!-- Ticket Header -->
                  <div class="flex items-center justify-between border-b border-white/10 pb-4">
                    <div class="flex items-center gap-3">
                      <div class="h-12 w-12 rounded-xl bg-gradient-to-tr from-cyan-500 to-indigo-600 flex items-center justify-center font-extrabold text-slate-950 text-xl shadow-md">
                        {{ currentShowcaseJob.logo }}
                      </div>
                      <div>
                        <h3 class="text-sm font-extrabold text-white">{{ currentShowcaseJob.company }}</h3>
                        <p class="text-xs font-semibold text-slate-400">{{ currentShowcaseJob.industry }}</p>
                      </div>
                    </div>
                    <span class="inline-flex items-center gap-1.5 rounded-full border border-emerald-400/30 bg-emerald-400/10 px-3 py-1 text-xs font-extrabold text-emerald-300">
                      <span class="h-1.5 w-1.5 rounded-full bg-emerald-400 animate-ping"></span>
                      Đã xác thực KYB
                    </span>
                  </div>

                  <!-- Ticket Body Job Highlight -->
                  <div class="space-y-2.5">
                    <span class="rounded-full bg-cyan-400/10 border border-cyan-400/20 px-3 py-1 text-[11px] font-black text-cyan-300">
                      {{ currentShowcaseJob.badge }}
                    </span>
                    <h4 class="text-lg font-extrabold text-white leading-snug">
                      {{ currentShowcaseJob.title }}
                    </h4>
                    <p class="text-xs text-slate-400 leading-relaxed">
                      {{ currentShowcaseJob.desc }}
                    </p>

                    <div class="pt-2 flex items-center justify-between">
                      <div>
                        <p class="text-[10px] font-black uppercase tracking-wider text-slate-400">Mức lương đề xuất</p>
                        <p class="text-lg font-extrabold text-emerald-300 mt-0.5">{{ currentShowcaseJob.salary }}</p>
                      </div>
                      <a href="#explore-jobs" class="rounded-xl bg-cyan-400 hover:bg-cyan-300 px-4 py-2 text-xs font-extrabold text-slate-950 shadow-md transition-all cursor-pointer">
                        Ứng Tuyển Nhanh
                      </a>
                    </div>
                  </div>
                </div>
              </transition>

              <!-- Carousel Dots Indicator & Live Ticker Bar -->
              <div class="pt-4 border-t border-white/10 space-y-3">
                <!-- Dots Indicator -->
                <div class="flex items-center justify-center gap-2">
                  <button
                    v-for="(_, idx) in showcaseJobs"
                    :key="idx"
                    @click="activeShowcaseIndex = idx"
                    :class="[
                      activeShowcaseIndex === idx ? 'w-6 bg-cyan-400' : 'w-2 bg-white/20 hover:bg-white/40',
                      'h-2 rounded-full transition-all duration-300 cursor-pointer'
                    ]"
                    :title="`Xem công việc ${idx + 1}`"
                  ></button>
                </div>

                <!-- LIVE FLOATING ACTIVITY TICKER -->
                <div class="rounded-xl border border-cyan-400/20 bg-slate-950/70 p-2.5 text-center overflow-hidden">
                  <transition name="fade" mode="out-in">
                    <p :key="activeActivityIndex" class="text-[11px] font-bold text-slate-300 truncate">
                      <span>{{ currentActivity.icon }}</span>
                      <span class="ml-1 text-slate-200">{{ currentActivity.text }}</span>
                    </p>
                  </transition>
                </div>
              </div>

            </div>
          </div>

        </div>
      </div>
    </section>

    <!-- SECTION: TOP CATEGORIES EXPLORER (VIBRANT MULTI-COLOR ACCENTS) -->
    <section id="categories" class="py-16 px-4 sm:px-6 lg:px-8 border-b border-white/10 bg-slate-900/40">
      <div class="mx-auto max-w-7xl">
        <div class="flex flex-col sm:flex-row items-start sm:items-end justify-between gap-4 mb-10">
          <div>
            <p class="text-xs font-black uppercase tracking-[0.24em] text-cyan-300">Khám Phá Xu Hướng</p>
            <h2 class="mt-2 text-2xl sm:text-3xl font-extrabold text-white tracking-tight">
              Ngành Nghề Tuyển Dụng Sinh Viên Phổ Biến
            </h2>
          </div>
          <a href="#explore-jobs" class="text-xs font-bold text-cyan-300 hover:text-cyan-200 flex items-center gap-1">
            <span>Xem tất cả danh mục</span>
            <span>→</span>
          </a>
        </div>

        <div class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-6 gap-4">
          <div
            v-for="cat in popularCategories"
            :key="cat.name"
            @click="filterByCategory(cat.keyword)"
            :class="[
              cat.cardStyle,
              'group relative overflow-hidden flex flex-col justify-end p-4 h-48 sm:h-52 rounded-2xl border shadow-xl shadow-slate-950/40 transition-all duration-500 hover:-translate-y-2 cursor-pointer'
            ]"
          >
            <!-- Background Cover Image with Zoom Effect -->
            <img
              :src="cat.image"
              :alt="cat.name"
              class="absolute inset-0 h-full w-full object-cover transition-transform duration-700 ease-out group-hover:scale-125"
            />

            <!-- Gradient Contrast Overlay -->
            <div class="absolute inset-0 bg-gradient-to-t from-slate-950 via-slate-950/70 to-transparent transition-opacity group-hover:opacity-90"></div>

            <!-- Card Content Overlay -->
            <div class="relative z-10 space-y-1.5 text-left">
              <h4 class="text-xs sm:text-sm font-black text-white group-hover:text-cyan-300 transition-colors leading-tight">
                {{ cat.name }}
              </h4>
              <p class="text-[11px] font-extrabold text-cyan-200 bg-slate-950/80 w-fit px-2.5 py-0.5 rounded-md border border-white/10 backdrop-blur-md">
                {{ cat.jobsCount }} việc làm
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- SECTION: LIVE FEATURED JOBS BOARD -->
    <ExploreJobsSection :show-hero="false" :standalone="false" section-id="explore-jobs" />

    <!-- SECTION: WHY QUICKWORK (CORE VALUES & INTERACTIVE AI GEMINI SCANNER SIMULATOR) -->
    <section id="features" class="py-16 lg:py-24 px-4 sm:px-6 lg:px-8 border-y border-white/10 bg-slate-900/50 relative overflow-hidden">
      <div class="mx-auto max-w-7xl">
        <div class="text-center max-w-3xl mx-auto mb-12">
          <p class="text-xs font-black uppercase tracking-[0.24em] text-cyan-300">Tính Năng Độc Quyền</p>
          <h2 class="mt-3 text-3xl sm:text-4xl font-extrabold text-white tracking-tight">
            Giải Pháp Công Nghệ Tối Ưu Cho Sinh Viên
          </h2>
          <p class="mt-3 text-sm sm:text-base font-medium text-slate-400">
            Hệ thống sinh thái tích hợp từ AI chấm điểm CV, Chấm công điểm danh ca làm cho đến Ví nhận lương tự động.
          </p>
        </div>

        <!-- INTERACTIVE LIVE AI SCANNER DEMO WIDGET -->
        <div class="mb-14 mx-auto max-w-3xl rounded-3xl border border-cyan-400/30 bg-slate-950/90 p-6 shadow-2xl shadow-cyan-950/50 backdrop-blur-xl">
          <div class="flex flex-col sm:flex-row items-center justify-between gap-4 border-b border-white/10 pb-4">
            <div class="flex items-center gap-3">
              <span class="flex h-10 w-10 items-center justify-center rounded-xl bg-cyan-400/20 text-2xl text-cyan-300">🤖</span>
              <div>
                <h3 class="text-base font-extrabold text-white">Mô Phỏng Trực Tiếp AI Gemini CV Scanner</h3>
                <p class="text-xs text-slate-400">Thử nghiệm công nghệ quét hồ sơ tự động</p>
              </div>
            </div>
            <button
              @click="runSimulatedAiScan"
              :disabled="isAnalyzing"
              class="w-full sm:w-auto rounded-xl bg-gradient-to-r from-cyan-400 to-emerald-400 px-5 py-2.5 text-xs font-extrabold text-slate-950 shadow-md hover:brightness-110 disabled:opacity-50 transition-all cursor-pointer"
            >
              {{ isAnalyzing ? '⚡ AI Đang Phân Tích...' : '⚡ Thử Phân Tích CV Mẫu' }}
            </button>
          </div>

          <!-- Laser Scan Progress Animation Bar -->
          <div v-if="isAnalyzing" class="my-6 space-y-2">
            <div class="flex justify-between text-xs font-bold text-cyan-300">
              <span>Đang quét từ khóa & chỉ số ATS...</span>
              <span>{{ scanProgress }}%</span>
            </div>
            <div class="h-2 w-full overflow-hidden rounded-full bg-slate-800 relative">
              <div
                class="h-full bg-gradient-to-r from-cyan-400 via-emerald-400 to-indigo-500 transition-all duration-300"
                :style="{ width: `${scanProgress}%` }"
              ></div>
            </div>
          </div>

          <!-- Simulated Result Output -->
          <div v-if="aiScanResult" class="mt-5 p-4 rounded-2xl border border-emerald-400/30 bg-emerald-950/30 space-y-2 animate-in fade-in zoom-in-95">
            <div class="flex items-center justify-between">
              <span class="text-xs font-extrabold uppercase text-emerald-300">KẾT QUẢ ĐÁNH GIÁ CHUẨN ATS</span>
              <span class="rounded-full bg-emerald-400/20 px-3 py-1 text-xs font-black text-emerald-200">Điểm ATS: {{ aiScanResult.score }}/100</span>
            </div>
            <p class="text-xs text-slate-200 font-medium"><strong>Kỹ năng nổi bật phát hiện:</strong> {{ aiScanResult.skills.join(', ') }}</p>
            <p class="text-xs text-cyan-300 italic">"{{ aiScanResult.feedback }}"</p>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <!-- Card 1: AI Gemini CV Rating (Neon Cyan/Purple) -->
          <div class="rounded-3xl border border-cyan-400/30 bg-gradient-to-b from-indigo-950/70 via-slate-900 to-slate-950 p-7 shadow-xl shadow-cyan-950/30 relative overflow-hidden group hover:border-cyan-400/60 transition-all hover:-translate-y-1">
            <div class="h-12 w-12 rounded-2xl bg-cyan-400/20 border border-cyan-400/30 flex items-center justify-center text-2xl mb-6 text-cyan-300 shadow-md">
              🤖
            </div>
            <span class="rounded-full bg-cyan-400/10 px-3 py-1 text-[11px] font-black text-cyan-300 border border-cyan-400/20">Công Nghệ AI Gemini</span>
            <h3 class="mt-4 text-xl font-extrabold text-white">Phân Tích & Chấm Điểm CV</h3>
            <p class="mt-2 text-xs font-medium text-slate-300/90 leading-relaxed">
              Trí tuệ nhân tạo Gemini quét trực tiếp file CV PDF của bạn, tính toán chỉ số phù hợp ATS và đưa ra lời khuyên cải thiện hồ sơ chi tiết.
            </p>
            <div class="mt-6 pt-4 border-t border-white/10 flex items-center justify-between text-xs font-bold text-cyan-300">
              <NuxtLink to="/student/dashboard" class="hover:underline">Đánh giá CV ngay →</NuxtLink>
            </div>
          </div>

          <!-- Card 2: Verified KYB Enterprises (Emerald Mint) -->
          <div class="rounded-3xl border border-emerald-400/30 bg-gradient-to-b from-emerald-950/70 via-slate-900 to-slate-950 p-7 shadow-xl shadow-emerald-950/30 relative overflow-hidden group hover:border-emerald-400/60 transition-all hover:-translate-y-1">
            <div class="h-12 w-12 rounded-2xl bg-emerald-400/20 border border-emerald-400/30 flex items-center justify-center text-2xl mb-6 text-emerald-300 shadow-md">
              🛡️
            </div>
            <span class="rounded-full bg-emerald-400/10 px-3 py-1 text-[11px] font-black text-emerald-300 border border-emerald-400/20">An Toàn 100%</span>
            <h3 class="mt-4 text-xl font-extrabold text-white">Doanh Nghiệp Đã Xác Thực KYB</h3>
            <p class="mt-2 text-xs font-medium text-slate-300/90 leading-relaxed">
              100% Nhà tuyển dụng trên QuickWork đều trải qua quy trình xác minh hồ sơ doanh nghiệp nghiêm ngặt. Cam kết không có tin tuyển dụng rác hay lừa đảo.
            </p>
            <div class="mt-6 pt-4 border-t border-white/10 flex items-center justify-between text-xs font-bold text-emerald-300">
              <span>Hồ sơ minh bạch</span>
            </div>
          </div>

          <!-- Card 3: Wallet & Timekeeping (Amber Orange) -->
          <div class="rounded-3xl border border-amber-400/30 bg-gradient-to-b from-amber-950/70 via-slate-900 to-slate-950 p-7 shadow-xl shadow-amber-950/30 relative overflow-hidden group hover:border-amber-400/60 transition-all hover:-translate-y-1">
            <div class="h-12 w-12 rounded-2xl bg-amber-400/20 border border-amber-400/30 flex items-center justify-center text-2xl mb-6 text-amber-300 shadow-md">
              💳
            </div>
            <span class="rounded-full bg-amber-400/10 px-3 py-1 text-[11px] font-black text-amber-300 border border-amber-400/20">Thanh Toán Linh Hoạt</span>
            <h3 class="mt-4 text-xl font-extrabold text-white">Ví Lương & Chấm Công 1-Chạm</h3>
            <p class="mt-2 text-xs font-medium text-slate-300/90 leading-relaxed">
              Theo dõi thời gian làm việc hàng ngày bằng Check-in/Check-out. Lương được quyết toán minh bạch trực tiếp về Ví QuickWork cá nhân.
            </p>
            <div class="mt-6 pt-4 border-t border-white/10 flex items-center justify-between text-xs font-bold text-amber-300">
              <span>Rút tiền nhanh 24/7</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- SECTION: INFINITE SMOOTH HIRING PARTNERS LOGO MARQUEE -->
    <section class="py-16 px-4 sm:px-6 lg:px-8 border-b border-white/10 bg-slate-950 overflow-hidden">
      <div class="mx-auto max-w-7xl">
        <p class="text-center text-xs font-black uppercase tracking-[0.24em] text-slate-500 mb-8">
          Đối Tác Tuyển Dụng Hàng Đầu Trên QuickWork
        </p>

        <!-- Marquee Track Container -->
        <div class="relative w-full overflow-hidden">
          <div class="flex w-max items-center gap-6 animate-marquee">
            <div
              v-for="(partner, i) in [...topPartners, ...topPartners]"
              :key="i"
              class="flex items-center gap-3 rounded-2xl border border-white/10 bg-slate-900/80 px-6 py-4 hover:border-cyan-400/50 transition-all flex-shrink-0"
            >
              <span class="text-2xl">{{ partner.logo }}</span>
              <div>
                <p class="text-xs font-extrabold text-white">{{ partner.name }}</p>
                <span class="text-[10px] font-semibold text-cyan-300">{{ partner.openJobs }} việc làm mở tuyển</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- DUAL CTA BANNER (VIBRANT HIGH-ENERGY DUAL CTA) -->
    <section class="py-16 lg:py-20 px-4 sm:px-6 lg:px-8">
      <div class="mx-auto max-w-7xl grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- Student CTA (Vibrant Cyan-Indigo Gradient) -->
        <div class="rounded-3xl border border-cyan-400/40 bg-gradient-to-br from-cyan-950/80 via-slate-900 to-indigo-950 p-8 sm:p-10 shadow-2xl relative overflow-hidden flex flex-col justify-between hover:border-cyan-400/60 transition-all">
          <div>
            <span class="rounded-full bg-cyan-400/20 border border-cyan-400/30 px-3 py-1 text-xs font-black text-cyan-200">Dành Cho Sinh Viên</span>
            <h3 class="mt-4 text-2xl sm:text-3xl font-extrabold text-white">Sẵn Sàng Bắt Đầu Hành Trình Sự Nghiệp?</h3>
            <p class="mt-3 text-xs sm:text-sm font-medium text-slate-300 leading-relaxed">
              Tạo tài khoản sinh viên hoàn toàn miễn phí, tải lên CV PDF của bạn và ứng tuyển ngay vào hàng nghìn vị trí bán thời gian & thực tập phù hợp.
            </p>
          </div>
          <div class="mt-8 flex flex-wrap gap-3">
            <NuxtLink to="/register" class="rounded-xl bg-cyan-400 px-6 py-3 text-xs font-extrabold text-slate-950 shadow-lg shadow-cyan-500/30 hover:bg-cyan-300 transition-all cursor-pointer">
              Đăng Ký Sinh Viên Ngay
            </NuxtLink>
            <a href="#explore-jobs" class="rounded-xl border border-white/15 bg-white/10 px-6 py-3 text-xs font-bold text-white hover:bg-white/20 transition-all cursor-pointer">
              Xem Tất Cả Việc Làm
            </a>
          </div>
        </div>

        <!-- Enterprise CTA (Vibrant Emerald-Teal Gradient) -->
        <div class="rounded-3xl border border-emerald-400/40 bg-gradient-to-br from-emerald-950/80 via-slate-900 to-teal-950 p-8 sm:p-10 shadow-2xl relative overflow-hidden flex flex-col justify-between hover:border-emerald-400/60 transition-all">
          <div>
            <span class="rounded-full bg-emerald-400/20 border border-emerald-400/30 px-3 py-1 text-xs font-black text-emerald-200">Dành Cho Nhà Tuyển Dụng</span>
            <h3 class="mt-4 text-2xl sm:text-3xl font-extrabold text-white">Tuyển Dụng Nhân Sự Sinh Viên Nhanh Chóng</h3>
            <p class="mt-3 text-xs sm:text-sm font-medium text-slate-300 leading-relaxed">
              Đăng tin tuyển dụng không giới hạn, tiếp cận hàng ngàn hồ sơ ứng viên sinh viên năng động và quản lý quy trình duyệt Offer dễ dàng.
            </p>
          </div>
          <div class="mt-8 flex flex-wrap gap-3">
            <NuxtLink to="/employer-register" class="rounded-xl bg-emerald-400 px-6 py-3 text-xs font-extrabold text-slate-950 shadow-lg shadow-emerald-500/30 hover:bg-emerald-300 transition-all cursor-pointer">
              Đăng Ký Doanh Nghiệp
            </NuxtLink>
            <NuxtLink to="/login" class="rounded-xl border border-white/15 bg-white/10 px-6 py-3 text-xs font-bold text-white hover:bg-white/20 transition-all cursor-pointer">
              Đăng Nhập Quản Trị
            </NuxtLink>
          </div>
        </div>
      </div>
    </section>

    <!-- FOOTER -->
    <footer id="contact" class="px-4 py-12 sm:px-6 lg:px-8 border-t border-white/10 bg-slate-950">
      <div class="mx-auto grid max-w-7xl gap-8 rounded-3xl border border-white/10 bg-slate-900/80 p-8 md:grid-cols-[1.2fr_0.8fr_0.8fr_1fr]">
        <div>
          <div class="flex items-center gap-3">
            <span class="flex h-10 w-10 items-center justify-center rounded-xl bg-gradient-to-br from-cyan-400 to-emerald-400 font-extrabold text-slate-950 shadow-md">Q</span>
            <span class="text-xl font-extrabold text-white">QuickWork</span>
          </div>
          <p class="mt-4 max-w-sm text-xs font-medium leading-relaxed text-slate-400">
            Nền tảng tuyển dụng hiện đại dành cho việc làm sinh viên, thực tập, AI phân tích CV, trao đổi trực tiếp, chấm công ca làm và ví tiền tiện lợi.
          </p>
        </div>

        <div>
          <h3 class="text-xs font-extrabold uppercase tracking-wider text-white">Liên kết nhanh</h3>
          <div class="mt-4 space-y-2 text-xs font-medium text-slate-400">
            <p><a href="#home" class="hover:text-cyan-300 transition-colors">Trang chủ</a></p>
            <p><a href="#explore-jobs" class="hover:text-cyan-300 transition-colors">Khám phá việc làm</a></p>
            <p><a href="#categories" class="hover:text-cyan-300 transition-colors">Danh mục ngành nghề</a></p>
            <p><a href="#features" class="hover:text-cyan-300 transition-colors">Tính năng hệ thống</a></p>
          </div>
        </div>

        <div>
          <h3 class="text-xs font-extrabold uppercase tracking-wider text-white">Dành cho người dùng</h3>
          <div class="mt-4 space-y-2 text-xs font-medium text-slate-400">
            <p><NuxtLink to="/register" class="hover:text-cyan-300 transition-colors">Tài khoản Sinh viên</NuxtLink></p>
            <p><NuxtLink to="/employer-register" class="hover:text-cyan-300 transition-colors">Doanh nghiệp tuyển dụng</NuxtLink></p>
            <p><NuxtLink to="/login" class="hover:text-cyan-300 transition-colors">Đăng nhập tài khoản</NuxtLink></p>
          </div>
        </div>

        <div>
          <h3 class="text-xs font-extrabold uppercase tracking-wider text-white">Liên hệ hỗ trợ</h3>
          <div class="mt-4 space-y-2 text-xs font-medium text-slate-400">
            <p>📧 support@quickwork.vn</p>
            <p>📞 0123 456 789</p>
            <p>📍 TP. Hồ Chí Minh, Việt Nam</p>
          </div>
          <div class="mt-5 flex gap-2">
            <span v-for="social in ['Facebook', 'LinkedIn', 'GitHub']" :key="social" class="rounded-lg border border-white/10 bg-white/5 px-2.5 py-1 text-[11px] font-bold text-slate-300 hover:border-cyan-400/40 cursor-pointer transition-colors">{{ social }}</span>
          </div>
        </div>
      </div>

      <p class="mx-auto mt-8 max-w-7xl text-center text-xs font-medium text-slate-500">© 2026 QuickWork Platform. Bản quyền thuộc về QuickWork.</p>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import ExploreJobsSection from '~/components/home/ExploreJobsSection.vue'

import catIt from '~/assets/images/categories/cat_it.jpg'
import catMarketing from '~/assets/images/categories/cat_marketing.jpg'
import catDesign from '~/assets/images/categories/cat_design.jpg'
import catFb from '~/assets/images/categories/cat_fb.jpg'
import catEducation from '~/assets/images/categories/cat_education.jpg'
import catSales from '~/assets/images/categories/cat_sales.jpg'

const heroSearchKeyword = ref('')
const heroSelectedLocation = ref('')

// Mouse Position tracking for Interactive Spotlight Aura
const mousePos = ref({ x: 400, y: 300 })
const handleMouseMove = (e: MouseEvent) => {
  mousePos.value = { x: e.clientX, y: e.clientY }
}

// Animated Counter Numbers
const countJobs = ref(0)
const countBusinesses = ref(0)
const countStudents = ref(0)

const animateCounters = () => {
  const duration = 1500
  const steps = 30
  const stepTime = duration / steps
  let currentStep = 0

  const interval = setInterval(() => {
    currentStep++
    countJobs.value = Math.min(300, Math.floor((300 / steps) * currentStep))
    countBusinesses.value = Math.min(150, Math.floor((150 / steps) * currentStep))
    countStudents.value = Math.min(500, Math.floor((500 / steps) * currentStep))

    if (currentStep >= steps) {
      clearInterval(interval)
    }
  }, stepTime)
}

// Interactive AI Gemini Live CV Scanner Simulation
const isAnalyzing = ref(false)
const scanProgress = ref(0)
const aiScanResult = ref<any | null>(null)

const runSimulatedAiScan = () => {
  if (isAnalyzing.value) return
  isAnalyzing.value = true
  scanProgress.value = 0
  aiScanResult.value = null

  const progressInterval = setInterval(() => {
    scanProgress.value += 10
    if (scanProgress.value >= 100) {
      clearInterval(progressInterval)
      isAnalyzing.value = false
      aiScanResult.value = {
        score: 95,
        skills: ['Backend Developer', 'Go (Golang)', 'Node.js', 'PostgreSQL', 'Docker'],
        feedback: 'Hồ sơ CV mẫu có cấu trúc xuất sắc! Độ tương thích 98% với các vị trí Thực tập sinh Backend tại FPT Software & VNG.'
      }
    }
  }, 200)
}

// Title Auto-Flipping Keyword State
const titleWords = ['Nhanh Chóng', 'Linh Hoạt', 'Lương Cao', 'AI Gemini']
const activeWordIndex = ref(0)
let wordTimer: any = null

// Auto-Rotating Showcase Jobs
const showcaseJobs = [
  {
    logo: 'F',
    company: 'FPT Software',
    industry: 'Công nghệ & Phần mềm',
    badge: 'Việc Làm Bán Thời Gian',
    title: 'Thực Tập Sinh Backend Developer (Go / Node.js)',
    desc: 'Phù hợp sinh viên · Thời gian ca làm linh hoạt · Nhận việc ngay sau phỏng vấn.',
    salary: '9.500.000 VNĐ / tháng'
  },
  {
    logo: 'S',
    company: 'Shopee Vietnam',
    industry: 'Thương mại điện tử',
    badge: 'Thực Tập Sinh',
    title: 'Marketing & Content Campaign Assistant',
    desc: 'Môi trường làm việc trẻ trung · Tự do sáng tạo · Hỗ trợ con đường sự nghiệp.',
    salary: '8.000.000 VNĐ / tháng'
  },
  {
    logo: 'V',
    company: 'VNG Corporation',
    industry: 'Game & Media',
    badge: 'Bán Thời Gian',
    title: 'Graphic & UI/UX Designer Part-Time',
    desc: 'Thiết kế ấn phẩm truyền thông · Được đào tạo bài bản bởi Senior Mentors.',
    salary: '10.000.000 VNĐ / tháng'
  },
  {
    logo: 'H',
    company: 'Highlands Coffee',
    industry: 'F&B & Phục vụ',
    badge: 'Theo Ca Linh Hoạt',
    title: 'Nhân Viên Barista & Thu Ngân Ca Linh Hoạt',
    desc: 'Xoay ca theo lịch học · Đào tạo pha chế miễn phí · Thưởng doanh số hàng tháng.',
    salary: '35.000 VNĐ / giờ'
  }
]
const activeShowcaseIndex = ref(0)
let showcaseTimer: any = null

const currentShowcaseJob = computed(() => showcaseJobs[activeShowcaseIndex.value])

// Live Floating Activity Feed Ticker
const liveActivities = [
  { icon: '🟢', text: 'Minh Anh vừa ứng tuyển FPT Software (12 giây trước)' },
  { icon: '🚀', text: 'Quốc Bảo vừa nhận Offer từ VNG Corporation (1 phút trước)' },
  { icon: '💳', text: 'Thanh Hà vừa nhận lương 1.200.000 VNĐ về Ví QuickWork (3 phút trước)' },
  { icon: '⚡', text: 'Hoàng Nam vừa được AI Gemini chấm CV 95/100 ATS (5 phút trước)' }
]
const activeActivityIndex = ref(0)
let activityTimer: any = null

const currentActivity = computed(() => liveActivities[activeActivityIndex.value])

const popularTags = ['Backend', 'Marketing', 'GraphicDesign', 'GiaSư', 'Frontend', 'Sales', 'Remote']

const popularCategories = [
  { icon: '💻', name: 'CNTT & Phần mềm', jobsCount: 45, keyword: 'it', image: catIt, cardStyle: 'border-indigo-500/40 hover:border-indigo-400 hover:shadow-indigo-500/30' },
  { icon: '📈', name: 'Marketing & Media', jobsCount: 38, keyword: 'marketing', image: catMarketing, cardStyle: 'border-rose-500/40 hover:border-rose-400 hover:shadow-rose-500/30' },
  { icon: '🎨', name: 'Thiết Kế Đồ Họa', jobsCount: 24, keyword: 'design', image: catDesign, cardStyle: 'border-amber-500/40 hover:border-amber-400 hover:shadow-amber-500/30' },
  { icon: '☕', name: 'F&B & Phục Vụ', jobsCount: 52, keyword: 'f&b', image: catFb, cardStyle: 'border-emerald-500/40 hover:border-emerald-400 hover:shadow-emerald-500/30' },
  { icon: '📚', name: 'Gia Sư & Giáo Dục', jobsCount: 19, keyword: 'giasu', image: catEducation, cardStyle: 'border-sky-500/40 hover:border-sky-400 hover:shadow-sky-500/30' },
  { icon: '🛍️', name: 'Bán Hàng & Sales', jobsCount: 31, keyword: 'sales', image: catSales, cardStyle: 'border-purple-500/40 hover:border-purple-400 hover:shadow-purple-500/30' }
]

const topPartners = [
  { logo: '🏢', name: 'FPT Software', openJobs: 12 },
  { logo: '🌐', name: 'Viettel Telecom', openJobs: 8 },
  { logo: '🚀', name: 'VNG Corporation', openJobs: 15 },
  { logo: '🛍️', name: 'Shopee Vietnam', openJobs: 20 },
  { logo: '💻', name: 'Techcombank', openJobs: 6 },
  { logo: '⚡', name: 'Momo Wallet', openJobs: 9 },
  { logo: '📱', name: 'Samsung R&D', openJobs: 10 },
  { logo: '🚕', name: 'Grab Vietnam', openJobs: 14 }
]

const handleSearch = () => {
  const element = document.getElementById('explore-jobs')
  if (element) {
    element.scrollIntoView({ behavior: 'smooth' })
  }
}

const applyTagSearch = (tag: string) => {
  heroSearchKeyword.value = tag
  handleSearch()
}

const filterByCategory = (keyword: string) => {
  heroSearchKeyword.value = keyword
  handleSearch()
}

// Lifecycle Timers for Dynamic Animations & Mouse Events
onMounted(() => {
  if (import.meta.client) {
    window.addEventListener('mousemove', handleMouseMove)
  }

  animateCounters()

  wordTimer = setInterval(() => {
    activeWordIndex.value = (activeWordIndex.value + 1) % titleWords.length
  }, 2400)

  showcaseTimer = setInterval(() => {
    activeShowcaseIndex.value = (activeShowcaseIndex.value + 1) % showcaseJobs.length
  }, 3500)

  activityTimer = setInterval(() => {
    activeActivityIndex.value = (activeActivityIndex.value + 1) % liveActivities.length
  }, 4000)
})

onUnmounted(() => {
  if (import.meta.client) {
    window.removeEventListener('mousemove', handleMouseMove)
  }
  if (wordTimer) clearInterval(wordTimer)
  if (showcaseTimer) clearInterval(showcaseTimer)
  if (activityTimer) clearInterval(activityTimer)
})
</script>

<style scoped>
/* Vue Transitions for Dynamic Micro-Animations */
.flip-fade-enter-active,
.flip-fade-leave-active {
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}
.flip-fade-enter-from {
  opacity: 0;
  transform: translateY(12px) rotateX(-20deg);
}
.flip-fade-leave-to {
  opacity: 0;
  transform: translateY(-12px) rotateX(20deg);
}

.slide-up-enter-active,
.slide-up-leave-active {
  transition: all 0.35s cubic-bezier(0.16, 1, 0.3, 1);
}
.slide-up-enter-from {
  opacity: 0;
  transform: translateY(16px);
}
.slide-up-leave-to {
  opacity: 0;
  transform: translateY(-16px);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* Infinite Smooth Logo Marquee Animation */
@keyframes marquee {
  0% {
    transform: translateX(0%);
  }
  100% {
    transform: translateX(-50%);
  }
}
.animate-marquee {
  animation: marquee 25s linear infinite;
}
.animate-marquee:hover {
  animation-play-state: paused;
}
</style>
