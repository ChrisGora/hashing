package hashing

import (
	"fmt"
	"testing"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Test1(t *testing.T) {
	m := NewBasicHashMap()
	err := m.Put("some key", 123)
	check(err)
	err = m.Put("some key2", 12)
	check(err)
	err = m.Put("some key3", 1)
	check(err)
	err = m.Put("some key4", 1235)
	check(err)
	err = m.Put("some key5", 67)
	check(err)

	var val int
	err = m.Get("some key", &val)
	check(err)
	fmt.Println(val)
}
