<template>
  <section id="features" ref="sectionRef" class="relative border-b border-white/10 bg-slate-900/40 py-16 lg:py-24">
    <div class="pointer-events-none absolute inset-0 opacity-[0.04] [background-image:linear-gradient(to_right,white_1px,transparent_1px),linear-gradient(to_bottom,white_1px,transparent_1px)] [background-size:56px_56px] [mask-image:radial-gradient(ellipse_70%_60%_at_50%_0%,black_30%,transparent_100%)]"></div>
    <div class="pointer-events-none absolute -top-24 right-1/4 h-[28rem] w-[28rem] rounded-full bg-brand-500/10 blur-[130px]"></div>

    <div class="relative mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
      <Transition name="reveal" appear>
        <div v-if="visible" class="mx-auto mb-14 max-w-2xl text-center">
          <p class="text-xs font-semibold uppercase tracking-[0.24em] text-brand-300">Không chắc mình phù hợp?</p>
          <h2 class="mt-3 text-2xl font-bold tracking-tight text-white sm:text-3xl lg:text-[2.25rem]">
            Để QuickWork giúp bạn nhìn rõ<br class="hidden sm:block" />
            điểm mạnh và hướng đi tiếp theo
          </h2>
        </div>
      </Transition>

      <div class="grid grid-cols-1 items-center gap-10 lg:grid-cols-12 lg:gap-12">
        <!-- LEFT: pitch + checklist + demo trigger + CTA -->
        <Transition name="reveal" appear>
          <div v-if="visible" class="lg:col-span-5">
            <div class="inline-flex items-center gap-2 rounded-full border border-brand-400/25 bg-brand-400/[0.07] px-3.5 py-1.5 text-xs font-semibold text-brand-200">
              <span>✦</span>
              <span>AI CV Analysis</span>
            </div>

            <h3 class="mt-5 text-xl font-bold text-white sm:text-2xl">Tải CV của bạn</h3>
            <p class="mt-2 text-sm text-slate-400 sm:text-base">
              QuickWork phân tích:
            </p>

            <ul class="mt-4 space-y-2.5">
              <li v-for="item in checklist" :key="item" class="flex items-center gap-2.5 text-sm text-slate-300">
                <svg class="h-4 w-4 flex-shrink-0 text-emerald-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 13l4 4L19 7" />
                </svg>
                <span>{{ item }}</span>
              </li>
            </ul>

            <button
              type="button"
              @click="runDemo"
              :disabled="stage === 'analyzing'"
              class="mt-7 inline-flex items-center gap-2 rounded-xl bg-brand-400 px-5 py-3 text-sm font-bold text-slate-950 transition-all hover:bg-brand-300 disabled:cursor-not-allowed disabled:opacity-60 active:scale-[0.98]"
            >
              <span>{{ stage === 'analyzing' ? statusLabel : 'Phân tích CV mẫu' }}</span>
              <svg v-if="stage !== 'analyzing'" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M14 5l7 7m0 0l-7 7m7-7H3" />
              </svg>
              <svg v-else class="h-4 w-4 animate-spin" fill="none" viewBox="0 0 24 24" aria-hidden="true">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v4a4 4 0 00-4 4H4z"></path>
              </svg>
            </button>

            <p class="mt-3 text-xs text-slate-500">
              Đây là mô phỏng trải nghiệm AI — không sử dụng CV thật của bạn.
            </p>

            <div class="mt-8 space-y-2 border-t border-white/10 pt-6">
              <p class="text-sm font-semibold text-white">Muốn phân tích CV của bạn?</p>
              <div class="flex flex-wrap items-center gap-x-5 gap-y-2">
                <NuxtLink to="/register" class="rounded-xl bg-white/[0.06] border border-white/10 px-4 py-2 text-xs font-bold text-white transition-colors hover:bg-white/10">
                  Đăng ký miễn phí
                </NuxtLink>
                <NuxtLink to="/student/dashboard" class="text-xs font-bold text-brand-300 hover:text-brand-200 hover:underline">
                  Đánh giá CV thật tại Dashboard →
                </NuxtLink>
              </div>
            </div>
          </div>
        </Transition>

        <!-- RIGHT: CV Analysis result card -->
        <Transition name="reveal-scale" appear>
          <div v-if="visible" class="lg:col-span-7">
            <div class="relative overflow-hidden rounded-3xl border border-white/10 bg-slate-900/70 p-6 shadow-2xl shadow-slate-950/60 backdrop-blur-xl sm:p-7">
              <div class="absolute inset-0 opacity-[0.05] [background-image:linear-gradient(to_right,white_1px,transparent_1px),linear-gradient(to_bottom,white_1px,transparent_1px)] [background-size:22px_22px]"></div>

              <div class="relative">
                <div class="flex items-center justify-between">
                  <div class="flex items-center gap-2 text-xs font-bold uppercase tracking-wider text-brand-300">
                    <span>✦</span>
                    <span>CV Analysis</span>
                  </div>
                  <span class="rounded-md border border-white/10 bg-white/[0.03] px-2 py-0.5 text-[9px] font-bold uppercase tracking-wider text-slate-500">
                    Demo · Minh họa
                  </span>
                </div>

                <!-- Idle state -->
                <div v-if="stage === 'idle'" class="mt-8 flex flex-col items-center justify-center py-10 text-center">
                  <span class="flex h-12 w-12 items-center justify-center rounded-xl bg-white/[0.04] text-slate-500">
                    <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                    </svg>
                  </span>
                  <p class="mt-3 text-sm font-semibold text-slate-300">CV mẫu sẵn sàng</p>
                  <p class="mt-1 text-xs text-slate-500">Bấm "Phân tích CV mẫu" để xem QuickWork hoạt động</p>
                </div>

                <!-- Analyzing state -->
                <div v-else-if="stage === 'analyzing'" class="mt-8 py-10">
                  <p class="text-center text-sm font-semibold text-brand-200">{{ statusLabel }}</p>
                  <div class="mx-auto mt-4 h-1.5 w-full max-w-xs overflow-hidden rounded-full bg-white/5">
                    <div class="h-full rounded-full bg-gradient-to-r from-brand-400 to-emerald-400 transition-all duration-200 ease-linear" :style="{ width: `${progress}%` }"></div>
                  </div>
                </div>

                <!-- Result state -->
                <div v-else class="mt-6 space-y-5">
                  <div class="space-y-3">
                    <div v-for="(bar, i) in demoResult.bars" :key="bar.label" class="space-y-1">
                      <div class="flex items-center justify-between text-xs">
                        <span class="font-semibold text-slate-300">{{ bar.label }}</span>
                        <span class="font-bold text-white">{{ bar.value }}%</span>
                      </div>
                      <div class="h-1.5 w-full overflow-hidden rounded-full bg-white/5">
                        <div
                          class="skill-bar-fill h-full rounded-full bg-gradient-to-r from-brand-400 to-emerald-400 transition-all ease-out"
                          :style="{ width: resultRevealed ? `${bar.value}%` : '0%', transitionDuration: '900ms', transitionDelay: `${i * 120}ms` }"
                        ></div>
                      </div>
                    </div>
                  </div>

                  <div class="border-t border-white/5 pt-4">
                    <p class="text-[10px] font-semibold uppercase tracking-wide text-slate-500">Gợi ý cải thiện</p>
                    <ul class="mt-2 space-y-1.5">
                      <li v-for="s in demoResult.suggestions" :key="s" class="flex items-start gap-2 text-xs text-slate-300">
                        <span class="mt-0.5 text-brand-300">•</span>
                        <span>{{ s }}</span>
                      </li>
                    </ul>
                  </div>

                  <p class="border-t border-white/5 pt-3 text-[11px] text-slate-500">Dựa trên CV mẫu minh họa, không phải hồ sơ thật của bạn.</p>
                </div>
              </div>
            </div>
          </div>
        </Transition>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
