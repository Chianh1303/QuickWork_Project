<template>
  <div class="min-h-screen bg-slate-50 flex flex-col md:flex-row">
    <!-- Sidebar Navigation -->
    <aside class="w-full md:w-64 bg-white border-r border-slate-200 flex-shrink-0">
      <div class="p-6 border-b border-slate-100 flex items-center space-x-3">
        <span class="h-8 w-8 rounded-lg bg-gradient-to-tr from-violet-600 to-indigo-500 flex items-center justify-center text-white font-bold text-sm">
          B
        </span>
        <div>
          <h1 class="font-bold text-slate-900 text-sm">Employer Console</h1>
          <p class="text-xs text-slate-500 font-medium">QuickWork Business</p>
        </div>
      </div>
      <nav class="p-4 space-y-1">
        <button
          v-for="item in navItems"
          :key="item.id"
          @click="activeSection = item.id"
          :class="[
            activeSection === item.id
              ? 'bg-violet-50 text-violet-700 font-semibold'
              : 'text-slate-600 hover:text-slate-900 hover:bg-slate-50',
            'w-full flex items-center space-x-3 px-4 py-2.5 text-sm rounded-xl transition-all duration-150'
          ]"
        >
          <!-- Render icons inline to avoid runtime compilation errors in Nuxt -->
          <svg v-if="item.id === 'dashboard'" class="h-5 w-5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 002 2h2a2 2 0 002-2z"/>
          </svg>
          <svg v-else-if="item.id === 'profile'" class="h-5 w-5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"/>
          </svg>
          <svg v-else-if="item.id === 'jobs'" class="h-5 w-5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
          </svg>
          <svg v-else-if="item.id === 'applicants'" class="h-5 w-5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"/>
          </svg>
          <span>{{ item.name }}</span>
        </button>
      </nav>
    </aside>

    <!-- Main Content Area -->
    <main class="flex-grow p-6 sm:p-8 bg-slate-50 overflow-y-auto">
      <!-- Section 1: Dashboard (Metrics) -->
      <div v-if="activeSection === 'dashboard'" class="space-y-6">
        <div class="pb-6 border-b border-slate-200">
          <h2 class="text-3xl font-extrabold text-slate-900 tracking-tight font-sans">Performance Overview</h2>
          <p class="mt-1 text-sm text-slate-500 font-medium">Real-time candidate metrics and application actions.</p>
        </div>

        <!-- Metrics Cards -->
        <div class="grid grid-cols-2 lg:grid-cols-4 gap-5">
          <div v-for="card in metricsCards" :key="card.title" class="bg-white p-6 rounded-2xl border border-slate-200 shadow-sm flex flex-col justify-between">
            <span class="text-sm font-semibold text-slate-500 uppercase tracking-wider">{{ card.title }}</span>
            <div class="flex items-baseline justify-between mt-4">
              <span class="text-3xl font-extrabold text-slate-900">{{ card.value }}</span>
              <span :class="[card.color, 'text-xs font-semibold px-2 py-0.5 rounded-full border']">{{ card.label }}</span>
            </div>
          </div>
        </div>

        <!-- Job Postings Summary -->
        <div class="bg-white rounded-xl border border-slate-200 p-6 shadow-sm">
          <h3 class="text-lg font-bold text-slate-900 mb-4">Postings Summary</h3>
          <div class="grid grid-cols-3 gap-4 text-center">
            <div class="p-4 bg-slate-50 rounded-xl">
              <p class="text-xs font-semibold text-slate-500 uppercase">Active Jobs</p>
              <p class="text-2xl font-extrabold text-slate-900 mt-1">{{ jobs.length }}</p>
            </div>
            <div class="p-4 bg-slate-50 rounded-xl">
              <p class="text-xs font-semibold text-slate-500 uppercase">Incoming Applicants</p>
              <p class="text-2xl font-extrabold text-slate-900 mt-1">{{ applications.length }}</p>
            </div>
            <div class="p-4 bg-slate-50 rounded-xl">
              <p class="text-xs font-semibold text-slate-500 uppercase">Fill Ratio</p>
              <p class="text-2xl font-extrabold text-slate-900 mt-1">{{ fillRatio }}%</p>
            </div>
          </div>
        </div>
      </div>

    <!-- Section 2: Company Profile -->
      <div v-show="activeSection === 'profile'" class="max-w-2xl mx-auto space-y-6">
        
        <!-- Header Tổng -->
        <div class="flex items-center justify-between pb-6 border-b border-slate-200">
          <div>
            <h2 class="text-3xl font-extrabold text-slate-900 tracking-tight">Company Profile</h2>
            <p class="mt-1 text-sm text-slate-500 font-medium">Update public info, logo, and address details.</p>
          </div>
          <!-- Nút chuyển sang chế độ sửa -->
          <button 
            v-if="!isEditing" 
            @click="isEditing = true"
            class="px-4 py-2 text-sm font-bold text-violet-600 bg-violet-50 hover:bg-violet-100 rounded-lg transition-all flex items-center space-x-1"
          >
            <span>✏️</span> <span>Edit Profile</span>
          </button>
        </div>

        <!-- ======================================================== -->
        <!-- TRẠNG THÁI 1: HIỂN THỊ DỮ LIỆU CŨ (VIEW MODE) -->
        <!-- ======================================================== -->
        <div v-if="!isEditing" class="space-y-6">
          <!-- Card thông tin công ty chính -->
          <div class="bg-white rounded-2xl border border-slate-200 shadow-sm p-6 flex flex-col sm:flex-row items-center sm:items-start space-y-4 sm:space-y-0 sm:space-x-6">
            <img 
              :src="profileForm.logo_url || 'https://images.unsplash.com/photo-1486406146926-c627a92ad1ab?auto=format&fit=crop&w=150&h=150&q=80'" 
              class="w-24 h-24 rounded-xl object-cover border-2 border-slate-100 shadow-sm bg-slate-50 p-1"
            />
            <div class="flex-1 text-center sm:text-left space-y-3">
              <h3 class="text-2xl font-bold text-slate-900">{{ profileForm.company_name || 'Your Company Name' }}</h3>
              
              <div class="space-y-2 text-sm text-slate-600 font-medium">
                <p class="flex items-center justify-center sm:justify-start">
                  <span class="mr-2">🏢</span> {{ profileForm.address || 'Address not updated yet' }}
                </p>
                <p class="flex items-center justify-center sm:justify-start">
                  <span class="mr-2">📞</span> {{ profileForm.phone || 'Corporate phone not updated yet' }}
                </p>
              </div>
            </div>
          </div>
        </div>

        <!-- ======================================================== -->
        <!-- TRẠNG THÁI 2: FORM CHỈNH SỬA (EDIT MODE) -->
        <!-- ======================================================== -->
        <div v-else class="bg-white rounded-2xl border border-slate-200 shadow-sm p-6">
          <form @submit.prevent="handleUpdateProfile" class="space-y-6">
            
            <!-- Upload Logo thực tế -->
            <div class="flex items-center space-x-6 bg-slate-50 p-4 rounded-xl border border-slate-100">
              <img 
                :src="logoPreview || profileForm.logo_url || 'https://images.unsplash.com/photo-1486406146926-c627a92ad1ab?auto=format&fit=crop&w=150&h=150&q=80'" 
                class="w-20 h-20 rounded-xl object-cover border border-slate-200 shadow-sm bg-white p-1"
              />
              <div>
                <input type="file" id="profile_logo_file" accept="image/*" class="hidden" @change="onLogoFileChange" />
                <label for="profile_logo_file" class="px-4 py-2 text-xs font-bold text-slate-700 bg-white border border-slate-300 rounded-lg shadow-sm hover:bg-slate-50 cursor-pointer transition-all">
                  Change Company Logo
                </label>
                <p class="text-xs text-slate-400 mt-1.5">JPG or PNG. Max 2MB.</p>
              </div>
            </div>

            <!-- Các ô nhập dữ liệu chữ -->
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
              <div>
                <label for="profile_company" class="block text-sm font-semibold text-slate-700 mb-1">Company Name</label>
                <input
                  id="profile_company"
                  type="text"
                  v-model="profileForm.company_name"
                  required
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-violet-500 focus:bg-white"
                />
              </div>
              <div>
                <label for="profile_phone" class="block text-sm font-semibold text-slate-700 mb-1">Corporate Phone</label>
                <input
                  id="profile_phone"
                  type="tel"
                  v-model="profileForm.phone"
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-violet-500 focus:bg-white"
                />
              </div>
              <div class="sm:col-span-2">
                <label for="profile_address" class="block text-sm font-semibold text-slate-700 mb-1">Company Address</label>
                <input
                  id="profile_address"
                  type="text"
                  v-model="profileForm.address"
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-violet-500 focus:bg-white"
                  placeholder="e.g. 123 Corporate Blvd, Hanoi"
                />
              </div>
            </div>

            <!-- Nút điều hướng Lưu / Hủy -->
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
                class="px-6 py-2.5 border border-transparent text-sm font-semibold rounded-lg text-white bg-violet-600 hover:bg-violet-500 focus:ring-2 focus:ring-violet-500 disabled:opacity-50 disabled:cursor-not-allowed shadow-sm transition-all duration-150"
              >
                <span v-if="isSavingProfile" class="flex items-center space-x-2">
                  <svg class="animate-spin h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  <span>Saving Profile...</span>
                </span>
                <span v-else>Save Changes</span>
              </button>
            </div>
          </form>
        </div>
      </div>      <!-- Section 3: Jobs (List and Post Form) -->
      <div v-show="activeSection === 'jobs'" class="space-y-6">
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between pb-6 border-b border-slate-200">
          <div>
            <h2 class="text-3xl font-extrabold text-slate-900 tracking-tight">Active Listings</h2>
            <p class="mt-1 text-sm text-slate-500 font-medium">Publish, configure, and evaluate job listings.</p>
          </div>
          <button
            @click="showCreateForm = !showCreateForm"
            class="mt-4 sm:mt-0 px-4 py-2 bg-violet-600 hover:bg-violet-500 text-white font-semibold text-sm rounded-lg shadow-sm focus-ring transition-colors"
          >
            {{ showCreateForm ? 'Back to Listings' : '+ Post a New Job' }}
          </button>
        </div>

        <!-- Post job form -->
        <div v-if="showCreateForm" class="bg-white rounded-xl border border-slate-200 shadow-sm p-6 max-w-2xl mx-auto">
          <h3 class="text-lg font-bold text-slate-900 border-b border-slate-100 pb-3 mb-6">Configure Job Parameters</h3>
          <form @submit.prevent="handleCreateJob" class="space-y-6">
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
              <div class="sm:col-span-2">
                <label for="job_title" class="block text-sm font-semibold text-slate-700 mb-1">Job Title</label>
                <input
                  id="job_title"
                  type="text"
                  v-model="jobForm.title"
                  required
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-violet-500"
                  placeholder="e.g. Golang Backend developer Trainee"
                />
              </div>

              <div>
                <label for="job_location" class="block text-sm font-semibold text-slate-700 mb-1">Location</label>
                <input
                  id="job_location"
                  type="text"
                  v-model="jobForm.location"
                  required
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-violet-500"
                  placeholder="e.g. Cau Giay, Hanoi"
                />
              </div>

              <div>
                <label for="job_working_date" class="block text-sm font-semibold text-slate-700 mb-1">Working Dates</label>
                <input
                  id="job_working_date"
                  type="text"
                  v-model="jobForm.working_date"
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-violet-500"
                  placeholder="e.g. Monday - Friday"
                />
              </div>

              <div>
                <label for="job_salary" class="block text-sm font-semibold text-slate-700 mb-1">Salary ($)</label>
                <input
                  id="job_salary"
                  type="number"
                  step="0.01"
                  v-model="jobForm.salary"
                  required
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-violet-500"
                  placeholder="e.g. 500"
                />
              </div>

              <div>
                <label for="job_slots" class="block text-sm font-semibold text-slate-700 mb-1">Slots Available</label>
                <input
                  id="job_slots"
                  type="number"
                  v-model="jobForm.slots"
                  required
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-violet-500"
                  placeholder="1"
                />
              </div>

              <div class="sm:col-span-2">
                <label for="job_description" class="block text-sm font-semibold text-slate-700 mb-1">Detailed Description</label>
                <textarea
                  id="job_description"
                  rows="4"
                  v-model="jobForm.description"
                  required
                  class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-violet-500"
                  placeholder="Responsibilities, requirements, skills..."
                ></textarea>
              </div>
            </div>

            <div class="pt-4 border-t border-slate-100 flex justify-end">
              <button
                type="submit"
                :disabled="isCreatingJob"
                class="px-6 py-2.5 border border-transparent text-sm font-semibold rounded-lg text-white bg-violet-600 hover:bg-violet-500 focus-ring disabled:opacity-50 disabled:cursor-not-allowed shadow-sm transition-all duration-150"
              >
                <span v-if="isCreatingJob" class="flex items-center space-x-2">
                  <svg class="animate-spin h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  <span>Creating Job...</span>
                </span>
                <span v-else>Publish Job Listing</span>
              </button>
            </div>
          </form>
        </div>

        <!-- Job listings table -->
        <div v-else class="bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden">
          <div v-if="jobs.length === 0" class="text-center py-16 text-slate-500">
            <p class="font-bold text-slate-700 text-lg">No active jobs posted</p>
            <p class="text-sm text-slate-400 mt-1">Publish job openings to discover talent.</p>
          </div>
          <div v-else class="overflow-x-auto">
            <table class="min-w-full divide-y divide-slate-200">
              <thead class="bg-slate-50">
                <tr>
                  <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Job Title</th>
                  <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Location</th>
                  <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Salary</th>
                  <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Slots</th>
                  <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Status</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-slate-100">
                <tr v-for="job in jobs" :key="job.id">
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-semibold text-slate-900">{{ job.title }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">{{ job.location }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-bold text-emerald-600">${{ job.salary.toLocaleString() }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">{{ job.slots }}</td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span :class="[
                      job.status === 'approved' ? 'bg-emerald-50 text-emerald-700 border-emerald-200' : 'bg-amber-50 text-amber-700 border-amber-200',
                      'inline-flex items-center px-2 py-0.5 rounded-full text-xs font-semibold border'
                    ]">
                      {{ job.status || 'pending' }}
                    </span>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- Section 4: Employer Applicant Management (B6) -->
      <div v-show="activeSection === 'applicants'" class="space-y-6">
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between pb-6 border-b border-slate-200">
          <div>
            <h2 class="text-3xl font-extrabold text-slate-900 tracking-tight">Applicant Management</h2>
            <p class="mt-1 text-sm text-slate-500 font-medium">Review and process student applications.</p>
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
              v-model="applicantSearchQuery"
              type="text"
              class="block w-full pl-10 pr-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 placeholder-slate-400 text-slate-900 focus:outline-none focus:ring-2 focus:ring-violet-500 focus:bg-white transition-all duration-200"
              placeholder="Search by student name or job..."
            />
          </div>
          <div>
            <select
              v-model="applicantStatusFilter"
              class="block w-full px-3 py-2 border border-slate-200 rounded-lg text-sm bg-slate-50 text-slate-900 focus:outline-none focus:ring-2 focus:ring-violet-500 focus:bg-white transition-all duration-200"
            >
              <option value="all">All Statuses</option>
              <option value="applied">Applied (Pending)</option>
              <option value="approved">Approved</option>
              <option value="rejected">Rejected</option>
            </select>
          </div>
        </div>

        <!-- Loading state skeleton -->
        <div v-if="isLoadingApps" class="bg-white rounded-xl border border-slate-200 p-6 space-y-4 animate-pulse">
          <div class="h-6 bg-slate-200 rounded w-1/4"></div>
          <div class="space-y-3">
            <div class="h-10 bg-slate-100 rounded"></div>
            <div class="h-10 bg-slate-100 rounded"></div>
            <div class="h-10 bg-slate-100 rounded"></div>
          </div>
        </div>

        <!-- Empty state -->
        <div v-else-if="filteredApps.length === 0" class="bg-white text-center py-16 px-4 rounded-xl border border-slate-200 shadow-sm">
          <div class="inline-flex p-4 rounded-full bg-slate-50 border border-slate-100 text-slate-400 mb-4">
            <svg class="h-12 w-12" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
          </div>
          <h3 class="text-lg font-bold text-slate-900">No applicants found</h3>
          <p class="text-sm text-slate-500 mt-1 max-w-sm mx-auto font-medium">
            There are currently no incoming applications matching these parameters.
          </p>
        </div>

       <!-- Desktop Applicants Table -->
