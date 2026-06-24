<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <!-- Header -->
    <div class="md:flex md:items-center md:justify-between mb-8 pb-6 border-b border-slate-200">
      <div class="flex-1 min-w-0">
        <h2 class="text-3xl font-bold leading-7 text-slate-900 sm:text-4xl sm:truncate">
          Employer Dashboard
        </h2>
        <p class="mt-1 text-sm text-slate-500">
          Publish job opportunities, review student applications, and keep your company profile updated.
        </p>
      </div>
    </div>

    <!-- Navigation Tabs -->
    <div class="mb-8">
      <nav class="flex space-x-4" aria-label="Tabs">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          @click="activeTab = tab.id"
          :class="[
            activeTab === tab.id
              ? 'bg-accent-600 text-white shadow-md shadow-accent-500/10'
              : 'text-slate-600 hover:text-slate-900 hover:bg-slate-100',
            'px-4 py-2.5 font-semibold text-sm rounded-lg transition-all duration-150 focus:outline-none'
          ]"
        >
          {{ tab.name }}
        </button>
      </nav>
    </div>

    <!-- Feedback Banner -->
    <div v-if="feedback" :class="[
      feedback.type === 'success' ? 'bg-emerald-50 border-emerald-300 text-emerald-800' : 'bg-red-50 border-red-300 text-red-800',
      'border-l-4 p-4 rounded-r-lg mb-6 flex justify-between items-start transition-all duration-300 shadow-sm'
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

    <!-- Tab Contents -->
    <div class="mt-6">
      <!-- Tab 1: Manage My Jobs -->
      <div v-show="activeTab === 'manage'" class="space-y-6">
        <div class="bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden">
          <div class="p-6 border-b border-slate-100 flex justify-between items-center">
            <h3 class="text-lg font-bold text-slate-900">Your Posted Jobs</h3>
            <button @click="activeTab = 'create'" class="inline-flex items-center space-x-1 text-sm font-semibold text-accent-600 hover:text-accent-700">
              <span>+ Post a New Job</span>
            </button>
          </div>

          <div v-if="isLoadingJobs" class="flex justify-center py-12">
            <svg class="animate-spin h-8 w-8 text-accent-600" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
          </div>

          <div v-else-if="jobs.length === 0" class="text-center py-12 text-slate-500">
            <svg class="mx-auto h-12 w-12 text-slate-300 mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 4a2 2 0 00-2 2v3m2-3a2 2 0 012 2v3m-2-3a2 2 0 00-2-3m-9 3h.01M9 16h.01" />
            </svg>
            <p class="font-medium text-slate-700">You haven't posted any jobs yet.</p>
            <p class="text-sm text-slate-400 mt-1">Publish listings to hire student workers.</p>
          </div>

          <div v-else class="overflow-x-auto">
            <table class="min-w-full divide-y divide-slate-200">
              <thead class="bg-slate-50">
                <tr>
                  <th scope="col" class="px-6 py-3 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Job Title</th>
                  <th scope="col" class="px-6 py-3 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Location</th>
                  <th scope="col" class="px-6 py-3 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Salary</th>
                  <th scope="col" class="px-6 py-3 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Slots</th>
                  <th scope="col" class="px-6 py-3 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Status</th>
                  <th scope="col" class="px-6 py-3 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Date</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-slate-200">
                <tr v-for="job in jobs" :key="job.id">
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="text-sm font-semibold text-slate-900">{{ job.title }}</div>
                    <div class="text-xs text-slate-400 line-clamp-1 max-w-xs">{{ job.description }}</div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">
                    {{ job.location }}
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-extrabold text-emerald-600">
                    ${{ job.salary.toLocaleString() }}
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">
                    {{ job.slots }}
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span :class="[
                      job.status === 'approved' ? 'bg-emerald-50 text-emerald-700' : 'bg-amber-50 text-amber-700',
                      'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-semibold'
                    ]">
                      {{ job.status || 'pending' }}
                    </span>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">
                    {{ job.working_date || 'N/A' }}
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- Tab 2: Create Job Listing -->
      <div v-show="activeTab === 'create'" class="bg-white rounded-xl border border-slate-200 shadow-sm p-6 max-w-2xl mx-auto">
        <h3 class="text-lg font-bold text-slate-900 border-b border-slate-100 pb-3 mb-6">Create New Job Posting</h3>
        <form @submit.prevent="handleCreateJob" class="space-y-6">
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
            <div class="sm:col-span-2">
              <label for="job_title" class="block text-sm font-semibold text-slate-700 mb-1">Job Title</label>
              <input
                id="job_title"
                type="text"
                v-model="jobForm.title"
                required
                placeholder="e.g. Frontend Developer Assistant"
                class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-accent-500"
              />
            </div>

            <div>
              <label for="job_location" class="block text-sm font-semibold text-slate-700 mb-1">Location</label>
              <input
                id="job_location"
                type="text"
                v-model="jobForm.location"
                required
                placeholder="e.g. Hanoi, Remote"
                class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-accent-500"
              />
            </div>

            <div>
              <label for="job_working_date" class="block text-sm font-semibold text-slate-700 mb-1">Working Date / Duration</label>
              <input
                id="job_working_date"
                type="text"
                v-model="jobForm.working_date"
                placeholder="e.g. July 1st - Dec 31st"
                class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-accent-500"
              />
            </div>

            <div>
              <label for="job_salary" class="block text-sm font-semibold text-slate-700 mb-1">Salary ($)</label>
              <input
                id="job_salary"
                type="number"
                step="0.01"
                v-model="jobForm.salary"
                required
                placeholder="e.g. 500.00"
                class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-accent-500"
              />
            </div>

            <div>
              <label for="job_slots" class="block text-sm font-semibold text-slate-700 mb-1">Target Hires (Slots)</label>
              <input
                id="job_slots"
                type="number"
                v-model="jobForm.slots"
                required
                placeholder="1"
                class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-accent-500"
              />
            </div>

            <div class="sm:col-span-2">
              <label for="job_description" class="block text-sm font-semibold text-slate-700 mb-1">Detailed Description</label>
              <textarea
                id="job_description"
                rows="4"
                v-model="jobForm.description"
                required
                placeholder="Provide a comprehensive job description, responsibilities, requirements, and benefits..."
                class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-accent-500"
              ></textarea>
            </div>
          </div>

          <div class="pt-4 border-t border-slate-100 flex justify-end">
            <button
              type="submit"
              :disabled="isCreatingJob"
              class="px-6 py-2 border border-transparent text-sm font-semibold rounded-lg text-white bg-accent-600 hover:bg-accent-500 focus-ring disabled:opacity-50 disabled:cursor-not-allowed shadow-sm transition-colors"
            >
              <span v-if="isCreatingJob" class="flex items-center space-x-2">
                <svg class="animate-spin h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                <span>Submitting Listing...</span>
              </span>
              <span v-else>Publish Job Listing</span>
            </button>
          </div>
        </form>
      </div>

      <!-- Tab 3: Review Applications -->
      <div v-show="activeTab === 'review'" class="bg-white rounded-xl border border-slate-200 shadow-sm p-6 max-w-2xl mx-auto">
        <div class="border-b border-slate-100 pb-3 mb-6">
          <h3 class="text-lg font-bold text-slate-900">Review Candidate Application</h3>
          <p class="text-xs text-slate-500 mt-1">
            Note: Fetching application lists requires admin/developer configuration. Standardized direct application review by ID is enabled below.
          </p>
        </div>

        <form @submit.prevent="handleReviewApplication" class="space-y-6">
          <div>
            <label for="review_app_id" class="block text-sm font-semibold text-slate-700 mb-1">Application ID</label>
            <input
              id="review_app_id"
              type="number"
              v-model="reviewForm.application_id"
              required
              placeholder="e.g. 5"
              class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-accent-500"
            />
          </div>

          <div>
            <label class="block text-sm font-semibold text-slate-700 mb-3">Decision / Evaluation Status</label>
            <div class="grid grid-cols-2 gap-4">
              <label
                :class="[
                  reviewForm.status === 'approved'
                    ? 'border-emerald-500 bg-emerald-50'
                    : 'border-slate-200 hover:bg-slate-50',
                  'border cursor-pointer rounded-lg p-3 text-center transition-all duration-150'
                ]"
              >
                <input type="radio" value="approved" v-model="reviewForm.status" class="sr-only" />
                <span class="text-sm font-semibold text-emerald-700">Accept / Approve</span>
              </label>

              <label
                :class="[
                  reviewForm.status === 'rejected'
                    ? 'border-red-500 bg-red-50'
                    : 'border-slate-200 hover:bg-slate-50',
                  'border cursor-pointer rounded-lg p-3 text-center transition-all duration-150'
                ]"
              >
                <input type="radio" value="rejected" v-model="reviewForm.status" class="sr-only" />
                <span class="text-sm font-semibold text-red-700">Reject / Decline</span>
              </label>
            </div>
          </div>

          <div class="pt-4 border-t border-slate-100 flex justify-end">
            <button
              type="submit"
              :disabled="isReviewing"
              class="px-6 py-2 border border-transparent text-sm font-semibold rounded-lg text-white bg-accent-600 hover:bg-accent-500 focus-ring disabled:opacity-50 disabled:cursor-not-allowed shadow-sm transition-colors"
            >
              <span v-if="isReviewing" class="flex items-center space-x-2">
                <svg class="animate-spin h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                <span>Evaluating...</span>
              </span>
              <span v-else>Submit Evaluation</span>
            </button>
          </div>
        </form>
      </div>

      <!-- Tab 4: Employer Profile Settings -->
      <div v-show="activeTab === 'profile'" class="bg-white rounded-xl border border-slate-200 shadow-sm p-6 max-w-2xl mx-auto">
        <h3 class="text-lg font-bold text-slate-900 border-b border-slate-100 pb-3 mb-6">Company / Business Information</h3>
        <form @submit.prevent="handleUpdateProfile" class="space-y-6">
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
            <div>
              <label for="profile_company" class="block text-sm font-semibold text-slate-700 mb-1">Company Name</label>
              <input
                id="profile_company"
                type="text"
                v-model="profileForm.company_name"
                required
                class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-accent-500"
              />
            </div>
            <div>
              <label for="profile_phone" class="block text-sm font-semibold text-slate-700 mb-1">Corporate Phone</label>
              <input
                id="profile_phone"
                type="tel"
                v-model="profileForm.phone"
                class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-accent-500"
              />
            </div>
            <div class="sm:col-span-2">
              <label for="profile_address" class="block text-sm font-semibold text-slate-700 mb-1">Company Address</label>
              <input
                id="profile_address"
                type="text"
                v-model="profileForm.address"
                placeholder="123 Corporate Blvd, Hanoi"
                class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-accent-500"
              />
            </div>
            <div class="sm:col-span-2">
              <label for="profile_logo" class="block text-sm font-semibold text-slate-700 mb-1">Logo URL</label>
              <input
                id="profile_logo"
                type="url"
                v-model="profileForm.logo_url"
                placeholder="https://example.com/logo.png"
                class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-accent-500"
              />
            </div>
          </div>

          <div class="pt-4 border-t border-slate-100 flex justify-end">
            <button
              type="submit"
              :disabled="isSavingProfile"
              class="px-6 py-2 border border-transparent text-sm font-semibold rounded-lg text-white bg-accent-600 hover:bg-accent-500 focus-ring disabled:opacity-50 disabled:cursor-not-allowed shadow-sm transition-colors"
            >
              <span v-if="isSavingProfile" class="flex items-center space-x-2">
                <svg class="animate-spin h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                <span>Saving Profile...</span>
              </span>
              <span v-else>Save Changes</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { useApi } from '~/composables/useApi'

