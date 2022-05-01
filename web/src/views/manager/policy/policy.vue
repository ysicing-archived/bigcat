<!--
 Copyright (c) 2022 ysicing All rights reserved.
 Use of this source code is governed by AGPL-3.0-or-later
 license that can be found in the LICENSE file.
-->

<template>
      <PageHeader :title="title.title" :subTitle="title.subTitle"></PageHeader>
      <RuleTable></RuleTable>
</template>
<script lang="ts" setup>
import PageHeader from "@/components/pageHeader/pageHeader.vue";
import RuleTable from "./policyTable.vue";
const title = {
      title: "权限组",
      subTitle: "权限组管理页面"
}
</script>