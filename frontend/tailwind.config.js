/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./app.vue",
    "./error.vue",
    "./composables/**/*.{js,ts}",
    "./stores/**/*.{js,ts}"
  ],
  theme: {
    extend: {
      colors: {
        brand: {
          50: '#f4f6fc',
          100: '#e8edf7',
          200: '#cbd7ee',
          300: '#9cb5e0',
          400: '#688dcf',
          500: '#466ebe',
          600: '#35559d', // Deep blue
          700: '#2b447e',
          800: '#273b6a',
          900: '#233359',
          950: '#151e36',
        },
        accent: {
          50: '#fdf4f5',
          100: '#fbe8eb',
          200: '#f7d5d9',
          300: '#f0b4bc',
          400: '#e58897',
          500: '#d55f75', // Crimson Coral
          600: '#c0445d',
          700: '#a13149',
          800: '#862c3f',
          900: '#732939',
          950: '#40121c',
        }
      },
      fontFamily: {
        sans: ['Inter', 'ui-sans-serif', 'system-ui', '-apple-system', 'BlinkMacSystemFont', 'Segoe UI', 'Roboto', 'Helvetica Neue', 'Arial', 'sans-serif'],
      }
    },
  },
  plugins: [],
}
