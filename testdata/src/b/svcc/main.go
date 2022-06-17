package svcc

import (
	"crypto/rand"
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("random: %s\n", RandomString(8))
}

func RandomString(length uint) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	var sb strings.Builder
	for _, v := range b {
		sb.WriteByte(letters[int(v)%len(letters)])
	}
	return sb.String()
}
