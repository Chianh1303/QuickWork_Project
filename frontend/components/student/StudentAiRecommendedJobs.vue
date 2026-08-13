<template>
  <div class="bg-slate-900/82 rounded-2xl border border-cyan-400/20 p-6 shadow-xl shadow-slate-950/40 backdrop-blur space-y-6">
    <!-- Header Banner -->
    <div class="flex flex-wrap items-center justify-between gap-4 border-b border-white/10 pb-4">
      <div class="flex items-center space-x-3">
        <div class="inline-flex h-10 w-10 items-center justify-center rounded-xl bg-gradient-to-tr from-cyan-500 to-emerald-400 text-slate-950 shadow-md shadow-cyan-950/50">
          <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
          </svg>
        </div>
        <div>
          <div class="flex items-center gap-2">
            <h2 class="text-xl font-extrabold text-white tracking-tight">Việc Làm AI Gợi Ý</h2>
            <span class="inline-flex items-center rounded-full bg-cyan-400/10 px-2.5 py-0.5 text-xs font-bold text-cyan-300 ring-1 ring-cyan-400/30">
              AI Matching V1
            </span>
          </div>
          <p class="text-xs font-medium text-slate-400 mt-0.5">Phân tích mức độ phù hợp giữa Kỹ năng & CV của bạn với tin tuyển dụng</p>
        </div>
      </div>
      <button
        @click="fetchRecommendations"
        :disabled="isLoading"
        class="inline-flex items-center gap-1.5 rounded-lg border border-white/10 bg-white/5 px-3 py-1.5 text-xs font-semibold text-slate-300 transition-colors hover:bg-white/10 hover:text-white disabled:opacity-50"
      >
        <svg :class="{ 'animate-spin': isLoading }" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
        </svg>
        <span>Làm mới</span>
      </button>
    </div>

    <!-- STATE 1: LOADING (Skeleton Pulse) -->
    <div v-if="isLoading" class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-5">
      <div v-for="n in 3" :key="n" class="rounded-xl border border-white/10 bg-slate-950/60 p-5 space-y-4 animate-pulse">
        <div class="flex justify-between items-center">
          <div class="h-6 bg-slate-800 rounded w-1/3"></div>
          <div class="h-6 bg-slate-800 rounded w-1/4"></div>
        </div>
        <div class="h-5 bg-slate-800 rounded w-3/4"></div>
        <div class="h-4 bg-slate-800 rounded w-1/2"></div>
        <div class="space-y-2 pt-2">
          <div class="h-3 bg-slate-800 rounded w-1/3"></div>
          <div class="flex gap-1.5">
            <div class="h-6 bg-slate-800 rounded-full w-12"></div>
            <div class="h-6 bg-slate-800 rounded-full w-16"></div>
          </div>
        </div>
      </div>
    </div>

    <!-- STATE 2: ERROR -->
    <div v-else-if="hasError" class="rounded-xl border border-red-500/20 bg-red-500/10 p-6 text-center space-y-3">
      <svg class="mx-auto h-10 w-10 text-red-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
      </svg>
      <p class="text-sm font-bold text-red-200">Không thể tải danh sách công việc gợi ý từ AI.</p>
      <button
        @click="fetchRecommendations"
        class="inline-flex items-center justify-center rounded-lg bg-red-500/20 px-4 py-2 text-xs font-bold text-red-200 ring-1 ring-red-500/40 transition-colors hover:bg-red-500/30"
      >
        Thử lại
      </button>
    </div>

    <!-- STATE 3: EMPTY -->
    <div v-else-if="jobs.length === 0" class="rounded-xl border border-white/10 bg-slate-950/40 p-8 text-center space-y-2">
      <svg class="mx-auto h-12 w-12 text-slate-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
      </svg>
      <p class="text-sm font-bold text-slate-300">Chưa có công việc nào phù hợp với hồ sơ của bạn lúc này.</p>
    </div>

    <!-- STATE 4: SUCCESS -->
    <div v-else class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-5">
      <div
        v-for="job in jobs"
        :key="job.job_id"
        class="group relative rounded-xl border border-white/10 bg-slate-950/60 p-5 shadow-md transition-all duration-200 hover:border-cyan-400/40 hover:bg-slate-950/80 flex flex-col justify-between"
      >
        <div class="space-y-3">
          <!-- Top Row: Company & Match Score Badge -->
          <div class="flex items-start justify-between gap-2">
            <span class="inline-flex items-center px-2.5 py-1 rounded-md text-xs font-semibold bg-white/5 text-slate-300 border border-white/10 truncate max-w-[60%]">
              {{ job.company }}
            </span>
            <span
              :class="getScoreBadgeClass(job.match_score)"
              class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-extrabold ring-1 shadow-sm whitespace-nowrap"
            >
              {{ job.match_score }}% Phù hợp
            </span>
          </div>

          <!-- Job Title -->
          <h3 class="text-base font-bold text-white group-hover:text-cyan-300 transition-colors line-clamp-2">
            {{ job.job_title }}
          </h3>

          <!-- Matching Skills Section -->
          <div v-if="job.matching_skills && job.matching_skills.length > 0" class="space-y-1.5 pt-1">
            <span class="text-[11px] font-bold uppercase tracking-wider text-emerald-400/90 flex items-center gap-1">
              <svg class="h-3 w-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 13l4 4L19 7" />
              </svg>
              Kỹ năng phù hợp
            </span>
            <div class="flex flex-wrap gap-1.5">
              <span
                v-for="skill in job.matching_skills"
                :key="skill"
                class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-emerald-500/10 text-emerald-300 border border-emerald-500/20"
              >
                {{ skill }}
              </span>
            </div>
          </div>

          <!-- Missing Skills Section -->
          <div v-if="job.missing_skills && job.missing_skills.length > 0" class="space-y-1.5">
            <span class="text-[11px] font-bold uppercase tracking-wider text-slate-400 flex items-center gap-1">
              <svg class="h-3 w-3 text-slate-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01" />
              </svg>
              Kỹ năng còn thiếu
            </span>
            <div class="flex flex-wrap gap-1.5">
              <span
                v-for="skill in job.missing_skills"
                :key="skill"
                class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-slate-800/80 text-slate-400 border border-slate-700/60"
              >
                {{ skill }}
              </span>
            </div>
          </div>
        </div>

        <!-- Card Footer / CTA Button -->
        <div class="mt-5 pt-3 border-t border-white/10">
          <button
            @click="handleApplyJob(job)"
            class="w-full flex items-center justify-center gap-1.5 py-2 px-3 rounded-lg text-xs font-bold text-slate-950 bg-cyan-400 hover:bg-cyan-300 transition-colors focus-ring shadow-sm"
          >
            <span>Xem chi tiết & Ứng tuyển</span>
            <svg class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
            </svg>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useAi, type RecommendedJobItem } from '~/composables/useAi'

const emit = defineEmits<{
  (e: 'apply', job: RecommendedJobItem): void
}>()

const { getRecommendedJobs } = useAi()

const jobs = ref<RecommendedJobItem[]>([])
const isLoading = ref<boolean>(true)
const hasError = ref<boolean>(false)

const fetchRecommendations = async () => {
  isLoading.value = true
  hasError.value = false
  try {
    const res = await getRecommendedJobs()
    jobs.value = res.jobs || []
  } catch (err) {
    hasError.value = true
    jobs.value = []
  } finally {
    isLoading.value = false
  }
}

const getScoreBadgeClass = (score: number): string => {
  if (score >= 80) {
    return 'bg-emerald-400/10 text-emerald-300 ring-emerald-400/30'
  }
  if (score >= 50) {
    return 'bg-cyan-400/10 text-cyan-300 ring-cyan-400/30'
  }
  return 'bg-slate-800 text-slate-400 ring-slate-700'
}

const handleApplyJob = (job: RecommendedJobItem) => {
  emit('apply', job)
}

onMounted(() => {
  fetchRecommendations()
})
</script>
