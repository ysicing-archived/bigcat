<!--
 Copyright (c) 2022 ysicing All rights reserved.
 Use of this source code is governed by AGPL-3.0-or-later
 license that can be found in the LICENSE file.
-->

<template>
      <a-row>
            <a-col :span="24">
                  <a-form layout="inline">
                        <a-form-item>
                              <a-button type="primary" @click="p.newTask()">{{ $t("auto.create") }}</a-button>
                        </a-form-item>
                        <a-form-item>
                              <a-input-search
                                    :placeholder="$t('auto.search.tips')"
                                    enter-button
                                    allowClear
                                    v-model:value="tblRef.expr.text"
                                    @search="tbl.manual()"
                              />
                        </a-form-item>
                  </a-form>
            </a-col>
      </a-row>

      <br />
      <c-table :tblRef="tblRef" ref="tbl">
            <template #bodyCell="{ column, text, record }">
                  <template v-if="column.dataIndex === 'action'">
                        <a-space>
                              <a-button
                                    ghost
                                    size="small"
                                    @click="p.editTask(record)"
                              >{{ $t('common.edit') }}</a-button>
                              <a-popconfirm
                                    :title="$t('auto.delete.tips')"
                                    @confirm="request.Delete(record.task_id).then(() => tbl.manual())"
                              >
                                    <a-button ghost size="small" danger>{{ $t('common.delete') }}</a-button>
                              </a-popconfirm>
                        </a-space>
                  </template>
                  <template
                        v-if="column.dataIndex === 'tp'"
                  >{{ taskTp.filter(item => item.v === text)[0].title }}</template>
                  <template v-if="column.dataIndex === 'status'">
                        <a-switch
                              v-model:checked="record.status"
                              :checkedValue="1"
                              :unCheckedValue="0"
                              @change="() => request.Post('active', record)"
                        ></a-switch>
                  </template>
            </template>
      </c-table>
      <AutotaskModal ref="p" @success="tbl.manual()"></AutotaskModal>
</template>

<script lang="ts" setup>

import { onMounted, reactive, ref } from "vue"
import CommonMixins from "@/mixins/common"
import { AutoTaskExpr, AutoTask, AutoTaskParams, AutoTaskResp, Request } from "@/apis/autotask"
import { AxiosResponse } from "axios"
import { Res } from "@/config/request"
import AutotaskModal from "./autotaskModal.vue"
import { useI18n } from 'vue-i18n';
import { tableRef } from "@/components/table"

const { t } = useI18n()

const tblRef = reactive<tableRef>({
      col: [
            {
                  title: t('common.table.name'),
                  dataIndex: 'name',
            },
            {
                  title: t('common.table.type'),
                  dataIndex: 'tp',
            },
            {
                  title: t('common.table.env'),
                  dataIndex: 'idc',
            },
            {
                  title: t('common.table.source'),
                  dataIndex: 'source',
            },
            {
                  title: t('common.table.schema'),
                  dataIndex: 'data_base',
            },
            {
                  title: t('common.table.table'),
                  dataIndex: 'table',
            },
            {
                  title: t('common.table.max'),
                  dataIndex: 'affect_rows',
            },
            {
                  title: t('common.table.state'),
                  dataIndex: 'status',
            },
            {
                  title: t('common.action'),
                  dataIndex: 'action',
            },
      ],
      data: [] as AutoTask[],
      pageCount: 0,
      expr: {} as AutoTaskExpr,
      fn: (expr: AutoTaskParams) => {
            request.List(expr).then((res: AxiosResponse<Res<AutoTaskResp>>) => {
                  tblRef.data = res.data.data.data
                  tblRef.pageCount = res.data.data.page
            })
      }
})

const p = ref()

const tbl = ref()

const { taskTp } = CommonMixins()

const request = new Request


</script>
