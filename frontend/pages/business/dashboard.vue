<template>
  <div class="dashboard-body min-h-screen bg-slate-950 text-slate-100 lg:flex font-sans">
    <!-- Sidebar navigation (desktop) -->
    <aside class="hidden lg:flex lg:w-64 lg:flex-shrink-0 lg:flex-col lg:border-r lg:border-indigo-500/15 lg:bg-slate-950/90 lg:px-4 lg:py-6 backdrop-blur-xl sticky top-0 h-screen overflow-y-auto z-30">
      <!-- Enterprise Platform Logo Badge -->
      <div class="px-3 pb-6 border-b border-indigo-500/10 mb-4">
        <button @click="activeSection = 'dashboard'" class="flex items-center gap-2.5 group text-left w-full">
          <div class="h-9 w-9 rounded-xl bg-gradient-to-tr from-blue-600 via-indigo-500 to-emerald-400 flex items-center justify-center text-white font-black text-lg shadow-lg shadow-indigo-500/30 group-hover:scale-105 transition-transform">
            QW
          </div>
          <div>
            <span class="text-[10px] font-black uppercase tracking-wider text-indigo-400">Business SaaS</span>
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
          <svg v-if="item.id === 'dashboard'" class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 3.055A9.001 9.001 0 1020.945 13H11V3.055z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.488 9H15V3.512A9.025 9.025 0 0120.488 9z" />
          </svg>
          <svg v-else-if="item.id === 'profile'" class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
          </svg>
          <svg v-else-if="item.id === 'jobs'" class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
          </svg>
          <svg v-else class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
          </svg>
          <span>{{ item.name }}</span>
        </button>
      </nav>

      <!-- Bottom Profile & Logout Box -->
      <div class="mt-auto space-y-3 pt-4 border-t border-indigo-500/10">
        <div class="rounded-2xl border border-indigo-500/20 bg-gradient-to-b from-indigo-950/40 to-slate-900/60 p-3.5 ring-1 ring-indigo-500/10">
          <div class="flex items-center justify-between mb-1">
            <p class="truncate text-xs font-bold text-white max-w-[130px]">{{ profileForm.company_name || 'Doanh Nghiệp' }}</p>
            <span class="rounded-full bg-emerald-400/10 px-2 py-0.5 text-[10px] font-extrabold text-emerald-300 ring-1 ring-emerald-400/30">
              Đã Duyệt
            </span>
          </div>
          <p class="text-[11px] font-medium text-slate-400 truncate">MST: {{ profileForm.tax_code || 'Đã xác minh' }}</p>
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

    <!-- Main Content Area -->
    <div class="flex-1 min-w-0">
      <section class="border-b border-indigo-500/15 bg-gradient-to-r from-slate-950 via-indigo-950/40 to-slate-950">
        <div class="w-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
          <div class="grid gap-5 lg:grid-cols-[0.95fr_1.05fr] lg:items-center">
            <div>
              <div class="flex flex-wrap items-center gap-2">
                <span class="inline-flex items-center rounded-full bg-indigo-500/10 px-3 py-1 text-xs font-black uppercase tracking-wide text-indigo-300 ring-1 ring-indigo-500/30">
                  Không gian Doanh nghiệp
                </span>
                <span class="text-xs font-semibold text-slate-400">Trung tâm quản lý tuyển dụng & Smart Escrow</span>
              </div>
              <h1 class="mt-2 text-2xl sm:text-3xl font-extrabold tracking-tight text-white">
                {{ businessHero.title }}
              </h1>
              <p class="mt-2 max-w-2xl text-xs sm:text-sm font-semibold leading-relaxed text-slate-300">
                {{ businessHero.description }}
              </p>
              <div class="mt-4 flex flex-wrap gap-3">
                <button
                  @click="handleBusinessHeroAction"
                  class="inline-flex items-center justify-center rounded-xl bg-gradient-to-r from-indigo-500 via-blue-600 to-emerald-500 px-5 py-2.5 text-xs font-black text-white shadow-lg shadow-indigo-500/25 transition-all hover:from-indigo-400 hover:to-emerald-400 focus-ring"
                >
                  {{ businessHero.cta }}
                </button>
                <button
                  v-if="activeSection !== 'applicants'"
                  @click="activeSection = 'applicants'"
                  class="inline-flex items-center justify-center rounded-xl border border-indigo-500/30 bg-indigo-500/10 px-5 py-2.5 text-xs font-black text-indigo-200 hover:bg-indigo-500/20 hover:text-white transition-all focus-ring"
                >
                  Xem danh sách ứng viên
                </button>
              </div>
            </div>

            <div class="grid grid-cols-2 gap-3">
              <div
                v-for="stat in businessHeroStats"
                :key="stat.label"
                class="rounded-2xl border border-indigo-500/20 bg-slate-900/80 px-4 py-3.5 shadow-xl shadow-indigo-950/30 backdrop-blur"
              >
                <div class="flex items-start justify-between gap-3">
                  <div>
                    <p class="text-[10px] font-black uppercase tracking-wider text-indigo-300">{{ stat.label }}</p>
                    <p class="mt-1 text-2xl font-black text-white">{{ stat.value }}</p>
                    <p class="mt-1 text-[11px] font-semibold text-slate-400">{{ stat.caption }}</p>
                  </div>
                  <span :class="[stat.iconClass, 'inline-flex h-9 w-9 flex-shrink-0 items-center justify-center rounded-xl bg-indigo-500/10 text-indigo-300 ring-1 ring-indigo-500/20']">
                    <svg v-if="stat.icon === 'jobs'" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                    </svg>
                    <svg v-else-if="stat.icon === 'apps'" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
                    </svg>
                    <svg v-else-if="stat.icon === 'pending'" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    <svg v-else class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 3.055A9.001 9.001 0 1020.945 13H11V3.055z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.488 9H15V3.512A9.025 9.025 0 0120.488 9z" />
                    </svg>
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <main class="w-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <BusinessOverviewSection :state="businessDashboardState" />
        <BusinessProfileSection :state="businessDashboardState" />
        <BusinessJobsSection :state="businessDashboardState" />
        <BusinessApplicantsSection :state="businessDashboardState" />
      </main>
    </div>
  </div>

  <BusinessDashboardModals :state="businessDashboardState" />
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, watch } from 'vue'
import BusinessOverviewSection from '~/components/business/BusinessOverviewSection.vue'
import BusinessProfileSection from '~/components/business/BusinessProfileSection.vue'
import BusinessJobsSection from '~/components/business/BusinessJobsSection.vue'
import BusinessApplicantsSection from '~/components/business/BusinessApplicantsSection.vue'
import BusinessDashboardModals from '~/components/business/BusinessDashboardModals.vue'
import { useApi } from '~/composables/useApi'
import { useAuth } from '~/composables/useAuth'

