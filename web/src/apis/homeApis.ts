/**
 * Copyright (c) 2022 ysicing All rights reserved.
 * Use of this source code is governed by AGPL-3.0-or-later
 * license that can be found in the LICENSE file.
 */

import { request } from "@/config/request"

export interface SelfCount {
      order: number;
      query: number;
      order_audit: number;
      query_audit: number;
}