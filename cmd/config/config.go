// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/cfichtmueller/storctl/cli"
	"github.com/spf13/cobra"
)

func NewCommand(storCli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Interact with the storectl config",
	}

	cmd.AddCommand(
		newCreateContextCommand(storCli),
		newDeleteContextCommand(storCli),
		newGetContextsCommand(storCli),
		newRenameContextCommand(storCli),
		newSetCredentialsCommand(storCli),
		newUseContextCommand(storCli),
	)

	return cmd
}
