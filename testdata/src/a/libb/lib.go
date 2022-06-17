package libb

import (
	"a/liba"
)

func sum(vals ...int) int {
	result := 0
	for _, val := range vals {
		result = liba.Add(result, val)
	}
	return result
}
