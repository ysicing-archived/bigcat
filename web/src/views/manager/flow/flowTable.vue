<!--
 Copyright (c) 2022 ysicing All rights reserved.
 Use of this source code is governed by AGPL-3.0-or-later
 license that can be found in the LICENSE file.
-->

<template>
      <a-button type="primary" @click="flow.newFlow()">新建流程</a-button>
      <br />
      <br />
      <a-table :columns="tblRef.col" :dataSource="tblRef.data" bordered :pagination="{
            total: tblRef.pageCount,
            showTotal: total => $t('common.count', { 'count': total }),
            position: ['bottomLeft'],
            showSizeChanger: true,
            defaultPageSize: 10
      }">
            <template #bodyCell="{ column, text, record }">
                  <template v-if="column.dataIndex === 'action'">
                        <a-space size="small">
                              <a-button type="primary" size="small" ghost @click="flow.editFlow(record)">详情</a-button>
                              <a-popconfirm title="确认要删除该流程吗?"
                                    @confirm="() => request.Delete(record.id).finally(() => currentPage())">
                                    <a-button type="primary" size="small" danger ghost>删除</a-button>
                              </a-popconfirm>
                        </a-space>
                  </template>
            </template>
            <template #customFilterDropdown="{ setSelectedKeys, selectedKeys, confirm, clearFilters, column }">
                  <div style="padding: 8px">
                        <a-input ref="searchInput" :placeholder="`Search ${column.dataIndex}`" :value="selectedKeys[0]"
                              style="width: 188px; margin-bottom: 8px; display: block"
                              @change="e => setSelectedKeys(e.target.value ? [e.target.value] : [])"
                              @pressEnter="handleSearch(selectedKeys, confirm, column.dataIndex)" />
                        <a-button type="primary" size="small" style="width: 90px; margin-right: 8px"
                              @click="handleSearch(selectedKeys, confirm, column.dataIndex)">
                              <template #icon>
                                    <SearchOutlined />
                              </template>
                              Search
                        </a-button>
                        <a-button size="small" style="width: 90px" @click="handleReset(clearFilters)">Reset</a-button>
                  </div>
            </template>
            <template #customFilterIcon="{ filtered }">
                  <search-outlined :style="{ color: filtered ? '#108ee9' : undefined }" />
            </template>
      </a-table>
      <FlowModel ref="flow" :title="title" @success="currentPage"></FlowModel>
</template>

<script lang="ts" setup>
import { Res } from "@/config/request";
import { AxiosResponse } from "axios";
import { ref, reactive, onMounted } from "vue";
import { SearchOutlined } from '@ant-design/icons-vue';
import { Request, RespTPLs } from "@/apis/flow";
import FlowModel from "./flowModal.vue"
import { tableRef } from "@/components/table";


const tblRef = reactive<tableRef>({
      col: [
            {
                  title: "流程名称",
                  dataIndex: "source",
                  customFilterDropdown: true,
                  onFilter: (value: string, record: { source: { toString: () => string; }; }) =>
                        record.source.toString().toLowerCase().includes(value.toLowerCase()),
                  onFilterDropdownVisibleChange: (visible: any) => {
                        if (visible) {
                              setTimeout(() => {
                                    searchInput.value.focus();
                              }, 100);
                        }
                  },
            },
            { title: "操作", dataIndex: "action" },
      ] as any,
      data: [] as RespTPLs[],
      pageCount: 0
})

const searchInput = ref();

const title = ref("新建流程")

const flow = ref()

const request = new Request


const state = reactive({
      searchText: '',
      searchedColumn: '',
});

const handleSearch = (selectedKeys: string[], confirm: () => void, dataIndex: string) => {
      confirm();
      state.searchText = selectedKeys[0];
      state.searchedColumn = dataIndex;
};

const handleReset = (clearFilters: () => void) => {
      clearFilters();
      state.searchText = '';
};

const currentPage = () => {
      request.List().then((res: AxiosResponse<Res<RespTPLs[]>>) => {
            tblRef.data = res.data.data
            tblRef.pageCount = res.data.data.length
      })
}

onMounted(() => {
      currentPage()
})

</script>
