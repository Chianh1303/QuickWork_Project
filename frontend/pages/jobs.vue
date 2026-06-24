<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
    <!-- Hero Header -->
    <div class="text-center max-w-3xl mx-auto mb-16">
      <h2 class="text-base font-semibold text-brand-600 tracking-wide uppercase">Open Positions</h2>
      <p class="mt-2 text-3xl font-extrabold text-slate-900 sm:text-4xl tracking-tight">
        Find the perfect role for your schedule
      </p>
      <p class="mt-4 text-lg text-slate-500">
        Browse active student job vacancies and application offers posted by verified local employers.
      </p>
    </div>

    <!-- Search & Filters -->
    <div class="bg-white p-4 rounded-xl border border-slate-200 shadow-sm max-w-4xl mx-auto mb-10 grid grid-cols-1 sm:grid-cols-2 gap-4">
      <div class="relative">
        <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </span>
        <input
          v-model="searchQuery"
          type="text"
          class="block w-full pl-10 pr-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 placeholder-slate-400 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:bg-white transition-all duration-200"
          placeholder="Filter by title, keywords..."
        />
      </div>
      <div class="relative">
        <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
        </span>
        <input
          v-model="locationQuery"
          type="text"
          class="block w-full pl-10 pr-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 placeholder-slate-400 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:bg-white transition-all duration-200"
          placeholder="Filter by location (e.g. Hanoi, Remote)..."
        />
      </div>
    </div>

    <!-- Alert / Toast -->
    <div v-if="feedback" :class="[
      feedback.type === 'success' ? 'bg-emerald-50 border-emerald-300 text-emerald-800' : 'bg-red-50 border-red-300 text-red-800',
      'border-l-4 p-4 rounded-r-lg max-w-4xl mx-auto mb-8 flex justify-between items-start transition-all duration-300'
    ]">
      <div class="flex items-center space-x-3">
        <svg v-if="feedback.type === 'success'" class="h-5 w-5 text-emerald-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <svg v-else class="h-5 w-5 text-red-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <span class="text-sm font-medium">{{ feedback.message }}</span>
      </div>
      <button @click="feedback = null" class="text-slate-400 hover:text-slate-600">
        <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    </div>

    <!-- Job Grid -->
    <div v-if="isLoadingJobs" class="flex justify-center py-12">
      <svg class="animate-spin h-8 w-8 text-brand-600" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
    </div>

    <div v-else-if="filteredJobs.length === 0" class="text-center py-12 text-slate-500">
      <svg class="mx-auto h-12 w-12 text-slate-300 mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
      </svg>
      <p class="font-semibold text-slate-700">No active job listings found.</p>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 max-w-6xl mx-auto">
      <div
        v-for="job in filteredJobs"
        :key="job.id"
        class="bg-white rounded-xl border border-slate-200 shadow-sm p-6 flex flex-col justify-between hover:shadow-md transition-shadow duration-200"
      >
        <div>
          <div class="flex justify-between items-start mb-3">
            <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-semibold bg-brand-50 text-brand-700">
              Slots: {{ job.slots }}
            </span>
            <span class="text-sm font-extrabold text-emerald-600 bg-emerald-50 px-2.5 py-1 rounded-lg">
              ${{ job.salary.toLocaleString() }}
            </span>
          </div>
          <h3 class="text-lg font-bold text-slate-900 line-clamp-1">{{ job.title }}</h3>
          <p class="text-sm font-medium text-slate-500 mt-1 flex items-center gap-1">
            <svg class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
            </svg>
            {{ job.location }}
          </p>
          <p class="text-sm font-medium text-slate-500 mt-1 flex items-center gap-1" v-if="job.working_date">
            <svg class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
            </svg>
            Date: {{ job.working_date }}
          </p>
          <p class="text-sm text-slate-600 mt-4 line-clamp-4 bg-slate-50 p-3 rounded-lg border border-slate-100">
            {{ job.description }}
          </p>
        </div>

        <div class="mt-6 pt-4 border-t border-slate-100">
          <template v-if="isAuthenticated && userRole === 'student'">
            <button
              @click="handleApply(job.id)"
              :disabled="isApplying === job.id"
              class="w-full flex justify-center py-2 px-4 border border-transparent text-sm font-semibold rounded-lg text-white bg-brand-600 hover:bg-brand-500 transition-colors focus-ring disabled:opacity-50 disabled:cursor-not-allowed shadow-sm"
            >
              <span v-if="isApplying === job.id" class="flex items-center space-x-2">
                <svg class="animate-spin h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                <span>Applying...</span>
              </span>
              <span v-else>Apply Instantly</span>
            </button>
          </template>
          <template v-else-if="isAuthenticated && userRole === 'business'">
            <div class="text-xs text-center font-medium text-slate-400 italic py-2">
              Viewing as Employer
            </div>
          </template>
          <template v-else>
            <NuxtLink
              to="/login"
              class="w-full flex justify-center py-2 px-4 border border-slate-200 text-sm font-semibold rounded-lg text-brand-600 hover:bg-brand-50 hover:text-brand-700 transition-all text-center"
            >
              Sign In to Apply
            </NuxtLink>
          </template>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useApi } from '~/composables/useApi'
import { useAuth } from '~/composables/useAuth'

const api = useApi()
const { isAuthenticated, userRole } = useAuth()

const jobs = ref<any[]>([])
const isLoadingJobs = ref(false)
const isApplying = ref<number | null>(null)
const feedback = ref<{ type: 'success' | 'error'; message: string } | null>(null)

const searchQuery = ref('')
const locationQuery = ref('')

const filteredJobs = computed(() => {
  return jobs.value.filter(job => {
    const matchesSearch = !searchQuery.value ||
      job.title.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      job.description.toLowerCase().includes(searchQuery.value.toLowerCase())
    
    const matchesLocation = !locationQuery.value ||
      job.location.toLowerCase().includes(locationQuery.value.toLowerCase())
      
    return matchesSearch && matchesLocation
  })
})

const fetchJobs = async () => {
  isLoadingJobs.value = true
  try {
    const res = await api.get('/api/jobs')
    jobs.value = res.data || []
  } catch (err: any) {
    console.error('Error fetching jobs:', err)
  } finally {
    isLoadingJobs.value = false
  }
}

const handleApply = async (jobId: number) => {
  isApplying.value = jobId
  feedback.value = null

  try {
    const res = await api.post('/api/jobs/apply', { job_id: jobId })
    feedback.value = {
      type: 'success',
      message: res.message || '🚀 Applied successfully! Waiting for Employer review.'
    }
  } catch (err: any) {
    feedback.value = {
      type: 'error',
      message: err.response?._data?.message || 'Failed to submit application.'
    }
  } finally {
    isApplying.value = null
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

onMounted(() => {
  fetchJobs()
})
</script>
