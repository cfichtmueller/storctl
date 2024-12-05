// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rm

import (
	"context"
	"fmt"

	"github.com/cfichtmueller/stor-go-client/stor"
	"github.com/cfichtmueller/storctl/cli"
	"github.com/spf13/cobra"
)

func NewCommand(storCli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rm BUCKET KEY...",
		Short:   "Delete object",
		Aliases: []string{"delete"},
		Args:    cobra.MatchAll(cobra.MinimumNArgs(2), cobra.MaximumNArgs(100)),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runRm(storCli, args[0], args[1:]); err != nil {
				storCli.Out.FailAndExit(err)
			}
		},
	}

	return cmd
}

func runRm(storCli *cli.Cli, bucket string, keys []string) error {
	refs := make([]stor.ObjectReference, len(keys))
	for i, key := range keys {
		refs[i] = stor.ObjectReference{Key: key}
	}
	result, err := storCli.Client.DeleteObjects(context.Background(), stor.DeleteObjectsCommand{
		Bucket:  bucket,
		Objects: refs,
	})
	if err != nil {
		return err
	}

	if len(result.Results) == 0 {
		return fmt.Errorf("didn't receive delete result from server")
	}

	for _, r := range result.Results {
		if !r.Deleted {
			storCli.Out.Errorf("didn't delete %s: %v\n", r.Key, r.Error.Code)
		} else {
			storCli.Out.Println("deleted", r.Key)
		}
	}

	return nil
}
