import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue';


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8125',
        changeOrigin: true,
      },
      "/l/": {
        target: 'http://localhost:8125',
        changeOrigin: true,
      },
    }
  },
  resolve: {
    alias: {
      "@": __dirname + "/src",
    },
  },
})
