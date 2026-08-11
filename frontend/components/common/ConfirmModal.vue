<template>
  <Transition
    enter-active-class="transition duration-200 ease-out"
    enter-from-class="opacity-0 scale-95"
    enter-to-class="opacity-100 scale-100"
    leave-active-class="transition duration-150 ease-in"
    leave-from-class="opacity-100 scale-100"
    leave-to-class="opacity-0 scale-95"
  >
    <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4">
      <!-- Backdrop -->
      <div class="fixed inset-0 bg-slate-950/80 backdrop-blur-sm" @click="handleCancel"></div>

      <!-- Modal Card -->
      <div class="relative w-full max-w-md rounded-3xl border border-white/10 bg-slate-900/95 p-6 shadow-2xl backdrop-blur ring-1 ring-white/10 text-left">
        <div class="flex items-center gap-4">
          <div
            class="flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-2xl"
            :class="type === 'danger' ? 'bg-rose-500/10 text-rose-400 border border-rose-500/20' : 'bg-cyan-400/10 text-cyan-300 border border-cyan-400/20'"
          >
            <svg v-if="type === 'danger'" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
            </svg>
            <svg v-else class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div>
            <h3 class="text-lg font-extrabold text-white">{{ title }}</h3>
            <p class="mt-1 text-xs text-slate-300">{{ message }}</p>
          </div>
        </div>

        <div class="mt-6 flex justify-end space-x-3">
          <button
            type="button"
            @click="handleCancel"
            class="rounded-xl border border-white/10 bg-white/5 px-4 py-2.5 text-xs font-bold text-slate-300 hover:bg-white/10 transition-colors"
          >
            {{ cancelText }}
          </button>
          <button
            type="button"
            @click="handleConfirm"
            :disabled="isLoading"
            class="rounded-xl px-4 py-2.5 text-xs font-extrabold transition-all disabled:opacity-50"
            :class="type === 'danger' ? 'bg-rose-500 text-white hover:bg-rose-600 shadow-md shadow-rose-500/20' : 'bg-cyan-400 text-slate-950 hover:bg-cyan-300 shadow-md shadow-cyan-500/20'"
          >
            {{ isLoading ? 'Đang xử lý...' : confirmText }}
          </button>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
withDefaults(defineProps<{
  isOpen: boolean
  title: string
  message: string
  confirmText?: string
  cancelText?: string
  type?: 'info' | 'danger'
  isLoading?: boolean
}>(), {
  confirmText: 'Xác nhận',
  cancelText: 'Hủy bỏ',
  type: 'info',
  isLoading: false
})

const emit = defineEmits(['confirm', 'cancel'])

const handleConfirm = () => emit('confirm')
const handleCancel = () => emit('cancel')
</script>
