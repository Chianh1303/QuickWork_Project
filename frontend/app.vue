<template>
  <div class="app-shell min-h-screen flex flex-col text-slate-900 font-sans">
    <!-- Premium Navigation Bar -->
    <header v-if="!isStudentDashboard && !isBusinessDashboard && !isAdminRoute" class="bg-slate-950/95 backdrop-blur-md border-b border-cyan-400/10 sticky top-0 z-40 transition-all duration-200 shadow-lg shadow-slate-950/10">
      <div v-if="isLandingRoute" class="hidden border-b border-white/10 bg-slate-900/80 text-xs font-semibold text-slate-300 lg:block">
        <div class="mx-auto flex max-w-7xl items-center justify-between px-4 py-2 sm:px-6 lg:px-8">
          <div class="flex items-center gap-5">
            <span>SĐT: 0123 456 789</span>
            <span>Email: support@quickwork.vn</span>
            <span>Địa chỉ: TP. Hồ Chí Minh, Việt Nam</span>
          </div>
          <div class="flex items-center gap-4">
            <span>Facebook</span>
            <span>LinkedIn</span>
          </div>
        </div>
      </div>
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex items-center space-x-8">
            <NuxtLink to="/" class="flex items-center space-x-2">
              <span class="h-9 w-9 rounded-lg bg-gradient-to-tr from-cyan-400 to-slate-700 flex items-center justify-center text-slate-950 font-bold text-lg shadow-md shadow-cyan-500/20">
                Q
              </span>
              <span class="font-extrabold text-xl tracking-tight bg-gradient-to-r from-white to-cyan-200 bg-clip-text text-transparent">
                QuickWork
              </span>
            </NuxtLink>

            <!-- Navigation Links -->
            <nav class="hidden md:flex items-center space-x-1">
              <template v-if="isAdminRoute">
                <NuxtLink
                  v-for="item in adminNavItems"
                  :key="item.to"
                  :to="item.to"
                  class="relative rounded-lg px-3 py-2 text-sm font-bold transition-colors"
                  :class="route.path === item.to
                    ? 'bg-cyan-400 text-slate-950 shadow-sm shadow-cyan-950/30'
                    : 'text-slate-300 hover:bg-white/10 hover:text-white'"
                >
                  {{ item.name }}
                </NuxtLink>
              </template>

              <template v-else-if="isStudentDashboard">
                <span class="px-3 py-2 text-sm font-medium text-slate-400">Không gian Sinh viên</span>
              </template>
              <template v-else-if="isDashboardRoute">
                <button
                  v-for="item in dashboardNavItems"
                  :key="item.id"
                  @click="setDashboardSection(item.id)"
                  :class="[
                    dashboardActiveSection === item.id
                      ? 'bg-cyan-400 text-slate-950 font-semibold shadow-sm shadow-cyan-950/30'
                      : 'text-slate-300 hover:bg-white/10 hover:text-white',
                    'px-3 py-2 rounded-lg text-sm font-medium transition-colors'
                  ]"
                >
                  {{ item.name }}
                </button>
              </template>

              <!-- Shared Public Links -->
              <template v-else>
                <template v-if="isLandingRoute || isPublicAuthRoute">
                  <a
                    v-for="item in landingNavItems"
                    :key="item.id"
                    :href="item.href"
                    @click.prevent="scrollToLandingSection(item.id)"
                    :class="[
                      activeLandingSection === item.id ? 'text-cyan-200' : 'text-slate-300 hover:text-white',
                      'group relative px-3 py-2 rounded-lg text-sm font-medium hover:bg-white/10 transition-colors'
                    ]"
                  >
                    <span>{{ item.name }}</span>
                    <span
                      :class="activeLandingSection === item.id ? 'scale-x-100 opacity-100' : 'scale-x-0 opacity-0 group-hover:scale-x-75 group-hover:opacity-60'"
                      class="absolute bottom-1 left-3 right-3 h-0.5 origin-center rounded-full bg-cyan-300 transition-all duration-300 ease-out"
                    ></span>
                  </a>
                </template>
                <NuxtLink
                  v-else
                  to="/jobs"
                  class="px-3 py-2 rounded-lg text-sm font-medium text-slate-300 hover:bg-white/10 hover:text-white transition-colors"
                  active-class="bg-cyan-400 text-slate-950 font-semibold"
                >
                  Tìm việc làm
                </NuxtLink>

                <!-- Student specific navigation links -->
                <template v-if="isAuthenticated && userRole === 'student'">
                  <NuxtLink
                    to="/student/dashboard"
                    class="px-3 py-2 rounded-lg text-sm font-medium text-slate-300 hover:bg-white/10 hover:text-white transition-colors"
                    active-class="bg-cyan-400 text-slate-950 font-semibold"
                  >
                    Trang cá nhân
                  </NuxtLink>
                </template>

                <!-- Business specific navigation links -->
                <template v-if="isAuthenticated && userRole === 'business'">
                  <NuxtLink
                    to="/business/dashboard"
                    class="px-3 py-2 rounded-lg text-sm font-medium text-slate-300 hover:bg-white/10 hover:text-white transition-colors"
                    active-class="bg-cyan-400 text-slate-950 font-semibold"
                  >
                    Nhà tuyển dụng
                  </NuxtLink>
                </template>

                <!-- Admin specific navigation links -->
                <template v-if="isAuthenticated && userRole === 'admin' && !isAdminRoute">
                  <NuxtLink
                    to="/admin/dashboard"
                    class="px-3 py-2 rounded-lg text-sm font-medium text-slate-300 hover:bg-white/10 hover:text-white transition-colors"
                    active-class="bg-cyan-400 text-slate-950 font-semibold"
                  >
                    Quản trị viên
                  </NuxtLink>
                </template>
              </template>
            </nav>
          </div>

          <!-- Right side Auth elements -->
          <div class="flex items-center space-x-4">
            <template v-if="isAuthenticated">
              <NuxtLink
                v-if="!isDashboardRoute && !isAdminRoute"
                :to="dashboardPath"
                class="hidden sm:inline-flex items-center justify-center rounded-lg bg-cyan-400 px-4 py-2 text-sm font-extrabold text-slate-950 hover:bg-cyan-300"
              >
                Bảng điều khiển
              </NuxtLink>
              <!-- Notification Bell Button -->
              <div class="relative">
                <button
                  @click="isNotifDropdownOpen = !isNotifDropdownOpen"
                  class="relative p-2 rounded-xl border border-white/10 bg-white/5 hover:bg-white/10 text-slate-300 hover:text-white transition-all cursor-pointer"
                  title="Thông báo"
                >
                  <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
                  </svg>
                  <span
                    v-if="unreadNotifCount > 0"
                    class="absolute -top-1 -right-1 flex h-4 w-4 items-center justify-center rounded-full bg-rose-500 text-[10px] font-black text-white ring-2 ring-slate-950 animate-pulse"
                  >
                    {{ unreadNotifCount > 9 ? '9+' : unreadNotifCount }}
                  </span>
                </button>

                <!-- Notification Dropdown Window -->
                <div
                  v-if="isNotifDropdownOpen"
                  class="absolute right-0 mt-2 w-80 sm:w-96 rounded-2xl border border-white/10 bg-slate-900 shadow-2xl shadow-slate-950/80 backdrop-blur-xl z-50 overflow-hidden animate-in fade-in duration-150"
                >
                  <div class="p-3.5 border-b border-white/10 bg-slate-950/80 flex items-center justify-between">
                    <div class="flex items-center gap-2">
                      <span class="text-xs font-black uppercase tracking-wider text-cyan-300">Thông báo mới</span>
                      <span v-if="unreadNotifCount > 0" class="px-2 py-0.5 rounded-full text-[10px] font-bold bg-rose-500/20 text-rose-300 border border-rose-500/30">
                        {{ unreadNotifCount }} chưa đọc
                      </span>
                    </div>
                    <button
                      v-if="unreadNotifCount > 0"
                      @click="markAllNotifsAsRead"
                      class="text-[11px] font-bold text-cyan-300 hover:text-cyan-200 transition-colors cursor-pointer"
                    >
                      Đánh dấu tất cả đã đọc
                    </button>
                  </div>

                  <div class="max-h-80 overflow-y-auto divide-y divide-white/5 custom-scrollbar">
                    <div
                      v-for="notif in notificationsList"
                      :key="notif.id"
                      @click="handleNotificationClick(notif)"
                      :class="[
                        !notif.is_read ? 'bg-cyan-400/5' : 'bg-transparent',
                        'p-3.5 hover:bg-white/5 transition-colors cursor-pointer flex items-start gap-3'
                      ]"
                    >
                      <span class="text-base flex-shrink-0">
                        {{ notifTypeIcon(notif.type) }}
                      </span>
                      <div class="min-w-0 flex-1">
                        <div class="flex items-center justify-between gap-1">
                          <h4 class="text-xs font-bold text-white truncate">{{ notif.title }}</h4>
                          <span v-if="!notif.is_read" class="h-2 w-2 rounded-full bg-cyan-400 flex-shrink-0"></span>
                        </div>
                        <p class="mt-0.5 text-xs text-slate-300 font-medium line-clamp-2 leading-relaxed">{{ notif.message }}</p>
                        <span class="mt-1 block text-[10px] text-slate-500 font-semibold">{{ formatNotifTime(notif.created_at) }}</span>
                      </div>
                    </div>

                    <div v-if="notificationsList.length === 0" class="p-8 text-center text-xs text-slate-500 font-medium">
                      Chưa có thông báo nào.
                    </div>
                  </div>
                </div>
              </div>

              <!-- Logged In Status -->
              <div class="hidden sm:flex flex-col text-right">
                <span class="text-xs font-semibold uppercase tracking-wider text-cyan-300">
                  {{ roleLabel }}
                </span>
                <span class="text-sm font-medium text-slate-200">{{ userEmail || 'Người dùng' }}</span>
              </div>

              <!-- Profile Dropdown trigger / Logout button -->
              <button
                @click="handleLogout"
                class="inline-flex items-center space-x-1.5 px-3 py-2 border border-white/10 rounded-lg text-sm font-semibold text-slate-200 hover:bg-white/10 hover:text-white focus-ring"
              >
                <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                </svg>
                <span>Đăng xuất</span>
              </button>
            </template>

            <template v-else>
              <!-- Guest Links -->
              <div class="flex items-center gap-2.5">
                <NuxtLink
                  v-if="route.path !== '/login'"
                  to="/login"
                  class="inline-flex items-center justify-center py-2 px-4 rounded-xl text-xs font-extrabold border border-white/15 bg-white/5 hover:bg-white/10 text-slate-200 hover:text-white transition-all shadow-sm leading-normal"
                >
                  Đăng nhập
                </NuxtLink>

                <NuxtLink
                  v-if="route.path !== '/register'"
                  to="/register"
                  class="inline-flex items-center justify-center py-2 px-4.5 rounded-xl text-xs font-extrabold border border-transparent text-slate-950 bg-cyan-400 hover:bg-cyan-300 shadow-md shadow-cyan-500/25 transition-all hover:scale-105 active:scale-95 leading-normal"
                >
                  Đăng ký
                </NuxtLink>

                <NuxtLink
                  to="/employer-register"
                  class="hidden lg:inline-flex items-center text-xs font-bold text-slate-300 hover:text-cyan-300 px-3 py-2 transition-colors border-l border-white/10 pl-4 ml-1"
                >
                  Đăng tuyển & tìm hồ sơ
                </NuxtLink>
              </div>
            </template>

            <!-- Mobile menu button -->
            <button
              @click="mobileMenuOpen = !mobileMenuOpen"
              class="md:hidden p-2 rounded-lg text-slate-300 hover:bg-white/10 focus:outline-none focus:ring-2 focus:ring-cyan-400"
            >
              <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path v-if="!mobileMenuOpen" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
                <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>
      </div>

      <!-- Mobile Menu -->
      <div v-if="mobileMenuOpen" class="md:hidden border-t border-white/10 bg-slate-950 px-2 pt-2 pb-3 space-y-1 shadow-lg">
        <template v-if="isAdminRoute">
          <div class="px-3 py-2">
            <p class="text-xs font-semibold uppercase tracking-wider text-cyan-300">Quản trị hệ thống</p>
          </div>
          <NuxtLink
            v-for="item in adminNavItems"
            :key="item.to"
            :to="item.to"
            @click="mobileMenuOpen = false"
            class="block px-3 py-2 rounded-lg text-base font-medium"
            :class="route.path === item.to
              ? 'bg-cyan-400 text-slate-950 font-semibold'
              : 'text-slate-300 hover:bg-white/10 hover:text-white'"
          >
            {{ item.name }}
          </NuxtLink>
        </template>

        <template v-else-if="isStudentDashboard">
          <div class="px-3 py-2">
            <p class="text-xs font-semibold uppercase tracking-wider text-cyan-300">Không gian Sinh viên</p>
            <p class="mt-1 text-xs text-slate-500">Dùng menu bên trong trang để chuyển mục.</p>
          </div>
        </template>
        <template v-else-if="isDashboardRoute">
          <div class="px-3 py-2">
            <p class="text-xs font-semibold uppercase tracking-wider text-cyan-300">{{ dashboardWorkspaceLabel }}</p>
          </div>
          <button
            v-for="item in dashboardNavItems"
            :key="item.id"
            @click="setDashboardSection(item.id); mobileMenuOpen = false"
            :class="[
              dashboardActiveSection === item.id
                ? 'bg-cyan-400 text-slate-950 font-semibold'
                : 'text-slate-300 hover:bg-white/10 hover:text-white',
              'block w-full text-left px-3 py-2 rounded-lg text-base font-medium'
            ]"
          >
            {{ item.name }}
          </button>
        </template>

        <template v-else>
          <template v-if="isLandingRoute || isPublicAuthRoute">
            <a
              v-for="item in landingNavItems"
              :key="item.id"
              :href="item.href"
              @click.prevent="scrollToLandingSection(item.id); mobileMenuOpen = false"
              :class="[
                activeLandingSection === item.id
                  ? 'bg-cyan-400/10 text-cyan-200 ring-1 ring-cyan-400/20'
                  : 'text-slate-300 hover:bg-white/10 hover:text-white',
                'relative block px-3 py-2 rounded-lg text-base font-medium transition-colors'
              ]"
            >
              {{ item.name }}
              <span
                :class="activeLandingSection === item.id ? 'opacity-100 translate-x-0' : 'opacity-0 -translate-x-2'"
                class="absolute bottom-1 left-3 h-0.5 w-10 rounded-full bg-cyan-300 transition-all duration-300"
              ></span>
            </a>
          </template>
          <NuxtLink
            v-else
            to="/jobs"
            @click="mobileMenuOpen = false"
            class="block px-3 py-2 rounded-lg text-base font-medium hover:bg-white/10 text-slate-300"
            active-class="bg-cyan-400 text-slate-950 font-semibold"
          >
            Tìm việc làm
          </NuxtLink>

          <template v-if="isAuthenticated && userRole === 'student'">
            <NuxtLink
              to="/student/dashboard"
              @click="mobileMenuOpen = false"
              class="block px-3 py-2 rounded-lg text-base font-medium hover:bg-white/10 text-slate-300"
              active-class="bg-cyan-400 text-slate-950 font-semibold"
            >
              Trang cá nhân
            </NuxtLink>
          </template>
          <template v-if="isAuthenticated && userRole === 'business'">
            <NuxtLink
              to="/business/dashboard"
              @click="mobileMenuOpen = false"
              class="block px-3 py-2 rounded-lg text-base font-medium hover:bg-white/10 text-slate-300"
              active-class="bg-cyan-400 text-slate-950 font-semibold"
            >
              Nhà tuyển dụng
            </NuxtLink>
          </template>
          <template v-if="isAuthenticated && userRole === 'admin' && !isAdminRoute">
            <NuxtLink
              to="/admin/dashboard"
              @click="mobileMenuOpen = false"
              class="block px-3 py-2 rounded-lg text-base font-medium hover:bg-white/10 text-slate-300"
              active-class="bg-cyan-400 text-slate-950 font-semibold"
            >
              Quản trị viên
            </NuxtLink>
          </template>
        </template>
      </div>
    </header>

    <!-- Main Content App Router -->
    <main class="flex-grow pb-16 md:pb-0">
      <NuxtPage />
    </main>

    <!-- Global Toast Container (Mục 4) -->
    <ToastContainer />

    <!-- Mobile Bottom Navigation (Hiện trên di động < 768px) -->
    <MobileBottomNav />

    <!-- Simple Modern Footer -->
    <footer v-if="showFooter" class="bg-slate-900 text-slate-400 py-6 border-t border-slate-800">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 flex flex-col sm:flex-row justify-between items-center space-y-4 sm:space-y-0 text-sm">
        <div class="flex items-center space-x-2">
          <span class="font-semibold text-white">QuickWork</span>
          <span>&copy; 2026. Bản quyền thuộc về QuickWork.</span>
        </div>
        <div class="flex space-x-6">
          <a href="#" class="hover:text-white transition-colors">Chính sách bảo mật</a>
          <a href="#" class="hover:text-white transition-colors">Điều khoản dịch vụ</a>
          <a href="#" class="hover:text-white transition-colors">Hỗ trợ khách hàng</a>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount, nextTick, watch } from 'vue'
