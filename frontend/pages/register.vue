<template>
  <div class="min-h-screen flex items-center justify-center bg-slate-900 py-12 px-4 sm:px-6 lg:px-8 relative overflow-hidden">
    <!-- Glow backgrounds for career-focused student theme -->
    <div class="absolute -top-40 -left-40 w-96 h-96 rounded-full bg-blue-500 opacity-20 blur-3xl pointer-events-none"></div>
    <div class="absolute -bottom-40 -right-40 w-96 h-96 rounded-full bg-emerald-500 opacity-20 blur-3xl pointer-events-none"></div>

    <div class="max-w-md w-full space-y-8 p-8 rounded-2xl shadow-2xl relative border border-slate-700 bg-slate-950 bg-opacity-90">
      <div>
        <div class="flex justify-center">
          <div class="h-12 w-12 rounded-xl bg-gradient-to-tr from-blue-600 to-emerald-500 flex items-center justify-center text-white font-extrabold text-2xl shadow-lg">
            Q
          </div>
        </div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-white tracking-tight">
          Create Student Account
        </h2>
        <p class="mt-2 text-center text-sm text-slate-300">
          Find student jobs and launch your career.
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

      <form class="mt-8 space-y-6" @submit.prevent="handleRegister" v-if="!registrationSuccess">
        <div class="space-y-4">
          <!-- Full Name -->
          <div>
            <label for="full_name" class="block text-sm font-semibold text-slate-200 mb-1">Full Name</label>
            <input
              id="full_name"
              type="text"
              required
              v-model="form.full_name"
              class="appearance-none block w-full px-3 py-2.5 border border-slate-700 placeholder-slate-400 text-white rounded-lg bg-slate-900 hover:bg-slate-800/80 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 sm:text-sm"
              placeholder="e.g. Nguyen Van A"
            />
          </div>

          <!-- Email -->
          <div>
            <label for="email" class="block text-sm font-semibold text-slate-200 mb-1">Email Address</label>
            <input
              id="email"
              type="email"
              required
              v-model="form.email"
              class="appearance-none block w-full px-3 py-2.5 border border-slate-700 placeholder-slate-400 text-white rounded-lg bg-slate-900 hover:bg-slate-800/80 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 sm:text-sm"
              placeholder="e.g. student@university.edu"
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
              class="appearance-none block w-full px-3 py-2.5 border border-slate-700 placeholder-slate-400 text-white rounded-lg bg-slate-900 hover:bg-slate-800/80 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 sm:text-sm"
              placeholder="Minimum 6 characters"
            />
          </div>
        </div>

        <div>
          <button
            type="submit"
            :disabled="isLoading"
            class="group relative w-full flex justify-center py-2.5 px-4 border border-transparent text-sm font-bold rounded-lg text-white bg-gradient-to-r from-blue-600 to-blue-500 hover:from-blue-500 hover:to-blue-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-all duration-150 disabled:opacity-50 disabled:cursor-not-allowed shadow-lg shadow-blue-500/20"
          >
            <span v-if="isLoading" class="flex items-center space-x-2">
              <svg class="animate-spin h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              <span>Creating Student Account...</span>
            </span>
            <span v-else>Register as Student</span>
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
        <h3 class="text-xl font-bold text-white">Student Account Created!</h3>
        <p class="text-slate-300 text-sm max-w-sm mx-auto font-medium">
          Your credentials have been successfully saved. You can now log in.
        </p>
        <div class="pt-4">
          <NuxtLink
            to="/login"
            class="inline-flex justify-center items-center py-2.5 px-6 font-bold text-sm text-white bg-gradient-to-r from-blue-600 to-blue-500 rounded-lg hover:from-blue-500 hover:to-blue-400 shadow-md shadow-blue-500/10 focus-ring"
          >
            Go to Login
          </NuxtLink>
        </div>
      </div>

      <!-- Navigation Links -->
      <div class="border-t border-slate-800 pt-6 flex flex-col space-y-3 text-center text-sm">
        <div class="text-slate-300">
          Already have an account?
          <NuxtLink to="/login" class="font-bold text-blue-400 hover:text-blue-300 transition-colors duration-150 ml-1">
            Sign In here
          </NuxtLink>
        </div>
        <div class="text-slate-400">
          Are you hiring?
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

const { register } = useAuth()

const form = reactive({
  full_name: '',
  email: '',
  password: ''
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
      role: 'student',
      full_name: form.full_name
    })
    successMessage.value = res.message || '🎉 Registration successful! Welcome to QuickWork.'
    registrationSuccess.value = true
  } catch (err: any) {
    errorMessage.value = err.response?._data?.message || 'Registration failed. Please check the form fields.'
  } finally {
    isLoading.value = false
  }
}
</script>
