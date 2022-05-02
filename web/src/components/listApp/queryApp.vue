<!--
 Copyright (c) 2022 ysicing All rights reserved.
 Use of this source code is governed by AGPL-3.0-or-later
 license that can be found in the LICENSE file.
-->

<template>
      <component :is="com" type="query" :id="110" @enter="clickPage" :isExport="isExport" />
</template>


<script lang="ts" setup>
import { Res } from "@/config/request"
import { AxiosResponse } from "axios"
import { onMounted, ref, shallowRef } from "vue"
import { QueryPost, Request as Query } from "@/apis/query";
import ListApp from "./listApp.vue";
import QueryBanner from "./queryBanner.vue";
import QueryOrder from "./queryOrder.vue";

const com = shallowRef<any>(QueryBanner)

const isQuery = ref(false)

const isExport = ref<boolean>(false)

const query = new Query

const clickPage = () => {
      if (isQuery.value) {
            com.value = QueryOrder
      } else {
            com.value = ListApp
            query.Post({
                  export: isExport ? 1 : 0
            } as QueryPost)
      }
}

const fetchState = () => {
      query.IsQuery().then((res: AxiosResponse<Res<any>>) => {
            isQuery.value = res.data.data.status
            isExport.value = res.data.data.export
      })
}

onMounted(() => {
      query.QueryStatus().then((res: AxiosResponse<Res<boolean>>) => {
            !res.data.data ? com.value = ListApp : fetchState()
      })
})

</script>

<style scoped>
.desc p {
      margin-bottom: 1em;
}
</style>
