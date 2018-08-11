package auth

import (
	cryptoRand "crypto/rand"
	"encoding/base64"
	"io"
	"log"
	mathRand "math/rand"
	"time"
)

//return a random string. It is not cryptographicaly secure unless randomlevel
//is set to 1 in the ocnfig file
func getRandomString(size int) string {
	switch Settings.RandomLevel {
	case 0:
		mathRand.Seed(time.Now().UnixNano())

		b := make([]byte, size)
		for i := 0; i < size; i++ {
			b[i] = byte(mathRand.Int())
		}
		return base64.StdEncoding.EncodeToString((b))
	case 1:
		b := make([]byte, size)
		_, err := io.ReadFull(cryptoRand.Reader, b[:])
		if err != nil {
			log.Fatal(err)
		}
		return base64.StdEncoding.EncodeToString((b))
	}

	return ""
}
