package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/nrnrk/banimport"
)

func main() {
	unitchecker.Main(banimport.Analyzer)
}