<div v-else class="hidden lg:block bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden">
  <table class="min-w-full divide-y divide-slate-200">
    <thead class="bg-slate-50">
      <tr>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Candidate</th>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Position Applied</th>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Contact Phone</th>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Applied Date</th>
        <th class="px-6 py-4 text-left text-xs font-bold text-slate-500 uppercase tracking-wider">Status</th>
        <th class="px-6 py-4 text-center text-xs font-bold text-slate-500 uppercase tracking-wider">Actions</th>
      </tr>
    </thead>
    <tbody class="bg-white divide-y divide-slate-100">
      <tr v-for="app in filteredApps" :key="app.id">
        <td class="px-6 py-4 whitespace-nowrap">
          <div class="flex items-center">
            <img v-if="app.student?.avatar_url" :src="app.student.avatar_url" class="h-8 w-8 rounded-full border border-slate-200 mr-3" />
            <div v-else class="h-8 w-8 rounded-full bg-slate-100 border border-slate-200 flex items-center justify-center text-slate-500 font-bold mr-3">
              {{ app.student?.full_name?.charAt(0) || 'S' }}
            </div>
            <div>
              <div class="text-sm font-semibold text-slate-900">{{ app.student?.full_name || 'N/A' }}</div>
              <div class="text-xs text-slate-500">Skills: {{ app.student?.skills || 'N/A' }}</div>
            </div>
          </div>
        </td>
        <td class="px-6 py-4 whitespace-nowrap text-sm font-semibold text-slate-800">
          {{ app.job?.title || jobTitleLookup(app.job_id) }}
        </td>
        <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-600 font-medium">
          {{ app.student?.phone || 'N/A' }}
        </td>
        <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500 font-medium">
          {{ formatDate(app.applied_at || app.id) }}
        </td>
        <td class="px-6 py-4 whitespace-nowrap">
          <span :class="[
            statusBadgeClass(app.status),
            'inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold border capitalize'
          ]">
            {{ app.status }}
          </span>
        </td>
        <td class="px-6 py-4 whitespace-nowrap text-center text-sm font-medium">
          <div class="flex justify-center items-center">
            <!-- 🌟 ĐÃ SỬA: Nếu trạng thái là pending thì hiện nút mở Modal xử lý và điền offer -->
            <button
              v-if="app.status?.toLowerCase() === 'pending'"
              @click="openReviewModal(app)"
              class="px-3 py-1.5 text-xs font-bold text-indigo-600 bg-indigo-50 hover:bg-indigo-100 rounded-lg transition-all duration-150"
            >
              Review & Offer
            </button>
            
            <!-- Nếu đã approved hoặc rejected thì hiện nút xem lại thông tin cũ thay vì khóa chết chữ Evaluated -->
            <button
              v-else
              @click="openReviewModal(app)"
              class="px-3 py-1.5 text-xs font-medium text-slate-500 bg-slate-100 hover:bg-slate-200 rounded-lg transition-all duration-150"
            >
              View Details
            </button>
          </div>
        </td>
      </tr>
    </tbody>
  </table>
