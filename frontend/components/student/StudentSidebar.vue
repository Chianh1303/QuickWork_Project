<template>
  <!-- DESKTOP SIDEBAR -->
  <aside class="hidden lg:flex lg:w-64 lg:flex-shrink-0 lg:flex-col lg:border-r lg:border-white/10 lg:bg-slate-950 lg:px-4 lg:py-6 sticky top-0 h-screen overflow-y-auto z-30">
    <button @click="activeSection = 'jobs'" class="flex items-center gap-2.5 px-2 pb-6 mb-4 border-b border-white/10 text-left w-full group">
      <span class="flex h-9 w-9 flex-shrink-0 items-center justify-center rounded-xl bg-brand-400 text-sm font-bold text-slate-950">QW</span>
      <div class="min-w-0">
        <p class="text-sm font-bold leading-tight text-white">QuickWork</p>
        <p class="text-xs text-slate-500">Student Workspace</p>
      </div>
    </button>

    <nav class="flex flex-1 flex-col gap-5">
      <div v-for="group in navGroups" :key="group.label">
        <p class="px-3 mb-1.5 text-[10px] font-semibold uppercase tracking-wider text-slate-500">{{ group.label }}</p>
        <div class="space-y-0.5">
          <button
            v-for="item in group.items"
            :key="item.id"
            @click="activeSection = item.id"
            :class="[
              activeSection === item.id
                ? 'border-brand-400 bg-brand-400/10 text-brand-200 font-semibold'
                : 'border-transparent text-slate-400 hover:bg-white/5 hover:text-white font-medium',
              'flex w-full items-center gap-3 rounded-r-lg border-l-2 py-2.5 pl-3 pr-2.5 text-sm transition-colors'
            ]"
          >
            <NavIcon :icon="item.id" class="h-4 w-4 flex-shrink-0" />
            <span class="flex-1 truncate text-left">{{ item.name }}</span>
            <span v-if="badgeCount(item.id) > 0" class="rounded-full bg-white/10 px-2 py-0.5 text-[10px] font-semibold text-slate-300">
              {{ badgeCount(item.id) }}
            </span>
          </button>
        </div>
      </div>
    </nav>

    <UserArea :display-name="displayName" :avatar-url="avatarUrl" :initials="initials" :profile-readiness="profileReadinessValue" @logout="$emit('logout')" />
  </aside>

  <!-- MOBILE TOP BAR -->
  <div class="flex items-center justify-between border-b border-white/10 bg-slate-950/95 px-4 py-3 backdrop-blur lg:hidden">
    <button
      type="button"
      @click="isDrawerOpen = true"
      aria-label="Mở menu điều hướng"
      class="flex h-9 w-9 items-center justify-center rounded-lg border border-white/10 text-slate-300 hover:bg-white/5 focus:outline-none focus-visible:ring-2 focus-visible:ring-brand-400/50"
    >
      <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
      </svg>
    </button>

    <span class="flex items-center gap-2 text-sm font-bold text-white">
      <span class="flex h-6 w-6 items-center justify-center rounded-md bg-brand-400 text-[10px] font-bold text-slate-950">QW</span>
      QuickWork
    </span>

    <div class="flex items-center gap-2">
      <NotificationBell />
      <button
        type="button"
        @click="activeSection = 'profile'"
        :aria-label="`Hồ sơ của ${displayName}`"
        class="flex h-9 w-9 items-center justify-center rounded-full bg-white/10 text-xs font-bold text-white focus:outline-none focus-visible:ring-2 focus-visible:ring-brand-400/50"
      >
        {{ initials }}
      </button>
    </div>
  </div>

  <!-- MOBILE DRAWER -->
  <Transition name="drawer-fade">
    <div v-if="isDrawerOpen" class="fixed inset-0 z-40 bg-slate-950/70 lg:hidden" @click="isDrawerOpen = false"></div>
  </Transition>
  <Transition name="drawer-slide">
    <aside
      v-if="isDrawerOpen"
      class="fixed inset-y-0 left-0 z-50 flex w-72 max-w-[85vw] flex-col border-r border-white/10 bg-slate-950 px-4 py-6 lg:hidden"
    >
      <div class="flex items-center justify-between pb-6 mb-4 border-b border-white/10">
        <span class="flex items-center gap-2.5">
          <span class="flex h-9 w-9 items-center justify-center rounded-xl bg-brand-400 text-sm font-bold text-slate-950">QW</span>
          <span>
            <p class="text-sm font-bold leading-tight text-white">QuickWork</p>
            <p class="text-xs text-slate-500">Student Workspace</p>
          </span>
        </span>
        <button
          type="button"
          @click="isDrawerOpen = false"
          aria-label="Đóng menu"
          class="flex h-8 w-8 items-center justify-center rounded-lg text-slate-400 hover:bg-white/5 hover:text-white focus:outline-none focus-visible:ring-2 focus-visible:ring-brand-400/50"
        >
          <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <nav class="flex flex-1 flex-col gap-5 overflow-y-auto">
        <div v-for="group in navGroups" :key="group.label">
          <p class="px-3 mb-1.5 text-[10px] font-semibold uppercase tracking-wider text-slate-500">{{ group.label }}</p>
          <div class="space-y-0.5">
            <button
              v-for="item in group.items"
              :key="item.id"
              @click="selectMobileSection(item.id)"
              :class="[
                activeSection === item.id
                  ? 'border-brand-400 bg-brand-400/10 text-brand-200 font-semibold'
                  : 'border-transparent text-slate-400 hover:bg-white/5 hover:text-white font-medium',
                'flex w-full items-center gap-3 rounded-r-lg border-l-2 py-2.5 pl-3 pr-2.5 text-sm transition-colors'
              ]"
            >
              <NavIcon :icon="item.id" class="h-4 w-4 flex-shrink-0" />
              <span class="flex-1 truncate text-left">{{ item.name }}</span>
              <span v-if="badgeCount(item.id) > 0" class="rounded-full bg-white/10 px-2 py-0.5 text-[10px] font-semibold text-slate-300">
                {{ badgeCount(item.id) }}
              </span>
            </button>
          </div>
        </div>
      </nav>

      <UserArea :display-name="displayName" :avatar-url="avatarUrl" :initials="initials" :profile-readiness="profileReadinessValue" @logout="$emit('logout')" />
    </aside>
  </Transition>
