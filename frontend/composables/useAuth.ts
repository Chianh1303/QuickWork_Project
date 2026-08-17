import { computed } from 'vue'
import { useAuthStore } from '~/stores/auth'
import { useApi } from '~/composables/useApi'

export const useAuth = () => {
  const authStore = useAuthStore()
  const api = useApi()
  
  const token = useCookie<string | null>('auth_token', {
    maxAge: 60 * 60 * 24 * 3, // 3 days
    path: '/',
    sameSite: 'lax',
    secure: process.env.NODE_ENV === 'production'
  })

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
      authStore.setUser(null)
      return null
    }

    try {
      const response = await api.get('/api/users/me')
      if (response && response.user_id) {
        authStore.setUser({
          id: response.user_id,
          email: '', // Not returned by endpoint, but ID and role are populated
          role: response.role
        })
      }
      return authStore.user
    } catch (error) {
      authStore.clearAuth()
      return null
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
    register,
    logout,
    fetchUser,
    redirectByUserRole,
    user: computed(() => authStore.user),
    isAuthenticated: computed(() => authStore.isAuthenticated),
    userRole: computed(() => authStore.userRole)
  }
}
