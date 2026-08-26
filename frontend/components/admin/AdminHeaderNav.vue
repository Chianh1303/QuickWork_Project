<template>
  <header class="border-b border-cyan-500/20 bg-slate-950/95 backdrop-blur-xl sticky top-0 z-40 shadow-xl shadow-slate-950/30">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">
        
        <!-- Logo & Portal Badge -->
        <div class="flex items-center space-x-3">
          <NuxtLink to="/admin/dashboard" class="flex items-center space-x-2.5 group">
            <div class="h-9 w-9 rounded-xl bg-gradient-to-tr from-cyan-500 via-blue-600 to-emerald-400 flex items-center justify-center text-white font-black text-lg shadow-lg shadow-cyan-500/30 group-hover:scale-105 transition-transform">
              QW
            </div>
            <div>
              <span class="font-extrabold text-lg tracking-tight text-white group-hover:text-cyan-300 transition-colors">
                QuickWork
              </span>
              <span class="ml-2 inline-flex items-center rounded-full bg-rose-500/15 px-2.5 py-0.5 text-[10px] font-black uppercase tracking-wider text-rose-300 ring-1 ring-rose-500/30">
                Admin Control
              </span>
            </div>
          </NuxtLink>
        </div>

        <!-- Navigation Links -->
        <nav class="hidden md:flex items-center space-x-1.5">
          <NuxtLink
            v-for="item in navItems"
            :key="item.to"
            :to="item.to"
            :class="[
              route.path === item.to
                ? 'bg-gradient-to-r from-cyan-500 via-blue-600 to-emerald-500 text-white font-extrabold shadow-md shadow-cyan-500/20'
                : 'text-slate-300 hover:bg-cyan-500/10 hover:text-white font-semibold',
              'px-3.5 py-2 rounded-xl text-xs transition-all'
            ]"
          >
            {{ item.name }}
          </NuxtLink>
        </nav>

        <!-- Admin Logout Button -->
        <div class="flex items-center space-x-3">
          <button
            @click="handleLogout"
            class="inline-flex items-center space-x-1.5 px-3.5 py-2 border border-rose-500/20 rounded-xl text-xs font-extrabold text-rose-300 bg-rose-500/10 hover:bg-rose-500/20 hover:text-rose-200 transition-all shadow-sm"
          >
            <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
            </svg>
            <span class="hidden sm:inline">Đăng xuất</span>
          </button>
        </div>

      </div>

      <!-- Mobile Navigation Sub-bar -->
      <div class="md:hidden border-t border-indigo-500/10 py-2.5 flex overflow-x-auto space-x-1.5 scrollbar-none">
        <NuxtLink
          v-for="item in navItems"
          :key="item.to"
          :to="item.to"
          :class="[
            route.path === item.to
              ? 'bg-indigo-500 text-white font-black'
              : 'text-slate-300 hover:bg-white/10 font-semibold',
            'whitespace-nowrap px-3 py-1.5 rounded-lg text-xs transition-all'
          ]"
        >
          {{ item.name }}
        </NuxtLink>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import { useAuth } from '~/composables/useAuth'

const route = useRoute()
const { logout } = useAuth()

const handleLogout = () => {
  logout()
}

const navItems = [
  { name: 'Dashboard Tổng Quan', to: '/admin/dashboard' },
  { name: 'Duyệt Doanh Nghiệp', to: '/admin/businesses/pending' },
  { name: 'Duyệt Tin Tuyển Dụng', to: '/admin/jobs/pending' },
  { name: 'Quản Lý Doanh Nghiệp', to: '/admin/businesses' },
  { name: 'Quản Lý Sinh Viên', to: '/admin/students' },
  { name: 'Xử Lý Khiếu Nại', to: '/admin/tickets' },
  { name: 'Quản Lý Danh Mục', to: '/admin/categories' }
]
</script>