</template>

<script setup lang="ts">
import { toRefs, computed, h, defineComponent, ref } from 'vue'
import NotificationBell from '~/components/common/NotificationBell.vue'
import { useMedia } from '~/composables/useMedia'

const { getMediaUrl } = useMedia()
const props = defineProps<{ state: Record<string, any> }>()
defineEmits<{ logout: [] }>()

const { activeSection, navItems, filteredJobs, filteredApps, profileForm, profileReadiness } = toRefs(props.state)

const isDrawerOpen = ref(false)
const selectMobileSection = (id: string) => {
  activeSection.value = id
  isDrawerOpen.value = false
}

const navGroups = computed(() => {
  const items = navItems?.value || []
  const byId = (id: string) => items.find((i: any) => i.id === id)
  const overview = [byId('jobs'), byId('saved-jobs'), byId('applications')].filter(Boolean)
  const profile = [byId('profile'), byId('wallet')].filter(Boolean)
  const known = new Set(['jobs', 'saved-jobs', 'applications', 'profile', 'wallet'])
  const rest = items.filter((i: any) => !known.has(i.id))
  const groups = [
    { label: 'Overview', items: overview },
    { label: 'Profile', items: profile }
  ]
  if (rest.length) groups.push({ label: 'Khác', items: rest })
  return groups.filter(g => g.items.length > 0)
})

const badgeCount = (id: string): number => {
  if (id === 'applications') return filteredApps?.value?.length || 0
  if (id === 'jobs') return filteredJobs?.value?.length || 0
  return 0
}

