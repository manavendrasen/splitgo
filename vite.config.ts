import path from "path";
import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";
import { VitePWA } from "vite-plugin-pwa";

export default defineConfig({
  plugins: [
    react(),
    VitePWA({
      registerType: "autoUpdate",
      devOptions: {
        enabled: true,
      },
      includeAssets: ["**/*.{png}"],
      manifest: {
        name: "SplitGo",
        start_url: "/",
        id: "/",
        short_name: "SplitGo",
        description: "Money Split",
        theme_color: "#031c36",
        icons: [
          {
            src: "/android/android-launchericon-512-512.png",
            sizes: "512x512",
            type: "image/png",
            purpose: "maskable",
          },
          {
            src: "/android/android-launchericon-192-192.png",
            sizes: "192x192",
            type: "image/png",
            purpose: "maskable",
          },
        ],
        shortcuts: [{ name: "New Split", url: "/add-expense" }],
      },
    }),
  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
});
