// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"os"

	"github.com/cfichtmueller/storctl/cli"
	"github.com/cfichtmueller/storctl/conf"
	"github.com/spf13/cobra"
)

func newUseContextCommand(storCli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "use-context NAME",
		Short:   "Set the context to use",
		Aliases: []string{"use"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runUseContext(storCli, args[0]); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}

	return cmd
}

func runUseContext(storCli *cli.Cli, name string) error {
	if err := storCli.Config.SetCurrentContext(name); err != nil {
		return err
	}
	if err := conf.Save(storCli.Config); err != nil {
		return err
	}
	return nil
}
