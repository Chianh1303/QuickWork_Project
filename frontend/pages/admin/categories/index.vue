<template>
  <AdminShell>
    <div class="space-y-6">
      <!-- Header Banner -->
      <section class="rounded-3xl border border-white/10 bg-slate-950/80 p-6 shadow-2xl shadow-slate-950/40 ring-1 ring-cyan-400/10 lg:p-8">
        <div class="flex flex-col gap-6">
          <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
            <div>
              <h1 class="text-3xl font-extrabold tracking-tight text-white sm:text-4xl">
                Quản lý Danh mục & Kỹ năng Chuẩn hóa
              </h1>
              <p class="mt-2 text-sm font-semibold text-slate-300">
                Quản lý danh sách Ngành nghề tuyển dụng và Kỹ năng chuẩn hóa phục vụ cho Bộ máy AI Match Job.
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

      <!-- Tabs Header -->
      <section class="rounded-2xl border border-white/10 bg-slate-950/70 p-2 shadow-xl shadow-slate-950/25 ring-1 ring-cyan-400/10">
        <div class="flex gap-2">
          <button
            @click="activeTab = 'categories'"
            class="flex-1 rounded-xl px-5 py-3 text-sm font-extrabold transition-all"
            :class="activeTab === 'categories' ? 'bg-cyan-400 text-slate-950 shadow-lg shadow-cyan-500/20' : 'text-slate-400 hover:text-white hover:bg-white/5'"
          >
            📁 Ngành nghề Tuyển dụng
          </button>
          <button
            @click="activeTab = 'skills'"
            class="flex-1 rounded-xl px-5 py-3 text-sm font-extrabold transition-all"
            :class="activeTab === 'skills' ? 'bg-cyan-400 text-slate-950 shadow-lg shadow-cyan-500/20' : 'text-slate-400 hover:text-white hover:bg-white/5'"
          >
            ⚡ Kỹ năng Chuẩn hóa (AI Taxonomy)
          </button>
        </div>
      </section>

      <!-- Tab 1: Categories Management -->
      <div v-if="activeTab === 'categories'" class="space-y-6">
        <!-- Add Category Form -->
        <section class="rounded-2xl border border-white/10 bg-slate-950/70 p-5 shadow-xl ring-1 ring-cyan-400/10">
          <h3 class="text-sm font-extrabold uppercase tracking-wide text-cyan-300 mb-4">Thêm Ngành nghề tuyển dụng mới</h3>
          <div class="grid gap-4 md:grid-cols-[1fr_1.5fr_auto]">
            <input
              v-model="newCategoryName"
              type="text"
              placeholder="Tên ngành nghề (VD: Lập trình Mobile...)"
              class="rounded-xl border border-white/10 bg-slate-900 px-4 py-2.5 text-sm text-white placeholder-slate-400 focus:border-cyan-400 focus:outline-none"
            />
            <input
              v-model="newCategoryDesc"
              type="text"
              placeholder="Mô tả ngành nghề..."
              class="rounded-xl border border-white/10 bg-slate-900 px-4 py-2.5 text-sm text-white placeholder-slate-400 focus:border-cyan-400 focus:outline-none"
            />
            <button
              @click="addCategory"
              class="rounded-xl bg-cyan-400 px-5 py-2.5 text-sm font-extrabold text-slate-950 hover:bg-cyan-300 transition-colors"
            >
              + Thêm Ngành nghề
            </button>
          </div>
        </section>

        <!-- Categories List -->
        <section class="overflow-hidden rounded-2xl border border-white/10 bg-slate-950/70 shadow-xl ring-1 ring-cyan-400/10">
          <div v-if="categories.length === 0" class="p-8 text-center text-slate-400">
            Chưa có ngành nghề nào.
          </div>
          <div v-else class="overflow-x-auto">
            <table class="w-full text-left text-sm text-slate-200">
              <thead class="bg-slate-900/90 text-xs uppercase text-slate-400">
                <tr>
                  <th class="px-5 py-3.5 font-extrabold">ID</th>
                  <th class="px-5 py-3.5 font-extrabold">Tên Ngành nghề</th>
                  <th class="px-5 py-3.5 font-extrabold">Mô tả</th>
                  <th class="px-5 py-3.5 font-extrabold text-right">Thao tác</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-white/5">
                <tr v-for="cat in categories" :key="cat.id" class="hover:bg-white/[0.02]">
                  <td class="px-5 py-4 font-mono text-cyan-300 font-bold">#{{ cat.id }}</td>
                  <td class="px-5 py-4 font-extrabold text-white">{{ cat.name }}</td>
                  <td class="px-5 py-4 text-xs text-slate-300">{{ cat.description || 'N/A' }}</td>
                  <td class="px-5 py-4 text-right">
                    <button
                      @click="deleteCategory(cat.id)"
                      class="rounded-lg bg-rose-400/20 px-3 py-1.5 text-xs font-extrabold text-rose-300 hover:bg-rose-400/30 transition-colors"
                    >
                      Xóa
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </section>
      </div>

      <!-- Tab 2: Skills Management -->
      <div v-if="activeTab === 'skills'" class="space-y-6">
        <!-- Add Skill Form -->
        <section class="rounded-2xl border border-white/10 bg-slate-950/70 p-5 shadow-xl ring-1 ring-cyan-400/10">
          <h3 class="text-sm font-extrabold uppercase tracking-wide text-cyan-300 mb-4">Thêm Kỹ năng chuẩn hóa mới</h3>
          <div class="grid gap-4 md:grid-cols-[1fr_1fr_auto]">
            <input
              v-model="newSkillName"
              type="text"
              placeholder="Tên kỹ năng (VD: Golang, VueJS...)"
              class="rounded-xl border border-white/10 bg-slate-900 px-4 py-2.5 text-sm text-white placeholder-slate-400 focus:border-cyan-400 focus:outline-none"
            />
            <input
              v-model="newSkillCategory"
              type="text"
              placeholder="Phân nhóm (VD: Backend, Frontend...)"
              class="rounded-xl border border-white/10 bg-slate-900 px-4 py-2.5 text-sm text-white placeholder-slate-400 focus:border-cyan-400 focus:outline-none"
            />
            <button
              @click="addSkill"
              class="rounded-xl bg-cyan-400 px-5 py-2.5 text-sm font-extrabold text-slate-950 hover:bg-cyan-300 transition-colors"
            >
              + Thêm Kỹ năng
            </button>
          </div>
        </section>

        <!-- Skills List -->
        <section class="overflow-hidden rounded-2xl border border-white/10 bg-slate-950/70 shadow-xl ring-1 ring-cyan-400/10">
          <div v-if="skills.length === 0" class="p-8 text-center text-slate-400">
            Chưa có kỹ năng nào.
          </div>
          <div v-else class="overflow-x-auto">
            <table class="w-full text-left text-sm text-slate-200">
              <thead class="bg-slate-900/90 text-xs uppercase text-slate-400">
                <tr>
                  <th class="px-5 py-3.5 font-extrabold">ID</th>
                  <th class="px-5 py-3.5 font-extrabold">Tên Kỹ năng</th>
                  <th class="px-5 py-3.5 font-extrabold">Phân nhóm</th>
                  <th class="px-5 py-3.5 font-extrabold text-right">Thao tác</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-white/5">
                <tr v-for="sk in skills" :key="sk.id" class="hover:bg-white/[0.02]">
                  <td class="px-5 py-4 font-mono text-cyan-300 font-bold">#{{ sk.id }}</td>
                  <td class="px-5 py-4 font-extrabold text-white">{{ sk.name }}</td>
                  <td class="px-5 py-4 text-xs font-bold text-cyan-300">
                    <span class="rounded-full bg-cyan-400/10 px-2.5 py-1 ring-1 ring-cyan-400/20">{{ sk.category }}</span>
                  </td>
                  <td class="px-5 py-4 text-right">
                    <button
                      @click="deleteSkill(sk.id)"
                      class="rounded-lg bg-rose-400/20 px-3 py-1.5 text-xs font-extrabold text-rose-300 hover:bg-rose-400/30 transition-colors"
                    >
                      Xóa
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </section>
      </div>
    </div>
  </AdminShell>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import AdminShell from '~/components/admin/AdminShell.vue'
