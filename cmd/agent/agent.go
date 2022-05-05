// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package agent

import "github.com/spf13/cobra"

func Agent() *cobra.Command {
	return &cobra.Command{
		Use:   "agent",
		Short: "agent by ysicing",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO
		},
	}
}
