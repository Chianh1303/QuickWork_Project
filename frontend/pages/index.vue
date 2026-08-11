<template>
  <div id="home" class="overflow-hidden bg-slate-950 text-white">
    <section class="relative border-b border-white/10">
      <div class="absolute inset-0 bg-[radial-gradient(circle_at_18%_18%,rgba(34,211,238,0.22),transparent_26rem),radial-gradient(circle_at_82%_16%,rgba(16,185,129,0.14),transparent_30rem),linear-gradient(135deg,#020617_0%,#0f172a_48%,#062f3a_100%)]"></div>
      <div class="absolute left-1/2 top-20 h-64 w-64 -translate-x-1/2 rounded-full bg-cyan-400/10 blur-3xl animate-pulse"></div>

      <div class="relative mx-auto grid max-w-7xl gap-10 px-4 py-20 sm:px-6 lg:grid-cols-[0.95fr_1.05fr] lg:items-center lg:px-8 lg:py-28">
        <div class="space-y-7 text-center lg:text-left">
          <div class="inline-flex items-center gap-2 rounded-full border border-cyan-400/20 bg-cyan-400/10 px-3 py-1 text-sm font-bold text-cyan-100">
            <span class="h-2 w-2 rounded-full bg-cyan-300"></span>
            Nền tảng việc làm & tuyển dụng sinh viên hàng đầu
          </div>
          <div class="space-y-5">
            <h1 class="text-4xl font-extrabold leading-[0.98] tracking-tight sm:text-6xl lg:text-7xl">
              Tìm Việc Bán Thời Gian & Thực Tập Nhanh Chóng
            </h1>
            <p class="mx-auto max-w-2xl text-base font-medium leading-8 text-slate-300 sm:text-lg lg:mx-0">
              QuickWork kết nối sinh viên với doanh nghiệp một cách nhanh chóng, an toàn và hiệu quả thông qua hệ thống tìm việc, ứng tuyển, nhắn tin, chấm công, ví tiền và đánh giá.
            </p>
          </div>
          <div class="flex flex-col justify-center gap-3 sm:flex-row lg:justify-start">
            <a href="#explore-jobs" class="focus-ring inline-flex items-center justify-center rounded-xl bg-cyan-400 px-7 py-3.5 text-base font-extrabold text-slate-950 shadow-lg shadow-cyan-500/20 hover:bg-cyan-300">
              Khám Phá Việc Làm
            </a>
            <NuxtLink to="/employer-register" class="focus-ring inline-flex items-center justify-center rounded-xl border border-white/10 bg-white/10 px-7 py-3.5 text-base font-bold text-white hover:bg-white/15">
              Đăng Ký Doanh Nghiệp
            </NuxtLink>
          </div>
        </div>

        <div class="relative">
          <div class="absolute -inset-5 rounded-[2rem] bg-cyan-400/10 blur-3xl"></div>
          <div class="relative overflow-hidden rounded-3xl border border-white/10 bg-slate-900/80 p-5 shadow-2xl shadow-slate-950/60 backdrop-blur">
            <div class="flex items-center justify-between border-b border-white/10 pb-4">
              <div>
                <p class="text-xs font-bold uppercase tracking-wider text-cyan-200">Kết nối QuickWork</p>
                <h2 class="mt-1 text-lg font-extrabold text-white">Không gian tìm việc trực tuyến</h2>
              </div>
              <span class="rounded-full border border-emerald-400/20 bg-emerald-400/10 px-3 py-1 text-xs font-bold text-emerald-200">Đã xác thực KYB</span>
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
                    <h3 class="mt-3 text-lg font-extrabold text-white">Thực tập sinh Backend</h3>
                    <p class="mt-2 text-sm font-medium text-slate-400">Bán thời gian · Phù hợp sinh viên · Nhận việc ngay</p>
                  </div>
                  <span class="rounded-lg bg-emerald-400/10 px-3 py-1 text-sm font-extrabold text-emerald-200">9.500.000 VNĐ</span>
                </div>
                <div class="mt-4 grid grid-cols-3 gap-2">
                  <span class="rounded-lg bg-white/[0.06] px-3 py-2 text-center text-xs font-bold text-slate-300">Nhắn tin</span>
                  <span class="rounded-lg bg-white/[0.06] px-3 py-2 text-center text-xs font-bold text-slate-300">Nhận Offer</span>
                  <span class="rounded-lg bg-white/[0.06] px-3 py-2 text-center text-xs font-bold text-slate-300">Ví lương</span>
                </div>
              </div>
              <div class="grid grid-cols-2 gap-3">
                <div class="rounded-xl border border-white/10 bg-white/[0.05] p-4">
                  <p class="text-xs font-bold uppercase text-slate-500">Chấm công</p>
                  <p class="mt-1 text-sm font-bold text-slate-200">Điểm danh 1-Chạm</p>
                </div>
                <div class="rounded-xl border border-white/10 bg-white/[0.05] p-4">
                  <p class="text-xs font-bold uppercase text-slate-500">Đánh giá</p>
                  <p class="mt-1 text-sm font-bold text-slate-200">Xây dựng uy tín</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <section id="about" class="px-4 py-16 sm:px-6 lg:px-8">
      <div class="mx-auto max-w-7xl">
        <SectionHeading eyebrow="Tại sao chọn QuickWork?" title="Giải pháp tối ưu cho sinh viên và nhà tuyển dụng" />
        <div class="mt-10 grid gap-5 md:grid-cols-2 xl:grid-cols-4">
          <InfoCard v-for="item in introCards" :key="item.title" v-bind="item" />
        </div>
      </div>
    </section>

    <section id="features" class="border-y border-white/10 bg-slate-900/45 px-4 py-16 sm:px-6 lg:px-8">
      <div class="mx-auto max-w-7xl">
        <SectionHeading eyebrow="Tính năng nổi bật" title="Trọn bộ công cụ từ tìm việc đến nhận lương" />
        <div class="mt-10 grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
          <FeatureCard v-for="feature in features" :key="feature.title" v-bind="feature" />
        </div>
      </div>
    </section>

    <ExploreJobsSection :show-hero="false" :standalone="false" section-id="explore-jobs" />

    <section class="border-y border-white/10 bg-slate-900/45 px-4 py-16 sm:px-6 lg:px-8">
      <div class="mx-auto max-w-7xl">
        <SectionHeading eyebrow="Quy trình hoạt động" title="Quy trình minh bạch và dễ dàng cho từng đơn ứng tuyển" />
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
        <SectionHeading eyebrow="Đánh giá từ người dùng" title="Được tin dùng bởi hàng ngàn sinh viên & doanh nghiệp" />
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
            Nền tảng tuyển dụng hiện đại dành cho việc làm sinh viên, thực tập, trao đổi trực tiếp, chấm công ca làm và ví tiền tiện lợi.
          </p>
        </div>
        <FooterGroup title="Liên kết nhanh" :items="['Trang chủ', 'Khám phá việc làm', 'Tính năng', 'Liên hệ']" />
        <FooterGroup title="Dành cho người dùng" :items="['Sinh viên', 'Doanh nghiệp', 'Đăng ký', 'Đăng nhập']" />
        <div>
          <h3 class="font-extrabold text-white">Liên hệ</h3>
          <div class="mt-4 space-y-2 text-sm font-medium text-slate-400">
            <p>support@quickwork.vn</p>
            <p>0123 456 789</p>
            <p>TP. Hồ Chí Minh, Việt Nam</p>
          </div>
          <div class="mt-5 flex gap-3">
            <span v-for="social in ['Facebook', 'LinkedIn', 'GitHub']" :key="social" class="rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-xs font-bold text-slate-300">{{ social }}</span>
          </div>
        </div>
      </div>
      <p class="mx-auto mt-6 max-w-7xl text-center text-sm font-medium text-slate-500">© 2026 QuickWork. Bản quyền thuộc về QuickWork.</p>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { defineComponent, h } from 'vue'
