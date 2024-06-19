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

	it := m.Iterator()
	fmt.Printf("Result: %T{", m)
	for it.Next() {
		fmt.Printf("\"%+v\": %+v, ", it.Key(), it.Value())
	}
	fmt.Println("}")

	fmt.Print("Actual: ")
	litter.Dump(m)
}
