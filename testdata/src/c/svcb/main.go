package svcb

import (
	"fmt"

	"c/liba/inner"
	"c/svcc" // want `import of "c/svcc" prohibited`
)

func main() {
	fmt.Printf("random: %s\n", svcc.RandomString(4))
	p := 7
	fmt.Printf("%d is prime?: %t\n", p, inner.IsPrime(uint(p)))
}
