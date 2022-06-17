package inner

import (
	"fmt"

	innera "c/liba/inner"
)

func Exec(n uint) {
	if innera.IsPrime(n) {
		fmt.Println("prime number")
	}
}
