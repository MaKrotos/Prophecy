import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import vueDevTools from "vite-plugin-vue-devtools";
import fs from "fs";

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    watch: {
      usePolling: true,
      interval: 100,
    },
    port: 5173,
    host: "0.0.0.0",
    
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
});
