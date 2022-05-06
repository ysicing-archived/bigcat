// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package rpc

import (
	"context"

	"github.com/ergoapi/util/zos"
	pb "github.com/ysicing/bigcat/proto"
)

type GreeterServer struct {
}

func (h *GreeterServer) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Response: zos.GetHostname()}, nil
}
