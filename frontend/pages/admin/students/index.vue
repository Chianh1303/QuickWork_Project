<template>
  <AdminShell>
    <div class="space-y-6">
      <!-- Header Banner -->
      <section class="rounded-3xl border border-white/10 bg-slate-950/80 p-6 shadow-2xl shadow-slate-950/40 ring-1 ring-cyan-400/10 lg:p-8">
        <div class="flex flex-col gap-6">
          <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
            <div>
              <h1 class="text-3xl font-extrabold tracking-tight text-white sm:text-4xl">
                Quản lý Sinh viên Hệ thống
              </h1>
              <p class="mt-2 text-sm font-semibold text-slate-300">
                Quản lý danh sách sinh viên, xem thông tin chi tiết, khóa hoặc mở khóa tài khoản người dùng.
              </p>
            </div>

            <NuxtLink
              to="/admin/dashboard"
              class="inline-flex items-center justify-center rounded-xl border border-white/10 bg-white/10 px-4 py-2.5 text-sm font-bold text-slate-200 transition-colors hover:bg-white/15"
            >
              Quay lại Dashboard
            </NuxtLink>
          </div>
        </div>
      </section>

      <!-- Filters & Search Bar -->
      <section class="rounded-2xl border border-white/10 bg-slate-950/70 p-5 shadow-xl shadow-slate-950/25 ring-1 ring-cyan-400/10">
        <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
          <div class="relative flex-1">
            <input
              v-model="searchQuery"
              type="text"
              placeholder="Tìm kiếm theo Tên, Email hoặc Số điện thoại..."
              class="w-full rounded-xl border border-white/10 bg-slate-900/90 px-4 py-2.5 text-sm text-white placeholder-slate-400 focus:border-cyan-400 focus:outline-none focus:ring-1 focus:ring-cyan-400"
              @keyup.enter="fetchStudents"
            />
          </div>

          <div class="flex items-center gap-3">
            <select
              v-model="statusFilter"
              class="rounded-xl border border-white/10 bg-slate-900/90 px-3.5 py-2.5 text-sm text-white focus:border-cyan-400 focus:outline-none"
              @change="fetchStudents"
            >
              <option value="">Tất cả trạng thái</option>
              <option value="approved">🟢 Đang hoạt động (Approved)</option>
              <option value="locked">🔴 Đã bị khóa (Locked)</option>
            </select>

            <button
              @click="fetchStudents"
              class="inline-flex items-center justify-center rounded-xl bg-cyan-400 px-4 py-2.5 text-sm font-extrabold text-slate-950 transition-colors hover:bg-cyan-300"
            >
              Tìm kiếm
            </button>
          </div>
        </div>
      </section>

      <!-- Student Data Table -->
      <section class="overflow-hidden rounded-2xl border border-white/10 bg-slate-950/70 shadow-xl shadow-slate-950/25 ring-1 ring-cyan-400/10">
        <div v-if="isLoading" class="p-8 text-center text-slate-400">
          <span class="animate-pulse text-sm font-bold">Đang tải danh sách sinh viên...</span>
        </div>

        <div v-else-if="students.length === 0" class="p-8 text-center text-slate-400">
          <p class="text-base font-semibold">Không tìm thấy sinh viên nào phù hợp.</p>
        </div>

        <div v-else class="overflow-x-auto">
          <table class="w-full text-left text-sm text-slate-200">
            <thead class="bg-slate-900/90 text-xs uppercase text-slate-400">
              <tr>
                <th class="px-5 py-3.5 font-extrabold">Sinh viên</th>
                <th class="px-5 py-3.5 font-extrabold">SĐT & Giới tính</th>
                <th class="px-5 py-3.5 font-extrabold">Kỹ năng</th>
                <th class="px-5 py-3.5 font-extrabold">Trạng thái</th>
                <th class="px-5 py-3.5 font-extrabold text-right">Thao tác</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-white/5">
              <tr v-for="st in students" :key="st.student_id" class="hover:bg-white/[0.02]">
                <td class="px-5 py-4">
                  <div class="flex items-center gap-3">
                    <img :src="st.avatar_url || 'https://api.dicebear.com/7.x/adventurer/svg?seed=' + st.full_name" class="h-10 w-10 rounded-full bg-slate-800" />
                    <div>
                      <p class="font-extrabold text-white">{{ st.full_name }}</p>
                      <p class="text-xs text-slate-400">{{ st.email }}</p>
                    </div>
                  </div>
                </td>

                <td class="px-5 py-4">
                  <p class="font-bold text-slate-200">{{ st.phone || 'Chưa cập nhật' }}</p>
                  <p class="text-xs text-slate-400">{{ st.gender || 'Chưa rõ' }}</p>
                </td>

                <td class="px-5 py-4 max-w-xs truncate text-xs text-slate-300">
                  {{ st.skills || 'Chưa có thông tin kỹ năng' }}
                </td>

                <td class="px-5 py-4">
                  <span
                    class="inline-flex items-center gap-1.5 rounded-full px-2.5 py-1 text-xs font-bold border"
                    :class="st.status === 'locked' ? 'bg-rose-400/10 text-rose-300 border-rose-400/30' : 'bg-emerald-400/10 text-emerald-300 border-emerald-400/30'"
                  >
                    <span>{{ st.status === 'locked' ? '🔴 Đã bị khóa' : '🟢 Hoạt động' }}</span>
                  </span>
                </td>

                <td class="px-5 py-4 text-right space-x-2">
                  <button
                    @click="viewDetail(st)"
                    class="rounded-lg bg-white/10 px-3 py-1.5 text-xs font-bold text-slate-200 hover:bg-white/20 transition-colors"
                  >
                    Xem chi tiết
                  </button>

                  <button
                    v-if="st.status === 'locked'"
                    @click="toggleStatus(st, 'approved')"
                    class="rounded-lg bg-emerald-400/20 px-3 py-1.5 text-xs font-extrabold text-emerald-300 hover:bg-emerald-400/30 border border-emerald-400/30 transition-colors"
                  >
                    Mở khóa
                  </button>

                  <button
                    v-else
                    @click="toggleStatus(st, 'locked')"
                    class="rounded-lg bg-rose-400/20 px-3 py-1.5 text-xs font-extrabold text-rose-300 hover:bg-rose-400/30 border border-rose-400/30 transition-colors"
                  >
                    Khóa tài khoản
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </section>
    </div>

    <!-- Student Detail Modal -->
    <div v-if="selectedStudent" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-sm">
      <div class="w-full max-w-lg rounded-2xl border border-white/10 bg-slate-900 p-6 shadow-2xl space-y-4">
        <div class="flex items-center justify-between border-b border-white/10 pb-3">
          <h3 class="text-lg font-extrabold text-white">Chi tiết Hồ sơ Sinh viên</h3>
          <button @click="selectedStudent = null" class="text-slate-400 hover:text-white text-lg font-bold">✕</button>
        </div>

        <div class="space-y-3 text-sm text-slate-300">
          <div class="flex items-center gap-3">
            <img :src="selectedStudent.avatar_url" class="h-14 w-14 rounded-full bg-slate-800 border border-cyan-400/30" />
            <div>
              <p class="text-base font-extrabold text-white">{{ selectedStudent.full_name }}</p>
              <p class="text-xs text-cyan-300 font-semibold">{{ selectedStudent.email }}</p>
            </div>
          </div>

          <div class="grid grid-cols-2 gap-2 bg-slate-950/60 p-3 rounded-xl border border-white/5">
            <div>
              <span class="text-xs text-slate-400 block font-bold">SỐ ĐIỆN THOẠI</span>
              <span class="font-semibold text-white">{{ selectedStudent.phone || 'N/A' }}</span>
            </div>
            <div>
              <span class="text-xs text-slate-400 block font-bold">GIỚI TÍNH</span>
              <span class="font-semibold text-white">{{ selectedStudent.gender || 'N/A' }}</span>
            </div>
          </div>

          <div>
            <span class="text-xs text-slate-400 block font-bold mb-1">KỸ NĂNG & MÔ TẢ</span>
            <p class="bg-slate-950/60 p-3 rounded-xl border border-white/5 text-xs font-medium text-slate-200">
              {{ selectedStudent.skills || 'Chưa cập nhật' }}
            </p>
          </div>

          <div v-if="selectedStudent.cv_url">
            <span class="text-xs text-slate-400 block font-bold mb-1">FILE CV PDF</span>
            <a :href="selectedStudent.cv_url" target="_blank" class="inline-flex items-center gap-1.5 text-xs font-bold text-cyan-300 hover:underline">
              📄 Xem file CV PDF ứng viên
            </a>
          </div>
        </div>

        <div class="pt-4 border-t border-white/10 flex justify-end">
          <button @click="selectedStudent = null" class="rounded-xl bg-cyan-400 px-5 py-2 text-xs font-extrabold text-slate-950">
            Đóng cửa sổ
          </button>
        </div>
      </div>
    </div>
  </AdminShell>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import AdminShell from '~/components/admin/AdminShell.vue'
