// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/cfichtmueller/storctl/cli"
	"github.com/cfichtmueller/storctl/conf"
	"github.com/spf13/cobra"
)

type createContextOptions struct {
	name   string
	server string
	apiKey string
}

func newCreateContextCommand(storCli *cli.Cli) *cobra.Command {
	options := createContextOptions{}

	cmd := &cobra.Command{
		Use:   "create-context",
		Short: "Create a new context",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runCreateContext(storCli, options); err != nil {
				storCli.Out.FailAndExit(err)
			}
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&options.name, "name", "n", "", "Name of the context")
	flags.StringVarP(&options.server, "server", "s", "", "URL of the server")
	flags.StringVarP(&options.apiKey, "api-key", "k", "", "API key")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("server")
	cmd.MarkFlagRequired("api-key")

	return cmd
}

func runCreateContext(storCli *cli.Cli, options createContextOptions) error {
	if err := storCli.Config.CreateContext(options.name, options.server, options.apiKey); err != nil {
		return err
	}
	if err := conf.Save(storCli.Config); err != nil {
		return err
	}
	return nil
}
