package auth

import (
	"crypto/sha512"
	"encoding/base64"
)

//returns the sha512 of the input string
func mySha512(toHash string) string {
	hasher := sha512.New()
	hasher.Write([]byte(toHash))
	b64 := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return b64
}
