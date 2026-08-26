<template>
  <section
    ref="heroRef"
    class="relative overflow-hidden border-b border-white/10 pt-14 pb-20 lg:pt-20 lg:pb-28"
  >
    <!-- BACKGROUND: single depth system — grid + one brand glow, no competing blobs -->
    <div class="pointer-events-none absolute inset-0">
      <div class="absolute inset-0 bg-[linear-gradient(180deg,#020617_0%,#0b1424_55%,#020617_100%)]"></div>
      <div class="absolute inset-0 opacity-[0.05] [background-image:linear-gradient(to_right,white_1px,transparent_1px),linear-gradient(to_bottom,white_1px,transparent_1px)] [background-size:56px_56px] [mask-image:radial-gradient(ellipse_60%_50%_at_50%_0%,black_40%,transparent_100%)]"></div>
      <div class="absolute -top-24 left-1/2 h-[32rem] w-[32rem] -translate-x-1/2 rounded-full bg-brand-500/10 blur-[140px]"></div>
    </div>

    <!-- MOUSE SPOTLIGHT (desktop only, RAF-throttled, respects reduced motion) -->
    <div
      v-if="showSpotlight"
      class="pointer-events-none absolute inset-0 z-10 hidden lg:block"
      :style="{
        background: `radial-gradient(500px circle at ${spotlight.x}px ${spotlight.y}px, rgba(34, 211, 238, 0.10), transparent 75%)`
      }"
    ></div>

    <div class="relative z-20 mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
      <div class="grid grid-cols-1 items-center gap-16 lg:grid-cols-12 lg:gap-8">

        <!-- LEFT: value statement -> headline -> description -> search -> tags -> social proof -->
        <div class="lg:col-span-7">
          <Transition name="stage" appear>
            <div v-if="stage >= 1" class="inline-flex items-center gap-2 rounded-full border border-brand-400/25 bg-brand-400/[0.07] px-4 py-1.5 text-xs font-semibold tracking-wide text-brand-200 backdrop-blur-md">
              <span class="text-brand-300">✦</span>
              <span>Nền tảng nghề nghiệp dành cho sinh viên</span>
            </div>
          </Transition>

          <Transition name="stage" appear>
            <h1
              v-if="stage >= 2"
              class="mt-6 text-4xl font-bold leading-[1.1] tracking-tight text-white sm:text-5xl lg:text-[3.5rem]"
            >
              Việc làm tốt hơn.<br />
              Sự nghiệp bắt đầu
              <span class="bg-gradient-to-r from-brand-300 to-emerald-300 bg-clip-text text-transparent">thông minh hơn.</span>
            </h1>
          </Transition>

          <Transition name="stage" appear>
            <p v-if="stage >= 3" class="mt-5 max-w-xl text-base leading-relaxed text-slate-400 sm:text-lg">
              Giúp sinh viên tìm cơ hội phù hợp và xây dựng sự nghiệp với dữ liệu, AI và những nhà tuyển dụng đáng tin cậy.
            </p>
          </Transition>

          <!-- SEARCH -->
          <Transition name="stage-up" appear>
            <div v-if="stage >= 4" class="mt-8">
              <form
                @submit.prevent="submitSearch"
                class="flex flex-col gap-2 rounded-2xl border border-white/10 bg-white/[0.03] p-2 shadow-xl shadow-slate-950/40 backdrop-blur-xl transition-colors focus-within:border-brand-400/40 sm:flex-row sm:items-center"
              >
                <div class="relative flex-1">
                  <svg class="pointer-events-none absolute left-3.5 top-1/2 h-4 w-4 -translate-y-1/2 text-slate-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                  </svg>
                  <input
                    v-model="keyword"
                    type="text"
                    class="w-full rounded-xl bg-transparent py-3 pl-10 pr-3 text-sm font-medium text-white placeholder-slate-500 focus:outline-none"
                    placeholder="Tìm việc, kỹ năng hoặc vị trí..."
                  />
                </div>

                <div class="hidden h-8 w-px bg-white/10 sm:block"></div>

                <div class="relative sm:w-48">
                  <svg class="pointer-events-none absolute left-3.5 top-1/2 h-4 w-4 -translate-y-1/2 text-slate-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.828 0l-4.243-4.243a8 8 0 1111.314 0z" />
                  </svg>
                  <select
                    v-model="location"
                    class="w-full cursor-pointer appearance-none rounded-xl bg-transparent py-3 pl-9 pr-6 text-sm font-medium text-white focus:outline-none"
                  >
                    <option class="bg-slate-900" value="">Tất cả địa điểm</option>
                    <option class="bg-slate-900" value="Hà Nội">Hà Nội</option>
                    <option class="bg-slate-900" value="Hồ Chí Minh">TP. Hồ Chí Minh</option>
                    <option class="bg-slate-900" value="Đà Nẵng">Đà Nẵng</option>
                    <option class="bg-slate-900" value="Remote">Làm việc từ xa (Remote)</option>
                  </select>
                </div>

                <button
                  type="submit"
                  class="flex items-center justify-center gap-1.5 rounded-xl bg-brand-400 px-5 py-3 text-sm font-bold text-slate-950 transition-all hover:bg-brand-300 active:scale-[0.98]"
                >
                  <span>Tìm việc ngay</span>
                  <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M14 5l7 7m0 0l-7 7m7-7H3" />
                  </svg>
                </button>
              </form>

              <!-- SEARCH TAGS: understated, single accent, no per-tag colors -->
              <div class="mt-3.5 flex flex-wrap items-center gap-2">
                <button
                  v-for="tag in searchTags"
                  :key="tag"
                  type="button"
                  @click="selectTag(tag)"
                  class="rounded-lg border border-white/10 bg-white/[0.02] px-3 py-1.5 text-xs font-medium text-slate-400 transition-colors hover:border-brand-400/30 hover:bg-brand-400/[0.06] hover:text-brand-200"
                >
                  {{ tag }}
                </button>
              </div>
            </div>
          </Transition>

          <!-- SOCIAL PROOF -->
          <Transition name="stage" appear>
            <div v-if="stage >= 4" class="mt-8 flex flex-wrap items-center gap-x-8 gap-y-3 border-t border-white/5 pt-6 text-sm">
              <div class="flex items-baseline gap-1.5">
                <span class="font-bold text-white">{{ countJobs }}+</span>
                <span class="text-slate-500">việc làm mở mới</span>
              </div>
              <div class="flex items-baseline gap-1.5">
                <span class="font-bold text-white">{{ countBusinesses }}+</span>
                <span class="text-slate-500">doanh nghiệp KYB</span>
              </div>
              <div class="flex items-baseline gap-1.5">
                <span class="font-bold text-white">100%</span>
                <span class="text-slate-500">nhà tuyển dụng xác minh</span>
              </div>
            </div>
          </Transition>
        </div>

        <!-- RIGHT: single focal point — AI Match illustrative showcase -->
        <div class="relative lg:col-span-5">
          <Transition name="stage-scale" appear>
            <div v-if="stage >= 5" class="relative mx-auto max-w-sm lg:max-w-none">
              <!-- ambient glow behind the card only -->
              <div class="absolute -inset-6 rounded-[2rem] bg-brand-500/10 blur-3xl"></div>

              <!-- floating badge: CV score (desktop only, single entrance, no infinite bounce) -->
              <Transition name="float-in">
                <div
                  v-if="stage >= 6"
                  class="absolute -left-6 -top-6 z-10 hidden items-center gap-2.5 rounded-2xl border border-white/10 bg-slate-900/90 px-3.5 py-2.5 opacity-90 shadow-xl backdrop-blur-md lg:flex"
                >
                  <span class="flex h-8 w-8 items-center justify-center rounded-lg bg-emerald-400/15 text-emerald-300">
                    <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/></svg>
                  </span>
                  <div>
                    <p class="text-[10px] font-semibold uppercase tracking-wide text-slate-500">CV Score</p>
                    <p class="text-sm font-bold text-white">95<span class="text-slate-500">/100</span></p>
                  </div>
                </div>
              </Transition>

              <!-- floating badge: jobs matched -->
              <Transition name="float-in">
                <div
                  v-if="stage >= 6"
                  class="absolute -bottom-5 -right-4 z-10 hidden items-center gap-2.5 rounded-2xl border border-white/10 bg-slate-900/90 px-3.5 py-2.5 opacity-90 shadow-xl backdrop-blur-md lg:flex"
                >
                  <span class="flex h-8 w-8 items-center justify-center rounded-lg bg-brand-400/15 text-brand-300">
                    <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" /></svg>
                  </span>
                  <div>
                    <p class="text-[10px] font-semibold uppercase tracking-wide text-slate-500">Việc phù hợp</p>
                    <p class="text-sm font-bold text-white">3 jobs</p>
                  </div>
                </div>
              </Transition>

              <!-- MAIN CARD: AI Match illustration -->
              <div class="relative overflow-hidden rounded-3xl border border-white/10 bg-slate-900/70 p-6 shadow-2xl shadow-slate-950/60 backdrop-blur-xl">
                <div class="absolute inset-0 opacity-[0.06] [background-image:linear-gradient(to_right,white_1px,transparent_1px),linear-gradient(to_bottom,white_1px,transparent_1px)] [background-size:22px_22px]"></div>

                <div class="relative">
                  <div class="flex items-center justify-between">
                    <div class="flex items-center gap-2 text-xs font-bold uppercase tracking-wider text-brand-300">
                      <span>✦</span>
                      <span>AI Match</span>
                    </div>
                    <span class="text-2xl font-bold text-white">94<span class="text-base text-slate-500">%</span></span>
                  </div>

                  <div class="mt-5 border-t border-white/5 pt-5">
                    <p class="text-sm font-bold text-white">Backend Developer</p>
                    <p class="text-xs text-slate-500">Software Company</p>
                  </div>

                  <div class="mt-4 space-y-2">
                    <p class="text-[10px] font-semibold uppercase tracking-wide text-slate-500">Skills matched</p>
                    <TransitionGroup v-if="stage >= 6" name="skill-row" tag="div" appear class="space-y-1.5">
                      <div v-for="(skill, i) in skills" :key="skill" :style="{ transitionDelay: `${i * 90}ms` }" class="flex items-center justify-between rounded-lg bg-white/[0.02] px-2.5 py-1.5 text-xs">
                        <span class="font-medium text-slate-300">{{ skill }}</span>
                        <svg class="h-3.5 w-3.5 text-emerald-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 13l4 4L19 7" /></svg>
                      </div>
                    </TransitionGroup>
                  </div>

                  <div class="mt-5 border-t border-white/5 pt-4">
                    <div class="flex items-center justify-between text-xs">
                      <span class="font-semibold text-slate-400">Match quality</span>
                      <span class="font-bold text-emerald-300">Excellent</span>
                    </div>
                    <div class="mt-2 h-1.5 w-full overflow-hidden rounded-full bg-white/5">
                      <div
                        class="h-full rounded-full bg-gradient-to-r from-brand-400 to-emerald-400 transition-all duration-[1200ms] ease-out"
                        :style="{ width: stage >= 6 ? '94%' : '0%' }"
                      ></div>
                    </div>
                  </div>

                  <div class="mt-4 flex items-center justify-between border-t border-white/5 pt-3.5">
                    <p class="text-[11px] text-slate-500">Based on skills &amp; experience</p>
                    <span class="rounded-md border border-white/10 bg-white/[0.03] px-2 py-0.5 text-[9px] font-bold uppercase tracking-wider text-slate-500">
                      Demo · Minh họa
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </Transition>
        </div>

      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
