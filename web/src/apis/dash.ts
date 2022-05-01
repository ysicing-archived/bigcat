/**
 * Copyright (c) 2022 ysicing All rights reserved.
 * Use of this source code is governed by AGPL-3.0-or-later
 * license that can be found in the LICENSE file.
 */

import { request, COMMON_URI } from "@/config/request"
import { AxiosPromise } from "axios"

export class Request {
      Banner (): AxiosPromise {
            return request({
                  method: 'get',
                  url: `${COMMON_URI}/dash/banner`,
            })
      }

      Top (): AxiosPromise {
            return request({
                  method: 'get',
                  url: `${COMMON_URI}/dash/top`,
            })
      }
}

