<!--
 Copyright (c) 2022 ysicing All rights reserved.
 Use of this source code is governed by AGPL-3.0-or-later
 license that can be found in the LICENSE file.
-->

<template>
      <PageHeader :title="title.title" :subTitle="title.subTitle"></PageHeader>
      <FlowTable></FlowTable>
</template>
<script lang="ts" setup>
import PageHeader from "@/components/pageHeader/pageHeader.vue";
import FlowTable from "./flowTable.vue";
const title = {
      title: "流程",
      subTitle: "流程管理页面"
}
</script>