import { ref } from 'vue'
import { useToast } from '~/composables/useToast'

export const useSavedJobs = () => {
  const { success, info } = useToast()
  const savedJobIds = ref<number[]>([])

  // Initialize from localStorage
  const initSavedJobs = () => {
    if (import.meta.client) {
      try {
        const stored = localStorage.getItem('qw_saved_jobs')
        if (stored) {
          savedJobIds.value = JSON.parse(stored)
        }
      } catch (e) {
        console.error('Error reading saved jobs from localStorage', e)
      }
    }
  }

  const saveToStorage = () => {
    if (import.meta.client) {
      try {
        localStorage.setItem('qw_saved_jobs', JSON.stringify(savedJobIds.value))
      } catch (e) {
        console.error('Error saving to localStorage', e)
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

  const toggleSaveJob = (job: any) => {
    const id = getJobId(job)
    if (!id) return

    const index = savedJobIds.value.indexOf(id)
    if (index > -1) {
      savedJobIds.value.splice(index, 1)
      saveToStorage()
      info(`Đã bỏ lưu công việc "${job?.title || job?.job_title || '#' + id}" khỏi danh sách yêu thích.`)
    } else {
      savedJobIds.value.push(id)
      saveToStorage()
      success(`❤️ Đã lưu công việc "${job?.title || job?.job_title || '#' + id}" vào mục Yêu thích!`)
    }
  }

  const shareJob = async (job: any) => {
    if (import.meta.client) {
      const url = `${window.location.origin}/jobs?jobId=${job?.id || ''}`
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

  initSavedJobs()

  return {
    savedJobIds,
    isJobSaved,
    toggleSaveJob,
    shareJob
  }
}
