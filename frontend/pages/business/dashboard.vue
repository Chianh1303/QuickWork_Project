<template>
  <div class="min-h-screen bg-slate-50 flex flex-col md:flex-row">
    <BusinessSidebar :state="businessDashboardState" />

    <main class="flex-grow p-6 sm:p-8 bg-slate-50 overflow-y-auto">
      <BusinessOverviewSection :state="businessDashboardState" />
      <BusinessProfileSection :state="businessDashboardState" />
      <BusinessJobsSection :state="businessDashboardState" />
      <BusinessApplicantsSection :state="businessDashboardState" />
    </main>
  </div>

  <BusinessDashboardModals :state="businessDashboardState" />
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import BusinessSidebar from '~/components/business/BusinessSidebar.vue'
import BusinessOverviewSection from '~/components/business/BusinessOverviewSection.vue'
import BusinessProfileSection from '~/components/business/BusinessProfileSection.vue'
import BusinessJobsSection from '~/components/business/BusinessJobsSection.vue'
import BusinessApplicantsSection from '~/components/business/BusinessApplicantsSection.vue'
import BusinessDashboardModals from '~/components/business/BusinessDashboardModals.vue'
import { useApi } from '~/composables/useApi'

definePageMeta({
  middleware: 'auth'
})

// Navigation setups
const activeSection = ref('dashboard')

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


// Trạng thái quản lý đóng mở form và file logo của Business
const isEditing = ref(false)
const logoFileSelected = ref<File | null>(null)
const logoPreview = ref<string | null>(null)


const selectedApp = ref<any>(null)
const reviewStatus = ref<'approved' | 'rejected' | null>(null)
const isSubmitting = ref(false)

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
  const pending = applications.value.filter(app => app.status?.toLowerCase() === 'applied').length
  const approved = applications.value.filter(app => app.status?.toLowerCase() === 'approved').length
  const rejected = applications.value.filter(app => app.status?.toLowerCase() === 'rejected').length

  return [
    { title: 'Total Applications', value: total, label: 'Applications', color: 'bg-slate-50 border-slate-200 text-slate-700' },
    { title: 'Pending Applications', value: pending, label: 'Awaiting Review', color: 'bg-amber-50 border-amber-200 text-amber-700' },
    { title: 'Approved Applications', value: approved, label: 'Accepted', color: 'bg-emerald-50 border-emerald-200 text-emerald-700' },
    { title: 'Rejected Applications', value: rejected, label: 'Declined', color: 'bg-rose-50 border-rose-200 text-rose-700' }
  ]
})

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

const statusBadgeClass = (status: string): string => {
  const norm = status?.toLowerCase() || ''
  if (norm === 'approved') return 'bg-emerald-50 border-emerald-200 text-emerald-700'
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
  applicantSearchQuery,
  applicantStatusFilter,
  filteredApps,
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