const displayName = computed(() => profileForm?.value?.full_name || 'Hồ sơ Sinh viên')
const avatarUrl = computed(() => profileForm?.value?.avatar_url || '')
const initials = computed(() => {
  const name = (profileForm?.value?.full_name || '').trim()
  if (!name) return 'SV'
  const parts = name.split(/\s+/)
  const last = parts[parts.length - 1]?.[0] || ''
  const first = parts[0]?.[0] || ''
  return (first + last).toUpperCase() || 'SV'
})
const profileReadinessValue = computed(() => profileReadiness?.value ?? 0)

const iconPaths: Record<string, string> = {
  jobs: 'M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z',
  'saved-jobs': 'M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z',
  profile: 'M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z',
  applications: 'M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z',
  wallet: 'M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-2m0-6h4v6h-4a3 3 0 010-6z'
}
const NavIcon = defineComponent({
  props: { icon: { type: String, required: true } },
  setup(p) {
    return () => h('svg', { class: 'flex-shrink-0', fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'aria-hidden': 'true' }, [
      h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', 'stroke-width': '2', d: iconPaths[p.icon] || iconPaths.jobs })
    ])
  }
})

const UserArea = defineComponent({
  props: {
    displayName: { type: String, required: true },
    avatarUrl: { type: String, default: '' },
    initials: { type: String, required: true },
    profileReadiness: { type: Number, required: true }
  },
  emits: ['logout'],
  setup(p, { emit }) {
    return () => h('div', { class: 'mt-auto space-y-3 pt-4 border-t border-white/10' }, [
      h('div', { class: 'flex items-center gap-3 px-1' }, [
        p.avatarUrl
          ? h('img', {
              src: getMediaUrl(p.avatarUrl),
              alt: p.displayName,
              class: 'h-9 w-9 flex-shrink-0 rounded-full object-cover border border-brand-400/40 shadow-sm'
            })
          : h('span', { class: 'flex h-9 w-9 flex-shrink-0 items-center justify-center rounded-full bg-white/10 text-xs font-bold text-white' }, p.initials),
        h('div', { class: 'min-w-0 flex-1' }, [
          h('p', { class: 'truncate text-sm font-semibold text-white' }, p.displayName),
          h('p', { class: 'text-xs text-slate-500' }, `${p.profileReadiness}% hồ sơ hoàn thiện`)
        ])
      ]),
      h('div', { class: 'h-1.5 w-full overflow-hidden rounded-full bg-white/10' }, [
        h('div', { class: 'h-full rounded-full bg-brand-400 transition-all duration-500', style: { width: `${p.profileReadiness}%` } })
      ]),
      h('button', {
        type: 'button',
        onClick: () => emit('logout'),
        class: 'flex w-full items-center justify-center gap-2 rounded-xl border border-white/10 px-4 py-2.5 text-xs font-semibold text-slate-300 transition-colors hover:bg-white/5 hover:text-white'
      }, [
        h('svg', { class: 'h-4 w-4', fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'aria-hidden': 'true' }, [
          h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', 'stroke-width': '2', d: 'M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1' })
        ]),
        h('span', 'Đăng xuất')
      ])
    ])
  }
})
</script>

<style scoped>
.drawer-fade-enter-active,
.drawer-fade-leave-active {
  transition: opacity 0.25s ease;
}
.drawer-fade-enter-from,
.drawer-fade-leave-to {
  opacity: 0;
}

.drawer-slide-enter-active,
.drawer-slide-leave-active {
  transition: transform 0.25s cubic-bezier(0.16, 1, 0.3, 1);
}
.drawer-slide-enter-from,
.drawer-slide-leave-to {
  transform: translateX(-100%);
}

@media (prefers-reduced-motion: reduce) {
  .drawer-fade-enter-active,
  .drawer-fade-leave-active,
  .drawer-slide-enter-active,
  .drawer-slide-leave-active {
    transition: none !important;
  }
}
</style>
