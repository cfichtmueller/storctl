// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/cfichtmueller/storctl/cli"
	"github.com/cfichtmueller/storctl/conf"
	"github.com/spf13/cobra"
)

func newDeleteContextCommand(storCli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-context NAME",
		Short: "Remove a context from the configuration",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			if err := runDeleteContext(storCli, name); err != nil {
				storCli.Out.FailAndExit(err)
			}
		},
	}

	return cmd
}

func runDeleteContext(storCli *cli.Cli, name string) error {
	if err := storCli.Config.DeleteContext(name); err != nil {
		return err
	}
	if err := conf.Save(storCli.Config); err != nil {
		return err
	}
	return nil
}
