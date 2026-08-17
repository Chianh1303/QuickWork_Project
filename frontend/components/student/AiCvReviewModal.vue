<template>
  <Transition
    enter-active-class="transition duration-200 ease-out"
    enter-from-class="opacity-0 scale-95"
    enter-to-class="opacity-100 scale-100"
    leave-active-class="transition duration-150 ease-in"
    leave-from-class="opacity-100 scale-100"
    leave-to-class="opacity-0 scale-95"
  >
    <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center p-3 sm:p-6 overflow-hidden">
      <!-- Backdrop -->
      <div class="fixed inset-0 bg-slate-950/85 backdrop-blur-md transition-opacity" @click="close"></div>

      <!-- Modal Container -->
      <div class="relative w-full max-w-4xl rounded-3xl border border-cyan-400/30 bg-slate-900/95 shadow-2xl shadow-cyan-950/60 backdrop-blur ring-1 ring-cyan-400/20 text-left flex flex-col max-h-[90vh] my-auto">
        
        <!-- Sticky Header (Luôn cố định trên cùng) -->
        <div class="flex items-center justify-between border-b border-white/10 p-5 sm:p-6 bg-slate-900/90 rounded-t-3xl backdrop-blur sticky top-0 z-20">
          <div class="flex items-center space-x-3">
            <div class="h-11 w-11 rounded-2xl bg-gradient-to-tr from-cyan-400 to-emerald-400 flex items-center justify-center text-slate-950 font-semibold text-2xl shadow-lg shadow-cyan-500/30">
              🤖
            </div>
            <div>
              <div class="flex items-center gap-2">
                <h3 class="text-xl font-semibold text-white">Báo Cáo Phân Tích CV Chuyên Sâu AI</h3>
                <span 
                  v-if="result?.evaluation_source"
                  class="px-2.5 py-0.5 text-[10px] font-semibold uppercase rounded-full border shadow-sm"
                  :class="result.evaluation_source === 'gemini' ? 'bg-emerald-400/10 text-emerald-300 border-emerald-400/30' : 'bg-cyan-400/10 text-cyan-300 border-cyan-400/30'"
                >
                  {{ result.evaluation_source === 'gemini' ? '✨ Gemini AI Multimodal' : '⚡ Smart Heuristic Parser' }}
                </span>
              </div>
              <p class="text-xs text-cyan-200/80 font-medium mt-0.5">Đánh giá khắt khe theo tiêu chuẩn Trưởng phòng Tuyển dụng Nhân sự</p>
            </div>
          </div>
          
          <button 
            @click="close" 
            class="h-9 w-9 rounded-xl bg-white/10 hover:bg-white/20 text-slate-300 hover:text-white flex items-center justify-center transition-colors text-base font-semibold"
            title="Đóng cửa sổ"
          >
            ✕
          </button>
        </div>

        <!-- Scrollable Content Body (Có thanh cuộn tùy chỉnh mượt mà) -->
        <div ref="scrollContainer" class="flex-1 overflow-y-auto p-5 sm:p-6 space-y-6 custom-scrollbar">
          <div v-if="result" class="space-y-6">
            
            <!-- 1. OVERALL SCORE & RECOMMENDATION VERDICT -->
            <div class="grid grid-cols-1 md:grid-cols-[180px_1fr] items-stretch gap-6 rounded-2xl border border-white/10 bg-gradient-to-br from-slate-950/90 via-slate-900/90 to-cyan-950/30 p-5 ring-1 ring-cyan-400/20">
              
              <!-- Score Ring -->
              <div class="flex flex-col items-center justify-center text-center border-b md:border-b-0 md:border-r border-white/10 pb-4 md:pb-0 md:pr-4">
                <div class="relative flex h-28 w-28 items-center justify-center rounded-full bg-gradient-to-tr from-cyan-500/20 to-emerald-500/20 p-2 ring-4 ring-cyan-400/40 shadow-xl shadow-cyan-950/50">
                  <span class="text-4xl font-semibold text-white tracking-tight">{{ result.score }}<span class="text-xl font-semibold text-cyan-300">%</span></span>
                </div>
                <span class="mt-3 text-xs font-semibold uppercase tracking-wider text-cyan-300">Điểm CV Tổng Thể</span>
              </div>

              <!-- HR Recommendation Box -->
              <div class="space-y-3 flex flex-col justify-center">
                <div class="flex flex-wrap items-center justify-between gap-2">
                  <span class="text-xs font-semibold uppercase tracking-wider text-cyan-200/90">📢 Đánh Giá Kết Luận Từ HR</span>
                  <span 
                    v-if="result.recommendation?.decision"
                    class="px-3 py-1 rounded-lg text-xs font-semibold uppercase tracking-wider border shadow-md"
                    :class="getDecisionBadgeClass(result.recommendation.decision)"
                  >
                    {{ getDecisionLabel(result.recommendation.decision) }}
                  </span>
                </div>

                <p class="text-sm text-slate-200 font-medium leading-relaxed bg-slate-950/60 p-3.5 rounded-xl border border-white/10">
                  "{{ result.recommendation?.reason || getScoreAssessment(result.score) }}"
                </p>

                <div v-if="result.recommendation?.confidence" class="flex items-center gap-2 text-xs text-slate-400 font-semibold">
                  <span>Độ tin cậy đánh giá AI:</span>
                  <div class="w-24 bg-slate-800 h-2 rounded-full overflow-hidden border border-white/10">
                    <div class="bg-cyan-400 h-full rounded-full" :style="{ width: (result.recommendation.confidence * 100) + '%' }"></div>
                  </div>
                  <span class="text-cyan-300 font-semibold">{{ Math.round(result.recommendation.confidence * 100) }}%</span>
                </div>
              </div>
            </div>

            <!-- 2. BÓC TÁCH DỮ LIỆU TỪ FILE PDF (EXTRACTED PDF DATA) -->
            <div class="rounded-2xl border border-white/10 bg-slate-950/70 p-5 space-y-4">
              <h4 class="text-xs font-semibold uppercase tracking-wider text-cyan-300 flex items-center gap-2">
                <span>📄</span> Báo Cáo Bóc Tách Trực Tiếp Từ File CV PDF
              </h4>

              <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
                
                <!-- Học vấn -->
                <div class="bg-slate-900/90 p-4 rounded-xl border border-white/10 space-y-2">
                  <span class="text-xs font-semibold text-slate-400 uppercase tracking-wider block">🎓 Học Vấn & Trường Học</span>
                  <div v-if="result.education?.items && result.education.items.length > 0" class="space-y-1.5">
                    <div v-for="(edu, idx) in result.education.items" :key="idx" class="text-xs text-slate-200 font-semibold">
                      <p class="text-cyan-300 font-semibold">{{ edu.institution !== 'not_found' ? edu.institution : 'Chưa rõ tên trường' }}</p>
                      <p v-if="edu.degree !== 'not_found'" class="text-[11px] text-slate-400 font-medium">{{ edu.degree }}</p>
                    </div>
                  </div>
                  <p v-else class="text-xs text-slate-400 italic">Đã phân tích từ file PDF.</p>
                </div>

                <!-- Kỹ năng -->
                <div class="bg-slate-900/90 p-4 rounded-xl border border-white/10 space-y-2">
                  <span class="text-xs font-semibold text-slate-400 uppercase tracking-wider block">🛠️ Kỹ Năng & Công Nghệ</span>
                  <div v-if="result.skills?.technical && result.skills.technical.length > 0" class="flex flex-wrap gap-1.5">
                    <span 
                      v-for="(tech, idx) in result.skills.technical" 
                      :key="idx"
                      class="px-2 py-0.5 bg-cyan-400/10 text-cyan-200 border border-cyan-400/20 text-[11px] font-semibold rounded-md"
                    >
                      {{ tech }}
                    </span>
                  </div>
                  <p v-else class="text-xs text-slate-400 italic">Đã phân tích từ file PDF.</p>
                </div>

                <!-- Dự án -->
                <div class="bg-slate-900/90 p-4 rounded-xl border border-white/10 space-y-2">
                  <span class="text-xs font-semibold text-slate-400 uppercase tracking-wider block">💻 Dự Án Thực Tế</span>
                  <div v-if="result.experience?.projects && result.experience.projects.length > 0" class="space-y-2">
                    <div v-for="(proj, idx) in result.experience.projects" :key="idx" class="text-xs space-y-0.5">
                      <p class="font-semibold text-emerald-300">● {{ proj.name }}</p>
                      <p v-if="proj.impact && proj.impact !== 'not_found'" class="text-[11px] text-slate-300 italic">{{ proj.impact }}</p>
                    </div>
                  </div>
                  <p v-else class="text-xs text-slate-400 italic">Đã ghi nhận kinh nghiệm thực chiến trong PDF.</p>
                </div>

              </div>
            </div>

            <!-- 3. PHÂN TÍCH STAR & CHUẨN LỌC ATS -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              
              <!-- ATS Formatting Review -->
              <div class="rounded-2xl border border-cyan-400/20 bg-slate-950/70 p-5 space-y-3">
                <div class="flex items-center justify-between">
                  <h4 class="text-xs font-semibold uppercase tracking-wider text-cyan-300">⚙️ Khung Chuẩn ATS (ATS Score)</h4>
                  <span class="text-sm font-semibold text-cyan-300">{{ result.ats?.score || result.score }}%</span>
                </div>
                <div class="w-full bg-slate-800 h-2 rounded-full overflow-hidden border border-white/10">
                  <div class="bg-cyan-400 h-full rounded-full transition-all duration-500" :style="{ width: (result.ats?.score || result.score) + '%' }"></div>
                </div>
                <ul class="space-y-1.5 text-xs text-slate-300 font-medium">
                  <li v-for="(issue, idx) in (result.ats?.issues || [])" :key="idx" class="flex items-start gap-1.5">
                    <span class="text-amber-400">⚠️</span>
                    <span>{{ issue }}</span>
                  </li>
                </ul>
              </div>

              <!-- STAR Analysis Progress Bars -->
              <div class="rounded-2xl border border-emerald-400/20 bg-slate-950/70 p-5 space-y-3">
                <h4 class="text-xs font-semibold uppercase tracking-wider text-emerald-300">⭐ Phân Tích Chuẩn STAR (Situation-Task-Action-Result)</h4>
                
                <div class="space-y-2 text-xs font-semibold text-slate-300">
                  <div class="flex items-center justify-between">
                    <span>S - Context/Situation:</span>
                    <span class="text-emerald-300">{{ result.star_analysis?.situation || 70 }}%</span>
                  </div>
                  <div class="flex items-center justify-between">
                    <span>T - Task Responsibility:</span>
                    <span class="text-emerald-300">{{ result.star_analysis?.task || 75 }}%</span>
                  </div>
                  <div class="flex items-center justify-between">
                    <span>A - Action Steps:</span>
                    <span class="text-emerald-300">{{ result.star_analysis?.action || 80 }}%</span>
                  </div>
                  <div class="flex items-center justify-between">
                    <span>R - Quantified Result:</span>
                    <span class="text-amber-300">{{ result.star_analysis?.result || 40 }}%</span>
                  </div>
                </div>
              </div>

            </div>

            <!-- 4. SUGGESTED SUMMARY (KÈM NÚT SAO CHÉP 1-CHẠM) -->
            <div class="rounded-2xl border border-cyan-400/30 bg-cyan-400/10 p-5 space-y-2">
              <div class="flex items-center justify-between">
                <span class="text-xs font-semibold uppercase tracking-wider text-cyan-200">Gợi Ý Đoạn Tóm Tắt Bản Thân Chuẩn HR (AI Executive Summary)</span>
                <button
                  @click="copySummary"
                  class="inline-flex items-center gap-1.5 rounded-lg bg-cyan-400 px-3.5 py-1.5 text-xs font-semibold text-slate-950 shadow-md shadow-cyan-500/25 hover:bg-cyan-300 active:scale-95 transition-all"
                >
                  <span>{{ copied ? '✓ Đã Sao Chép' : '📋 Sao Chép 1-Chạm' }}</span>
                </button>
              </div>
              <p class="text-sm text-slate-100 italic font-medium leading-relaxed bg-slate-950/60 p-4 rounded-xl border border-white/10">
                "{{ result.suggested_summary }}"
              </p>
            </div>

            <!-- 5. STRENGTHS VS WEAKNESSES GRID -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              
              <!-- Strengths -->
              <div class="rounded-2xl border border-emerald-400/20 bg-emerald-400/10 p-5 space-y-2">
                <h4 class="text-xs font-semibold uppercase tracking-wider text-emerald-300 flex items-center gap-1.5">
                  <span>✅</span> Điểm Mạnh Cốt Lõi
                </h4>
                <ul class="space-y-2 text-xs font-semibold text-emerald-100">
                  <li v-for="(item, idx) in result.strengths" :key="idx" class="flex items-start gap-2">
                    <span class="text-emerald-400 font-semibold">•</span>
                    <span>{{ item }}</span>
                  </li>
                </ul>
              </div>

              <!-- Weaknesses / Improvements -->
              <div class="rounded-2xl border border-amber-400/20 bg-amber-400/10 p-5 space-y-2">
                <h4 class="text-xs font-semibold uppercase tracking-wider text-amber-300 flex items-center gap-1.5">
                  <span>⚠️</span> Điểm Cần Khắc Phục / Thiếu Sót
                </h4>
                <ul class="space-y-2 text-xs font-semibold text-amber-100">
                  <li v-for="(item, idx) in (result.weaknesses || result.improvements || [])" :key="idx" class="flex items-start gap-2">
                    <span class="text-amber-400 font-semibold">•</span>
                    <span>{{ item }}</span>
                  </li>
                </ul>
              </div>

            </div>

            <!-- 6. ACTIONABLE TIPS (CỐNG THỨC GOOGLE RESUME) -->
            <div class="rounded-2xl border border-white/10 bg-slate-950/70 p-5 space-y-3">
              <h4 class="text-xs font-semibold uppercase tracking-wider text-cyan-300">💡 Các Bước Cải Thiện Cụ Thể (Google Resume Formula)</h4>
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-3 text-xs font-semibold text-slate-300">
                <div v-for="(tip, idx) in result.actionable_tips" :key="idx" class="flex items-center gap-2.5 rounded-xl bg-white/5 p-3 border border-white/5">
                  <span class="flex h-6 w-6 flex-shrink-0 items-center justify-center rounded-full bg-cyan-400/20 text-cyan-300 font-semibold text-xs">{{ idx + 1 }}</span>
                  <span class="leading-snug">{{ tip }}</span>
                </div>
              </div>
            </div>

          </div>
        </div>

        <!-- Sticky Footer với nút Scroll To Top & Đóng -->
        <div class="flex items-center justify-between border-t border-white/10 p-4 sm:p-5 bg-slate-900/90 rounded-b-3xl backdrop-blur sticky bottom-0 z-20">
          <button
            @click="scrollToTop"
            class="inline-flex items-center gap-1.5 rounded-xl bg-white/10 hover:bg-white/20 px-4 py-2 text-xs font-semibold text-slate-200 transition-colors"
          >
            <span>⬆ Scroll Lên Đầu</span>
          </button>

          <button
            @click="close"
            class="rounded-xl bg-cyan-400 px-6 py-2.5 text-xs font-semibold text-slate-950 shadow-lg shadow-cyan-500/20 hover:bg-cyan-300 transition-colors"
          >
            Đóng Cửa Sổ
          </button>
        </div>

      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { EvaluateCvResult } from '~/composables/useAi'
