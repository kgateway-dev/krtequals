// Package krtequals provides a golangci-lint plugin for checking Equals() method implementations.
//
// This file exists to satisfy golangci-lint's module plugin system, which expects
// a package at the module root. The actual plugin registration happens in the
// pkg/analyzer package's init() function.
package krtequals

import (
	// Import the analyzer package to trigger plugin registration via init().
	_ "github.com/kgateway-dev/krtequals/pkg/analyzer"
)