</div>
        <!-- Mobile Applicants Cards -->
        <div class="lg:hidden space-y-4">
          <div
            v-for="app in filteredApps"
            :key="app.id"
            class="bg-white border border-slate-200 rounded-xl p-5 shadow-sm space-y-4"
          >
            <div class="flex items-center justify-between">
              <div class="flex items-center">
                <img v-if="app.student?.avatar_url" :src="app.student.avatar_url" class="h-10 w-10 rounded-full border border-slate-200 mr-3" />
                <div v-else class="h-10 w-10 bg-slate-100 rounded-full border border-slate-200 flex items-center justify-center font-bold text-slate-500 mr-3">
                  {{ app.student?.full_name?.charAt(0) || 'S' }}
                </div>
                <div>
                  <h4 class="text-sm font-bold text-slate-900">{{ app.student?.full_name || 'N/A' }}</h4>
                  <p class="text-xs text-slate-500">Skills: {{ app.student?.skills || 'N/A' }}</p>
                </div>
              </div>
              <span :class="[
                statusBadgeClass(app.status),
                'inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold border'
              ]">
                {{ app.status }}
              </span>
            </div>

            <div class="pt-2 border-t border-slate-100 text-xs font-medium text-slate-600 space-y-1">
              <div>Position: <span class="text-slate-900 font-bold">{{ jobTitleLookup(app.job_id) }}</span></div>
              <div>Contact: <span class="text-slate-900">{{ app.student?.phone }}</span></div>
              <div>Applied: <span class="text-slate-900">{{ formatDate(app.applied_at || app.id) }}</span></div>
            </div>

            <!-- Mobile actions -->
            <div v-if="app.status?.toLowerCase() === 'applied'" class="flex space-x-2 pt-2">
              <button
                @click="triggerConfirmModal(app, 'approved')"
                class="flex-1 text-center py-2 text-xs font-bold text-white bg-emerald-600 hover:bg-emerald-500 rounded-lg shadow-sm"
              >
                Approve
              </button>
              <button
                @click="triggerConfirmModal(app, 'rejected')"
                class="flex-1 text-center py-2 text-xs font-bold text-white bg-rose-600 hover:bg-rose-500 rounded-lg shadow-sm"
              >
                Reject
              </button>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Confirmation Modal Dialog -->
    <div v-if="showConfirmModal" class="fixed inset-0 z-50 flex items-center justify-center bg-slate-900 bg-opacity-50 p-4">
      <div class="bg-white rounded-2xl max-w-sm w-full p-6 shadow-2xl border border-slate-200 space-y-4">
        <h3 class="text-lg font-bold text-slate-900">Confirm Action</h3>
        <p class="text-sm text-slate-600">
          Are you sure you want to <span class="font-bold uppercase" :class="confirmAction === 'approved' ? 'text-emerald-600' : 'text-rose-600'">{{ confirmAction }}</span> the application for <span class="font-bold text-slate-900">{{ confirmTarget?.student?.full_name }}</span>?
        </p>
        <div class="flex justify-end space-x-3 pt-2">
          <button
            @click="showConfirmModal = false"
            :disabled="isReviewing"
            class="px-4 py-2 border border-slate-200 hover:bg-slate-50 text-slate-700 font-semibold rounded-lg text-sm transition-colors"
          >
            Cancel
          </button>
          <button
            @click="handleReviewApplication"
            :disabled="isReviewing"
            :class="[
              confirmAction === 'approved' ? 'bg-emerald-600 hover:bg-emerald-500' : 'bg-rose-600 hover:bg-rose-500',
              'px-4 py-2 text-white font-bold rounded-lg text-sm shadow-sm flex items-center space-x-1'
            ]"
          >
            <span v-if="isReviewing">Processing...</span>
            <span v-else>Confirm</span>
          </button>
        </div>
      </div>
    </div>
  </div>
  <div v-if="selectedApp" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm">
  <div class="bg-white rounded-2xl border border-slate-200 shadow-xl max-w-lg w-full overflow-hidden flex flex-col max-h-[90vh] text-left animate-in fade-in zoom-in-95 duration-150">
    
    <div class="p-6 border-b border-slate-100 flex justify-between items-center bg-slate-50">
      <div>
        <h3 class="text-base font-bold text-slate-900">Đánh giá hồ sơ & Gửi Offer</h3>
        <p class="text-xs text-slate-500 font-medium mt-0.5">Ứng viên: <span class="text-slate-700 font-bold">{{ selectedApp.student?.full_name || 'N/A' }}</span></p>
      </div>
      <button @click="closeModal" class="text-slate-400 hover:text-slate-600 font-bold text-xl">&times;</button>
    </div>

    <div class="p-6 space-y-4 overflow-y-auto flex-1">
      <div>
        <label class="text-xs font-bold text-slate-400 uppercase tracking-wider block mb-1">Lời nhắn từ ứng viên (Cover Note)</label>
        <div class="p-3 bg-blue-50/60 border border-blue-100 rounded-xl text-xs font-medium text-slate-700 whitespace-pre-line">
          {{ selectedApp.cover_note || 'Ứng viên không để lại lời nhắn.' }}
        </div>
      </div>

      <div>
        <label class="text-xs font-bold text-slate-400 uppercase tracking-wider block mb-2">Quyết định</label>
        <div class="grid grid-cols-2 gap-3">
          <button 
            type="button"
            @click="reviewStatus = 'approved'"
            :class="[reviewStatus === 'approved' ? 'border-emerald-500 bg-emerald-50 text-emerald-700 ring-2 ring-emerald-500/20' : 'border-slate-200 bg-white text-slate-700 hover:bg-slate-50']"
            class="py-2 px-4 text-xs font-bold border rounded-xl transition-all flex items-center justify-center space-x-1"
          >
            <span>👍 Chấp nhận & Gửi Offer</span>
          </button>
          <button 
            type="button"
            @click="reviewStatus = 'rejected'"
            :class="[reviewStatus === 'rejected' ? 'border-rose-500 bg-rose-50 text-rose-700 ring-2 ring-rose-500/20' : 'border-slate-200 bg-white text-slate-700 hover:bg-slate-50']"
            class="py-2 px-4 text-xs font-bold border rounded-xl transition-all flex items-center justify-center space-x-1"
          >
            <span>👎 Từ chối đơn</span>
          </button>
        </div>
      </div>

      <div v-if="reviewStatus === 'approved'" class="p-4 bg-emerald-50/30 border border-emerald-100 rounded-xl space-y-3 animate-in fade-in duration-150">
        <h4 class="text-xs font-bold text-emerald-800">✉️ Chi tiết Offer gửi Sinh viên</h4>
        
        <div class="grid grid-cols-2 gap-3">
          <div>
            <label class="text-[11px] font-bold text-slate-600 block mb-1">Mức lương Offer</label>
            <input 
              v-model="offerForm.salary"
              type="text" 
              placeholder="Ví dụ: 15,000,000 VND"
              class="w-full text-xs px-3 py-2 border border-slate-200 rounded-lg bg-white text-slate-800 focus:outline-none focus:border-emerald-500"
            />
          </div>
          <div>
            <label class="text-[11px] font-bold text-slate-600 block mb-1">Ngày đi làm dự kiến</label>
            <input 
              v-model="offerForm.startDate"
              type="date" 
              class="w-full text-xs px-3 py-2 border border-slate-200 rounded-lg bg-white text-slate-800 focus:outline-none focus:border-emerald-500"
            />
          </div>
        </div>

        <div>
          <label class="text-[11px] font-bold text-slate-600 block mb-1">Lời nhắn chào mừng từ HR</label>
          <textarea 
            v-model="offerForm.message"
            rows="2"
            placeholder="Chào mừng bạn gia nhập đội ngũ..."
            class="w-full text-xs p-2.5 border border-slate-200 rounded-lg bg-white text-slate-800 focus:outline-none focus:border-emerald-500"
          ></textarea>
        </div>
      </div>
    </div>

    <div class="p-4 bg-slate-50 border-t border-slate-100 flex items-center justify-end space-x-2">
      <button 
        @click="closeModal" 
        class="px-4 py-2 text-xs font-bold text-slate-600 hover:bg-slate-100 rounded-lg transition-all"
      >
        Đóng
      </button>
      <button 
        @click="submitReview"
        :disabled="!reviewStatus || isSubmitting"
        class="px-4 py-2 text-xs font-bold text-white bg-indigo-600 hover:bg-indigo-700 disabled:opacity-50 rounded-lg shadow-sm transition-all"
      >
        {{ isSubmitting ? 'Đang gửi...' : 'Xác nhận phản hồi' }}
      </button>
    </div>

  </div>
