package password

import (
	"crypto/sha512"
	"encoding/base64"
)

// EncodeAndHash ...
func EncodeAndHash(password string) string {
	return encode(hash(password))
}

func encode(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password))
}

func hash(password string) string {
	s512 := sha512.New()
	s512.Write([]byte(password))
	return string(s512.Sum(nil))
}