/**
 * HomeHero.vue
 * Presentation-only: Hero visuals + Hero-local interaction/animation.
 * No business logic, no real job fetching, no AI API calls, no real match scoring.
 * All "AI Match" figures in this component are illustrative demo data (labelled "Demo · Minh họa").
 */
import { ref, reactive, onMounted, onBeforeUnmount } from 'vue'

const emit = defineEmits<{
  search: [payload: { keyword: string; location: string }]
  'tag-search': [tag: string]
}>()

const keyword = ref('')
const location = ref('')

const searchTags = ['Backend Developer', 'Frontend Developer', 'Data Analyst', 'UI/UX Designer']

// Illustrative-only skill list for the AI Match demo card.
const skills = ['Go', 'MySQL', 'Docker', 'REST API']

const submitSearch = () => {
  emit('search', { keyword: keyword.value, location: location.value })
}

const selectTag = (tag: string) => {
  keyword.value = tag
  emit('tag-search', tag)
}

// --- Staged entrance sequence (runs once, no repeat) ---
// 1 badge -> 2 headline -> 3 description -> 4 search+proof -> 5 showcase card -> 6 skills/floating badges
const stage = ref(0)
const prefersReducedMotion = ref(false)
let stageTimers: ReturnType<typeof setTimeout>[] = []