</div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { useApi } from '~/composables/useApi'

definePageMeta({
  middleware: 'auth'
})

// Navigation setups
const activeSection = ref('dashboard')

const navItems = [
  { id: 'dashboard', name: 'Dashboard' },
  { id: 'profile', name: 'Company Profile' },
  { id: 'jobs', name: 'Jobs' },
  { id: 'applicants', name: 'Applicants' }
]

// API Client references
const api = useApi()
const jobs = ref<any[]>([])
const applications = ref<any[]>([])
const isLoadingJobs = ref(false)
const isLoadingApps = ref(false)
const isSavingProfile = ref(false)
const isCreatingJob = ref(false)
const showCreateForm = ref(false)
const feedback = ref<{ type: 'success' | 'error'; message: string } | null>(null)

// Confirmation Modal States
const showConfirmModal = ref(false)
const confirmTarget = ref<any | null>(null)
const confirmAction = ref<'approved' | 'rejected'>('approved')
const isReviewing = ref(false)

// Query filters
const applicantSearchQuery = ref('')
const applicantStatusFilter = ref('all')


// Trạng thái quản lý đóng mở form và file logo của Business
const isEditing = ref(false)
const logoFileSelected = ref<File | null>(null)
const logoPreview = ref<string | null>(null)


const selectedApp = ref<any>(null)
const reviewStatus = ref<'approved' | 'rejected' | null>(null)
const isSubmitting = ref(false)

