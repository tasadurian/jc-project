package helpers

import (
	"errors"
	"strconv"
	"strings"
)

func GetPasswordString(b []byte) (string, error) {
	str := strings.Split(string(b), "=")
	if len(str) != 2 {
		return "", errors.New("Error getting password")
	}
	return str[1], nil
}

// ParseURL takes in a URL path ex: /hash/42 and
// returns 42.
func ParseURL(url string) (int, error) {
	str := strings.Split(url[1:], "/")
	if len(str) != 2 {
		return 0, errors.New("Error getting key")
	}

	i, err := strconv.Atoi(str[1])
	if err != nil {
		return 0, err
	}

	return i, nil
}
