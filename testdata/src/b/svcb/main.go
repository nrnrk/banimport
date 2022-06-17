package svcb

import (
	"fmt"

	"b/svcc"
)

func main() {
	fmt.Printf("random: %s\n", svcc.RandomString(4))
}
