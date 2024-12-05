// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mv

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/cfichtmueller/stor-go-client/stor"
	"github.com/cfichtmueller/storctl/cli"
	"github.com/spf13/cobra"
)

type mvopts struct {
	contentType string
}

func NewCommand(storCli *cli.Cli) *cobra.Command {
	opts := mvopts{}
	cmd := &cobra.Command{
		Use:     "mv SOURCE TARGET",
		Short:   "Move object",
		Aliases: []string{"move"},
		Args:    cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			from, to := args[0], args[1]
			if err := runMv(storCli, opts, from, to); err != nil {
				storCli.Out.FailAndExit(err)
			}
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&opts.contentType, "content-type", "application/octet-stream", "the content type to use when moving to STOR")
	return cmd
}

func runMv(storCli *cli.Cli, opts mvopts, from, to string) error {
	fromStor := storCli.In.IsStorUri(from)
	toStor := storCli.In.IsStorUri(to)
	if !fromStor && !toStor {
		return fmt.Errorf("either source or target needs to be a STOR location")
	}
	if !fromStor {
		return runMvToStor(storCli, opts, from, to)
	}
	if !toStor {
		return runMvFromStor(storCli, from, to)
	}
	return runMvStorStor(storCli, from, to)
}

func runMvToStor(storCli *cli.Cli, opts mvopts, from, to string) error {
	bucket, key, _ := storCli.In.ParseStorUri(to)
	f, err := os.Open(from)
	if err != nil {
		return err
	}
	if _, err := storCli.Client.CreateObject(context.Background(), stor.CreateObjectCommand{
		Bucket:      bucket,
		Key:         key,
		ContentType: opts.contentType,
		Data:        f,
	}); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	if err := os.Remove(from); err != nil {
		return err
	}

	return nil
}

func runMvFromStor(storCli *cli.Cli, from, to string) error {
	ctx := context.Background()
	bucket, key, _ := storCli.In.ParseStorUri(from)
	wr, err := storCli.Client.ReadObject(ctx, bucket, key)
	if err != nil {
		return err
	}
	f, err := os.Create(to)
	defer f.Close()
	if err != nil {
		return err
	}

	if _, err := io.Copy(f, wr); err != nil {
		return err
	}

	dr, err := storCli.Client.DeleteObjects(ctx, stor.DeleteObjectsCommand{
		Bucket: bucket,
		Objects: []stor.ObjectReference{
			{Key: key},
		},
	})
	if err != nil {
		return err
	}

	for _, drr := range dr.Results {
		if !drr.Deleted {
			return fmt.Errorf("failed to delete %s/%s: %s", bucket, key, drr.Error.Code)
		}
	}

	return nil
}

func runMvStorStor(storCli *cli.Cli, from, to string) error {
	ctx := context.Background()
	fromBucket, fromKey, _ := storCli.In.ParseStorUri(from)
	toBucket, toKey, _ := storCli.In.ParseStorUri(to)

	if fromBucket != toBucket {
		return fmt.Errorf("cannot move between buckets")
	}

	if _, err := storCli.Client.CopyObject(ctx, stor.CopyObjectCommand{
		Bucket:    fromBucket,
		SourceKey: fromKey,
		DestKey:   toKey,
	}); err != nil {
		return err
	}

	dr, err := storCli.Client.DeleteObjects(ctx, stor.DeleteObjectsCommand{
		Bucket: fromBucket,
		Objects: []stor.ObjectReference{
			{Key: fromKey},
		},
	})
	if err != nil {
		return err
	}

	for _, drr := range dr.Results {
		if !drr.Deleted {
			return fmt.Errorf("failed to delete %s/%s: %s", fromBucket, fromKey, drr.Error.Code)
		}
	}

	return nil
}
