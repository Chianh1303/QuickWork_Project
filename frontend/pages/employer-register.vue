<template>
  <div class="min-h-screen flex items-center justify-center bg-slate-900 py-12 px-4 sm:px-6 lg:px-8 relative overflow-hidden">
    <!-- Glow backgrounds for premium corporate employer theme -->
    <div class="absolute -top-40 -left-40 w-96 h-96 rounded-full bg-violet-600 opacity-20 blur-3xl pointer-events-none"></div>
    <div class="absolute -bottom-40 -right-40 w-96 h-96 rounded-full bg-indigo-500 opacity-20 blur-3xl pointer-events-none"></div>

    <div class="max-w-md w-full space-y-8 p-8 rounded-2xl shadow-2xl relative border border-slate-700 bg-slate-950 bg-opacity-90">
      <div>
        <div class="flex justify-center">
          <div class="h-12 w-12 rounded-xl bg-gradient-to-tr from-violet-600 to-indigo-500 flex items-center justify-center text-white font-extrabold text-2xl shadow-lg">
            Q
          </div>
        </div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-white tracking-tight">
          Employer Registration
        </h2>
        <p class="mt-2 text-center text-sm text-slate-300">
          Source top student talent and manage job postings.
        </p>
      </div>

      <!-- Success / Error Alert -->
      <div v-if="successMessage" class="bg-emerald-950 border border-emerald-500 text-emerald-100 px-4 py-3 rounded-lg text-sm flex items-center space-x-3" role="alert">
        <svg class="h-5 w-5 text-emerald-400 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <span class="font-medium">{{ successMessage }}</span>
      </div>

      <div v-if="errorMessage" class="bg-red-950 border border-red-500 text-red-100 px-4 py-3 rounded-lg text-sm flex items-center space-x-3" role="alert">
        <svg class="h-5 w-5 text-red-400 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <span class="font-medium">{{ errorMessage }}</span>
      </div>

      <form class="mt-8 space-y-5" @submit.prevent="handleRegister" v-if="!registrationSuccess">
        <div class="space-y-4">
          <!-- Full Name / HR Name -->
          <div>
            <label for="full_name" class="block text-sm font-semibold text-slate-200 mb-1">HR Manager / Recruiter Name</label>
            <input
              id="full_name"
              type="text"
              required
              v-model="form.full_name"
              class="appearance-none block w-full px-3 py-2.5 border border-slate-700 placeholder-slate-400 text-white rounded-lg bg-slate-900 hover:bg-slate-800/80 focus:outline-none focus:ring-2 focus:ring-violet-500 focus:border-transparent transition-all duration-200 sm:text-sm"
              placeholder="e.g. Ms. Jane Doe"
            />
          </div>

          <!-- Corporate Email -->
          <div>
            <label for="email" class="block text-sm font-semibold text-slate-200 mb-1">Corporate Email Address</label>
            <input
              id="email"
              type="email"
              required
              v-model="form.email"
              class="appearance-none block w-full px-3 py-2.5 border border-slate-700 placeholder-slate-400 text-white rounded-lg bg-slate-900 hover:bg-slate-800/80 focus:outline-none focus:ring-2 focus:ring-violet-500 focus:border-transparent transition-all duration-200 sm:text-sm"
              placeholder="e.g. hr@company.com"
            />
          </div>

          <!-- Password -->
          <div>
            <label for="password" class="block text-sm font-semibold text-slate-200 mb-1">Password</label>
            <input
              id="password"
              type="password"
              required
              v-model="form.password"
              class="appearance-none block w-full px-3 py-2.5 border border-slate-700 placeholder-slate-400 text-white rounded-lg bg-slate-900 hover:bg-slate-800/80 focus:outline-none focus:ring-2 focus:ring-violet-500 focus:border-transparent transition-all duration-200 sm:text-sm"
              placeholder="Minimum 6 characters"
            />
          </div>

          <!-- Company Name -->
          <div>
            <label for="company_name" class="block text-sm font-semibold text-slate-200 mb-1">Company Name</label>
            <input
              id="company_name"
              type="text"
              required
              v-model="form.company_name"
              class="appearance-none block w-full px-3 py-2.5 border border-slate-700 placeholder-slate-400 text-white rounded-lg bg-slate-900 hover:bg-slate-800/80 focus:outline-none focus:ring-2 focus:ring-violet-500 focus:border-transparent transition-all duration-200 sm:text-sm"
              placeholder="e.g. QuickWork Tech Corp"
            />
          </div>

          <!-- Tax Code -->
          <div>
            <label for="tax_code" class="block text-sm font-semibold text-slate-200 mb-1">Corporate Tax Code</label>
            <input
              id="tax_code"
              type="text"
              required
              v-model="form.tax_code"
              class="appearance-none block w-full px-3 py-2.5 border border-slate-700 placeholder-slate-400 text-white rounded-lg bg-slate-900 hover:bg-slate-800/80 focus:outline-none focus:ring-2 focus:ring-violet-500 focus:border-transparent transition-all duration-200 sm:text-sm"
              placeholder="e.g. 0101234567"
            />
          </div>
        </div>

        <div class="pt-2">
          <button
            type="submit"
            :disabled="isLoading"
            class="group relative w-full flex justify-center py-2.5 px-4 border border-transparent text-sm font-bold rounded-lg text-white bg-gradient-to-r from-violet-600 to-indigo-600 hover:from-violet-500 hover:to-indigo-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-violet-500 transition-all duration-150 disabled:opacity-50 disabled:cursor-not-allowed shadow-lg shadow-violet-500/20"
          >
            <span v-if="isLoading" class="flex items-center space-x-2">
              <svg class="animate-spin h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              <span>Creating Employer Account...</span>
            </span>
            <span v-else>Register as Employer</span>
          </button>
        </div>
      </form>

      <!-- Success Screen -->
      <div v-else class="text-center py-8 space-y-4">
        <div class="inline-flex items-center justify-center p-3 bg-emerald-950 border border-emerald-500 rounded-full text-emerald-400 mb-2">
          <svg class="h-12 w-12" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </div>
        <h3 class="text-xl font-bold text-white">Employer Account Created!</h3>
        <p class="text-slate-300 text-sm max-w-sm mx-auto font-medium">
          Corporate registrations are active. You can now log in and publish listings.
        </p>
        <div class="pt-4">
          <NuxtLink
            to="/login"
            class="inline-flex justify-center items-center py-2.5 px-6 font-bold text-sm text-white bg-gradient-to-r from-violet-600 to-indigo-500 rounded-lg hover:from-violet-500 hover:to-indigo-400 shadow-md shadow-violet-500/10 focus-ring"
          >
            Go to Login
          </NuxtLink>
        </div>
      </div>

      <!-- Navigation Links -->
      <div class="border-t border-slate-800 pt-6 flex flex-col space-y-3 text-center text-sm">
        <div class="text-slate-300">
          Already have an account?
          <NuxtLink to="/login" class="font-bold text-violet-400 hover:text-violet-300 transition-colors duration-150 ml-1">
            Sign In here
          </NuxtLink>
        </div>
        <div class="text-slate-400">
          Looking for a job?
          <NuxtLink to="/register" class="font-bold text-blue-400 hover:text-blue-300 transition-colors duration-150 ml-1">
            Register as Student
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

const { register } = useAuth()

const form = reactive({
  full_name: '',
  email: '',
  password: '',
  company_name: '',
  tax_code: ''
})

const isLoading = ref(false)
const errorMessage = ref('')
const successMessage = ref('')
const registrationSuccess = ref(false)

const handleRegister = async () => {
  isLoading.value = true
  errorMessage.value = ''
  successMessage.value = ''

  try {
    const res = await register({
      email: form.email,
      password: form.password,
      role: 'business',
      full_name: form.full_name,
      company_name: form.company_name,
      tax_code: form.tax_code
    })
    successMessage.value = res.message || '🎉 Employer registration successful!'
    registrationSuccess.value = true
  } catch (err: any) {
    errorMessage.value = err.response?._data?.message || 'Registration failed. Please check the corporate data fields.'
  } finally {
    isLoading.value = false
  }
}
</script>
