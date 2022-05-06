// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package user

import (
	"github.com/ergoapi/exgin"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/bigcat/internal/dash/model"
	"github.com/ysicing/bigcat/internal/dash/routes/handler"
	"github.com/ysicing/bigcat/internal/util"
)

type Handler struct{}

func NewHandler() (handler.RouteRegister, error) {
	return &Handler{}, nil
}

func (h *Handler) ApplyRoute(r *gin.Engine) {
	group := r.Group("/api/v1/user")
	group.GET("/list", h.UserList)
}

type page struct {
	Current  int ` json:"current"`
	PageSize int `json:"pageSize"`
}

func (Handler) UserList(c *gin.Context) {
	start, end := util.GinPage(c)
	var users []model.User
	var count int64
	model.DB().Model(&model.User{}).Count(&count).Offset(start).Limit(end).Find(&users)
	exgin.GinsData(c, map[string]interface{}{
		"page": count,
		"data": users,
	}, nil)
}
