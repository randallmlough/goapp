package rand

import (
	"math/rand"
	"strings"
	"time"
)

const (
	azLower = "abcdefghijklmnopqrstuvwxyz"
	azUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits  = "123456789"
	charset = azLower + azUpper + digits
)

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandomString(length int) string {
	sb := strings.Builder{}
	sb.Grow(length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, seededRand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = seededRand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(charset) {
			sb.WriteByte(charset[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return sb.String()
}
