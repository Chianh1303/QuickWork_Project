import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export interface UserState {
  id: number
  email: string
  role: 'student' | 'business' | 'admin'
}

export const useAuthStore = defineStore('auth', () => {
  const tokenCookie = useCookie<string | null>('auth_token', {
    maxAge: 60 * 60 * 24 * 3,
    path: '/',
    sameSite: 'lax'
  })

  const userCookie = useCookie<UserState | null>('auth_user', {
    maxAge: 60 * 60 * 24 * 3,
    path: '/',
    sameSite: 'lax'
  })

  const user = ref<UserState | null>(userCookie.value || null)

  // Reactive state getters
  const isAuthenticated = computed(() => !!tokenCookie.value)
  const userRole = computed(() => user.value?.role || userCookie.value?.role || null)

  /**
   * Set user information in state and persist to cookie.
   */
  function setUser(userData: UserState | null) {
    user.value = userData
    userCookie.value = userData
  }

  /**
   * Clear authentication state and remove cookies.
   */
  function clearAuth() {
    user.value = null
    userCookie.value = null
    tokenCookie.value = null
  }

  return {
    user,
    isAuthenticated,
    userRole,
    setUser,
    clearAuth
  }
})
