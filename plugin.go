// Package krtequals is a golangci-lint plugin for checking Equals() method implementations.
// It is built as a module to be used with golangci-lint.
// See https://golangci-lint.run/plugins/module-plugins/ for more information.
//
// This pattern follows the same approach used by kube-api-linter (KAL):
// https://github.com/kubernetes-sigs/kube-api-linter/blob/main/plugin.go
package krtequals

import (
	// Importing the analyzer package triggers plugin registration via init().
	"github.com/kgateway-dev/krtequals/pkg/analyzer"
)

// New is the entrypoint for the plugin.
var New = analyzer.New
