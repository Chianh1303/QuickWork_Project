<template>
  <div
    v-if="totalPages > 1"
    class="flex flex-col gap-3 rounded-xl border border-white/10 bg-slate-900/82 px-4 py-3 shadow-lg shadow-slate-950/25 backdrop-blur sm:flex-row sm:items-center sm:justify-between"
  >
    <p class="text-sm font-semibold text-slate-400">
      Showing
      <span class="text-white">{{ startItem }}</span>
      -
      <span class="text-white">{{ endItem }}</span>
      of
      <span class="text-white">{{ totalItems }}</span>
    </p>

    <div class="flex items-center gap-2">
      <button
        type="button"
        :disabled="page <= 1"
        @click="emit('update:page', page - 1)"
        class="inline-flex h-9 items-center justify-center rounded-lg border border-white/10 px-3 text-sm font-bold text-slate-300 transition-colors hover:border-cyan-300/40 hover:bg-cyan-400/10 hover:text-cyan-100 disabled:cursor-not-allowed disabled:opacity-40"
      >
        Prev
      </button>

      <button
        v-for="pageNumber in visiblePages"
        :key="pageNumber"
        type="button"
        @click="emit('update:page', pageNumber)"
        :class="[
          page === pageNumber
            ? 'border-cyan-300 bg-cyan-400 text-slate-950 shadow-sm shadow-cyan-950/30'
            : 'border-white/10 text-slate-300 hover:border-cyan-300/40 hover:bg-cyan-400/10 hover:text-cyan-100',
          'inline-flex h-9 min-w-9 items-center justify-center rounded-lg border px-3 text-sm font-bold transition-colors'
        ]"
      >
        {{ pageNumber }}
      </button>

      <button
        type="button"
        :disabled="page >= totalPages"
        @click="emit('update:page', page + 1)"
        class="inline-flex h-9 items-center justify-center rounded-lg border border-white/10 px-3 text-sm font-bold text-slate-300 transition-colors hover:border-cyan-300/40 hover:bg-cyan-400/10 hover:text-cyan-100 disabled:cursor-not-allowed disabled:opacity-40"
      >
        Next
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  page: number
  pageSize: number
  totalItems: number
}>()

const emit = defineEmits<{
  'update:page': [page: number]
}>()

const totalPages = computed(() => Math.max(1, Math.ceil(props.totalItems / props.pageSize)))
const startItem = computed(() => (props.page - 1) * props.pageSize + 1)
const endItem = computed(() => Math.min(props.totalItems, props.page * props.pageSize))
const visiblePages = computed(() => {
  const start = Math.max(1, props.page - 2)
  const end = Math.min(totalPages.value, start + 4)
  return Array.from({ length: end - start + 1 }, (_, index) => start + index)
})
</script>