const runEntrance = () => {
  if (prefersReducedMotion.value) {
    stage.value = 6
    return
  }
  const delays = [0, 120, 260, 400, 560, 900]
  delays.forEach((delay, idx) => {
    stageTimers.push(setTimeout(() => { stage.value = idx + 1 }, delay))
  })
}

// --- Social proof counters (illustrative marketing figures, matches prior Home behavior) ---
const countJobs = ref(0)
const countBusinesses = ref(0)
let counterTimer: ReturnType<typeof setInterval> | null = null

const animateCounters = () => {
  if (prefersReducedMotion.value) {
    countJobs.value = 300
    countBusinesses.value = 150
    return
  }
  const duration = 1200
  const steps = 24
  const stepTime = duration / steps
  let currentStep = 0
  counterTimer = setInterval(() => {
    currentStep++
    countJobs.value = Math.min(300, Math.floor((300 / steps) * currentStep))
    countBusinesses.value = Math.min(150, Math.floor((150 / steps) * currentStep))
    if (currentStep >= steps && counterTimer) {
      clearInterval(counterTimer)
      counterTimer = null
    }
  }, stepTime)
}

// --- Mouse spotlight: RAF-throttled, desktop-only, disabled on reduced motion ---
const heroRef = ref<HTMLElement | null>(null)
const spotlight = reactive({ x: 0, y: 0 })
const showSpotlight = ref(false)
let rafId: number | null = null
let pendingX = 0
let pendingY = 0
let hasPendingMove = false

