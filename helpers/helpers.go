package helpers

import (
	"errors"
	"strings"
)

func getPasswordString(b []byte) (string, error) {
	str := strings.Split(string(b), "=")
	if len(str) != 2 {
		return "", errors.New("Error getting password")
	}
	return str[1], nil
}
