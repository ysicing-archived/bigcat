// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package web

import (
	"github.com/gin-gonic/gin"
	"github.com/ysicing/bigcat/internal/dash/routes/handler"
)

type Handler struct{}

func NewHandler() (handler.RouteRegister, error) {
	return &Handler{}, nil
}

func (h *Handler) ApplyRoute(r *gin.Engine) {
	r.StaticFile("/", "web/dist/index.html")
	r.StaticFile("/icon.png", "web/dist/icon.png")
	r.StaticFile("/favicon.ico", "web/dist/favicon.ico")
	r.Static("/assets", "web/dist/assets")
}
