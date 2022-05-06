// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package dash

import (
	"github.com/ergoapi/exgin"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/bigcat/internal/dash/routes/handler"
)

type Handler struct{}

func NewHandler() (handler.RouteRegister, error) {
	return &Handler{}, nil
}

func (h *Handler) ApplyRoute(r *gin.Engine) {
	group := r.Group("/api/v1/dash")
	group.GET("/top", h.DashTop)
	group.GET("/banner", h.DashBanner)
}

type groupBy struct {
	Source string `json:"source"`
	C      int    `json:"count"`
	Time   string `json:"time"`
	Type   string `json:"type"`
}

type bannerCount struct {
	User      int `json:"user"`
	Order     int `json:"order"`
	Query     int `json:"query"`
	Source    int `json:"source"`
	SelfDDL   int `json:"self_ddl"`
	SelfDML   int `json:"self_dml"`
	SelfAudit int `json:"self_audit"`
	SelfQuery int `json:"self_query"`
}

func (Handler) DashBanner(c *gin.Context) {
	var bannerCount bannerCount
	bannerCount.User = 1
	bannerCount.Order = 1
	bannerCount.Query = 1
	bannerCount.Source = 1
	bannerCount.SelfDDL = 1
	bannerCount.SelfDML = 1
	bannerCount.SelfAudit = 1
	bannerCount.SelfQuery = 1
	exgin.GinsData(c, bannerCount, nil)
}

func (Handler) DashTop(c *gin.Context) {
	var source []groupBy
	exgin.GinsData(c, source, nil)
}
