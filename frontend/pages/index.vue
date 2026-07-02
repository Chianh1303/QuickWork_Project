<template>
  <div class="relative overflow-hidden text-white flex-grow min-h-[calc(100vh-4rem)] bg-slate-950">
    <div class="absolute inset-0 bg-[radial-gradient(circle_at_15%_12%,rgba(34,211,238,0.18),transparent_28rem),radial-gradient(circle_at_85%_20%,rgba(14,116,144,0.18),transparent_30rem),linear-gradient(180deg,#020617_0%,#0f172a_52%,#020617_100%)]"></div>
    <div class="absolute inset-x-0 top-0 h-px bg-gradient-to-r from-transparent via-cyan-300/40 to-transparent"></div>

    <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16 lg:py-24">
      <div class="grid lg:grid-cols-[0.95fr_1.05fr] gap-10 lg:gap-14 items-center">
        <section class="text-center lg:text-left space-y-7">
          <div class="inline-flex items-center gap-2 rounded-full border border-cyan-400/20 bg-cyan-400/10 px-3 py-1 text-xs sm:text-sm font-bold text-cyan-200">
            <span class="flex h-2 w-2 rounded-full bg-cyan-300 animate-pulse"></span>
            Student work and employer hiring in one workspace
          </div>

          <div class="space-y-5">
            <h1 class="text-4xl sm:text-6xl lg:text-7xl font-extrabold tracking-tight text-white leading-[0.95]">
              Start working faster with
              <span class="block bg-gradient-to-r from-cyan-300 via-white to-cyan-200 bg-clip-text text-transparent">
                QuickWork
              </span>
            </h1>
            <p class="text-base sm:text-lg text-slate-300 max-w-xl mx-auto lg:mx-0 leading-8 font-medium">
              A focused job marketplace for students and businesses: discover roles, manage applications, chat with HR, and track accepted work from one modern dashboard.
            </p>
          </div>

          <div class="flex flex-col sm:flex-row justify-center lg:justify-start gap-3">
          <template v-if="isAuthenticated">
            <button
              @click="goToDashboard"
              class="inline-flex items-center justify-center px-7 py-3.5 border border-transparent text-base font-bold rounded-xl text-slate-950 bg-cyan-400 hover:bg-cyan-300 shadow-lg shadow-cyan-500/20 focus-ring"
            >
              Go to Portal Console
            </button>
          </template>
          <template v-else>
            <NuxtLink
              to="/register"
              class="inline-flex items-center justify-center px-7 py-3.5 border border-transparent text-base font-bold rounded-xl text-slate-950 bg-cyan-400 hover:bg-cyan-300 shadow-lg shadow-cyan-500/20 focus-ring"
            >
              Find Student Work
            </NuxtLink>
            <NuxtLink
              to="/employer-register"
              class="inline-flex items-center justify-center px-7 py-3.5 border border-white/10 hover:border-cyan-300/50 text-base font-bold rounded-xl text-slate-100 bg-white/10 hover:bg-white/15 transition-colors focus-ring"
            >
              Hire Student Talent
            </NuxtLink>
          </template>
          </div>

          <div class="grid grid-cols-3 gap-3 pt-6 max-w-xl mx-auto lg:mx-0">
            <div class="rounded-xl border border-white/10 bg-white/[0.06] px-4 py-3 backdrop-blur">
              <div class="text-2xl font-extrabold text-white">11</div>
              <div class="text-xs font-semibold text-slate-400">Open roles</div>
            </div>
            <div class="rounded-xl border border-white/10 bg-white/[0.06] px-4 py-3 backdrop-blur">
              <div class="text-2xl font-extrabold text-white">4</div>
              <div class="text-xs font-semibold text-slate-400">Applications</div>
            </div>
            <div class="rounded-xl border border-white/10 bg-white/[0.06] px-4 py-3 backdrop-blur">
              <div class="text-2xl font-extrabold text-white">100%</div>
              <div class="text-xs font-semibold text-slate-400">Profile ready</div>
            </div>
          </div>
        </section>

        <section class="relative">
          <div class="absolute -inset-4 rounded-[2rem] bg-cyan-400/10 blur-3xl"></div>
          <div class="relative rounded-3xl border border-white/10 bg-slate-900/80 p-4 sm:p-5 shadow-2xl shadow-slate-950/50 backdrop-blur">
            <div class="flex items-center justify-between border-b border-white/10 pb-4">
              <div>
                <p class="text-xs font-bold uppercase tracking-wider text-cyan-200">Live workspace preview</p>
                <h2 class="mt-1 text-lg font-extrabold text-white">Student Dashboard</h2>
              </div>
              <span class="rounded-full border border-emerald-400/20 bg-emerald-400/10 px-3 py-1 text-xs font-bold text-emerald-200">Verified</span>
            </div>

            <div class="grid grid-cols-2 gap-3 py-4">
              <div class="rounded-xl border border-white/10 bg-white/[0.06] p-4">
                <p class="text-xs font-bold uppercase text-slate-400">Open jobs</p>
                <p class="mt-1 text-3xl font-extrabold text-white">11</p>
              </div>
              <div class="rounded-xl border border-white/10 bg-white/[0.06] p-4">
                <p class="text-xs font-bold uppercase text-slate-400">Accepted</p>
                <p class="mt-1 text-3xl font-extrabold text-white">1</p>
              </div>
            </div>

            <div class="space-y-3">
              <div class="rounded-2xl border border-cyan-400/20 bg-slate-950/70 p-4">
                <div class="flex items-start justify-between gap-4">
                  <div>
                    <span class="rounded-full bg-cyan-400/10 px-2.5 py-1 text-xs font-bold text-cyan-200">FPT Software</span>
                    <h3 class="mt-3 text-base font-extrabold text-white">Marketing</h3>
                    <p class="mt-2 text-sm font-medium text-slate-400">247 Cau Giay, Ha Noi</p>
                  </div>
                  <span class="rounded-lg bg-emerald-400/10 px-3 py-1 text-sm font-extrabold text-emerald-200">5.000.000 VND</span>
                </div>
                <button class="mt-4 w-full rounded-lg bg-cyan-400 px-4 py-2.5 text-sm font-bold text-slate-950">Apply Now</button>
              </div>

              <div class="grid grid-cols-2 gap-3">
                <div class="rounded-xl border border-white/10 bg-white/[0.05] p-4">
                  <p class="text-xs font-bold uppercase text-slate-500">Chat</p>
                  <p class="mt-1 text-sm font-bold text-slate-200">HR response ready</p>
                </div>
                <div class="rounded-xl border border-white/10 bg-white/[0.05] p-4">
                  <p class="text-xs font-bold uppercase text-slate-500">Attendance</p>
                  <p class="mt-1 text-sm font-bold text-slate-200">Check-in enabled</p>
                </div>
              </div>
            </div>
          </div>
        </section>
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