/**
 * HomeAiShowcase.vue
 * Presentation-only: shows what QuickWork's AI CV analysis looks like using a local, illustrative demo.
 * No API calls, no Gemini, no CV upload, no auth logic, no real user data.
 */
import { ref, nextTick, onMounted, onBeforeUnmount } from 'vue'

const checklist = ['Skills', 'Experience', 'Projects', 'Career signals']

// Illustrative-only result — clearly labeled "Demo · Minh họa" in the UI.
const demoResult = {
  bars: [
    { label: 'Skills', value: 92 },
    { label: 'Experience', value: 78 },
    { label: 'Projects', value: 85 },
    { label: 'ATS readiness', value: 95 }
  ],
  suggestions: [
    'Làm nổi bật kỹ năng Go (Golang) & Docker ngay đầu CV',
    'Thêm số liệu định lượng vào phần dự án đã thực hiện',
    'Bổ sung chứng chỉ hoặc khóa học liên quan đến Cloud/DevOps'
  ]
}

const stage = ref<'idle' | 'analyzing' | 'done'>('idle')
const progress = ref(0)
const resultRevealed = ref(false)
let progressTimer: ReturnType<typeof setInterval> | null = null
let revealFrame: number | null = null

const milestones = [
  { at: 30, label: 'Đang phát hiện kỹ năng...' },
  { at: 60, label: 'Đang phân tích kinh nghiệm...' },
  { at: 90, label: 'Đang tạo gợi ý cải thiện...' }
]
const statusLabel = ref('Đang phân tích CV...')

const prefersReducedMotion = ref(false)

const clearProgressTimer = () => {
  if (progressTimer) {
    clearInterval(progressTimer)
    progressTimer = null
  }
}

const revealResult = () => {
  resultRevealed.value = false
  if (prefersReducedMotion.value) {
    resultRevealed.value = true
    return
  }
  nextTick(() => {
    revealFrame = requestAnimationFrame(() => {
      resultRevealed.value = true
      revealFrame = null
    })
  })
}

const runDemo = () => {
  if (stage.value === 'analyzing') return

  if (prefersReducedMotion.value) {
    // Reduced motion: skip the progress animation, reveal the result immediately.
    stage.value = 'done'
    revealResult()
    return
  }

  stage.value = 'analyzing'
  progress.value = 0
  statusLabel.value = 'Đang phân tích CV...'

  progressTimer = setInterval(() => {
    progress.value += 10
    const reached = [...milestones].reverse().find(m => progress.value >= m.at)
    if (reached) statusLabel.value = reached.label

    if (progress.value >= 100) {
      clearProgressTimer()
      stage.value = 'done'
      revealResult()
    }
  }, 160) // ~1.6s total — matches the previous simulation's pacing
}

// Entrance: reveal once when scrolled into view.
const sectionRef = ref<HTMLElement | null>(null)
const visible = ref(false)
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
  clearProgressTimer()
  if (revealFrame) cancelAnimationFrame(revealFrame)
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
  transform: scale(0.97) translateY(14px);
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
  .skill-bar-fill {
    transition-duration: 0.01ms !important;
    transition-delay: 0ms !important;
  }
}
</style>
