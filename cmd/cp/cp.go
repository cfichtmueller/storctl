// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cp

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/cfichtmueller/stor-go-client/stor"
	"github.com/cfichtmueller/storctl/cli"
	"github.com/spf13/cobra"
)

type cpopts struct {
	contentType string
}

func NewCommand(storCli *cli.Cli) *cobra.Command {
	opts := cpopts{}

	cmd := &cobra.Command{
		Use:     "cp SOURCE TARGET",
		Short:   "Copy object",
		Aliases: []string{"copy"},
		Args:    cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			from, to := args[0], args[1]
			if err := runCp(storCli, opts, from, to); err != nil {
				storCli.Out.FailAndExit(err)
			}
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&opts.contentType, "content-type", "application/octet-stream", "the content type to use when copying to STOR")
	return cmd
}

func runCp(storCli *cli.Cli, opts cpopts, from, to string) error {
	fromStor := storCli.In.IsStorUri(from)
	toStor := storCli.In.IsStorUri(to)
	if !fromStor && !toStor {
		return fmt.Errorf("either source or target needs to be a STOR location")
	}
	if !fromStor {
		return runCpToStor(storCli, opts, from, to)
	}
	if !toStor {
		return runCpFromStor(storCli, from, to)
	}
	return runCpStorStor(storCli, from, to)
}

func runCpToStor(storCli *cli.Cli, opts cpopts, from, to string) error {
	bucket, key, _ := storCli.In.ParseStorUri(to)
	f, err := os.Open(from)
	defer f.Close()
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

	return nil
}

func runCpFromStor(storCli *cli.Cli, from, to string) error {
	bucket, key, _ := storCli.In.ParseStorUri(from)
	r, err := storCli.Client.ReadObject(context.Background(), bucket, key)
	if err != nil {
		return err
	}
	f, err := os.Create(to)
	defer f.Close()
	if err != nil {
		return err
	}

	if _, err := io.Copy(f, r); err != nil {
		return err
	}

	return nil
}

func runCpStorStor(storCli *cli.Cli, from, to string) error {
	fromBucket, fromKey, _ := storCli.In.ParseStorUri(from)
	toBucket, toKey, _ := storCli.In.ParseStorUri(to)

	if fromBucket != toBucket {
		return fmt.Errorf("cannot copy between buckets")
	}

	if _, err := storCli.Client.CopyObject(context.Background(), stor.CopyObjectCommand{
		Bucket:    fromBucket,
		SourceKey: fromKey,
		DestKey:   toKey,
	}); err != nil {
		return err
	}

	return nil
}
