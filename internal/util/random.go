package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

/*func init() {
	rand.Seed(time.Now().UnixNano())
}*/

// RandomInt generates a random integer between min and max
func RandomInt(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		panic(err)
	}
	return nBig.Int64()
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := int64(len(alphabet))

	for i := 0; i < n; i++ {
		nBig, err := rand.Int(rand.Reader, big.NewInt(k))
		if err != nil {
			panic(err)
		}

		ind := nBig.Int64()
		c := alphabet[ind]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(1000)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