// Form thông tin Offer gửi đính kèm
const offerForm = ref({
  salary: '',
  startDate: '',
  message: ''
})
const openReviewModal = (app: any) => {
  selectedApp.value = app
  reviewStatus.value = app.status === 'pending' ? 'approved' : app.status // Mặc định mở ra chọn luôn approved cho tiện
  offerForm.value = {
    salary: app.offer_salary || '',
    startDate: app.offer_start_date || '',
    message: app.offer_message || ''
  }
}

const closeModal = () => {
  selectedApp.value = null
  reviewStatus.value = null
}

const submitReview = async () => {
  if (!selectedApp.value || !reviewStatus.value) return

  isSubmitting.value = true
  try {
    const payload = {
      application_id: selectedApp.value.id,
      status: reviewStatus.value,
      offer_salary: reviewStatus.value === 'approved' ? offerForm.value.salary : '',
      offer_start_date: reviewStatus.value === 'approved' ? offerForm.value.startDate : '',
      offer_message: reviewStatus.value === 'approved' ? offerForm.value.message : ''
    }

    await api.put('/api/jobs/review-application', payload)
    
    // 🌟 TỰ ĐỘNG ĐỒNG BỘ: Quét và gọi đúng tên hàm fetch hiện có trong file của Chanh
    const anyWindow = window as any
    if (typeof anyWindow.fetchApplications === 'function') {
      await anyWindow.fetchApplications()
    } else if (typeof anyWindow.fetchApplicants === 'function') {
      await anyWindow.fetchApplicants()
    } else if (typeof anyWindow.fetchEmployerApplications === 'function') {
      await anyWindow.fetchEmployerApplications()
    } else {
      // Giải pháp an toàn cuối cùng nếu không tìm thấy hàm nào: Reload lại trang
      window.location.reload()
    }
    
    closeModal()
  } catch (err) {
    console.error('Lỗi khi cập nhật trạng thái hoặc gửi offer:', err)
    alert('Có lỗi xảy ra khi phê duyệt.')
  } finally {
    isSubmitting.value = false
  }
}

  const onLogoFileChange = (e: Event) => {
  const target = e.target as HTMLInputElement
  const file = target.files?.[0]
  if (file) {
    logoFileSelected.value = file
    logoPreview.value = URL.createObjectURL(file) // Tạo preview cục bộ
  }
}
  const fetchProfile = async () => {
  try {
    const res = await api.get('/api/profile/business')
    const data = res && res.data ? res.data : res
    if (data) {
      profileForm.company_name = data.company_name || ''
      profileForm.phone = data.phone || ''
      profileForm.address = data.address || ''
      profileForm.logo_url = data.logo_url || ''
    }
  } catch (err) {
    console.error('Error fetching company profile:', err)
  }
}
// Forms Setup
const profileForm = reactive({
  company_name: '',
  phone: '',
  address: '',
  logo_url: ''
})

