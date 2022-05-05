// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package version

import "github.com/spf13/cobra"

func Version() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "show version",
	}
}
