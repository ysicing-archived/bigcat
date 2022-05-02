// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package model

import (
	"time"

	"github.com/ergoapi/glog"
	"github.com/ergoapi/util/color"
	"github.com/ergoapi/zlog"
	"github.com/ysicing/bigcat/pkg/util/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
)

var db *gorm.DB
var migrates []interface{}

func migrate(obj interface{}) {
	migrates = append(migrates, obj)
}

func Init() {
	var err error
	dbtype := config.GetString("db.type")
	dbdsn := config.GetString("db.dsn")
	dbmode := config.GetBool("debug")
	newLogger := glog.New(zlog.Zlog, dbmode)
	switch dbtype {
	case "mysql":
		zlog.Debug("dbtype is mysql")
		db, err = gorm.Open(mysql.Open(dbdsn), &gorm.Config{
			Logger: newLogger,
		})
	default:
		zlog.Debug("dbtype is %s", dbtype)
		db, err = gorm.Open(sqlite.Open(dbdsn), &gorm.Config{
			Logger: newLogger,
		})
	}
	if err != nil {
		zlog.Panic("setup db err: %v", err.Error())
	}
	if config.GetBool("db.metrics.enable") {
		dbname := config.GetString("db.metrics.name", "bigcat")

		db.Use(prometheus.New(prometheus.Config{
			DBName: dbname,
		}))
	}
	dbcfg, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	dbcfg.SetMaxIdleConns(15)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	dbcfg.SetMaxOpenConns(50)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	dbcfg.SetConnMaxLifetime(time.Minute * 10)

	if err := db.AutoMigrate(migrates...); err != nil {
		zlog.Error("auto migrate table err: %v", err.Error())
	}
	zlog.Info(color.SGreen("create db engine success..."))
	InitSalt()
}

func DB() *gorm.DB {
	return db
}