definePageMeta({
  middleware: 'auth'
})

const api = useApi()

const activeTab = ref('manage')
const tabs = [
  { id: 'manage', name: 'Manage Job Listings' },
  { id: 'create', name: 'Post a New Job' },
  { id: 'review', name: 'Review Applications' },
  { id: 'profile', name: 'Company Profile' }
]

const jobs = ref<any[]>([])
const isLoadingJobs = ref(false)
const isCreatingJob = ref(false)
const isReviewing = ref(false)
const isSavingProfile = ref(false)
const feedback = ref<{ type: 'success' | 'error'; message: string } | null>(null)

// Forms
const jobForm = reactive({
  title: '',
  description: '',
  location: '',
  salary: 0,
  slots: 1,
  working_date: ''
})

const reviewForm = reactive({
  application_id: null as number | null,
  status: 'approved' as 'approved' | 'rejected'
})

const profileForm = reactive({
  company_name: '',
  phone: '',
  address: '',
  logo_url: ''
})

const fetchJobs = async () => {
  isLoadingJobs.value = true
  try {
    const res = await api.get('/api/jobs')
    // As the public API lists all jobs, the employer sees all vacancies. 
    // In production, we filter or backend supports /api/jobs/my-jobs, 
    // but for now, we display jobs list from backend
    jobs.value = res.data || []
  } catch (err: any) {
    console.error('Error fetching jobs:', err)
  } finally {
    isLoadingJobs.value = false
  }
}

