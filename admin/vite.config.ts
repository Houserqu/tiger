import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { viteMockServe } from 'vite-plugin-mock'
import { ConfigEnv } from 'vite'
import { UserConfigExport } from 'vite'
import monacoEditorPlugin from "vite-plugin-monaco-editor"

// https://vitejs.dev/config/
export default ({ command }: ConfigEnv): UserConfigExport => {
  return ({
    base: 'admin',
    plugins: [
      react(),
      monacoEditorPlugin(),
      viteMockServe({
        mockPath: 'mock',
        localEnabled: false,
        ignore: /^mockUtil.ts$/g,
        injectCode: `
          import { setupProdMockServer } from './mockProdServer';
          setupProdMockServer();
        `,
      })
    ],
    build: {
      outDir: '../public',
    },
    server: {
      host: '0.0.0.0',
      port: 3001,
      proxy: {
        '/api': {
          target: 'http://localhost:8080',
          changeOrigin: true,
        },
        '/pub': {
          target: 'http://localhost:8080',
          changeOrigin: true,
        },
      },
    }
  })
}
