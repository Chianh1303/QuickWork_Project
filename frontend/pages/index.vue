<template>
  <div id="home" class="overflow-hidden bg-slate-950 text-white font-sans selection:bg-cyan-500 selection:text-slate-950 relative">
    
    <!-- TOP ACCENT LINE (brand-only, consistent with design system) -->
    <div class="fixed top-0 left-0 right-0 h-1 bg-brand-400/80 z-50" aria-hidden="true"></div>

    <!-- HERO SECTION (see components/home/HomeHero.vue) -->
    <HomeHero @search="handleSearch" @tag-search="applyTagSearch" />

    <!-- CATEGORIES SECTION (see components/home/HomeCategories.vue) -->
    <HomeCategories @category-select="filterByCategory" />

    <!-- SECTION: LIVE FEATURED JOBS BOARD -->
    <ExploreJobsSection
      :show-hero="false"
      :standalone="false"
      section-id="explore-jobs"
      :external-keyword="jobsSearchKeyword"
      :external-location="jobsSearchLocation"
    />

    <!-- AI SHOWCASE (see components/home/HomeAiShowcase.vue) -->
    <HomeAiShowcase />

    <!-- TRUST SECTION (see components/home/HomeTrustSection.vue) -->
    <HomeTrustSection />

    <!-- FINAL CTA (see components/home/HomeFinalCta.vue) -->
    <HomeFinalCta />

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
import { ref } from 'vue'
import ExploreJobsSection from '~/components/home/ExploreJobsSection.vue'
import HomeHero from '~/components/home/HomeHero.vue'
import HomeCategories from '~/components/home/HomeCategories.vue'
import HomeAiShowcase from '~/components/home/HomeAiShowcase.vue'
import HomeTrustSection from '~/components/home/HomeTrustSection.vue'
import HomeFinalCta from '~/components/home/HomeFinalCta.vue'

// Keyword/location that ExploreJobsSection should pick up (via its externalKeyword/externalLocation props).
const jobsSearchKeyword = ref('')
const jobsSearchLocation = ref('')

const scrollToJobs = () => {
  const element = document.getElementById('explore-jobs')
  if (element) {
    element.scrollIntoView({ behavior: 'smooth' })
  }
}

const handleSearch = (payload?: { keyword: string; location: string }) => {
  if (payload) {
    jobsSearchKeyword.value = payload.keyword
    jobsSearchLocation.value = payload.location
  }
  scrollToJobs()
}

const applyTagSearch = (tag: string) => {
  jobsSearchKeyword.value = tag
  scrollToJobs()
}

const filterByCategory = (keyword: string) => {
  jobsSearchKeyword.value = keyword
  scrollToJobs()
}

</script>
