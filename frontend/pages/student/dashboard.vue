<template>
  <div class="dashboard-body min-h-screen">
    <section class="border-b border-slate-800 bg-gradient-to-br from-slate-950 via-slate-900 to-cyan-950">
      <div class="w-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-5">
        <div class="grid gap-5 lg:grid-cols-[0.95fr_1.05fr] lg:items-center">
          <div>
            <div class="flex flex-wrap items-center gap-2">
              <span class="inline-flex items-center rounded-full bg-cyan-400/10 px-3 py-1 text-xs font-bold uppercase tracking-wide text-cyan-200 ring-1 ring-cyan-400/30">
                Student Workspace
              </span>
              <span class="text-xs font-semibold text-slate-400">Career dashboard</span>
            </div>
            <h1 class="mt-2 text-2xl sm:text-3xl font-extrabold tracking-tight text-white">
              {{ studentHero.title }}
            </h1>
            <p class="mt-2 max-w-2xl text-sm font-semibold leading-6 text-slate-300">
              {{ studentHero.description }}
            </p>
            <div class="mt-4 flex flex-wrap gap-3">
              <button
                @click="handleStudentHeroAction"
                class="inline-flex items-center justify-center rounded-lg bg-cyan-400 px-4 py-2.5 text-sm font-bold text-slate-950 shadow-sm shadow-cyan-950/30 transition-colors hover:bg-cyan-300 focus-ring"
              >
                {{ studentHero.cta }}
              </button>
              <button
                v-if="activeSection !== 'profile'"
                @click="activeSection = 'profile'"
                class="inline-flex items-center justify-center rounded-lg border border-white/10 bg-white/10 px-4 py-2.5 text-sm font-bold text-white transition-colors hover:bg-white/15 focus-ring"
              >
                Improve profile
              </button>
            </div>
          </div>

          <div class="grid grid-cols-2 gap-3">
            <div
              v-for="stat in studentHeroStats"
              :key="stat.label"
              class="rounded-xl border border-white/10 bg-white/[0.07] px-4 py-3 shadow-sm shadow-slate-950/20 backdrop-blur ring-1 ring-white/10"
            >
              <div class="flex items-start justify-between gap-3">
                <div>
                  <p class="text-xs font-bold uppercase tracking-wide text-slate-400">{{ stat.label }}</p>
                  <p class="mt-1 text-2xl font-extrabold text-white">{{ stat.value }}</p>
                  <p class="mt-1 text-xs font-semibold text-slate-300">{{ stat.caption }}</p>
                </div>
                <span :class="[stat.iconClass, 'inline-flex h-9 w-9 flex-shrink-0 items-center justify-center rounded-lg']">
                  <svg v-if="stat.icon === 'jobs'" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                  </svg>
                  <svg v-else-if="stat.icon === 'apps'" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                  </svg>
                  <svg v-else-if="stat.icon === 'accepted'" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  <svg v-else class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                  </svg>
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <main class="w-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <StudentJobsSection :state="studentDashboardState" />
      <StudentProfileSection :state="studentDashboardState" />
      <StudentApplicationsSection :state="studentDashboardState" />
    </main>
  </div>

  <StudentDashboardModals :state="studentDashboardState" />
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, onUnmounted, watch } from 'vue'
import StudentJobsSection from '~/components/student/StudentJobsSection.vue'
import StudentProfileSection from '~/components/student/StudentProfileSection.vue'
import StudentApplicationsSection from '~/components/student/StudentApplicationsSection.vue'
import StudentDashboardModals from '~/components/student/StudentDashboardModals.vue'

definePageMeta({
  middleware: 'auth'
})

const activeSection = useState<string>('studentDashboardActiveSection', () => 'jobs')

const navItems = [
  { id: 'jobs', name: 'Dashboard' },
  { id: 'profile', name: 'Profile' },
  { id: 'applications', name: 'My Applications' }
]

const filterSearch = ref('')
const filterLocation = ref('all')
const filterCategory = ref('all')
const filterJobType = ref('all')
const filterMinSalary = ref('')

