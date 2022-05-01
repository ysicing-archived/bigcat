// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package handler

import "github.com/gin-gonic/gin"

type RegisterFactory func() (RouteRegister, error)

type RouteRegister interface {
	ApplyRoute(r *gin.Engine)
}
