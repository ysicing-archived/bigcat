/**
 * Copyright (c) 2022 ysicing All rights reserved.
 * Use of this source code is governed by AGPL-3.0-or-later
 * license that can be found in the LICENSE file.
 */

import { Module } from "vuex";
import { RootStore } from "../types";
import { LoginRespPayload } from "@/types"

export interface userStore {
      account: LoginRespPayload
}

export const user: Module<userStore, RootStore> = {
      namespaced: true,
      state: {
            account: {
                  token: "",
                  real_name: "",
                  user: "",
                  is_record: 2
            } as LoginRespPayload
      },
      mutations: {
            USER_STORE (state, vl: LoginRespPayload) {
                  state.account = vl
            }
      }
}