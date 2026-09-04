<template>
  <div class="relative" ref="bellRef">
    <button
      @click="toggleDropdown"
      class="relative p-2.5 rounded-xl border border-white/10 bg-slate-900/90 hover:bg-slate-800 text-cyan-300 hover:text-white transition-all cursor-pointer shadow-md flex items-center justify-center"
      title="Thông báo hệ thống"
    >
      <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
      </svg>
      
      <!-- Unread Count Badge -->
      <span
        v-if="unreadCount > 0"
        class="absolute -top-1 -right-1 flex h-5 w-5 items-center justify-center rounded-full bg-rose-500 text-[10px] font-black text-white ring-2 ring-slate-950 animate-bounce shadow-lg shadow-rose-500/50"
      >
        {{ unreadCount > 9 ? '9+' : unreadCount }}
      </span>
    </button>

    <!-- Dropdown Window -->
    <div
      v-if="isOpen"
      :class="[
        align === 'left' ? 'left-0' : 'right-0',
        'absolute mt-2.5 w-[calc(100vw-2.5rem)] max-w-sm sm:w-96 rounded-2xl border border-white/20 bg-slate-950/95 shadow-2xl shadow-slate-950/95 backdrop-blur-2xl z-[100] overflow-hidden text-left'
      ]"
    >
      <div class="p-3.5 border-b border-white/10 bg-slate-900/90 flex items-center justify-between">
        <div class="flex items-center gap-2">
          <span class="text-xs font-black uppercase tracking-wider text-cyan-300">🔔 Thông báo hệ thống</span>
          <span v-if="unreadCount > 0" class="px-2 py-0.5 rounded-full text-[10px] font-bold bg-rose-500/20 text-rose-300 border border-rose-500/30">
            {{ unreadCount }} mới
          </span>
        </div>
        <button
          v-if="unreadCount > 0"
          @click="markAllAsRead"
          class="text-[11px] font-bold text-cyan-300 hover:text-cyan-200 transition-colors cursor-pointer"
        >
          Đánh dấu tất cả đã đọc
        </button>
      </div>

      <div class="max-h-80 overflow-y-auto divide-y divide-white/5">
        <div
          v-for="notif in notifications"
          :key="notif.id"
          @click="handleNotifClick(notif)"
          :class="[
            !notif.is_read ? 'bg-cyan-400/10' : 'bg-transparent',
            'p-3.5 hover:bg-white/5 transition-colors cursor-pointer flex items-start gap-3'
          ]"
        >
          <span class="text-lg flex-shrink-0">
            {{ getNotifIcon(notif.type) }}
          </span>
          <div class="min-w-0 flex-1">
            <div class="flex items-center justify-between gap-1">
              <h4 class="text-xs font-bold text-white truncate">{{ notif.title }}</h4>
              <span v-if="!notif.is_read" class="h-2 w-2 rounded-full bg-cyan-400 flex-shrink-0"></span>
            </div>
            <p class="mt-1 text-xs text-slate-300 font-medium line-clamp-2 leading-relaxed">{{ notif.message }}</p>
            <span class="mt-1.5 block text-[10px] text-slate-400 font-semibold">{{ formatTime(notif.created_at) }}</span>
          </div>
        </div>

        <div v-if="notifications.length === 0" class="p-8 text-center text-xs text-slate-400 font-medium">
          Chưa có thông báo nào.
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { useApi } from '~/composables/useApi'
import { useAuth } from '~/composables/useAuth'

const props = withDefaults(defineProps<{
  align?: 'left' | 'right'
}>(), {
  align: 'right'
})

const api = useApi()
const { isAuthenticated, userRole } = useAuth()

const bellRef = ref<HTMLElement | null>(null)
const isOpen = ref(false)
const notifications = ref<any[]>([])
const unreadCount = ref(0)
let timer: any = null

const handleClickOutside = (event: MouseEvent) => {
  if (isOpen.value && bellRef.value && !bellRef.value.contains(event.target as Node)) {
    isOpen.value = false
  }
}

const toggleDropdown = async () => {
  isOpen.value = !isOpen.value
  if (isOpen.value && unreadCount.value > 0) {
    await markAllAsRead()
  }
}

let isFetching = false
const fetchNotifs = async () => {
  if (!isAuthenticated.value || isFetching) return
  isFetching = true
  try {
    const res = await api.get('/api/notifications?limit=20')
    notifications.value = res.data || []
    unreadCount.value = res.unread_count || 0
  } catch (e) {
  } finally {
    isFetching = false
  }
}

const markAllAsRead = async () => {
  try {
    await api.patch('/api/notifications/read-all')
    unreadCount.value = 0
    notifications.value.forEach(n => n.is_read = true)
  } catch (e) {}
}

const handleNotifClick = async (notif: any) => {
  if (!notif.is_read) {
    try {
      await api.patch(`/api/notifications/${notif.id}/read`)
      notif.is_read = true
      if (unreadCount.value > 0) unreadCount.value--
    } catch (e) {}
  }

  isOpen.value = false

  const targetAppId = useState<number | null>('targetAppIdFromNotification', () => null)
  if (notif.reference_id) {
    targetAppId.value = Number(notif.reference_id)
  }

  if (userRole.value === 'business') {
    const activeSection = useState<string>('businessDashboardActiveSection', () => 'dashboard')
    activeSection.value = 'applicants'
    await navigateTo('/business/dashboard')
  } else if (userRole.value === 'student') {
    const activeSection = useState<string>('studentDashboardActiveSection', () => 'jobs')
    activeSection.value = 'applications'
    await navigateTo('/student/dashboard')
  }
}

const getNotifIcon = (type: string) => {
  switch (type) {
    case 'chat': return '💬'
    case 'offer': return '🎉'
    case 'application': return '📄'
    case 'escrow': return '💰'
    default: return '🔔'
  }
}

const formatTime = (timeStr: string) => {
  if (!timeStr) return ''
  const d = new Date(timeStr)
  return d.toLocaleTimeString('vi-VN', { hour: '2-digit', minute: '2-digit' }) + ' ' + d.toLocaleDateString('vi-VN')
}

onMounted(() => {
  fetchNotifs()
  timer = setInterval(fetchNotifs, 8000)
  if (typeof window !== 'undefined') {
    window.addEventListener('click', handleClickOutside)
  }
})

onBeforeUnmount(() => {
  if (timer) clearInterval(timer)
  if (typeof window !== 'undefined') {
    window.removeEventListener('click', handleClickOutside)
  }
})
</script>
