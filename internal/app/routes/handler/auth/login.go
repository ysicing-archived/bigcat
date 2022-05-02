// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package auth

import (
	"fmt"

	"github.com/ergoapi/exgin"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/bigcat/internal/app/model"
	"github.com/ysicing/bigcat/internal/util"
	libutil "github.com/ysicing/bigcat/pkg/util"
)

type loginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h Handler) UserGeneralLogin(c *gin.Context) {
	var login loginForm
	exgin.Bind(c, &login)

	var user model.User
	err := model.DB().Where("username = ?", login.Username).First(&user).Error
	if err != nil {
		exgin.GinsData(c, nil, err)
		return
	}
	if util.CheckPassword(&user, login.Password) {
		token, tokenerr := libutil.JwtAuth(libutil.Token{
			Username: login.Username,
		})
		if tokenerr != nil {
			exgin.GinsData(c, nil, tokenerr)
			return
		}
		dataStore := map[string]interface{}{
			"token":     token,
			"real_name": user.RealName,
			"user":      user.Username,
		}
		exgin.GinsData(c, dataStore, nil)
		return
	}
	exgin.GinsData(c, nil, fmt.Errorf("username or password error"))
}
