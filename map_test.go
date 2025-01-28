package gomap

import (
	"testing"

	"github.com/sanity-io/litter"
)

func TestGoMap(t *testing.T) {
	m := New[string, int]()

	m.Set("one", 1)
	m.Set("two", 2)
	m.Set("three", 3)
	m.Set("two", -2)

	for key, value := range m.All() {
		t.Log("{", key, ":", value, "}, ")
	}
	t.Log("\nlitter.Dumpump:", litter.Sdump(m), "\n") //Dump(m)
}
