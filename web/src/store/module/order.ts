/**
 * Copyright (c) 2022 ysicing All rights reserved.
 * Use of this source code is governed by AGPL-3.0-or-later
 * license that can be found in the LICENSE file.
 */

import { Module } from "vuex";
import { RootStore } from "../types";
import { OrderTableData } from "@/types"

export interface orderStore {
      order: OrderTableData
      enabled: boolean,
      history: string[]
}

export const order: Module<orderStore, RootStore> = {
      namespaced: true,
      state: {
            order: {} as OrderTableData,
            enabled: true,
            history: [] as string[]
      },
      mutations: {
            ORDER_STORE (state, vl: OrderTableData) {
                  state.order = vl
            },
            SET_ORDER_STATUS (state, vl: number) {
                  state.order.status = vl
            },
            IS_ENABLED (state, vl: boolean) {
                  state.enabled = vl
            },
            SET_QUERY_HISTORY (state, vl: string) {
                  if (state.history.length < 100) {
                        state.history.unshift(vl)
                  } else {
                        state.history.pop()
                        state.history.unshift(vl)
                  }
            }
      }
}