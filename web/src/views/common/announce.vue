<!--
 Copyright (c) 2022 ysicing All rights reserved.
 Use of this source code is governed by AGPL-3.0-or-later
 license that can be found in the LICENSE file.
-->

<template>
      <a-modal v-model:visible="is_open" title="欢迎使用大猫平台" @ok="close" :width="800">
            <p v-for="i in Announce">{{ i }}</p>
      </a-modal>
</template>

<script setup lang="ts">
import CommonMixin from "@/mixins/common";
import { Announce } from "@/config/vars"

const { is_open, close, open } = CommonMixin()

defineExpose({
      open
})

</script>