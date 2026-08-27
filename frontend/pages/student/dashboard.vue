<template>
  <div class="dashboard-body min-h-screen bg-slate-950 text-slate-100 lg:flex font-sans">
    <!-- Sidebar navigation (desktop aside + mobile top bar/drawer) -->
    <StudentSidebar :state="studentDashboardState" @logout="handleLogout" />

    <div class="min-w-0 flex-1">

      <!-- Dashboard Header -->
      <section class="border-b border-white/10 bg-slate-950">
        <div class="w-full px-4 py-6 sm:px-6 lg:px-8">
          <div class="grid gap-6 lg:grid-cols-[1.1fr_0.9fr] lg:items-center">
            <div>
              <div class="flex items-center justify-between gap-4">
                <p class="text-sm font-medium text-slate-400">
                  {{ profileForm.full_name ? `Chào mừng trở lại, ${profileForm.full_name} 👋` : 'Chào mừng bạn trở lại 👋' }}
                </p>
                <NotificationBell align="right" />
              </div>
              <h1 class="mt-1.5 text-2xl font-bold tracking-tight text-white sm:text-3xl">
                {{ studentHero.title }}
              </h1>
              <p class="mt-2 max-w-2xl text-sm leading-relaxed text-slate-400">
                {{ studentHero.description }}
              </p>
              <div class="mt-5 flex flex-wrap gap-3">
                <button
                  @click="handleStudentHeroAction"
                  class="inline-flex items-center justify-center rounded-xl bg-brand-400 px-5 py-2.5 text-sm font-bold text-slate-950 transition-colors hover:bg-brand-300 focus-ring"
                >
                  {{ studentHero.cta }}
                </button>
                <button
                  v-if="activeSection !== 'profile'"
                  @click="activeSection = 'profile'"
                  class="inline-flex items-center justify-center rounded-xl border border-white/10 px-5 py-2.5 text-sm font-semibold text-slate-300 transition-colors hover:bg-white/5 hover:text-white focus-ring"
                >
                  Cập nhật hồ sơ
                </button>
              </div>
            </div>

            <!-- Stats metric grid -->
            <div class="grid grid-cols-2 gap-3.5 sm:grid-cols-4 lg:grid-cols-2">
              <div
                v-for="stat in studentHeroStats"
                :key="stat.label"
                class="rounded-2xl border border-white/10 bg-white/[0.02] p-4 transition-colors hover:border-white/20"
              >
                <div class="flex items-start justify-between gap-2">
                  <div class="min-w-0">
                    <p class="truncate text-[11px] font-semibold uppercase tracking-wider text-slate-500">{{ stat.label }}</p>
                    <p class="mt-1.5 text-2xl font-bold text-white tracking-tight">{{ stat.value }}</p>
                    <p class="mt-1 truncate text-xs font-medium text-slate-500">{{ stat.caption }}</p>
                  </div>
                  <span :class="[stat.iconClass, 'inline-flex h-8 w-8 flex-shrink-0 items-center justify-center rounded-xl']">
                    <svg v-if="stat.icon === 'jobs'" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                    </svg>
                    <svg v-else-if="stat.icon === 'apps'" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                    </svg>
                    <svg v-else-if="stat.icon === 'accepted'" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    <svg v-else class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                    </svg>
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <main class="w-full max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
        <StudentJobsSection :state="studentDashboardState" />
        <StudentSavedJobsSection :state="studentDashboardState" @apply="openApplyModal" />
        <StudentProfileSection :state="studentDashboardState" />
        <StudentApplicationsSection :state="studentDashboardState" />
        <StudentWalletSection :state="studentDashboardState" />
      </main>
    </div>
  </div>

  <StudentDashboardModals
    :state="studentDashboardState"
    @upload-evidence="handleUploadEvidence"
    @upload-completion-proof="handleUploadCompletionProof"
  />
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, onUnmounted, watch } from 'vue'
import StudentJobsSection from '~/components/student/StudentJobsSection.vue'
import StudentSavedJobsSection from '~/components/student/StudentSavedJobsSection.vue'
import StudentProfileSection from '~/components/student/StudentProfileSection.vue'
import StudentApplicationsSection from '~/components/student/StudentApplicationsSection.vue'
import StudentDashboardModals from '~/components/student/StudentDashboardModals.vue'
import StudentWalletSection from '~/components/student/StudentWalletSection.vue'
import StudentSidebar from '~/components/student/StudentSidebar.vue'
import NotificationBell from '~/components/common/NotificationBell.vue'

