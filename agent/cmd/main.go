// Copyright 2024 The AuthPolicyController Authors.  All rights reserved.
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/auth-policy-controller/apc/agent/cmd/app"
)

func main() {
	if err := app.NewRootCommand().Execute(); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
