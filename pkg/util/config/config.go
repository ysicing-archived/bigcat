// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package config

import "github.com/spf13/viper"

func GetString(s string, callback ...string) string {
	if v := viper.GetString(s); v != "" {
		return v
	}
	if len(callback) > 0 {
		return callback[0]
	}
	return ""
}

func GetBool(s string, callback ...bool) bool {
	k := GetString(s)
	if k == "" {
		if len(callback) > 0 {
			return callback[0]
		}
		return false
	}
	return viper.GetBool(s)
}
