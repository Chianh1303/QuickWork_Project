import { useApi } from '~/composables/useApi'
import type { PendingBusinessListResponse } from '~/types/admin'

export interface PendingBusinessQuery {
  page?: number
  limit?: number
  search?: string
}

export const useAdminApi = () => {
  const api = useApi()

  const getPendingBusinesses = (query: PendingBusinessQuery = {}) => {
    return api.get<PendingBusinessListResponse>('/api/admin/businesses/pending', {
      query: {
        page: query.page ?? 1,
        limit: query.limit ?? 10,
        search: query.search ?? ''
      }
    })
  }

  return {
    getPendingBusinesses
  }
}
