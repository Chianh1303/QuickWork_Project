<template>
  <AdminShell>
    <div class="space-y-6">
      <section class="overflow-hidden rounded-3xl border border-white/10 bg-slate-950/80 shadow-2xl shadow-slate-950/40 ring-1 ring-cyan-400/10">
        <div class="grid gap-6 p-6 lg:grid-cols-[1.1fr_0.9fr] lg:p-8">
          <div class="flex flex-col justify-between gap-8">
            <div>
              <span class="inline-flex rounded-full bg-cyan-400/10 px-3 py-1 text-xs font-extrabold uppercase tracking-wide text-cyan-200 ring-1 ring-cyan-400/25">
                Admin Dashboard
              </span>
              <h1 class="mt-4 max-w-3xl text-3xl font-extrabold tracking-tight text-white sm:text-4xl">
                QuickWork operations command center
              </h1>
              <p class="mt-3 max-w-2xl text-sm font-semibold leading-6 text-slate-300 sm:text-base">
                Monitor business onboarding, keep marketplace access controlled, and prepare pending registrations for review.
              </p>
            </div>

            <div class="flex flex-col gap-3 sm:flex-row">
              <NuxtLink
                to="/admin/businesses/pending"
                class="inline-flex items-center justify-center rounded-xl bg-cyan-400 px-5 py-3 text-sm font-extrabold text-slate-950 shadow-lg shadow-cyan-500/20 transition-colors hover:bg-cyan-300"
              >
                Review Pending Business
              </NuxtLink>
              <div class="inline-flex items-center justify-center rounded-xl border border-white/10 bg-white/[0.06] px-5 py-3 text-sm font-bold text-slate-200">
                Admin-only workspace
              </div>
            </div>
          </div>

          <div class="grid gap-3 sm:grid-cols-2">
            <div
              v-for="metric in metrics"
              :key="metric.label"
              class="rounded-2xl border border-white/10 bg-white/[0.06] p-5"
            >
              <p class="text-xs font-extrabold uppercase tracking-wide text-slate-400">{{ metric.label }}</p>
              <p class="mt-3 text-3xl font-extrabold text-white">{{ metric.value }}</p>
              <p class="mt-2 text-sm font-semibold text-slate-300">{{ metric.caption }}</p>
            </div>
          </div>
        </div>
      </section>

      <section class="grid gap-4 lg:grid-cols-3">
        <article
          v-for="card in cards"
          :key="card.title"
          class="rounded-2xl border border-white/10 bg-slate-950/70 p-5 shadow-xl shadow-slate-950/25 ring-1 ring-cyan-400/10"
        >
          <div class="flex items-start justify-between gap-4">
            <div>
              <p class="text-xs font-extrabold uppercase tracking-wide text-slate-400">{{ card.label }}</p>
              <h2 class="mt-2 text-xl font-extrabold text-white">{{ card.title }}</h2>
            </div>
            <span class="flex h-10 w-10 items-center justify-center rounded-xl bg-cyan-400/10 text-sm font-extrabold text-cyan-200 ring-1 ring-cyan-400/20">
              {{ card.index }}
            </span>
          </div>
          <p class="mt-4 text-sm font-semibold leading-6 text-slate-300">{{ card.description }}</p>
        </article>
      </section>

      <section class="rounded-2xl border border-white/10 bg-slate-950/70 p-5 shadow-xl shadow-slate-950/25 ring-1 ring-cyan-400/10">
        <div class="flex flex-col gap-4 md:flex-row md:items-center md:justify-between">
          <div>
            <p class="text-xs font-extrabold uppercase tracking-wide text-cyan-200">Primary workflow</p>
            <h2 class="mt-2 text-xl font-extrabold text-white">Business onboarding queue</h2>
            <p class="mt-2 text-sm font-semibold leading-6 text-slate-300">
              Use the Pending Business page to search, scan, and select business registrations. Detail review and approval actions remain reserved for the next phase.
            </p>
          </div>
          <NuxtLink
            to="/admin/businesses/pending"
            class="inline-flex items-center justify-center rounded-xl border border-cyan-400/25 bg-cyan-400/10 px-5 py-3 text-sm font-extrabold text-cyan-100 transition-colors hover:bg-cyan-400 hover:text-slate-950"
          >
            Go to Queue
          </NuxtLink>
        </div>
      </section>
    </div>
  </AdminShell>
</template>

<script setup lang="ts">
import AdminShell from '~/components/admin/AdminShell.vue'

definePageMeta({
  middleware: 'auth'
})

const metrics = [
  { label: 'Access', value: 'Admin', caption: 'Role-protected controls' },
  { label: 'Queue', value: 'Pending', caption: 'Businesses awaiting review' },
  { label: 'Scope', value: 'List', caption: 'No approve or reject in this phase' },
  { label: 'API', value: 'JWT', caption: 'Authenticated REST endpoint' }
]

const cards = [
  {
    index: '01',
    label: 'Queue',
    title: 'Pending Business',
    description: 'View businesses waiting for admin review and prepare records for later approval workflow.'
  },
  {
    index: '02',
    label: 'Security',
    title: 'Admin-only access',
    description: 'All admin APIs are protected by JWT authentication and role-based authorization.'
  },
  {
    index: '03',
    label: 'Scope',
    title: 'Review list only',
    description: 'Approve, reject, and detailed profile review actions are intentionally left for a later task.'
  }
]
</script>
