// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cli

import (
	"io"

	"github.com/cfichtmueller/stor-go-client/stor"
	"github.com/cfichtmueller/storctl/conf"
)

type Cli struct {
	Config    *conf.Config
	Client    *stor.Client
	Formatter Formatter
	In        *In
	Out       *Out
}

func New(
	config *conf.Config,
	client *stor.Client,
	formatter Formatter,
	out, err io.Writer,
) *Cli {
	return &Cli{
		Config:    config,
		Client:    client,
		Formatter: formatter,
		In:        &In{},
		Out: &Out{
			out: out,
			err: err,
		},
	}
}
