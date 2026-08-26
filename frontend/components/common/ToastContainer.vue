<template>
  <div class="fixed top-5 right-5 z-50 flex flex-col space-y-3 max-w-sm w-full pointer-events-none px-4 sm:px-0">
    <TransitionGroup
      enter-active-class="transform ease-out duration-300 transition"
      enter-from-class="translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-4"
      enter-to-class="translate-y-0 opacity-100 sm:translate-x-0"
      leave-active-class="transition ease-in duration-200"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div
        v-for="toast in toasts"
        :key="toast.id"
        class="pointer-events-auto flex items-start gap-3 rounded-2xl border p-4 shadow-2xl backdrop-blur-md transition-all"
        :class="getToastStyles(toast.type)"
      >
        <div class="flex-shrink-0 mt-0.5">
          <!-- Success Icon -->
          <svg v-if="toast.type === 'success'" class="h-5 w-5 text-emerald-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <!-- Error Icon -->
          <svg v-else-if="toast.type === 'error'" class="h-5 w-5 text-rose-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <!-- Warning Icon -->
          <svg v-else-if="toast.type === 'warning'" class="h-5 w-5 text-amber-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
          </svg>
          <!-- Info Icon -->
          <svg v-else class="h-5 w-5 text-cyan-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </div>

        <div class="flex-1 text-sm font-medium">
          <p v-if="toast.title" class="font-bold text-white leading-snug">{{ toast.title }}</p>
          <p class="mt-0.5 text-slate-300 text-xs leading-relaxed">{{ toast.message }}</p>
        </div>

        <button @click="remove(toast.id)" class="text-slate-400 hover:text-white transition-colors">
          ✕
        </button>
      </div>
    </TransitionGroup>
  </div>
</template>

<script setup lang="ts">
import { useToast } from '~/composables/useToast'

const { toasts, remove } = useToast()

const getToastStyles = (type: string) => {
  switch (type) {
    case 'success':
      return 'bg-slate-950/90 border-emerald-400/30 text-emerald-100 ring-1 ring-emerald-400/20'
    case 'error':
      return 'bg-slate-950/90 border-rose-400/30 text-rose-100 ring-1 ring-rose-400/20'
    case 'warning':
      return 'bg-slate-950/90 border-amber-400/30 text-amber-100 ring-1 ring-amber-400/20'
    default:
      return 'bg-slate-950/90 border-cyan-400/30 text-cyan-100 ring-1 ring-cyan-400/20'
  }
}
</script>
