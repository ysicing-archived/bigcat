<!--
 Copyright (c) 2022 ysicing All rights reserved.
 Use of this source code is governed by AGPL-3.0-or-later
 license that can be found in the LICENSE file.
-->

<template>
      <a-card>
            <a-input-search v-model:value="searchValue" style="margin-bottom: 8px" placeholder="搜索" />
            <a-spin :spinning="spinning">
                  <a-tree v-model:expandedKey="expandedKeys" :auto-expand-parent="autoExpandParent" :tree-data="gData"
                        :load-data="onLoadData" :height="700" show-icon>
                        <template #switcherIcon="{ dataRef, defaultIcon }">
                              <template v-if="dataRef.meta === 'Schema'">
                                    <hdd-outlined />
                              </template>
                        </template>
                        <template #title="{ title, meta, key: treeKey, }">
                              <a-dropdown :trigger="['contextmenu']">
                                    <template v-if="title !== undefined">
                                          <span v-if="title.indexOf(searchValue) > -1">
                                                {{ title.substr(0, title.indexOf(searchValue)) }}
                                                <span style="color: #f50">{{ searchValue }}</span>
                                                {{ title.substr(title.indexOf(searchValue) + searchValue.length) }}
                                          </span>
                                          <span v-else>{{ title }}</span>
                                    </template>
                                    <template #overlay>
                                          <a-menu>
                                                <a-menu-item key="1" v-if="meta === 'Table'"
                                                      @click="showTableData(treeKey)">查看表数据</a-menu-item>
                                          </a-menu>
                                    </template>
                              </a-dropdown>
                        </template>
                  </a-tree>
            </a-spin>
      </a-card>
</template>

<script lang="ts" setup>

import { onMounted, ref, watch } from "vue"
import { Request } from "@/apis/query";
import { useRoute } from "vue-router";
import { AxiosResponse } from "axios";
import { Res } from "@/config/request";
import { HddOutlined, TableOutlined, CloudOutlined, ArrowLeftOutlined } from "@ant-design/icons-vue";
import { useStore } from "@/store";
import type { TreeProps } from 'ant-design-vue';

const emit = defineEmits(["showTableRef"])

const route = useRoute()

const searchValue = ref<string>("")

const store = useStore()

const expandedKeys = ref<any[]>([])

const autoExpandParent = ref<boolean>(false)

const spinning = ref(false)

let gData = ref<any>([])

let dataList = [] as any[]

const request = new Request

const schema = ref("")

watch(searchValue, value => {
      let expanded: string[] = []
      gData.value.forEach((item: any) => {
            if (item.children !== undefined) {
                  if (item.children[0].key !== undefined) {
                        item.children.forEach((el: any) => {
                              if (el.title.indexOf(value) > -1) {
                                    if (expanded.indexOf(item.title) == -1)
                                          expanded.push(item.title)
                              }
                        });

                  }
            }

      });
      expandedKeys.value = expanded
      searchValue.value = value;
      autoExpandParent.value = false;
});

const onLoadData: TreeProps['loadData'] = (treeNode: any) => {
      return new Promise(resolve => {
            if (treeNode.dataRef.meta === "Table") {
                  resolve();
                  return;
            }
            spin()
            request.QueryTable(route.query.source_id as string, treeNode.dataRef.title).then((res: AxiosResponse<Res<any>>) => {
                  treeNode.dataRef.children = res.data.data.table
                  gData.value = [...gData.value];
            }).finally(() => spin())
            expandedKeys.value = [treeNode.dataRef.title]
            resolve();
      });
}

const spin = () => {
      spinning.value = !spinning.value
}

const showTableData = (key: string) => {
      emit("showTableRef", { source_id: store.state.common.queryInfo.source_id, schema: schema.value, sql: `select * from ${key}` })
}

onMounted(() => {
      spin()
      request.QuerySchema(route.query.source_id as string).then((res: AxiosResponse<Res<any>>) => {
            gData.value = res.data.data.info
            dataList = res.data.data.info
            expandedKeys.value = [res.data.data.info.key]
            if (res.data.data.info.length > 0) {
                  store.commit("common/SET_SCHEMA_List",
                        {
                              schema: res.data.data.info.map((item: { key: string; }) => item.key),
                              source: route.query.source as string,
                              source_id: route.query.source_id as string
                        }
                  )
                  store.commit("common/SET_SCHEMA", "")
            }
      }
      ).finally(() => spin())

})

</script>