// Các biến điều khiển trạng thái Modal Cover Note
const selectedJobForApply = ref<any>(null)
const coverNoteText = ref('')
const isSubmittingApply = ref(false)

const resetFilters = () => {
  jobsSearchQuery.value = ''
  jobsLocationQuery.value = ''
  filterCategory.value = 'all'
  filterJobType.value = 'all'
  filterMinSalary.value = ''
  fetchJobs()
}

// Các lựa chọn địa điểm hiển thị lên giao diện
const locationsList = [
  { value: 'all', label: 'All locations' },
  { value: 'hanoi', label: 'Hanoi' },
  { value: 'hcm', label: 'Ho Chi Minh City' },
  { value: 'danang', label: 'Da Nang' }
]

// State setup
const api = useApi()
const jobs = ref<any[]>([])
const applications = ref<any[]>([])
const isLoadingJobs = ref(false)
const isLoadingApps = ref(false)
const isApplying = ref<number | null>(null)
const isSavingProfile = ref(false)
const feedback = ref<{ type: 'success' | 'error'; message: string } | null>(null)

// 🌟 TRẠNG THÁI UI PROFILE MỚI
const isEditing = ref(false)
const skillsText = ref('')
const skillsArray = ref<string[]>([])
const avatarFileSelected = ref<File | null>(null)
const avatarPreview = ref<string | null>(null)
const cvFileSelected = ref<File | null>(null)

// Search Query filters
const jobsSearchQuery = ref('')
const jobsLocationQuery = ref('')
const appSearchQuery = ref('')
const appStatusFilter = ref('all')
const jobsPage = ref(1)
const jobsPageSize = 6
const applicationsPage = ref(1)
const applicationsPageSize = 8

const appIdToCancel = ref<number | null>(null)
const isCancellingApp = ref<number | null>(null)

// Hàm kích hoạt mở Modal Hủy ứng tuyển
const triggerCancelConfirm = (id: number) => {
  appIdToCancel.value = id
  feedback.value = null
}

// Profile update form reactive state
const profileForm = reactive({
  full_name: '',
  phone: '',
  gender: '',
  avatar_url: '',
  skills: '',
  cv_url: ''
})

// Company Seed Lookup
const companyLookup: Record<number, string> = {
  1: "FPT Software",
  2: "VNG Corporation",
  3: "Viettel Digital",
  4: "MISA Joint Stock Company",
  5: "NashTech Vietnam"
}

const companyNameLookup = (id: number): string => {
  return companyLookup[id] || `Verified Employer #${id || ''}`
}

// Format Application Date
const formatDate = (dateVal: any): string => {
  return 'Jun 24, 2026'
}

// Status Badges Style config
const statusBadgeClass = (status: string): string => {
  const norm = (status || '').toLowerCase()
  if (norm === 'approved' || norm === 'offer_accepted') return 'bg-emerald-50 border-emerald-200 text-emerald-700'
  if (norm === 'student_completed') return 'bg-amber-50 border-amber-200 text-amber-700'
  if (norm === 'paid') return 'bg-cyan-50 border-cyan-200 text-cyan-700'
  if (norm === 'rejected') return 'bg-rose-50 border-rose-200 text-rose-700'
  return 'bg-slate-50 border-slate-200 text-slate-700'
}

// Computeds for filtering with defensive null-safety
const filteredJobs = computed(() => {
  const list = Array.isArray(jobs.value) ? jobs.value : []
  return list.filter(job => {
    if (!job) return false
    const title = (job.title || '').toLowerCase()
    const description = (job.description || '').toLowerCase()
    const location = (job.location || '').toLowerCase()

    const matchesSearch = !jobsSearchQuery.value ||
      title.includes(jobsSearchQuery.value.toLowerCase()) ||
      description.includes(jobsSearchQuery.value.toLowerCase())

    const matchesLocation = !jobsLocationQuery.value ||
      location.includes(jobsLocationQuery.value.toLowerCase())

    return matchesSearch && matchesLocation
  })
})

