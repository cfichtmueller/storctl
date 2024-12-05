// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/cfichtmueller/storctl/cli"
	"github.com/spf13/cobra"
)

func newGetContextsCommand(storCli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-contexts",
		Short: "List available contexts",
		Run: func(cmd *cobra.Command, args []string) {
			if err := runGetContexts(storCli); err != nil {
				storCli.Out.FailAndExit(err)
			}
		},
	}

	return cmd
}

func runGetContexts(storCli *cli.Cli) error {
	contexts := storCli.Config.Contexts
	if len(contexts) == 0 {
		storCli.Out.Println("No Contexts")
		return nil
	}
	w := storCli.Out.NewTabWriter()
	w.Writeln("NAME", "SERVER")
	for _, c := range contexts {
		w.Writeln(c.Name, c.Server)
	}
	w.Flush()
	return nil
}
