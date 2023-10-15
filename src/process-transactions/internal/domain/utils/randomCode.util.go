package utils

import (
	"crypto/rand"
	"math/big"
)

var GenerateRandomCode = func(n int) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	code := ""
	for i := 0; i < n; i++ {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))

		if err != nil {
			return ""
		}
		n2 := nBig.Int64()
		code += string(charset[int(n2)])
	}

	return code
}
