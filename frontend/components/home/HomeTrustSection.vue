<template>
  <section ref="sectionRef" class="relative border-b border-white/10 bg-slate-950 py-16 lg:py-24 overflow-hidden">
    <div class="pointer-events-none absolute inset-0 opacity-[0.04] [background-image:linear-gradient(to_right,white_1px,transparent_1px),linear-gradient(to_bottom,white_1px,transparent_1px)] [background-size:56px_56px] [mask-image:radial-gradient(ellipse_70%_60%_at_50%_0%,black_30%,transparent_100%)]"></div>

    <div class="relative mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
      <Transition name="reveal" appear>
        <div v-if="visible" class="mx-auto max-w-2xl text-center">
          <p class="text-xs font-semibold uppercase tracking-[0.24em] text-brand-300">Bạn không đi một mình</p>
          <h2 class="mt-3 text-2xl font-bold tracking-tight text-white sm:text-3xl lg:text-[2.25rem]">
            Kết nối sinh viên với những doanh nghiệp<br class="hidden sm:block" />
            đang tìm kiếm thế hệ nhân sự tiếp theo
          </h2>
        </div>
      </Transition>

      <!-- TRUST SIGNALS: existing product claims only, no new claims invented -->
      <Transition name="reveal" appear>
        <div v-if="visible" class="mx-auto mt-10 flex max-w-3xl flex-wrap items-center justify-center gap-x-8 gap-y-3">
          <div v-for="signal in trustSignals" :key="signal" class="flex items-center gap-2 text-sm text-slate-300">
            <svg class="h-4 w-4 flex-shrink-0 text-emerald-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 13l4 4L19 7" />
            </svg>
            <span>{{ signal }}</span>
          </div>
        </div>
      </Transition>

      <!-- PARTNER MARQUEE: illustrative data, clearly labelled -->
      <Transition name="reveal" appear>
        <div v-if="visible" class="mt-16">
          <p class="text-center text-xs font-medium text-slate-500">
            Minh họa các doanh nghiệp đang tuyển dụng trên nền tảng — dữ liệu minh họa
          </p>

          <div class="relative mt-6 w-full overflow-hidden [mask-image:linear-gradient(to_right,transparent,black_8%,black_92%,transparent)]">
            <div class="flex w-max items-center gap-4" :class="prefersReducedMotion ? '' : 'animate-marquee'">
              <div
                v-for="(partner, i) in [...partners, ...partners]"
                :key="i"
                class="flex flex-shrink-0 items-center gap-3 rounded-xl border border-white/10 bg-white/[0.03] px-5 py-3.5"
              >
                <span class="flex h-8 w-8 flex-shrink-0 items-center justify-center rounded-lg bg-brand-400/10 text-xs font-bold text-brand-300">
                  {{ partner.mark }}
                </span>
                <div>
                  <p class="text-xs font-bold text-white">{{ partner.name }}</p>
                  <span class="text-[10px] font-medium text-slate-500">{{ partner.openJobs }} việc làm mở tuyển</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </Transition>
    </div>
  </section>
</template>

<script setup lang="ts">
/**
 * HomeTrustSection.vue
 * Presentation-only: social proof / trusted-companies visuals.
 * Static illustrative data + existing product claims — no API calls, no new backend claims.
 */
import { ref, onMounted, onBeforeUnmount } from 'vue'

// Reuses only claims already present in the codebase (KYB verification, wallet payout),
// consolidated here as short trust signals rather than duplicated as separate feature cards.
const trustSignals = [
  'Nhà tuyển dụng đã xác thực KYB',
  'Hồ sơ minh bạch, không tin tuyển dụng rác',
  'Ví lương minh bạch, rút tiền nhanh 24/7'
]

// Illustrative company data (unchanged names/counts from the previous marquee) — clearly labelled above as demo data.
const partners = [
  { mark: 'F', name: 'FPT Software', openJobs: 12 },
  { mark: 'V', name: 'Viettel Telecom', openJobs: 8 },
  { mark: 'VN', name: 'VNG Corporation', openJobs: 15 },
  { mark: 'S', name: 'Shopee Vietnam', openJobs: 20 },
  { mark: 'T', name: 'Techcombank', openJobs: 6 },
  { mark: 'M', name: 'Momo Wallet', openJobs: 9 },
  { mark: 'SS', name: 'Samsung R&D', openJobs: 10 },
  { mark: 'G', name: 'Grab Vietnam', openJobs: 14 }
]

const sectionRef = ref<HTMLElement | null>(null)
const visible = ref(false)
const prefersReducedMotion = ref(false)
let observer: IntersectionObserver | null = null

onMounted(() => {
  if (!import.meta.client) return

  prefersReducedMotion.value = window.matchMedia('(prefers-reduced-motion: reduce)').matches

  if (prefersReducedMotion.value || !('IntersectionObserver' in window)) {
    visible.value = true
    return
  }

  observer = new IntersectionObserver(
    (entries) => {
      if (entries[0]?.isIntersecting) {
        visible.value = true
        observer?.disconnect()
        observer = null
      }
    },
    { threshold: 0.15 }
  )
  if (sectionRef.value) observer.observe(sectionRef.value)
})

onBeforeUnmount(() => {
  observer?.disconnect()
  observer = null
})
</script>

<style scoped>
.reveal-enter-active {
  transition: opacity 0.5s cubic-bezier(0.16, 1, 0.3, 1), transform 0.5s cubic-bezier(0.16, 1, 0.3, 1);
}
.reveal-enter-from {
  opacity: 0;
  transform: translateY(14px);
}

@keyframes marquee {
  0% { transform: translateX(0%); }
  100% { transform: translateX(-50%); }
}
.animate-marquee {
  animation: marquee 34s linear infinite;
}
.animate-marquee:hover {
  animation-play-state: paused;
}

@media (prefers-reduced-motion: reduce) {
  .reveal-enter-active {
    transition: none !important;
  }
  .reveal-enter-from {
    opacity: 1 !important;
    transform: none !important;
  }
  .animate-marquee {
    animation: none !important;
  }
}
</style>