import { useApi } from '~/composables/useApi'
import { useToast } from '~/composables/useToast'

interface StudentItem {
  student_id: number
  user_id: number
  full_name: string
  email: string
  phone: string
  gender: string
  avatar_url: string
  skills: string
  cv_url: string
  status: string
  created_at: string
}

const api = useApi()
const { success, error } = useToast()

const students = ref<StudentItem[]>([])
const isLoading = ref(false)
const searchQuery = ref('')
const statusFilter = ref('')
const selectedStudent = ref<StudentItem | null>(null)

const fetchStudents = async () => {
  isLoading.value = true
  try {
    const params = new URLSearchParams()
    if (searchQuery.value) params.append('search', searchQuery.value)
    if (statusFilter.value) params.append('status', statusFilter.value)

    const res = await api.get<{ items: StudentItem[] }>(`/api/admin/students?${params.toString()}`)
    students.value = res.items || []
  } catch (err: any) {
    error('Không thể tải danh sách sinh viên')
  } finally {
    isLoading.value = false
  }
}

const viewDetail = (st: StudentItem) => {
  selectedStudent.value = st
}

const toggleStatus = async (st: StudentItem, newStatus: string) => {
  const actionText = newStatus === 'locked' ? 'Khóa' : 'Mở khóa'
  if (!confirm(`Bạn có chắc chắn muốn ${actionText} tài khoản của sinh viên "${st.full_name}"?`)) {
    return
  }

  try {
    await api.put(`/api/admin/students/${st.student_id}/status`, { status: newStatus })
    st.status = newStatus
    success(`Đã ${actionText} tài khoản sinh viên thành công!`)
  } catch (err: any) {
    error(`Không thể ${actionText} tài khoản sinh viên`)
  }
}

onMounted(() => {
  fetchStudents()
})
</script>
