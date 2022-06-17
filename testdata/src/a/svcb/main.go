package svcb

import (
	"fmt"

	"a/svcc" // want `import of "a/svcc" prohibited`
)

func main() {
	fmt.Printf("random: %s\n", svcc.RandomString(4))
}
