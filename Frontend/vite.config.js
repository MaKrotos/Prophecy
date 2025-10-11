import { fileURLToPath, URL } from "node:url";
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    }
  },
  server: {
    proxy: {
      '/api': {
        target: 'http://backend-prophecy:8080',
        changeOrigin: true,
        secure: false,
        rewrite: (path) => path.replace(/^\/api/, ''),
        configure: (proxy, options) => {
          proxy.on('proxyReq', (proxyReq, req, res) => {
            // Безопасная установка заголовков
            const remoteAddress = req.socket.remoteAddress || '';
            const forwardedFor = req.headers['x-forwarded-for'] || remoteAddress || '';
            const protocol = req.protocol || (req.connection?.encrypted ? 'https' : 'http') || 'http';
            
            proxyReq.setHeader('X-Real-IP', remoteAddress);
            proxyReq.setHeader('X-Forwarded-For', forwardedFor);
            proxyReq.setHeader('X-Forwarded-Proto', protocol);
          });
        },
        proxyTimeout: 30000,
        timeout: 30000,
      }
    }
  }
});