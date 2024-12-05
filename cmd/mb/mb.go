// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mb

import (
	"context"

	"github.com/cfichtmueller/stor-go-client/stor"
	"github.com/cfichtmueller/storctl/cli"
	"github.com/spf13/cobra"
)

func NewCommand(storCli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mb NAME",
		Short: "Create bucket",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			bucket := args[0]
			if err := runMb(storCli, bucket); err != nil {
				storCli.Out.FailAndExit(err)
			}
		},
	}

	return cmd
}

func runMb(storCli *cli.Cli, bucket string) error {
	r, err := storCli.Client.CreateBucket(context.Background(), stor.CreateBucketCommand{
		Name: bucket,
	})
	if err != nil {
		return err
	}

	storCli.Out.Println("Created", r.Name)
	return nil
}
