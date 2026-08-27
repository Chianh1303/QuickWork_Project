import { ref } from 'vue'
import { useToast } from '~/composables/useToast'

export const useSavedJobs = () => {
  const { success, info } = useToast()
  const savedJobIds = ref<number[]>([])
  const savedJobsList = ref<any[]>([])

  const saveToStorage = () => {
    if (import.meta.client) {
      try {
        localStorage.setItem('qw_saved_jobs', JSON.stringify(savedJobIds.value))
        localStorage.setItem('qw_saved_jobs_list', JSON.stringify(savedJobsList.value))
      } catch (e) {
        console.error('Error writing localStorage', e)
      }
    }
  }

  const getJobId = (job: any): number => {
    if (typeof job === 'number') return job
    return Number(job?.id || job?.ID || job?.job_id || job?.JobID || 0)
  }

  const isJobSaved = (job: any): boolean => {
    const id = getJobId(job)
    if (!id) return false
    return savedJobIds.value.includes(id)
  }

  // Load from localStorage & sync from API safely on client side
  const initSavedJobs = async () => {
    if (!import.meta.client) return

    // 1. Read from localStorage
    try {
      const storedIds = localStorage.getItem('qw_saved_jobs')
      if (storedIds) {
        const parsed = JSON.parse(storedIds)
        if (Array.isArray(parsed)) {
          savedJobIds.value = parsed.map(n => Number(n)).filter(n => n > 0)
        }
      }

      const storedJobs = localStorage.getItem('qw_saved_jobs_list')
      if (storedJobs) {
        const parsedJobs = JSON.parse(storedJobs)
        if (Array.isArray(parsedJobs)) {
          savedJobsList.value = parsedJobs
        }
      }
    } catch (e) {
      // Ignore
    }

    // 2. Fetch from backend API if token exists
    try {
      const token = useCookie<string | null>('auth_token')
      if (token.value) {
        const api = useApi()
        const res: any = await api.get('/api/saved-jobs', { skipAutoLogout: true })
        if (res) {
          if (Array.isArray(res.saved_ids)) {
            const remoteIds = res.saved_ids.map((id: any) => Number(id)).filter((id: number) => id > 0)
            savedJobIds.value = Array.from(new Set([...savedJobIds.value, ...remoteIds]))
          }
          if (Array.isArray(res.jobs) && res.jobs.length > 0) {
            // Deduplicate backend jobs with local jobs
            const map = new Map<number, any>()
            savedJobsList.value.forEach(j => {
              const id = getJobId(j)
              if (id > 0) map.set(id, j)
            })
            res.jobs.forEach((j: any) => {
              const id = getJobId(j)
              if (id > 0) map.set(id, j)
            })
            savedJobsList.value = Array.from(map.values())
          }
          saveToStorage()
        }
      }
    } catch (e) {
      // Fail silently (guest user or offline mode)
    }
  }

  const toggleSaveJob = async (job: any) => {
    const id = getJobId(job)
    if (!id) return

    const index = savedJobIds.value.indexOf(id)
    const title = job?.title || job?.job_title || '#' + id

    if (index > -1) {
      // Remove
      savedJobIds.value.splice(index, 1)
      savedJobsList.value = savedJobsList.value.filter(j => getJobId(j) !== id)
      saveToStorage()
      info(`Đã bỏ lưu công việc "${title}" khỏi mục yêu thích.`)
    } else {
      // Add
      savedJobIds.value.push(id)
      if (typeof job === 'object' && job !== null) {
        savedJobsList.value.unshift(job)
      }
      saveToStorage()
      success(`❤️ Đã lưu công việc "${title}" vào mục Yêu thích!`)
    }

    // Backend API sync
    if (import.meta.client) {
      try {
        const token = useCookie<string | null>('auth_token')
        if (token.value) {
          const api = useApi()
          await api.post(`/api/saved-jobs/${id}`, null, { skipAutoLogout: true })
        }
      } catch (err) {
        // Fail silently
      }
    }
  }

  const shareJob = async (job: any) => {
    if (import.meta.client) {
      const id = getJobId(job)
      const url = `${window.location.origin}/jobs?jobId=${id || ''}`
      try {
        if (navigator.clipboard) {
          await navigator.clipboard.writeText(url)
          success('📋 Đã sao chép liên kết việc làm vào bộ nhớ tạm!')
        } else {
          success('Liên kết công việc: ' + url)
        }
      } catch (err) {
        success('Liên kết công việc: ' + url)
      }
    }
  }

  if (import.meta.client) {
    initSavedJobs()
  }

  return {
    savedJobIds,
    savedJobsList,
    isJobSaved,
    toggleSaveJob,
    shareJob,
    fetchSavedJobs: initSavedJobs
  }
}
