<template>
  <div class="min-h-screen flex items-center justify-center bg-slate-900 py-12 px-4 sm:px-6 lg:px-8 relative overflow-hidden">
    <!-- Glow backgrounds for premium high-contrast feel -->
    <div class="absolute -top-40 -right-40 w-96 h-96 rounded-full bg-brand-500 opacity-25 blur-3xl pointer-events-none"></div>
    <div class="absolute -bottom-40 -left-40 w-96 h-96 rounded-full bg-accent-500 opacity-25 blur-3xl pointer-events-none"></div>

    <div class="max-w-md w-full space-y-8 p-8 rounded-2xl shadow-2xl relative border border-slate-700 bg-slate-950 bg-opacity-95">
      <div>
        <div class="flex justify-center">
          <div class="h-12 w-12 rounded-xl bg-gradient-to-tr from-brand-600 to-accent-500 flex items-center justify-center text-white font-extrabold text-2xl shadow-lg">
            Q
          </div>
        </div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-white tracking-tight">
          Welcome to QuickWork
        </h2>
        <p class="mt-2 text-center text-sm text-slate-300">
          Access your workspace and listings.
        </p>
      </div>

      <!-- Alert for error messages -->
      <div v-if="errorMessage" class="bg-red-950 border border-red-500 text-red-100 px-4 py-3 rounded-lg text-sm flex items-center space-x-3" role="alert">
        <svg class="h-5 w-5 text-red-400 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <span class="font-medium">{{ errorMessage }}</span>
      </div>

      <form class="mt-8 space-y-6" @submit.prevent="handleLogin">
        <div class="space-y-4">
          <!-- Email Address -->
          <div>
            <label for="email-address" class="block text-sm font-semibold text-slate-200 mb-1">Email Address</label>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <svg class="h-5 w-5 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                </svg>
              </div>
              <input
                id="email-address"
                name="email"
                type="email"
                autocomplete="email"
                required
                v-model="form.email"
                class="appearance-none relative block w-full pl-10 pr-3 py-2.5 border border-slate-700 placeholder-slate-400 text-white rounded-lg bg-slate-900 hover:bg-slate-800/80 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-transparent transition-all duration-200 sm:text-sm"
                placeholder="name@example.com"
              />
            </div>
          </div>

          <!-- Password -->
          <div>
            <label for="password" class="block text-sm font-semibold text-slate-200 mb-1">Password</label>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <svg class="h-5 w-5 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 7a2 2 0 012 2m-2 2a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2V9a2 2 0 00-2-2h-2zm0 0V5.5A2.5 2.5 0 0012.5 3h-5A2.5 2.5 0 005 5.5V7m10 4V9" />
                </svg>
              </div>
              <input
                id="password"
                name="password"
                type="password"
                autocomplete="current-password"
                required
                v-model="form.password"
                class="appearance-none relative block w-full pl-10 pr-3 py-2.5 border border-slate-700 placeholder-slate-400 text-white rounded-lg bg-slate-900 hover:bg-slate-800/80 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-transparent transition-all duration-200 sm:text-sm"
                placeholder="••••••••"
              />
            </div>
          </div>
        </div>

        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <input
              id="remember-me"
              name="remember-me"
              type="checkbox"
              class="h-4 w-4 text-brand-500 focus:ring-brand-500 border-slate-700 rounded bg-slate-900 cursor-pointer"
            />
            <label for="remember-me" class="ml-2 block text-sm text-slate-300 font-medium cursor-pointer">
              Remember me
            </label>
          </div>

          <div class="text-sm">
            <a href="#" class="font-bold text-brand-400 hover:text-brand-300 transition-colors">
              Forgot password?
            </a>
          </div>
        </div>

        <div>
          <button
            type="submit"
            :disabled="isLoading"
            class="group relative w-full flex justify-center py-2.5 px-4 border border-transparent text-sm font-bold rounded-lg text-white bg-gradient-to-r from-brand-600 to-brand-500 hover:from-brand-500 hover:to-brand-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500 transition-all duration-150 disabled:opacity-50 disabled:cursor-not-allowed shadow-lg shadow-brand-500/20"
          >
            <span v-if="isLoading" class="flex items-center space-x-2">
              <svg class="animate-spin h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              <span>Signing in...</span>
            </span>
            <span v-else>Sign In</span>
          </button>
        </div>
      </form>

      <!-- Navigation Links -->
      <div class="border-t border-slate-800 pt-6 flex flex-col space-y-3 text-center text-sm">
        <div class="text-slate-300">
          Need a student account?
          <NuxtLink to="/register" class="font-bold text-blue-400 hover:text-blue-300 transition-colors duration-150 ml-1">
            Create Student Account
          </NuxtLink>
        </div>
        <div class="text-slate-400">
          Want to hire?
          <NuxtLink to="/employer-register" class="font-bold text-emerald-400 hover:text-emerald-300 transition-colors duration-150 ml-1">
            Register as Employer
          </NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useAuth } from '~/composables/useAuth'

definePageMeta({
  middleware: 'auth'
})

const { login } = useAuth()

const form = reactive({
  email: '',
  password: ''
})

const isLoading = ref(false)
const errorMessage = ref('')

const handleLogin = async () => {
  isLoading.value = true
  errorMessage.value = ''

  try {
    await login({
      email: form.email,
      password: form.password
    })
  } catch (err: any) {
    errorMessage.value = err.response?._data?.message || 'Failed to sign in. Please check your credentials.'
  } finally {
    isLoading.value = false
  }
}
</script>
