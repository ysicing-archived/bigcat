// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ergoapi/exgin"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/bigcat/internal/app/model"
)

type loginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserGeneralLogin(c *gin.Context) {
	var login loginForm
	exgin.Bind(c, &login)

	var account model.User
	err := model.DB().Where("username = ?", login.Username).First(&account).Error
	if err != nil {
		exgin.GinsData(c, nil, err)
		return
	}
	dataStore := map[string]interface{}{
		"token":     "",
		"real_name": account.RealName,
		"user":      account.Username,
	}
	exgin.GinsData(c, dataStore, nil)
}