const jobForm = reactive({
  title: '',
  description: '',
  location: '',
  salary: 0,
  slots: 1,
  working_date: ''
})

// Metrics and compute operations
const jobTitleLookup = (jobId: number): string => {
  const job = jobs.value.find(j => j.id === jobId)
  return job ? job.title : `Position #${jobId || ''}`
}

const fillRatio = computed(() => {
  if (jobs.value.length === 0) return 0
  const approvedTotal = applications.value.filter(app => app.status?.toLowerCase() === 'approved').length
  const slotsTotal = jobs.value.reduce((acc, j) => acc + (j.slots || 1), 0)
  return Math.min(100, Math.round((approvedTotal / slotsTotal) * 100))
})

const metricsCards = computed(() => {
  const total = applications.value.length
  const pending = applications.value.filter(app => app.status?.toLowerCase() === 'applied').length
  const approved = applications.value.filter(app => app.status?.toLowerCase() === 'approved').length
  const rejected = applications.value.filter(app => app.status?.toLowerCase() === 'rejected').length

  return [
    { title: 'Total Applications', value: total, label: 'Applications', color: 'bg-slate-50 border-slate-200 text-slate-700' },
    { title: 'Pending Applications', value: pending, label: 'Awaiting Review', color: 'bg-amber-50 border-amber-200 text-amber-700' },
    { title: 'Approved Applications', value: approved, label: 'Accepted', color: 'bg-emerald-50 border-emerald-200 text-emerald-700' },
    { title: 'Rejected Applications', value: rejected, label: 'Declined', color: 'bg-rose-50 border-rose-200 text-rose-700' }
  ]
})

