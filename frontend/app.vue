<template>
  <div class="app-shell min-h-screen flex flex-col text-slate-900 font-sans">
    <!-- Premium Navigation Bar -->
    <header class="bg-slate-950/95 backdrop-blur-md border-b border-cyan-400/10 sticky top-0 z-40 transition-all duration-200 shadow-lg shadow-slate-950/10">
      <div v-if="isLandingRoute" class="hidden border-b border-white/10 bg-slate-900/80 text-xs font-semibold text-slate-300 lg:block">
        <div class="mx-auto flex max-w-7xl items-center justify-between px-4 py-2 sm:px-6 lg:px-8">
          <div class="flex items-center gap-5">
            <span>Phone: 0123 456 789</span>
            <span>Email: support@quickwork.vn</span>
            <span>Address: Ho Chi Minh City, Vietnam</span>
          </div>
          <div class="flex items-center gap-4">
            <span>Facebook</span>
            <span>LinkedIn</span>
          </div>
        </div>
      </div>
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
                <template v-if="isLandingRoute">
                  <a
                    v-for="item in landingNavItems"
                    :key="item.id"
                    :href="item.href"
                    @click.prevent="scrollToLandingSection(item.id)"
                    :class="[
                      activeLandingSection === item.id ? 'text-cyan-200' : 'text-slate-300 hover:text-white',
                      'group relative px-3 py-2 rounded-lg text-sm font-medium hover:bg-white/10 transition-colors'
                    ]"
                  >
                    <span>{{ item.name }}</span>
                    <span
                      :class="activeLandingSection === item.id ? 'scale-x-100 opacity-100' : 'scale-x-0 opacity-0 group-hover:scale-x-75 group-hover:opacity-60'"
                      class="absolute bottom-1 left-3 right-3 h-0.5 origin-center rounded-full bg-cyan-300 transition-all duration-300 ease-out"
                    ></span>
                  </a>
                </template>
                <NuxtLink
                  v-else
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
              <NuxtLink
                v-if="!isDashboardRoute"
                :to="dashboardPath"
                class="hidden sm:inline-flex items-center justify-center rounded-lg bg-cyan-400 px-4 py-2 text-sm font-extrabold text-slate-950 hover:bg-cyan-300"
              >
                Dashboard
              </NuxtLink>
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
                Register
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
          <template v-if="isLandingRoute">
            <a
              v-for="item in landingNavItems"
              :key="item.id"
              :href="item.href"
              @click.prevent="scrollToLandingSection(item.id); mobileMenuOpen = false"
              :class="[
                activeLandingSection === item.id
                  ? 'bg-cyan-400/10 text-cyan-200 ring-1 ring-cyan-400/20'
                  : 'text-slate-300 hover:bg-white/10 hover:text-white',
                'relative block px-3 py-2 rounded-lg text-base font-medium transition-colors'
              ]"
            >
              {{ item.name }}
              <span
                :class="activeLandingSection === item.id ? 'opacity-100 translate-x-0' : 'opacity-0 -translate-x-2'"
                class="absolute bottom-1 left-3 h-0.5 w-10 rounded-full bg-cyan-300 transition-all duration-300"
              ></span>
            </a>
          </template>
          <NuxtLink
            v-else
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
import { ref, computed, onMounted, onBeforeUnmount, nextTick, watch } from 'vue'
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

const landingNavItems = [
  { id: 'home', name: 'Home', href: '#home' },
  { id: 'explore-jobs', name: 'Explore Jobs', href: '#explore-jobs' },
  { id: 'about', name: 'About', href: '#about' },
  { id: 'features', name: 'Features', href: '#features' },
  { id: 'contact', name: 'Contact', href: '#contact' }
]

const activeLandingSection = ref('home')
let landingObserver: IntersectionObserver | null = null

// Keep auth pages focused while preserving the global header.
const showFooter = computed(() => {
  return !['/', '/login', '/register', '/employer-register'].includes(route.path)
})

const userEmail = computed(() => user.value?.email)
const isLandingRoute = computed(() => route.path === '/')
const isStudentDashboard = computed(() => route.path === '/student/dashboard')
const isBusinessDashboard = computed(() => route.path === '/business/dashboard')
const isDashboardRoute = computed(() => isStudentDashboard.value || isBusinessDashboard.value)
const dashboardPath = computed(() => {
  return userRole.value === 'business' ? '/business/dashboard' : '/student/dashboard'
})
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

const getHeaderOffset = () => {
  if (!import.meta.client) return 80
  const header = document.querySelector('header')
  return (header?.getBoundingClientRect().height || 80) + 8
}

const scrollToLandingSection = async (sectionId: string) => {
  activeLandingSection.value = sectionId

  if (!isLandingRoute.value) {
    await navigateTo(`/#${sectionId}`)
    return
  }

  await nextTick()
  const target = document.getElementById(sectionId)
  if (!target) return

  const top = target.getBoundingClientRect().top + window.scrollY - getHeaderOffset()
  window.history.replaceState(null, '', `#${sectionId}`)
  window.scrollTo({ top, behavior: 'smooth' })
}

const setupLandingObserver = async () => {
  if (!import.meta.client) return

  landingObserver?.disconnect()
  landingObserver = null

  if (!isLandingRoute.value) return

  await nextTick()

  const sections = landingNavItems
    .map(item => document.getElementById(item.id))
    .filter((section): section is HTMLElement => Boolean(section))

  if (sections.length === 0) return

  const observerOptions = {
    root: null,
    rootMargin: `-${getHeaderOffset()}px 0px -55% 0px`,
    threshold: [0.12, 0.35, 0.6]
  }

  landingObserver = new IntersectionObserver(entries => {
    const visible = entries
      .filter(entry => entry.isIntersecting)
      .sort((a, b) => b.intersectionRatio - a.intersectionRatio)[0]

    if (visible?.target?.id) {
      activeLandingSection.value = visible.target.id
    }
  }, observerOptions)

  sections.forEach(section => landingObserver?.observe(section))

  const hashSection = window.location.hash.replace('#', '')
  if (landingNavItems.some(item => item.id === hashSection)) {
    activeLandingSection.value = hashSection
  }
}

// Fetch user profile on layout mounting to ensure session state hydration
onMounted(async () => {
  try {
    await fetchUser()
  } catch (e) {
    console.error('Session restoration failed:', e)
  }

  await setupLandingObserver()
})

watch(() => route.fullPath, async () => {
  mobileMenuOpen.value = false
  await setupLandingObserver()
})

onBeforeUnmount(() => {
  landingObserver?.disconnect()
})

const handleLogout = async () => {
  await logout()
}
</script>
