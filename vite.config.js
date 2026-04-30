import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src')
    }
  },
  server: {
    proxy: {
      '/login': {
        target: 'http://localhost:3000',
        changeOrigin: true
      },
      '/refresh': {
        target: 'http://localhost:3000',
        changeOrigin: true
      },
      '/logout': {
        target: 'http://localhost:3000',
        changeOrigin: true
      },
      '/admin': {
        target: 'http://localhost:3000',
        changeOrigin: true
      },
      '/profile': {
        target: 'http://localhost:3000',
        changeOrigin: true
      }
    }
  }
})