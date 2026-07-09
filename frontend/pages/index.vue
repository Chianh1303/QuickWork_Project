<template>
  <div id="home" class="overflow-hidden bg-slate-950 text-white">
    <section class="relative border-b border-white/10">
      <div class="absolute inset-0 bg-[radial-gradient(circle_at_18%_18%,rgba(34,211,238,0.22),transparent_26rem),radial-gradient(circle_at_82%_16%,rgba(16,185,129,0.14),transparent_30rem),linear-gradient(135deg,#020617_0%,#0f172a_48%,#062f3a_100%)]"></div>
      <div class="absolute left-1/2 top-20 h-64 w-64 -translate-x-1/2 rounded-full bg-cyan-400/10 blur-3xl animate-pulse"></div>

      <div class="relative mx-auto grid max-w-7xl gap-10 px-4 py-20 sm:px-6 lg:grid-cols-[0.95fr_1.05fr] lg:items-center lg:px-8 lg:py-28">
        <div class="space-y-7 text-center lg:text-left">
          <div class="inline-flex items-center gap-2 rounded-full border border-cyan-400/20 bg-cyan-400/10 px-3 py-1 text-sm font-bold text-cyan-100">
            <span class="h-2 w-2 rounded-full bg-cyan-300"></span>
            Student jobs and hiring, in one platform
          </div>
          <div class="space-y-5">
            <h1 class="text-4xl font-extrabold leading-[0.98] tracking-tight sm:text-6xl lg:text-7xl">
              Find Part-time & Internship Jobs Faster
            </h1>
            <p class="mx-auto max-w-2xl text-base font-medium leading-8 text-slate-300 sm:text-lg lg:mx-0">
              QuickWork connects students with businesses quickly, safely and efficiently through job search, applications, chat, attendance, wallet and reviews.
            </p>
          </div>
          <div class="flex flex-col justify-center gap-3 sm:flex-row lg:justify-start">
            <a href="#explore-jobs" class="focus-ring inline-flex items-center justify-center rounded-xl bg-cyan-400 px-7 py-3.5 text-base font-extrabold text-slate-950 shadow-lg shadow-cyan-500/20 hover:bg-cyan-300">
              Explore Jobs
            </a>
            <NuxtLink to="/employer-register" class="focus-ring inline-flex items-center justify-center rounded-xl border border-white/10 bg-white/10 px-7 py-3.5 text-base font-bold text-white hover:bg-white/15">
              Create Employer Account
            </NuxtLink>
          </div>
        </div>

        <div class="relative">
          <div class="absolute -inset-5 rounded-[2rem] bg-cyan-400/10 blur-3xl"></div>
          <div class="relative overflow-hidden rounded-3xl border border-white/10 bg-slate-900/80 p-5 shadow-2xl shadow-slate-950/60 backdrop-blur">
            <div class="flex items-center justify-between border-b border-white/10 pb-4">
              <div>
                <p class="text-xs font-bold uppercase tracking-wider text-cyan-200">QuickWork matching</p>
                <h2 class="mt-1 text-lg font-extrabold text-white">Live job search workspace</h2>
              </div>
              <span class="rounded-full border border-emerald-400/20 bg-emerald-400/10 px-3 py-1 text-xs font-bold text-emerald-200">Verified</span>
            </div>

            <div class="grid gap-3 py-4 sm:grid-cols-3">
              <div v-for="stat in heroStats" :key="stat.label" class="rounded-xl border border-white/10 bg-white/[0.06] p-4">
                <p class="text-2xl font-extrabold text-white">{{ stat.value }}</p>
                <p class="mt-1 text-xs font-bold uppercase tracking-wide text-slate-500">{{ stat.label }}</p>
              </div>
            </div>

            <div class="space-y-3">
              <div class="rounded-2xl border border-cyan-400/20 bg-slate-950/80 p-4">
                <div class="flex items-start justify-between gap-4">
                  <div>
                    <span class="rounded-full bg-cyan-400/10 px-2.5 py-1 text-xs font-bold text-cyan-200">FPT Software</span>
                    <h3 class="mt-3 text-lg font-extrabold text-white">Backend Intern</h3>
                    <p class="mt-2 text-sm font-medium text-slate-400">Hybrid · Student-friendly · Start this month</p>
                  </div>
                  <span class="rounded-lg bg-emerald-400/10 px-3 py-1 text-sm font-extrabold text-emerald-200">9.500.000 VND</span>
                </div>
                <div class="mt-4 grid grid-cols-3 gap-2">
                  <span class="rounded-lg bg-white/[0.06] px-3 py-2 text-center text-xs font-bold text-slate-300">Chat</span>
                  <span class="rounded-lg bg-white/[0.06] px-3 py-2 text-center text-xs font-bold text-slate-300">Offer</span>
                  <span class="rounded-lg bg-white/[0.06] px-3 py-2 text-center text-xs font-bold text-slate-300">Wallet</span>
                </div>
              </div>
              <div class="grid grid-cols-2 gap-3">
                <div class="rounded-xl border border-white/10 bg-white/[0.05] p-4">
                  <p class="text-xs font-bold uppercase text-slate-500">Attendance</p>
                  <p class="mt-1 text-sm font-bold text-slate-200">Check-in ready</p>
                </div>
                <div class="rounded-xl border border-white/10 bg-white/[0.05] p-4">
                  <p class="text-xs font-bold uppercase text-slate-500">Reviews</p>
                  <p class="mt-1 text-sm font-bold text-slate-200">Trust after work</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <section id="about" class="px-4 py-16 sm:px-6 lg:px-8">
      <div class="mx-auto max-w-7xl">
        <SectionHeading eyebrow="Why QuickWork?" title="Built for students and employers who move fast" />
        <div class="mt-10 grid gap-5 md:grid-cols-2 xl:grid-cols-4">
          <InfoCard v-for="item in introCards" :key="item.title" v-bind="item" />
        </div>
      </div>
    </section>

    <section id="features" class="border-y border-white/10 bg-slate-900/45 px-4 py-16 sm:px-6 lg:px-8">
      <div class="mx-auto max-w-7xl">
        <SectionHeading eyebrow="Platform Features" title="Everything needed from discovery to payment" />
        <div class="mt-10 grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
          <FeatureCard v-for="feature in features" :key="feature.title" v-bind="feature" />
        </div>
      </div>
    </section>

    <ExploreJobsSection :show-hero="false" :standalone="false" section-id="explore-jobs" />

    <section class="border-y border-white/10 bg-slate-900/45 px-4 py-16 sm:px-6 lg:px-8">
      <div class="mx-auto max-w-7xl">
        <SectionHeading eyebrow="How It Works" title="A clear workflow for every application" />
        <div class="mt-10 grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
          <div v-for="(step, index) in timeline" :key="step.title" class="rounded-2xl border border-white/10 bg-white/[0.04] p-5">
            <div class="flex h-10 w-10 items-center justify-center rounded-xl bg-cyan-400 text-sm font-extrabold text-slate-950">{{ index + 1 }}</div>
            <h3 class="mt-4 text-base font-extrabold text-white">{{ step.title }}</h3>
            <p class="mt-2 text-sm font-medium leading-6 text-slate-400">{{ step.text }}</p>
          </div>
        </div>
      </div>
    </section>

    <section class="px-4 py-16 sm:px-6 lg:px-8">
      <div class="mx-auto max-w-7xl">
        <div class="grid gap-4 md:grid-cols-4">
          <div v-for="stat in stats" :key="stat.label" class="rounded-2xl border border-cyan-400/15 bg-cyan-400/10 p-6 text-center">
            <p class="text-4xl font-extrabold text-white">{{ stat.value }}</p>
            <p class="mt-2 text-sm font-bold uppercase tracking-wide text-cyan-200">{{ stat.label }}</p>
          </div>
        </div>
      </div>
    </section>

    <section class="border-y border-white/10 bg-slate-900/45 px-4 py-16 sm:px-6 lg:px-8">
      <div class="mx-auto max-w-7xl">
        <SectionHeading eyebrow="Testimonials" title="Trusted by students and local teams" />
        <div class="mt-10 grid gap-5 md:grid-cols-3">
          <div v-for="testimonial in testimonials" :key="testimonial.name" class="rounded-2xl border border-white/10 bg-white/[0.04] p-6">
            <p class="text-sm font-medium leading-7 text-slate-300">"{{ testimonial.quote }}"</p>
            <div class="mt-5">
              <p class="font-extrabold text-white">{{ testimonial.name }}</p>
              <p class="text-sm font-semibold text-cyan-200">{{ testimonial.role }}</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <footer id="contact" class="px-4 py-12 sm:px-6 lg:px-8">
      <div class="mx-auto grid max-w-7xl gap-8 rounded-3xl border border-white/10 bg-slate-900/80 p-8 md:grid-cols-[1.2fr_0.8fr_0.8fr_1fr]">
        <div>
          <div class="flex items-center gap-3">
            <span class="flex h-10 w-10 items-center justify-center rounded-lg bg-cyan-400 font-extrabold text-slate-950">Q</span>
            <span class="text-xl font-extrabold">QuickWork</span>
          </div>
          <p class="mt-4 max-w-sm text-sm font-medium leading-7 text-slate-400">
            A modern recruitment workspace for student jobs, internships, hiring conversations, attendance, wallet and reviews.
          </p>
        </div>
        <FooterGroup title="Quick Links" :items="['Home', 'Explore Jobs', 'Features', 'Contact']" />
        <FooterGroup title="Users" :items="['Student', 'Business', 'Register', 'Login']" />
        <div>
          <h3 class="font-extrabold text-white">Contact</h3>
          <div class="mt-4 space-y-2 text-sm font-medium text-slate-400">
            <p>support@quickwork.vn</p>
            <p>0123 456 789</p>
            <p>Ho Chi Minh City, Vietnam</p>
          </div>
          <div class="mt-5 flex gap-3">
            <span v-for="social in ['Facebook', 'LinkedIn', 'GitHub']" :key="social" class="rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-xs font-bold text-slate-300">{{ social }}</span>
          </div>
        </div>
      </div>
      <p class="mx-auto mt-6 max-w-7xl text-center text-sm font-medium text-slate-500">© 2026 QuickWork. All rights reserved.</p>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { defineComponent, h } from 'vue'
