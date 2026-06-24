import type { FetchOptions } from 'ofetch'

export const useApi = () => {
  // Store JWT in cookie for 3 days (matching backend JWT expiration)
  const token = useCookie<string | null>('auth_token', {
    maxAge: 60 * 60 * 24 * 3, // 3 days
    path: '/',
    sameSite: 'lax',
    secure: process.env.NODE_ENV === 'production'
  })

  /**
   * Main request runner using Nuxt's $fetch with automatic JWT attachment.
   * Leverages the routeRules proxy configured in nuxt.config.ts.
   */
  const request = async <T = any>(url: string, options: FetchOptions = {}): Promise<T> => {
    // Standardize headers
    const headers: Record<string, string> = {
      Accept: 'application/json',
      ...((options.headers as Record<string, string>) || {})
    }

    // Automatically append Bearer token if it exists in cookies
    if (token.value) {
      headers['Authorization'] = `Bearer ${token.value}`
    }

    try {
      return await $fetch<T>(url, {
        ...options,
        headers
      })
    } catch (error: any) {
      // If we get an Unauthorized response, clear token and redirect
      if (error.response?.status === 401) {
        token.value = null
        if (import.meta.client) {
          navigateTo('/login')
        }
      }
      throw error
    }
  }

  return {
    request,
    get: <T = any>(url: string, options?: FetchOptions) => request<T>(url, { method: 'GET', ...options }),
    post: <T = any>(url: string, body?: any, options?: FetchOptions) => request<T>(url, { method: 'POST', body, ...options }),
    put: <T = any>(url: string, body?: any, options?: FetchOptions) => request<T>(url, { method: 'PUT', body, ...options }),
    delete: <T = any>(url: string, options?: FetchOptions) => request<T>(url, { method: 'DELETE', ...options })
  }
}
