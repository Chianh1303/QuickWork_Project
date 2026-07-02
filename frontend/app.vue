<template>
  <div class="app-shell min-h-screen flex flex-col text-slate-900 font-sans">
    <!-- Premium Navigation Bar -->
    <header class="bg-slate-950/95 backdrop-blur-md border-b border-cyan-400/10 sticky top-0 z-40 transition-all duration-200 shadow-lg shadow-slate-950/10">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex items-center space-x-8">
            <NuxtLink to="/" class="flex items-center space-x-2">
              <span class="h-9 w-9 rounded-lg bg-gradient-to-tr from-cyan-400 to-slate-700 flex items-center justify-center text-slate-950 font-bold text-lg shadow-md shadow-cyan-500/20">
                Q
              </span>
              <span class="font-extrabold text-xl tracking-tight bg-gradient-to-r from-white to-cyan-200 bg-clip-text text-transparent">
                QuickWork
              </span>
            </NuxtLink>

            <!-- Navigation Links -->
            <nav class="hidden md:flex items-center space-x-1">
              <template v-if="isDashboardRoute">
                <button
                  v-for="item in dashboardNavItems"
                  :key="item.id"
                  @click="setDashboardSection(item.id)"
                  :class="[
                    dashboardActiveSection === item.id
                      ? 'bg-cyan-400 text-slate-950 font-semibold shadow-sm shadow-cyan-950/30'
                      : 'text-slate-300 hover:bg-white/10 hover:text-white',
                    'px-3 py-2 rounded-lg text-sm font-medium transition-colors'
                  ]"
                >
                  {{ item.name }}
                </button>
              </template>

              <!-- Shared Public Links -->
              <template v-else>
                <NuxtLink
                  to="/jobs"
                  class="px-3 py-2 rounded-lg text-sm font-medium text-slate-300 hover:bg-white/10 hover:text-white transition-colors"
                  active-class="bg-cyan-400 text-slate-950 font-semibold"
                >
                  Explore Jobs
                </NuxtLink>

                <!-- Student specific navigation links -->
                <template v-if="isAuthenticated && userRole === 'student'">
                  <NuxtLink
                    to="/student/dashboard"
                    class="px-3 py-2 rounded-lg text-sm font-medium text-slate-300 hover:bg-white/10 hover:text-white transition-colors"
                    active-class="bg-cyan-400 text-slate-950 font-semibold"
                  >
                    My Dashboard
                  </NuxtLink>
                </template>

                <!-- Business specific navigation links -->
                <template v-if="isAuthenticated && userRole === 'business'">
                  <NuxtLink
                    to="/business/dashboard"
                    class="px-3 py-2 rounded-lg text-sm font-medium text-slate-300 hover:bg-white/10 hover:text-white transition-colors"
                    active-class="bg-cyan-400 text-slate-950 font-semibold"
                  >
                    Employer Console
                  </NuxtLink>
                </template>
              </template>
            </nav>
          </div>

          <!-- Right side Auth elements -->
          <div class="flex items-center space-x-4">
            <template v-if="isAuthenticated">
              <!-- Logged In Status -->
              <div class="hidden sm:flex flex-col text-right">
                <span class="text-xs font-semibold uppercase tracking-wider text-cyan-300">
                  {{ userRole === 'student' ? 'Student' : 'Employer' }}
                </span>
                <span class="text-sm font-medium text-slate-200">{{ userEmail || 'Active User' }}</span>
              </div>

              <!-- Profile Dropdown trigger / Logout button -->
              <button
                @click="handleLogout"
                class="inline-flex items-center space-x-1.5 px-3 py-2 border border-white/10 rounded-lg text-sm font-semibold text-slate-200 hover:bg-white/10 hover:text-white focus-ring"
              >
                <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                </svg>
                <span>Sign Out</span>
              </button>
            </template>

            <template v-else>
              <!-- Guest Links -->
              <NuxtLink
                v-if="route.path !== '/login'"
                to="/login"
                class="text-sm font-semibold text-slate-200 hover:text-white px-3 py-2"
              >
                Sign In
              </NuxtLink>
              <NuxtLink
                v-if="route.path !== '/register'"
                to="/register"
                class="inline-flex items-center justify-center py-2 px-4 border border-transparent text-sm font-semibold rounded-lg text-slate-950 bg-cyan-400 hover:bg-cyan-300 shadow-md shadow-cyan-500/10 focus-ring"
              >
                Sign Up
              </NuxtLink>
            </template>

            <!-- Mobile menu button -->
            <button
              @click="mobileMenuOpen = !mobileMenuOpen"
              class="md:hidden p-2 rounded-lg text-slate-300 hover:bg-white/10 focus:outline-none focus:ring-2 focus:ring-cyan-400"
            >
              <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path v-if="!mobileMenuOpen" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
                <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>
      </div>

      <!-- Mobile Menu -->
      <div v-if="mobileMenuOpen" class="md:hidden border-t border-white/10 bg-slate-950 px-2 pt-2 pb-3 space-y-1 shadow-lg">
        <template v-if="isDashboardRoute">
          <div class="px-3 py-2">
            <p class="text-xs font-semibold uppercase tracking-wider text-cyan-300">{{ dashboardWorkspaceLabel }}</p>
          </div>
          <button
            v-for="item in dashboardNavItems"
            :key="item.id"
            @click="setDashboardSection(item.id); mobileMenuOpen = false"
            :class="[
              dashboardActiveSection === item.id
                ? 'bg-cyan-400 text-slate-950 font-semibold'
                : 'text-slate-300 hover:bg-white/10 hover:text-white',
              'block w-full text-left px-3 py-2 rounded-lg text-base font-medium'
            ]"
          >
            {{ item.name }}
          </button>
        </template>

        <template v-else>
          <NuxtLink
            to="/jobs"
            @click="mobileMenuOpen = false"
            class="block px-3 py-2 rounded-lg text-base font-medium hover:bg-white/10 text-slate-300"
            active-class="bg-cyan-400 text-slate-950 font-semibold"
          >
            Explore Jobs
          </NuxtLink>

          <template v-if="isAuthenticated && userRole === 'student'">
            <NuxtLink
              to="/student/dashboard"
              @click="mobileMenuOpen = false"
              class="block px-3 py-2 rounded-lg text-base font-medium hover:bg-white/10 text-slate-300"
              active-class="bg-cyan-400 text-slate-950 font-semibold"
            >
              My Dashboard
            </NuxtLink>
          </template>
          <template v-if="isAuthenticated && userRole === 'business'">
            <NuxtLink
              to="/business/dashboard"
              @click="mobileMenuOpen = false"
              class="block px-3 py-2 rounded-lg text-base font-medium hover:bg-white/10 text-slate-300"
              active-class="bg-cyan-400 text-slate-950 font-semibold"
            >
              Employer Console
            </NuxtLink>
          </template>
        </template>
      </div>
    </header>

    <!-- Main Content App Router -->
    <main class="flex-grow">
      <NuxtPage />
    </main>

    <!-- Simple Modern Footer -->
    <footer v-if="showFooter" class="bg-slate-900 text-slate-400 py-6 border-t border-slate-800">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0 text-sm">
        <div class="flex items-center space-x-2">
          <span class="font-semibold text-white">QuickWork</span>
          <span>&copy; 2026. All rights reserved.</span>
        </div>
        <div class="flex space-x-6">
          <a href="#" class="hover:text-white transition-colors">Privacy Policy</a>
          <a href="#" class="hover:text-white transition-colors">Terms of Service</a>
          <a href="#" class="hover:text-white transition-colors">Contact Support</a>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useAuth } from '~/composables/useAuth'

