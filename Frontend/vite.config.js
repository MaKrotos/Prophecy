import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';

export default defineConfig({
  plugins: [vue()],
  server: {
    allowedHosts: ['rhwfn-85-198-105-19.a.free.pinggy.link']
  }
});