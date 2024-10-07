import { defineConfig, ViteCustomizableConfig } from "@solidjs/start/config";
import type { HmrOptions } from "vite";

export default defineConfig({
  
  ssr: true,
  
  vite({ router }) {
    const viteConf: ViteCustomizableConfig = {
      
      resolve: {
        alias: {
          "~": "/src",
        },
      },
      
      server: {
        host: "0.0.0.0",
        port: 3000,
        strictPort: true,
      },
      
      plugins: [
      ],
      
    };
    if (router === 'client' && viteConf.server) {
      (viteConf.server as any).hmr = {
        protocol: "ws",
        port: 44447,
        clientPort: 3023,
        path: "/_hmr",
      } satisfies HmrOptions;
    }
    return viteConf;
  },
});