definePageMeta({
  middleware: 'auth'
})

const { logout } = useAuth()
const handleLogout = () => {
  logout()
}

// Navigation setups
const activeSection = useState<string>('businessDashboardActiveSection', () => 'dashboard')

const navItems = [
  { id: 'dashboard', name: 'Dashboard' },
  { id: 'profile', name: 'Company Profile' },
  { id: 'jobs', name: 'Jobs' },
  { id: 'applicants', name: 'Applicants' }
]

// API Client references
const api = useApi()
const jobs = ref<any[]>([])
const applications = ref<any[]>([])
const isLoadingJobs = ref(false)
const isLoadingApps = ref(false)
const isSavingProfile = ref(false)
const isCreatingJob = ref(false)
const showCreateForm = ref(false)
const feedback = ref<{ type: 'success' | 'error'; message: string } | null>(null)

// Confirmation Modal States
const showConfirmModal = ref(false)
const confirmTarget = ref<any | null>(null)
const confirmAction = ref<'approved' | 'rejected'>('approved')
const isReviewing = ref(false)

// Query filters
const applicantSearchQuery = ref('')
const applicantStatusFilter = ref('all')
const businessJobsPage = ref(1)
const businessJobsPageSize = 10
const applicantsPage = ref(1)
const applicantsPageSize = 8


