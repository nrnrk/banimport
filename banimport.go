package banimport

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const (
	doc            = "banimport is a tool to check import dependencies"
	configFileName = `.banimport.yaml`
)

var (
	Analyzer = &analysis.Analyzer{
		Name: "banimport",
		Doc:  doc,
		Run:  run,
	}
	configJSON = `{"pattern":{}}`
)

func init() {
	Analyzer.Flags.StringVar(&configJSON, "config", configJSON, "config json")
}

func run(pass *analysis.Pass) (interface{}, error) {
	var config *Config
	if err := json.Unmarshal([]byte(configJSON), &config); err != nil {
		return nil, fmt.Errorf("JSON parse error: %w", err)
	}
	banList := config.getBanList(pass.Pkg.Path())
	if len(banList) == 0 {
		return nil, nil
	}
	for _, f := range pass.Files {
		for _, imported := range f.Imports {
			pathStr, err := strconv.Unquote(imported.Path.Value)
			if err != nil {
				return nil, fmt.Errorf("unexpected unquote error: %w", err)
			}
			if banList.ban(pathStr) {
				pass.Reportf(imported.Pos(), "import of %s prohibited", imported.Path.Value)
			}
		}
	}
	return nil, nil
}

type (
	BanList    []string
	BanPattern map[string]BanList
	Config     struct {
		BanPattern BanPattern `json:"pattern"`
	}
)

func (c *Config) getBanList(from string) BanList {
	list := make([]string, 0)
	for source, dists := range c.BanPattern {
		if strings.HasPrefix(from, source) {
			list = append(list, dists...)
		}
	}
	return list
}

func (bl BanList) ban(target string) bool {
	for _, v := range bl {
		if strings.HasPrefix(target, v) {
			return true
		}
	}
	return false
}
