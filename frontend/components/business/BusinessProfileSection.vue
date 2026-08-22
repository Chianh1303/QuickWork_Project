<template>
  <!-- Section 2: Company Profile (Enterprise SaaS Style) -->
  <div v-show="activeSection === 'profile'" class="max-w-4xl mx-auto space-y-6">
    
    <!-- ======================================================== -->
    <!-- TRẠNG THÁI 1: VIEW MODE (HIỂN THỊ HỒ SƠ CHUYÊN NGHIỆP) -->
    <!-- ======================================================== -->
    <div v-if="!isEditing" class="space-y-6">
      
      <!-- Top Card: Header & KYB Verification Status -->
      <div class="bg-slate-900/90 rounded-3xl border border-indigo-500/20 shadow-2xl p-6 sm:p-8 backdrop-blur-xl space-y-6">
        <div class="flex flex-col sm:flex-row items-center sm:items-start justify-between gap-6 border-b border-white/10 pb-6">
          <div class="flex flex-col sm:flex-row items-center sm:items-start gap-5 text-center sm:text-left">
            <!-- Company Logo -->
            <div class="relative h-24 w-24 rounded-2xl overflow-hidden border-2 border-cyan-400/30 bg-slate-950 p-1 shadow-xl flex-shrink-0">
              <img 
                :src="profileForm.logo_url || 'https://images.unsplash.com/photo-1486406146926-c627a92ad1ab?auto=format&fit=crop&w=150&h=150&q=80'" 
                class="h-full w-full object-cover rounded-xl"
              />
            </div>

            <div class="space-y-1.5">
              <div class="flex flex-wrap items-center justify-center sm:justify-start gap-2.5">
                <h3 class="text-2xl font-extrabold text-white tracking-tight">{{ profileForm.company_name || 'Doanh Nghiệp Chưa Đặt Tên' }}</h3>
                
                <!-- KYB Verification Status Badge -->
                <span v-if="profileForm.is_verified" class="inline-flex items-center gap-1.5 rounded-full border border-emerald-400/30 bg-emerald-400/10 px-3 py-1 text-xs font-extrabold text-emerald-300">
                  <span class="h-2 w-2 rounded-full bg-emerald-400 animate-ping"></span>
                  Đã Xác Thực KYB 🟢
                </span>
                <span v-else class="inline-flex items-center gap-1.5 rounded-full border border-amber-400/30 bg-amber-400/10 px-3 py-1 text-xs font-extrabold text-amber-300">
                  <span class="h-2 w-2 rounded-full bg-amber-400"></span>
                  Đang Chờ Duyệt KYB 🟡
                </span>
              </div>

              <p class="text-xs font-semibold text-cyan-300">
                Mã Số Thuế (MST): <span class="font-extrabold text-white">{{ profileForm.tax_code || 'Chưa cập nhật MST' }}</span>
              </p>
              <p class="text-xs font-medium text-slate-400">
                Quy mô: <span class="text-slate-200 font-bold">{{ profileForm.company_size || '20-50' }} nhân sự</span>
              </p>
            </div>
          </div>

          <button
            @click="isEditing = true"
            class="flex items-center gap-2 px-5 py-2.5 bg-gradient-to-r from-indigo-500 to-cyan-400 hover:brightness-110 text-white font-extrabold text-xs rounded-xl shadow-lg shadow-indigo-500/20 transition-all cursor-pointer whitespace-nowrap"
          >
            <!-- Pencil SVG Icon -->
            <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
            </svg>
            <span>Chỉnh Sửa Hồ Sơ</span>
          </button>
        </div>

        <!-- Corporate Info Grid -->
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 text-xs font-medium text-slate-300">
          <div class="flex items-center gap-3 p-3.5 rounded-xl border border-white/5 bg-slate-950/60">
            <span class="text-lg">📍</span>
            <div class="min-w-0">
              <p class="text-[10px] font-black uppercase text-slate-400">Trụ sở doanh nghiệp</p>
              <p class="font-bold text-white truncate">{{ profileForm.address || 'Chưa cập nhật địa chỉ' }}</p>
            </div>
          </div>

          <div class="flex items-center gap-3 p-3.5 rounded-xl border border-white/5 bg-slate-950/60">
            <span class="text-lg">📞</span>
            <div class="min-w-0">
              <p class="text-[10px] font-black uppercase text-slate-400">Số điện thoại liên hệ</p>
              <p class="font-bold text-white truncate">{{ profileForm.phone || 'Chưa cập nhật SĐT' }}</p>
            </div>
          </div>

          <div class="flex items-center gap-3 p-3.5 rounded-xl border border-white/5 bg-slate-950/60">
            <span class="text-lg">🌐</span>
            <div class="min-w-0">
              <p class="text-[10px] font-black uppercase text-slate-400">Website chính thức</p>
              <a v-if="profileForm.website" :href="profileForm.website" target="_blank" class="font-bold text-cyan-300 hover:underline truncate block">
                {{ profileForm.website }}
              </a>
              <p v-else class="font-bold text-slate-500">Chưa cập nhật website</p>
            </div>
          </div>

          <div class="flex items-center gap-3 p-3.5 rounded-xl border border-white/5 bg-slate-950/60">
            <span class="text-lg">📧</span>
            <div class="min-w-0">
              <p class="text-[10px] font-black uppercase text-slate-400">Email tuyển dụng HR</p>
              <p class="font-bold text-white truncate">{{ profileForm.contact_email || 'Chưa cập nhật email' }}</p>
            </div>
          </div>
        </div>

        <!-- Corporate Description -->
        <div class="pt-2 border-t border-white/10 space-y-2">
          <h4 class="text-xs font-black uppercase tracking-wider text-cyan-300">Giới Thiệu Doanh Nghiệp</h4>
          <p class="text-xs text-slate-300 leading-relaxed bg-slate-950/40 p-4 rounded-2xl border border-white/5">
            {{ profileForm.description || 'Doanh nghiệp chưa cập nhật bài giới thiệu tổng quan. Hãy bấm "Chỉnh Sửa Hồ Sơ" để cập nhật tầm nhìn và môi trường làm việc.' }}
          </p>
        </div>

      </div>
    </div>

    <!-- ======================================================== -->
    <!-- TRẠNG THÁI 2: EDIT MODE (FORM CẬP NHẬT HỒ SƠ) -->
    <!-- ======================================================== -->
    <div v-else class="bg-slate-900/95 rounded-3xl border border-cyan-400/30 shadow-2xl p-6 sm:p-8 backdrop-blur-xl space-y-6">
      <div class="border-b border-white/10 pb-4 flex items-center justify-between">
        <div>
          <span class="text-[10px] font-black uppercase tracking-wider text-cyan-300 bg-cyan-400/10 px-2.5 py-0.5 rounded-full border border-cyan-400/20">
            Cập Nhật Thông Tin
          </span>
          <h3 class="mt-1 text-lg font-extrabold text-white">Chỉnh Sửa Hồ Sơ Doanh Nghiệp</h3>
        </div>
      </div>

      <form @submit.prevent="handleUpdateProfile" class="space-y-6">
        
        <!-- Logo Upload Row -->
        <div class="flex items-center gap-6 bg-slate-950/80 p-4 rounded-2xl border border-white/10">
          <img 
            :src="logoPreview || profileForm.logo_url || 'https://images.unsplash.com/photo-1486406146926-c627a92ad1ab?auto=format&fit=crop&w=150&h=150&q=80'" 
            class="h-20 w-20 rounded-2xl object-cover border-2 border-cyan-400/30 bg-slate-900 p-1 shadow-md"
          />
          <div>
            <input type="file" id="profile_logo_file" accept="image/*" class="hidden" @change="onLogoFileChange" />
            <label for="profile_logo_file" class="px-4 py-2 text-xs font-extrabold text-white bg-indigo-600 hover:bg-indigo-500 rounded-xl shadow-md cursor-pointer transition-all inline-block">
              Tải Logo Mới
            </label>
            <p class="text-[11px] text-slate-400 mt-1.5">Định dạng JPG, PNG. Dung lượng tối đa 2MB.</p>
          </div>
        </div>

        <!-- Form Inputs Grid -->
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <!-- Company Name -->
          <div>
            <label for="profile_company" class="block text-xs font-bold text-slate-300 mb-1">Tên Doanh nghiệp <span class="text-rose-400">*</span></label>
            <input
              id="profile_company"
              type="text"
              v-model="profileForm.company_name"
              required
              class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-bold"
              placeholder="VD: CÔNG TY CỔ PHẦN FPT SOFTWARE"
            />
          </div>

          <!-- Tax Code -->
          <div>
            <label for="profile_tax_code" class="block text-xs font-bold text-slate-300 mb-1">Mã Số Thuế (MST / KYB) <span class="text-rose-400">*</span></label>
            <input
              id="profile_tax_code"
              type="text"
              v-model="profileForm.tax_code"
              required
              class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-bold"
              placeholder="VD: 0101234567"
            />
          </div>

          <!-- Phone -->
          <div>
            <label for="profile_phone" class="block text-xs font-bold text-slate-300 mb-1">Số điện thoại doanh nghiệp</label>
            <input
              id="profile_phone"
              type="tel"
              v-model="profileForm.phone"
              class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium"
              placeholder="VD: 028 1234 5678"
            />
          </div>

          <!-- Contact Email -->
          <div>
            <label for="profile_email" class="block text-xs font-bold text-slate-300 mb-1">Email tuyển dụng HR</label>
            <input
              id="profile_email"
              type="email"
              v-model="profileForm.contact_email"
              class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium"
              placeholder="VD: hr@company.com"
            />
          </div>

          <!-- Address -->
          <div class="sm:col-span-2">
            <label for="profile_address" class="block text-xs font-bold text-slate-300 mb-1">Địa chỉ trụ sở chính</label>
            <input
              id="profile_address"
              type="text"
              v-model="profileForm.address"
              class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium"
              placeholder="VD: Tòa nhà FPT, Quận Kế Nối, TP. Hồ Chí Minh"
            />
          </div>

          <!-- Website -->
          <div>
            <label for="profile_website" class="block text-xs font-bold text-slate-300 mb-1">Website chính thức</label>
            <input
              id="profile_website"
              type="url"
              v-model="profileForm.website"
              class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium"
              placeholder="https://company.com"
            />
          </div>

          <!-- Company Size -->
          <div>
            <label for="profile_size" class="block text-xs font-bold text-slate-300 mb-1">Quy mô Doanh nghiệp</label>
            <select
              id="profile_size"
              v-model="profileForm.company_size"
              class="block w-full px-4 py-2.5 border border-white/10 rounded-xl text-xs bg-slate-950 text-slate-100 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium cursor-pointer"
            >
              <option value="Dưới 20">Dưới 20 nhân sự</option>
              <option value="20-50">20 - 50 nhân sự</option>
              <option value="50-200">50 - 200 nhân sự</option>
              <option value="200-500">200 - 500 nhân sự</option>
              <option value="Trên 500">Trên 500 nhân sự</option>
            </select>
          </div>
        </div>

        <!-- Description -->
        <div>
          <label for="profile_desc" class="block text-xs font-bold text-slate-300 mb-1">Giới thiệu tổng quan về Doanh nghiệp</label>
          <textarea
            id="profile_desc"
            rows="5"
            v-model="profileForm.description"
            class="block w-full px-4 py-3 border border-white/10 rounded-xl text-xs bg-slate-950 text-slate-100 placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-400 font-medium leading-relaxed"
            placeholder="Giới thiệu tầm nhìn, lịch sử phát triển, văn hóa doanh nghiệp và các quyền lợi nổi bật..."
          ></textarea>
        </div>

        <!-- Buttons -->
        <div class="pt-4 border-t border-white/10 flex justify-end gap-3">
          <button 
            type="button" 
            @click="isEditing = false"
            class="px-5 py-2.5 border border-white/10 text-xs font-bold rounded-xl text-slate-300 bg-white/5 hover:bg-white/10 transition-all cursor-pointer"
          >
            Hủy bỏ
          </button>
          <button
            type="submit"
            :disabled="isSavingProfile"
            class="px-7 py-2.5 rounded-xl text-xs font-black text-slate-950 bg-gradient-to-r from-cyan-400 to-emerald-400 hover:brightness-110 shadow-lg shadow-cyan-500/25 transition-all disabled:opacity-50 cursor-pointer"
          >
            <span v-if="isSavingProfile" class="flex items-center space-x-2">
              <svg class="animate-spin h-4 w-4 text-slate-950" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              <span>Đang lưu hồ sơ...</span>
            </span>
            <span v-else>💾 Lưu Thay Đổi Hồ Sơ</span>
          </button>
        </div>
      </form>
    </div>

  </div>
</template>

<script setup lang="ts">
import { toRefs } from 'vue'

const props = defineProps<{ state: Record<string, any> }>()
const {
  activeSection,
  isEditing,
  profileForm,
  logoPreview,
  onLogoFileChange,
  isSavingProfile,
  handleUpdateProfile
} = toRefs(props.state)
</script>
