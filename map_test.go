package gomap

import (
	"fmt"
	"testing"

	"github.com/sanity-io/litter"
)

func TestPrint(t *testing.T) {
	m := New[string, int]()

	m.Set("one", 1)
	m.Set("two", 2)
	m.Set("three", 3)
	m.Set("two", -2)

	m.Range(func(key string, value int) bool {
		print("{", key, ":", value, "}, ")
		return true
	})

	fmt.Printf("\nActual: ")
	litter.Dump(m)
}
