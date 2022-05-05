// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package cache

import (
	"time"

	"github.com/ergoapi/zlog"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
)

var MemCache *cache.Cache

func Setup() {
	zlog.Log.Info("init cache", zap.String("plugin", "memcache"))
	MemCache = cache.New(5*time.Minute, 10*time.Minute) // cache 5min, clean 10min
}
