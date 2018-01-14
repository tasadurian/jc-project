package helpers

import "testing"

const criteria = "password=angryMonkey"

func TestGetPasswordString(t *testing.T) {
	b := []byte(criteria)
	pw, err := GetPasswordString(b)
	if err != nil {
		t.Fatalf("err should have been nil, but was: %s", err.Error())
		t.Fail()
	}
	if pw != "angryMonkey" {
		t.Fatalf("pw should have been angryMonkey, but was: %s", pw)
		t.Fail()
	}
}

func TestParseURL(t *testing.T) {
	url := "/hash/42"
	key, err := ParseURL(url)
	if err != nil {
		t.Fatalf("err should have been nil, but was: %s", err.Error())
		t.Fail()
	}
	if key != 42 {
		t.Fatalf("key should have been 42, but was: %d", key)
		t.Fail()
	}
}
