<template>
  <div class="rounded-3xl border border-white/10 bg-slate-900/80 p-8 text-center shadow-xl backdrop-blur">
    <div class="mx-auto flex h-16 w-16 items-center justify-center rounded-2xl bg-cyan-400/10 text-cyan-300 ring-1 ring-cyan-400/20">
      <svg class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.8" :d="iconPath" />
      </svg>
    </div>
    <h3 class="mt-4 text-lg font-extrabold text-white">{{ title }}</h3>
    <p class="mt-2 text-sm font-medium leading-relaxed text-slate-400 max-w-sm mx-auto">{{ description }}</p>
    <div v-if="actionLabel" class="mt-6">
      <button
        @click="$emit('action')"
        class="inline-flex items-center justify-center rounded-xl bg-cyan-400 px-5 py-2.5 text-sm font-extrabold text-slate-950 shadow-md shadow-cyan-500/20 hover:bg-cyan-300 transition-colors"
      >
        {{ actionLabel }}
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
withDefaults(defineProps<{
  title: string
  description: string
  actionLabel?: string
  iconPath?: string
}>(), {
  iconPath: 'M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4'
})

defineEmits(['action'])
</script>
