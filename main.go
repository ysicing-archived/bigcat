// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package main

import (
	"runtime"

	"github.com/ergoapi/errors"
	"github.com/ysicing/bigcat/cmd"
)

// @title BigCat API
// @version 0.0.1
// @description This is a sample server Petstore server.

// @contact.name ysicing
// @contact.url http://github.com/ysicing
// @contact.email i@ysicing.me

// @license.name AGPL
// @license.url https://www.gnu.org/licenses/agpl-3.0.en.html

func main() {
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)
	errors.CheckAndExit(cmd.Execute())
}