const flushSpotlight = () => {
  spotlight.x = pendingX
  spotlight.y = pendingY
  hasPendingMove = false
  rafId = null
}

const handleMouseMove = (e: MouseEvent) => {
  if (!heroRef.value) return
  const rect = heroRef.value.getBoundingClientRect()
  pendingX = e.clientX - rect.left
  pendingY = e.clientY - rect.top
  if (!hasPendingMove) {
    hasPendingMove = true
    rafId = requestAnimationFrame(flushSpotlight)
  }
}

let mediaQuery: MediaQueryList | null = null
const handleMotionPreferenceChange = () => {
  prefersReducedMotion.value = mediaQuery?.matches ?? false
  showSpotlight.value = !prefersReducedMotion.value && window.innerWidth >= 1024
}

onMounted(() => {
  if (!import.meta.client) return

  mediaQuery = window.matchMedia('(prefers-reduced-motion: reduce)')
  prefersReducedMotion.value = mediaQuery.matches
  mediaQuery.addEventListener('change', handleMotionPreferenceChange)

  showSpotlight.value = !prefersReducedMotion.value && window.innerWidth >= 1024
  if (showSpotlight.value) {
    window.addEventListener('mousemove', handleMouseMove, { passive: true })
  }

  runEntrance()
  animateCounters()
})

onBeforeUnmount(() => {
  if (!import.meta.client) return
  window.removeEventListener('mousemove', handleMouseMove)
  mediaQuery?.removeEventListener('change', handleMotionPreferenceChange)
  if (rafId) cancelAnimationFrame(rafId)
  if (counterTimer) clearInterval(counterTimer)
  stageTimers.forEach(clearTimeout)
  stageTimers = []
})
</script>

<style scoped>
.stage-enter-active {
  transition: opacity 0.5s cubic-bezier(0.16, 1, 0.3, 1), transform 0.5s cubic-bezier(0.16, 1, 0.3, 1);
}
.stage-enter-from {
  opacity: 0;
  transform: translateY(10px);
}

.stage-up-enter-active {
  transition: opacity 0.55s cubic-bezier(0.16, 1, 0.3, 1), transform 0.55s cubic-bezier(0.16, 1, 0.3, 1);
}
.stage-up-enter-from {
  opacity: 0;
  transform: translateY(18px);
}

.stage-scale-enter-active {
  transition: opacity 0.6s cubic-bezier(0.16, 1, 0.3, 1), transform 0.6s cubic-bezier(0.16, 1, 0.3, 1);
}
.stage-scale-enter-from {
  opacity: 0;
  transform: scale(0.96) translateY(12px);
}

.float-in-enter-active {
  transition: opacity 0.45s ease-out, transform 0.45s ease-out;
}
.float-in-enter-from {
  opacity: 0;
  transform: translateY(8px);
}

.skill-row-enter-active {
  transition: opacity 0.35s ease-out, transform 0.35s ease-out;
}
.skill-row-enter-from {
  opacity: 0;
  transform: translateX(-6px);
}

@media (prefers-reduced-motion: reduce) {
  .stage-enter-active,
  .stage-up-enter-active,
  .stage-scale-enter-active,
  .float-in-enter-active,
  .skill-row-enter-active {
    transition: none !important;
  }
  .stage-enter-from,
  .stage-up-enter-from,
  .stage-scale-enter-from,
  .float-in-enter-from,
  .skill-row-enter-from {
    opacity: 1 !important;
    transform: none !important;
  }
}
</style>
