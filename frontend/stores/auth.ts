import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export interface UserState {
  id: number
  email: string
  role: 'student' | 'business' | 'admin'
}

export const useAuthStore = defineStore('auth', () => {
  const user = ref<UserState | null>(null)
  
  // Retrieve token cookie
  const tokenCookie = useCookie<string | null>('auth_token')

  // Reactive state getters
  const isAuthenticated = computed(() => !!tokenCookie.value)
  const userRole = computed(() => user.value?.role || null)

  /**
   * Set user information in state.
   */
  function setUser(userData: UserState | null) {
    user.value = userData
  }

  /**
   * Clear authentication state and remove cookies.
   */
  function clearAuth() {
    user.value = null
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
