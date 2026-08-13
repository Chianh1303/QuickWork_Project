import { defineNuxtConfig } from 'nuxt/config'

export default defineNuxtConfig({
  compatibilityDate: '2026-06-24',

  app: {
    head: {
      link: [
        { rel: 'preconnect', href: 'https://fonts.googleapis.com' },
        { rel: 'preconnect', href: 'https://fonts.gstatic.com', crossorigin: '' },
        { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700;800&family=Outfit:wght@400;500;600;700;800&display=swap' }
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

  // Proxy /api requests to the Go Fiber backend running at http://localhost:3000
  routeRules: {
    '/api/**': { proxy: 'http://localhost:3000/api/**' }
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
