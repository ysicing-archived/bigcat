<!--
 Copyright (c) 2022 ysicing All rights reserved.
 Use of this source code is governed by AGPL-3.0-or-later
 license that can be found in the LICENSE file.
-->

<template>
      <a-page-header :title="props.title" @back="() => $router.go(-1)">
            <span style="opacity: 0.8;">{{ props.subTitle }}</span>
      </a-page-header>
</template>
<script lang="ts" setup>


const props = defineProps<{
      title: string,
      subTitle: string
}>()

</script>