<template>
  <AdminShell>
    <div class="space-y-6">
      <!-- Header Banner -->
      <section class="rounded-3xl border border-white/10 bg-slate-950/80 p-6 shadow-2xl shadow-slate-950/40 ring-1 ring-cyan-400/10 lg:p-8">
        <div class="flex flex-col gap-6">
          <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
            <div>
              <h1 class="text-3xl font-extrabold tracking-tight text-white sm:text-4xl">
                Xử lý Khiếu nại & Tranh chấp
              </h1>
              <p class="mt-2 text-sm font-semibold text-slate-300">
                Xem chi tiết các ticket khiếu nại giữa Sinh viên & Doanh nghiệp và đưa ra phán quyết giải quyết chính thức.
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

      <!-- Filter Bar -->
      <section class="rounded-2xl border border-white/10 bg-slate-950/70 p-5 shadow-xl shadow-slate-950/25 ring-1 ring-cyan-400/10">
        <div class="flex items-center justify-between">
          <h2 class="text-base font-extrabold text-white">Danh sách Ticket khiếu nại</h2>

          <select
            v-model="statusFilter"
            class="rounded-xl border border-white/10 bg-slate-900/90 px-3.5 py-2 text-sm text-white focus:border-cyan-400 focus:outline-none"
            @change="fetchTickets"
          >
            <option value="">Tất cả trạng thái</option>
            <option value="pending">🟡 Đang chờ xử lý (Pending)</option>
            <option value="resolved">🟢 Đã ra phán quyết (Resolved)</option>
            <option value="rejected">🔴 Từ chối khiếu nại (Rejected)</option>
          </select>
        </div>
      </section>

      <!-- Tickets Data Table -->
      <section class="overflow-hidden rounded-2xl border border-white/10 bg-slate-950/70 shadow-xl shadow-slate-950/25 ring-1 ring-cyan-400/10">
        <div v-if="isLoading" class="p-8 text-center text-slate-400">
          <span class="animate-pulse text-sm font-bold">Đang tải danh sách khiếu nại...</span>
        </div>

        <div v-else-if="tickets.length === 0" class="p-8 text-center text-slate-400">
          <p class="text-base font-semibold">Chưa có ticket khiếu nại nào.</p>
        </div>

        <div v-else class="overflow-x-auto">
          <table class="w-full text-left text-sm text-slate-200">
            <thead class="bg-slate-900/90 text-xs uppercase text-slate-400">
              <tr>
                <th class="px-5 py-3.5 font-extrabold">Mã Ticket</th>
                <th class="px-5 py-3.5 font-extrabold">Người gửi khiếu nại</th>
                <th class="px-5 py-3.5 font-extrabold">Đối tượng bị khiếu nại</th>
                <th class="px-5 py-3.5 font-extrabold">Lý do khiếu nại</th>
                <th class="px-5 py-3.5 font-extrabold">Trạng thái</th>
                <th class="px-5 py-3.5 font-extrabold text-right">Thao tác</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-white/5">
              <tr v-for="t in tickets" :key="t.ticket_id" class="hover:bg-white/[0.02]">
                <td class="px-5 py-4 font-mono font-extrabold text-cyan-300">
                  #TK-{{ t.ticket_id }}
                </td>

                <td class="px-5 py-4">
                  <p class="font-extrabold text-white">{{ t.reporter_email }}</p>
                  <span class="text-xs uppercase font-bold text-cyan-400/80">({{ t.reporter_role }})</span>
                </td>

                <td class="px-5 py-4">
                  <p class="font-extrabold text-white">{{ t.target_email }}</p>
                </td>

                <td class="px-5 py-4 max-w-xs truncate font-semibold text-slate-200">
                  {{ t.reason }}
                </td>

                <td class="px-5 py-4">
                  <span
                    class="inline-flex items-center gap-1.5 rounded-full px-2.5 py-1 text-xs font-bold border"
                    :class="{
                      'bg-amber-400/10 text-amber-300 border-amber-400/30': t.status === 'pending' && !t.is_reappealed,
                      'bg-rose-500/20 text-rose-300 border-rose-500/40 animate-pulse': t.status === 'pending' && t.is_reappealed,
                      'bg-emerald-400/10 text-emerald-300 border-emerald-400/30': t.status === 'resolved',
                      'bg-rose-400/10 text-rose-300 border-rose-400/30': t.status === 'rejected'
                    }"
                  >
                    <span>{{ t.is_reappealed && t.status === 'pending' ? '⚠️ Tái khiếu nại' : getStatusLabel(t.status) }}</span>
                  </span>
                </td>

                <td class="px-5 py-4 text-right">
                  <button
                    @click="openResolveModal(t)"
                    class="rounded-lg bg-cyan-400/20 border border-cyan-400/30 px-3.5 py-1.5 text-xs font-extrabold text-cyan-200 hover:bg-cyan-400 hover:text-slate-950 transition-colors cursor-pointer"
                  >
                    {{ t.status === 'pending' ? (t.is_reappealed ? '⚠️ Xử lý Tái khiếu nại' : '⚖️ Ra phán quyết') : 'Xem chi tiết' }}
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </section>
    </div>

    <!-- Ticket Resolve Modal -->
    <div v-if="selectedTicket" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-950/80 backdrop-blur-sm">
      <div class="w-full max-w-xl rounded-2xl border border-white/10 bg-slate-900 p-6 shadow-2xl space-y-4 max-h-[90vh] overflow-y-auto">
        <div class="flex items-center justify-between border-b border-white/10 pb-3">
          <h3 class="text-lg font-extrabold text-white">Xử lý Ticket Khiếu nại #TK-{{ selectedTicket.ticket_id }}</h3>
          <button @click="selectedTicket = null" class="text-slate-400 hover:text-white text-lg font-bold">✕</button>
        </div>

        <div class="space-y-3 text-sm text-slate-300">
          <div class="grid grid-cols-2 gap-3 bg-slate-950/60 p-3.5 rounded-xl border border-white/5">
            <div>
              <span class="text-xs text-slate-400 block font-bold">NGƯỜI GỬI</span>
              <span class="font-semibold text-white">{{ selectedTicket.reporter_email }}</span>
            </div>
            <div>
              <span class="text-xs text-slate-400 block font-bold">ĐỐI TƯỢNG BỊ KHIẾU NẠI</span>
              <span class="font-semibold text-white">{{ selectedTicket.target_email }}</span>
            </div>
          </div>

          <div>
            <span class="text-xs text-slate-400 block font-bold mb-1">LÝ DO KHIẾU NẠI BAN ĐẦU</span>
            <p class="font-bold text-white text-base">{{ selectedTicket.reason }}</p>
          </div>

          <div>
            <span class="text-xs text-slate-400 block font-bold mb-1">MÔ TẢ CHI TIẾT TRANH CHẤP</span>
            <p class="bg-slate-950/60 p-3 rounded-xl border border-white/5 text-xs text-slate-200 leading-relaxed">
              {{ selectedTicket.description }}
            </p>
          </div>

          <div v-if="selectedTicket.requested_action">
            <span class="text-xs text-cyan-300 block font-bold mb-1">ĐỀ XUẤT MONG MUỐN XỬ LÝ</span>
            <p class="bg-cyan-950/40 p-3 rounded-xl border border-cyan-500/20 text-xs text-cyan-100 font-medium">
              {{ selectedTicket.requested_action }}
            </p>
          </div>

          <div v-if="selectedTicket.evidence_url">
            <span class="text-xs text-cyan-300 block font-bold mb-1">BẰNG CHỨNG / MINH CHỨNG ĐÍNH KÈM</span>
            <a
              :href="selectedTicket.evidence_url"
              target="_blank"
              rel="noopener noreferrer"
              class="inline-flex items-center gap-2 bg-cyan-950/40 hover:bg-cyan-900/60 px-3.5 py-2.5 rounded-xl border border-cyan-500/30 text-xs font-bold text-cyan-200 transition-colors"
            >
              🔗 Xem Bằng Chứng Đính Kèm (Link ngoài)
            </a>
          </div>

          <!-- Section Yêu cầu Tái xem xét nếu có -->
          <div v-if="selectedTicket.is_reappealed" class="rounded-xl border border-amber-500/40 bg-amber-950/40 p-4 space-y-2">
            <div class="flex items-center gap-2 font-extrabold text-amber-300 text-xs uppercase tracking-wider">
              <span>⚠️ YÊU CẦU TÁI XEM XÉT PHÁN QUYẾT TỪ NGƯỜI DÙNG</span>
            </div>
            <p class="text-xs font-semibold text-amber-100 leading-relaxed">
              <strong>Lý do phản hồi:</strong> {{ selectedTicket.reappeal_reason }}
            </p>
            <div v-if="selectedTicket.reappeal_evidence">
              <a :href="selectedTicket.reappeal_evidence" target="_blank" class="text-xs font-bold text-cyan-300 underline">
                🔗 Xem Bằng Chứng Bổ Sung Tái Khiếu Nại
              </a>
            </div>
          </div>

          <!-- Previous Verdict Display if already resolved once -->
          <div v-if="selectedTicket.verdict && selectedTicket.status === 'pending'" class="rounded-xl border border-slate-700 bg-slate-950/80 p-3">
            <span class="text-xs font-bold text-slate-400">Phán quyết cũ trước đó:</span>
            <p class="mt-1 text-xs font-medium text-slate-300 italic">{{ selectedTicket.verdict }}</p>
          </div>

          <!-- Readonly Verdict display if resolved/rejected -->
          <div v-if="selectedTicket.status !== 'pending' && selectedTicket.verdict" class="rounded-xl border border-emerald-400/30 bg-emerald-400/10 p-4">
            <span class="text-xs font-extrabold uppercase text-emerald-300">Phán quyết chính thức đã ban hành</span>
            <p class="mt-1 text-sm font-semibold text-white">{{ selectedTicket.verdict }}</p>
          </div>

          <!-- Form ban hành phán quyết if pending -->
          <div v-if="selectedTicket.status === 'pending'" class="space-y-3 pt-2 border-t border-white/10">
            <div>
              <label class="block text-xs font-bold text-slate-300 mb-1">QUYẾT ĐỊNH XỬ LÝ (STATUS):</label>
              <select
                v-model="verdictStatusInput"
                class="w-full rounded-xl border border-white/10 bg-slate-950 p-2.5 text-xs text-white focus:border-cyan-400 focus:outline-none"
              >
                <option value="resolved">🟢 Duyệt khiếu nại (Resolved)</option>
                <option value="rejected">🔴 Bác bỏ khiếu nại (Rejected)</option>
              </select>
            </div>

            <div>
              <label class="block text-xs font-bold text-slate-300 mb-1">NHẬP PHÁN QUYẾT XỬ LÝ CỦA ADMIN:</label>
              <textarea
                v-model="verdictInput"
                rows="3"
                placeholder="Nhập nội dung phán quyết giải quyết tranh chấp (Ví dụ: Chấp nhận giải ngân đền bù cho sinh viên...)"
                class="w-full rounded-xl border border-white/10 bg-slate-950 p-3 text-xs text-white placeholder-slate-500 focus:border-cyan-400 focus:outline-none"
              />
            </div>
          </div>
        </div>

        <div class="pt-4 border-t border-white/10 flex justify-end gap-3">
          <button @click="selectedTicket = null" class="rounded-xl border border-white/10 px-4 py-2 text-xs font-bold text-slate-300 hover:bg-white/10">
            Đóng
          </button>

          <button
            v-if="selectedTicket.status === 'pending'"
            @click="submitVerdict"
            class="rounded-xl bg-cyan-400 px-5 py-2 text-xs font-extrabold text-slate-950 hover:bg-cyan-300 transition-colors cursor-pointer"
          >
            ⚖️ Ban hành Phán quyết
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

