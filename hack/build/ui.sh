#!/usr/bin/env bash
# Copyright (c) 2022 ysicing All rights reserved.
# Use of this source code is governed by AGPL-3.0-or-later
# license that can be found in the LICENSE file.


version=${1}

[ -d "web" ] && (
pushd web
  cat package.json  | jq '.version="'${version}'"' > package.new.json
  mv package.new.json package.json
  [ ! -d "node_modules" ] && yarn install
  yarn build
popd
)