const filteredApps = computed(() => {
  const list = Array.isArray(applications.value) ? applications.value : []
  return list.filter(app => {
    if (!app) return false
    const jobTitle = (app.job?.title || '').toLowerCase()
    const businessId = app.job?.business_id || 0
    const compName = companyNameLookup(businessId).toLowerCase()

    const matchesSearch = !appSearchQuery.value ||
      jobTitle.includes(appSearchQuery.value.toLowerCase()) ||
      compName.includes(appSearchQuery.value.toLowerCase())

    const matchesStatus = appStatusFilter.value === 'all' ||
      (app.status || '').toLowerCase() === appStatusFilter.value.toLowerCase()

    return matchesSearch && matchesStatus
  })
})

const paginatedJobs = computed(() => {
  const start = (jobsPage.value - 1) * jobsPageSize
  return filteredJobs.value.slice(start, start + jobsPageSize)
})

const paginatedApps = computed(() => {
  const start = (applicationsPage.value - 1) * applicationsPageSize
  return filteredApps.value.slice(start, start + applicationsPageSize)
})

watch(filteredJobs, () => {
  jobsPage.value = 1
})

watch(filteredApps, () => {
  applicationsPage.value = 1
})

const acceptedApplicationsCount = computed(() => {
  return applications.value.filter(app => {
    const status = app.status?.toLowerCase()
    return status === 'approved' || status === 'offer_accepted'
  }).length
})

const pendingApplicationsCount = computed(() => {
  return applications.value.filter(app => {
    const status = app.status?.toLowerCase()
    return status === 'applied' || status === 'pending'
  }).length
})

const profileReadiness = computed(() => {
  const fields = [
    profileForm.full_name,
    profileForm.phone,
    profileForm.gender,
    profileForm.avatar_url || avatarPreview.value,
    skillsArray.value.length > 0 || skillsText.value,
    profileForm.cv_url || cvFileSelected.value
  ]
  const completed = fields.filter(Boolean).length
  return Math.round((completed / fields.length) * 100)
})

const studentHero = computed(() => {
  if (activeSection.value === 'profile') {
    return {
      title: profileForm.full_name ? `${profileForm.full_name}'s profile` : 'Build a stronger student profile',
      description: 'Keep your contact details, skills, avatar, and CV ready before applying to new opportunities.',
      cta: isEditing.value ? 'Editing profile' : 'Edit profile'
    }
  }

  if (activeSection.value === 'applications') {
    return {
      title: 'Track your applications',
      description: 'Follow every application, respond to offers, and manage active work sessions from one place.',
      cta: 'Find more jobs'
    }
  }

  return {
    title: 'Find student-friendly jobs',
    description: 'Browse verified roles, compare salary and location, then apply with your QuickWork profile.',
    cta: 'Browse jobs'
  }
})

const studentHeroStats = computed(() => [
  {
    label: 'Open jobs',
    value: jobs.value.length,
    caption: 'Available now',
    icon: 'jobs',
    iconClass: 'bg-brand-50 text-brand-700'
  },
  {
    label: 'Applications',
    value: applications.value.length,
    caption: 'Submitted total',
    icon: 'apps',
    iconClass: 'bg-sky-50 text-sky-700'
  },
  {
    label: 'Accepted',
    value: acceptedApplicationsCount.value,
    caption: `${pendingApplicationsCount.value} pending`,
    icon: 'accepted',
    iconClass: 'bg-emerald-50 text-emerald-700'
  },
  {
    label: 'Profile',
    value: `${profileReadiness.value}%`,
    caption: 'Readiness score',
    icon: 'profile',
    iconClass: 'bg-amber-50 text-amber-700'
  }
])

const handleStudentHeroAction = () => {
  if (activeSection.value === 'profile') {
    isEditing.value = true
    return
  }

  activeSection.value = 'jobs'
}

