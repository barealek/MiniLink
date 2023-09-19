package utility

import (
	"math/rand"
	"regexp"
)

func GetRandomID() string {
	length := 8
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = byte(rand.Intn(26) + 65)
	}
	return string(bytes)
}

func SanitizeURL(url string) bool {
	pattern := `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
	matched, err := regexp.Match(pattern, []byte(url))
	if err != nil {
		return false
	}

	return matched
}
