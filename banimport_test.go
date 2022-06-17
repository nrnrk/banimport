package banimport_test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/nrnrk/banimport"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	t.Parallel()
	for _, p := range []string{"a", "b", "c"} {
		testdataPath := analysistest.TestData()
		buf, err := ioutil.ReadFile(filepath.Join(testdataPath, "src", p, ".banimport.json"))
		if err != nil {
			t.Fatal(err)
		}
		if err := banimport.Analyzer.Flags.Set("config", string(buf)); err != nil {
			t.Fatal(err)
		}
		testdata := testutil.WithModules(t, testdataPath, nil)
		analysistest.Run(t, testdata, banimport.Analyzer, fmt.Sprintf("%s/...", p))
	}
}
