/**
 * Copyright (c) 2022 ysicing All rights reserved.
 * Use of this source code is governed by AGPL-3.0-or-later
 * license that can be found in the LICENSE file.
 */

import { Module } from "vuex";
import { RootStore } from "../types";

export interface menuStore {
      selectedKey: string[],
      activeKey: string
}

export const menu: Module<menuStore, RootStore> = {
      namespaced: true,
      state: {
            selectedKey: ["/home"],
            activeKey: "dml"
      },
      mutations: {
            CHANGE_SELECTED (state, vl: string[]) {
                  state.selectedKey = vl
            },
            CHANGE_ACTIVE_TABS (state, vl: string) {
                  state.activeKey = vl
            }
      }
}