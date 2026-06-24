<template>
  <div class="min-h-screen flex flex-col bg-slate-50 text-slate-900 font-sans">
    <!-- Premium Navigation Bar, hidden on Login/Register pages -->
    <header v-if="showNavigation" class="bg-white/80 backdrop-blur-md border-b border-slate-200/80 sticky top-0 z-40 transition-all duration-200">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex items-center space-x-8">
            <NuxtLink to="/" class="flex items-center space-x-2">
              <span class="h-9 w-9 rounded-lg bg-gradient-to-tr from-brand-600 to-accent-500 flex items-center justify-center text-white font-bold text-lg shadow-md shadow-brand-500/10">
                Q
              </span>
              <span class="font-extrabold text-xl tracking-tight bg-gradient-to-r from-brand-700 to-slate-900 bg-clip-text text-transparent">
                QuickWork
              </span>
            </NuxtLink>

            <!-- Navigation Links -->
            <nav class="hidden md:flex space-x-1">
              <!-- Shared Public Links -->
              <NuxtLink
                to="/jobs"
                class="px-3 py-2 rounded-lg text-sm font-medium hover:bg-slate-100 hover:text-slate-900 transition-colors"
                active-class="bg-brand-50 text-brand-700 font-semibold"
              >
                Explore Jobs
              </NuxtLink>

              <!-- Student specific navigation links -->
              <template v-if="isAuthenticated && userRole === 'student'">
                <NuxtLink
                  to="/student/dashboard"
                  class="px-3 py-2 rounded-lg text-sm font-medium hover:bg-slate-100 hover:text-slate-900 transition-colors"
                  active-class="bg-brand-50 text-brand-700 font-semibold"
                >
                  My Dashboard
                </NuxtLink>
              </template>

              <!-- Business specific navigation links -->
              <template v-if="isAuthenticated && userRole === 'business'">
                <NuxtLink
                  to="/business/dashboard"
                  class="px-3 py-2 rounded-lg text-sm font-medium hover:bg-slate-100 hover:text-slate-900 transition-colors"
                  active-class="bg-brand-50 text-brand-700 font-semibold"
                >
                  Employer Console
                </NuxtLink>
              </template>
            </nav>
          </div>

          <!-- Right side Auth elements -->
          <div class="flex items-center space-x-4">
            <template v-if="isAuthenticated">
              <!-- Logged In Status -->
              <div class="hidden sm:flex flex-col text-right">
                <span class="text-xs font-semibold uppercase tracking-wider text-slate-400">
                  {{ userRole === 'student' ? 'Student' : 'Employer' }}
                </span>
                <span class="text-sm font-medium text-slate-700">{{ userEmail || 'Active User' }}</span>
              </div>

              <!-- Profile Dropdown trigger / Logout button -->
              <button
                @click="handleLogout"
                class="inline-flex items-center space-x-1.5 px-3 py-2 border border-slate-200 rounded-lg text-sm font-semibold text-slate-700 hover:bg-slate-50 hover:text-slate-950 focus-ring"
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
                to="/login"
                class="text-sm font-semibold text-slate-700 hover:text-slate-950 px-3 py-2"
              >
                Sign In
              </NuxtLink>
              <NuxtLink
                to="/register"
                class="inline-flex items-center justify-center py-2 px-4 border border-transparent text-sm font-semibold rounded-lg text-white bg-gradient-to-r from-brand-600 to-brand-500 hover:from-brand-500 hover:to-brand-400 shadow-md shadow-brand-500/10 focus-ring"
              >
                Sign Up
              </NuxtLink>
            </template>

            <!-- Mobile menu button -->
            <button
              @click="mobileMenuOpen = !mobileMenuOpen"
              class="md:hidden p-2 rounded-lg text-slate-500 hover:bg-slate-100 focus:outline-none focus:ring-2 focus:ring-brand-500"
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
      <div v-if="mobileMenuOpen" class="md:hidden border-t border-slate-200 bg-white px-2 pt-2 pb-3 space-y-1 shadow-lg">
        <NuxtLink
          to="/jobs"
          @click="mobileMenuOpen = false"
          class="block px-3 py-2 rounded-lg text-base font-medium hover:bg-slate-100 text-slate-700"
          active-class="bg-brand-50 text-brand-700 font-semibold"
        >
          Explore Jobs
        </NuxtLink>
        <template v-if="isAuthenticated && userRole === 'student'">
          <NuxtLink
            to="/student/dashboard"
            @click="mobileMenuOpen = false"
            class="block px-3 py-2 rounded-lg text-base font-medium hover:bg-slate-100 text-slate-700"
            active-class="bg-brand-50 text-brand-700 font-semibold"
          >
            My Dashboard
          </NuxtLink>
        </template>
        <template v-if="isAuthenticated && userRole === 'business'">
          <NuxtLink
            to="/business/dashboard"
            @click="mobileMenuOpen = false"
            class="block px-3 py-2 rounded-lg text-base font-medium hover:bg-slate-100 text-slate-700"
            active-class="bg-brand-50 text-brand-700 font-semibold"
          >
            Employer Console
          </NuxtLink>
        </template>
      </div>
    </header>

    <!-- Main Content App Router -->
    <main class="flex-grow">
      <NuxtPage />
    </main>

    <!-- Simple Modern Footer -->
    <footer v-if="showNavigation" class="bg-slate-900 text-slate-400 py-6 border-t border-slate-800">
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

// Hide navigation bar and footer on auth screens
const showNavigation = computed(() => {
  return !['/login', '/register'].includes(route.path)
})

const userEmail = computed(() => user.value?.email)

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