import ExploreJobsSection from '~/components/home/ExploreJobsSection.vue'

const heroStats = [
  { value: '300+', label: 'Việc làm' },
  { value: '500+', label: 'Sinh viên' },
  { value: '150+', label: 'Doanh nghiệp' }
]

const introCards = [
  { icon: 'S', title: 'Sinh viên', text: 'Tìm kiếm công việc thực tập và bán thời gian linh hoạt phù hợp với lịch học.' },
  { icon: 'D', title: 'Doanh nghiệp', text: 'Tuyển dụng nhân sự sinh viên nhanh chóng với quy trình nộp đơn và offer tập trung.' },
  { icon: 'C', title: 'Nhắn tin Real-time', text: 'Trao đổi công việc trực tiếp giữa ứng viên và nhà tuyển dụng qua WebSocket.' },
  { icon: 'V', title: 'Ví & Thanh toán', text: 'Nhận lương và theo dõi lịch sử thu nhập minh bạch ngay sau khi hoàn thành công việc.' }
]

const features = [
  ['Tìm việc làm', 'Khám phá các vị trí tuyển dụng uy tín từ các doanh nghiệp.'],
  ['Bộ lọc thông minh', 'Lọc công việc theo từ khóa, địa điểm, hình thức và mức lương.'],
  ['Theo dõi đơn ứng tuyển', 'Theo dõi tiến trình xét duyệt đơn ứng tuyển minh bạch.'],
  ['Quản lý Offer', 'Nhận và phản hồi Offer nhận việc ngay trên màn hình cá nhân.'],
  ['Chấm công ca làm', 'Điểm danh Check-in và Check-out trong suốt thời gian làm việc.'],
  ['Ví tiền cá nhân', 'Theo dõi số dư thu nhập và lịch sử giao dịch thanh toán lương.'],
  ['Hệ thống Đánh giá', 'Tạo dựng uy tín cá nhân và thương hiệu sau mỗi công việc hoàn thành.'],
  ['Chat trực tiếp', 'Trò chuyện nhanh chóng với nhà tuyển dụng hoặc ứng viên.']
].map(([title, text], index) => ({ icon: String(index + 1), title, text }))

