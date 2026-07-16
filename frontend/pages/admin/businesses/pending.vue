<template>
  <AdminShell>
    <div class="space-y-6">
      <section class="rounded-3xl border border-white/10 bg-slate-950/80 p-6 shadow-2xl shadow-slate-950/40 ring-1 ring-cyan-400/10 lg:p-7">
        <div class="flex flex-col gap-5 lg:flex-row lg:items-center lg:justify-between">
          <div>
            <span class="inline-flex rounded-full bg-amber-400/10 px-3 py-1 text-xs font-extrabold uppercase tracking-wide text-amber-200 ring-1 ring-amber-400/25">
              Admin Review Queue
            </span>
            <h1 class="mt-3 text-3xl font-extrabold tracking-tight text-white sm:text-4xl">
              Pending Businesses
            </h1>
            <p class="mt-2 max-w-2xl text-sm font-semibold leading-6 text-slate-300 sm:text-base">
              Search and scan business registrations waiting for admin review. Approval actions are intentionally not included in this phase.
            </p>
          </div>

          <div class="grid min-w-0 grid-cols-2 gap-3 sm:min-w-[360px]">
            <div class="rounded-2xl border border-cyan-400/20 bg-cyan-400/10 px-5 py-4">
              <p class="text-xs font-extrabold uppercase tracking-wide text-cyan-200">Total Results</p>
              <p class="mt-1 text-2xl font-extrabold text-white">{{ pagination.total }}</p>
            </div>
            <div class="rounded-2xl border border-white/10 bg-white/[0.06] px-5 py-4">
              <p class="text-xs font-extrabold uppercase tracking-wide text-slate-400">Current Page</p>
              <p class="mt-1 text-2xl font-extrabold text-white">{{ pagination.page }}</p>
            </div>
          </div>
        </div>
      </section>

      <section class="rounded-2xl border border-white/10 bg-slate-950/70 p-4 shadow-xl shadow-slate-950/30 ring-1 ring-cyan-400/10 sm:p-5">
        <div class="grid gap-3 xl:grid-cols-[1fr_auto_auto] xl:items-end">
          <label class="block">
            <span class="mb-2 block text-xs font-extrabold uppercase tracking-wide text-slate-400">
              Search by Company Name, Tax Code, or Email
            </span>
            <input
              v-model="searchInput"
              type="search"
              class="w-full rounded-xl border border-white/10 bg-slate-950/80 px-4 py-3 text-sm font-semibold text-white outline-none transition focus:border-cyan-300 focus:ring-2 focus:ring-cyan-400/30"
              placeholder="Enter company name, tax code, or email"
              @keydown.enter.prevent="applySearchNow"
            >
          </label>

          <button
            type="button"
            class="inline-flex h-12 items-center justify-center rounded-xl bg-cyan-400 px-7 text-sm font-extrabold text-slate-950 transition-colors hover:bg-cyan-300 disabled:cursor-not-allowed disabled:opacity-60"
            :disabled="isLoading"
            @click="applySearchNow"
          >
            Search
          </button>

          <button
            type="button"
            class="inline-flex h-12 items-center justify-center rounded-xl border border-white/10 bg-white/5 px-5 text-sm font-extrabold text-slate-200 transition-colors hover:bg-white/10 hover:text-white disabled:cursor-not-allowed disabled:opacity-60"
            :disabled="isLoading || (!searchInput && !activeSearch)"
            @click="clearSearch"
          >
            Clear
          </button>
        </div>
      </section>

      <section class="overflow-hidden rounded-2xl border border-white/10 bg-slate-950/70 shadow-xl shadow-slate-950/30 ring-1 ring-cyan-400/10">
        <div class="flex flex-col gap-3 border-b border-white/10 px-5 py-4 sm:flex-row sm:items-center sm:justify-between">
          <div>
            <h2 class="text-base font-extrabold text-white">Business registrations</h2>
            <p class="mt-1 text-sm font-semibold text-slate-400">
              {{ activeSearch ? `Filtered by "${activeSearch}"` : 'Showing all pending businesses' }}
            </p>
          </div>
          <span class="inline-flex w-fit rounded-full bg-white/[0.06] px-3 py-1 text-xs font-extrabold uppercase tracking-wide text-slate-300 ring-1 ring-white/10">
            {{ pagination.limit }} per page
          </span>
        </div>

        <div v-if="errorMessage" class="p-6">
          <div class="rounded-2xl border border-rose-400/20 bg-rose-400/10 p-5">
            <h2 class="text-base font-extrabold text-rose-100">Unable to load pending businesses</h2>
            <p class="mt-2 text-sm font-semibold text-rose-200">{{ errorMessage }}</p>
            <button
              type="button"
              class="mt-4 rounded-xl bg-rose-300 px-4 py-2 text-sm font-extrabold text-slate-950 transition-colors hover:bg-rose-200"
              @click="fetchPendingBusinesses"
            >
              Retry
            </button>
          </div>
        </div>

        <div v-else-if="isLoading" class="space-y-3 p-5">
          <div
            v-for="index in 5"
            :key="index"
            class="h-16 animate-pulse rounded-xl border border-white/10 bg-white/[0.06]"
          />
        </div>

        <div v-else-if="items.length === 0" class="p-10 text-center">
          <div class="mx-auto flex h-14 w-14 items-center justify-center rounded-2xl bg-white/5 text-xl font-extrabold text-cyan-200 ring-1 ring-white/10">
            Q
          </div>
          <h2 class="mt-4 text-xl font-extrabold text-white">
            {{ activeSearch ? 'No pending businesses match your search.' : 'No pending businesses found.' }}
          </h2>
          <p class="mt-2 text-sm font-semibold text-slate-400">
            New business registrations with pending status will appear here.
          </p>
        </div>

        <div v-else>
          <div class="hidden overflow-x-auto lg:block">
            <table class="min-w-full divide-y divide-white/10">
              <thead class="bg-white/[0.04]">
                <tr>
                  <th class="px-5 py-3 text-left text-xs font-extrabold uppercase tracking-wide text-slate-400">Company Name</th>
                  <th class="px-5 py-3 text-left text-xs font-extrabold uppercase tracking-wide text-slate-400">Tax Code</th>
                  <th class="px-5 py-3 text-left text-xs font-extrabold uppercase tracking-wide text-slate-400">Email</th>
                  <th class="px-5 py-3 text-left text-xs font-extrabold uppercase tracking-wide text-slate-400">Submitted At</th>
                  <th class="px-5 py-3 text-left text-xs font-extrabold uppercase tracking-wide text-slate-400">Status</th>
                  <th class="px-5 py-3 text-right text-xs font-extrabold uppercase tracking-wide text-slate-400">Action</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-white/10">
                <tr
                  v-for="item in items"
                  :key="item.business_id"
                  class="cursor-pointer transition-colors hover:bg-cyan-400/10"
                  :class="selectedBusiness?.business_id === item.business_id ? 'bg-cyan-400/10' : ''"
                  @click="selectBusiness(item)"
                >
                  <td class="px-5 py-4 align-middle">
                    <p class="text-sm font-extrabold text-white">{{ item.company_name }}</p>
                    <p class="text-xs font-semibold text-slate-400">ID {{ item.business_id }}</p>
                  </td>
                  <td class="px-5 py-4 align-middle text-sm font-semibold text-slate-300">{{ item.tax_code }}</td>
                  <td class="px-5 py-4 align-middle text-sm font-semibold text-slate-300">{{ item.email }}</td>
                  <td class="px-5 py-4 align-middle text-sm font-semibold text-slate-300">{{ formatDate(item.created_at) }}</td>
                  <td class="px-5 py-4 align-middle">
                    <span class="inline-flex rounded-full bg-amber-400/10 px-3 py-1 text-xs font-extrabold uppercase tracking-wide text-amber-200 ring-1 ring-amber-400/25">
                      Pending
                    </span>
                  </td>
                  <td class="px-5 py-4 text-right align-middle">
                    <button
                      type="button"
                      title="Business detail will be implemented later"
                      class="rounded-xl border border-white/10 bg-white/5 px-4 py-2 text-sm font-extrabold text-slate-300 transition-colors hover:bg-white/10 hover:text-white"
                      @click.stop="selectBusiness(item)"
                    >
                      View
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <div class="space-y-3 p-4 lg:hidden">
            <article
              v-for="item in items"
              :key="item.business_id"
              class="rounded-2xl border border-white/10 bg-white/[0.04] p-4"
              :class="selectedBusiness?.business_id === item.business_id ? 'ring-2 ring-cyan-400/40' : ''"
              @click="selectBusiness(item)"
            >
              <div class="flex items-start justify-between gap-3">
                <div>
                  <h2 class="text-base font-extrabold text-white">{{ item.company_name }}</h2>
                  <p class="mt-1 text-sm font-semibold text-slate-400">{{ item.email }}</p>
                </div>
                <span class="rounded-full bg-amber-400/10 px-3 py-1 text-xs font-extrabold uppercase tracking-wide text-amber-200 ring-1 ring-amber-400/25">
                  Pending
                </span>
              </div>
              <dl class="mt-4 grid gap-3 text-sm">
                <div>
                  <dt class="text-xs font-extrabold uppercase tracking-wide text-slate-500">Tax Code</dt>
                  <dd class="mt-1 font-semibold text-slate-200">{{ item.tax_code }}</dd>
                </div>
                <div>
                  <dt class="text-xs font-extrabold uppercase tracking-wide text-slate-500">Submitted At</dt>
                  <dd class="mt-1 font-semibold text-slate-200">{{ formatDate(item.created_at) }}</dd>
                </div>
              </dl>
              <button
                type="button"
                title="Business detail will be implemented later"
                class="mt-4 w-full rounded-xl border border-white/10 bg-white/5 px-4 py-2 text-sm font-extrabold text-slate-300 transition-colors hover:bg-white/10 hover:text-white"
                @click.stop="selectBusiness(item)"
              >
                View
              </button>
            </article>
          </div>
        </div>

        <div
          v-if="!errorMessage && !isLoading"
          class="flex flex-col gap-3 border-t border-white/10 px-5 py-4 sm:flex-row sm:items-center sm:justify-between"
        >
          <p class="text-sm font-semibold text-slate-400">
            Page {{ pagination.page }} of {{ pagination.total_pages || 1 }} · {{ pagination.total }} result{{ pagination.total === 1 ? '' : 's' }}
          </p>
          <div class="flex gap-2">
            <button
              type="button"
              class="rounded-xl border border-white/10 bg-white/5 px-4 py-2 text-sm font-extrabold text-slate-200 transition-colors hover:bg-white/10 disabled:cursor-not-allowed disabled:opacity-40"
              :disabled="pagination.page <= 1"
              @click="goToPage(pagination.page - 1)"
            >
              Previous
            </button>
            <button
              type="button"
              class="rounded-xl bg-cyan-400 px-4 py-2 text-sm font-extrabold text-slate-950 transition-colors hover:bg-cyan-300 disabled:cursor-not-allowed disabled:opacity-40"
              :disabled="pagination.total_pages === 0 || pagination.page >= pagination.total_pages"
              @click="goToPage(pagination.page + 1)"
            >
              Next
            </button>
          </div>
        </div>
      </section>

      <section
        v-if="selectedBusiness"
        class="rounded-2xl border border-cyan-400/20 bg-cyan-400/10 p-5 shadow-xl shadow-slate-950/20"
      >
        <p class="text-xs font-extrabold uppercase tracking-wide text-cyan-200">Selected Business</p>
        <div class="mt-2 flex flex-col gap-2 md:flex-row md:items-center md:justify-between">
          <div>
            <h2 class="text-lg font-extrabold text-white">{{ selectedBusiness.company_name }}</h2>
            <p class="text-sm font-semibold text-slate-300">
              {{ selectedBusiness.email }} · Tax code {{ selectedBusiness.tax_code }}
            </p>
          </div>
          <p class="text-sm font-bold text-cyan-100">Business detail will be implemented later.</p>
        </div>
      </section>
    </div>
  </AdminShell>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import AdminShell from '~/components/admin/AdminShell.vue'
