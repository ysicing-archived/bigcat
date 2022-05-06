// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/ergoapi/zlog"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/bigcat/internal/dash/model"
	"github.com/ysicing/bigcat/internal/dash/routes"
	rpcService "github.com/ysicing/bigcat/internal/dash/rpc"
	"github.com/ysicing/bigcat/pkg/cache"
	"github.com/ysicing/bigcat/pkg/cron"
	"github.com/ysicing/bigcat/pkg/util/config"
	pb "github.com/ysicing/bigcat/proto"
	"google.golang.org/grpc"
)

func Serve(ctx context.Context) error {
	defer cron.Cron.Stop()
	cron.Cron.Start()
	cache.Setup()
	model.Init()
	g := routes.SetupRoutes()
	if config.GetBool("server.ssl.enable") {
		go startTLS(ctx, g)
	}
	if config.GetBool("server.rpc.enable") {
		go startRPC(ctx)
	}
	addr := ":" + config.GetString("server.listen", "65001")
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

func startTLS(ctx context.Context, e *gin.Engine) {
	tlsaddr := ":" + config.GetString("server.ssl.listen", "65002")
	srv := &http.Server{
		Addr:    tlsaddr,
		Handler: e,
	}
	tlscfile := config.GetString("server.ssl.cert")
	tlskfile := config.GetString("server.ssl.key")
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			zlog.Error("failed to stop tls server, error: %s", err)
		}
		zlog.Info("tls server exited.")
	}()
	zlog.Info("tls listen to %v, pid is %v", tlsaddr, os.Getpid())
	if err := srv.ListenAndServeTLS(tlscfile, tlskfile); err != nil && err != http.ErrServerClosed {
		zlog.Fatal("failed to start tls server, error: %s", err)
	}
}

func startRPC(ctx context.Context) {
	addr := ":" + config.GetString("server.rpc.listen", "65003")
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &rpcService.GreeterServer{})
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		zlog.Fatal("start rpc err: %v", err)
	}
	zlog.Info("rpc listen to %v, pid is %v", addr, os.Getpid())
	server.Serve(listen)
}
