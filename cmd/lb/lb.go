// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package lb

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cfichtmueller/stor-go-client/stor"
	"github.com/cfichtmueller/storctl/cli"
	"github.com/spf13/cobra"
)

type lbopts struct {
	quiet   bool
	summary bool
}

func NewCommand(storCli *cli.Cli) *cobra.Command {
	opts := lbopts{}

	cmd := &cobra.Command{
		Use:   "lb ",
		Short: "List buckets",
		Run: func(cmd *cobra.Command, args []string) {
			if err := runLb(storCli, opts); err != nil {
				storCli.Out.FailAndExit(err)
			}
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opts.quiet, "quiet", "q", false, "only print bucket names")
	flags.BoolVar(&opts.summary, "summary", false, "print summary row")

	return cmd
}

func runLb(storCli *cli.Cli, opts lbopts) error {
	r, err := storCli.Client.ListBuckets(context.Background(), stor.ListBucketsCommand{})
	if err != nil {
		return err
	}
	if opts.quiet {
		for _, b := range r.Buckets {
			storCli.Out.Println(b.Name)
		}
		return nil
	}
	var objects int64 = 0
	var size int64 = 0
	w := storCli.Out.NewTabWriter()
	w.Writeln("NAME", "OBJECTS", "SIZE")
	for _, b := range r.Buckets {
		w.Writeln(b.Name, strconv.FormatInt(int64(b.Objects), 10), storCli.Formatter.FormatBytes(b.Size))
		objects += b.Objects
		size += b.Size
	}
	if opts.summary {
		w.Writeln(fmt.Sprintf("%d Buckets", len(r.Buckets)), strconv.FormatInt(objects, 10), storCli.Formatter.FormatBytes(size))
	}
	w.Flush()
	return nil
}