// Trạng thái quản lý đóng mở form và file logo của Business
const isEditing = ref(false)
const logoFileSelected = ref<File | null>(null)
const logoPreview = ref<string | null>(null)


const selectedApp = ref<any>(null)
const reviewStatus = ref<'approved' | 'rejected' | null>(null)
const isSubmitting = ref(false)
const selectedCompletionApp = ref<any | null>(null)
const isCompletingJob = ref(false)
const selectedBusinessReviewApp = ref<any | null>(null)
const businessReviewRating = ref(5)
const businessReviewComment = ref('')
const isSubmittingBusinessReview = ref(false)
const selectedManagedApplicant = ref<any | null>(null)
const selectedReviewsApp = ref<any | null>(null)
const reviewsList = ref<any[]>([])
const reviewSummary = ref<{ average_rating: number; total_reviews: number }>({
  average_rating: 0,
  total_reviews: 0
})
const isLoadingReviews = ref(false)
const isReviewsModalOpen = ref(false)

// Form thông tin Offer gửi đính kèm
const offerForm = ref({
  salary: '',
  startDate: '',
  message: ''
})
const openReviewModal = (app: any) => {
  selectedApp.value = app
  reviewStatus.value = app.status === 'pending' ? 'approved' : app.status // Mặc định mở ra chọn luôn approved cho tiện
  offerForm.value = {
    salary: app.offer_salary || '',
    startDate: app.offer_start_date || '',
    message: app.offer_message || ''
  }
}

const closeModal = () => {
  selectedApp.value = null
  reviewStatus.value = null
}

const openCompletionModal = (app: any) => {
  selectedCompletionApp.value = app
  feedback.value = null
}

const closeCompletionModal = () => {
  selectedCompletionApp.value = null
}

const openBusinessReviewModal = (app: any) => {
  selectedBusinessReviewApp.value = app
  businessReviewRating.value = 5
  businessReviewComment.value = ''
  feedback.value = null
}

const closeBusinessReviewModal = () => {
  selectedBusinessReviewApp.value = null
}

const openManagedApplicantModal = (app: any) => {
  selectedManagedApplicant.value = app
  feedback.value = null
  fetchReviewsByApplication(app.id)
}

