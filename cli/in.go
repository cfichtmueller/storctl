// Copyright 2024 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cli

import (
	"fmt"
	"regexp"
)

var (
	storUriPattern = regexp.MustCompile("stor://([^/]+)/(.+)")
)

type In struct{}

// IsStorUri determines if the given string is a STOR uri.
func (i *In) IsStorUri(v string) bool {
	return storUriPattern.MatchString(v)
}

// ParseStorUri tries to parse the given string.
// If IsStorUri returns true for the string, then ParseStorUri will also succeed on the same input.
func (i *In) ParseStorUri(v string) (string, string, error) {
	matches := storUriPattern.FindStringSubmatch(v)
	if matches == nil {
		return "", "", fmt.Errorf("invalid stor URI")
	}
	return matches[1], matches[2], nil
}
