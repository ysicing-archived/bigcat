<!--
 Copyright (c) 2022 ysicing All rights reserved.
 Use of this source code is governed by AGPL-3.0-or-later
 license that can be found in the LICENSE file.
-->

<template>
      <PageHeader :title="title.title" :subTitle="title.subTitle"></PageHeader>
      <OrderTable></OrderTable>
</template>
<script lang="ts"  setup>
import { ref } from "@vue/runtime-core";
import { useRoute, onBeforeRouteUpdate } from "vue-router";
import PageHeader from "@/components/pageHeader/pageHeader.vue";
import OrderTable from "@/components/table/orderTable.vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n()

const route = useRoute()

const checkTitle = (audit: string): { [key: string]: string } => {
      return audit === "audit" ? { title: t('menu.order.order') + t('menu.order.title'), subTitle: t('order.desc.audit') } : { title: t('menu.order.self'), subTitle: t('order.desc.self') }
}


let title = ref(checkTitle(route.params.tp as string))

onBeforeRouteUpdate((to) => {
      title.value = checkTitle(to.params.tp as string)
})

</script>