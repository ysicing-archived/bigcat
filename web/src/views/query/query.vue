<!--
 Copyright (c) 2022 ysicing All rights reserved.
 Use of this source code is governed by AGPL-3.0-or-later
 license that can be found in the LICENSE file.
-->

<template>
      <a-spin :spinning="spinning">
            <a-row>
                  <a-col :span="5">
                        <a-space>
                              <a-button size="small" @click="() => hide = !hide" type="primary">{{ $t('common.hide')
                              }}/{{ $t('common.visible') }}</a-button>
                              <a-button size="small" @click="m.turnState()">{{ $t('common.new') }}{{ $t('common.clip')
                              }}</a-button>
                        </a-space>
                  </a-col>
            </a-row>
            <br />
            <a-row>
                  <a-col :span="hide ? 0 : 5">
                        <a-tabs v-model:activeKey="tool">
                              <a-tab-pane key="tree" tab="数据库">
                                    <Tree @showTableRef="showTableRef"></Tree>
                              </a-tab-pane>
                              <a-tab-pane key="history" tab="历史记录">
                                    <History></History>
                              </a-tab-pane>
                        </a-tabs>
                  </a-col>
                  <a-col :span="hide ? 24 : 18" :offset="hide ? 0 : 1">
                        <a-tabs v-model:activeKey="feat">
                              <a-tab-pane key="edit" :tab="$t('query.query')">
                                    <a-tabs v-model:activeKey="activeKey" type="editable-card" @edit="onEdit">
                                          <a-tab-pane v-for="pane in panes" :key="pane.key" :tab="pane.title"
                                                :closable="pane.closable">
                                                <Input :id="pane.title" />
                                          </a-tab-pane>
                                    </a-tabs>
                              </a-tab-pane>
                              <a-tab-pane key="table" :tab="$t('query.table')" forceRender>
                                    <Table ref="tbl" :height="800" id="tblInfo"></Table>
                              </a-tab-pane>
                        </a-tabs>
                  </a-col>
            </a-row>
            <Clip></Clip>
            <Modal ref="m"></Modal>
      </a-spin>
</template>

<script lang="ts" setup>
import History from "./history.vue"
import Tree from "./tree.vue"
import Input from "./input.vue"
import Table from "./table.vue"
import Clip from "./clip.vue"
import Modal from "./modal.vue"
import Socket from "@/socket"
import { computed, onMounted, onUnmounted, ref } from "vue"
import { useStore } from "@/store"
import { encode } from "@msgpack/msgpack";

const panes = ref([{ title: 'Untitled 1', key: '1', closable: false }])

const activeKey = ref()

const hide = ref(false)

const newTabIndex = ref(1);

const feat = ref("edit")

const tool = ref("tree")

const tbl = ref()

const m = ref()

const store = useStore()

const spinning = computed(() => store.state.common.spinning)

const onEdit = (targetKey: string, action: string) => {
      if (action === 'add') {
            add();
      } else {
            remove(targetKey);
      }
}

const showTableRef = (vl: any) => {
      tbl.value.runResults(vl.source_id, vl.schema, vl.sql)
      feat.value = "table"
}

const add = () => {
      activeKey.value = `Untitled${++newTabIndex.value}`;
      panes.value.push({
            title: activeKey.value, key: activeKey.value, closable: true
      });
};

const remove = (targetKey: string) => {
      let lastIndex = 0;
      panes.value.forEach((pane, i) => {
            if (pane.key === targetKey) {
                  lastIndex = i - 1;
            }
      });
      panes.value = panes.value.filter(pane => pane.key !== targetKey);
      if (panes.value.length && activeKey.value === targetKey) {
            if (lastIndex >= 0) {
                  activeKey.value = panes.value[lastIndex].key;
            } else {
                  activeKey.value = panes.value[0].key;
            }
      }
}

onMounted(() => {
      const sock = new Socket(`/query/results?user=${store.state.user.account.user}`)
      sock.create()
      store.commit("common/QUERY_CONN", sock)
})

onUnmounted(() => {
      const encoded: Uint8Array = encode({ "type": "1", "sql": "", "schema": "", "source_id": "" });
      store.state.common.sock?.send(encoded)
      store.state.common.sock.close()
})


</script>