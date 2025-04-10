// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/ergoapi/util/color"
	"github.com/ergoapi/util/zos"
	"github.com/ergoapi/zlog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ysicing/bigcat/cmd/agent"
	"github.com/ysicing/bigcat/cmd/core"
	"github.com/ysicing/bigcat/cmd/version"
)

var (
	rootCmd = &cobra.Command{
		Use:   "bigcat",
		Short: "bigcat by ysicing",
	}
	CfgFile string
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	logcfg := &zlog.Config{Simple: true, WriteLog: true, WriteJSON: false, ServiceName: "bigcat"}
	zlog.InitZlog(logcfg)
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "config file (default is /conf/bigcat.yml)")
	rootCmd.AddCommand(core.Core())
	rootCmd.AddCommand(agent.Agent())
	rootCmd.AddCommand(version.Version())
}

func initConfig() {
	if CfgFile == "" {
		CfgFile = "/conf/bigcat.yml"
		if zos.IsMacOS() {
			CfgFile = "./conf/bigcat.yml"
		}
	}
	viper.SetEnvPrefix("bc")
	viper.SetConfigFile(CfgFile)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		zlog.Debug("Using config file: %v", color.SGreen(viper.ConfigFileUsed()))
	}
	// reload
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zlog.Debug("config changed: %v", color.SGreen(in.Name))
	})
}