// API Operations
const fetchJobs = async () => {
  isLoadingJobs.value = true
  try {
    const queryParams: Record<string, string> = {}
    if (jobsSearchQuery.value && jobsSearchQuery.value.trim()) {
      queryParams.search = jobsSearchQuery.value.trim()
    }
    if (jobsLocationQuery.value && jobsLocationQuery.value.trim()) {
      queryParams.location = jobsLocationQuery.value.trim()
    }
    if (filterCategory.value && filterCategory.value !== 'all') {
      queryParams.category = filterCategory.value
    }
    if (filterJobType.value && filterJobType.value !== 'all') {
      queryParams.job_type = filterJobType.value
    }
    if (filterMinSalary.value) {
      queryParams.max_salary = filterMinSalary.value
    }

    const res = await api.get('/api/jobs', { params: queryParams })
    jobs.value = res.data || []
} catch (err) {
    console.error('Error fetching filtered jobs:', err)
  } finally {  // Đã xóa chữ gõ nhầm
    isLoadingJobs.value = false
  }
}

const fetchApplications = async () => {
  isLoadingApps.value = true
  try {
    const res = await api.get('/api/applications/my-applications')
    applications.value = Array.isArray(res)
      ? res
      : (res && Array.isArray(res.data) ? res.data : [])
  } catch (err: any) {
    console.error('Error fetching student applications:', err)
  } finally {
    isLoadingApps.value = false
  }
}

const fetchProfile = async () => {
  try {
    const res = await api.get('/api/profile/student')
    const data = res && res.data ? res.data : res

    if (data) {
      profileForm.full_name = data.full_name || ''
      profileForm.phone = data.phone || ''
      profileForm.gender = data.gender || ''
      profileForm.avatar_url = data.avatar_url || ''
      profileForm.cv_url = data.cv_url || ''
      profileForm.skills = data.skills || ''
      currentUserId.value = data.user_id || data.id

      if (data.skills) {
        try {
          const parsed = JSON.parse(data.skills)
          if (Array.isArray(parsed)) {
            skillsArray.value = parsed
            skillsText.value = parsed.join(', ')
          } else {
            skillsText.value = data.skills
            skillsArray.value = data.skills.split(',').map((s: string) => s.trim())
          }
        } catch {
          skillsText.value = data.skills
          skillsArray.value = data.skills.split(',').map((s: string) => s.trim())
        }
      }
    }
  } catch (err: any) {
    console.error('Error fetching profile details:', err)
  }
}

const handleApply = (job: any) => {
  selectedJobForApply.value = job
  coverNoteText.value = ''
  feedback.value = null
}

