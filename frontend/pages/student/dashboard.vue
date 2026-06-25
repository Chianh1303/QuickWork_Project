<template>
  <div class="min-h-screen bg-slate-50 flex flex-col md:flex-row">
    <!-- Sidebar Navigation -->
    <aside class="w-full md:w-64 bg-white border-r border-slate-200 flex-shrink-0">
      <div class="p-6 border-b border-slate-100 flex items-center space-x-3">
        <span class="h-8 w-8 rounded-lg bg-gradient-to-tr from-blue-600 to-emerald-500 flex items-center justify-center text-white font-bold text-sm">
          S
        </span>
        <div>
          <h1 class="font-bold text-slate-900 text-sm">Student Portal</h1>
          <p class="text-xs text-slate-500 font-medium">QuickWork Platform</p>
        </div>
      </div>
      <nav class="p-4 space-y-1">
        <button
          v-for="item in navItems"
          :key="item.id"
          @click="activeSection = item.id"
          :class="[
            activeSection === item.id
              ? 'bg-blue-50 text-blue-700 font-semibold'
              : 'text-slate-600 hover:text-slate-900 hover:bg-slate-50',
            'w-full flex items-center space-x-3 px-4 py-2.5 text-sm rounded-xl transition-all duration-150'
          ]"
        >
          <!-- Render icons inline to avoid runtime compilation errors in Nuxt -->
          <svg v-if="item.id === 'jobs'" class="h-5 w-5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
          </svg>
          <svg v-else-if="item.id === 'profile'" class="h-5 w-5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
          </svg>
          <svg v-else-if="item.id === 'applications'" class="h-5 w-5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
          </svg>
          <span>{{ item.name }}</span>
        </button>
      </nav>
    </aside>

    <!-- Main Content Area -->
    <main class="flex-grow p-6 sm:p-8 bg-slate-50 overflow-y-auto">
      <!-- Section 1: Dashboard (Explore Jobs) -->
      <div v-if="activeSection === 'jobs'" class="space-y-6">
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between pb-6 border-b border-slate-200">
          <div>
            <h2 class="text-3xl font-extrabold text-slate-900 tracking-tight">Available Gigs & Jobs</h2>
            <p class="mt-1 text-sm text-slate-500 font-medium">Explore current work opportunities in real-time.</p>
          </div>
        </div>

        <!-- Feedback Banner -->
        <div v-if="feedback" :class="[
          feedback.type === 'success' ? 'bg-emerald-50 border-emerald-300 text-emerald-800' : 'bg-red-50 border-red-300 text-red-800',
          'border-l-4 p-4 rounded-r-lg flex justify-between items-start transition-all duration-300 shadow-sm'
        ]">
          <div class="flex items-center space-x-3">
            <svg v-if="feedback.type === 'success'" class="h-5 w-5 text-emerald-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <svg v-else class="h-5 w-5 text-red-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <span class="text-sm font-semibold">{{ feedback.message }}</span>
          </div>
          <button @click="feedback = null" class="text-slate-400 hover:text-slate-600">
            <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Filters -->
        <div class="bg-white p-4 rounded-xl border border-slate-200 shadow-sm grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div class="relative">
            <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
              <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </span>
            <input
              v-model="jobsSearchQuery"
              type="text"
              class="block w-full pl-10 pr-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 placeholder-slate-400 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white transition-all duration-200"
              placeholder="Search by job title or keyword..."
            />
          </div>
          <div class="relative">
            <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
              <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
            </span>
            <input
              v-model="jobsLocationQuery"
              type="text"
              class="block w-full pl-10 pr-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 placeholder-slate-400 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white transition-all duration-200"
              placeholder="Filter by location..."
            />
          </div>
        </div>

        <!-- Skeleton / Loading -->
        <div v-if="isLoadingJobs" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div v-for="n in 6" :key="n" class="bg-white rounded-xl border border-slate-200 p-6 space-y-4 animate-pulse">
            <div class="h-4 bg-slate-200 rounded w-2/3"></div>
            <div class="h-3 bg-slate-200 rounded w-1/2"></div>
            <div class="h-16 bg-slate-100 rounded"></div>
            <div class="h-8 bg-slate-200 rounded"></div>
          </div>
        </div>

        <!-- Jobs Grid -->
        <div v-else-if="filteredJobs.length === 0" class="bg-white text-center py-16 px-4 rounded-xl border border-slate-200 shadow-sm text-slate-500">
          <svg class="mx-auto h-12 w-12 text-slate-300 mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <p class="font-bold text-slate-700 text-lg">No work listings found</p>
          <p class="text-sm text-slate-400 mt-1">Try modifying your search or filters.</p>
        </div>

        <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div
            v-for="job in filteredJobs"
            :key="job.id"
            class="bg-white rounded-xl border border-slate-200 shadow-sm p-6 flex flex-col justify-between hover:shadow-md transition-shadow duration-200"
          >
            <div>
              <div class="flex justify-between items-start mb-3">
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-semibold bg-blue-50 text-blue-700">
                  {{ companyNameLookup(job.business_id) }}
                </span>
                <span class="text-sm font-extrabold text-emerald-600 bg-emerald-50 px-2.5 py-1 rounded-lg">
                  ${{ Number(job.salary || 0).toLocaleString() }}
                </span>
              </div>
              <h3 class="text-lg font-bold text-slate-900 line-clamp-1">{{ job.title || 'Untitled Job' }}</h3>
              <p class="text-sm font-medium text-slate-500 mt-1 flex items-center gap-1">
                <svg class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                </svg>
                {{ job.location || 'Location N/A' }}
              </p>
              <p class="text-sm font-medium text-slate-500 mt-1 flex items-center gap-1" v-if="job.working_date">
                <svg class="h-4 w-4 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
                {{ job.working_date }}
              </p>
              <p class="text-sm text-slate-600 mt-4 line-clamp-4 bg-slate-50 p-3 rounded-lg border border-slate-100 font-medium">
                {{ job.description || 'No description provided.' }}
              </p>
            </div>

            <div class="mt-6 pt-4 border-t border-slate-100">
              <button
                @click="handleApply(job.id)"
                :disabled="isApplying === job.id"
                class="w-full flex justify-center py-2.5 px-4 border border-transparent text-sm font-semibold rounded-lg text-white bg-blue-600 hover:bg-blue-500 transition-colors focus-ring disabled:opacity-50 disabled:cursor-not-allowed shadow-sm"
              >
                <span v-if="isApplying === job.id" class="flex items-center space-x-2">
                  <svg class="animate-spin h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  <span>Applying...</span>
                </span>
                <span v-else>Apply Instantly</span>
              </button>
            </div>
          </div>
        </div>
      </div>

     <!-- Section 2: Profile Settings -->
      <div v-show="activeSection === 'profile'" class="max-w-2xl mx-auto space-y-6">
        
        <!-- Header Tổng -->
        <div class="flex items-center justify-between pb-6 border-b border-slate-200">
          <div>
            <h2 class="text-3xl font-extrabold text-slate-900 tracking-tight">My Professional Profile</h2>
            <p class="mt-1 text-sm text-slate-500 font-medium">View and manage your portfolio details, core skills, and active CV.</p>
          </div>
          <!-- Nút chuyển đổi trạng thái nếu không trong chế độ sửa -->
          <button 
            v-if="!isEditing" 
            @click="isEditing = true"
            class="px-4 py-2 text-sm font-bold text-blue-600 bg-blue-50 hover:bg-blue-100 rounded-lg transition-all flex items-center space-x-1"
          >
            <span>✏️</span> <span>Edit Profile</span>
          </button>
        </div>

        <!-- ======================================================== -->
        <!-- TRẠNG THÁI 1: HIỂN THỊ DỮ LIỆU CŨ (VIEW MODE) -->
        <!-- ======================================================== -->
        <div v-if="!isEditing" class="space-y-6">
          <!-- Thẻ Card thông tin chính -->
          <div class="bg-white rounded-2xl border border-slate-200 shadow-sm p-6 flex flex-col sm:flex-row items-center sm:items-start space-y-4 sm:space-y-0 sm:space-x-6">
            <img 
              :src="profileForm.avatar_url || 'https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?auto=format&fit=crop&w=150&h=150&q=80'" 
              class="w-24 h-24 rounded-full object-cover border-4 border-slate-50 shadow-md"
            />
            <div class="flex-1 text-center sm:text-left space-y-2">
              <h3 class="text-2xl font-bold text-slate-900">{{ profileForm.full_name || 'Your Full Name' }}</h3>
              <div class="flex flex-wrap justify-center sm:justify-start gap-4 text-sm text-slate-600 font-medium">
                <span class="flex items-center">📞 {{ profileForm.phone || 'Not updated yet' }}</span>
                <span class="flex items-center">🚻 {{ profileForm.gender || 'Not specified' }}</span>
              </div>
            </div>
          </div>

          <!-- Khối Kỹ năng -->
          <div class="bg-white rounded-2xl border border-slate-200 shadow-sm p-6 space-y-3">
            <h4 class="text-sm font-bold text-slate-400 uppercase tracking-wider">Core Skills & Technologies</h4>
            <div v-if="skillsArray.length > 0" class="flex flex-wrap gap-2">
              <span 
                v-for="skill in skillsArray" 
                :key="skill"
                class="px-3 py-1 bg-slate-100 text-slate-800 text-xs font-semibold rounded-full border border-slate-200"
              >
                {{ skill }}
              </span>
            </div>
            <p v-else class="text-sm text-slate-400 italic">No skills added yet. Click edit to add your stack.</p>
          </div>

          <!-- Khối File CV Đính Kèm -->
          <div class="bg-white rounded-2xl border border-slate-200 shadow-sm p-6 space-y-3">
            <h4 class="text-sm font-bold text-slate-400 uppercase tracking-wider">Attached Curriculum Vitae (CV)</h4>
            <div v-if="profileForm.cv_url" class="flex items-center justify-between bg-emerald-50 border border-emerald-100 p-4 rounded-xl">
              <div class="flex items-center space-x-3">
                <span class="text-3xl">📄</span>
                <div>
                  <p class="text-sm font-bold text-emerald-950">Your CV is Active</p>
                  <p class="text-xs text-emerald-700">Ready to apply for jobs on QuickWork</p>
                </div>
              </div>
              <a 
                :href="profileForm.cv_url" 
                target="_blank" 
                class="px-4 py-2 bg-emerald-600 hover:bg-emerald-700 text-white text-xs font-bold rounded-lg shadow-sm transition-colors"
              >
                View CV ↗
              </a>
            </div>
            <div v-else class="text-center p-6 border-2 border-dashed border-slate-200 rounded-xl bg-slate-50">
              <p class="text-sm text-slate-500 font-medium">You haven't uploaded any CV PDF file yet.</p>
            </div>
          </div>
        </div>

        <!-- ======================================================== -->
        <!-- TRẠNG THÁI 2: FORM CHỈNH SỬA (EDIT MODE) -->
        <!-- ======================================================== -->
        <div v-else class="bg-white rounded-2xl border border-slate-200 shadow-sm p-6">
          <form @submit.prevent="handleUpdateProfile" class="space-y-6">
            
            <!-- Upload Ảnh -->
            <div class="flex items-center space-x-6 bg-slate-50 p-4 rounded-xl border border-slate-100">
              <img 
                :src="avatarPreview || profileForm.avatar_url || 'https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?auto=format&fit=crop&w=150&h=150&q=80'" 
                class="w-20 h-20 rounded-full object-cover border-2 border-slate-200 shadow-sm"
              />
              <div>
                <input type="file" id="profile_avatar_file" accept="image/*" class="hidden" @change="onAvatarFileChange" />
                <label for="profile_avatar_file" class="px-4 py-2 text-xs font-bold text-slate-700 bg-white border border-slate-300 rounded-lg shadow-sm hover:bg-slate-50 cursor-pointer transition-all">
                  Change Avatar Image
                </label>
                <p class="text-xs text-slate-400 mt-1.5">JPG or PNG. Max 2MB.</p>
              </div>
            </div>

            <!-- Các input chữ -->
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
              <div>
                <label for="profile_fullname" class="block text-sm font-semibold text-slate-700 mb-1">Full Name</label>
                <input id="profile_fullname" type="text" v-model="profileForm.full_name" required class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white" />
              </div>
              <div>
                <label for="profile_phone" class="block text-sm font-semibold text-slate-700 mb-1">Phone Number</label>
                <input id="profile_phone" type="tel" v-model="profileForm.phone" class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white" />
              </div>
              <div>
                <label for="profile_gender" class="block text-sm font-semibold text-slate-700 mb-1">Gender</label>
                <select id="profile_gender" v-model="profileForm.gender" class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white">
                  <option value="">Select Gender</option>
                  <option value="Male">Male</option>
                  <option value="Female">Female</option>
                  <option value="Other">Other</option>
                </select>
              </div>
            </div>

            <!-- Kỹ năng -->
            <div>
              <label for="profile_skills" class="block text-sm font-semibold text-slate-700 mb-1">Skills (Comma-separated)</label>
              <input id="profile_skills" type="text" v-model="skillsText" class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white" placeholder="Go, Vue, JavaScript" />
            </div>

            <!-- Upload File CV -->
            <div class="space-y-2">
              <label class="block text-sm font-semibold text-slate-700">CV File (PDF Only)</label>
              <div class="border-2 border-dashed border-slate-200 hover:border-blue-400 rounded-xl p-5 transition-colors text-center relative bg-slate-50">
                <input type="file" id="profile_cv_file" accept=".pdf" class="absolute inset-0 w-full h-full opacity-0 cursor-pointer" @change="onCvFileChange" />
                <div class="space-y-1">
                  <span class="text-xl">📄</span>
                  <p class="text-sm font-medium text-slate-700">
                    {{ cvFileSelected ? cvFileSelected.name : 'Click to upload your new CV' }}
                  </p>
                </div>
              </div>
            </div>

            <!-- Các nút Lưu / Hủy -->
            <div class="pt-4 border-t border-slate-100 flex justify-end space-x-3">
              <button 
                type="button" 
                @click="isEditing = false"
                class="px-5 py-2.5 border border-slate-200 text-sm font-semibold rounded-lg text-slate-700 bg-white hover:bg-slate-50 transition-all"
              >
                Cancel
              </button>
              <button
                type="submit"
                :disabled="isSavingProfile"
                class="px-6 py-2.5 border border-transparent text-sm font-semibold rounded-lg text-white bg-blue-600 hover:bg-blue-500 focus:ring-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed shadow-sm transition-all duration-150"
              >
                <span v-if="isSavingProfile">Saving Profile...</span>
                <span v-else>Save Changes</span>
              </button>
            </div>
          </form>
        </div>

      </div>      <!-- Section 3: Student Application History (C6) -->
      <div v-if="activeSection === 'applications'" class="space-y-6">
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between pb-6 border-b border-slate-200">
          <div>
            <h2 class="text-3xl font-extrabold text-slate-900 tracking-tight">My Applications</h2>
            <p class="mt-1 text-sm text-slate-500 font-medium">Track every opportunity you've applied for.</p>
          </div>
          <div class="mt-4 sm:mt-0 bg-white border border-slate-200 px-4 py-2 rounded-xl shadow-sm text-sm font-semibold text-slate-700">
            Total Applications: {{ filteredApps.length }}
          </div>
        </div>

        <!-- Search & Filter Controls -->
        <div class="bg-white p-4 rounded-xl border border-slate-200 shadow-sm grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div class="relative">
            <span class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
              <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </span>
            <input
              v-model="appSearchQuery"
              type="text"
              class="block w-full pl-10 pr-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 placeholder-slate-400 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white transition-all duration-200"
              placeholder="Search by job title or company..."
            />
          </div>
          <div>
            <select
              v-model="appStatusFilter"
              class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:bg-white transition-all duration-200"
            >
              <option value="all">All Application Statuses</option>
              <option value="applied">Applied (Pending)</option>
              <option value="approved">Approved</option>
              <option value="rejected">Rejected</option>
            </select>
          </div>
        </div>

        <!-- Loading / Skeleton -->
        <div v-if="isLoadingApps" class="bg-white rounded-xl border border-slate-200 p-6 space-y-4 animate-pulse">
          <div class="h-6 bg-slate-200 rounded w-1/4"></div>
          <div class="space-y-3">
            <div class="h-10 bg-slate-100 rounded"></div>
            <div class="h-10 bg-slate-100 rounded"></div>
            <div class="h-10 bg-slate-100 rounded"></div>
          </div>
        </div>

        <!-- Empty State -->
        <div v-else-if="filteredApps.length === 0" class="bg-white text-center py-16 px-4 rounded-xl border border-slate-200 shadow-sm">
          <div class="inline-flex p-4 rounded-full bg-slate-50 border border-slate-100 text-slate-400 mb-4">
            <svg class="h-12 w-12" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
          </div>
          <h3 class="text-lg font-bold text-slate-900">No applications matched</h3>
          <p class="text-sm text-slate-500 mt-1 max-w-sm mx-auto font-medium">
            You haven't applied to any roles matching these filters. Go to the jobs listing tab to apply!
          </p>
        </div>

        <!-- Desktop Applications Table -->
