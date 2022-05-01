<!--
 Copyright (c) 2022 ysicing All rights reserved.
 Use of this source code is governed by AGPL-3.0-or-later
 license that can be found in the LICENSE file.
-->

<template>
      <a-modal v-model:visible="is_open" title="让大猫(BigCat)持续提供更好的功能与服务" @ok="close">
            感谢您支持大猫(BigCat)。目前已添加以下功能:
            <p v-for="i in FeatureList">{{ i }}</p>
            <p>联系方式</p>
            <p>i@ysicing.me</p>
      </a-modal>
</template>

<script setup lang="ts">
import CommonMixin from "@/mixins/common";
import { FeatureList } from "@/config/vars"

const { is_open, close, open } = CommonMixin()

defineExpose({
      open
})

</script>