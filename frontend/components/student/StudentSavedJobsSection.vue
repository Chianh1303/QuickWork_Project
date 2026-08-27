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

      <div class="flex items-center gap-2">
        <button
          @click="fetchSavedJobs"
          title="Làm mới danh sách"
          class="p-2 bg-slate-800 hover:bg-slate-700 text-cyan-300 border border-slate-700 rounded-xl text-xs font-bold transition-all cursor-pointer"
        >
          🔄
        </button>
        <button
          @click="activeSection = 'jobs'"
          class="px-4 py-2 bg-cyan-500/10 hover:bg-cyan-500/20 text-cyan-300 border border-cyan-500/30 rounded-xl text-xs font-extrabold transition-all cursor-pointer"
        >
          🔍 Tìm thêm việc làm
        </button>
      </div>
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

    <!-- Compact 3-Column Job Card Grid -->
    <div v-else class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-5">
      <div
        v-for="job in savedJobs"
        :key="getJobId(job)"
        class="group relative rounded-2xl border border-cyan-500/15 bg-slate-950/80 p-5 shadow-xl transition-all duration-300 flex flex-col justify-between hover:border-cyan-400/50 hover:bg-slate-950 hover:shadow-2xl hover:shadow-cyan-500/10"
      >
        <div class="space-y-3.5">
          <!-- Top Row: Logo & Company Name & Heart Button -->
          <div class="flex items-start justify-between gap-3">
            <div class="flex items-center gap-3 min-w-0 flex-1">
              <div class="relative h-11 w-11 flex-shrink-0 overflow-hidden rounded-xl border border-cyan-500/20 bg-slate-800 shadow-md flex items-center justify-center">
                <img
                  v-if="job.business?.logo_url || job.logo_url"
                  :src="getMediaUrl(job.business?.logo_url || job.logo_url)"
                  :alt="getCompanyName(job)"
                  class="h-full w-full object-cover relative z-10"
                  @error="handleImgError"
                />
                <div class="absolute inset-0 flex items-center justify-center bg-gradient-to-br from-cyan-600 via-blue-600 to-emerald-500 text-xs font-black text-white">
                  {{ getCompanyInitial(getCompanyName(job)) }}
                </div>
              </div>

              <div class="min-w-0 flex-1">
                <span class="inline-block max-w-full truncate rounded-lg bg-cyan-500/10 px-2.5 py-0.5 text-xs font-bold text-cyan-200 border border-cyan-500/20">
                  {{ getCompanyName(job) }}
                </span>
                <h3 class="mt-1 text-base font-extrabold text-white group-hover:text-cyan-300 transition-colors line-clamp-1">
                  {{ displayJobTitle(job.title || job.job_title) }}
                </h3>
              </div>
            </div>

            <!-- Remove from saved Heart Button -->
            <button
              @click="toggleSaveJob(job)"
              title="Bỏ lưu công việc"
              class="p-2 rounded-xl bg-rose-500/10 hover:bg-rose-500/20 text-rose-400 border border-rose-500/20 transition-all hover:scale-110 cursor-pointer flex-shrink-0"
            >
              <svg class="h-4 w-4 fill-rose-500 text-rose-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
              </svg>
            </button>
          </div>

          <!-- Salary & Location Row -->
          <div class="flex flex-wrap items-center gap-2 pt-1 border-t border-cyan-500/10 text-xs">
            <span class="font-black text-emerald-400">
              💰 {{ formatCurrency(job.salary) }}
            </span>
            <span v-if="job.location" class="text-slate-400 font-medium truncate max-w-[150px]">
              📍 {{ job.location }}
            </span>
          </div>

          <!-- Description Preview -->
          <p class="text-xs text-slate-300 font-medium line-clamp-2 leading-relaxed">
            {{ job.description }}
          </p>
        </div>

        <!-- Action Footer -->
        <div class="mt-4 pt-3 border-t border-cyan-500/10 flex items-center justify-between gap-2">
          <button
            @click="shareJob(job)"
            title="Sao chép liên kết chia sẻ"
            class="p-2 rounded-xl bg-slate-800 hover:bg-slate-700 text-slate-300 border border-slate-700 transition-all text-xs font-bold cursor-pointer"
          >
            🔗
          </button>

          <button
            @click="$emit('apply', job)"
            class="flex-1 py-2 px-4 rounded-xl text-xs font-black text-slate-950 bg-gradient-to-r from-cyan-400 to-emerald-400 hover:from-cyan-300 hover:to-emerald-300 shadow-lg shadow-cyan-500/20 transition-all text-center cursor-pointer"
          >
            Ứng Tuyển Ngay ⚡
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, toRefs, watch } from 'vue'
import { useMedia } from '~/composables/useMedia'
import { useSavedJobs } from '~/composables/useSavedJobs'

const props = defineProps<{ state: Record<string, any> }>()
defineEmits<{ (e: 'apply', job: any): void }>()

const { getMediaUrl, getCompanyInitial } = useMedia()
const { activeSection, jobs } = toRefs(props.state)
const { savedJobIds, savedJobsList, toggleSaveJob, shareJob, fetchSavedJobs } = useSavedJobs()

// Auto-reload saved jobs list when student switches to 'saved-jobs' section
watch(activeSection, (newSection) => {
  if (newSection === 'saved-jobs') {
    fetchSavedJobs()
  }
}, { immediate: true })

const handleImgError = (event: Event) => {
  const target = event.target as HTMLImageElement
  if (target) {
    target.style.display = 'none'
  }
}

const formatCurrency = (value: number | string | null | undefined) => {
  const amount = Number(value || 0)
  return `${amount.toLocaleString('vi-VN')} VNĐ`
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
