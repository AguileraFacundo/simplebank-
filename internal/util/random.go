package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmopqtuvxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func RandomNumber(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int64) string {
	var sb strings.Builder
	k := len(alphabet)

	for range n {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomCurrencies() string {
	currencies := []string{"USD", "EUR", "ARS"}
	n := len(currencies)
	return currencies[rand.Intn(n)]

}

func RandomMoney() int64 {
	return RandomNumber(10, 1000)
}

func RandomOwner() string {
	return RandomString(6)
}
