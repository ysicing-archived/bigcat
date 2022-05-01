/**
 * Copyright (c) 2022 ysicing All rights reserved.
 * Use of this source code is governed by AGPL-3.0-or-later
 * license that can be found in the LICENSE file.
 */

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import * as path from "path"

// https://vitejs.dev/config/
export default defineConfig({
      base: "/",
      server: {
            proxy: {
                  '/fetch': {
                        target: 'http://127.0.0.1:65001'
                  },
                  '/login': {
                        target: 'http://127.0.0.1:65001'
                  },
                  '/api': {
                        target: 'http://127.0.0.1:65001'
                  },
                  '/downlaod/*': {
                        target: 'http://127.0.0.1:65001'
                  },
            },
      },
      plugins: [
            vue(),
      ],
      resolve: {
            alias: {
                  '@': path.resolve(__dirname, './src/'),
            }
      },
      css: {
            preprocessorOptions: {
                  less: {
                        javascriptEnabled: true
                  },
            }
      }
})