const { logout, user } = useAuth()
const handleLogout = () => {
  logout()
}

const activeSection = useState<string>('studentDashboardActiveSection', () => 'jobs')

const navItems = [
  { id: 'jobs', name: 'Trang chủ tìm việc' },
  { id: 'saved-jobs', name: 'Việc làm đã lưu' },
  { id: 'profile', name: 'Hồ sơ cá nhân' },
  { id: 'applications', name: 'Đơn ứng tuyển' },
  { id: 'wallet', name: 'Ví thu nhập' }
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
const wallet = ref<any>(null)
const walletTransactions = ref<any[]>([])
const isLoadingWallet = ref(false)
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


const selectedReviewApp = ref<any>(null)
const reviewRating = ref(5)
const reviewComment = ref('')
const isSubmittingReview = ref(false)
const selectedManagedApplication = ref<any | null>(null)
const selectedReviewsApp = ref<any | null>(null)
const reviewsList = ref<any[]>([])
const reviewSummary = ref<{ average_rating: number; total_reviews: number }>({
  average_rating: 0,
  total_reviews: 0
})
const isLoadingReviews = ref(false)
const isReviewsModalOpen = ref(false)
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


const openReviewModal = (app: any) => {
  selectedReviewApp.value = app
  reviewRating.value = 5
  reviewComment.value = ''
}

const submitReview = async () => {
  if (!selectedReviewApp.value) return

  isSubmittingReview.value = true

  try {
    const res = await api.post('/api/reviews', {
      application_id: selectedReviewApp.value.id,
      rating: reviewRating.value,
      comment: reviewComment.value
    })

    showToast(res?.message || 'Đánh giá thành công.')

    selectedReviewApp.value = null
    await fetchApplications()
  } catch (error: any) {
    console.error('Review error:', error)
    showToast(
      error.data?.error ||
      error.response?._data?.error ||
      error.message ||
      'Không thể gửi đánh giá.',
      'error'
    )
  } finally {
    isSubmittingReview.value = false
  }
}

const openManagedApplicationModal = async (app: any) => {
  selectedManagedApplication.value = app
  feedback.value = null
  await fetchReviewsByApplication(app.id)
}

const closeManagedApplicationModal = () => {
  selectedManagedApplication.value = null
}

const fetchReviewsByApplication = async (applicationId: number) => {
  isLoadingReviews.value = true
  try {
    const res = await api.get(`/api/reviews/application/${applicationId}`)
    const list = Array.isArray(res?.reviews) ? res.reviews : []
    reviewsList.value = list

    const total = list.length
    const sum = list.reduce((acc: number, review: any) => acc + Number(review.rating || review.Rating || 0), 0)
    reviewSummary.value = {
      average_rating: total > 0 ? sum / total : 0,
      total_reviews: total
    }
  } catch (error) {
    console.error('Error fetching application reviews:', error)
    reviewsList.value = []
    reviewSummary.value = { average_rating: 0, total_reviews: 0 }
  } finally {
    isLoadingReviews.value = false
  }
}

const fetchReviewsByUser = async (userId: number) => {
  isLoadingReviews.value = true
  try {
    const res = await api.get(`/api/reviews/user/${userId}`)
    reviewsList.value = Array.isArray(res?.reviews) ? res.reviews : []
    reviewSummary.value = {
      average_rating: Number(res?.average_rating || 0),
      total_reviews: Number(res?.total_reviews || 0)
    }
  } catch (error) {
    console.error('Error fetching user reviews:', error)
    reviewsList.value = []
    reviewSummary.value = { average_rating: 0, total_reviews: 0 }
  } finally {
    isLoadingReviews.value = false
  }
}

const openReviewsModal = async (app: any) => {
  selectedReviewsApp.value = app
  isReviewsModalOpen.value = true
  await fetchReviewsByApplication(app.id)
}

const closeReviewsModal = () => {
  isReviewsModalOpen.value = false
  selectedReviewsApp.value = null
}
const companyNameLookup = (source: any): string => {
  if (source?.business?.company_name) {
    return source.business.company_name
  }

  if (source?.company_name) {
    return source.company_name
  }

  const id = Number(source?.business_id ?? source)

  const fallback: Record<number, string> = {
    1: "Công ty TNHH Phần mềm iNET",
    2: "FPT Software",
    3: "VNG Corporation",
    4: "Viettel Digital",
    5: "MISA Joint Stock Company",
    6: "NashTech Vietnam"
  }

  return fallback[id] || `Verified Employer #${id || ''}`
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

const checkIfApplied = (jobId: number): boolean => {
  const list = Array.isArray(applications.value) ? applications.value : []
  return list.some(app => app.job_id === jobId)
}

const filterApplyStatus = ref('all') // 'all' | 'unapplied' | 'applied'

// Computeds for filtering with defensive null-safety & smart sorting
const filteredJobs = computed(() => {
  const list = Array.isArray(jobs.value) ? jobs.value : []
  const result = list.filter(job => {
    if (!job) return false
    const title = (job.title || '').toLowerCase()
    const description = (job.description || '').toLowerCase()
    const location = (job.location || '').toLowerCase()
    const isApplied = checkIfApplied(job.id)

    const matchesSearch = !jobsSearchQuery.value ||
      title.includes(jobsSearchQuery.value.toLowerCase()) ||
      description.includes(jobsSearchQuery.value.toLowerCase())

    const matchesLocation = !jobsLocationQuery.value ||
      location.includes(jobsLocationQuery.value.toLowerCase())

    let matchesApplyStatus = true
    if (filterApplyStatus.value === 'unapplied') {
      matchesApplyStatus = !isApplied
    } else if (filterApplyStatus.value === 'applied') {
      matchesApplyStatus = isApplied
    }

    return matchesSearch && matchesLocation && matchesApplyStatus
  })

  // SẮP XẾP THÔNG MINH:
  // Nhóm 1: Các việc làm CHƯA ỨNG TUYỂN -> đưa lên trên cùng (xếp theo ID giảm dần = mới nhất lên đầu)
  // Nhóm 2: Các việc làm ĐÃ ỨNG TUYỂN -> đẩy xuống dưới cùng để không chiếm chỗ
  return result.sort((a, b) => {
    const aApplied = checkIfApplied(a.id) ? 1 : 0
    const bApplied = checkIfApplied(b.id) ? 1 : 0
    if (aApplied !== bApplied) {
      return aApplied - bApplied // 0 (chưa nộp) trước 1 (đã nộp)
    }
    return (b.id || 0) - (a.id || 0)
  })
})

const filteredApps = computed(() => {
  const list = Array.isArray(applications.value) ? applications.value : []
  return list.filter(app => {
    if (!app) return false
    const jobTitle = (app.job?.title || '').toLowerCase()
    const compName = companyNameLookup(app.job).toLowerCase()

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
      title: profileForm.full_name ? `Hồ sơ cá nhân: ${profileForm.full_name}` : 'Tạo hồ sơ sinh viên ấn tượng',
      description: 'Cập nhật thông tin liên hệ, kỹ năng, ảnh đại diện và CV của bạn trước khi ứng tuyển các cơ hội mới.',
      cta: isEditing.value ? 'Đang chỉnh sửa' : 'Chỉnh sửa hồ sơ'
    }
  }

  if (activeSection.value === 'applications') {
    return {
      title: 'Theo dõi đơn ứng tuyển',
      description: 'Theo dõi tiến trình từng đơn ứng tuyển, phản hồi offer nhận việc và quản lý các ca làm việc tập trung.',
      cta: 'Tìm thêm việc làm'
    }
  }

  return {
    title: 'Tìm việc làm bán thời gian cho sinh viên',
    description: 'Duyệt các công việc đã xác thực, so sánh mức lương và địa điểm, sau đó nộp đơn ứng tuyển nhanh chóng.',
    cta: 'Khám phá công việc'
  }
})

const studentHeroStats = computed(() => [
  {
    label: 'Việc làm mở tuyển',
    value: jobs.value.length,
    caption: 'Đang tuyển dụng',
    icon: 'jobs',
    iconClass: 'bg-brand-400/10 text-brand-300'
  },
  {
    label: 'Đơn ứng tuyển',
    value: applications.value.length,
    caption: 'Tổng số đã nộp',
    icon: 'apps',
    iconClass: 'bg-white/10 text-slate-300'
  },
  {
    label: 'Đã nhận việc',
    value: acceptedApplicationsCount.value,
    caption: `${pendingApplicationsCount.value} đơn chờ duyệt`,
    icon: 'accepted',
    iconClass: 'bg-emerald-400/10 text-emerald-300'
  },
  {
    label: 'Hồ sơ cá nhân',
    value: `${profileReadiness.value}%`,
    caption: 'Mức độ hoàn thiện',
    icon: 'profile',
    iconClass: 'bg-amber-400/10 text-amber-300'
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
  if (jobs.value.length === 0) {
    isLoadingJobs.value = true
  }
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
  } finally {
    isLoadingJobs.value = false
  }
}

const targetAppId = useState<number | null>('targetAppIdFromNotification', () => null)

const checkAndOpenTargetNotifApp = () => {
  if (targetAppId.value && applications.value.length > 0) {
    const foundApp = applications.value.find(a => Number(a.id) === Number(targetAppId.value))
    if (foundApp) {
      selectedManagedApplication.value = foundApp
    }
    targetAppId.value = null
  }
}

watch(applications, () => {
  checkAndOpenTargetNotifApp()
}, { deep: true, immediate: true })

watch(targetAppId, () => {
  checkAndOpenTargetNotifApp()
})

const fetchApplications = async () => {
  if (applications.value.length === 0) {
    isLoadingApps.value = true
  }
  try {
    const res = await api.get('/api/applications/my-applications')
    applications.value = Array.isArray(res)
      ? res
      : (res && Array.isArray(res.data) ? res.data : [])
    checkAndOpenTargetNotifApp()
  } catch (err: any) {
    console.error('Error fetching student applications:', err)
  } finally {
    isLoadingApps.value = false
  }
}
const fetchWallet = async () => {
  if (!wallet.value) {
    isLoadingWallet.value = true
  }

  try {
    const res = await api.get('/api/wallet/me')

    wallet.value = res.wallet || null
    walletTransactions.value = Array.isArray(res.transactions)
      ? res.transactions
      : []
  } catch (err) {
    console.error('Error fetching wallet:', err)
  } finally {
    isLoadingWallet.value = false
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

    const msg = res.message || '🎉 Nộp đơn ứng tuyển thành công! Hồ sơ CV của bạn đã được chuyển tới Nhà tuyển dụng.'
    feedback.value = {
      type: 'success',
      message: msg
    }
    showToast(msg, 'success')
    selectedJobForApply.value = null
    await Promise.all([
      fetchApplications(),
      fetchJobs(),
      fetchMyTickets()
    ])
  } catch (err: any) {
    const errMsg = err.response?._data?.message || err.message || 'Có lỗi xảy ra khi gửi đơn ứng tuyển.'
    feedback.value = {
      type: 'error',
      message: errMsg
    }
    showToast(errMsg, 'error')
  } finally {
    isSubmittingApply.value = false
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
      profileForm.avatar_url = updatedData.avatar_url || updatedData.AvatarUrl || profileForm.avatar_url
      profileForm.cv_url = updatedData.cv_url || updatedData.CvUrl || profileForm.cv_url
    }

    await fetchProfile()

    skillsArray.value = parsedSkills
    isEditing.value = false
    avatarFileSelected.value = null
    avatarPreview.value = null
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

    const res = await api.post('/api/applications/respond-offer', payload)
    showToast(res.message || (responseType === 'accept' ? '🎉 Đã chấp nhận Offer thành công!' : '❌ Đã từ chối Offer.'))
    selectedOffer.value = null
    await Promise.all([
      fetchApplications(),
      fetchJobs(),
      fetchWallet(),
      fetchMyTickets()
    ])
  } catch (err: any) {
    console.error('Lỗi khi gửi phản hồi Offer:', err)
    showToast(err.response?._data?.message || err.message || 'Không thể gửi phản hồi lúc này, vui lòng thử lại!', 'error')
  } finally {
    isResponding.value = false
  }
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

const showCheckOutConfirmModal = ref(false)
const checkOutJobIdTarget = ref<number | null>(null)
const isSubmittingCheckOut = ref(false)

const triggerCheckOutConfirm = (jobId: number) => {
  checkOutJobIdTarget.value = jobId
  showCheckOutConfirmModal.value = true
}

const executeCheckOut = async () => {
  if (!checkOutJobIdTarget.value) return
  const jobId = checkOutJobIdTarget.value
  isSubmittingCheckOut.value = true

  try {
    const res = await api.post('/api/attendance/check-out', { job_id: Number(jobId) })

    delete activeAttendance.value[jobId]
    delete timers.value[jobId]
    localStorage.removeItem(`active_work_${jobId}`)

    showToast(res?.message || '🛑 Check-out kết thúc ca thành công!')
    showCheckOutConfirmModal.value = false
    checkOutJobIdTarget.value = null
  } catch (error: any) {
    showToast(error.response?._data?.error || 'Không thể check-out', 'error')
  } finally {
    isSubmittingCheckOut.value = false
  }
}
const selectedCompletionAppForReport = ref<any | null>(null)
const completionForm = reactive({
  note: '',
  proofUrl: ''
})
const isSubmittingCompletion = ref(false)

const openStudentCompletionModal = (app: any) => {
  let targetApp = app
  if (typeof app !== 'object') {
    targetApp = applications.value.find(a => Number(a.id) === Number(app)) || { id: app }
  }
  selectedCompletionAppForReport.value = targetApp
  completionForm.note = targetApp.completion_note || targetApp.CompletionNote || ''
  completionForm.proofUrl = targetApp.completion_proof_url || targetApp.CompletionProofUrl || ''
}

const submitStudentCompletionReport = async () => {
  if (!selectedCompletionAppForReport.value) return
  if (!completionForm.note.trim()) {
    showToast('Vui lòng nhập nội dung báo cáo kết quả công việc', 'error')
    return
  }

  isSubmittingCompletion.value = true
  try {
    const res = await api.post('/api/applications/student-complete', {
      application_id: selectedCompletionAppForReport.value.id,
      completion_note: completionForm.note,
      completion_proof_url: completionForm.proofUrl
    })

    showToast(res?.message || '🎉 Đã nộp báo cáo hoàn thành công việc! Đang chờ doanh nghiệp đối soát giải ngân.')
    selectedCompletionAppForReport.value = null

    await fetchApplications()
    await fetchWallet()
  } catch (error: any) {
    console.error('Student complete error:', error)
    showToast(
      error.data?.error ||
      error.response?._data?.error ||
      error.message ||
      'Không thể xác nhận hoàn thành.',
      'error'
    )
  } finally {
    isSubmittingCompletion.value = false
  }
}

const isUploadingCompletionProof = ref(false)

const handleUploadCompletionProof = async (event: any) => {
  const file = event.target.files?.[0]
  if (!file) return

  if (file.size > 10 * 1024 * 1024) {
    showToast('File bài nộp không được vượt quá 10MB', 'error')
    return
  }

  isUploadingCompletionProof.value = true
  try {
    const formData = new FormData()
    formData.append('evidence', file)
    const res = await api.post('/api/tickets/upload-evidence', formData)
    if (res.url) {
      completionForm.proofUrl = res.url
      showToast('🎉 Tải file bài nộp hoàn thành thành công!')
    }
  } catch (err: any) {
    const errMsg = err.response?._data?.message || err.message || 'Không thể tải file bài nộp'
    showToast(errMsg, 'error')
  } finally {
    isUploadingCompletionProof.value = false
  }
}

const isDisputeModalOpen = ref(false)
const disputeTargetApp = ref<any | null>(null)
const isSubmittingDispute = ref(false)
const disputeForm = reactive({
  reason: '',
  description: '',
  requested_action: '',
  evidence_url: ''
})

const isUploadingEvidence = ref(false)

const handleUploadEvidence = async (event: any) => {
  const file = event.target.files?.[0]
  if (!file) return

  if (file.size > 10 * 1024 * 1024) {
    showToast('File bằng chứng không được vượt quá 10MB', 'error')
    return
  }

  isUploadingEvidence.value = true
  try {
    const formData = new FormData()
    formData.append('evidence', file)
    const res = await api.post('/api/tickets/upload-evidence', formData)
    if (res.url) {
      disputeForm.evidence_url = res.url
      showToast('🎉 Tải file bằng chứng thành công!')
    }
  } catch (err: any) {
    const errMsg = err.response?._data?.message || err.message || 'Không thể tải file bằng chứng'
    showToast(errMsg, 'error')
  } finally {
    isUploadingEvidence.value = false
  }
}

const openDisputeModal = (app: any) => {
  disputeTargetApp.value = app
  disputeForm.reason = ''
  disputeForm.description = ''
  disputeForm.requested_action = ''
  disputeForm.evidence_url = ''
  isDisputeModalOpen.value = true
}

const showDisputeErrorModal = ref(false)
const disputeErrorMsg = ref('')

const openDisputeError = (msg: string) => {
  disputeErrorMsg.value = msg
  showDisputeErrorModal.value = true
}

const myTickets = ref<any[]>([])
const isLoadingTickets = ref(false)
const showReappealModal = ref(false)
const selectedTicketForView = ref<any | null>(null)
const isSubmittingReappeal = ref(false)
const reappealForm = reactive({
  reason: ''
})

const fetchMyTickets = async () => {
  if (myTickets.value.length === 0) {
    isLoadingTickets.value = true
  }
  try {
    const res = await api.get('/api/tickets/my-tickets')
    myTickets.value = res.data || []
  } catch (err) {
    console.error('Fetch tickets error:', err)
  } finally {
    isLoadingTickets.value = false
  }
}

const openReappealModal = (ticket: any) => {
  selectedTicketForView.value = ticket
  reappealForm.reason = ''
  showReappealModal.value = true
}

const submitReappealTicket = async () => {
  if (!selectedTicketForView.value || !reappealForm.reason.trim()) {
    showToast('Vui lòng nhập lý do yêu cầu tái xem xét phán quyết', 'error')
    return
  }
  if (reappealForm.reason.trim().length < 10) {
    showToast('Lý do yêu cầu tái xem xét phải có ít nhất 10 ký tự', 'error')
    return
  }
  isSubmittingReappeal.value = true
  try {
    const ticketId = selectedTicketForView.value.id || selectedTicketForView.value.ticket_id || selectedTicketForView.value.ID
    const res = await api.post(`/api/tickets/${ticketId}/reappeal`, {
      reason: reappealForm.reason.trim()
    })
    showToast(res.message || '🎉 Đã gửi yêu cầu tái xem xét thành công!')
    showReappealModal.value = false
    await fetchMyTickets()
  } catch (err: any) {
    const errMsg = err.response?._data?.message || err.message || 'Không thể gửi yêu cầu tái xem xét'
    showToast(errMsg, 'error')
  } finally {
    isSubmittingReappeal.value = false
  }
}

const submitDisputeTicket = async () => {
  if (!disputeTargetApp.value || !disputeForm.reason || !disputeForm.description.trim()) {
    openDisputeError('Vui lòng chọn lý do và nhập mô tả chi tiết khiếu nại.')
    return
  }
  isSubmittingDispute.value = true

  try {
    const res = await api.post('/api/tickets', {
      application_id: Number(disputeTargetApp.value.id),
      reason: disputeForm.reason,
      description: disputeForm.description.trim(),
      requested_action: disputeForm.requested_action.trim(),
      evidence_url: disputeForm.evidence_url.trim()
    })

    showToast(res.message || '🎉 Gửi khiếu nại thành công!')
    isDisputeModalOpen.value = false
    disputeTargetApp.value = null
    await fetchMyTickets()
  } catch (err: any) {
    const errMsg = err.response?._data?.message || err.message || 'Không thể gửi khiếu nại'
    showToast(errMsg, 'error')
    openDisputeError(errMsg)
  } finally {
    isSubmittingDispute.value = false
  }
}

const handleStudentComplete = async (app: any) => {
  if (typeof app === 'object' && app?.id) {
    openStudentCompletionModal(app)
  } else {
    openStudentCompletionModal({ id: app })
  }
}
// Vòng đời Hook duy nhất bọc toàn bộ logic
const studentDashboardState = reactive({
  activeSection,
  navItems,
  handleLogout,
  profileReadiness,
  feedback,
  jobsSearchQuery,
  jobsLocationQuery,
  filterCategory,
  filterJobType,
  filterMinSalary,
  filterApplyStatus,
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
  applications,
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
  triggerCheckOutConfirm,
  executeCheckOut,
  showCheckOutConfirmModal,
  isSubmittingCheckOut,
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
  selectedCompletionAppForReport,
  completionForm,
  isSubmittingCompletion,
  isUploadingCompletionProof,
  openStudentCompletionModal,
  submitStudentCompletionReport,
  isDisputeModalOpen,
  disputeTargetApp,
  isSubmittingDispute,
  disputeForm,
  isUploadingEvidence,
  myTickets,
  isLoadingTickets,
  showReappealModal,
  selectedTicketForView,
  isSubmittingReappeal,
  reappealForm,
  openReappealModal,
  submitReappealTicket,
  fetchMyTickets,
  openDisputeModal,
  submitDisputeTicket,
  showDisputeErrorModal,
  disputeErrorMsg,
  selectedChatApp,
  toast,
  wallet,
  walletTransactions,
  isLoadingWallet,
  selectedReviewApp,
  reviewRating,
  reviewComment,
  isSubmittingReview,
  openReviewModal,
  submitReview,
  selectedManagedApplication,
  openManagedApplicationModal,
  closeManagedApplicationModal,
  selectedReviewsApp,
  reviewsList,
  reviewSummary,
  isLoadingReviews,
  isReviewsModalOpen,
  openReviewsModal,
  closeReviewsModal,
  fetchReviewsByApplication,
  fetchReviewsByUser,
  fetchWallet
})


let autoSyncStudentInterval: any = null

const syncStudentDashboardData = async () => {
  try {
    await Promise.all([
      fetchJobs(),
      fetchApplications(),
      fetchProfile(),
      fetchWallet(),
      fetchMyTickets()
    ])
  } catch (e) {
    // silent background sync
  }
}

const handleStudentVisibilityChange = () => {
  if (document.visibilityState === 'visible') {
    syncStudentDashboardData()
  }
}

onUnmounted(() => {
  if (timerInterval) clearInterval(timerInterval)
  if (autoSyncStudentInterval) clearInterval(autoSyncStudentInterval)
  if (import.meta.client) {
    document.removeEventListener('visibilitychange', handleStudentVisibilityChange)
  }
})

onMounted(async () => {
  await syncStudentDashboardData()

  if (import.meta.client) {
    document.addEventListener('visibilitychange', handleStudentVisibilityChange)

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
