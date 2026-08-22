<template>
  <section id="categories" ref="sectionRef" class="relative border-b border-white/10 bg-slate-950 py-16 lg:py-24">
    <!-- Same visual world as Hero: neutral base + very light grid, no separate background effect -->
    <div class="pointer-events-none absolute inset-0 opacity-[0.04] [background-image:linear-gradient(to_right,white_1px,transparent_1px),linear-gradient(to_bottom,white_1px,transparent_1px)] [background-size:56px_56px] [mask-image:radial-gradient(ellipse_70%_60%_at_50%_0%,black_30%,transparent_100%)]"></div>

    <div class="relative mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
      <!-- SECTION INTRO -->
      <Transition name="reveal" appear>
        <div v-if="visible" class="mx-auto mb-12 max-w-2xl text-center">
          <p class="text-xs font-semibold uppercase tracking-[0.24em] text-brand-300">Khám phá cơ hội</p>
          <h2 class="mt-3 text-2xl font-bold tracking-tight text-white sm:text-3xl lg:text-[2.25rem]">
            Tìm công việc phù hợp với<br class="hidden sm:block" />
            hướng đi của bạn
          </h2>
          <p class="mt-3 text-sm text-slate-400 sm:text-base">
            Chọn một lĩnh vực để bắt đầu khám phá.
          </p>
        </div>
      </Transition>

      <!-- CATEGORY CARDS -->
      <div class="grid grid-cols-2 gap-3 sm:gap-4 lg:grid-cols-3">
        <Transition
          v-for="(cat, i) in categories"
          :key="cat.name"
          name="reveal"
          appear
        >
          <button
            v-if="visible"
            type="button"
            :style="{ transitionDelay: `${i * 70}ms` }"
            @click="$emit('category-select', cat.keyword)"
            class="group relative flex flex-col items-start gap-4 rounded-2xl border border-white/10 bg-white/[0.03] p-5 text-left transition-all duration-300 hover:-translate-y-1 hover:border-brand-400/30 hover:bg-white/[0.05] hover:shadow-lg hover:shadow-brand-500/10 sm:p-6"
          >
            <span class="flex h-11 w-11 items-center justify-center rounded-xl bg-brand-400/10 text-xl text-brand-300 transition-colors group-hover:bg-brand-400/15">
              {{ cat.icon }}
            </span>

            <div class="min-w-0">
              <h3 class="text-sm font-bold leading-snug text-white sm:text-[15px]">
                {{ cat.name }}
              </h3>
              <p class="mt-1 text-xs font-medium text-slate-500">
                {{ cat.jobsCount }}+ việc làm
              </p>
            </div>

            <span class="flex items-center gap-1 text-xs font-semibold text-slate-500 transition-all group-hover:gap-2 group-hover:text-brand-300">
              <span>Khám phá</span>
              <svg class="h-3.5 w-3.5 transition-transform group-hover:translate-x-0.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M14 5l7 7m0 0l-7 7m7-7H3" />
              </svg>
            </span>
          </button>
        </Transition>
      </div>

      <!-- MICRO-TRANSITION: bridges Categories (orientation) -> Jobs (action) -->
      <Transition name="reveal" appear>
        <p v-if="visible" class="mt-12 text-center text-sm font-medium text-slate-500">
          <span class="text-brand-300">↓</span>
          Những cơ hội đang chờ bạn khám phá
        </p>
      </Transition>
    </div>
  </section>
</template>

<script setup lang="ts">
/**
 * HomeCategories.vue
 * Presentation-only: renders category navigation and emits the user's choice.
 * Owns no job data, no fetching, no filtering — filterByCategory() stays in pages/index.vue,
 * which is the existing contract this component reuses via `category-select`.
 */
import { ref, onMounted, onBeforeUnmount } from 'vue'

defineEmits<{
  'category-select': [keyword: string]
}>()

const categories = [
  { icon: '💻', name: 'CNTT & Phần mềm', jobsCount: 45, keyword: 'it' },
  { icon: '📈', name: 'Marketing & Media', jobsCount: 38, keyword: 'marketing' },
  { icon: '🎨', name: 'Thiết Kế Đồ Họa', jobsCount: 24, keyword: 'design' },
  { icon: '☕', name: 'F&B & Phục Vụ', jobsCount: 52, keyword: 'f&b' },
  { icon: '📚', name: 'Gia Sư & Giáo Dục', jobsCount: 19, keyword: 'giasu' },
  { icon: '🛍️', name: 'Bán Hàng & Sales', jobsCount: 31, keyword: 'sales' }
]

// Entrance animation: reveal once when scrolled into view, then disconnect. No re-trigger on scroll.
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

@media (prefers-reduced-motion: reduce) {
  .reveal-enter-active {
    transition: none !important;
  }
  .reveal-enter-from {
    opacity: 1 !important;
    transform: none !important;
  }
}
</style>
