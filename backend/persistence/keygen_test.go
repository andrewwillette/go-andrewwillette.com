package persistence

import (
	"testing"
)

func TestShaGeneration(t *testing.T) {
	var shaMap = make(map[string]int)
	for i := 0; i <= 100; i++ {
		sha1A := newSHA1Hash(2)
		_, exists := shaMap[sha1A]
		if exists {
			t.Log("sha already created")
			t.Fail()
		}
		shaMap[sha1A] = 1
	}
}
