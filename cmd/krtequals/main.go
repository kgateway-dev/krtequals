// Command krtequals runs the krtequals analyzer as a standalone tool.
package main

import (
	"github.com/kgateway-dev/krtequals/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
