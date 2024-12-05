// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cli

import (
	"fmt"
	"strconv"
)

type Formatter interface {
	FormatBytes(bytes int64) string
}

type defaultFormatter struct{}

func DefaultFormatter() Formatter {
	return &defaultFormatter{}
}

func (f *defaultFormatter) FormatBytes(bytes int64) string {
	return strconv.FormatInt(bytes, 10)
}

type humanFormatter struct{}

func HumanFormatter() Formatter {
	return &humanFormatter{}
}

func (f *humanFormatter) FormatBytes(bytes int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
		TB = GB * 1024
	)

	switch {
	case bytes >= TB:
		return fmt.Sprintf("%.2f TB", float64(bytes)/float64(TB))
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}
