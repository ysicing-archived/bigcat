<!--
 Copyright (c) 2022 ysicing All rights reserved.
 Use of this source code is governed by AGPL-3.0-or-later
 license that can be found in the LICENSE file.
-->

<template>
      <a-config-provider :locale="locale">
            <router-view />
      </a-config-provider>
</template>

<script lang="ts" setup>
import zhCN from 'ant-design-vue/es/locale/zh_CN';
import enUS from 'ant-design-vue/es/locale/en_US'
import dayjs from 'dayjs';
import 'dayjs/locale';
import { defaultLang } from "@/lang"
import { onMounted } from 'vue';
import { useStore } from './store';
dayjs.locale(defaultLang);
const locale = defaultLang === "en-US" ? enUS : zhCN
</script>