<!--
 Copyright (c) 2022 ysicing All rights reserved.
 Use of this source code is governed by AGPL-3.0-or-later
 license that can be found in the LICENSE file.
-->

<template>
      <a-modal v-model:visible="is_open" :title="$t('auto.title')" @ok="postOk">
            <a-form layout="vertical" :model="autotask" :rules="rules" ref="formRef">
                  <a-form-item :label="$t('auto.edit.name')" name="name">
                        <a-input v-model:value="autotask.name"></a-input>
                  </a-form-item>
                  <a-form-item :label="$t('common.table.type')">
                        <a-select v-model:value="autotask.tp">
                              <a-select-option v-for="i in taskTp" :value="i.v">{{ i.title }}</a-select-option>
                        </a-select>
                  </a-form-item>
                  <a-form-item :label="$t('common.table.env')" name="idc">
                        <a-select v-model:value="autotask.idc" @change="fetchSource">
                              <a-select-option
                                    v-for="i in fetchList.idc"
                                    :key="i"
                                    :value="i"
                              >{{ i }}</a-select-option>
                        </a-select>
                  </a-form-item>
                  <a-form-item :label="$t('common.table.source')" name="sourceLabel">
                        <a-select
                              v-model:value="autotask.sourceLabel"
                              @change="fetchSchema"
                              labelInValue
                              optionLabelProp="label"
                        >
                              <a-select-option
                                    v-for="i in fetchList.source"
                                    :value="i.source_id"
                                    :label="i.source"
                              >{{ i.source }}</a-select-option>
                        </a-select>
                  </a-form-item>
                  <a-form-item :label="$t('common.table.schema')" name="data_base">
                        <a-select v-model:value="autotask.data_base" @change="fetchTable">
                              <a-select-option
                                    v-for="i in fetchList.schema"
                                    :key="i"
                                    :value="i"
                              >{{ i }}</a-select-option>
                        </a-select>
                  </a-form-item>
                  <a-form-item :label="$t('common.table.table')" name="table">
                        <a-select v-model:value="autotask.table">
                              <a-select-option
                                    v-for="i in fetchList.tables"
                                    :key="i"
                                    :value="i"
                              >{{ i }}</a-select-option>
                        </a-select>
                  </a-form-item>
                  <a-form-item :label="$t('common.table.max')">
                        <a-input-number v-model:value="autotask.affect_rows" :min="0"></a-input-number>
                  </a-form-item>
                  <a-form-item :label="$t('auto.edit.enabled')">
                        <a-switch
                              v-model:checked="autotask.status"
                              :checkedValue="1"
                              :unCheckedValue="0"
                        ></a-switch>
                  </a-form-item>
            </a-form>
      </a-modal>
</template>

<script lang="ts" setup>

import { AutoTask, Request as Re } from "@/apis/autotask";
import { DBRelated, Request } from "@/apis/fetchSchema";
import { Res } from "@/config/request";
import CommonMixins from "@/mixins/common"
import { LabelInValue } from "@/types";
import { RuleObject } from 'ant-design-vue/es/form';
import { AxiosResponse } from "axios";
import { onMounted, reactive, unref, ref } from "vue";
import { useI18n } from 'vue-i18n';

const { t } = useI18n()

const formRef = ref()

const fetchList = reactive({
      source: [] as AutoTask[],
      idc: [] as string[],
      schema: [] as string[],
      tables: [] as string[],
})

const autotask = ref<AutoTask>({
      name: "",
      tp: 0,
      sourceLabel: {} as LabelInValue,
      data_base: "",
      table: "",
      affect_rows: 0,
      idc: "",
})

const checkSource = async (_rule: RuleObject, value: LabelInValue) => {
      if (value.value === undefined) {
            return Promise.reject(t('common.check.source'));
      }
      return Promise.resolve()
}

const rules = {
      name: [{ required: true, trigger: 'blur', message: t('common.check.name') }],
      idc: [{ required: true, trigger: 'change', message: t('common.check.env') }],
      sourceLabel: [{ required: true, trigger: 'change', validator: checkSource }],
      data_base: [{ required: true, trigger: 'change', message: t('common.check.schema') }],
      table: [{ required: true, trigger: 'change', message: t('common.check.table') }],

}

const initTask = Object.assign({}, autotask)

const request = new Request

const postTask = new Re

const emit = defineEmits(['success'])

const { is_open, turnState, taskTp } = CommonMixins()

const newTask = () => {
      Object.assign(autotask, unref(initTask))
      turnState()
}

const editTask = (task: AutoTask) => {
      autotask.value = Object.assign({}, task)
      turnState()
}

const postOk = () => {
      formRef.value.validateFields().then(() => {
            postTask.Post('curd', autotask.value).then(() => {
                  turnState()
                  emit('success')
            })
      })
}

const fetchSource = (vl: string) => {
      request.Source('idc', vl).then((res: AxiosResponse<Res<any[]>>) => {
            fetchList.source = res.data.data
      })
}

const fetchSchema = (vl: LabelInValue) => {
      request.Schema(vl.value).then((res: AxiosResponse<Res<DBRelated>>) => fetchList.schema = res.data.data.results)
}

const fetchTable = (vl: string) => {
      request.Table(autotask.value.sourceLabel.value, vl).then((res: AxiosResponse<Res<DBRelated>>) => fetchList.tables = res.data.data.results)
}

onMounted(() => {
      request.IDC().then((res: AxiosResponse<Res<string[]>>) => fetchList.idc = res.data.data)
})


defineExpose({
      newTask,
      editTask
})


</script>
