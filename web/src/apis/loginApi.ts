/**
 * Copyright (c) 2022 ysicing All rights reserved.
 * Use of this source code is governed by AGPL-3.0-or-later
 * license that can be found in the LICENSE file.
 */

import { request } from "@/config/request"

export interface LoginFrom {
      username: string;
      password: string;
      is_ldap: boolean;
}

export const LoginApi = (login: LoginFrom) => {
      return request({
            url: login.is_ldap ? "/ldap" : "/login",
            method: "POST",
            data: login
      })
}

export const IsRegister = () => {
      return request({
            url: "/fetch",
            method: "GET",
      })
}