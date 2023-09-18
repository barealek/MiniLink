package utility

import (
	"math/rand"
)

func GetRandomID() string {
	length := 8
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = byte(rand.Intn(26) + 65)
	}
	return string(bytes)
}