import ExploreJobsSection from '~/components/home/ExploreJobsSection.vue'

const heroStats = [
  { value: '300+', label: 'Jobs' },
  { value: '500+', label: 'Students' },
  { value: '150+', label: 'Businesses' }
]

const introCards = [
  { icon: 'S', title: 'Student', text: 'Find internships and part-time jobs that match your schedule.' },
  { icon: 'B', title: 'Business', text: 'Recruit students quickly with applications and offers in one place.' },
  { icon: 'C', title: 'Real-time Chat', text: 'Instant communication using WebSocket-powered conversations.' },
  { icon: 'W', title: 'Wallet & Payment', text: 'Simulated payroll after job completion and confirmation.' }
]

const features = [
  ['Job Search', 'Browse verified openings from local businesses.'],
  ['Smart Filter', 'Find jobs by keyword, location, type and salary.'],
  ['Application Tracking', 'Follow every application status clearly.'],
  ['Offer Management', 'Review offers and respond from one workspace.'],
  ['Attendance', 'Check in and check out during accepted work.'],
  ['Wallet', 'Track simulated earnings and payment history.'],
  ['Review System', 'Build trust after each completed job.'],
  ['Real-time Chat', 'Talk with HR or candidates instantly.']
].map(([title, text], index) => ({ icon: String(index + 1), title, text }))

