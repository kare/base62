package base62 // import "kkn.fi/base62"

import (
	"fmt"
	"strings"
)

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length   = int64(len(alphabet))
)

// Encode decoded integer to base62 string.
func Encode(n int64) string {
	if n == 0 {
		return "0"
	}

	s := ""
	for ; n > 0; n = n / length {
		s = string(alphabet[n%length]) + s
	}
	return s
}

// Decode a base62 encoded string to int.
// Returns an error if input s is not valid base62 literal [0-9a-zA-Z].
func Decode(s string) (int64, error) {
	var r int64
	for _, c := range []byte(s) {
		i := strings.IndexByte(alphabet, c)
		if i < 0 {
			return 0, fmt.Errorf("unexpected character %c in base62 literal", c)
		}
		r = length*r + int64(i)
	}
	return r, nil
}
