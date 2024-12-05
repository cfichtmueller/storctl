// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/cfichtmueller/stor-go-client/stor"
	"github.com/cfichtmueller/storctl/cli"
	"github.com/cfichtmueller/storctl/cmd/config"
	"github.com/cfichtmueller/storctl/cmd/cp"
	"github.com/cfichtmueller/storctl/cmd/lb"
	"github.com/cfichtmueller/storctl/cmd/ls"
	"github.com/cfichtmueller/storctl/cmd/mb"
	"github.com/cfichtmueller/storctl/cmd/mv"
	"github.com/cfichtmueller/storctl/cmd/rb"
	"github.com/cfichtmueller/storctl/cmd/rm"
	"github.com/cfichtmueller/storctl/conf"
	"github.com/spf13/cobra"
)

type rootOpts struct {
	human bool
}

func Execute() error {
	return newRootCommand().Execute()
}

func newRootCommand() *cobra.Command {
	storConfig, err := conf.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var client *stor.Client
	current := storConfig.GetCurrentContext()
	if current != nil {
		opts := stor.NewClientOptions().
			SetHost(current.Server).
			SetApiKey(current.ApiKey)

		client = stor.NewClient(opts)
	}
	storCli := cli.New(
		storConfig,
		client,
		cli.DefaultFormatter(),
		os.Stdout,
		os.Stderr,
	)

	opts := rootOpts{}

	cmd := &cobra.Command{
		Use:   "storctl",
		Short: "storctl is the CLI for the STOR object store",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if opts.human {
				storCli.Formatter = cli.HumanFormatter()
			}
		},
	}

	persistentFlags := cmd.PersistentFlags()
	persistentFlags.BoolVar(&opts.human, "human", false, "use unit suffixes: Byte, Kilobyte, Megabyte, Gigabyte and Terabyte")

	cmd.AddCommand(
		config.NewCommand(storCli),
		cp.NewCommand(storCli),
		lb.NewCommand(storCli),
		ls.NewCommand(storCli),
		mb.NewCommand(storCli),
		mv.NewCommand(storCli),
		rb.NewCommand(storCli),
		rm.NewCommand(storCli),
	)

	return cmd
}