interface TicketItem {
  ticket_id: number
  application_id: number
  reporter_id: number
  reporter_email: string
  reporter_role: string
  target_id: number
  target_email: string
  reason: string
  description: string
  requested_action?: string
  evidence_url?: string
  status: string
  verdict: string
  is_reappealed?: boolean
  reappeal_reason?: string
  reappeal_evidence?: string
  reappealed_at?: string
  resolved_at?: string
  created_at: string
}

const api = useApi()
const { success, error } = useToast()

const tickets = ref<TicketItem[]>([])
const isLoading = ref(false)
const statusFilter = ref('')
const selectedTicket = ref<TicketItem | null>(null)
const verdictInput = ref('')
const verdictStatusInput = ref<'resolved' | 'rejected'>('resolved')

const fetchTickets = async () => {
  isLoading.value = true
  try {
    const params = new URLSearchParams()
    if (statusFilter.value) params.append('status', statusFilter.value)

    const res = await api.get<{ items: TicketItem[] }>(`/api/admin/tickets?${params.toString()}`)
    tickets.value = res.items || []
  } catch (err: any) {
    error('Không thể tải danh sách ticket khiếu nại')
  } finally {
    isLoading.value = false
  }
}

const openResolveModal = (t: TicketItem) => {
  selectedTicket.value = t
  verdictInput.value = t.verdict || ''
  verdictStatusInput.value = 'resolved'
}

const submitVerdict = async () => {
  if (!selectedTicket.value) return
  if (!verdictInput.value.trim() || verdictInput.value.trim().length < 5) {
    error('Phán quyết Admin phải có ít nhất 5 ký tự!')
    return
  }

  try {
    await api.put(`/api/admin/tickets/${selectedTicket.value.ticket_id}/resolve`, {
      verdict: verdictInput.value,
      status: verdictStatusInput.value
    })
    selectedTicket.value.verdict = verdictInput.value
    selectedTicket.value.status = verdictStatusInput.value
    success('Đã ban hành phán quyết xử lý khiếu nại thành công!')
    selectedTicket.value = null
    fetchTickets()
  } catch (err: any) {
    error('Không thể gửi phán quyết xử lý: ' + (err.response?._data?.message || err.message))
  }
}

const getStatusLabel = (status: string) => {
  switch (status) {
    case 'pending': return '🟡 Chờ xử lý'
    case 'resolved': return '🟢 Đã giải quyết'
    case 'rejected': return '🔴 Từ chối khiếu nại'
    default: return status
  }
}

onMounted(() => {
  fetchTickets()
})
</script>