<div
  v-if="filteredApps.length > 0"
  class="hidden md:block bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden"
>          <table class="min-w-full divide-y divide-slate-200">
            <thead class="bg-slate-50">
              <tr>
                <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Position & Company</th>
                <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Location</th>
                <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Compensation</th>
                <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Applied Date</th>
                <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Status</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-slate-100">
              <tr v-for="app in filteredApps" :key="app.id">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm font-semibold text-slate-900">{{ app.job?.title || 'Unknown Position' }}</div>
                  <div class="text-xs text-slate-500 font-medium">{{ companyNameLookup(app.job?.business_id) }}</div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-600 font-medium">
                  {{ app.job?.location || 'Location N/A' }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-extrabold text-emerald-600">
                  ${{ Number(app.job?.salary || 0).toLocaleString() }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500 font-medium">
                  {{ formatDate(app.applied_at || app.id) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="[
                    statusBadgeClass(app.status),
                    'inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold border'
                  ]">
                    {{ app.status }}
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Mobile Applications Cards -->
        <div v-show="filteredApps.length > 0" class="md:hidden space-y-4">
          <div
            v-for="app in filteredApps"
            :key="app.id"
            class="bg-white border border-slate-200 rounded-xl p-5 shadow-sm space-y-3"
          >
            <div class="flex justify-between items-start">
              <div>
                <h4 class="text-base font-bold text-slate-900">{{ app.job?.title || 'Unknown Position' }}</h4>
                <p class="text-sm text-slate-500 font-medium">{{ companyNameLookup(app.job?.business_id) }}</p>
              </div>
              <span :class="[
                statusBadgeClass(app.status),
                'inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold border'
              ]">
                {{ app.status }}
              </span>
            </div>

            <div class="grid grid-cols-2 gap-2 pt-2 text-xs border-t border-slate-100 font-medium text-slate-500">
              <div>Location: <span class="text-slate-900">{{ app.job?.location || 'Location N/A' }}</span></div>
              <div>Salary: <span class="text-emerald-600 font-bold">${{ Number(app.job?.salary || 0).toLocaleString() }}</span></div>
              <div class="col-span-2">Applied Date: <span class="text-slate-900">{{ formatDate(app.applied_at || app.id) }}</span></div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { useApi } from '~/composables/useApi'

definePageMeta({
  middleware: 'auth'
})

const activeSection = ref('jobs')

const navItems = [
  { id: 'jobs', name: 'Dashboard' },
  { id: 'profile', name: 'Profile' },
  { id: 'applications', name: 'My Applications' }
]

// State setup
const api = useApi()
const jobs = ref<any[]>([])
const applications = ref<any[]>([])
const isLoadingJobs = ref(false)
const isLoadingApps = ref(false)
const isApplying = ref<number | null>(null)
const isSavingProfile = ref(false)
const feedback = ref<{ type: 'success' | 'error'; message: string } | null>(null)

// 🌟 TRẠNG THÁI UI PROFILE MỚI (VIEW / EDIT MODE & FILE MANAGEMENT)
const isEditing = ref(false)
const skillsText = ref('') // Dùng để binding riêng cho ô input text (Edit Mode)
const skillsArray = ref<string[]>([]) // Dùng để hiển thị danh sách tag (View Mode)
const avatarFileSelected = ref<File | null>(null)
const avatarPreview = ref<string | null>(null)
const cvFileSelected = ref<File | null>(null)

// Search Query filters
const jobsSearchQuery = ref('')
const jobsLocationQuery = ref('')
const appSearchQuery = ref('')
const appStatusFilter = ref('all')

// Profile update form reactive state
const profileForm = reactive({
  full_name: '',
  phone: '',
  gender: '',
  avatar_url: '',
  skills: '',
  cv_url: ''
})

// Company Seed Lookup
const companyLookup: Record<number, string> = {
  1: "FPT Software",
  2: "VNG Corporation",
  3: "Viettel Digital",
  4: "MISA Joint Stock Company",
  5: "NashTech Vietnam"
}

const companyNameLookup = (id: number): string => {
  return companyLookup[id] || `Verified Employer #${id || ''}`
}

// Format Application Date
const formatDate = (dateVal: any): string => {
  return 'Jun 24, 2026'
}

// Status Badges Style config
const statusBadgeClass = (status: string): string => {
  const norm = (status || '').toLowerCase()
  if (norm === 'approved') return 'bg-emerald-50 border-emerald-200 text-emerald-700'
  if (norm === 'rejected') return 'bg-rose-50 border-rose-200 text-rose-700'
  return 'bg-slate-50 border-slate-200 text-slate-700'
}

// Computeds for filtering with defensive null-safety
const filteredJobs = computed(() => {
  const list = Array.isArray(jobs.value) ? jobs.value : []
  return list.filter(job => {
    if (!job) return false
    const title = (job.title || '').toLowerCase()
    const description = (job.description || '').toLowerCase()
    const location = (job.location || '').toLowerCase()

    const matchesSearch = !jobsSearchQuery.value ||
      title.includes(jobsSearchQuery.value.toLowerCase()) ||
      description.includes(jobsSearchQuery.value.toLowerCase())
    
    const matchesLocation = !jobsLocationQuery.value ||
      location.includes(jobsLocationQuery.value.toLowerCase())
      
    return matchesSearch && matchesLocation
  })
})

const filteredApps = computed(() => {
  const list = Array.isArray(applications.value) ? applications.value : []
  return list.filter(app => {
    if (!app) return false
    const jobTitle = (app.job?.title || '').toLowerCase()
    const businessId = app.job?.business_id || 0
    const compName = companyNameLookup(businessId).toLowerCase()
    
    const matchesSearch = !appSearchQuery.value ||
      jobTitle.includes(appSearchQuery.value.toLowerCase()) ||
      compName.includes(appSearchQuery.value.toLowerCase())
    
    const matchesStatus = appStatusFilter.value === 'all' ||
      (app.status || '').toLowerCase() === appStatusFilter.value.toLowerCase()

    return matchesSearch && matchesStatus
  })
})

// API Operations
const fetchJobs = async () => {
  isLoadingJobs.value = true
  try {
    const res = await api.get('/api/jobs')
    jobs.value = Array.isArray(res)
      ? res
      : (res && Array.isArray(res.data) ? res.data : [])
  } catch (err: any) {
    console.error('Error fetching jobs:', err)
  } finally {
    isLoadingJobs.value = false
  }
}

const fetchApplications = async () => {
  isLoadingApps.value = true
  try {
    const res = await api.get('/api/applications/my-applications')
    applications.value = Array.isArray(res)
      ? res
      : (res && Array.isArray(res.data) ? res.data : [])
  } catch (err: any) {
    console.error('Error fetching student applications:', err)
  } finally {
    isLoadingApps.value = false
  }
}

// 🌟 HÀM TẢI PROFILE ĐỂ ĐỒNG BỘ DỮ LIỆU CŨ LÊN MÀN HÌNH CHÍNH
const fetchProfile = async () => {
  try {
    const res = await api.get('/api/profile/student')
    const data = res && res.data ? res.data : res
    
    if (data) {
      profileForm.full_name = data.full_name || ''
      profileForm.phone = data.phone || ''
      profileForm.gender = data.gender || ''
      profileForm.avatar_url = data.avatar_url || ''
      profileForm.cv_url = data.cv_url || ''
      profileForm.skills = data.skills || ''

      // Xử lý bóc tách Skills phục vụ cho cả 2 Mode (View và Edit)
      if (data.skills) {
        try {
          const parsed = JSON.parse(data.skills)
          if (Array.isArray(parsed)) {
            skillsArray.value = parsed
            skillsText.value = parsed.join(', ')
          } else {
            skillsText.value = data.skills
            skillsArray.value = data.skills.split(',').map((s: string) => s.trim())
          }
        } catch {
          skillsText.value = data.skills
          skillsArray.value = data.skills.split(',').map((s: string) => s.trim())
        }
      }
    }
  } catch (err: any) {
    console.error('Error fetching profile details:', err)
  }
}

const handleApply = async (jobId: number) => {
  isApplying.value = jobId
  feedback.value = null

  try {
    const res = await api.post('/api/jobs/apply', { job_id: jobId })
    feedback.value = {
      type: 'success',
      message: res.message || '🚀 Applied successfully! Waiting for Employer review.'
    }
    await fetchApplications()
  } catch (err: any) {
    feedback.value = {
      type: 'error',
      message: err.response?._data?.message || 'Failed to submit application.'
    }
  } finally {
    isApplying.value = null
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

// 🌟 LOGIC ĐỒNG BỘ FILE KHI KHÁCH HÀNG CHỌN FILE TRÊN INTERFACE
const onAvatarFileChange = (e: Event) => {
  const target = e.target as HTMLInputElement
  const file = target.files?.[0]
  if (file) {
    avatarFileSelected.value = file
    avatarPreview.value = URL.createObjectURL(file) // Tạo ảnh xem trước tạm thời
  }
}

const onCvFileChange = (e: Event) => {
  const target = e.target as HTMLInputElement
  const file = target.files?.[0]
  if (file) {
    if (file.type !== 'application/pdf') {
      alert('Please select a valid PDF file for your CV!')
      return
    }
    cvFileSelected.value = file
  }
}

// 🌟 HÀM CẬP NHẬT PROFILE NÂNG CẤP SANG FORMDATA ĐỂ TRUYỀN FILE
const handleUpdateProfile = async () => {
  isSavingProfile.value = true
  feedback.value = null

  try {
    // Khởi tạo FormData để bọc file vật lý
    const formData = new FormData()
    formData.append('full_name', profileForm.full_name)
    formData.append('phone', profileForm.phone)
    formData.append('gender', profileForm.gender)

    // Đồng bộ chuỗi nhập từ ô input text thành chuỗi mảng JSON gửi xuống Go
    const parsedSkills = skillsText.value.split(',').map(s => s.trim()).filter(s => s !== '')
    formData.append('skills', JSON.stringify(parsedSkills))

    // Đính kèm các tệp tin thực tế nếu có thay đổi
    if (avatarFileSelected.value) formData.append('avatar', avatarFileSelected.value)
    if (cvFileSelected.value) formData.append('cv', cvFileSelected.value)

    // Bắn API lên Server thông qua useApi cơ bản nhưng truyền body dạng FormData
    const res = await api.put('/api/profile/student', formData)
    
    feedback.value = {
      type: 'success',
      message: res.message || '🎉 Profile updated successfully with files!'
    }

    // Cập nhật lại dữ liệu local từ Server trả về sau khi upload
    const updatedData = res.data ? res.data : res
    if (updatedData) {
      profileForm.avatar_url = updatedData.AvatarUrl || updatedData.avatar_url || profileForm.avatar_url
      profileForm.cv_url = updatedData.CvUrl || updatedData.cv_url || profileForm.cv_url
    }

    // Đồng bộ lại danh sách tag hiển thị và đóng form
    skillsArray.value = parsedSkills
    isEditing.value = false
    
    // Clear dữ liệu file tạm thời sau khi lưu xong
    avatarFileSelected.value = null
    cvFileSelected.value = null
  } catch (err: any) {
    feedback.value = {
      type: 'error',
      message: err.response?._data?.message || 'Failed to update profile.'
    }
  } finally {
    isSavingProfile.value = false
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

onMounted(() => {
  fetchJobs()
  fetchApplications()
  fetchProfile() // Gọi hàm tải profile cũ ngay khi mở trang
})
</script>