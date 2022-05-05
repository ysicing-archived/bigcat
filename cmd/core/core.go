// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package core

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/ergoapi/zlog"
	"github.com/spf13/cobra"
	"github.com/ysicing/bigcat/internal/app/server"
)

func Core() *cobra.Command {
	return &cobra.Command{
		Use:   "core",
		Short: "core by ysicing",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
			go func() {
				<-ctx.Done()
				stop()
			}()

			if err := server.Serve(ctx); err != nil {
				zlog.Fatal("run serve: %v", err)
			}
		},
	}
}
