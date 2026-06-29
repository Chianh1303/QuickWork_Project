<template>
  <div class="w-full h-[450px] bg-slate-900 border border-slate-800 rounded-2xl flex flex-col overflow-hidden shadow-inner">
    
    <!-- Danh sách tin nhắn -->
    <div ref="chatScrollContainer" class="flex-1 p-4 overflow-y-auto space-y-3 bg-slate-950/40">
      <div v-if="chatMessages.length === 0" class="h-full flex items-center justify-center text-slate-600 text-xs font-medium uppercase tracking-wider">
        Chưa có tin nhắn nào. Hãy bắt đầu cuộc trò chuyện!
      </div>
      
      <div 
        v-for="(msg, index) in chatMessages" 
       :key="msg.id"
        :class="[msg.sender_id === currentUserId ? 'justify-end' : 'justify-start', 'flex items-end']"
      >
        <div 
          :class="[
            msg.sender_id === currentUserId 
              ? 'bg-blue-600 text-white rounded-br-none' 
              : 'bg-slate-900 text-slate-200 border border-slate-800 rounded-bl-none',
            'max-w-xs md:max-w-md px-4 py-2.5 rounded-2xl text-xs font-medium leading-relaxed shadow-sm'
          ]"
        >
          <p>{{ msg.message_text }}</p>
        </div>
      </div>
    </div>

    <!-- Thanh nhập dữ liệu -->
    <div class="p-4 bg-slate-950 border-t border-slate-800/60">
      <form @submit.prevent="sendMessage" class="flex gap-3">
        <input 
          v-model="newMessageText"
          type="text" 
          placeholder="Nhập nội dung tin nhắn trò chuyện..."
          class="flex-1 px-4 py-3 border border-slate-800 rounded-xl text-xs bg-slate-900 text-slate-200 placeholder-slate-600 focus:outline-none focus:border-blue-500 transition-all font-medium"
        />
        <button 
          type="submit" 
          class="px-5 py-2.5 bg-blue-600 hover:bg-blue-500 text-white text-xs font-bold uppercase tracking-wider rounded-xl transition-all shadow-lg shadow-blue-600/10"
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
    applicationId:number
    targetId:number
    currentUserId:number
}>()

const api = useApi()

const chatMessages = ref<any[]>([])

const newMessageText = ref("")

const chatScrollContainer = ref<HTMLElement | null>(null)

let socket:WebSocket | null = null

//----------------------------------------------------
// Scroll
//----------------------------------------------------

const scrollToBottom = async()=>{

    await nextTick()

    if(chatScrollContainer.value){

        chatScrollContainer.value.scrollTop=
        chatScrollContainer.value.scrollHeight

    }

}

//----------------------------------------------------
// Load History
//----------------------------------------------------

const fetchHistory = async()=>{

    try{

        const result:any = await api.get(
            "/api/chat/history",
            {
                params:{
                    application_id:props.applicationId
                }
            }
        )

        chatMessages.value=result

        scrollToBottom()

    }catch(err){

        console.error(err)

    }

}

//----------------------------------------------------
// Connect Socket
//----------------------------------------------------

const connectSocket = () => {

    console.log("CONNECT SOCKET")

    if (
        socket &&
        (
            socket.readyState === WebSocket.OPEN ||
            socket.readyState === WebSocket.CONNECTING
        )
    ) {
        console.log("Socket already exists")
        return
    }

    socket = new WebSocket(
        `ws://localhost:3000/api/chat/ws?userId=${props.currentUserId}`
    )

    socket.onopen = () => {
        console.log("OPEN", socket)
    }

    socket.onmessage = (event) => {
        console.log("MESSAGE")

        const message = JSON.parse(event.data)

        if(message.application_id!==props.applicationId){
            return
        }

        chatMessages.value.push(message)
    }

    socket.onclose = () => {
        console.log("CLOSE")
        socket = null
    }
}
//----------------------------------------------------
// Send
//----------------------------------------------------

const sendMessage = () => {
  if (!socket) return

  if (socket.readyState !== WebSocket.OPEN) return
  console.log("CURRENT =", props.currentUserId)
console.log("TARGET =", props.targetId)

  const text = newMessageText.value.trim()
  if (text === "") return

  socket.send(
    JSON.stringify({
      application_id: props.applicationId,
      receiver_id: props.targetId,
      message_text: text
    })
    
  )

  // Không push vào chatMessages nữa
  // Chờ server gửi lại qua WebSocket

  newMessageText.value = ""
}
//----------------------------------------------------

watch(
  () => props.currentUserId,
  (id) => {
    if (!id) return

    if (socket) return

    connectSocket()
  },
  {
    immediate: true
  }
)

//----------------------------------------------------

onMounted(()=>{

    fetchHistory()

})

//----------------------------------------------------

onUnmounted(()=>{

    if(socket){

        socket.close()

    }

})
</script>