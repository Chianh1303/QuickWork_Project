<template>
  <div class="relative bg-slate-900 overflow-hidden text-white flex-grow flex flex-col justify-center min-h-[calc(100vh-4rem)]">
    <!-- Background visual assets -->
    <div class="absolute inset-0 bg-[radial-gradient(ellipse_at_top_right,_var(--tw-gradient-stops))] from-brand-900/40 via-slate-950 to-slate-950"></div>
    <div class="absolute top-0 right-0 w-1/2 h-1/2 bg-gradient-to-br from-brand-500/10 to-transparent blur-3xl pointer-events-none"></div>

    <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-20 lg:py-32 flex flex-col lg:flex-row items-center justify-between gap-12">
      <!-- Left: Copywriting -->
      <div class="max-w-2xl text-center lg:text-left space-y-8">
        <div class="inline-flex items-center space-x-2 bg-slate-800/80 border border-slate-700/50 rounded-full px-3 py-1 text-xs sm:text-sm text-brand-300">
          <span class="flex h-2 w-2 rounded-full bg-brand-400 animate-pulse"></span>
          <span>Connecting top student talent with forward-thinking companies.</span>
        </div>
        <h1 class="text-4xl sm:text-6xl font-extrabold tracking-tight text-white leading-none">
          Find your next <br />
          <span class="bg-gradient-to-r from-brand-400 via-brand-200 to-accent-300 bg-clip-text text-transparent">
            QuickWork opportunity
          </span>
        </h1>
        <p class="text-lg text-slate-300 max-w-lg mx-auto lg:mx-0">
          The instant gig and job platform built for modern students and agile businesses. Create an account, publish listings, and start working immediately.
        </p>

        <!-- Dynamic Action Buttons -->
        <div class="flex flex-col sm:flex-row justify-center lg:justify-start gap-4">
          <template v-if="isAuthenticated">
            <button
              @click="goToDashboard"
              class="inline-flex items-center justify-center px-8 py-3.5 border border-transparent text-base font-semibold rounded-xl text-white bg-gradient-to-r from-brand-600 to-brand-500 hover:from-brand-500 hover:to-brand-400 shadow-lg shadow-brand-500/20 focus-ring"
            >
              Go to Portal Console
            </button>
          </template>
          <template v-else>
            <NuxtLink
              to="/register"
              class="inline-flex items-center justify-center px-8 py-3.5 border border-transparent text-base font-semibold rounded-xl text-white bg-gradient-to-r from-brand-600 to-brand-500 hover:from-brand-500 hover:to-brand-400 shadow-lg shadow-brand-500/20 focus-ring"
            >
              Find Student Work
            </NuxtLink>
            <NuxtLink
              to="/employer-register"
              class="inline-flex items-center justify-center px-8 py-3.5 border border-slate-700 hover:border-slate-500 text-base font-semibold rounded-xl text-slate-200 bg-slate-800/40 hover:bg-slate-800 transition-colors focus-ring"
            >
              Hire Student Talent
            </NuxtLink>
          </template>
        </div>

        <!-- Trust factors / Stats -->
        <div class="grid grid-cols-3 gap-6 pt-8 border-t border-slate-800/60 max-w-md mx-auto lg:mx-0">
          <div>
            <div class="text-2xl sm:text-3xl font-bold text-white">10K+</div>
            <div class="text-xs text-slate-400">Students Registered</div>
          </div>
          <div>
            <div class="text-2xl sm:text-3xl font-bold text-white">500+</div>
            <div class="text-xs text-slate-400">Trusted Employers</div>
          </div>
          <div>
            <div class="text-2xl sm:text-3xl font-bold text-white">4.9/5</div>
            <div class="text-xs text-slate-400">Match Rating</div>
          </div>
        </div>
      </div>

      <!-- Right: Visual Cards Stack -->
      <div class="relative w-full max-w-md lg:max-w-lg hidden lg:block">
        <div class="absolute -inset-1 rounded-2xl bg-gradient-to-r from-brand-500 to-accent-500 opacity-20 blur-xl"></div>
        <div class="relative bg-slate-900 border border-slate-800/80 rounded-2xl p-6 shadow-2xl space-y-6">
          <h3 class="text-lg font-bold text-white flex items-center gap-2">
            <svg class="h-5 w-5 text-brand-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
            Active Listings Feed
          </h3>

          <div class="space-y-4">
            <!-- Simulated Listing 1 -->
            <div class="p-4 rounded-xl bg-slate-950/60 border border-slate-800 flex justify-between items-start gap-4">
              <div>
                <span class="text-xs font-semibold text-brand-400 uppercase">Tech Corp</span>
                <h4 class="font-bold text-white text-sm">Fullstack Intern (Nuxt/Go)</h4>
                <p class="text-xs text-slate-400 mt-1">Ho Chi Minh City • Remote ok</p>
              </div>
              <div class="text-right">
                <span class="text-xs font-bold text-emerald-400 bg-emerald-500/10 px-2.5 py-1 rounded-full">$15/hr</span>
              </div>
            </div>

            <!-- Simulated Listing 2 -->
            <div class="p-4 rounded-xl bg-slate-950/60 border border-slate-800 flex justify-between items-start gap-4">
              <div>
                <span class="text-xs font-semibold text-accent-400 uppercase">Design Studio</span>
                <h4 class="font-bold text-white text-sm">UI/UX Designer Helper</h4>
                <p class="text-xs text-slate-400 mt-1">Hanoi • Part-time</p>
              </div>
              <div class="text-right">
                <span class="text-xs font-bold text-emerald-400 bg-emerald-500/10 px-2.5 py-1 rounded-full">$12/hr</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useAuth } from '~/composables/useAuth'

const { isAuthenticated, userRole, redirectByUserRole } = useAuth()

const goToDashboard = async () => {
  if (userRole.value) {
    await redirectByUserRole(userRole.value)
  }
}

onMounted(async () => {
  if (isAuthenticated.value && userRole.value) {
    await redirectByUserRole(userRole.value)
  }
})
</script>
