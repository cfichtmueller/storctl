// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/cfichtmueller/storctl/cli"
	"github.com/spf13/cobra"
)

func newSetCredentialsCommand(storCli *cli.Cli) *cobra.Command {
	//TODO: implement
	cmd := &cobra.Command{
		Use:   "set-credentials API_KEY",
		Short: "Set the credentials for the current context",
		Run: func(cmd *cobra.Command, args []string) {
			storCli.Out.FailAndExitf("Not implemented yet")
		},
	}

	return cmd
}
