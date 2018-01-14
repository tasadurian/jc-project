package store

import "testing"

func TestStore(t *testing.T) {
	db := NewStore()
	db.Put(5, "world")

	val, err := db.Get(5)
	if err != nil {
		t.Fatalf("err should be nil, but was %s", err.Error())
		t.Fail()
	}
	if val != "world" {
		t.Fatalf("val should be world, but was %s", val)
		t.Fail()
	}
}
