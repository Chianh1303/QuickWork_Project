<template>
  <div class="dashboard-body min-h-screen bg-slate-950 text-slate-100 lg:flex font-sans">
    <!-- Sidebar navigation (desktop) -->
    <aside class="hidden lg:flex lg:w-64 lg:flex-shrink-0 lg:flex-col lg:border-r lg:border-indigo-500/15 lg:bg-slate-950/90 lg:px-4 lg:py-6 backdrop-blur-xl sticky top-0 h-screen overflow-y-auto z-30">
      <!-- Enterprise Platform Logo Badge -->
      <div class="px-3 pb-6 border-b border-indigo-500/10 mb-4">
        <button @click="activeSection = 'jobs'" class="flex items-center gap-2.5 group text-left w-full">
          <div class="h-9 w-9 rounded-xl bg-gradient-to-tr from-indigo-500 to-emerald-400 flex items-center justify-center text-slate-950 font-black text-lg shadow-lg shadow-indigo-500/30 group-hover:scale-105 transition-transform">
            QW
          </div>
          <div>
            <span class="text-[10px] font-black uppercase tracking-wider text-indigo-400">Enterprise AI</span>
            <p class="text-sm font-extrabold text-white tracking-tight group-hover:text-indigo-300 transition-colors">QuickWork Portal</p>
          </div>
        </button>
      </div>

      <nav class="flex flex-1 flex-col gap-1.5">
        <button
          v-for="item in navItems"
          :key="item.id"
          @click="activeSection = item.id"
          :class="[
            activeSection === item.id
              ? 'bg-gradient-to-r from-indigo-500/20 to-emerald-500/10 text-indigo-200 border-l-4 border-indigo-400 shadow-md shadow-indigo-500/10 font-extrabold'
              : 'text-slate-400 hover:bg-white/[0.04] hover:text-white font-medium',
            'flex items-center gap-3 rounded-xl px-3.5 py-3 text-sm transition-all text-left'
          ]"
        >
          <svg v-if="item.id === 'jobs'" class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
          </svg>
          <svg v-else-if="item.id === 'profile'" class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
          </svg>
          <svg v-else-if="item.id === 'applications'" class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          <svg v-else class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-2m0-6h4v6h-4a3 3 0 010-6z" />
          </svg>
          <div class="flex items-center justify-between w-full">
            <span>{{ item.name }}</span>
            <span
              v-if="item.id === 'applications' && filteredApps.length > 0"
              class="rounded-full bg-indigo-500/20 px-2 py-0.5 text-[10px] font-black text-indigo-300 ring-1 ring-indigo-500/30"
            >
              {{ filteredApps.length }}
            </span>
            <span
              v-else-if="item.id === 'jobs' && filteredJobs.length > 0"
              class="rounded-full bg-emerald-500/15 px-2 py-0.5 text-[10px] font-black text-emerald-300 ring-1 ring-emerald-500/30"
            >
              {{ filteredJobs.length }}
            </span>
          </div>
        </button>
      </nav>

      <!-- Bottom Profile & Logout Box -->
      <div class="mt-auto space-y-3 pt-4 border-t border-indigo-500/10">
        <div class="rounded-2xl border border-indigo-500/20 bg-gradient-to-b from-indigo-950/40 to-slate-900/60 p-3.5 ring-1 ring-indigo-500/10">
          <div class="flex items-center justify-between mb-2">
            <p class="truncate text-xs font-bold text-white max-w-[130px]">{{ profileForm.full_name || 'Hồ sơ Sinh viên' }}</p>
            <span class="rounded-full bg-emerald-400/10 px-2 py-0.5 text-[10px] font-extrabold text-emerald-300 ring-1 ring-emerald-400/30">
              {{ profileReadiness }}% ATS
            </span>
          </div>
          <div class="h-1.5 w-full overflow-hidden rounded-full bg-slate-800">
            <div
              class="h-full rounded-full bg-gradient-to-r from-indigo-500 to-emerald-400 transition-all duration-500"
              :style="{ width: profileReadiness + '%' }"
            ></div>
          </div>
        </div>

        <!-- Integrated Logout Button -->
        <button
          @click="handleLogout"
          class="w-full flex items-center justify-center gap-2 rounded-xl border border-rose-500/20 bg-rose-500/10 px-4 py-2.5 text-xs font-extrabold text-rose-300 hover:bg-rose-500/20 hover:text-rose-200 transition-all"
        >
          <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
          </svg>
          <span>Đăng xuất tài khoản</span>
        </button>
      </div>
    </aside>

    <div class="min-w-0 flex-1">
      <!-- Section switcher (mobile / tablet) -->
      <div class="flex gap-2 overflow-x-auto border-b border-indigo-500/15 bg-slate-950/80 px-4 py-3 lg:hidden">
        <button
          v-for="item in navItems"
          :key="item.id"
          @click="activeSection = item.id"
          :class="[
            activeSection === item.id ? 'bg-gradient-to-r from-indigo-500 to-emerald-500 text-white font-extrabold shadow-lg shadow-indigo-500/20' : 'bg-white/5 text-slate-300',
            'flex-shrink-0 rounded-xl px-4 py-2 text-xs font-semibold transition-colors'
          ]"
        >
          {{ item.name }}
        </button>
      </div>

      <!-- Hero Header Banner -->
      <section class="border-b border-indigo-500/15 bg-gradient-to-r from-slate-950 via-indigo-950/40 to-slate-950">
        <div class="w-full px-4 py-6 sm:px-6 lg:px-8">
          <div class="grid gap-6 lg:grid-cols-[1.1fr_0.9fr] lg:items-center">
            <div>
              <span class="inline-flex rounded-full bg-indigo-500/10 px-3 py-1 text-xs font-black uppercase tracking-wider text-indigo-300 ring-1 ring-indigo-500/30">
                Enterprise AI Student Portal
              </span>
              <h1 class="mt-3 text-2xl font-extrabold tracking-tight text-white sm:text-3xl lg:text-4xl">
                {{ studentHero.title }}
              </h1>
              <p class="mt-2 max-w-2xl text-sm leading-relaxed text-slate-300 font-medium">
                {{ studentHero.description }}
              </p>
              <div class="mt-5 flex flex-wrap gap-3">
                <button
                  @click="handleStudentHeroAction"
                  class="inline-flex items-center justify-center rounded-xl bg-gradient-to-r from-indigo-500 via-blue-600 to-emerald-500 px-5 py-2.5 text-sm font-extrabold text-white shadow-xl shadow-indigo-500/25 transition-all hover:from-indigo-400 hover:to-emerald-400 focus-ring"
                >
                  {{ studentHero.cta }}
                </button>
                <button
                  v-if="activeSection !== 'profile'"
                  @click="activeSection = 'profile'"
                  class="inline-flex items-center justify-center rounded-xl border border-indigo-500/30 bg-indigo-500/10 px-5 py-2.5 text-sm font-bold text-indigo-200 transition-all hover:bg-indigo-500/20 hover:text-white focus-ring"
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
                class="rounded-2xl border border-indigo-500/20 bg-slate-900/80 p-4 shadow-xl shadow-indigo-950/30 backdrop-blur-md transition-all hover:border-indigo-500/40"
              >
                <div class="flex items-start justify-between gap-2">
                  <div class="min-w-0">
                    <p class="truncate text-[11px] font-extrabold uppercase tracking-wider text-indigo-300/80">{{ stat.label }}</p>
                    <p class="mt-1.5 text-2xl font-black text-white tracking-tight">{{ stat.value }}</p>
                    <p class="mt-1 truncate text-xs font-semibold text-slate-400">{{ stat.caption }}</p>
                  </div>
                  <span :class="[stat.iconClass, 'inline-flex h-8 w-8 flex-shrink-0 items-center justify-center rounded-xl bg-indigo-500/10 text-indigo-300 ring-1 ring-indigo-500/20']">
                    <svg v-if="stat.icon === 'jobs'" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                    </svg>
                    <svg v-else-if="stat.icon === 'apps'" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                    </svg>
                    <svg v-else-if="stat.icon === 'accepted'" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    <svg v-else class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
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
        <StudentProfileSection :state="studentDashboardState" />
        <StudentApplicationsSection :state="studentDashboardState" />
        <StudentWalletSection :state="studentDashboardState" />
      </main>
    </div>
  </div>

  <StudentDashboardModals :state="studentDashboardState" />
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, onUnmounted, watch } from 'vue'
import StudentJobsSection from '~/components/student/StudentJobsSection.vue'
import StudentProfileSection from '~/components/student/StudentProfileSection.vue'
import StudentApplicationsSection from '~/components/student/StudentApplicationsSection.vue'
import StudentDashboardModals from '~/components/student/StudentDashboardModals.vue'
import StudentWalletSection from '~/components/student/StudentWalletSection.vue'

const { logout, user } = useAuth()
const handleLogout = () => {
  logout()
}

const activeSection = useState<string>('studentDashboardActiveSection', () => 'jobs')

const navItems = [
  { id: 'jobs', name: 'Trang chủ tìm việc' },
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
    iconClass: 'bg-brand-50 text-brand-700'
  },
  {
    label: 'Đơn ứng tuyển',
    value: applications.value.length,
    caption: 'Tổng số đã nộp',
    icon: 'apps',
    iconClass: 'bg-sky-50 text-sky-700'
  },
  {
    label: 'Đã nhận việc',
    value: acceptedApplicationsCount.value,
    caption: `${pendingApplicationsCount.value} đơn chờ duyệt`,
    icon: 'accepted',
    iconClass: 'bg-emerald-50 text-emerald-700'
  },
  {
    label: 'Hồ sơ cá nhân',
    value: `${profileReadiness.value}%`,
    caption: 'Mức độ hoàn thiện',
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
const fetchWallet = async () => {
  isLoadingWallet.value = true

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


onUnmounted(() => {
  if (timerInterval) clearInterval(timerInterval)
})
onMounted(async () => {
  await fetchJobs()
  await fetchApplications()
  await fetchProfile()
  await fetchWallet()

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
