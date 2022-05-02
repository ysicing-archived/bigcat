// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package routes

import (
	"fmt"

	"github.com/ergoapi/exgin"
	"github.com/ergoapi/util/exnet"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/bigcat/internal/app/routes/handler"
	"github.com/ysicing/bigcat/internal/app/routes/handler/auth"
	"github.com/ysicing/bigcat/internal/app/routes/handler/dash"
	"github.com/ysicing/bigcat/internal/app/routes/handler/healthz"
	"github.com/ysicing/bigcat/internal/app/routes/handler/user"
	"github.com/ysicing/bigcat/internal/app/routes/handler/web"
	"github.com/ysicing/bigcat/pkg/util/config"
)

func SetupRoutes() *gin.Engine {
	g := exgin.Init(config.GetBool("debug"))
	g.Use(exgin.ExCors())
	g.Use(exgin.ExLog("/healthz", "/metrics"))
	g.Use(exgin.ExRecovery())
	pprof.Register(g, fmt.Sprintf("/hostdebug/%v/entry", exnet.LocalIPs()[0]))
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
