<template>
  <div class="min-h-screen bg-slate-50 flex flex-col md:flex-row">
    <StudentSidebar :state="studentDashboardState" />

    <main class="flex-grow p-6 sm:p-8 bg-slate-50 overflow-y-auto">
      <StudentJobsSection :state="studentDashboardState" />
      <StudentProfileSection :state="studentDashboardState" />
      <StudentApplicationsSection :state="studentDashboardState" />
    </main>
  </div>

  <StudentDashboardModals :state="studentDashboardState" />
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, onUnmounted } from 'vue'
import StudentSidebar from '~/components/student/StudentSidebar.vue'
import StudentJobsSection from '~/components/student/StudentJobsSection.vue'
import StudentProfileSection from '~/components/student/StudentProfileSection.vue'
import StudentApplicationsSection from '~/components/student/StudentApplicationsSection.vue'
import StudentDashboardModals from '~/components/student/StudentDashboardModals.vue'

definePageMeta({
  middleware: 'auth'
})

const activeSection = ref('jobs')

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
  { value: 'all', label: '📍 All Locations' },
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