import { useRoute } from 'vue-router'
import { useAuth } from '~/composables/useAuth'
import MobileBottomNav from '~/components/common/MobileBottomNav.vue'
import ToastContainer from '~/components/common/ToastContainer.vue'

const route = useRoute()
const { logout, user, isAuthenticated, userRole, fetchUser } = useAuth()

const mobileMenuOpen = ref(false)
const studentDashboardActiveSection = useState<string>('studentDashboardActiveSection', () => 'jobs')
const businessDashboardActiveSection = useState<string>('businessDashboardActiveSection', () => 'dashboard')

const studentDashboardNavItems = [
  { id: 'jobs', name: 'Tìm việc làm' },
  { id: 'profile', name: 'Hồ sơ cá nhân' },
  { id: 'applications', name: 'Đơn ứng tuyển' },
  { id: 'wallet', name: 'Ví tiền' }
]

const businessDashboardNavItems = [
  { id: 'dashboard', name: 'Tổng quan' },
  { id: 'profile', name: 'Hồ sơ công ty' },
  { id: 'jobs', name: 'Bài đăng việc làm' },
  { id: 'applicants', name: 'Danh sách ứng viên' }
]

const adminNavItems = [
  { name: 'Tổng quan', to: '/admin/dashboard' },
  { name: 'Duyệt Doanh nghiệp', to: '/admin/businesses/pending' }
]

