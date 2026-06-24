import { defineNuxtConfig } from 'nuxt/config'

export default defineNuxtConfig({
  compatibilityDate: '2026-06-24',
  
  // Enable Nuxt 4 compatibility version
  future: {
    compatibilityVersion: 4,
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
