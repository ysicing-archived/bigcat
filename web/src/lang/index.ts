/**
 * Copyright (c) 2022 ysicing All rights reserved.
 * Use of this source code is governed by AGPL-3.0-or-later
 * license that can be found in the LICENSE file.
 */


import enUS from "./en-us"
import zhCN from "./zh-cn"
import { createI18n } from "vue-i18n";

const messages = {
      'en-US': {
            ...enUS
      },
      'zh-CN': {
            ...zhCN
      }
}


export let defaultLang = 'zh-CN'


const i18n = createI18n({
      legacy: false,
      globalInjection: true,
      silentTranslationWarn: true,
      locale: defaultLang,
      fallbackLocale: defaultLang,
      messages
})

export default i18n