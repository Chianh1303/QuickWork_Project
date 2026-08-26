import { defineNuxtConfig } from 'nuxt/config'

export default defineNuxtConfig({
  compatibilityDate: '2026-06-24',
  app: {
    head: {
      link: [
        { rel: 'preconnect', href: 'https://fonts.googleapis.com' },
        { rel: 'preconnect', href: 'https://fonts.gstatic.com', crossorigin: '' },
        { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:ital,wght@0,300;0,400;0,500;0,600;0,700;0,800;1,400&family=Inter:wght@300;400;500;600;700;800&display=swap' }
      ],
      script: [
        { src: 'https://accounts.google.com/gsi/client', async: true, defer: true }
      ]
    }
  },

  // Since pages, composables, stores, middleware, and assets are at the root level of the frontend/ folder,
  // we point the app directory to the root directory '.'.
  dir: {
    app: '.'
  },

  // Modules
  modules: [
    '@nuxtjs/tailwindcss',
    '@pinia/nuxt'
  ],

  // Proxy /api and /uploads requests to the Go Fiber backend running at http://localhost:3000
  routeRules: {
    '/api/**': { proxy: 'http://localhost:3000/api/**' },
    '/uploads/**': { proxy: 'http://localhost:3000/uploads/**' }
  },

  // Tailwind configuration options
  tailwindcss: {
    cssPath: '~/assets/css/main.css',
    configPath: 'tailwind.config.js',
    exposeConfig: false,
    viewer: false
  },

  // Nuxt Dev Server Port
  devServer: {
    port: 3001
  }
})