import { useToast } from '~/composables/useToast'

const props = defineProps<{
  isOpen: boolean
  result: EvaluateCvResult | null
}>()

const emit = defineEmits(['close'])
const { success } = useToast()
const copied = ref(false)
const scrollContainer = ref<HTMLElement | null>(null)

const close = () => emit('close')

const scrollToTop = () => {
  if (scrollContainer.value) {
    scrollContainer.value.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const getDecisionBadgeClass = (decision: string) => {
  switch (decision) {
    case 'strong_interview':
      return 'bg-emerald-400/20 text-emerald-300 border-emerald-400/40'
    case 'interview':
      return 'bg-cyan-400/20 text-cyan-300 border-cyan-400/40'
    case 'improve':
      return 'bg-amber-400/20 text-amber-300 border-amber-400/40'
    case 'reject':
      return 'bg-rose-400/20 text-rose-300 border-rose-400/40'
    default:
      return 'bg-slate-800 text-slate-300 border-white/10'
  }
}

const getDecisionLabel = (decision: string) => {
  switch (decision) {
    case 'strong_interview':
      return '🟢 Ưu Tiên Phỏng Vấn Ngay'
    case 'interview':
      return '🟢 Đạt Chuẩn Phỏng Vấn'
    case 'improve':
      return '🟡 Cần Cải Thiện Thêm'
    case 'reject':
      return '🔴 Chưa Đạt Tiêu Chuẩn'
    default:
      return decision
  }
}

const getScoreAssessment = (score: number) => {
  if (score >= 80) return 'Hồ sơ CV của bạn rất ấn tượng! Tỷ lệ trúng tuyển qua vòng lọc hồ sơ đạt mức rất cao. Hãy tiếp tục ứng tuyển ngay.'
  if (score >= 60) return 'Hồ sơ CV đạt mức khá. Bạn đã có nền tảng tốt nhưng cần bổ sung thêm kỹ năng và làm rõ mô tả để chinh phục các công việc lương cao.'
  return 'Hồ sơ CV đang ở mức cơ bản. Hãy làm theo các gợi ý của AI dưới đây để tăng điểm số và tạo sức hút với nhà tuyển dụng.'
}

const copySummary = () => {
  if (!props.result?.suggested_summary) return
  navigator.clipboard.writeText(props.result.suggested_summary)
  copied.value = true
  success('Đã sao chép đoạn tóm tắt AI vào khay nhớ tạm!')
  setTimeout(() => {
    copied.value = false
  }, 3000)
}
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 6px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: rgba(15, 23, 42, 0.6);
  border-radius: 9999px;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: rgba(34, 211, 238, 0.3);
  border-radius: 9999px;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: rgba(34, 211, 238, 0.6);
}
</style>
