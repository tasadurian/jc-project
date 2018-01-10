package helpers

import "testing"

const criteria = "password=angryMonkey"

func TestGetPasswordString(t *testing.T) {
	b := []byte(criteria)
	pw, err := getPasswordString(b)
	if err != nil {
		t.Fatalf("err should have been nil, but was: %s", err.Error())
		t.Fail()
	}
	if pw != "angryMonkey" {
		t.Fatalf("pw should have been angryMonkey, but was: %s", pw)
		t.Fail()
	}
}
