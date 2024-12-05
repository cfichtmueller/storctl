// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cli

import (
	"strings"
	"text/tabwriter"
)

type TabWriter struct {
	tw *tabwriter.Writer
}

func (w *TabWriter) Writeln(a ...string) {
	if _, err := w.tw.Write([]byte(strings.Join(a, "\t") + "\n")); err != nil {
		panic(err)
	}
}

func (w *TabWriter) Flush() {
	if err := w.tw.Flush(); err != nil {
		panic(err)
	}
}
