/**
 * Copyright (c) 2022 ysicing All rights reserved.
 * Use of this source code is governed by AGPL-3.0-or-later
 * license that can be found in the LICENSE file.
 */

import { request, COMMON_URI } from "@/config/request"
import { AxiosPromise } from "axios"

export interface Source {
      idc: string
      source: string
      ip: string
      port: number
      username: string
      password: string
      is_query: number
      flow_id: number
      source_id: string
      exclude_db_list: string
      insulate_word_list: string
}

export interface DBParams {
      page: number
      find: DBExpr
}

export interface DBExpr {
      idc: string
      source: string
      ip: string
      is_query: number
}

export interface DBResp {
      data: Source[]
      page: number
}

export interface RequestDB {
      tp: string
      db: Source[] | Source
      encrypt?: boolean
}

export class Request {
      List (args: DBParams): AxiosPromise {
            return request({
                  method: 'put',
                  url: `${COMMON_URI}/manage/db`,
                  data: args
            })
      }

      Delete (args: string): AxiosPromise {
            return request({
                  method: 'DELETE',
                  url: `${COMMON_URI}/manage/db`,
                  params: { source_id: args }
            })
      }

      Ops (args: RequestDB): AxiosPromise {
            return request({
                  method: 'POST',
                  url: `${COMMON_URI}/manage/db`,
                  data: args
            })
      }

}