const landingNavItems = [
  { id: 'explore-jobs', name: 'Tìm việc làm', href: '#explore-jobs' },
  { id: 'categories', name: 'Danh mục ngành', href: '#categories' },
  { id: 'features', name: 'Tính năng & AI', href: '#features' },
  { id: 'contact', name: 'Liên hệ', href: '#contact' }
]

const activeLandingSection = ref('home')
let landingObserver: IntersectionObserver | null = null

// Only show footer on the public Landing Page (/) and hide on Dashboards & Auth pages
const showFooter = computed(() => {
  return route.path === '/'
})

const userEmail = computed(() => user.value?.email)
const isLandingRoute = computed(() => route.path === '/')
const isPublicAuthRoute = computed(() => ['/login', '/register', '/employer-register'].includes(route.path))
const isStudentDashboard = computed(() => route.path === '/student/dashboard')
const isBusinessDashboard = computed(() => route.path === '/business/dashboard')
const isAdminRoute = computed(() => route.path.startsWith('/admin'))
const isDashboardRoute = computed(() => isStudentDashboard.value || isBusinessDashboard.value)
const roleLabel = computed(() => {
  if (userRole.value === 'student') return 'Sinh viên'
  if (userRole.value === 'business') return 'Doanh nghiệp'
  if (userRole.value === 'admin') return 'Quản trị viên'
  return 'Người dùng'
})
const dashboardPath = computed(() => {
  if (userRole.value === 'business') return '/business/dashboard'
  if (userRole.value === 'admin') return '/admin/dashboard'
  return '/student/dashboard'
})
const dashboardNavItems = computed(() => {
  return isBusinessDashboard.value ? businessDashboardNavItems : studentDashboardNavItems
})
const dashboardActiveSection = computed(() => {
  return isBusinessDashboard.value
    ? businessDashboardActiveSection.value
    : studentDashboardActiveSection.value
})
const dashboardWorkspaceLabel = computed(() => {
  return isBusinessDashboard.value ? 'Không gian Doanh nghiệp' : 'Không gian Sinh viên'
})

