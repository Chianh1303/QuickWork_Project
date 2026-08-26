export interface PendingBusinessItem {
  business_id: number
  company_name: string
  tax_code: string
  email: string
  status: string
  created_at: string
}

export interface PaginationMeta {
  page: number
  limit: number
  total: number
  total_pages: number
}

export interface PendingBusinessListResponse {
  items: PendingBusinessItem[]
  pagination: PaginationMeta
}

export interface AdminDashboardStats {
  total_students: number
  total_businesses: number
  pending_businesses: number
  total_jobs: number
  pending_jobs: number
  total_disbursed: number
}

export interface BusinessKYBDetail {
  business_id: number
  user_id: number
  company_name: string
  tax_code: string
  email: string
  phone: string
  address: string
  logo_url: string
  status: string
  is_verified: boolean
  reject_reason: string
  created_at: string
  reviewed_at: string | null
}

export interface ReviewBusinessPayload {
  decision: 'approved' | 'rejected'
  reject_reason: string
}
