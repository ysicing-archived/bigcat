/**
 * Copyright (c) 2022 ysicing All rights reserved.
 * Use of this source code is governed by AGPL-3.0-or-later
 * license that can be found in the LICENSE file.
 */

import { COMMON_URI, request } from "@/config/request"
import { AxiosPromise } from "axios";

export class Request {
      Post (vl: string): AxiosPromise {
            return request({
                  method: 'post',
                  url: `${COMMON_URI}/manage/board/post`,
                  data: {
                        board: vl
                  }
            })
      }

      Get (): AxiosPromise {
            return request({
                  method: 'get',
                  url: `${COMMON_URI}/board/get`,
            })
      }
}