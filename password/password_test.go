package password

import "testing"

const pw = "angryMonkey"

func TestEncodeAndHash(t *testing.T) {
	encodedPW := EncodeAndHash(pw)
	if encodedPW != "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==" {
		t.Fatalf("encodedPW test fail - encodedPW was %s", encodedPW)
		t.Fail()
	}
}

func TestEncode(t *testing.T) {
	encodedPW := EncodeAndHash(pw)
	if encodedPW == "" {
		t.Fatalf("encodedPW  was empty")
		t.Fail()
	}
}

func TestHash(t *testing.T) {
	hashedPW := EncodeAndHash(pw)
	if hashedPW == "" {
		t.Fatalf("hashedPW  was empty")
		t.Fail()
	}
}
