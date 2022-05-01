/**
 * Copyright (c) 2022 ysicing All rights reserved.
 * Use of this source code is governed by AGPL-3.0-or-later
 * license that can be found in the LICENSE file.
 */

declare module '*.vue' {
      import { DefineComponent } from 'vue'
      // eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/ban-types
      const component: DefineComponent<{}, {}, any>
      export default component
}

declare module '*.json' {
      const value: any;
      export default value;
}

declare module "nprogress"

declare module "vue-super-flow"
