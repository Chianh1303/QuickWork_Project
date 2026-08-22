<template>
  <div class="w-full h-[500px] bg-slate-950 flex flex-col overflow-hidden rounded-2xl border border-white/10">
    
    <!-- Messages Scroll Area -->
    <div ref="chatScrollContainer" class="flex-1 p-5 overflow-y-auto space-y-3.5 bg-slate-950/90">
      <div v-if="chatMessages.length === 0" class="h-full flex flex-col items-center justify-center text-slate-500 text-xs font-medium space-y-2">
        <span class="text-3xl block opacity-40">💬</span>
        <p class="uppercase tracking-wider font-extrabold text-[11px]">Chưa có tin nhắn nào</p>
        <p class="text-slate-600 text-[10px]">Hãy nhập tin nhắn để mở đầu cuộc trao đổi với ứng viên.</p>
      </div>
      
      <div 
        v-for="msg in chatMessages" 
        :key="msg.id"
        :class="[msg.sender_id === currentUserId ? 'justify-end' : 'justify-start', 'flex items-end gap-2']"
      >
        <div 
          :class="[
            msg.sender_id === currentUserId 
              ? 'bg-gradient-to-r from-indigo-600 to-cyan-500 text-white font-semibold rounded-2xl rounded-br-none shadow-md shadow-cyan-500/10'
              : 'bg-slate-900 text-slate-100 border border-white/10 rounded-2xl rounded-bl-none',
            'max-w-xs sm:max-w-md px-4.5 py-3 text-xs leading-relaxed font-medium'
          ]"
        >
          <p class="whitespace-pre-line">{{ msg.message_text }}</p>
        </div>
      </div>
    </div>

    <!-- Message Input Bar -->
    <div class="p-4 bg-slate-900/90 border-t border-white/10">
      <form @submit.prevent="sendMessage" class="flex gap-3">
        <input 
          v-model="newMessageText"
          type="text" 
          placeholder="Nhập nội dung tin nhắn trò chuyện..."
          class="flex-1 px-4 py-3 border border-white/10 rounded-xl text-xs bg-slate-950 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 transition-all font-medium"
        />
        <button 
          type="submit" 
          :disabled="!newMessageText.trim()"
          class="px-6 py-2.5 bg-gradient-to-r from-cyan-400 to-emerald-400 hover:brightness-110 disabled:opacity-40 text-slate-950 text-xs font-black uppercase tracking-wider rounded-xl transition-all shadow-md shadow-cyan-500/25 cursor-pointer"
        >
          Gửi
        </button>
      </form>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, nextTick, watch } from "vue"
import { useApi } from "~/composables/useApi"

const props = defineProps<{
    applicationId: number
    targetId: number
    currentUserId: number
}>()

const api = useApi()
const chatMessages = ref<any[]>([])
const newMessageText = ref("")
const chatScrollContainer = ref<HTMLElement | null>(null)
let socket: WebSocket | null = null

const scrollToBottom = async () => {
  await nextTick()
  if (chatScrollContainer.value) {
    chatScrollContainer.value.scrollTop = chatScrollContainer.value.scrollHeight
  }
}

const fetchHistory = async () => {
  try {
    const result: any = await api.get(
      "/api/chat/history",
      {
        params: {
          application_id: props.applicationId
        }
      }
    )
    chatMessages.value = result || []
    scrollToBottom()
  } catch (err) {
    console.error(err)
  }
}

const connectSocket = () => {
  if (
    socket &&
    (
      socket.readyState === WebSocket.OPEN ||
      socket.readyState === WebSocket.CONNECTING
    )
  ) {
    return
  }

  socket = new WebSocket(
    `ws://localhost:3000/api/chat/ws?userId=${props.currentUserId}`
  )

  socket.onopen = () => {
    console.log("WebSocket connected")
  }

  socket.onmessage = (event) => {
    const message = JSON.parse(event.data)
    if (message.application_id !== props.applicationId) {
      return
    }
    chatMessages.value.push(message)
    scrollToBottom()
  }

  socket.onclose = () => {
    socket = null
  }
}

const sendMessage = () => {
  if (!socket || socket.readyState !== WebSocket.OPEN) return
  const text = newMessageText.value.trim()
  if (text === "") return

  socket.send(
    JSON.stringify({
      application_id: props.applicationId,
      receiver_id: props.targetId,
      message_text: text
    })
  )

  newMessageText.value = ""
}

watch(
  () => props.currentUserId,
  (id) => {
    if (!id || socket) return
    connectSocket()
  },
  { immediate: true }
)

onMounted(() => {
  fetchHistory()
})

onUnmounted(() => {
  if (socket) {
    socket.close()
  }
})
</script>
