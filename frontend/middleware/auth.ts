import { useAuthStore } from '~/stores/auth'

export default defineNuxtRouteMiddleware(async (to, from) => {
  const authStore = useAuthStore()
  const { fetchUser } = useAuth()
  const token = useCookie('auth_token')

  const dashboardForRole = (role: string) => {
    if (role === 'business') return '/business/dashboard'
    if (role === 'admin') return '/admin/dashboard'
    return '/student/dashboard'
  }

  // Attempt to fetch current user profile if the cookie exists but store state is empty (e.g., page reload)
  if (token.value && !authStore.user) {
    await fetchUser()
  }

  const isAuthRoute = ['/login', '/register', '/employer-register'].includes(to.path)

  // 1. Unauthenticated users handling
  if (!token.value) {
    if (!isAuthRoute) {
      // Redirect to login if attempting to access protected dashboard/route
      return navigateTo('/login')
    }
  } 
  // 2. Authenticated users handling
  else {
    // Redirect authenticated users away from login/register to their dashboard
    if (isAuthRoute && authStore.user) {
      const role = authStore.user.role
      if (role === 'student') {
        return navigateTo('/student/dashboard')
      } else if (role === 'business') {
        return navigateTo('/business/dashboard')
      } else if (role === 'admin') {
        return navigateTo('/admin/dashboard')
      }
      return navigateTo('/')
    }

    // Role-based route guard checks
    if (authStore.user) {
      const role = authStore.user.role
      
      // Guard student routes
      if (to.path.startsWith('/student') && role !== 'student') {
        return navigateTo(dashboardForRole(role))
      }
      
      // Guard business routes
      if (to.path.startsWith('/business') && role !== 'business') {
        return navigateTo(dashboardForRole(role))
      }

      // Guard admin routes
      if (to.path.startsWith('/admin') && role !== 'admin') {
        return navigateTo(dashboardForRole(role))
      }
    }
  }
})
