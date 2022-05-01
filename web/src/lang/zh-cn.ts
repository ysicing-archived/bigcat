/**
 * Copyright (c) 2022 ysicing All rights reserved.
 * Use of this source code is governed by AGPL-3.0-or-later
 * license that can be found in the LICENSE file.
 */

import setting from "@/lang/zh-cn/setting"
import common from "@/lang/zh-cn/common"
import menu from "@/lang/zh-cn/menu"
import autoTask from "./zh-cn/autoTask";
import order from "./zh-cn/order";
import user from "./zh-cn/user";
import antdZhCN from 'ant-design-vue/es/locale/zh_CN';
import record from "./zh-cn/record"
import query from "./zh-cn/query";


const components = {
      antLocale: antdZhCN,
      DayJsName: 'zh-cn',
}

export default {
      ...components,
      ...setting,
      ...common,
      ...menu,
      ...autoTask,
      ...order,
      ...user,
      ...record,
      ...query
}