import { useAdminApi } from '~/composables/useAdminApi'
import type { PendingBusinessItem, PaginationMeta } from '~/types/admin'

definePageMeta({
  middleware: 'auth'
})

const adminApi = useAdminApi()

const items = ref<PendingBusinessItem[]>([])
const selectedBusiness = ref<PendingBusinessItem | null>(null)
const searchInput = ref('')
const activeSearch = ref('')
const isLoading = ref(false)
const errorMessage = ref('')
const pagination = ref<PaginationMeta>({
  page: 1,
  limit: 10,
  total: 0,
  total_pages: 0
})

let debounceTimer: ReturnType<typeof setTimeout> | null = null

const fetchPendingBusinesses = async () => {
  isLoading.value = true
  errorMessage.value = ''

  try {
    const response = await adminApi.getPendingBusinesses({
      page: pagination.value.page,
      limit: pagination.value.limit,
      search: activeSearch.value
    })

    items.value = response.items || []
    pagination.value = response.pagination
  } catch (error: any) {
    const status = error?.response?.status
    if (status === 403) {
      errorMessage.value = 'You do not have permission to view pending businesses.'
    } else {
      errorMessage.value = error?.response?._data?.message || 'Failed to load pending businesses.'
    }
    items.value = []
  } finally {
    isLoading.value = false
  }
}

const applySearchNow = () => {
  activeSearch.value = searchInput.value.trim()
  pagination.value.page = 1
  fetchPendingBusinesses()
}

const clearSearch = () => {
  searchInput.value = ''
  activeSearch.value = ''
  pagination.value.page = 1
  fetchPendingBusinesses()
}

const goToPage = (page: number) => {
  if (page < 1 || (pagination.value.total_pages > 0 && page > pagination.value.total_pages)) return
  pagination.value.page = page
  fetchPendingBusinesses()
}

const selectBusiness = (item: PendingBusinessItem) => {
  selectedBusiness.value = item
}

const formatDate = (value: string) => {
  if (!value) return 'N/A'
  return new Intl.DateTimeFormat('en', {
    year: 'numeric',
    month: 'short',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  }).format(new Date(value))
}

watch(searchInput, () => {
  if (debounceTimer) clearTimeout(debounceTimer)
  debounceTimer = setTimeout(() => {
    if (searchInput.value.trim() !== activeSearch.value) {
      applySearchNow()
    }
  }, 400)
})

onMounted(fetchPendingBusinesses)
</script>