const submitApplication = async () => {
  if (!selectedJobForApply.value) return
  isSubmittingApply.value = true
  feedback.value = null

  try {
    const res = await api.post('/api/jobs/apply', {
      job_id: selectedJobForApply.value.id,
      cover_note: coverNoteText.value.trim()
    })

    feedback.value = {
      type: 'success',
      message: res.message || '🚀 Applied successfully!'
    }
    selectedJobForApply.value = null
    await fetchApplications()
  } catch (err: any) {
    feedback.value = {
      type: 'error',
      message: err.response?._data?.message || 'Failed to submit application.'
    }
  } finally {
    isSubmittingApply.value = false
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const isChatModalOpen = ref(false)
const selectedChatApp = ref<any>(null)
const currentUserId = ref<number | null>(null)

const openChatModal = (app: any) => {
  selectedChatApp.value = app
  isChatModalOpen.value = true
}

const confirmCancelApplication = async () => {
  if (!appIdToCancel.value) return
  const targetId = appIdToCancel.value
  isCancellingApp.value = targetId
  feedback.value = null
  appIdToCancel.value = null

  try {
    const res = await api.post(`/api/applications/${targetId}/cancel`)
    feedback.value = {
      type: 'success',
      message: res.message || '❌ Đã hủy đơn ứng tuyển thành công.'
    }
    await fetchApplications()
  } catch (err: any) {
    console.error('Error cancelling application:', err)
    feedback.value = {
      type: 'error',
      message: err.response?._data?.message || 'Failed to cancel application.'
    }
  } finally {
    isCancellingApp.value = null
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const onAvatarFileChange = (e: Event) => {
  const target = e.target as HTMLInputElement
  const file = target.files?.[0]
  if (file) {
    avatarFileSelected.value = file
    avatarPreview.value = URL.createObjectURL(file)
  }
}

const onCvFileChange = (e: Event) => {
  const target = e.target as HTMLInputElement
  const file = target.files?.[0]
  if (file) {
    if (file.type !== 'application/pdf') {
      alert('Please select a valid PDF file for your CV!')
      return
    }
    cvFileSelected.value = file
  }
}

const handleUpdateProfile = async () => {
  isSavingProfile.value = true
  feedback.value = null

  try {
    const formData = new FormData()
    formData.append('full_name', profileForm.full_name)
    formData.append('phone', profileForm.phone)
    formData.append('gender', profileForm.gender)

    const parsedSkills = skillsText.value.split(',').map(s => s.trim()).filter(s => s !== '')
    formData.append('skills', JSON.stringify(parsedSkills))

    if (avatarFileSelected.value) formData.append('avatar', avatarFileSelected.value)
    if (cvFileSelected.value) formData.append('cv', cvFileSelected.value)

    const res = await api.put('/api/profile/student', formData)

    feedback.value = {
      type: 'success',
      message: res.message || '🎉 Profile updated successfully!'
    }

    const updatedData = res.data ? res.data : res
    if (updatedData) {
      profileForm.avatar_url = updatedData.AvatarUrl || updatedData.avatar_url || profileForm.avatar_url
      profileForm.cv_url = updatedData.CvUrl || updatedData.cv_url || profileForm.cv_url
    }

    skillsArray.value = parsedSkills
    isEditing.value = false
    avatarFileSelected.value = null
    cvFileSelected.value = null
  } catch (err: any) {
    feedback.value = {
      type: 'error',
      message: err.response?._data?.message || 'Failed to update profile.'
    }
  } finally {
    isSavingProfile.value = false
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const selectedOffer = ref<any>(null)
const isResponding = ref(false)

const openOfferModal = (app: any) => {
  selectedOffer.value = app
}

const handleOfferResponse = async (responseType: 'accept' | 'decline') => {
  if (!selectedOffer.value) return
  isResponding.value = true
  try {
    const payload = {
      application_id: selectedOffer.value.id,
      response: responseType
    }

    await api.post('/api/applications/respond-offer', payload)
    await fetchApplications()
    selectedOffer.value = null
  } catch (err) {
    console.error('Lỗi khi gửi phản hồi Offer:', err)
    alert('Không thể gửi phản hồi lúc này, vui lòng thử lại!')
  } finally {
    isResponding.value = false
  }
}

const checkIfApplied = (jobId: number): boolean => {
  const list = Array.isArray(applications.value) ? applications.value : []
  return list.some(app => app.job_id === jobId)
}

// ⏰ QUẢN LÝ ATTENDANCE & REALTIME TIMERS
const activeAttendance = ref<{ [key: number]: string }>({})
const timers = ref<{ [key: number]: string }>({})
let timerInterval: any = null

const isWorking = (jobId: number) => {
  return !!activeAttendance.value[jobId]
}

const getTimer = (jobId: number) => {
  return timers.value[jobId] || '00:00:00'
}

const updateTimers = () => {
  for (const jobId in activeAttendance.value) {
    const startTime = new Date(activeAttendance.value[jobId]).getTime()
    const now = new Date().getTime()
    const diff = now - startTime

    if (diff > 0) {
      const hours = Math.floor(diff / (1000 * 60 * 60)).toString().padStart(2, '0')
      const minutes = Math.floor((diff % (1000 * 60 * 60)) / (1000 * 60)).toString().padStart(2, '0')
      const seconds = Math.floor((diff % (1000 * 60)) / 1000).toString().padStart(2, '0')
      timers.value[jobId] = `${hours}:${minutes}:${seconds}`
    }
  }
}
// 🔔 Hệ thống Toast tùy biến thay thế alert()
const toast = ref<{ show: boolean; message: string; type: 'success' | 'error' }>({
  show: false,
  message: '',
  type: 'success'
})

// Hàm kích hoạt hiển thị Toast và tự động đóng sau 3 giây
const showToast = (message: string, type: 'success' | 'error' = 'success') => {
  toast.value = { show: true, message, type }
  setTimeout(() => {
    toast.value.show = false
  }, 3000)
}
// Gọi API Check-in chuẩn hóa theo useApi của Chanh
const handleCheckIn = async (jobId: number) => {
  try {
    const res = await api.post('/api/attendance/check-in', { job_id: Number(jobId) })

    // Lưu vào state
    activeAttendance.value[jobId] = new Date().toISOString()
    // 🌟 Lưu thêm vào localStorage để F5 không mất trạng thái nút
    localStorage.setItem(`active_work_${jobId}`, activeAttendance.value[jobId])

    showToast(res?.message || '⚡ Check-in thành công!')
  } catch (error: any) {
    showToast(error.response?._data?.error || 'Không thể check-in', 'error')
  }
}

const handleCheckOut = async (jobId: number) => {
  if (!confirm('Bạn có chắc chắn muốn kết thúc ca làm việc?')) return
  try {
    const res = await api.post('/api/attendance/check-out', { job_id: Number(jobId) })

    delete activeAttendance.value[jobId]
    delete timers.value[jobId]
    // 🌟 Xóa khỏi localStorage khi ra ca
    localStorage.removeItem(`active_work_${jobId}`)

    showToast(res?.message || '🛑 Check-out thành công!')
  } catch (error: any) {
    showToast(error.response?._data?.error || 'Không thể check-out', 'error')
  }
}
const handleStudentComplete = async (applicationId: number) => {
  try {
    const res = await api.post('/api/applications/student-complete', {
      application_id: applicationId
    })

    showToast(res?.message || 'Bạn đã xác nhận hoàn thành công việc.')

    await fetchApplications()
  } catch (error: any) {
    console.error('Student complete error:', error)

    showToast(
      error.data?.error ||
      error.response?._data?.error ||
      error.message ||
      'Không thể xác nhận hoàn thành.',
      'error'
    )
  }
}
// Vòng đời Hook duy nhất bọc toàn bộ logic
const studentDashboardState = reactive({
  activeSection,
  navItems,
  feedback,
  jobsSearchQuery,
  jobsLocationQuery,
  filterCategory,
  filterJobType,
  filterMinSalary,
  resetFilters,
  fetchJobs,
  isLoadingJobs,
  filteredJobs,
  paginatedJobs,
  jobsPage,
  jobsPageSize,
  companyNameLookup,
  checkIfApplied,
  handleApply,
  isApplying,
  isEditing,
  profileForm,
  skillsArray,
  avatarPreview,
  onAvatarFileChange,
  skillsText,
  onCvFileChange,
  isSavingProfile,
  handleUpdateProfile,
  isLoadingApps,
  filteredApps,
  paginatedApps,
  applicationsPage,
  applicationsPageSize,
  appSearchQuery,
  appStatusFilter,
  openChatModal,
  openOfferModal,
  triggerCancelConfirm,
  isCancellingApp,
  formatDate,
  statusBadgeClass,
  isWorking,
  getTimer,
  handleCheckIn,
  handleCheckOut,
  selectedJobForApply,
  coverNoteText,
  submitApplication,
  isSubmittingApply,
  appIdToCancel,
  confirmCancelApplication,
  selectedOffer,
  currentUserId,
  handleOfferResponse,
  isResponding,
  isChatModalOpen,
  handleStudentComplete,
  selectedChatApp,
  toast
})

onMounted(() => {
  fetchJobs()
  fetchApplications()
  fetchProfile()
  timerInterval = setInterval(updateTimers, 1000)
})

onUnmounted(() => {
  if (timerInterval) clearInterval(timerInterval)
})
onMounted(() => {
  fetchJobs()
  fetchApplications()
  fetchProfile()

  // 🌟 Khôi phục các ca làm việc đang mở từ localStorage
  if (import.meta.client) {
    applications.value.forEach(app => {
      if (app.job_id) {
        const savedTime = localStorage.getItem(`active_work_${app.job_id}`)
        if (savedTime) {
          activeAttendance.value[app.job_id] = savedTime
        }
      }
    })
  }

  timerInterval = setInterval(updateTimers, 1000)
})
</script>
