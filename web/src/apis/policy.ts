/**
 * Copyright (c) 2022 ysicing All rights reserved.
 * Use of this source code is governed by AGPL-3.0-or-later
 * license that can be found in the LICENSE file.
 */

import { request, COMMON_URI, Res } from "@/config/request"
import { AxiosPromise, AxiosResponse } from "axios"

export interface PolicyExpr {
      text: string
}

export interface PolicyParams {
      page: number
      find: PolicyExpr
}

interface line {
      id: number
      name: string
}

export interface Policy extends line {
      permissions: PolicyPost
}

export interface checkList {
      source?: string
      source_id?: string
}

export interface PolicyPost extends line {
      dml_source: string[]
      ddl_source: string[]
      query_source: string[]
}

export interface checkProps {
      source: checkList[]
      query: checkList[]
}

export interface PolicyResp {
      page: number
      data: Policy[]
}

export interface PolicyRuse {
      source: checkList[]
      query: checkList[]
}

export class Request {

      List (args: PolicyParams): AxiosPromise {
            return request({
                  method: 'put',
                  url: `${COMMON_URI}/manage/policy`,
                  data: args
            })
      }

      Get (): AxiosPromise {
            return request({
                  method: 'get',
                  url: `${COMMON_URI}/manage/policy/source`,
            })
      }

      Post (args: PolicyPost): AxiosPromise {
            return request({
                  method: 'post',
                  url: `${COMMON_URI}/manage/policy`,
                  data: args
            })
      }

      Drop (args: string): AxiosPromise {
            return request({
                  method: 'delete',
                  url: `${COMMON_URI}/manage/policy`,
                  params: {
                        group_id: args
                  }
            })
      }

}