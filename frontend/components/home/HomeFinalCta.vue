<template>
  <section ref="sectionRef" class="relative py-16 lg:py-24 px-4 sm:px-6 lg:px-8 overflow-hidden">
    <div class="pointer-events-none absolute left-1/2 top-1/2 h-[36rem] w-[36rem] -translate-x-1/2 -translate-y-1/2 rounded-full bg-brand-500/10 blur-[150px]"></div>

    <div class="relative mx-auto max-w-7xl">
      <Transition name="reveal" appear>
        <div v-if="visible" class="mx-auto mb-12 max-w-2xl text-center">
          <p class="text-xs font-semibold uppercase tracking-[0.24em] text-brand-300">Sẵn sàng bắt đầu?</p>
          <h2 class="mt-3 text-3xl font-bold tracking-tight text-white sm:text-4xl">
            Sẵn sàng cho bước tiếp theo?
          </h2>
          <p class="mt-3 text-base text-slate-400">
            Tìm cơ hội phù hợp. Xây dựng tương lai.
          </p>
        </div>
      </Transition>

      <Transition name="reveal-scale" appear>
        <div v-if="visible" class="grid grid-cols-1 gap-5 lg:grid-cols-2">
          <!-- Student -->
          <div class="rounded-3xl border border-white/10 bg-white/[0.03] p-8 sm:p-10 transition-colors hover:border-brand-400/30">
            <span class="rounded-full border border-brand-400/25 bg-brand-400/10 px-3 py-1 text-xs font-bold text-brand-200">Dành cho sinh viên</span>
            <h3 class="mt-4 text-xl font-bold text-white sm:text-2xl">Tìm việc phù hợp với bạn</h3>
            <p class="mt-2.5 text-sm text-slate-400">
              Tạo tài khoản miễn phí, tải CV và ứng tuyển vào hàng nghìn vị trí bán thời gian &amp; thực tập.
            </p>
            <div class="mt-7 flex flex-wrap gap-3">
              <NuxtLink to="/register" class="rounded-xl bg-brand-400 px-6 py-3 text-sm font-bold text-slate-950 transition-all hover:bg-brand-300 active:scale-[0.98]">
                Tìm việc
              </NuxtLink>
              <a href="#explore-jobs" class="rounded-xl border border-white/15 px-6 py-3 text-sm font-semibold text-white transition-colors hover:bg-white/10">
                Xem tất cả việc làm
              </a>
            </div>
          </div>

          <!-- Business -->
          <div class="rounded-3xl border border-white/10 bg-white/[0.03] p-8 sm:p-10 transition-colors hover:border-emerald-400/30">
            <span class="rounded-full border border-emerald-400/25 bg-emerald-400/10 px-3 py-1 text-xs font-bold text-emerald-200">Dành cho doanh nghiệp</span>
            <h3 class="mt-4 text-xl font-bold text-white sm:text-2xl">Tuyển đúng người, đúng lúc</h3>
            <p class="mt-2.5 text-sm text-slate-400">
              Đăng tin không giới hạn, tiếp cận hồ sơ sinh viên và quản lý quy trình duyệt Offer dễ dàng.
            </p>
            <div class="mt-7 flex flex-wrap gap-3">
              <NuxtLink to="/employer-register" class="rounded-xl bg-emerald-400 px-6 py-3 text-sm font-bold text-slate-950 transition-all hover:bg-emerald-300 active:scale-[0.98]">
                Đăng tuyển
              </NuxtLink>
              <NuxtLink to="/login" class="rounded-xl border border-white/15 px-6 py-3 text-sm font-semibold text-white transition-colors hover:bg-white/10">
                Đăng nhập quản trị
              </NuxtLink>
            </div>
          </div>
        </div>
      </Transition>
    </div>
  </section>
</template>

<script setup lang="ts">
/**
 * HomeFinalCta.vue
 * Presentation-only: final dual CTA. Links reuse existing routes (/register, #explore-jobs,
 * /employer-register, /login) — no new routes, no auth logic.
 */
import { ref, onMounted, onBeforeUnmount } from 'vue'

const sectionRef = ref<HTMLElement | null>(null)
const visible = ref(false)
let observer: IntersectionObserver | null = null

onMounted(() => {
  if (!import.meta.client) return

  const prefersReducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches
  if (prefersReducedMotion || !('IntersectionObserver' in window)) {
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

.reveal-scale-enter-active {
  transition: opacity 0.55s cubic-bezier(0.16, 1, 0.3, 1), transform 0.55s cubic-bezier(0.16, 1, 0.3, 1);
  transition-delay: 0.1s;
}
.reveal-scale-enter-from {
  opacity: 0;
  transform: scale(0.98) translateY(14px);
}

@media (prefers-reduced-motion: reduce) {
  .reveal-enter-active,
  .reveal-scale-enter-active {
    transition: none !important;
  }
  .reveal-enter-from,
  .reveal-scale-enter-from {
    opacity: 1 !important;
    transform: none !important;
  }
}
</style>
