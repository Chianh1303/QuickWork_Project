<template>
  <div v-show="activeSection === 'saved-jobs'" class="space-y-6 max-w-7xl mx-auto">
    <!-- Header Banner -->
    <div class="rounded-3xl border border-cyan-500/20 bg-slate-900/90 p-6 flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4 shadow-2xl backdrop-blur-xl">
      <div class="flex items-center space-x-4">
        <div class="h-12 w-12 rounded-2xl bg-gradient-to-tr from-rose-500 via-pink-500 to-amber-400 flex items-center justify-center text-white text-xl shadow-lg shadow-rose-500/25 flex-shrink-0">
          ❤️
        </div>
        <div>
          <div class="flex items-center gap-2">
            <h2 class="text-xl font-extrabold text-white tracking-tight">Danh Sách Việc Làm Đã Lưu</h2>
            <span class="inline-flex items-center rounded-full bg-rose-500/15 px-3 py-0.5 text-xs font-black text-rose-300 ring-1 ring-rose-500/30 uppercase tracking-wider">
              {{ savedJobs.length }} Việc làm
            </span>
          </div>
          <p class="text-xs font-medium text-slate-400 mt-0.5">Các công việc bạn đã thả tim lưu lại để theo dõi và ứng tuyển sau</p>
        </div>
      </div>

      <button
        @click="activeSection = 'jobs'"
        class="px-4 py-2 bg-cyan-500/10 hover:bg-cyan-500/20 text-cyan-300 border border-cyan-500/30 rounded-xl text-xs font-extrabold transition-all cursor-pointer"
      >
        🔍 Tìm thêm việc làm
      </button>
    </div>

    <!-- Empty State -->
    <div v-if="savedJobs.length === 0" class="bg-slate-900/90 text-center py-16 px-4 rounded-2xl border border-cyan-500/15 text-slate-400 shadow-xl">
      <svg class="mx-auto h-12 w-12 text-slate-400 mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
      </svg>
      <p class="font-bold text-slate-100 text-lg">Chưa có việc làm yêu thích nào</p>
      <p class="text-sm text-slate-400 mt-1 max-w-md mx-auto">Bạn chưa bấm lưu công việc nào. Hãy tìm kiếm công việc phù hợp và bấm biểu tượng ❤️ để lưu lại!</p>
      <button
        @click="activeSection = 'jobs'"
        class="mt-4 px-6 py-2.5 bg-gradient-to-r from-cyan-500 to-emerald-400 text-slate-950 font-black text-xs rounded-xl shadow-lg shadow-cyan-500/20 hover:scale-105 transition-all cursor-pointer"
      >
        Khám Phá Việc Làm Ngay ⚡
      </button>
    </div>

    <!-- Saved Jobs List (Identical Card Layout matching StudentJobsSection) -->
    <div v-else class="grid grid-cols-1 lg:grid-cols-2 gap-4">
      <div
        v-for="job in savedJobs"
        :key="getJobId(job)"
        class="group relative flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4 rounded-2xl border border-cyan-500/15 bg-slate-900/90 p-4 shadow-md shadow-cyan-950/20 backdrop-blur-xl transition-all duration-200 hover:border-cyan-400/50 hover:bg-slate-900 hover:shadow-xl hover:shadow-cyan-500/10"
      >
        <!-- Left & Middle: Business Logo + Main Details (Clickable to Apply/Detail) -->
        <div @click="$emit('apply', job)" class="flex items-start sm:items-center gap-3.5 min-w-0 flex-1 w-full cursor-pointer">
          <!-- Business Logo -->
          <div class="relative h-14 w-14 sm:h-16 sm:w-16 flex-shrink-0 overflow-hidden rounded-xl border border-cyan-500/20 bg-slate-800 shadow-md flex items-center justify-center group-hover:scale-105 transition-transform">
            <img
              v-if="job.business?.logo_url || job.logo_url"
              :src="getMediaUrl(job.business?.logo_url || job.logo_url)"
              :alt="getCompanyName(job)"
              class="h-full w-full object-cover relative z-10"
              @error="handleImgError"
            />
            <div class="absolute inset-0 flex items-center justify-center bg-gradient-to-br from-cyan-600 via-blue-600 to-emerald-500 text-sm font-black text-white">
              {{ getCompanyInitial(getCompanyName(job)) }}
            </div>
          </div>

          <!-- Main Details -->
          <div class="min-w-0 flex-1 space-y-1">
            <h3 class="text-sm sm:text-base font-extrabold text-white group-hover:text-cyan-300 transition-colors truncate" :title="displayJobTitle(job.title || job.job_title)">
              {{ displayJobTitle(job.title || job.job_title) }}
            </h3>

            <p class="text-xs font-semibold text-slate-300 truncate" :title="getCompanyName(job)">
              {{ getCompanyName(job) }}
            </p>

            <!-- Tags / Badges Row -->
            <div class="flex flex-wrap items-center gap-1.5 pt-0.5">
              <span class="inline-flex items-center text-xs font-black text-emerald-300 bg-emerald-500/10 border border-emerald-500/20 px-2 py-0.5 rounded-md whitespace-nowrap">
                {{ formatCurrency(job.salary) }}
              </span>

              <span v-if="job.location" class="inline-flex items-center gap-1 text-[11px] font-medium text-slate-400 bg-slate-800/80 border border-slate-700/60 px-2 py-0.5 rounded-md truncate max-w-[130px]" :title="job.location">
                <svg class="h-3 w-3 flex-shrink-0 text-indigo-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.243-4.243a8 8 0 1111.314 0z" />
                </svg>
                <span class="truncate">{{ job.location }}</span>
              </span>

              <span v-if="job.working_date" class="hidden sm:inline-flex items-center gap-1 text-[11px] font-medium text-slate-400 bg-slate-800/80 border border-slate-700/60 px-2 py-0.5 rounded-md truncate max-w-[130px]" :title="job.working_date">
                <svg class="h-3 w-3 flex-shrink-0 text-indigo-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
                <span class="truncate">{{ job.working_date }}</span>
              </span>
            </div>
          </div>
        </div>

        <!-- Right Side: Bookmark, Share & Action Button -->
        <div class="w-full sm:w-auto flex-shrink-0 pt-2 sm:pt-0 border-t sm:border-t-0 border-indigo-500/10 flex items-center justify-between sm:justify-end gap-2">
          <div class="flex items-center gap-1.5">
            <!-- Bookmark Heart Button -->
            <button
              type="button"
              @click.stop="toggleSaveJob(job)"
              title="Bỏ lưu việc làm khỏi mục Yêu thích"
              class="p-2 rounded-xl border border-indigo-500/20 bg-slate-950/60 hover:bg-slate-950 transition-colors cursor-pointer"
            >
              <svg
                class="h-4 w-4 transition-transform active:scale-125 text-rose-500 fill-rose-500"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
              </svg>
            </button>

            <!-- Share Link Button -->
            <button
              type="button"
              @click.stop="shareJob(job)"
              title="Sao chép liên kết chia sẻ"
              class="p-2 rounded-xl border border-indigo-500/20 bg-slate-950/60 hover:bg-slate-950 text-slate-400 hover:text-indigo-300 transition-colors cursor-pointer"
            >
              <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
              </svg>
            </button>
          </div>

          <!-- Apply Button -->
          <button
            @click="$emit('apply', job)"
            class="flex-1 sm:flex-initial flex items-center justify-center gap-1.5 py-2 px-4 rounded-xl text-xs font-extrabold bg-gradient-to-r from-indigo-500 to-emerald-500 hover:from-indigo-400 hover:to-emerald-400 text-white shadow-md shadow-indigo-500/20 transition-all cursor-pointer whitespace-nowrap"
          >
            <span>Ứng tuyển</span>
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
import { computed, toRefs } from 'vue'
import { useMedia } from '~/composables/useMedia'
import { useSavedJobs } from '~/composables/useSavedJobs'

