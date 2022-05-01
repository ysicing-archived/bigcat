<!--
 Copyright (c) 2022 ysicing All rights reserved.
 Use of this source code is governed by AGPL-3.0-or-later
 license that can be found in the LICENSE file.
-->

<template>
      <PageHeader :title="title.title" :subTitle="title.subTitle"></PageHeader>
      <AutotaskTable></AutotaskTable>
</template>
<script lang="ts" setup>
import PageHeader from '@/components/pageHeader/pageHeader.vue';
import AutotaskTable from './autotaskTable.vue';
import { useI18n } from 'vue-i18n';

const { t } = useI18n()

const title = {
      title: t('auto.title'),
      subTitle: t('auto.desc')
}
</script>