const handleCreateJob = async () => {
  isCreatingJob.value = true
  feedback.value = null

  try {
    const res = await api.post('/api/jobs', {
      title: jobForm.title,
      description: jobForm.description,
      location: jobForm.location,
      salary: Number(jobForm.salary),
      slots: Number(jobForm.slots),
      working_date: jobForm.working_date
    })
    
    feedback.value = {
      type: 'success',
      message: res.message || '🎉 Job posted successfully! Awaiting validation.'
    }

    // Reset Form
    jobForm.title = ''
    jobForm.description = ''
    jobForm.location = ''
    jobForm.salary = 0
    jobForm.slots = 1
    jobForm.working_date = ''

    // Refresh Jobs List and redirect
    await fetchJobs()
    activeTab.value = 'manage'
  } catch (err: any) {
    feedback.value = {
      type: 'error',
      message: err.response?._data?.message || 'Failed to submit job listing.'
    }
  } finally {
    isCreatingJob.value = false
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const handleReviewApplication = async () => {
  if (reviewForm.application_id === null) return

  isReviewing.value = true
  feedback.value = null

  try {
    const res = await api.put('/api/jobs/review-application', {
      application_id: Number(reviewForm.application_id),
      status: reviewForm.status
    })

    feedback.value = {
      type: 'success',
      message: res.message || '🎉 Candidate application reviewed successfully!'
    }

    // Reset Form
    reviewForm.application_id = null
  } catch (err: any) {
    feedback.value = {
      type: 'error',
      message: err.response?._data?.message || 'Failed to submit application review.'
    }
  } finally {
    isReviewing.value = false
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const handleUpdateProfile = async () => {
  isSavingProfile.value = true
  feedback.value = null

  try {
    const res = await api.put('/api/profile/business', {
      company_name: profileForm.company_name,
      phone: profileForm.phone,
      address: profileForm.address,
      logo_url: profileForm.logo_url
    })
    feedback.value = {
      type: 'success',
      message: res.message || '🎉 Corporate profile updated successfully!'
    }
  } catch (err: any) {
    feedback.value = {
      type: 'error',
      message: err.response?._data?.message || 'Failed to update company profile.'
    }
  } finally {
    isSavingProfile.value = false
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

onMounted(() => {
  fetchJobs()
})
</script>
