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

func CreateRandomNumber(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func createRandomString(n int64) string {
	var sb strings.Builder
	k := len(alphabet)

	for range n {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func CreateRandomCurrencies() string {
	currencies := []string{"USD", "EUR", "ARS"}
	n := len(currencies)
	return currencies[rand.Intn(n)]

}

func CreateRandomMoney() int64 {
	return CreateRandomNumber(10, 1000)
}

func CreateRandomOwner() string {
	return createRandomString(6)
}