const filteredApps = computed(() => {
  return applications.value.filter(app => {
    const studentName = app.student?.full_name?.toLowerCase() || ''
    const jobTitle = jobTitleLookup(app.job_id).toLowerCase()

    const matchesSearch = !applicantSearchQuery.value ||
      studentName.includes(applicantSearchQuery.value.toLowerCase()) ||
      jobTitle.includes(applicantSearchQuery.value.toLowerCase())

    const matchesStatus = applicantStatusFilter.value === 'all' ||
      app.status?.toLowerCase() === applicantStatusFilter.value.toLowerCase()

    return matchesSearch && matchesStatus
  })
})

const statusBadgeClass = (status: string): string => {
  const norm = status?.toLowerCase() || ''
  if (norm === 'approved') return 'bg-emerald-50 border-emerald-200 text-emerald-700'
  if (norm === 'rejected') return 'bg-rose-50 border-rose-200 text-rose-700'
  return 'bg-slate-50 border-slate-200 text-slate-700'
}

const formatDate = (dateVal: any): string => {
  return 'Jun 24, 2026'
}

// API Methods
const fetchJobs = async () => {
  isLoadingJobs.value = true
  try {
    const res = await api.get('/api/jobs')
    jobs.value = res.data || []
  } catch (err) {
    console.error('Error fetching jobs:', err)
  } finally {
    isLoadingJobs.value = false
  }
}

