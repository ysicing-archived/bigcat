// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package healthz

import (
	"fmt"

	"github.com/ergoapi/exgin"
	"github.com/ergoapi/util/version"

	// version prometheus
	_ "github.com/ergoapi/util/version/prometheus"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	// docs
	_ "github.com/ysicing/bigcat/docs"
	"github.com/ysicing/bigcat/internal/app/routes/handler"
)

type Handler struct{}

func NewHandler() (handler.RouteRegister, error) {
	return &Handler{}, nil
}

func (h *Handler) ApplyRoute(r *gin.Engine) {
	r.GET("/healthz", func(c *gin.Context) {
		exgin.GinsData(c, map[string]string{
			"healthz": "healthz",
		}, nil)
	})
	r.GET("/rv", func(c *gin.Context) {
		v := version.Get()
		exgin.GinsData(c, map[string]string{
			"builddate": v.BuildDate,
			"gitcommit": v.GitCommit,
			"version":   v.GitVersion,
		}, nil)
	})
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.NoMethod(func(c *gin.Context) {
		msg := fmt.Sprintf("not found: %v", c.Request.Method)
		exgin.GinsAbortWithCode(c, 404, msg)
	})
	r.NoRoute(func(c *gin.Context) {
		msg := fmt.Sprintf("not found: %v", c.Request.URL.Path)
		exgin.GinsAbortWithCode(c, 404, msg)
	})
}
