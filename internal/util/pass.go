// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package util

import (
	"strconv"
	"strings"

	"github.com/ergoapi/exgin"
	"github.com/ergoapi/util/expass"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/bigcat/internal/dash/model"
)

func CheckPassword(user *model.User, password string) bool {
	sl := strings.Split(user.Password, "$")[2]
	checkPasswordToken := expass.SaltPbkdf2Pass(sl, password)
	return user.Password == checkPasswordToken
}

func Paging(page interface{}, total int) (start int, end int) {
	var i int
	switch v := page.(type) {
	case string:
		i, _ = strconv.Atoi(v)
	case int:
		i = v
	}
	start = i*total - total
	end = total
	return
}

func GinPage(c *gin.Context) (page int, pageSize int) {
	page = exgin.GinsQueryInt(c, "page", 1)
	pageSize = exgin.GinsQueryInt(c, "pageSize", 10)
	return Paging(page, pageSize)
}
