# Copyright (c) 2022 ysicing All rights reserved.
# Use of this source code is governed by AGPL-3.0-or-later
# license that can be found in the LICENSE file.

FROM ysicing/god as build

WORKDIR /go/src

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.mod

COPY go.sum go.sum

RUN go mod download

COPY . .

RUN make build

FROM ysicing/debian

COPY --from=build /go/src/_output/bin/bigcat /root/bigcat

COPY --from=build /go/src/web/dist /root/web/dist

RUN chmod +x /root/bigcat

CMD /root/bigcat core
