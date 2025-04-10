// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/ergoapi/exgin"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/bigcat/internal/dash/routes/handler"
	"github.com/ysicing/bigcat/internal/dash/routes/handler/auth"
	"github.com/ysicing/bigcat/internal/dash/routes/handler/dash"
	"github.com/ysicing/bigcat/internal/dash/routes/handler/healthz"
	"github.com/ysicing/bigcat/internal/dash/routes/handler/user"
	"github.com/ysicing/bigcat/internal/dash/routes/handler/web"
	"github.com/ysicing/bigcat/pkg/util/config"
)

func SetupRoutes() *gin.Engine {
	g := exgin.Init(&exgin.Config{
		Debug:   config.GetBool("debug", true),
		Metrics: true,
		Cors:    true,
		Gops:    config.GetBool("server.gops.enable", true),
		Pprof:   config.GetBool("server.pprof.enable", true),
	})
	g.Use(exgin.ExZLog("/healthz", "/metrics", "/docs", "/hostdebug"))
	g.Use(exgin.ExZRecovery())
	factories := []handler.RegisterFactory{
		healthz.NewHandler,
		web.NewHandler,
		auth.NewHandler,
		dash.NewHandler,
		user.NewHandler,
	}
	for i := range factories {
		h, err := factories[i]()
		if err != nil {
			panic(err)
		}
		h.ApplyRoute(g)
	}
	return g
}