const props = defineProps<{ state: Record<string, any> }>()
defineEmits<{ (e: 'apply', job: any): void }>()

const { getMediaUrl, getCompanyInitial } = useMedia()
const { activeSection, jobs } = toRefs(props.state)
const { savedJobIds, savedJobsList, toggleSaveJob, shareJob } = useSavedJobs()

const handleImgError = (event: Event) => {
  const target = event.target as HTMLImageElement
  if (target) {
    target.style.display = 'none'
  }
}

const formatCurrency = (value: number | string | null | undefined) => {
  const amount = Number(value || 0)
  return `${amount.toLocaleString('vi-VN')} VND`
}

const displayJobTitle = (title: string | null | undefined) => {
  return (title || 'Untitled Job').replace(/\bMarketting\b/gi, 'Marketing')
}

const getCompanyName = (job: any): string => {
  return job.company || job.business?.company_name || 'Doanh Nghiệp QuickWork'
}

const getJobId = (job: any): number => {
  if (typeof job === 'number') return job
  return Number(job?.id || job?.ID || job?.job_id || job?.JobID || 0)
}

const savedJobs = computed(() => {
  const ids = savedJobIds?.value || []
  if (!Array.isArray(ids) || ids.length === 0) return []

  const map = new Map<number, any>()

  // 1. Add full job objects stored in savedJobsList
  if (Array.isArray(savedJobsList?.value)) {
    savedJobsList.value.forEach((j: any) => {
      const id = getJobId(j)
      if (id > 0 && ids.includes(id)) {
        map.set(id, j)
      }
    })
  }

  // 2. Add matching jobs from state.jobs
  if (Array.isArray(jobs?.value)) {
    jobs.value.forEach((j: any) => {
      const id = getJobId(j)
      if (id > 0 && ids.includes(id)) {
        map.set(id, j)
      }
    })
  }

  return Array.from(map.values())
})
</script>