const fetchApplications = async () => {
  isLoadingApps.value = true
  try {
    const res = await api.get('/api/applications/employer')
    applications.value = res.data || []
  } catch (err) {
    console.error('Error fetching employer applications:', err)
  } finally {
    isLoadingApps.value = false
  }
}

const handleUpdateProfile = async () => {
  isSavingProfile.value = true
  feedback.value = null

  try {
    const formData = new FormData()
    formData.append('company_name', profileForm.company_name)
    formData.append('phone', profileForm.phone)
    formData.append('address', profileForm.address)

    // Nếu doanh nghiệp có chọn file logo mới thì đính kèm vào payload
    if (logoFileSelected.value) {
      formData.append('logo', logoFileSelected.value)
    }

    const res = await api.put('/api/profile/business', formData)
    
    feedback.value = {
      type: 'success',
      message: res.message || '🎉 Company profile updated successfully!'
    }

    // Đồng bộ lại đường dẫn logo mới nhất trả về từ server để hiển thị ngay
    const updatedData = res.data ? res.data : res
    if (updatedData) {
      profileForm.logo_url = updatedData.LogoUrl || updatedData.logo_url || profileForm.logo_url
    }

    // Tắt chế độ sửa, quay về chế độ xem và dọn dẹp file tạm
    isEditing.value = false
    logoFileSelected.value = null
  } catch (err: any) {
    feedback.value = {
      type: 'error',
      message: err.response?._data?.message || 'Failed to update company profile.'
    }
  } finally {
    isSavingProfile.value = false
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const handleCreateJob = async () => {
  isCreatingJob.value = true
  feedback.value = null

  try {
    const res = await api.post('/api/jobs', {
      title: jobForm.title,
      description: jobForm.description,
      location: jobForm.location,
      salary: Number(jobForm.salary),
      slots: Number(jobForm.slots),
      working_date: jobForm.working_date
    })

    feedback.value = {
      type: 'success',
      message: res.message || '🎉 Job opening created successfully!'
    }

    // Reset Form fields
    jobForm.title = ''
    jobForm.description = ''
    jobForm.location = ''
    jobForm.salary = 0
    jobForm.slots = 1
    jobForm.working_date = ''

    // Return to main jobs overview
    showCreateForm.value = false
    await fetchJobs()
  } catch (err: any) {
    feedback.value = {
      type: 'error',
      message: err.response?._data?.message || 'Failed to publish job opening.'
    }
  } finally {
    isCreatingJob.value = false
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const triggerConfirmModal = (app: any, status: 'approved' | 'rejected') => {
  confirmTarget.value = app
  confirmAction.value = status
  showConfirmModal.value = true
}

const handleReviewApplication = async () => {
  if (!confirmTarget.value) return

  isReviewing.value = true
  feedback.value = null

  const appId = confirmTarget.value.id
  const targetStatus = confirmAction.value

  // Optimistic UI updates
  const originalAppsState = [...applications.value]
  const targetIndex = applications.value.findIndex(app => app.id === appId)
  if (targetIndex !== -1) {
    applications.value[targetIndex].status = targetStatus
  }

  try {
    const res = await api.put('/api/jobs/review-application', {
      application_id: Number(appId),
      status: targetStatus
    })

    feedback.value = {
      type: 'success',
      message: res.message || '🎉 Candidate status evaluated successfully!'
    }
  } catch (err: any) {
    // Revert optimistic update
    applications.value = originalAppsState
    feedback.value = {
      type: 'error',
      message: err.response?._data?.message || 'Failed to update candidate status.'
    }
  } finally {
    isReviewing.value = false
    showConfirmModal.value = false
    confirmTarget.value = null
    // Sync-refresh list data
    await fetchApplications()
  }
}

onMounted(() => {
  fetchJobs()
  fetchApplications()
  fetchProfile() // <-- Gọi nạp dữ liệu cũ của Công ty ở đây nha Chanh
})
</script>
