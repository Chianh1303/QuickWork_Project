import { useApi } from '~/composables/useApi'
import type {
  AdminDashboardStats,
  BusinessKYBDetail,
  PendingBusinessListResponse,
  ReviewBusinessPayload
} from '~/types/admin'

export interface PendingBusinessQuery {
  page?: number
  limit?: number
  search?: string
}

export const useAdminApi = () => {
  const api = useApi()

  const getDashboardStats = () => {
    return api.get<AdminDashboardStats>('/api/admin/dashboard/stats')
  }

  const getPendingBusinesses = (query: PendingBusinessQuery = {}) => {
    return api.get<PendingBusinessListResponse>('/api/admin/businesses/pending', {
      query: {
        page: query.page ?? 1,
        limit: query.limit ?? 10,
        search: query.search ?? ''
      }
    })
  }

  const getBusinessDetail = (businessId: number) => {
    return api.get<BusinessKYBDetail>(`/api/admin/businesses/${businessId}`)
  }

  const reviewBusiness = (businessId: number, payload: ReviewBusinessPayload) => {
    return api.put<{ message: string; decision: ReviewBusinessPayload['decision'] }>(
      `/api/admin/businesses/${businessId}/review`,
      payload
    )
  }

  return {
    getDashboardStats,
    getPendingBusinesses,
    getBusinessDetail,
    reviewBusiness
  }
}