const route = useRoute()
const { logout, user, isAuthenticated, userRole, fetchUser } = useAuth()

const mobileMenuOpen = ref(false)
const studentDashboardActiveSection = useState<string>('studentDashboardActiveSection', () => 'jobs')
const businessDashboardActiveSection = useState<string>('businessDashboardActiveSection', () => 'dashboard')

const studentDashboardNavItems = [
  { id: 'jobs', name: 'Find Jobs' },
  { id: 'profile', name: 'Profile' },
  { id: 'applications', name: 'Applications' },
  { id: 'wallet', name: 'Wallet' }
]

const businessDashboardNavItems = [
  { id: 'dashboard', name: 'Dashboard' },
  { id: 'profile', name: 'Company Profile' },
  { id: 'jobs', name: 'Jobs' },
  { id: 'applicants', name: 'Applicants' }
]

// Keep auth pages focused while preserving the global header.
const showFooter = computed(() => {
  return !['/login', '/register', '/employer-register'].includes(route.path)
})

const userEmail = computed(() => user.value?.email)
const isStudentDashboard = computed(() => route.path === '/student/dashboard')
const isBusinessDashboard = computed(() => route.path === '/business/dashboard')
const isDashboardRoute = computed(() => isStudentDashboard.value || isBusinessDashboard.value)
const dashboardNavItems = computed(() => {
  return isBusinessDashboard.value ? businessDashboardNavItems : studentDashboardNavItems
})
const dashboardActiveSection = computed(() => {
  return isBusinessDashboard.value
    ? businessDashboardActiveSection.value
    : studentDashboardActiveSection.value
})
const dashboardWorkspaceLabel = computed(() => {
  return isBusinessDashboard.value ? 'Employer Workspace' : 'Student Workspace'
})

const setDashboardSection = (sectionId: string) => {
  if (isBusinessDashboard.value) {
    businessDashboardActiveSection.value = sectionId
    return
  }

  studentDashboardActiveSection.value = sectionId
}

// Fetch user profile on layout mounting to ensure session state hydration
onMounted(async () => {
  try {
    await fetchUser()
  } catch (e) {
    console.error('Session restoration failed:', e)
  }
})

const handleLogout = async () => {
  await logout()
}
</script>