const setDashboardSection = (sectionId: string) => {
  if (isBusinessDashboard.value) {
    businessDashboardActiveSection.value = sectionId
    return
  }

  studentDashboardActiveSection.value = sectionId
}

const getHeaderOffset = () => {
  if (!import.meta.client) return 80
  const header = document.querySelector('header')
  return (header?.getBoundingClientRect().height || 80) + 8
}

const scrollToLandingSection = async (sectionId: string) => {
  activeLandingSection.value = sectionId

  if (!isLandingRoute.value) {
    await navigateTo(`/#${sectionId}`)
    return
  }

  await nextTick()
  const target = document.getElementById(sectionId)
  if (!target) return

  const top = target.getBoundingClientRect().top + window.scrollY - getHeaderOffset()
  window.history.replaceState(null, '', `#${sectionId}`)
  window.scrollTo({ top, behavior: 'smooth' })
}

const setupLandingObserver = async () => {
  if (!import.meta.client) return

  landingObserver?.disconnect()
  landingObserver = null

  if (!isLandingRoute.value) return

  await nextTick()

  const sections = landingNavItems
    .map(item => document.getElementById(item.id))
    .filter((section): section is HTMLElement => Boolean(section))

  if (sections.length === 0) return

  const observerOptions = {
    root: null,
    rootMargin: `-${getHeaderOffset()}px 0px -55% 0px`,
    threshold: [0.12, 0.35, 0.6]
  }

  landingObserver = new IntersectionObserver(entries => {
    const visible = entries
      .filter(entry => entry.isIntersecting)
      .sort((a, b) => b.intersectionRatio - a.intersectionRatio)[0]

    if (visible?.target?.id) {
      activeLandingSection.value = visible.target.id
    }
  }, observerOptions)

  sections.forEach(section => landingObserver?.observe(section))

  const hashSection = window.location.hash.replace('#', '')
  if (landingNavItems.some(item => item.id === hashSection)) {
    activeLandingSection.value = hashSection
  }
}

