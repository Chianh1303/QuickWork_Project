<template>
  <div class="dashboard-body min-h-screen bg-slate-950 text-slate-100 font-sans">
    <AdminHeaderNav />
    <main class="mx-auto w-full max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
      <slot />
    </main>
  </div>
</template>

<script setup lang="ts">
import AdminHeaderNav from '~/components/admin/AdminHeaderNav.vue'
</script>
