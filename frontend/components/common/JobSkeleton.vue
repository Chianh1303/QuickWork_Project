<template>
  <div class="grid max-w-6xl grid-cols-1 gap-5 md:grid-cols-2 xl:grid-cols-3">
    <div
      v-for="n in count"
      :key="n"
      class="flex min-h-[340px] flex-col rounded-lg border border-white/10 bg-slate-900/82 p-5 shadow-lg backdrop-blur animate-pulse"
    >
      <div class="flex-1 space-y-4">
        <div class="flex items-center justify-between">
          <div class="h-6 w-28 rounded-full bg-slate-800"></div>
          <div class="h-7 w-24 rounded-lg bg-slate-800"></div>
        </div>
        <div class="h-6 w-3/4 rounded-lg bg-slate-800"></div>
        <div class="h-4 w-1/2 rounded bg-slate-800"></div>
        <div class="space-y-2 pt-2">
          <div class="h-3 w-full rounded bg-slate-800"></div>
          <div class="h-3 w-5/6 rounded bg-slate-800"></div>
          <div class="h-3 w-4/6 rounded bg-slate-800"></div>
        </div>
      </div>
      <div class="mt-6 border-t border-white/10 pt-4">
        <div class="h-10 w-full rounded-lg bg-slate-800"></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
withDefaults(defineProps<{ count?: number }>(), {
  count: 6
})
</script>
