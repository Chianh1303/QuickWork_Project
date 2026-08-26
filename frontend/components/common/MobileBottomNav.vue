<template>
  <nav class="fixed bottom-0 left-0 right-0 z-50 md:hidden bg-slate-950/95 backdrop-blur-lg border-t border-cyan-400/15 shadow-2xl shadow-cyan-950/50 px-2 py-1.5">
    <div class="grid grid-cols-5 items-center text-center">
      <template v-if="isAuthenticated && userRole === 'student'">
        <button
          @click="navigateToTab('jobs')"
          class="flex flex-col items-center justify-center py-1 px-1 transition-colors"
          :class="isActive('jobs') ? 'text-cyan-300 font-bold' : 'text-slate-400 hover:text-slate-200'"
        >
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
          </svg>
          <span class="text-[10px] mt-0.5 tracking-tight">Tìm việc</span>
        </button>

        <button
          @click="navigateToTab('applications')"
          class="flex flex-col items-center justify-center py-1 px-1 transition-colors"
          :class="isActive('applications') ? 'text-cyan-300 font-bold' : 'text-slate-400 hover:text-slate-200'"
        >
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          <span class="text-[10px] mt-0.5 tracking-tight">Đơn hàng</span>
        </button>

        <!-- Quick Check-in Center Button -->
        <button
          @click="navigateToTab('applications')"
          class="flex flex-col items-center justify-center -mt-4 transition-transform active:scale-95"
        >
          <div class="h-12 w-12 rounded-full bg-gradient-to-tr from-cyan-400 to-emerald-400 flex items-center justify-center text-slate-950 shadow-lg shadow-cyan-500/30 ring-4 ring-slate-950">
            <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <span class="text-[10px] font-bold text-cyan-300 mt-0.5">Điểm danh</span>
        </button>

        <button
          @click="navigateToTab('wallet')"
          class="flex flex-col items-center justify-center py-1 px-1 transition-colors"
          :class="isActive('wallet') ? 'text-cyan-300 font-bold' : 'text-slate-400 hover:text-slate-200'"
        >
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
          </svg>
          <span class="text-[10px] mt-0.5 tracking-tight">Ví lương</span>
        </button>

        <button
          @click="navigateToTab('profile')"
          class="flex flex-col items-center justify-center py-1 px-1 transition-colors"
          :class="isActive('profile') ? 'text-cyan-300 font-bold' : 'text-slate-400 hover:text-slate-200'"
        >
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
          </svg>
          <span class="text-[10px] mt-0.5 tracking-tight">Hồ sơ</span>
        </button>
      </template>

      <template v-else-if="isAuthenticated && userRole === 'business'">
        <button
          @click="navigateToTab('dashboard')"
          class="flex flex-col items-center justify-center py-1 px-1 transition-colors"
          :class="isActive('dashboard') ? 'text-cyan-300 font-bold' : 'text-slate-400 hover:text-slate-200'"
        >
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
          </svg>
          <span class="text-[10px] mt-0.5 tracking-tight">Tổng quan</span>
        </button>

        <button
          @click="navigateToTab('jobs')"
          class="flex flex-col items-center justify-center py-1 px-1 transition-colors"
          :class="isActive('jobs') ? 'text-cyan-300 font-bold' : 'text-slate-400 hover:text-slate-200'"
        >
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
          </svg>
          <span class="text-[10px] mt-0.5 tracking-tight">Đăng tin</span>
        </button>

        <button
          @click="navigateToTab('applicants')"
          class="flex flex-col items-center justify-center py-1 px-1 transition-colors"
          :class="isActive('applicants') ? 'text-cyan-300 font-bold' : 'text-slate-400 hover:text-slate-200'"
        >
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
          <span class="text-[10px] mt-0.5 tracking-tight">Ứng viên</span>
        </button>

        <button
          @click="navigateToTab('profile')"
          class="flex flex-col items-center justify-center py-1 px-1 transition-colors"
          :class="isActive('profile') ? 'text-cyan-300 font-bold' : 'text-slate-400 hover:text-slate-200'"
        >
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
          </svg>
          <span class="text-[10px] mt-0.5 tracking-tight">Công ty</span>
        </button>

        <NuxtLink
          to="/"
          class="flex flex-col items-center justify-center py-1 px-1 text-slate-400 hover:text-slate-200 transition-colors"
        >
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
          </svg>
          <span class="text-[10px] mt-0.5 tracking-tight">Trang chủ</span>
        </NuxtLink>
      </template>

      <template v-else>
        <NuxtLink
          to="/"
          class="flex flex-col items-center justify-center py-1 px-1 transition-colors"
          :class="route.path === '/' ? 'text-cyan-300 font-bold' : 'text-slate-400 hover:text-slate-200'"
        >
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
          </svg>
          <span class="text-[10px] mt-0.5 tracking-tight">Trang chủ</span>
        </NuxtLink>

        <NuxtLink
          to="/#explore-jobs"
          class="flex flex-col items-center justify-center py-1 px-1 transition-colors"
          :class="route.path === '/jobs' ? 'text-cyan-300 font-bold' : 'text-slate-400 hover:text-slate-200'"
        >
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
          <span class="text-[10px] mt-0.5 tracking-tight">Tìm việc</span>
        </NuxtLink>

        <NuxtLink
          to="/login"
          class="flex flex-col items-center justify-center col-span-2 py-1 px-3 bg-cyan-400 rounded-lg text-slate-950 font-bold text-xs shadow-md shadow-cyan-500/20"
        >
          Đăng nhập
        </NuxtLink>

        <NuxtLink
          to="/register"
          class="flex flex-col items-center justify-center py-1 px-1 text-slate-400 hover:text-slate-200 transition-colors"
        >
          <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z" />
          </svg>
          <span class="text-[10px] mt-0.5 tracking-tight">Đăng ký</span>
        </NuxtLink>
      </template>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import { useAuth } from '~/composables/useAuth'

const route = useRoute()
const { isAuthenticated, userRole } = useAuth()

const studentDashboardActiveSection = useState<string>('studentDashboardActiveSection', () => 'jobs')
const businessDashboardActiveSection = useState<string>('businessDashboardActiveSection', () => 'dashboard')

const isActive = (tabId: string) => {
  if (userRole.value === 'student' && route.path === '/student/dashboard') {
    return studentDashboardActiveSection.value === tabId
  }
  if (userRole.value === 'business' && route.path === '/business/dashboard') {
    return businessDashboardActiveSection.value === tabId
  }
  return false
}

const navigateToTab = (tabId: string) => {
  if (userRole.value === 'student') {
    studentDashboardActiveSection.value = tabId
    if (route.path !== '/student/dashboard') {
      navigateTo('/student/dashboard')
    }
  } else if (userRole.value === 'business') {
    businessDashboardActiveSection.value = tabId
    if (route.path !== '/business/dashboard') {
      navigateTo('/business/dashboard')
    }
  }
}
</script>