const timeline = [
  { title: 'Register', text: 'Create a student or employer account.' },
  { title: 'Apply Job', text: 'Students apply to verified openings.' },
  { title: 'Business Reviews', text: 'Employers review applications and profiles.' },
  { title: 'Receive Offer', text: 'Approved students receive offer details.' },
  { title: 'Work', text: 'Accepted students begin the job.' },
  { title: 'Attendance', text: 'Work sessions are tracked with check-in.' },
  { title: 'Salary Payment', text: 'Payment is released after confirmation.' },
  { title: 'Review', text: 'Both sides leave reviews after completion.' }
]

const stats = [
  { value: '150+', label: 'Businesses' },
  { value: '500+', label: 'Students' },
  { value: '1200+', label: 'Applications' },
  { value: '300+', label: 'Jobs' }
]

const testimonials = [
  { name: 'Minh Anh', role: 'Student', quote: 'QuickWork helped me find a flexible internship and manage the offer without messy messages.' },
  { name: 'Thanh Nguyen', role: 'HR Executive', quote: 'The applicant pipeline and chat flow make student hiring much easier for our team.' },
  { name: 'Quoc Bao', role: 'Student Developer', quote: 'I liked having attendance, payment status and reviews inside the same dashboard.' }
]

const SectionHeading = defineComponent({
  props: { eyebrow: String, title: String },
  setup(props) {
    return () => h('div', { class: 'mx-auto max-w-3xl text-center' }, [
      h('p', { class: 'text-sm font-bold uppercase tracking-[0.24em] text-cyan-300' }, props.eyebrow),
      h('h2', { class: 'mt-3 text-3xl font-extrabold tracking-tight text-white sm:text-4xl' }, props.title)
    ])
  }
})

const InfoCard = defineComponent({
  props: { icon: String, title: String, text: String },
  setup(props) {
    return () => h('div', { class: 'rounded-2xl border border-white/10 bg-white/[0.04] p-6 transition-all duration-200 hover:-translate-y-1 hover:border-cyan-300/40 hover:bg-white/[0.07]' }, [
      h('div', { class: 'flex h-11 w-11 items-center justify-center rounded-xl bg-cyan-400 text-sm font-extrabold text-slate-950' }, props.icon),
      h('h3', { class: 'mt-5 text-lg font-extrabold text-white' }, props.title),
      h('p', { class: 'mt-2 text-sm font-medium leading-6 text-slate-400' }, props.text)
    ])
  }
})

const FeatureCard = defineComponent({
  props: { icon: String, title: String, text: String },
  setup(props) {
    return () => h('div', { class: 'rounded-2xl border border-white/10 bg-slate-950/60 p-5 transition-all duration-200 hover:-translate-y-1 hover:border-cyan-300/40 hover:shadow-lg hover:shadow-cyan-950/30' }, [
      h('div', { class: 'flex h-10 w-10 items-center justify-center rounded-lg bg-white/10 text-sm font-extrabold text-cyan-200' }, props.icon),
      h('h3', { class: 'mt-4 font-extrabold text-white' }, props.title),
      h('p', { class: 'mt-2 text-sm font-medium leading-6 text-slate-400' }, props.text)
    ])
  }
})

const FooterGroup = defineComponent({
  props: { title: String, items: Array },
  setup(props) {
    return () => h('div', [
      h('h3', { class: 'font-extrabold text-white' }, props.title),
      h('div', { class: 'mt-4 space-y-2 text-sm font-medium text-slate-400' }, (props.items as string[]).map(item => h('p', item)))
    ])
  }
})
</script>
