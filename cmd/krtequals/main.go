// Command krtequals runs the krtequals analyzer as a standalone tool.
package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/kgateway-dev/krtequals/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
