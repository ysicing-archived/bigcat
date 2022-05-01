<!--
 Copyright (c) 2022 ysicing All rights reserved.
 Use of this source code is governed by AGPL-3.0-or-later
 license that can be found in the LICENSE file.
-->

<template>
      <PageHeader :title="title.title" :subTitle="title.subTitle"></PageHeader>
      <a-row>
            <a-col :xs="24" :sm="6">
                  <a-divider orientation="left">{{ $t('user.form.title') }}</a-divider>
                  <RegisterForm isManager></RegisterForm>
            </a-col>
            <a-col :xs="24" :sm="17" :offset="1">
                  <UserTable></UserTable>
            </a-col>
      </a-row>
</template>

<script lang="ts" setup>
import PageHeader from "@/components/pageHeader/pageHeader.vue";
import RegisterForm from "@/components/user/registerForm.vue";
import UserTable from "@/views/manager/user/userTable.vue"
import { useI18n } from 'vue-i18n';

const { t } = useI18n()

const title = {
      title: t('user.title'),
      subTitle: t('user.desc')
}

</script>

<style scoped>
</style>