const closeManagedApplicantModal = () => {
  selectedManagedApplicant.value = null
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

const submitBusinessReview = async () => {
  if (!selectedBusinessReviewApp.value) return

  isSubmittingBusinessReview.value = true
  feedback.value = null

  try {
    const res = await api.post('/api/reviews', {
      application_id: selectedBusinessReviewApp.value.id,
      rating: businessReviewRating.value,
      comment: businessReviewComment.value
    })

    feedback.value = {
      type: 'success',
      message: res.message || 'Đánh giá ứng viên thành công.'
    }

    closeBusinessReviewModal()
    await fetchApplications()
  } catch (err: any) {
    feedback.value = {
      type: 'error',
      message: err.response?._data?.error || err.response?._data?.message || 'Không thể gửi đánh giá.'
    }
  } finally {
    isSubmittingBusinessReview.value = false
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const submitBusinessCompletion = async () => {
  if (!selectedCompletionApp.value) return

  isCompletingJob.value = true
  feedback.value = null

  try {
    const res = await api.post('/api/applications/business-complete', {
      application_id: selectedCompletionApp.value.id
    })

    feedback.value = {
      type: 'success',
      message: res.message || 'Đã xác nhận hoàn thành và giải ngân lương.'
    }

    closeCompletionModal()
    await fetchApplications()
  } catch (err: any) {
    feedback.value = {
      type: 'error',
      message: err.response?._data?.error || err.response?._data?.message || 'Không thể xác nhận hoàn thành công việc.'
    }
  } finally {
    isCompletingJob.value = false
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const submitReview = async () => {
  if (!selectedApp.value || !reviewStatus.value) return

  isSubmitting.value = true
  try {
    const payload = {
      application_id: selectedApp.value.id,
      status: reviewStatus.value,
      offer_salary: reviewStatus.value === 'approved' ? offerForm.value.salary : '',
      offer_start_date: reviewStatus.value === 'approved' ? offerForm.value.startDate : '',
      offer_message: reviewStatus.value === 'approved' ? offerForm.value.message : ''
    }

    await api.put('/api/jobs/review-application', payload)
    
    // 🌟 TỰ ĐỘNG ĐỒNG BỘ: Quét và gọi đúng tên hàm fetch hiện có trong file của Chanh
    const anyWindow = window as any
    if (typeof anyWindow.fetchApplications === 'function') {
      await anyWindow.fetchApplications()
    } else if (typeof anyWindow.fetchApplicants === 'function') {
      await anyWindow.fetchApplicants()
    } else if (typeof anyWindow.fetchEmployerApplications === 'function') {
      await anyWindow.fetchEmployerApplications()
    } else {
      // Giải pháp an toàn cuối cùng nếu không tìm thấy hàm nào: Reload lại trang
      window.location.reload()
    }
    
    closeModal()
  } catch (err) {
    console.error('Lỗi khi cập nhật trạng thái hoặc gửi offer:', err)
    alert('Có lỗi xảy ra khi phê duyệt.')
  } finally {
    isSubmitting.value = false
  }
}

  const onLogoFileChange = (e: Event) => {
  const target = e.target as HTMLInputElement
  const file = target.files?.[0]
  if (file) {
    logoFileSelected.value = file
    logoPreview.value = URL.createObjectURL(file) // Tạo preview cục bộ
  }
}
  const fetchProfile = async () => {
  try {
    const res = await api.get('/api/profile/business')
    const data = res && res.data ? res.data : res
    if (data) {
      profileForm.company_name = data.company_name || ''
      profileForm.phone = data.phone || ''
      profileForm.address = data.address || ''
      profileForm.logo_url = data.logo_url || ''
      currentBusinessUserId.value = data.user_id || data.id
    }
  } catch (err) {
    console.error('Error fetching company profile:', err)
  }
}

// 🌟 CẤU HÌNH ĐIỀU KHIỂN MODAL CHAT (THÊM VÀO SCRIPT PHÍA DOANH NGHIỆP)
const isChatModalOpen = ref(false)
const selectedChatApp = ref<any>(null)

// Biến hứng ID Doanh nghiệp (Chanh nhớ gán data.user_id vào biến này trong hàm fetchProfile/fetchCompany của bạn nhé)
const currentBusinessUserId = ref<number | null>(2) // Tạm thời để 1 để test, hoặc ref(null) nếu gán động

const openChatModal = (app: any) => {
  selectedChatApp.value = app
  isChatModalOpen.value = true
}
// Forms Setup
const profileForm = reactive({
  company_name: '',
  phone: '',
  address: '',
  logo_url: ''
})

const jobForm = reactive({
  title: '',
  description: '',
  location: '',
  salary: 0,
  slots: 1,
  working_date: ''
})

// Metrics and compute operations
const jobTitleLookup = (jobId: number): string => {
  const job = jobs.value.find(j => j.id === jobId)
  return job ? job.title : `Position #${jobId || ''}`
}

const fillRatio = computed(() => {
  if (jobs.value.length === 0) return 0
  const approvedTotal = applications.value.filter(app => app.status?.toLowerCase() === 'approved').length
  const slotsTotal = jobs.value.reduce((acc, j) => acc + (j.slots || 1), 0)
  return Math.min(100, Math.round((approvedTotal / slotsTotal) * 100))
})

const metricsCards = computed(() => {
  const total = applications.value.length
  const pending = applications.value.filter(app => {
    const status = app.status?.toLowerCase()
    return status === 'applied' || status === 'pending'
  }).length
  const approved = applications.value.filter(app => {
    const status = app.status?.toLowerCase()
    return status === 'approved' || status === 'offer_accepted' || status === 'student_completed' || status === 'paid'
  }).length
  const rejected = applications.value.filter(app => app.status?.toLowerCase() === 'rejected').length

  return [
    { title: 'Total Applications', value: total, label: 'Applications', color: 'bg-slate-50 border-slate-200 text-slate-700' },
    { title: 'Pending Applications', value: pending, label: 'Awaiting Review', color: 'bg-amber-50 border-amber-200 text-amber-700' },
    { title: 'Approved Applications', value: approved, label: 'Accepted', color: 'bg-emerald-50 border-emerald-200 text-emerald-700' },
    { title: 'Rejected Applications', value: rejected, label: 'Declined', color: 'bg-rose-50 border-rose-200 text-rose-700' }
  ]
})

const pendingApplicationsCount = computed(() => {
  return applications.value.filter(app => {
    const status = app.status?.toLowerCase()
    return status === 'applied' || status === 'pending'
  }).length
})

const acceptedApplicationsCount = computed(() => {
  return applications.value.filter(app => {
    const status = app.status?.toLowerCase()
    return status === 'approved' || status === 'offer_accepted' || status === 'student_completed' || status === 'paid'
  }).length
})

const businessHero = computed(() => {
  if (activeSection.value === 'profile') {
    return {
      title: profileForm.company_name ? `Hồ sơ công ty: ${profileForm.company_name}` : 'Hoàn thiện hồ sơ doanh nghiệp',
      description: 'Cập nhật thông tin công ty rõ ràng giúp sinh viên hiểu rõ thương hiệu trước khi ứng tuyển.',
      cta: 'Cập nhật hồ sơ'
    }
  }

  if (activeSection.value === 'jobs') {
    return {
      title: 'Quản lý bài đăng tuyển dụng',
      description: 'Đăng tin tuyển dụng rõ ràng, theo dõi nhu cầu nhân sự và duy trì các cơ hội phù hợp cho sinh viên.',
      cta: 'Đăng tin mới'
    }
  }

  if (activeSection.value === 'applicants') {
    return {
      title: 'Quản lý ứng viên nộp đơn',
      description: 'Xem xét hồ sơ ứng viên, chuyển đổi trạng thái nộp đơn và trao đổi trực tiếp từ một giao diện tập trung.',
      cta: 'Xem đơn chờ duyệt'
    }
  }

  return {
    title: 'Tổng quan hiệu quả tuyển dụng',
    description: 'Theo dõi bài đăng tuyển dụng, ứng viên nộp đơn và tiến độ nhận việc thực tế.',
    cta: 'Đăng tin tuyển dụng'
  }
})

const businessHeroStats = computed(() => [
  {
    label: 'Bài đăng hoạt động',
    value: jobs.value.length,
    caption: 'Đang mở tuyển',
    icon: 'jobs',
    iconClass: 'bg-brand-50 text-brand-700'
  },
  {
    label: 'Ứng viên nộp đơn',
    value: applications.value.length,
    caption: 'Tổng đơn nhận được',
    icon: 'apps',
    iconClass: 'bg-sky-50 text-sky-700'
  },
  {
    label: 'Đơn chờ duyệt',
    value: pendingApplicationsCount.value,
    caption: 'Cần xem xét',
    icon: 'pending',
    iconClass: 'bg-amber-50 text-amber-700'
  },
  {
    label: 'Tỷ lệ nhận việc',
    value: `${fillRatio.value}%`,
    caption: `${acceptedApplicationsCount.value} ứng viên nhận việc`,
    icon: 'fill',
    iconClass: 'bg-emerald-50 text-emerald-700'
  }
])

const handleBusinessHeroAction = () => {
  if (activeSection.value === 'profile') {
    isEditing.value = true
    return
  }

  if (activeSection.value === 'applicants') {
    applicantStatusFilter.value = 'pending'
    return
  }

  activeSection.value = 'jobs'
  showCreateForm.value = true
}

const filteredApps = computed(() => {
  return applications.value.filter(app => {
    const studentName = app.student?.full_name?.toLowerCase() || ''
    const jobTitle = jobTitleLookup(app.job_id).toLowerCase()

    const matchesSearch = !applicantSearchQuery.value ||
      studentName.includes(applicantSearchQuery.value.toLowerCase()) ||
      jobTitle.includes(applicantSearchQuery.value.toLowerCase())

    const matchesStatus = applicantStatusFilter.value === 'all' ||
      app.status?.toLowerCase() === applicantStatusFilter.value.toLowerCase()

    return matchesSearch && matchesStatus
  })
})

const paginatedBusinessJobs = computed(() => {
  const start = (businessJobsPage.value - 1) * businessJobsPageSize
  return jobs.value.slice(start, start + businessJobsPageSize)
})

const paginatedApplicants = computed(() => {
  const start = (applicantsPage.value - 1) * applicantsPageSize
  return filteredApps.value.slice(start, start + applicantsPageSize)
})

watch(jobs, () => {
  businessJobsPage.value = 1
})

watch(filteredApps, () => {
  applicantsPage.value = 1
})

const statusBadgeClass = (status: string): string => {
  const norm = status?.toLowerCase() || ''
  if (norm === 'approved' || norm === 'offer_accepted') return 'bg-emerald-50 border-emerald-200 text-emerald-700'
  if (norm === 'student_completed') return 'bg-amber-50 border-amber-200 text-amber-700'
  if (norm === 'paid') return 'bg-cyan-50 border-cyan-200 text-cyan-700'
  if (norm === 'rejected') return 'bg-rose-50 border-rose-200 text-rose-700'
  return 'bg-slate-50 border-slate-200 text-slate-700'
}

const formatDate = (dateVal: any): string => {
  return 'Jun 24, 2026'
}

// API Methods
const fetchJobs = async () => {
  isLoadingJobs.value = true
  try {
    const res = await api.get('/api/jobs')
    jobs.value = res.data || []
  } catch (err) {
    console.error('Error fetching jobs:', err)
  } finally {
    isLoadingJobs.value = false
  }
}

const fetchApplications = async () => {
  isLoadingApps.value = true
  try {
    const res = await api.get('/api/applications/employer')
    applications.value = res.data || []
  } catch (err) {
    console.error('Error fetching employer applications:', err)
  } finally {
    isLoadingApps.value = false
  }
}

const handleUpdateProfile = async () => {
  isSavingProfile.value = true
  feedback.value = null

  try {
    const formData = new FormData()
    formData.append('company_name', profileForm.company_name)
    formData.append('phone', profileForm.phone)
    formData.append('address', profileForm.address)

    // Nếu doanh nghiệp có chọn file logo mới thì đính kèm vào payload
    if (logoFileSelected.value) {
      formData.append('logo', logoFileSelected.value)
    }

    const res = await api.put('/api/profile/business', formData)
    
    feedback.value = {
      type: 'success',
      message: res.message || '🎉 Company profile updated successfully!'
    }

    // Đồng bộ lại đường dẫn logo mới nhất trả về từ server để hiển thị ngay
    const updatedData = res.data ? res.data : res
    if (updatedData) {
      profileForm.logo_url = updatedData.LogoUrl || updatedData.logo_url || profileForm.logo_url
    }

    // Tắt chế độ sửa, quay về chế độ xem và dọn dẹp file tạm
    isEditing.value = false
    logoFileSelected.value = null
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
      message: res.message || '🎉 Job opening created successfully!'
    }

    // Reset Form fields
    jobForm.title = ''
    jobForm.description = ''
    jobForm.location = ''
    jobForm.salary = 0
    jobForm.slots = 1
    jobForm.working_date = ''

    // Return to main jobs overview
    showCreateForm.value = false
    await fetchJobs()
  } catch (err: any) {
    feedback.value = {
      type: 'error',
      message: err.response?._data?.message || 'Failed to publish job opening.'
    }
  } finally {
    isCreatingJob.value = false
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const triggerConfirmModal = (app: any, status: 'approved' | 'rejected') => {
  confirmTarget.value = app
  confirmAction.value = status
  showConfirmModal.value = true
}

const handleReviewApplication = async () => {
  if (!confirmTarget.value) return

  isReviewing.value = true
  feedback.value = null

  const appId = confirmTarget.value.id
  const targetStatus = confirmAction.value

  // Optimistic UI updates
  const originalAppsState = [...applications.value]
  const targetIndex = applications.value.findIndex(app => app.id === appId)
  if (targetIndex !== -1) {
    applications.value[targetIndex].status = targetStatus
  }

  try {
    const res = await api.put('/api/jobs/review-application', {
      application_id: Number(appId),
      status: targetStatus
    })

    feedback.value = {
      type: 'success',
      message: res.message || '🎉 Candidate status evaluated successfully!'
    }
  } catch (err: any) {
    // Revert optimistic update
    applications.value = originalAppsState
    feedback.value = {
      type: 'error',
      message: err.response?._data?.message || 'Failed to update candidate status.'
    }
  } finally {
    isReviewing.value = false
    showConfirmModal.value = false
    confirmTarget.value = null
    // Sync-refresh list data
    await fetchApplications()
  }
}
const parseSkills = (skillsRaw: string): string[] => {
  if (!skillsRaw) return []
  try {
    const parsed = JSON.parse(skillsRaw)
    if (Array.isArray(parsed)) return parsed
    return skillsRaw.split(',').map(s => s.trim())
  } catch (e) {
    return skillsRaw.split(',').map(s => s.trim()).filter(s => s !== '')
  }
}

const businessDashboardState = reactive({
  activeSection,
  navItems,
  jobs,
  applications,
  metricsCards,
  fillRatio,
  isEditing,
  profileForm,
  logoPreview,
  onLogoFileChange,
  isSavingProfile,
  handleUpdateProfile,
  showCreateForm,
  jobForm,
  handleCreateJob,
  isCreatingJob,
  isLoadingJobs,
  isLoadingApps,
  paginatedBusinessJobs,
  businessJobsPage,
  businessJobsPageSize,
  applicantSearchQuery,
  applicantStatusFilter,
  filteredApps,
  paginatedApplicants,
  applicantsPage,
  applicantsPageSize,
  jobTitleLookup,
  formatDate,
  statusBadgeClass,
  parseSkills,
  openChatModal,
  openReviewModal,
  triggerConfirmModal,
  showConfirmModal,
  confirmTarget,
  confirmAction,
  isReviewing,
  handleReviewApplication,
  selectedApp,
  reviewStatus,
  offerForm,
  isSubmitting,
  selectedCompletionApp,
  isCompletingJob,
  openCompletionModal,
  closeCompletionModal,
  submitBusinessCompletion,
  selectedBusinessReviewApp,
  businessReviewRating,
  businessReviewComment,
  isSubmittingBusinessReview,
  openBusinessReviewModal,
  closeBusinessReviewModal,
  submitBusinessReview,
  selectedManagedApplicant,
  openManagedApplicantModal,
  closeManagedApplicantModal,
  selectedReviewsApp,
  reviewsList,
  reviewSummary,
  isLoadingReviews,
  isReviewsModalOpen,
  openReviewsModal,
  closeReviewsModal,
  fetchReviewsByApplication,
  fetchReviewsByUser,
  closeModal,
  submitReview,
  isChatModalOpen,
  selectedChatApp,
  currentBusinessUserId
})

onMounted(() => {
  fetchJobs()
  fetchApplications()
  fetchProfile() // <-- Gọi nạp dữ liệu cũ của Công ty ở đây nha Chanh
})
</script>
