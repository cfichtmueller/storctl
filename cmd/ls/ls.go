// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ls

import (
	"context"
	"fmt"

	"github.com/cfichtmueller/stor-go-client/stor"
	"github.com/cfichtmueller/storctl/cli"
	"github.com/spf13/cobra"
)

type lsopts struct {
	quiet   bool
	summary bool
}

func NewCommand(cli *cli.Cli) *cobra.Command {

	opts := lsopts{}

	cmd := &cobra.Command{
		Use:     "ls BUCKET [PREFIX]",
		Short:   "List objects",
		Aliases: []string{"list"},
		Args: cobra.MatchAll(
			cobra.MinimumNArgs(1),
			cobra.MaximumNArgs(2),
		),
		Run: func(cmd *cobra.Command, args []string) {
			bucket := args[0]
			prefix := ""
			if len(args) > 1 {
				prefix = args[1]
			}
			if err := runLs(cli, opts, bucket, prefix); err != nil {
				cli.Out.FailAndExit(err)
			}
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opts.quiet, "quiet", "q", false, "only print object keys")
	flags.BoolVar(&opts.summary, "summary", false, "print summary row")

	return cmd
}

func runLs(storCli *cli.Cli, opts lsopts, bucket, prefix string) error {
	delimiter := ""
	if prefix != "" {
		delimiter = "/"
	}
	r, err := storCli.Client.ListObjects(context.Background(), stor.ListObjectsCommand{
		Bucket:    bucket,
		Prefix:    prefix,
		Delimiter: delimiter,
	})
	if err != nil {
		return err
	}

	if opts.quiet {
		for _, o := range r.Objects {
			storCli.Out.Println(o.Key)
		}
		return nil
	}

	count := 0
	var size int64 = 0

	w := storCli.Out.NewTabWriter()
	w.Writeln("KEY", "CONTENT-TYPE", "SIZE")
	for _, o := range r.Objects {
		count += 1
		size += o.Size
		w.Writeln(o.Key, o.ContentType, storCli.Formatter.FormatBytes(o.Size))
	}
	if opts.summary {
		w.Writeln(fmt.Sprintf("%d Objects", count), "", storCli.Formatter.FormatBytes(size))
	}
	w.Flush()
	return nil
}