// Notifications state & methods
const isNotifDropdownOpen = ref(false)
const notificationsList = ref<any[]>([])
const unreadNotifCount = ref(0)
let notifPollInterval: any = null

watch(isNotifDropdownOpen, async (newVal) => {
  if (newVal && unreadNotifCount.value > 0) {
    await markAllNotifsAsRead()
  }
})

const fetchNotifications = async () => {
  if (!isAuthenticated.value) return
  try {
    const res = await api.get('/api/notifications?limit=20')
    notificationsList.value = res.data || []
    unreadNotifCount.value = res.unread_count || 0
  } catch (e) {
    // Silent polling
  }
}

const markAllNotifsAsRead = async () => {
  try {
    await api.patch('/api/notifications/read-all')
    unreadNotifCount.value = 0
    notificationsList.value.forEach(n => n.is_read = true)
  } catch (e) {
    console.error('Failed to mark notifications read:', e)
  }
}

const handleNotificationClick = async (notif: any) => {
  if (!notif.is_read) {
    try {
      await api.patch(`/api/notifications/${notif.id}/read`)
      notif.is_read = true
      if (unreadNotifCount.value > 0) unreadNotifCount.value--
    } catch (e) {}
  }

  isNotifDropdownOpen.value = false

  if (userRole.value === 'business') {
    await navigateTo('/business/dashboard')
  } else if (userRole.value === 'student') {
    await navigateTo('/student/dashboard')
  }
}

