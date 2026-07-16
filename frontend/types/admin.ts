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
