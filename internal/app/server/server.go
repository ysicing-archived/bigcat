// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/ergoapi/zlog"
	"github.com/ysicing/bigcat/internal/app/model"
	"github.com/ysicing/bigcat/internal/app/routes"
	"github.com/ysicing/bigcat/pkg/cache"
	"github.com/ysicing/bigcat/pkg/cron"
)

func Serve(ctx context.Context) error {
	defer cron.Cron.Stop()
	cron.Cron.Start()
	cache.Setup()
	model.Init()
	g := routes.SetupRoutes()

	addr := "0.0.0.0:65001"
	srv := &http.Server{
		Addr:    addr,
		Handler: g,
	}
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			zlog.Error("Failed to stop server, error: %s", err)
		}
		zlog.Info("server exited.")
	}()
	zlog.Info("http listen to %v, pid is %v", addr, os.Getpid())
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zlog.Error("Failed to start http server, error: %s", err)
		return err
	}
	return nil
}
