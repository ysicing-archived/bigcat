/**
 * Copyright (c) 2022 ysicing All rights reserved.
 * Use of this source code is governed by AGPL-3.0-or-later
 * license that can be found in the LICENSE file.
 */

import { COMMON_URI, request } from "@/config/request"
import { AxiosPromise } from "axios"


export interface Rules {
      [key: string]: string | number | boolean
}


export class Request {
      List (): AxiosPromise {
            return request({
                  method: 'put',
                  url: `${COMMON_URI}/manage/roles`
            })
      }

      Post (args: any): AxiosPromise {
            return request({
                  method: 'post',
                  url: `${COMMON_URI}/manage/roles`,
                  data: args
            })
      }
}