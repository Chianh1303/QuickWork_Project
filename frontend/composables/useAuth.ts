import { computed } from 'vue'
import { useAuthStore } from '~/stores/auth'
import { useApi } from '~/composables/useApi'

export const useAuth = () => {
  const authStore = useAuthStore()
  const api = useApi()
  
  const token = useCookie<string | null>('auth_token', {
    maxAge: 60 * 60 * 24 * 3, // 3 days
    path: '/',
    sameSite: 'lax'
  })

  /**
   * Helper to decode payload from JWT token directly as instant fallback.
   */
  const decodeJwtPayload = (jwtString: string): any => {
    try {
      const parts = jwtString.split('.')
      if (parts.length < 2) return null
      const base64Url = parts[1]
      const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/')
      const jsonPayload = decodeURIComponent(
        atob(base64)
          .split('')
          .map((c) => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2))
          .join('')
      )
      return JSON.parse(jsonPayload)
    } catch {
      return null
    }
  }

  /**
   * Log in user, store JWT token in cookie, update store state, and redirect based on role.
   */
  const login = async (credentials: { email: string; password: string }) => {
    try {
      const response = await api.post('/api/auth/login', credentials)
      
      if (response && response.token) {
        token.value = response.token
        
        authStore.setUser({
          id: response.user.id,
          email: response.user.email,
          role: response.user.role
        })

        // Route by role
        await redirectByUserRole(response.user.role)
      }
      return response
    } catch (error: any) {
      console.error('Login error:', error)
      throw error
    }
  }

  /**
   * Log in via Google OAuth, store JWT token in cookie, update store state, and redirect.
   */
  const loginWithGoogle = async (googleData: { id_token?: string; email: string; name?: string; picture?: string; role?: string }) => {
    try {
      const response = await api.post('/api/auth/google-login', googleData)
      
      if (response && response.token) {
        token.value = response.token
        
        authStore.setUser({
          id: response.user.id,
          email: response.user.email,
          role: response.user.role
        })

        // Route by role
        await redirectByUserRole(response.user.role)
      }
      return response
    } catch (error: any) {
      console.error('Google login error:', error)
      throw error
    }
  }

  /**
   * Register a new user (student or business).
   */
  const register = async (registrationData: any) => {
    try {
      // Clear old token to ensure clean guest registration
      token.value = null
      authStore.clearAuth()
      const response = await api.post('/api/auth/register', registrationData)
      return response
    } catch (error: any) {
      console.error('Registration error:', error)
      throw error
    }
  }

  /**
   * Sign out the user, clear local state and cookies, and redirect to /login.
   */
  const logout = async () => {
    authStore.clearAuth()
    await navigateTo('/login')
  }

  /**
   * Fetch current user info using JWT token.
   */
  const fetchUser = async () => {
    if (!token.value) {
      authStore.clearAuth()
      return null
    }

    // Try decoding role and user_id directly from JWT token payload as instant fallback
    if (!authStore.user) {
      const payload = decodeJwtPayload(token.value)
      if (payload && payload.user_id && payload.role) {
        authStore.setUser({
          id: payload.user_id,
          email: payload.email || '',
          role: payload.role
        })
      }
    }

    try {
      const response = await api.get('/api/users/me', { skipAutoLogout: true })
      if (response && response.user_id) {
        authStore.setUser({
          id: response.user_id,
          email: authStore.user?.email || '',
          role: response.role
        })
      }
      return authStore.user
    } catch (error: any) {
      // ONLY clear auth if the server explicitly tells us the token is invalid/expired (401)
      const status = error?.response?.status || error?.status || error?.statusCode
      if (status === 401) {
        authStore.clearAuth()
        return null
      }
      // If it's a network glitch or temporary 5xx server warmup, KEEP current session!
      return authStore.user
    }
  }

  /**
   * Direct the user to their respective dashboard depending on their role.
   */
  const redirectByUserRole = async (role: string) => {
    if (role === 'student') {
      await navigateTo('/student/dashboard')
    } else if (role === 'business') {
      await navigateTo('/business/dashboard')
    } else if (role === 'admin') {
      await navigateTo('/admin/dashboard')
    } else {
      await navigateTo('/')
    }
  }

  return {
    login,
    loginWithGoogle,
    register,
    logout,
    fetchUser,
    redirectByUserRole,
    user: computed(() => authStore.user),
    isAuthenticated: computed(() => authStore.isAuthenticated),
    userRole: computed(() => authStore.userRole)
  }
}
