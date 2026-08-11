<template>
  <div class="rounded-3xl border border-white/10 bg-slate-950/80 p-6 shadow-2xl shadow-slate-950/50 backdrop-blur ring-1 ring-cyan-400/10">
    <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between border-b border-white/10 pb-4">
      <div>
        <span class="inline-flex rounded-full bg-cyan-400/10 px-3 py-1 text-xs font-bold uppercase tracking-wider text-cyan-200 ring-1 ring-cyan-400/25">
          Trực quan hóa Dữ liệu (Data Visualization)
        </span>
        <h3 class="mt-2 text-xl font-extrabold text-white">Xu hướng Tiền lương Giải ngân & Tăng trưởng</h3>
      </div>
      <div class="flex gap-2">
        <button
          @click="activeChart = 'disbursement'"
          class="rounded-xl px-3.5 py-1.5 text-xs font-extrabold transition-all"
          :class="activeChart === 'disbursement' ? 'bg-cyan-400 text-slate-950 shadow-md shadow-cyan-500/20' : 'bg-white/5 text-slate-300 hover:bg-white/10'"
        >
          Biểu đồ Giải ngân
        </button>
        <button
          @click="activeChart = 'growth'"
          class="rounded-xl px-3.5 py-1.5 text-xs font-extrabold transition-all"
          :class="activeChart === 'growth' ? 'bg-cyan-400 text-slate-950 shadow-md shadow-cyan-500/20' : 'bg-white/5 text-slate-300 hover:bg-white/10'"
        >
          Tăng trưởng Người dùng
        </button>
      </div>
    </div>

    <!-- Chart 1: Biểu đồ Giải ngân Tiền lương SVG Area Chart -->
    <div v-if="activeChart === 'disbursement'" class="mt-6">
      <div class="mb-4 flex items-center justify-between text-xs font-bold text-slate-400">
        <span>Tổng giải ngân hiện tại: <strong class="text-emerald-300 text-sm ml-1">{{ formatCurrency(totalDisbursed) }}</strong></span>
        <span>Đơn vị: VNĐ</span>
      </div>

      <div class="relative h-64 w-full">
        <svg class="h-full w-full overflow-visible" viewBox="0 0 500 200" preserveAspectRatio="none">
          <defs>
            <linearGradient id="disbursementGradient" x1="0" y1="0" x2="0" y2="1">
              <stop offset="0%" stop-color="#22d3ee" stop-opacity="0.35" />
              <stop offset="100%" stop-color="#10b981" stop-opacity="0.0" />
            </linearGradient>
          </defs>

          <!-- Gridlines -->
          <line x1="0" y1="40" x2="500" y2="40" stroke="rgba(255,255,255,0.06)" stroke-dasharray="4 4" />
          <line x1="0" y1="90" x2="500" y2="90" stroke="rgba(255,255,255,0.06)" stroke-dasharray="4 4" />
          <line x1="0" y1="140" x2="500" y2="140" stroke="rgba(255,255,255,0.06)" stroke-dasharray="4 4" />

          <!-- Gradient Area -->
          <polygon
            points="0,180 0,140 83,110 166,130 250,80 333,65 416,40 500,20 500,180"
            fill="url(#disbursementGradient)"
          />

          <!-- Smooth Line -->
          <polyline
            points="0,140 83,110 166,130 250,80 333,65 416,40 500,20"
            fill="none"
            stroke="#22d3ee"
            stroke-width="3.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />

          <!-- Data Points -->
          <circle v-for="(pt, idx) in chartPoints" :key="idx" :cx="pt.x" :cy="pt.y" r="5" class="fill-slate-950 stroke-cyan-300 stroke-[3] transition-all hover:r-7 cursor-pointer" />
        </svg>

        <!-- Month Labels -->
        <div class="mt-3 flex justify-between text-xs font-bold text-slate-400">
          <span v-for="m in months" :key="m">{{ m }}</span>
        </div>
      </div>
    </div>

    <!-- Chart 2: Tăng trưởng Sinh viên vs Doanh nghiệp (Bar Chart) -->
    <div v-else class="mt-6 space-y-5">
      <div class="grid grid-cols-2 gap-4 text-center">
        <div class="rounded-2xl border border-cyan-400/20 bg-cyan-400/10 p-4">
          <p class="text-xs font-bold uppercase tracking-wider text-cyan-200">Tỷ lệ Sinh viên</p>
          <p class="mt-1 text-2xl font-extrabold text-white">{{ studentCount }} sinh viên</p>
        </div>
        <div class="rounded-2xl border border-emerald-400/20 bg-emerald-400/10 p-4">
          <p class="text-xs font-bold uppercase tracking-wider text-emerald-200">Tỷ lệ Doanh nghiệp</p>
          <p class="mt-1 text-2xl font-extrabold text-white">{{ businessCount }} công ty</p>
        </div>
      </div>

      <div class="space-y-3 pt-2">
        <div>
          <div class="mb-1 flex justify-between text-xs font-bold text-slate-300">
            <span>Sinh viên đăng ký mới ca tháng</span>
            <span class="text-cyan-300">76%</span>
          </div>
          <div class="h-3 w-full overflow-hidden rounded-full bg-slate-900 border border-white/10">
            <div class="h-full rounded-full bg-gradient-to-r from-cyan-400 to-blue-500 transition-all duration-700" style="width: 76%"></div>
          </div>
        </div>

        <div>
          <div class="mb-1 flex justify-between text-xs font-bold text-slate-300">
            <span>Doanh nghiệp xác minh KYB thành công</span>
            <span class="text-emerald-300">88%</span>
          </div>
          <div class="h-3 w-full overflow-hidden rounded-full bg-slate-900 border border-white/10">
            <div class="h-full rounded-full bg-gradient-to-r from-emerald-400 to-teal-500 transition-all duration-700" style="width: 88%"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps<{
  totalDisbursed: number
  studentCount: number
  businessCount: number
}>()

const activeChart = ref<'disbursement' | 'growth'>('disbursement')

const months = ['Tháng 3', 'Tháng 4', 'Tháng 5', 'Tháng 6', 'Tháng 7', 'Tháng 8 (Hiện tại)']
const chartPoints = [
  { x: 0, y: 140 },
  { x: 83, y: 110 },
  { x: 166, y: 130 },
  { x: 250, y: 80 },
  { x: 333, y: 65 },
  { x: 416, y: 40 },
  { x: 500, y: 20 }
]

const formatCurrency = (value: number) =>
  new Intl.NumberFormat('vi-VN', { style: 'currency', currency: 'VND', maximumFractionDigits: 0 }).format(value)
</script>
