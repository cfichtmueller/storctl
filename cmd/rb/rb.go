// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rb

import (
	"context"

	"github.com/cfichtmueller/stor-go-client/stor"
	"github.com/cfichtmueller/storctl/cli"
	"github.com/spf13/cobra"
)

func NewCommand(storCli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rb NAME",
		Short: "Delete bucket",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			bucket := args[0]
			if err := runRm(storCli, bucket); err != nil {
				storCli.Out.FailAndExit(err)
			}
		},
	}

	return cmd
}

func runRm(storCli *cli.Cli, bucket string) error {
	if err := storCli.Client.DeleteBucket(context.Background(), stor.DeleteBucketCommand{
		Name: bucket,
	}); err != nil {
		return err
	}

	storCli.Out.Println("Deleted", bucket)

	return nil
}
