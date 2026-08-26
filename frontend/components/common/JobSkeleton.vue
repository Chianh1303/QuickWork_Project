<template>
  <div class="grid max-w-6xl grid-cols-1 lg:grid-cols-2 gap-4">
    <div
      v-for="n in count"
      :key="n"
      class="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-3.5 rounded-2xl border border-white/10 bg-slate-900/85 p-4 shadow-md backdrop-blur animate-pulse"
    >
      <div class="flex items-start sm:items-center gap-3.5 min-w-0 flex-1 w-full">
        <div class="h-14 w-14 sm:h-16 sm:w-16 rounded-xl bg-slate-800 flex-shrink-0"></div>
        <div class="min-w-0 flex-1 space-y-2">
          <div class="h-4 w-3/4 rounded bg-slate-800"></div>
          <div class="h-3 w-1/2 rounded bg-slate-800"></div>
          <div class="flex gap-2 pt-1">
            <div class="h-5 w-20 rounded bg-slate-800"></div>
            <div class="h-5 w-24 rounded bg-slate-800"></div>
          </div>
        </div>
      </div>
      <div class="w-full sm:w-auto flex justify-end">
        <div class="h-8 w-24 rounded-xl bg-slate-800"></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
withDefaults(defineProps<{ count?: number }>(), {
  count: 6
})
</script>
