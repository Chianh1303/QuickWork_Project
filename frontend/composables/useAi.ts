import { ref } from 'vue'
import { useApi } from '~/composables/useApi'

export interface Recommendation {
  decision: string
  confidence: number
  reason: string
}

export interface ATSReview {
  score: number
  strengths: string[]
  issues: string[]
}

export interface SkillsReview {
  score: number
  technical: string[]
  soft: string[]
}

export interface ProjectReview {
  name: string
  technologies: string[]
  description: string
  impact: string
}

export interface ExperienceReview {
  score: number
  projects: ProjectReview[]
}

export interface EducationItem {
  institution: string
  degree: string
  graduation_year: string
}

export interface EducationReview {
  score: number
  items: EducationItem[]
}

export interface STARReview {
  situation: number
  task: number
  action: number
  result: number
}

export interface EvaluateCvResult {
  score: number
  recommendation?: Recommendation
  ats?: ATSReview
  skills?: SkillsReview
  experience?: ExperienceReview
  education?: EducationReview
  strengths?: string[]
  weaknesses?: string[]
  improvements?: string[]
  missing_information?: string[]
  star_analysis?: STARReview
  suggested_summary: string
  actionable_tips: string[]
  evaluation_source?: string
}

export interface MatchJobResult {
  match_score: number
  matching_reasons: string[]
  missing_skills: string[]
}

export interface GenerateJobResult {
  description: string
  requirements: string
  benefits: string
  suggested_salary: string
}

export interface RecommendedJobItem {
  job_id: number
  job_title: string
  company: string
  logo_url?: string
  description?: string
  salary?: number
  location?: string
  match_score: number
  matching_skills: string[]
  missing_skills: string[]
}

export interface RecommendedJobsResponse {
  jobs: RecommendedJobItem[]
}

export const useAi = () => {
  const api = useApi()
  const isLoading = ref(false)
  const errorMessage = ref('')

  const evaluateCv = async (payload?: Record<string, any>): Promise<EvaluateCvResult> => {
    isLoading.value = true
    errorMessage.value = ''
    try {
      const data = await api.post<EvaluateCvResult>('/api/ai/evaluate-cv', payload || {})
      return data
    } catch (err: any) {
      errorMessage.value = err?.data?.message || 'Không thể phân tích CV lúc này.'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const matchJob = async (jobId: number): Promise<MatchJobResult> => {
    isLoading.value = true
    errorMessage.value = ''
    try {
      const data = await api.post<MatchJobResult>('/api/ai/match-job', { job_id: jobId })
      return data
    } catch (err: any) {
      errorMessage.value = err?.data?.message || 'Không thể tính toán độ phù hợp công việc.'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const generateJobDescription = async (title: string): Promise<GenerateJobResult> => {
    isLoading.value = true
    errorMessage.value = ''
    try {
      const data = await api.post<GenerateJobResult>('/api/ai/generate-job', { title })
      return data
    } catch (err: any) {
      errorMessage.value = err?.data?.message || 'Không thể tạo mô tả tin tuyển dụng.'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const getRecommendedJobs = async (): Promise<RecommendedJobsResponse> => {
    isLoading.value = true
    errorMessage.value = ''
    try {
      const data = await api.get<RecommendedJobsResponse>('/api/ai/recommended-jobs')
      return data
    } catch (err: any) {
      errorMessage.value = err?.data?.message || 'Không thể lấy danh sách công việc gợi ý từ AI.'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  return {
    isLoading,
    errorMessage,
    evaluateCv,
    matchJob,
    generateJobDescription,
    getRecommendedJobs
  }
}
