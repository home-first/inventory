import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  server: {
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8090',
        changeOrigin: true,
      },
      "/_/": {
        target: 'http://127.0.0.1:8090',
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
