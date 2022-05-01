#!/usr/bin/env bash
# Copyright (c) 2022 ysicing All rights reserved.
# Use of this source code is governed by AGPL-3.0-or-later
# license that can be found in the LICENSE file.


addlicense -f hack/tpl/agpl.tpl -ignore web/** -ignore "**/*.md" -ignore vendor/** -ignore "**/*.yml" -ignore "**/*.yaml" -ignore "**/*.sh" ./**