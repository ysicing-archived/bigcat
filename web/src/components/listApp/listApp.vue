<!--
 Copyright (c) 2022 ysicing All rights reserved.
 Use of this source code is governed by AGPL-3.0-or-later
 license that can be found in the LICENSE file.
-->

<template>
      <a-row :gutter="24">
            <a-col :xs="24" :md="8">
                  <a-select
                        v-model:value="selected"
                        show-search
                        allowClear
                        placeholder="数据源搜索"
                        style="width: 100%;"
                        :filterOption="filterOption"
                        @change="handleChange"
                  >
                        <a-select-option value="all">{{ $t('order.state.all') }}</a-select-option>
                        <a-select-option v-for="i in options" :value="i.source">{{ i.source }}</a-select-option>
                  </a-select>
            </a-col>
      </a-row>
      <br />
      <a-list
            :loading="loading"
            :data-source="source"
            :grid="{ gutter: 16, xs: 1, sm: 1, md: 2, lg: 2, xl: 4, xxl: 4, xxxl: 4 }"
            :pagination="pagination"
      >
            <template #renderItem="{ item }">
                  <a-list-item>
                        <div
                              @click="() => router.push({ path: props.type !== 'query' ? '/apply/order' : '/apply/query', query: { type: props.id, idc: item.idc, source: item.source, source_id: item.source_id } })"
                        >
                              <a-card :body-style="{ paddingBottom: 20 }" hoverable>
                                    <a-card-meta :title="item.source">
                                          <template
                                                v-slot:description
                                          >{{ $t('order.apply.card.env', { 'env': item.idc }) }}</template>
                                          <template v-slot:avatar>
                                                <a-avatar :style="{ backgroundColor: '#Ff9900' }">
                                                      <template v-slot:icon>
                                                            <CodepenCircleOutlined />
                                                      </template>
                                                </a-avatar>
                                          </template>
                                    </a-card-meta>
                                    <template v-slot:actions>
                                          <a-tooltip
                                                :title="$t('order.apply.tab.source_id', { 'env': item.source_id })"
                                          >
                                                <SubnodeOutlined />
                                          </a-tooltip>
                                          <a-tooltip
                                                :title="$t('order.apply.card.env', { 'env': item.idc })"
                                          >
                                                <ShareAltOutlined />
                                          </a-tooltip>
                                          <a-dropdown>
                                                <a-tooltip :title="$t('order.apply.card.enter')">
                                                      <a
                                                            class="ant-dropdown-link"
                                                            @click="() => router.push({ path: props.type !== 'query' ? '/apply/order' : '/apply/query', query: { type: props.id, idc: item.idc, source: item.source, source_id: item.source_id } })"
                                                      >
                                                            <EnterOutlined />
                                                      </a>
                                                </a-tooltip>
                                          </a-dropdown>
                                    </template>
                              </a-card>
                        </div>
                  </a-list-item>
            </template>
      </a-list>
</template>

<script lang="ts" setup>
import { SubnodeOutlined, EnterOutlined, ShareAltOutlined, CodepenCircleOutlined } from "@ant-design/icons-vue"
import { AxiosResponse } from "axios"
import { Res } from "@/config/request"
import { onMounted, ref } from "@vue/runtime-core";
import { RespFetchSource } from "@/apis/listAppApis"
import { useRouter } from "vue-router"
import { Request } from "@/apis/fetchSchema";

const props = defineProps<{
      type: string,
      id: number,
      isExport?: boolean
}>()

const emit = defineEmits(['enter'])

const router = useRouter()

const pagination = {
      pageSize: 20,
}

const filterOption = (input: string, option: any) => {
      return option.value.toLowerCase().indexOf(input.toLowerCase()) >= 0;
};

const handleChange = (value: string) => {
      value === "" || value === undefined ? source.value = tmpSource : source.value = tmpSource.filter((item: RespFetchSource) => item.source === value)
};

const selected = ref("all")

const request = new Request

let tmpSource = [] as RespFetchSource[]

let source = ref([] as RespFetchSource[])

let options = ref([] as RespFetchSource[])

let loading = ref(true)

onMounted(() => {
      request.Source(props.type).then((res: AxiosResponse<Res<RespFetchSource[]>>) => {
            tmpSource = res.data.data
            source.value = res.data.data
            options.value = res.data.data
      }).finally(() => {
            loading.value = false
      })

})

</script>