const notifTypeIcon = (type: string) => {
  switch (type) {
    case 'chat': return '💬'
    case 'offer': return '🎉'
    case 'application': return '📄'
    case 'escrow': return '💰'
    default: return '🔔'
  }
}

const formatNotifTime = (timeStr: string) => {
  if (!timeStr) return ''
  const d = new Date(timeStr)
  return d.toLocaleTimeString('vi-VN', { hour: '2-digit', minute: '2-digit' }) + ' ' + d.toLocaleDateString('vi-VN')
}

// Fetch user profile on layout mounting to ensure session state hydration
onMounted(async () => {
  try {
    await fetchUser()
  } catch (e) {
    console.error('Session restoration failed:', e)
  }

  await setupLandingObserver()
})

watch(() => route.fullPath, async () => {
  mobileMenuOpen.value = false
  isNotifDropdownOpen.value = false
  await setupLandingObserver()
})

watch(isAuthenticated, async (newVal) => {
  if (newVal) {
    await fetchNotifications()
    if (!notifPollInterval) {
      notifPollInterval = setInterval(fetchNotifications, 10000)
    }
  } else {
    if (notifPollInterval) {
      clearInterval(notifPollInterval)
      notifPollInterval = null
    }
    notificationsList.value = []
    unreadNotifCount.value = 0
  }
})

onBeforeUnmount(() => {
  landingObserver?.disconnect()
  if (notifPollInterval) {
    clearInterval(notifPollInterval)
  }
})

const handleLogout = async () => {
  if (notifPollInterval) {
    clearInterval(notifPollInterval)
    notifPollInterval = null
  }
  await logout()
}
</script>

