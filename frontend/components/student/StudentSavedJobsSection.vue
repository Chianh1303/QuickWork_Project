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

      <NuxtLink
        to="/student/dashboard"
        @click="activeSection = 'jobs'"
        class="px-4 py-2 bg-cyan-500/10 hover:bg-cyan-500/20 text-cyan-300 border border-cyan-500/30 rounded-xl text-xs font-extrabold transition-all"
      >
        🔍 Tìm thêm việc làm
      </NuxtLink>
    </div>

    <!-- Empty State -->
    <div v-if="savedJobs.length === 0" class="rounded-3xl border border-cyan-500/15 bg-slate-900/80 p-12 text-center space-y-4 shadow-xl">
      <div class="h-16 w-16 mx-auto rounded-full bg-slate-800 flex items-center justify-center text-3xl">
        🤍
      </div>
      <div class="space-y-1">
        <h3 class="text-lg font-extrabold text-white">Chưa Có Việc Làm Yêu Thích Nào</h3>
        <p class="text-xs text-slate-400 max-w-md mx-auto">
          Bạn chưa bấm lưu công việc nào. Hãy tìm kiếm công việc phù hợp và bấm biểu tượng Nút Nút Lưu / ❤️ để theo dõi danh sách việc làm tốt nhất!
        </p>
      </div>
      <button
        @click="activeSection = 'jobs'"
        class="px-6 py-2.5 bg-gradient-to-r from-cyan-500 to-emerald-400 text-slate-950 font-black text-xs rounded-xl shadow-lg shadow-cyan-500/20 hover:scale-105 transition-all"
      >
        Khám Phá Việc Làm Ngay ⚡
      </button>
    </div>

    <!-- Saved Jobs Grid -->
    <div v-else class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6">
      <div
        v-for="job in savedJobs"
        :key="job.id"
        class="group relative rounded-3xl border border-cyan-500/15 bg-slate-900/90 p-6 shadow-xl hover:border-cyan-400/40 hover:shadow-2xl hover:shadow-cyan-500/10 transition-all duration-300 flex flex-col justify-between"
      >
        <div class="space-y-4">
          <!-- Top info -->
          <div class="flex items-start justify-between gap-3">
            <div class="flex items-center gap-3">
              <div class="h-12 w-12 rounded-2xl bg-gradient-to-br from-indigo-600 via-cyan-600 to-emerald-500 flex items-center justify-center text-sm font-black text-white shadow-md flex-shrink-0">
                {{ (job.title || 'Q').slice(0, 1).toUpperCase() }}
              </div>
              <div class="min-w-0">
                <span class="inline-block text-[11px] font-bold text-cyan-300 bg-cyan-500/10 px-2 py-0.5 rounded-md border border-cyan-500/20 truncate max-w-[180px]">
                  {{ job.business?.company_name || 'QuickWork Doanh Nghiệp' }}
                </span>
                <h3 class="text-base font-extrabold text-white group-hover:text-cyan-300 transition-colors line-clamp-1 mt-0.5">
                  {{ job.title }}
                </h3>
              </div>
            </div>
            
            <button
              @click="toggleSaveJob(job)"
              title="Bỏ lưu khỏi danh sách"
              class="p-2 rounded-xl bg-rose-500/10 hover:bg-rose-500/20 text-rose-400 border border-rose-500/20 transition-all hover:scale-110"
            >
              ❤️
            </button>
          </div>

          <!-- Location & Salary -->
          <div class="flex flex-wrap items-center gap-2 text-xs border-t border-cyan-500/10 pt-3">
            <span class="font-black text-emerald-400 bg-emerald-500/10 px-2.5 py-1 rounded-lg border border-emerald-500/20">
              💰 {{ (job.salary || 0).toLocaleString('vi-VN') }} VNĐ
            </span>
            <span class="text-slate-300 font-medium bg-slate-800/80 px-2.5 py-1 rounded-lg border border-slate-700/50">
              📍 {{ job.location }}
            </span>
          </div>

          <!-- Description preview -->
          <p class="text-xs text-slate-300 font-medium line-clamp-2 leading-relaxed">
            {{ job.description }}
          </p>
        </div>

        <!-- Action footer -->
        <div class="mt-5 pt-4 border-t border-cyan-500/10 flex items-center justify-between gap-3">
          <button
            @click="shareJob(job)"
            title="Chia sẻ công việc"
            class="p-2.5 rounded-xl bg-slate-800 hover:bg-slate-700 text-slate-300 border border-slate-700 transition-all text-xs font-bold"
          >
            🔗
          </button>

          <button
            @click="$emit('apply', job)"
            class="flex-1 py-2.5 px-4 rounded-xl text-xs font-black text-slate-950 bg-gradient-to-r from-cyan-400 to-emerald-400 hover:from-cyan-300 hover:to-emerald-300 shadow-lg shadow-cyan-500/20 transition-all text-center"
          >
            Ứng Tuyển Ngay ⚡
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, toRefs } from 'vue'
import { useSavedJobs } from '~/composables/useSavedJobs'

const props = defineProps<{ state: Record<string, any> }>()
defineEmits<{ (e: 'apply', job: any): void }>()

const { activeSection, jobs } = toRefs(props.state)
const { savedJobIds, toggleSaveJob, shareJob } = useSavedJobs()

const savedJobs = computed(() => {
  if (!jobs?.value || !savedJobIds.value || savedJobIds.value.length === 0) return []
  return jobs.value.filter((j: any) => {
    const id = Number(j?.id || j?.ID || j?.job_id || j?.JobID || 0)
    return id > 0 && savedJobIds.value.includes(id)
  })
})
</script>