import { useApi } from '~/composables/useApi'
import { useToast } from '~/composables/useToast'

interface CategoryItem {
  id: number
  name: string
  description: string
}

interface SkillItem {
  id: number
  name: string
  category: string
}

const api = useApi()
const { success, error } = useToast()

const activeTab = ref<'categories' | 'skills'>('categories')

const categories = ref<CategoryItem[]>([])
const newCategoryName = ref('')
const newCategoryDesc = ref('')

const skills = ref<SkillItem[]>([])
const newSkillName = ref('')
const newSkillCategory = ref('')

const fetchCategories = async () => {
  try {
    const res = await api.get<CategoryItem[]>('/api/admin/categories')
    categories.value = res || []
  } catch (err: any) {
    error('Không thể tải danh sách ngành nghề')
  }
}

const addCategory = async () => {
  if (!newCategoryName.value.trim()) {
    error('Vui lòng nhập tên ngành nghề!')
    return
  }
  try {
    await api.post('/api/admin/categories', {
      name: newCategoryName.value,
      description: newCategoryDesc.value
    })
    newCategoryName.value = ''
    newCategoryDesc.value = ''
    success('Đã thêm ngành nghề mới thành công!')
    fetchCategories()
  } catch (err: any) {
    error('Không thể thêm ngành nghề')
  }
}

const deleteCategory = async (id: number) => {
  if (!confirm('Bạn có chắc chắn muốn xóa ngành nghề này?')) return
  try {
    await api.delete(`/api/admin/categories/${id}`)
    success('Đã xóa ngành nghề thành công!')
    fetchCategories()
  } catch (err: any) {
    error('Không thể xóa ngành nghề')
  }
}

const fetchSkills = async () => {
  try {
    const res = await api.get<SkillItem[]>('/api/admin/skills')
    skills.value = res || []
  } catch (err: any) {
    error('Không thể tải danh sách kỹ năng')
  }
}

const addSkill = async () => {
  if (!newSkillName.value.trim()) {
    error('Vui lòng nhập tên kỹ năng!')
    return
  }
  try {
    await api.post('/api/admin/skills', {
      name: newSkillName.value,
      category: newSkillCategory.value || 'General'
    })
    newSkillName.value = ''
    newSkillCategory.value = ''
    success('Đã thêm kỹ năng mới thành công!')
    fetchSkills()
  } catch (err: any) {
    error('Không thể thêm kỹ năng')
  }
}

const deleteSkill = async (id: number) => {
  if (!confirm('Bạn có chắc chắn muốn xóa kỹ năng này?')) return
  try {
    await api.delete(`/api/admin/skills/${id}`)
    success('Đã xóa kỹ năng thành công!')
    fetchSkills()
  } catch (err: any) {
    error('Không thể xóa kỹ năng')
  }
}

onMounted(() => {
  fetchCategories()
  fetchSkills()
})
</script>
