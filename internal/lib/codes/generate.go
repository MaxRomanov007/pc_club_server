package codes

import (
	"math/rand/v2"
	"strings"
)

func Generate(length int) string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	sb := strings.Builder{}
	for i := 0; i < length; i++ {
		sb.WriteByte(letters[rand.N[int](36)])
	}
	return sb.String()
}