const timeline = [
  { title: 'Tạo tài khoản', text: 'Đăng ký tài khoản Sinh viên hoặc Doanh nghiệp.' },
  { title: 'Ứng tuyển', text: 'Sinh viên nộp đơn vào công việc phù hợp.' },
  { title: 'Duyệt hồ sơ', text: 'Doanh nghiệp xem xét hồ sơ và kỹ năng của ứng viên.' },
  { title: 'Nhận Offer', text: 'Sinh viên đủ điều kiện nhận được thông báo Offer công việc.' },
  { title: 'Bắt đầu làm việc', text: 'Sinh viên xác nhận chấp nhận và bắt đầu ca làm.' },
  { title: 'Chấm công', text: 'Điểm danh ca làm việc hàng ngày bằng Check-in/Check-out.' },
  { title: 'Thanh toán lương', text: 'Nhận lương trực tiếp vào ví sau khi công việc hoàn thành.' },
  { title: 'Viết đánh giá', text: 'Hai bên để lại đánh giá để nâng cao uy tín trên hệ thống.' }
]

const stats = [
  { value: '150+', label: 'Doanh nghiệp' },
  { value: '500+', label: 'Sinh viên' },
  { value: '1200+', label: 'Lượt nộp đơn' },
  { value: '300+', label: 'Công việc' }
]

const testimonials = [
  { name: 'Minh Anh', role: 'Sinh viên CNTT', quote: 'QuickWork giúp mình dễ dàng tìm được công việc thực tập linh hoạt và quản lý nhận lương rất rõ ràng.' },
  { name: 'Thanh Nguyễn', role: 'Trưởng phòng HR', quote: 'Quy trình tuyển dụng và trao đổi qua chat giúp công ty kết nối với sinh viên nhanh hơn bao giờ hết.' },
  { name: 'Quốc Bảo', role: 'Sinh viên Lập trình', quote: 'Mình rất thích việc chấm công ca làm và ví tiền tích hợp ngay trên cùng một trang điều hướng.' }
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

