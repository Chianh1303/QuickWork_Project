<template>
  <nav class="flex flex-wrap items-center gap-2 sm:gap-3">
    <NuxtLink
      to="/admin/businesses/pending"
      class="inline-flex items-center justify-center rounded-xl bg-cyan-400 px-4 py-2 text-xs sm:text-sm font-extrabold text-slate-950 shadow-md shadow-cyan-500/20 transition-all hover:bg-cyan-300"
    >
      Duyệt Doanh Nghiệp Pending
    </NuxtLink>

    <NuxtLink
      to="/admin/businesses"
      class="inline-flex items-center justify-center rounded-xl border border-cyan-400/30 bg-cyan-400/10 px-4 py-2 text-xs sm:text-sm font-extrabold text-cyan-200 transition-all hover:bg-cyan-400 hover:text-slate-950"
    >
      Quản lý Doanh nghiệp
    </NuxtLink>

    <NuxtLink
      to="/admin/students"
      class="inline-flex items-center justify-center rounded-xl border border-cyan-400/30 bg-cyan-400/10 px-4 py-2 text-xs sm:text-sm font-extrabold text-cyan-200 transition-all hover:bg-cyan-400 hover:text-slate-950"
    >
      Quản lý Sinh viên
    </NuxtLink>

    <NuxtLink
      to="/admin/tickets"
      class="inline-flex items-center justify-center rounded-xl border border-cyan-400/30 bg-cyan-400/10 px-4 py-2 text-xs sm:text-sm font-extrabold text-cyan-200 transition-all hover:bg-cyan-400 hover:text-slate-950"
    >
      Xử lý Khiếu nại
    </NuxtLink>

    <NuxtLink
      to="/admin/categories"
      class="inline-flex items-center justify-center rounded-xl border border-cyan-400/30 bg-cyan-400/10 px-4 py-2 text-xs sm:text-sm font-extrabold text-cyan-200 transition-all hover:bg-cyan-400 hover:text-slate-950"
    >
      Quản lý Danh mục
    </NuxtLink>
  </nav>
</template>
